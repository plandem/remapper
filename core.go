package remapper

import (
    "reflect"
    "errors"
    "fmt"
    "strings"
    "strconv"
)

var (
    // TagName is the name of the tag to use on struct fields
    TagName string = "remapper"

    // NameMapper is the function used to convert name of fields. The default mapper converts field names to lower case.
    NameMapper func(string) string = strings.ToLower
)

type mapperTypeI interface {
    create() (reflect.Value, error)                                //Create a new data of required type
    set(to reflect.Value, i int, name string, value reflect.Value) //Set value to
    get(from reflect.Value, i int, name string) (reflect.Value)    //Get value from
}

type mapperType struct {
    mapperTypeI
    dataType       reflect.Type            //Holds original type of data
    normalizedType reflect.Type            //Holds normalized type of data - no ptr and etc
    fields         map[string]*mapperField //Holds info for fields what must be mapped
}

func resolveType(v interface{}, types ...reflect.Kind) (reflect.Type, error) {
    var protoType reflect.Type

    if v != nil {
        protoType = reflect.TypeOf(v)

        if protoType.Kind() == reflect.Ptr {
            protoType = reflect.Indirect(reflect.ValueOf(v)).Type()
        }

        for _, kind := range types {
            if protoType.Kind() == kind {
                return protoType, nil
            }
        }
    }

    return nil, errors.New(fmt.Sprintf("Invalid type '%+v'. It must be a kind of: %+v", protoType.Kind(), types))
}

func resolveMappingField(fromType *mapperType, from string, toType *mapperType, to interface{}) (error) {
    fromFieldName := NameMapper(from)

    //is fromFieldName valid?
    if fromField, ok := fromType.fields[fromFieldName]; !ok {
        return errors.New(fmt.Sprintf("There is no field with name '%s' at %v", fromFieldName, fromType.normalizedType))
    } else {
        //mapping by index?
        if toFieldId, ok := to.(int); ok {
            fromField.reverseId = toFieldId
            return nil
        }

        //mapping by name or index with settings?
        toFieldMappingSettings, ok := to.(string)
        if !ok {
            return errors.New(fmt.Sprintf("Invalid type of field settings. Supports only int and string."))
        }

        toFieldName, options := parseFieldMapping(toFieldMappingSettings)
        if toFieldId, err := strconv.ParseInt(toFieldName, 10, 32); err == nil {
            fromField.reverseId = int(toFieldId)
        } else {
            toFieldName := NameMapper(toFieldName)
            if toField, ok := toType.fields[toFieldName]; !ok {
                return errors.New(fmt.Sprintf("There is no field with name '%s' at %v", toFieldName, toType.normalizedType))
            } else {
                fromField.reverseId = int(toField.id)
                fromField.reverseName = toFieldName

                toField.reverseId = fromField.id
                toField.reverseName = fromFieldName
                toField.resolveOptions(options)
            }
        }

        fromField.resolveOptions(options)
    }

    return nil
}

func resolveMapping(fromType *mapperType, toType *mapperType, fromToMapping interface{}) (error) {
    if fromToMapping != nil {
        mappingVal := reflect.ValueOf(fromToMapping)

        if mappingVal.Kind() == reflect.Map {
            var err error

            for _, keyVal := range mappingVal.MapKeys() {
                from := keyVal.Interface()
                to := mappingVal.MapIndex(keyVal).Interface()

                var (
                    fromString string
                    toString   string

                    isFromString bool
                    isToString   bool
                )

                fromString, isFromString = from.(string)
                toString, isToString = to.(string)

                if isFromString {
                    if _, err = strconv.ParseInt(fromString, 10, 32); err == nil {
                        isFromString = false
                    }
                }

                if isToString {
                    if _, err = strconv.ParseInt(toString, 10, 32); err == nil {
                        isToString = false
                    }
                }

                if isFromString {
                    err = resolveMappingField(fromType, fromString, toType, to)
                } else if isToString {
                    err = resolveMappingField(toType, toString, fromType, from)
                } else {
                    err = errors.New("You can't map index to index.")
                }

                if err != nil {
                    return err
                }
            }

            return nil
        }
    }

    return errors.New("Can't resolve mapping or invalid type of mapping. You must provide via 'mapping' argument or via tags.")
}

func unsupportedType(value reflect.Value) (error) {
    return errors.New(fmt.Sprintf("Unsupported type: %s", value.Kind()))
}

func unknownFieldName(fieldName string) (error) {
    return errors.New(fmt.Sprintf("Unknown field name: %s", fieldName))
}
