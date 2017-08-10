package remapper

import (
    "reflect"
)

// StructMapper is mapper to convert from/to struct
type StructMapper mapperType

// newStructMapper creates a new StructMapper to map from/to struct
func newStructMapper(dataType reflect.Type, normalizedType reflect.Type) (mapperType) {
    m := StructMapper{
        fields:         map[string]*mapperField{},
        dataType:       dataType,
        normalizedType: normalizedType,
    }

    for i, i_max := 0, normalizedType.NumField(); i < i_max; i++ {
        m.fields[NameMapper(normalizedType.Field(i).Name)] = &mapperField{
            id:        i,
            convert:   Convert,
            reverseId: -1,
        }
    }

    m.mapperTypeI = &m
    return mapperType(m)
}

// creates a new instance of struct of required type.
func (m *StructMapper) create() (reflect.Value, error) {
    return reflect.Indirect(reflect.New(m.normalizedType)), nil
}

// sets a value to a field of struct with i index
func (m *StructMapper) set(to reflect.Value, i int, name string, value reflect.Value) {
    to.Field(i).Set(value)
}

// gets a value from field of struct with i index
func (m *StructMapper) get(from reflect.Value, i int, name string) (reflect.Value) {
    return from.Field(i)
}

