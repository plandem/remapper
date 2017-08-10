package remapper

import (
    "reflect"
)

type MapMapper mapperType

func newMapMapper(dataType reflect.Type, normalizedType reflect.Type, names []string) (mapperType) {
    m := MapMapper{
        fields:         map[string]*mapperField{},
        dataType:       dataType,
        normalizedType: normalizedType,
    }

    for fieldIndex, fieldName := range names {
        m.fields[NameMapper(fieldName)] = &mapperField{
            id:        fieldIndex,
            convert:   ValueConverter,
            reverseId: -1,
        }
    }

    m.mapperTypeI = &m
    return mapperType(m)
}

// creates a new instance of map of required type.
func (m *MapMapper) create() (reflect.Value, error) {
    return reflect.MakeMap(m.dataType), nil
}

// sets a value to a map at index with name
func (m *MapMapper) set(to reflect.Value, i int, name string, value reflect.Value) {
    name = NameMapper(name)
    if _, ok := m.fields[name]; ok {
        to.SetMapIndex(reflect.ValueOf(name), value)
    }
}

// gets a value from a map at index with name
func (m *MapMapper) get(from reflect.Value, i int, name string) (reflect.Value) {
    name = NameMapper(name)
    if _, ok := m.fields[name]; ok {
        v := from.MapIndex(reflect.ValueOf(name))

        if !v.IsValid() {
            v = reflect.Zero(m.dataType.Elem())
        }

        return v
    }

    return reflect.Value{}
}
