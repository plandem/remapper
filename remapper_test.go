package remapper

import (
    "testing"
    "reflect"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

type TestStructNamed struct {
    IntVal   int        `remapper:"int_val"`
    UintVal  uint       `remapper:"uint_val"`
    StrVal   string     `remapper:"str_val"`
    FloatVal float64    `remapper:"float_val"`
    BoolVal  bool       `remapper:"bool_val"`
}

type TestStructIndexed struct {
    IntVal   int        `remapper:"1"`
    UintVal  uint    `remapper:"2"`
    StrVal   string    `remapper:"4"`
    FloatVal float64    `remapper:"5"`
    BoolVal  bool    `remapper:"6"`
}

var arrayUntyped = []interface{}{
    nil,
    -1,
    uint(1),
    nil,
    "test string",
    1.2345,
    true,
    nil,
}

var arrayTyped = []string{
    "",
    "-1",
    "1",
    "",
    "test string",
    "1.2345",
    "true",
    "",
}

var arrayFieldNames = []string{
    "_f1",
    "int_val",
    "uint_val",
    "_f2",
    "str_val",
    "float_val",
    "bool_val",
    "_f3",
}

var namedArrayToStructMapping = map[string]string{
    "int_val":   "IntVal",
    "uint_val":  "UintVal",
    "str_val":   "StrVal",
    "float_val": "FloatVal",
    "bool_val":  "BoolVal",
}

var structToNamedArrayMapping = map[string]string{
    "IntVal":   "int_val",
    "UintVal":  "uint_val",
    "StrVal":   "str_val",
    "FloatVal": "float_val",
    "BoolVal":  "bool_val",
}

var indexedArrayToStructMapping = map[int]string{
    1: "IntVal",
    2: "UintVal",
    4: "StrVal",
    5: "FloatVal",
    6: "BoolVal",
}

var structToIndexedArrayMapping = map[string]int{
    "IntVal":   1,
    "UintVal":  2,
    "StrVal":   4,
    "FloatVal": 5,
    "BoolVal":  6,
}

var mappedNamedStruct = TestStructNamed{
    -1,
    1,
    "test string",
    1.2345,
    true,
}

var mappedIndexedStruct = TestStructIndexed{
    -1,
    1,
    "test string",
    1.2345,
    true,
}

func TestEmptyMapping(t *testing.T) {
    type TestStructEmpty struct {
        IntVal   int
        UintVal  uint
        StrVal   string
        FloatVal float64
        BoolVal  bool
    }

    mapper, err := New(TestStructEmpty{}, Slice(arrayTyped, arrayFieldNames))
    require.Nil(t, err)
    require.NotNil(t, mapper)

    s, err := mapper.Map(arrayTyped)
    assert.Nil(t, err)
    assert.Nil(t, s)
}

func TestFromNamedUntypedArrayWithTags(t *testing.T) {
    mapper, err := New(Slice(arrayUntyped, arrayFieldNames), TestStructNamed{})
    require.Nil(t, err)
    require.NotNil(t, mapper)

    s, err := mapper.Map(100)
    require.Nil(t, s)
    assert.NotNil(t, err)

    s, err = mapper.Map(arrayTyped)
    require.Nil(t, s)
    assert.NotNil(t, err)

    s, err = mapper.Map(arrayUntyped)
    require.Nil(t, err)
    assert.NotNil(t, s)
    assert.IsType(t, TestStructNamed{}, s)
    assert.Equal(t, mappedNamedStruct, s)
}

func TestToNamedUntypedArrayWithTags(t *testing.T) {
    mapper, err := New(TestStructNamed{}, Slice(arrayUntyped, arrayFieldNames))
    require.Nil(t, err)
    require.NotNil(t, mapper)

    a, err := mapper.Map(100)
    assert.NotNil(t, err)
    assert.Nil(t, a)

    a, err = mapper.Map(mappedNamedStruct)
    require.Nil(t, err)
    assert.NotNil(t, a)
    assert.IsType(t, []interface{}{}, a)
    assert.Equal(t, arrayUntyped, a)
}

func TestFromNamedUntypedArrayWithMapping(t *testing.T) {
    mapper, err := New(Slice(arrayUntyped, arrayFieldNames), TestStructNamed{}, namedArrayToStructMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    s, err := mapper.Map(arrayUntyped)
    require.Nil(t, err)
    assert.NotNil(t, s)
    assert.IsType(t, TestStructNamed{}, s)
    assert.Equal(t, mappedNamedStruct, s)
}

func TestToNamedUntypedArrayWithMapping(t *testing.T) {
    mapper, err := New(TestStructNamed{}, Slice(arrayUntyped, arrayFieldNames), structToNamedArrayMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    a, err := mapper.Map(mappedNamedStruct)
    require.Nil(t, err)
    assert.NotNil(t, a)
    assert.IsType(t, []interface{}{}, a)
    assert.Equal(t, arrayUntyped, a)
}

func TestFromIndexedUntypedArrayWithTags(t *testing.T) {
    mapper, err := New(arrayUntyped, TestStructNamed{})
    assert.NotNil(t, err)
    assert.Nil(t, mapper)

    mapper, err = New(arrayUntyped, TestStructIndexed{})
    require.Nil(t, err)
    require.NotNil(t, mapper)

    s, err := mapper.Map(arrayUntyped)
    require.Nil(t, err)
    assert.NotNil(t, s)
    assert.IsType(t, TestStructIndexed{}, s)
    assert.Equal(t, mappedIndexedStruct, s)
}

func TestToIndexedUntypedArrayWithTags(t *testing.T) {
    mapper, err := New(TestStructNamed{}, arrayUntyped)
    assert.NotNil(t, err)
    assert.Nil(t, mapper)

    mapper, err = New(TestStructIndexed{}, arrayUntyped)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    a, err := mapper.Map(mappedIndexedStruct)
    assert.NotNil(t, err)
    assert.Nil(t, a)
}

func TestFromIndexedUntypedArrayWithMapping(t *testing.T) {
    mapper, err := New(arrayUntyped, TestStructIndexed{}, indexedArrayToStructMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    s, err := mapper.Map(arrayUntyped)
    require.Nil(t, err)
    assert.NotNil(t, s)
    assert.IsType(t, TestStructIndexed{}, s)
    assert.Equal(t, mappedIndexedStruct, s)
}

func TestToIndexedUntypedArrayWithMapping(t *testing.T) {
    mapper, err := New(TestStructIndexed{}, arrayUntyped, structToIndexedArrayMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    a, err := mapper.Map(mappedIndexedStruct)
    assert.NotNil(t, err)
    assert.Nil(t, a)
}

func TestFromNamedTypedArrayWithTags(t *testing.T) {
    mapper, err := New(Slice(arrayTyped, arrayFieldNames), TestStructNamed{})
    require.Nil(t, err)
    require.NotNil(t, mapper)

    s, err := mapper.Map(100)
    require.Nil(t, s)
    assert.NotNil(t, err)

    s, err = mapper.Map(arrayUntyped)
    require.Nil(t, s)
    assert.NotNil(t, err)

    s, err = mapper.Map(arrayTyped)
    require.Nil(t, err)
    assert.NotNil(t, s)
    assert.IsType(t, TestStructNamed{}, s)
    assert.Equal(t, mappedNamedStruct, s)
}

func TestToNamedTypedArrayWithTags(t *testing.T) {
    mapper, err := New(TestStructNamed{}, Slice(arrayTyped, arrayFieldNames))
    require.Nil(t, err)
    require.NotNil(t, mapper)

    a, err := mapper.Map(100)
    assert.NotNil(t, err)
    assert.Nil(t, a)

    a, err = mapper.Map(mappedNamedStruct)
    assert.Nil(t, err)
    assert.NotNil(t, a)
    assert.IsType(t, []string{}, a)
    assert.Equal(t, arrayTyped, a)
}

func TestFromNamedTypedArrayWithMapping(t *testing.T) {
    mapper, err := New(Slice(arrayTyped, arrayFieldNames), TestStructNamed{}, namedArrayToStructMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    s, err := mapper.Map(arrayTyped)
    require.Nil(t, err)
    assert.NotNil(t, s)
    assert.IsType(t, TestStructNamed{}, s)
    assert.Equal(t, mappedNamedStruct, s)
}

func TestToNamedTypedArrayWithMapping(t *testing.T) {
    mapper, err := New(TestStructNamed{}, Slice(arrayTyped, arrayFieldNames), structToNamedArrayMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    a, err := mapper.Map(mappedNamedStruct)
    require.Nil(t, err)
    assert.NotNil(t, a)
    assert.IsType(t, []string{}, a)
    assert.Equal(t, arrayTyped, a)
}

func TestFromIndexedTypedArrayWithTags(t *testing.T) {
    mapper, err := New(arrayTyped, TestStructNamed{})
    assert.NotNil(t, err)
    assert.Nil(t, mapper)

    mapper, err = New(arrayTyped, TestStructIndexed{})
    require.Nil(t, err)
    require.NotNil(t, mapper)

    s, err := mapper.Map(arrayTyped)
    require.Nil(t, err)
    assert.NotNil(t, s)
    assert.IsType(t, TestStructIndexed{}, s)
    assert.Equal(t, mappedIndexedStruct, s)
}

func TestToIndexedTypedArrayWithTags(t *testing.T) {
    mapper, err := New(TestStructNamed{}, arrayTyped)
    assert.NotNil(t, err)
    assert.Nil(t, mapper)

    mapper, err = New(TestStructIndexed{}, arrayTyped)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    a, err := mapper.Map(mappedIndexedStruct)
    assert.NotNil(t, err)
    assert.Nil(t, a)
}

func TestFromIndexedTypedArrayWithMapping(t *testing.T) {
    mapper, err := New(arrayTyped, TestStructIndexed{}, indexedArrayToStructMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    s, err := mapper.Map(arrayTyped)
    require.Nil(t, err)
    assert.NotNil(t, s)
    assert.IsType(t, TestStructIndexed{}, s)
    assert.Equal(t, mappedIndexedStruct, s)
}

func TestToIndexedTypedArrayWithMapping(t *testing.T) {
    mapper, err := New(TestStructIndexed{}, arrayTyped, structToIndexedArrayMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    a, err := mapper.Map(mappedIndexedStruct)
    assert.NotNil(t, err)
    assert.Nil(t, a)
}

//TODO: refactor. Now it's possible
func TestArraySameType(t *testing.T) {
    //mapper, err := New(TestStructIndexed{}, TestStructIndexed{}, structToIndexedArrayMapping)
    //require.NotNil(t, err)
    //require.Nil(t, mapper)
    //
    //mapper, err = New(arrayTyped, arrayTyped, structToIndexedArrayMapping)
    //require.NotNil(t, err)
    //require.Nil(t, mapper)
}

func TestNamedArrayNameByName(t *testing.T) {
    mapper, err := New(TestStructNamed{}, Slice(arrayTyped, arrayFieldNames), structToNamedArrayMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    reverseName, err := mapper.NameByName(mappedNamedStruct, "floatVal")
    require.Nil(t, err)
    assert.Equal(t, "float_val", reverseName)

    reverseName, err = mapper.NameByName(arrayTyped, "float_val")
    require.Nil(t, err)
    assert.Equal(t, "floatval", reverseName)
}

func TestIndexedArrayNameByName(t *testing.T) {
    mapper, err := New(TestStructIndexed{}, arrayTyped)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    _, err = mapper.NameByName(mappedIndexedStruct, "floatVal")
    assert.NotNil(t, err)

    _, err = mapper.NameByName(arrayTyped, "float_val")
    assert.NotNil(t, err)
}

func TestNamedArrayGetByName(t *testing.T) {
    mapper, err := New(TestStructNamed{}, Slice(arrayTyped, arrayFieldNames), structToNamedArrayMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    v, err := mapper.GetByName(mappedNamedStruct, "floatVal")
    assert.Nil(t, err)
    assert.Equal(t, 1.2345, v)

    v, err = mapper.GetByName(arrayTyped, "float_val")
    require.Nil(t, err)
    assert.Equal(t, "1.2345", v)

    mapper, err = New(TestStructNamed{}, Slice(arrayUntyped, arrayFieldNames), structToNamedArrayMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    v, err = mapper.GetByName(mappedNamedStruct, "floatVal")
    assert.Nil(t, err)
    assert.Equal(t, 1.2345, v)

    v, err = mapper.GetByName(arrayUntyped, "float_val")
    assert.Nil(t, err)
    assert.Equal(t, 1.2345, v)
}

func TestIndexedArrayGetByName(t *testing.T) {
    mapper, err := New(TestStructIndexed{}, arrayTyped)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    v, err := mapper.GetByName(mappedIndexedStruct, "floatVal")
    require.Nil(t, err)
    assert.Equal(t, 1.2345, v)

    _, err = mapper.GetByName(arrayTyped, "float_val")
    assert.NotNil(t, err)

    mapper, err = New(TestStructIndexed{}, arrayUntyped)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    v, err = mapper.GetByName(mappedIndexedStruct, "floatVal")
    require.Nil(t, err)
    assert.Equal(t, 1.2345, v)

    _, err = mapper.GetByName(arrayUntyped, "float_val")
    assert.NotNil(t, err)
}

func TestNamedArraySetByName(t *testing.T) {
    mapper, err := New(TestStructNamed{}, Slice(arrayTyped, arrayFieldNames), structToNamedArrayMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    mappedNamedStruct := TestStructNamed{}
    err = mapper.SetByName(&mappedNamedStruct, "floatVal", 5432.1)
    require.Nil(t, err)
    assert.Equal(t, TestStructNamed{
        FloatVal: 5432.1,
    }, mappedNamedStruct)

    arrayTyped := []string{"", "", "", "", "", "", "", ""}
    err = mapper.SetByName(&arrayTyped, "float_val", "5432.1")
    require.Nil(t, err)
    assert.Equal(t, []string{"", "", "", "", "", "5432.1", "", ""}, arrayTyped)

    mapper, err = New(TestStructNamed{}, Slice(arrayUntyped, arrayFieldNames), structToNamedArrayMapping)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    mappedNamedStruct = TestStructNamed{}
    err = mapper.SetByName(&mappedNamedStruct, "floatVal", 5432.1)
    require.Nil(t, err)
    assert.Equal(t, TestStructNamed{
        FloatVal: 5432.1,
    }, mappedNamedStruct)

    arrayUntyped := []interface{}{nil, nil, nil, nil, nil, nil, nil, nil}
    err = mapper.SetByName(&arrayUntyped, "float_val", 5432.1)
    require.Nil(t, err)
    assert.Equal(t, []interface{}{nil, nil, nil, nil, nil, 5432.1, nil, nil}, arrayUntyped)
}

func TestIndexedArraySetByName(t *testing.T) {
    mapper, err := New(TestStructIndexed{}, arrayTyped)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    mappedIndexedStruct := TestStructIndexed{}
    err = mapper.SetByName(&mappedIndexedStruct, "floatVal", 5432.1)
    require.Nil(t, err)
    assert.Equal(t, TestStructIndexed{
        FloatVal: 5432.1,
    }, mappedIndexedStruct)

    err = mapper.SetByName(&arrayTyped, "float_val", 5432.1)
    assert.NotNil(t, err)

    mapper, err = New(TestStructIndexed{}, arrayUntyped)
    require.Nil(t, err)
    require.NotNil(t, mapper)

    mappedIndexedStruct = TestStructIndexed{}
    err = mapper.SetByName(&mappedIndexedStruct, "floatVal", 5432.1)
    require.Nil(t, err)
    assert.Equal(t, TestStructIndexed{
        FloatVal: 5432.1,
    }, mappedIndexedStruct)

    err = mapper.SetByName(&arrayUntyped, "float_val", 5432.1)
    assert.NotNil(t, err)
}

func TestCustomTagArrayNameByName(t *testing.T) {
    type TestStructNamed struct {
        IntVal   int        `db:"int_val"`
        UintVal  uint    `db:"uint_val"`
        StrVal   string    `db:"str_val"`
        FloatVal float64    `db:"float_val"`
        BoolVal  bool    `db:"bool_val"`
    }

    mappedNamedStruct := TestStructNamed{
        -1,
        1,
        "test string",
        1.2345,
        true,
    }

    mapper, err := New(TestStructNamed{}, Slice(arrayTyped, arrayFieldNames), "db")
    require.Nil(t, err)
    require.NotNil(t, mapper)

    reverseName, err := mapper.NameByName(mappedNamedStruct, "floatVal")
    require.Nil(t, err)
    assert.Equal(t, "float_val", reverseName)

    reverseName, err = mapper.NameByName(arrayTyped, "float_val")
    require.Nil(t, err)
    assert.Equal(t, "floatval", reverseName)
}

func testMapperIget(t *testing.T, mapper mapperType, target reflect.Value, i int, n string, value interface{}) {
    target = reflect.Indirect(target)

    val := mapper.get(target, i, n)
    require.NotNil(t, val)
    assert.Equal(t, true, val.IsValid())
    assert.Equal(t, value, val.Interface())
}

func testMapperIset(t *testing.T, mapper mapperType, target reflect.Value, i int, n string, value interface{}) {
    target = reflect.Indirect(target)
    mapper.set(target, i, n, reflect.ValueOf(value))
}

func testMapperIcreate(t *testing.T, mapper mapperType) {
    result, err := mapper.create()
    require.Nil(t, err)
    require.NotNil(t, result)
    assert.IsType(t, mapper.dataType, reflect.TypeOf(result))
}

func testMapperIuntyped(t *testing.T, mapper mapperType, target reflect.Value) {
    testMapperIcreate(t, mapper)

    testMapperIget(t, mapper, target, 0, "IntVal", int(0))
    testMapperIset(t, mapper, target, 0, "IntVal", int(-1))
    testMapperIget(t, mapper, target, 0, "IntVal", int(-1))

    testMapperIget(t, mapper, target, 1, "UintVal", uint(0))
    testMapperIset(t, mapper, target, 1, "UintVal", uint(100))
    testMapperIget(t, mapper, target, 1, "UintVal", uint(100))

    testMapperIget(t, mapper, target, 2, "StrVal", "")
    testMapperIset(t, mapper, target, 2, "StrVal", "test string")
    testMapperIget(t, mapper, target, 2, "StrVal", "test string")

    testMapperIget(t, mapper, target, 3, "FloatVal", 0.0)
    testMapperIset(t, mapper, target, 3, "FloatVal", 1.2345)
    testMapperIget(t, mapper, target, 3, "FloatVal", 1.2345)

    testMapperIget(t, mapper, target, 4, "BoolVal", false)
    testMapperIset(t, mapper, target, 4, "BoolVal", true)
    testMapperIget(t, mapper, target, 4, "BoolVal", true)
}

func testMapperItyped(t *testing.T, mapper mapperType, target reflect.Value) {
    testMapperIcreate(t, mapper)

    testMapperIget(t, mapper, target, 0, "IntVal", "")
    testMapperIset(t, mapper, target, 0, "IntVal", "-1")
    testMapperIget(t, mapper, target, 0, "IntVal", "-1")

    testMapperIget(t, mapper, target, 1, "UintVal", "")
    testMapperIset(t, mapper, target, 1, "UintVal", "100")
    testMapperIget(t, mapper, target, 1, "UintVal", "100")

    testMapperIget(t, mapper, target, 2, "StrVal", "")
    testMapperIset(t, mapper, target, 2, "StrVal", "test string")
    testMapperIget(t, mapper, target, 2, "StrVal", "test string")

    testMapperIget(t, mapper, target, 3, "FloatVal", "")
    testMapperIset(t, mapper, target, 3, "FloatVal", "1.2345")
    testMapperIget(t, mapper, target, 3, "FloatVal", "1.2345")

    testMapperIget(t, mapper, target, 4, "BoolVal","")
    testMapperIset(t, mapper, target, 4, "BoolVal", "true")
    testMapperIget(t, mapper, target, 4, "BoolVal", "true")
}

