package remapper_test

import (
    "log"
    "fmt"

    "github.com/plandem/remapper"
)

func ExampleMapper() {
    type type1 struct {}
    type type2 struct {}

    //creates a mapper for: type1 <-> type2
    remapper.New(type1{}, type2{})

    //order of type1/type2 doesn't matter. Next one is same
    remapper.New(type2{}, type1{})
}

func ExampleSliceMapper() {
    //creates a mapper for: typed-indexed slice <-> struct
    remapper.New([]string{}, struct {}{})

    //creates a mapper for: untyped-indexed slice <-> struct
    remapper.New([]interface{}{}, struct {}{})

    //create a mapper for: typed-indexed slice <-> struct with known length(10)
    remapper.New([]string{}, struct {}{}, nil, 10)

    //create a mapper for: typed-named slice <-> struct with provided field names
    remapper.New([]string{}, struct {}{}, nil, []string{
        "int_val",
        "uint_val",
        "str_val",
        "float_val",
        "bool_val",
    })
}

func ExampleStructMapper() {
    type MyStruct struct {
        IntVal   int        `remapper:"0"`
        UintVal  uint       `remapper:"1"`
        StrVal   string     `remapper:"2"`
        FloatVal float64    `remapper:"3"`
        BoolVal  bool       `remapper:"4"`
    }

    //create a mapper for: struct <-> indexed-slice with mapping via tags
    remapper.New(&MyStruct{}, []string{})

    //create a mapper for: struct <-> indexed-slice with manual mapping
    remapper.New(&MyStruct{}, []string{}, map[string]int{
        "IntVal": 0,
        "UintVal": 1,
        "StrVal": 2,
        "FloatVal": 3,
        "BoolVal": 4,
    })

    //create a mapper for: struct <-> named-slice with manual mapping
    remapper.New(&MyStruct{}, []string{}, map[string]string{
        "IntVal": "int_val",
        "UintVal": "uint_val",
        "StrVal": "str_val",
        "FloatVal": "float_val",
        "BoolVal": "bool_val",
    }, []string{
        "int_val",
        "uint_val",
        "str_val",
        "float_val",
        "bool_val",
    })
}

func Example_indexedSliceAndStructWithTags() {
    type MyStruct struct {
        IntVal   int        `remapper:"0"`
        UintVal  uint       `remapper:"1"`
        StrVal   string     `remapper:"2"`
        FloatVal float64    `remapper:"3"`
        BoolVal  bool       `remapper:"4"`
    }

    mapper, err := remapper.New(&MyStruct{}, []string{})
    if err != nil {
        panic(err)
    }

    // convert slice -> struct
    s, err := mapper.Map([]string{
        "-1",
        "1",
        "test string",
        "1.2345",
        "true",
    })
    fmt.Println(s, err)

    //Output:
    // &{-1 1 test string 1.2345 true} <nil>
}

func Example_namedSliceAndStructWithCustomTags() {
    type MyStruct struct {
        IntVal   int        `db:"int_val"`
        UintVal  uint       `db:"uint_val"`
        StrVal   string     `db:"str_val"`
        FloatVal float64    `db:"float_val"`
        BoolVal  bool       `db:"bool_val"`
    }

    mapper, err := remapper.New(&MyStruct{}, remapper.Slice([]string{}, []string{
        "int_val",
        "uint_val",
        "str_val",
        "float_val",
        "bool_val",
    }), "db")

    if err != nil {
        panic(err)
    }

    // convert slice -> struct
    s, err := mapper.Map(&[]string{
        "-1",
        "1",
        "test string",
        "1.2345",
        "true",
    })

    fmt.Println(s, err)

    // convert struct -> slice
    arr, err := mapper.Map(s)
    fmt.Println(arr, err)

    //Output:
    // &{-1 1 test string 1.2345 true} <nil>
    // [-1 1 test string 1.2345 true] <nil>
}

func Example_namedSliceAndStructWithoutTags() {
    type MyStruct struct {
        IntVal   int
        UintVal  uint
        StrVal   string
        FloatVal float64
        BoolVal  bool
    }

    mapper, err := remapper.New(&MyStruct{}, remapper.Slice([]string{}, []string{
        "int_val",
        "uint_val",
        "str_val",
        "float_val",
        "bool_val",
    }), map[string]string{
        "IntVal":   "int_val",
        "UintVal":  "uint_val",
        "StrVal":   "str_val",
        "FloatVal": "float_val",
        "BoolVal":  "bool_val",
    })

    if err != nil {
        panic(err)
    }

    // convert slice -> struct
    s, err := mapper.Map([]string{
        "-1",
        "1",
        "test string",
        "1.2345",
        "true",
    })

    log.Printf("struct: %+v, error: %v", s, err)

    // convert struct -> slice
    arr, err := mapper.Map(s)
    log.Printf("slice: %+v, error: %v", arr, err)

    vs, _:= mapper.GetByName(s, "IntVal")
    varr, _:= mapper.GetByName(arr, "int_val")

    log.Printf("value from struct: %v, value from slice: %v", vs, varr)
}

func Example_untypedSliceAndStruct() {
    type MyStruct struct {
        IntVal   int        `remapper:"0"`
        UintVal  uint       `remapper:"1"`
        StrVal   string     `remapper:"2"`
        FloatVal float64    `remapper:"3"`
        BoolVal  bool       `remapper:"4"`
    }

    mapper, err := remapper.New(&MyStruct{}, []interface{}{})
    if err != nil {
        panic(err)
    }

    // convert slice -> struct
    s, err := mapper.Map([]interface{}{
        -1,
        1,
        "test string",
        1.2345,
        true,
    })

    fmt.Println(s, err)

    //Output:
    // &{-1 1 test string 1.2345 true} <nil>
}