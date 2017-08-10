package remapper_test

import (
    "testing"
    "reflect"
    "github.com/stretchr/testify/require"
    "github.com/stretchr/testify/assert"

    "github.com/plandem/remapper"
)

var (
    intType   = reflect.TypeOf(int(0))
    int8Type  = reflect.TypeOf(int8(0))
    int16Type = reflect.TypeOf(int16(0))
    int32Type = reflect.TypeOf(int32(0))
    int64Type = reflect.TypeOf(int64(0))

    uintType   = reflect.TypeOf(uint(0))
    uint8Type  = reflect.TypeOf(uint8(0))
    uint16Type = reflect.TypeOf(uint16(0))
    uint32Type = reflect.TypeOf(uint32(0))
    uint64Type = reflect.TypeOf(uint64(0))

    floatType   = reflect.TypeOf(0.0)
    float32Type = reflect.TypeOf(float32(0.0))
    float64Type = reflect.TypeOf(float64(0.0))

    boolType   = reflect.TypeOf(true)
    stringType = reflect.TypeOf("")
)

func TestConverterInt(t *testing.T) {
    //int <-
    testConverter(t, int(-1), intType, int(-1), false)
    testConverter(t, int(-1), int8Type, int8(-1), false)
    testConverter(t, int(-1), int16Type, int16(-1), false)
    testConverter(t, int(-1), int32Type, int32(-1), false)
    testConverter(t, int(-1), int64Type, int64(-1), false)

    testConverter(t, int8(-1), intType, int(-1), false)
    testConverter(t, int8(-1), int8Type, int8(-1), false)
    testConverter(t, int8(-1), int16Type, int16(-1), false)
    testConverter(t, int8(-1), int32Type, int32(-1), false)
    testConverter(t, int8(-1), int64Type, int64(-1), false)

    testConverter(t, int16(-1), intType, int(-1), false)
    testConverter(t, int16(-1), int8Type, int8(-1), false)
    testConverter(t, int16(-1), int16Type, int16(-1), false)
    testConverter(t, int16(-1), int32Type, int32(-1), false)
    testConverter(t, int16(-1), int64Type, int64(-1), false)

    testConverter(t, int32(-1), intType, int(-1), false)
    testConverter(t, int32(-1), int8Type, int8(-1), false)
    testConverter(t, int32(-1), int16Type, int16(-1), false)
    testConverter(t, int32(-1), int32Type, int32(-1), false)
    testConverter(t, int32(-1), int64Type, int64(-1), false)

    testConverter(t, int64(-1), intType, int(-1), false)
    testConverter(t, int64(-1), int8Type, int8(-1), false)
    testConverter(t, int64(-1), int16Type, int16(-1), false)
    testConverter(t, int64(-1), int32Type, int32(-1), false)
    testConverter(t, int64(-1), int64Type, int64(-1), false)

    testConverter(t, -1.0, intType, int(-1), false)
    testConverter(t, float32(-1.0), intType, int(-1), false)
    testConverter(t, float64(-1.0), intType, int(-1), false)

    testConverter(t, true, intType, nil, true)
    testConverter(t, false, intType, nil, true)

    testConverter(t, "-1", intType, int(-1), false)
    testConverter(t, "-1", int8Type, int8(-1), false)
    testConverter(t, "-1", int16Type, int16(-1), false)
    testConverter(t, "-1", int32Type, int32(-1), false)
    testConverter(t, "-1", int64Type, int64(-1), false)

    testConverter(t, "-1.0", intType, int(-1), false)
    testConverter(t, "-1.2", intType, int(-1), false)

    testConverter(t, "true", intType, nil, true)
    testConverter(t, "1 test string", intType, nil, true)

    //int ->
    testConverter(t, int(1), uintType, uint(1), false)
    testConverter(t, int(1), uint8Type, uint8(1), false)
    testConverter(t, int(1), uint16Type, uint16(1), false)
    testConverter(t, int(1), uint32Type, uint32(1), false)
    testConverter(t, int(1), uint64Type, uint64(1), false)

    testConverter(t, int8(1), uintType, uint(1), false)
    testConverter(t, int8(1), uint8Type, uint8(1), false)
    testConverter(t, int8(1), uint16Type, uint16(1), false)
    testConverter(t, int8(1), uint32Type, uint32(1), false)
    testConverter(t, int8(1), uint64Type, uint64(1), false)

    testConverter(t, int16(1), uintType, uint(1), false)
    testConverter(t, int16(1), uint8Type, uint8(1), false)
    testConverter(t, int16(1), uint16Type, uint16(1), false)
    testConverter(t, int16(1), uint32Type, uint32(1), false)
    testConverter(t, int16(1), uint64Type, uint64(1), false)

    testConverter(t, int32(1), uintType, uint(1), false)
    testConverter(t, int32(1), uint8Type, uint8(1), false)
    testConverter(t, int32(1), uint16Type, uint16(1), false)
    testConverter(t, int32(1), uint32Type, uint32(1), false)
    testConverter(t, int32(1), uint64Type, uint64(1), false)

    testConverter(t, int64(1), uintType, uint(1), false)
    testConverter(t, int64(1), uint8Type, uint8(1), false)
    testConverter(t, int64(1), uint16Type, uint16(1), false)
    testConverter(t, int64(1), uint32Type, uint32(1), false)
    testConverter(t, int64(1), uint64Type, uint64(1), false)

    testConverter(t, int(1), floatType, 1.0, false)
    testConverter(t, int(1), float32Type, float32(1.0), false)
    testConverter(t, int(1), float64Type, float64(1.0), false)

    testConverter(t, int(1), boolType, true, false)
    testConverter(t, int(0), boolType, false, false)

    testConverter(t, int(1), stringType, "1", false)
}

func TestConverterUint(t *testing.T) {
    //uint <-
    testConverter(t, uint(1), uintType, uint(1), false)
    testConverter(t, uint(1), uint8Type, uint8(1), false)
    testConverter(t, uint(1), uint16Type, uint16(1), false)
    testConverter(t, uint(1), uint32Type, uint32(1), false)
    testConverter(t, uint(1), uint64Type, uint64(1), false)

    testConverter(t, uint8(1), uintType, uint(1), false)
    testConverter(t, uint8(1), uint8Type, uint8(1), false)
    testConverter(t, uint8(1), uint16Type, uint16(1), false)
    testConverter(t, uint8(1), uint32Type, uint32(1), false)
    testConverter(t, uint8(1), uint64Type, uint64(1), false)

    testConverter(t, uint16(1), uintType, uint(1), false)
    testConverter(t, uint16(1), uint8Type, uint8(1), false)
    testConverter(t, uint16(1), uint16Type, uint16(1), false)
    testConverter(t, uint16(1), uint32Type, uint32(1), false)
    testConverter(t, uint16(1), uint64Type, uint64(1), false)

    testConverter(t, uint32(1), uintType, uint(1), false)
    testConverter(t, uint32(1), uint8Type, uint8(1), false)
    testConverter(t, uint32(1), uint16Type, uint16(1), false)
    testConverter(t, uint32(1), uint32Type, uint32(1), false)
    testConverter(t, uint32(1), uint64Type, uint64(1), false)

    testConverter(t, uint64(1), uintType, uint(1), false)
    testConverter(t, uint64(1), uint8Type, uint8(1), false)
    testConverter(t, uint64(1), uint16Type, uint16(1), false)
    testConverter(t, uint64(1), uint32Type, uint32(1), false)
    testConverter(t, uint64(1), uint64Type, uint64(1), false)

    testConverter(t, 1.0, uintType, uint(1), false)
    testConverter(t, float32(1.0), uintType, uint(1), false)
    testConverter(t, float64(1.0), uintType, uint(1), false)

    testConverter(t, true, uintType, nil, true)
    testConverter(t, false, uintType, nil, true)

    testConverter(t, "1", uintType, uint(1), false)
    testConverter(t, "1", uint8Type, uint8(1), false)
    testConverter(t, "1", uint16Type, uint16(1), false)
    testConverter(t, "1", uint32Type, uint32(1), false)
    testConverter(t, "1", uint64Type, uint64(1), false)

    testConverter(t, "1.0", uintType, uint(1), false)
    testConverter(t, "1.2", uintType, uint(1), false)

    testConverter(t, "true", uintType, nil, true)
    testConverter(t, "1 test string", uintType, nil, true)

    //uint ->
    testConverter(t, uint(1), intType, int(1), false)
    testConverter(t, uint(1), int8Type, int8(1), false)
    testConverter(t, uint(1), int16Type, int16(1), false)
    testConverter(t, uint(1), int32Type, int32(1), false)
    testConverter(t, uint(1), int64Type, int64(1), false)

    testConverter(t, uint8(1), intType, int(1), false)
    testConverter(t, uint8(1), int8Type, int8(1), false)
    testConverter(t, uint8(1), int16Type, int16(1), false)
    testConverter(t, uint8(1), int32Type, int32(1), false)
    testConverter(t, uint8(1), int64Type, int64(1), false)

    testConverter(t, uint16(1), intType, int(1), false)
    testConverter(t, uint16(1), int8Type, int8(1), false)
    testConverter(t, uint16(1), int16Type, int16(1), false)
    testConverter(t, uint16(1), int32Type, int32(1), false)
    testConverter(t, uint16(1), int64Type, int64(1), false)

    testConverter(t, uint32(1), intType, int(1), false)
    testConverter(t, uint32(1), int8Type, int8(1), false)
    testConverter(t, uint32(1), int16Type, int16(1), false)
    testConverter(t, uint32(1), int32Type, int32(1), false)
    testConverter(t, uint32(1), int64Type, int64(1), false)

    testConverter(t, uint64(1), intType, int(1), false)
    testConverter(t, uint64(1), int8Type, int8(1), false)
    testConverter(t, uint64(1), int16Type, int16(1), false)
    testConverter(t, uint64(1), int32Type, int32(1), false)
    testConverter(t, uint64(1), int64Type, int64(1), false)

    testConverter(t, uint(1), floatType, 1.0, false)
    testConverter(t, uint(1), float32Type, float32(1.0), false)
    testConverter(t, uint(1), float64Type, float64(1.0), false)

    testConverter(t, uint(1), boolType, true, false)
    testConverter(t, uint(0), boolType, false, false)

    testConverter(t, uint(1), stringType, "1", false)
}

func TestConverterFloat(t *testing.T) {
    //float <-
    testConverter(t, uint(1), floatType, 1.0, false)
    testConverter(t, uint(1), float32Type, float32(1.0), false)
    testConverter(t, uint(1), float64Type, float64(1.0), false)

    testConverter(t, uint8(1), floatType, 1.0, false)
    testConverter(t, uint8(1), float32Type, float32(1.0), false)
    testConverter(t, uint8(1), float64Type, float64(1.0), false)

    testConverter(t, uint16(1), floatType, 1.0, false)
    testConverter(t, uint16(1), float32Type, float32(1.0), false)
    testConverter(t, uint16(1), float64Type, float64(1.0), false)

    testConverter(t, uint32(1), floatType, 1.0, false)
    testConverter(t, uint32(1), float32Type, float32(1.0), false)
    testConverter(t, uint32(1), float64Type, float64(1.0), false)

    testConverter(t, uint64(1), floatType, 1.0, false)
    testConverter(t, uint64(1), float32Type, float32(1.0), false)
    testConverter(t, uint64(1), float64Type, float64(1.0), false)

    testConverter(t, int(1), floatType, 1.0, false)
    testConverter(t, int(1), float32Type, float32(1.0), false)
    testConverter(t, int(1), float64Type, float64(1.0), false)

    testConverter(t, int8(1), floatType, 1.0, false)
    testConverter(t, int8(1), float32Type, float32(1.0), false)
    testConverter(t, int8(1), float64Type, float64(1.0), false)

    testConverter(t, int16(1), floatType, 1.0, false)
    testConverter(t, int16(1), float32Type, float32(1.0), false)
    testConverter(t, int16(1), float64Type, float64(1.0), false)

    testConverter(t, int32(1), floatType, 1.0, false)
    testConverter(t, int32(1), float32Type, float32(1.0), false)
    testConverter(t, int32(1), float64Type, float64(1.0), false)

    testConverter(t, int64(1), floatType, 1.0, false)
    testConverter(t, int64(1), float32Type, float32(1.0), false)
    testConverter(t, int64(1), float64Type, float64(1.0), false)

    testConverter(t, 1.0, floatType, 1.0, false)
    testConverter(t, float32(1.0), float32Type, float32(1.0), false)
    testConverter(t, float64(1.0), float64Type, float64(1.0), false)

    testConverter(t, true, floatType, nil, true)
    testConverter(t, false, floatType, nil, true)

    testConverter(t, "1", floatType, 1.0, false)
    testConverter(t, "1.0", floatType, 1.0, false)

    //float ->
    testConverter(t, 1.0, intType, int(1), false)
    testConverter(t, float32(1.0), intType, int(1.0), false)
    testConverter(t, float64(1.0), intType, int(1.0), false)

    testConverter(t, 1.0, uintType, uint(1), false)
    testConverter(t, float32(1.0), uintType, uint(1.0), false)
    testConverter(t, float64(1.0), uintType, uint(1.0), false)

    testConverter(t, 1.0, boolType, true, false)
    testConverter(t, 0.0, boolType, false, false)

    testConverter(t, 1.2, stringType, "1.2000", false)
}

func TestConverterString(t *testing.T) {
    //string <-
    testConverter(t, uint(1), stringType, "1", false)
    testConverter(t, uint8(1), stringType, "1", false)
    testConverter(t, uint16(1), stringType, "1", false)
    testConverter(t, uint32(1), stringType, "1", false)
    testConverter(t, uint64(1), stringType, "1", false)

    testConverter(t, int(1), stringType, "1", false)
    testConverter(t, int8(1), stringType, "1", false)
    testConverter(t, int16(1), stringType, "1", false)
    testConverter(t, int32(1), stringType, "1", false)
    testConverter(t, int64(1), stringType, "1", false)

    testConverter(t, 1.0, stringType, "1.0000", false)
    testConverter(t, float32(1.0), stringType, "1.0000", false)
    testConverter(t, float32(1.0), stringType, "1.0000", false)

    testConverter(t, true, stringType, "true", false)
    testConverter(t, false, stringType, "false", false)

    //string ->
    testConverter(t, "1", uintType, uint(1), false)
    testConverter(t, "1", uint8Type, uint8(1), false)
    testConverter(t, "1", uint16Type, uint16(1), false)
    testConverter(t, "1", uint32Type, uint32(1), false)
    testConverter(t, "1", uint64Type, uint64(1), false)

    testConverter(t, "-1", intType, int(-1), false)
    testConverter(t, "-1", int8Type, int8(-1), false)
    testConverter(t, "-1", int16Type, int16(-1), false)
    testConverter(t, "-1", int32Type, int32(-1), false)
    testConverter(t, "-1", int64Type, int64(-1), false)

    testConverter(t, "1.2", floatType, 1.2, false)
    testConverter(t, "1.2", float32Type, float32(1.2), false)
    testConverter(t, "1.2", float64Type, float64(1.2), false)

    testConverter(t, "true", boolType, true, false)
    testConverter(t, "false", boolType, false, false)

    testConverter(t, "test string", intType, nil, true)
}

func TestConverterBool(t *testing.T) {
    //bool <-
    testConverter(t, uint(1), boolType, true, false)
    testConverter(t, uint8(1), boolType, true, false)
    testConverter(t, uint16(1), boolType, true, false)
    testConverter(t, uint32(1), boolType, true, false)
    testConverter(t, uint64(1), boolType, true, false)

    testConverter(t, int(1), boolType, true, false)
    testConverter(t, int8(1), boolType, true, false)
    testConverter(t, int16(1), boolType, true, false)
    testConverter(t, int32(1), boolType, true, false)
    testConverter(t, int64(1), boolType, true, false)

    testConverter(t, 1.0, boolType, true, false)
    testConverter(t, float32(1.0), boolType, true, false)
    testConverter(t, float64(1.0), boolType, true, false)

    testConverter(t, "true", boolType, true, false)
    testConverter(t, "1", boolType, true, false)

    //bool ->
    testConverter(t, true, intType, nil, true)
    testConverter(t, true, uintType, nil, true)
    testConverter(t, true, floatType, nil, true)

    testConverter(t, true, stringType, "true", false)
    testConverter(t, false, stringType, "false", false)
    testConverter(t, true, boolType, true, false)
}

func testConverter(t *testing.T, from interface{}, toType reflect.Type, result interface{}, hasError bool) {
    to, err := remapper.Convert(reflect.ValueOf(from), toType)

    if hasError {
        require.NotNil(t, err)
        require.NotNil(t, to)
        assert.Equal(t, false, to.IsValid())
    } else {
        require.Nil(t, err)
        require.NotNil(t, result)
        assert.IsType(t, result, to.Interface())
        assert.Equal(t, result, to.Interface())
    }
}
