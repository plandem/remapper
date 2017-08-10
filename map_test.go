package remapper

import (
    "testing"
    "reflect"
    "github.com/stretchr/testify/require"
)

func TestMapMapper(t *testing.T) {
    mapNames := []string{"IntVal", "UintVal", "StrVal", "FloatVal", "BoolVal"}

    data := map[string]string{}
    dataVal := reflect.Indirect(reflect.ValueOf(&data))

    //typed map
    dataType := reflect.TypeOf(data)
    require.Equal(t, reflect.Map, dataType.Kind())

    dataNormalizedType, err := resolveType(data, reflect.Map)
    require.Nil(t, err)
    require.Equal(t, reflect.Map, dataNormalizedType.Kind())

    mapper := newMapMapper(dataType, dataNormalizedType, mapNames)
    testTypedMapMethods(t, mapper, dataVal)

    //untyped map is not possible, so no tests for it.
}

func testTypedMapMethods(t *testing.T, mapper mapperType, target reflect.Value) {
    require.NotNil(t, mapper)
    require.IsType(t, mapperType{}, mapper)
    require.IsType(t, &MapMapper{}, mapper.mapperTypeI)

    testMapperItyped(t, mapper, target)
}
