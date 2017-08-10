package remapper

import (
    "testing"
    "reflect"
    "github.com/stretchr/testify/require"
)

func TestStructMapper(t *testing.T) {
    type MyStruct struct {
        IntVal   int
        UintVal  uint
        StrVal   string
        FloatVal float64
        BoolVal  bool
    }

    data := MyStruct{}
    dataVal := reflect.Indirect(reflect.ValueOf(&data))

    pData := &MyStruct{}
    pDataVal := reflect.Indirect(reflect.ValueOf(pData))

    //struct
    dataType := reflect.TypeOf(data)
    require.Equal(t, reflect.Struct, dataType.Kind())

    dataNormalizedType, err := resolveType(data, reflect.Struct)
    require.Nil(t, err)
    require.Equal(t, reflect.Struct, dataNormalizedType.Kind())

    mapper := newStructMapper(dataType, dataNormalizedType)
    testStructMethods(t, mapper, dataVal)

    //pointer to struct
    dataType = reflect.TypeOf(pData)
    require.Equal(t, reflect.Ptr, dataType.Kind())

    dataNormalizedType, err = resolveType(pData, reflect.Struct)
    require.Nil(t, err)
    require.Equal(t, reflect.Struct, dataNormalizedType.Kind())

    mapper = newStructMapper(dataType, dataNormalizedType)
    testStructMethods(t, mapper, pDataVal)

    //typed struct is possible, so no tests for it
}

func testStructMethods(t *testing.T, mapper mapperType, target reflect.Value) {
    require.NotNil(t, mapper)
    require.IsType(t, mapperType{}, mapper)
    require.IsType(t, &StructMapper{}, mapper.mapperTypeI)

    testMapperIuntyped(t, mapper, target)
}
