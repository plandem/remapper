remapper
=========

remapper provides some convenience functions to remap data from one type to another

```go
package main

import (
    "github.com/plandem/remapper"
    "log"
)

type MyStruct struct {
    IntVal   int        `db:"int_val"`
    UintVal  uint    `db:"uint_val"`
    StrVal   string    `db:"str_val"`
    FloatVal float64    `db:"float_val"`
    BoolVal  bool    `db:"bool_val"`
}

func main() {
    //named-typed-slice and struct with tags
    mapper, err := remapper.New(&MyStruct{}, remapper.Slice([]string{}, []string{"int_val","uint_val","str_val","float_val","bool_val"}), "db")

    if err != nil {
        panic(err)
    }

    // slice -> struct
    s, err := mapper.Map(&[]string{
        "-1",
        "1",
        "test string",
        "1.2345",
        "true",
    })

    log.Printf("struct: %+v, error: %v", s, err)

    // struct -> slice
    arr, err := mapper.Map(s)
    log.Printf("arr: %+v, error: %v", arr, err)

    //named-untyped-slice and struct with manual mapping
    mapper, err = remapper.New(&MyStruct{}, remapper.Slice([]interface{}{}, []string{"int_val","uint_val","str_val","float_val","bool_val"}), map[string]string{
        "IntVal":   "int_val",
        "UintVal":  "uint_val",
        "StrVal":   "str_val",
        "FloatVal": "float_val",
        "BoolVal":  "bool_val",
    })
    
    if err != nil {
        panic(err)
    }

    // slice -> struct
    s, err = mapper.Map([]interface{}{
        -1,
        1,
        "test string",
        1.2345,
        true,
    })

    log.Printf("struct: %+v, error: %v", s, err)

    // struct -> slice
    arr, err = mapper.Map(s)
    log.Printf("arr: %+v, error: %v", arr, err)
}
```

Documentation can be found at http://godoc.org/github.com/plandem/remapper