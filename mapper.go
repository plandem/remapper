package remapper

import (
    "fmt"
    "errors"
    "reflect"
)

type Mapper struct {
    types [2]*mapperType
}

func (m *Mapper) setType(tm *mapperType)(error) {
    var err error

    if m.types[0] == nil {
        m.types[0] = tm
    } else if m.types[1] == nil {
        m.types[1] = tm
    } else {
        err = errors.New("You can't set any type mapper anymore.")
    }

    return err
}

func (m *Mapper) getType(target interface{}) (*mapperType, error) {
    if targetType, err := resolveType(target, m.types[0].normalizedType.Kind(), m.types[1].normalizedType.Kind()); err != nil {
        return nil, err
    } else if targetType == m.types[0].normalizedType {
        return m.types[0], nil
    } else if targetType == m.types[1].normalizedType {
        return m.types[1], nil
    } else {
        return nil, errors.New(fmt.Sprintf("Type mismatch. Expected '%s' or '%s', but got '%s'", m.types[0].normalizedType, m.types[1].normalizedType, targetType))
    }
}

// Set value at target object for field with fieldName or return error if field was not mapped or value could not be converted
func (m *Mapper) SetByName(target interface{}, fieldName string, value interface{}) (error) {
    if targetType, err := m.getType(target); err == nil {
        fieldName := NameMapper(fieldName)

        if field, ok := targetType.fields[fieldName]; !ok {
            return unknownFieldName(fieldName)
        } else {
            value := reflect.ValueOf(value)

            //check if target is pointer to actual value
            target := reflect.ValueOf(target)
            if target.Kind() != reflect.Ptr {
                return errors.New("You can't directly mutate the 'target'. Use a pointer to target.")
            }

            target = reflect.Indirect(target)
            targetType.set(target, field.id, fieldName, value)
            return nil
        }
    } else {
        return err
    }
}

// Get value from target object for field with fieldName or return error if field was not mapped
func (m *Mapper) GetByName(target interface{}, fieldName string) (interface{}, error) {
    if targetType, err := m.getType(target); err == nil {
        fieldName := NameMapper(fieldName)
        if field, ok := targetType.fields[fieldName]; !ok {
            return nil, unknownFieldName(fieldName)
        } else {
            return targetType.get(reflect.ValueOf(target), field.id, fieldName).Interface(), nil
        }
    } else {
        return nil, err
    }
}

// Return reverse name from object for field with fieldName or return error if field was not mapped
func (m *Mapper) NameByName(from interface{}, fieldName string) (string, error) {
    if fromType, err := m.getType(from); err == nil {
        if field, ok := fromType.fields[NameMapper(fieldName)]; ok {
            if field.reverseId > 0 && len(field.reverseName) > 0 {
                return field.reverseName, nil
            }
        }
    }

    return "", unknownFieldName(fieldName)
}

// Map from object to reverse object or return error if mapping was failed
func (m *Mapper) Map(from interface{}) (interface{}, error) {
    if fromType, err := m.getType(from); err != nil {
        return nil, err
    } else {
        var toType *mapperType

        isToEmpty := true
        fromVal := reflect.Indirect(reflect.ValueOf(from))
        if fromType == m.types[0] {
            toType = m.types[1]
        } else {
            toType = m.types[0]
        }

        toVal, err := toType.create()
        if err != nil {
            return nil, err
        }

        for fieldName, field := range toType.fields {
            if field.reverseId < 0 || field.omit {
                continue
            }

            fromFieldVal := fromType.get(fromVal, field.reverseId, field.reverseName)
            if !fromFieldVal.IsValid() {
                continue
            }

            toFieldVal := toType.get(toVal, field.id, fieldName)

            if val, err := field.convert(fromFieldVal, toFieldVal.Type()); err != nil {
                return nil, errors.New(fmt.Sprintf("Could not convert '%s'. %s", fieldName, err.Error()))
            } else {
                if val.IsValid() {
                    isToEmpty = false
                    toFieldVal.Set(val)
                }
            }
        }

        if isToEmpty {
            return nil, nil
        }

        if toType.dataType.Kind() == reflect.Ptr {
            return toVal.Addr().Interface(), nil
        } else {
            return toVal.Interface(), nil
        }
    }
}

