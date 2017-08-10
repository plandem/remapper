package remapper

import (
    "testing"
)

type MyStruct struct {
    IntVal   int        `remapper:"int_val"`
    UintVal  uint       `remapper:"uint_val"`
    StrVal   string     `remapper:"str_val"`
    FloatVal float64    `remapper:"float_val"`
    BoolVal  bool       `remapper:"bool_val"`
}

var result interface{}

func BenchmarkFromNamedTypedArrayToStruct(b *testing.B) {
    mapper, err := New(&MyStruct{}, Slice([]string{}, []string{
        "int_val",
        "uint_val",
        "str_val",
        "float_val",
        "bool_val",
    }))

    if err != nil {
        panic(err)
    }

    var r interface{}
    for n := 0; n < b.N; n++ {
        r, err = mapper.Map(&[]string{
            "-1",
            "1",
            "test string",
            "1.2345",
            "true",
        })
    }

    result = r
}

func BenchmarkFromNamedUntypedArrayToStruct(b *testing.B) {
    mapper, err := New(&MyStruct{}, Slice([]interface{}{},[]string{
        "int_val",
        "uint_val",
        "str_val",
        "float_val",
        "bool_val",
    }))

    if err != nil {
        panic(err)
    }

    var r interface{}
    for n := 0; n < b.N; n++ {
        r, err = mapper.Map(&[]interface{}{
            -1,
            1,
            "test string",
            1.2345,
            true,
        })
    }

    result = r
}
