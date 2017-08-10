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
}

func testStructMethods(t *testing.T, mapper mapperType, target reflect.Value) {
    require.NotNil(t, mapper)
    require.IsType(t, mapperType{}, mapper)
    require.IsType(t, &StructMapper{}, mapper.mapperTypeI)

    testMapperIcreate(t, mapper)

    testMapperIget(t, mapper, target, 0, int(0))
    testMapperIset(t, mapper, target, 0, int(-1))
    testMapperIget(t, mapper, target, 0, int(-1))

    testMapperIget(t, mapper, target, 1, uint(0))
    testMapperIset(t, mapper, target, 1, uint(100))
    testMapperIget(t, mapper, target, 1, uint(100))

    testMapperIget(t, mapper, target, 2, "")
    testMapperIset(t, mapper, target, 2, "test string")
    testMapperIget(t, mapper, target, 2, "test string")

    testMapperIget(t, mapper, target, 3, 0.0)
    testMapperIset(t, mapper, target, 3, 1.2345)
    testMapperIget(t, mapper, target, 3, 1.2345)

    testMapperIget(t, mapper, target, 4, false)
    testMapperIset(t, mapper, target, 4, true)
    testMapperIget(t, mapper, target, 4, true)
}
