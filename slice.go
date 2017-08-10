package remapper

import (
    "errors"
    "reflect"
    "strconv"
)

// SliceMapper is mapper to convert from/to slice
//
// slice can be:
// - named: i.e. additional 'field names' were provided that will be using to convert from/to
// - indexed: i.e. slice doesn't have any names, so indexes will be using to convert from/to
//
// also additionally slice can be:
// - typed: i.e. slice has type, e.g. []string
// - untyped: i.e. slice has no any type, e.g. []interface{}
type SliceMapper mapperType

// newSliceMapper creates a new SliceMapper that configured with options to map from/to slice
//
// - for 'named' slice you must provide a 'field names' via []string. E.g.: []string{"first_name", "birthday",...}
// - for 'indexed' slice you can provide a length of slice. Omit it if you are going only to convert from this type of slice.
func newSliceMapper(dataType reflect.Type, normalizedType reflect.Type, options interface{}) (mapperType) {
    m := SliceMapper{
        fields:         map[string]*mapperField{},
        dataType:       dataType,
        normalizedType: normalizedType,
    }

    if options != nil {
        //is indexed array with length?
        if arrayLen, isIndexedArray := options.(int); isIndexedArray {
            for fieldIndex := 0; fieldIndex < arrayLen; fieldIndex++ {
                fieldName := strconv.FormatInt(int64(fieldIndex), 10)
                m.fields[NameMapper(fieldName)] = &mapperField{
                    id:        fieldIndex,
                    convert:   Convert,
                    reverseId: -1,
                }
            }
        } else {
            //is named array with names: []string?
            fieldNames := reflect.ValueOf(options)
            if fieldNames.Kind() == reflect.Slice || fieldNames.Kind() == reflect.Array {
                for fieldIndex, totalFields := 0, fieldNames.Len(); fieldIndex < totalFields; fieldIndex += 1 {
                    fieldName := fieldNames.Index(fieldIndex).Interface().(string)
                    m.fields[NameMapper(fieldName)] = &mapperField{
                        id:        fieldIndex,
                        convert:   Convert,
                        reverseId: -1,
                    }
                }
            }
        }
    }

    m.mapperTypeI = &m
    return mapperType(m)
}

// creates a new instance of slice of required type.
// N.B.: It's not possible to create an indexed slice without length - see options for 'indexed' slices
func (m *SliceMapper) create() (reflect.Value, error) {
    arrLen := len(m.fields)

    if arrLen == 0 {
        return reflect.Value{}, errors.New("Can't get length of a new slice to create.")
    }

    return reflect.MakeSlice(m.dataType, arrLen, arrLen), nil
}

// sets a value to a slice at i index
func (m *SliceMapper) set(to reflect.Value, i int, value reflect.Value) {
    if i < to.Len() {
        to.Index(i).Set(value)
    }
}

// gets a value from a slice at i index
func (m *SliceMapper) get(from reflect.Value, i int) (reflect.Value) {
    if i < from.Len() {
        return from.Index(i)
    }

    return reflect.Value{}
}
