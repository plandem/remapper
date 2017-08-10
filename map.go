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

    m.mapperTypeI = &m
    return mapperType(m)
}

// creates a new instance of map of required type.
func (m *MapMapper) create() (reflect.Value, error) {
    //mapS := len(m.fields)
    //
    //if arrLen == 0 {
    //    return reflect.Value{}, errors.New("Can't get length of a new slice to create.")
    //}
    //
    //return reflect.MakeSlice(m.dataType, arrLen, arrLen), nil
   return reflect.MakeMap(m.dataType), nil
}

// sets a value to a map at i index
func (m *MapMapper) set(to reflect.Value, i int, value reflect.Value) {
    //if i < to.Len() {
    //    to.MapIndex(i).Set(value)
    //}
}

// gets a value from a map at i index
func (m *MapMapper) get(from reflect.Value, i int) (reflect.Value) {
    //if i < from.Len() {
    //    return from.MapIndex(i)
    //}

    return reflect.Value{}
}
