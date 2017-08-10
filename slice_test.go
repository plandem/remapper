package remapper

import (
    "testing"
    "reflect"
    "github.com/stretchr/testify/require"
)

func TestSliceMapper(t *testing.T) {
    namedSliceNames := []string{"IntVal", "UintVal", "StrVal", "FloatVal", "BoolVal"}
    indexedArrayLen := len(namedSliceNames)

    data := []string{"", "", "", "", ""}
    dataVal := reflect.Indirect(reflect.ValueOf(&data))

    //named-typed array
    dataType := reflect.TypeOf(data)
    require.Equal(t, reflect.Slice, dataType.Kind())

    dataNormalizedType, err := resolveType(data, reflect.Slice)
    require.Nil(t, err)
    require.Equal(t, reflect.Slice, dataNormalizedType.Kind())

    mapper := newSliceMapper(dataType, dataNormalizedType, namedSliceNames)
    testTypedSliceMethods(t, mapper, dataVal)

    //indexed-typed array
    data = []string{"", "", "", "", ""}
    mapper = newSliceMapper(dataType, dataNormalizedType, indexedArrayLen)
    testTypedSliceMethods(t, mapper, dataVal)

    //named-untyped array
    data1 := []interface{}{int(0), uint(0), "", 0.0, false}
    data1Val := reflect.Indirect(reflect.ValueOf(&data1))
    dataType = reflect.TypeOf(data1)
    require.Equal(t, reflect.Slice, dataType.Kind())

    dataNormalizedType, err = resolveType(data1, reflect.Slice)
    require.Nil(t, err)
    require.Equal(t, reflect.Slice, dataNormalizedType.Kind())

    mapper = newSliceMapper(dataType, dataNormalizedType, namedSliceNames)
    testUntypedSliceMethods(t, mapper, data1Val)

    //indexed-untyped array
    data1 = []interface{}{int(0), uint(0), "", 0.0, false}
    mapper = newSliceMapper(dataType, dataNormalizedType, indexedArrayLen)
    testUntypedSliceMethods(t, mapper, data1Val)
}

func testTypedSliceMethods(t *testing.T, mapper mapperType, target reflect.Value) {
    require.NotNil(t, mapper)
    require.IsType(t, mapperType{}, mapper)
    require.IsType(t, &SliceMapper{}, mapper.mapperTypeI)

    testMapperItyped(t, mapper, target)
}

func testUntypedSliceMethods(t *testing.T, mapper mapperType, target reflect.Value) {
    require.NotNil(t, mapper)
    require.IsType(t, mapperType{}, mapper)
    require.IsType(t, &SliceMapper{}, mapper.mapperTypeI)

    testMapperIuntyped(t, mapper, target)
}
