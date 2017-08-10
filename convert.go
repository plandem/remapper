package remapper

import (
    "reflect"
    "strings"
    "strconv"
)

// An function that will be using to convert from a value to required type toType (E.g. string -> int)
type ConvertFunc func(from reflect.Value, toType reflect.Type) (reflect.Value, error)

// Converts a value to required type toType or return error in case of failure. This function is using by default.
func Convert(from reflect.Value, toType reflect.Type) (reflect.Value, error) {
    var err error

    fromKind := from.Kind()
    if fromKind == reflect.Interface {
        from = reflect.Indirect(from).Elem()
        fromKind = from.Kind()
    }

    if toType.Kind() == reflect.Interface {
        toType = from.Type()
    }
    to := reflect.Indirect(reflect.New(toType))

    switch to.Kind() {

    //-> string
    case reflect.String:
        var v string
        switch fromKind {
        case reflect.String:
            v = strings.TrimSpace(from.String())
            if len(v) == 0 {
                return reflect.Value{}, nil
            }
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            v = strconv.FormatInt(from.Int(), 10)
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            v = strconv.FormatUint(from.Uint(), 10)
        case reflect.Float32:
            v = strconv.FormatFloat(from.Float(), 'f', 4, 32)
        case reflect.Float64:
            v = strconv.FormatFloat(from.Float(), 'f', 4, 64)
        case reflect.Bool:
            v = strconv.FormatBool(from.Bool())
        default:
            return reflect.Value{}, unsupportedType(from)
        }

        to.SetString(v)

        //-> int
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        var v int64

        switch fromKind {
        case reflect.String:
            s := strings.TrimSpace(from.String())
            if len(s) == 0 {
                return reflect.Value{}, nil
            }

            if v, err = strconv.ParseInt(s, 10, 64); err != nil {
                //If type of field is int, but value is float, then try to convert it
                if vv, err2 := strconv.ParseFloat(s, 10); err2 != nil {
                    return reflect.Value{}, err
                } else {
                    v = int64(vv)
                }
            }
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            v = int64(from.Int())
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            v = int64(from.Uint())
        case reflect.Float32, reflect.Float64:
            v = int64(from.Float())
        default:
            return reflect.Value{}, unsupportedType(from)
        }

        to.SetInt(v)

        //-> uint
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        var v uint64

        switch fromKind {
        case reflect.String:
            s := strings.TrimSpace(from.String())
            if len(s) == 0 {
                return reflect.Value{}, nil
            }

            if v, err = strconv.ParseUint(s, 10, 64); err != nil {
                //If type of field is uint, but value is float, then try to convert it
                if vv, err2 := strconv.ParseFloat(s, 10); err2 != nil {
                    return reflect.Value{}, err
                } else {
                    v = uint64(vv)
                }
            }
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            v = uint64(from.Int())
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            v = uint64(from.Uint())
        case reflect.Float32, reflect.Float64:
            v = uint64(from.Float())
        default:
            return reflect.Value{}, unsupportedType(from)
        }

        to.SetUint(v)

        //-> float
    case reflect.Float32, reflect.Float64:
        var v float64

        switch fromKind {
        case reflect.String:
            s := strings.Replace(strings.TrimSpace(from.String()), ",", ".", -1)
            if len(s) == 0 {
                return reflect.Value{}, nil
            }

            if v, err = strconv.ParseFloat(s, 10); err != nil {
                return reflect.Value{}, err
            }
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            v = float64(from.Int())
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            v = float64(from.Uint())
        case reflect.Float32, reflect.Float64:
            v = from.Float()
        default:
            return reflect.Value{}, unsupportedType(from)
        }

        to.SetFloat(v)

        //-> bool
    case reflect.Bool:
        var v bool

        switch fromKind {
        case reflect.String:
            s := strings.TrimSpace(from.String())
            if len(s) == 0 {
                return reflect.Value{}, nil
            }

            if v, err = strconv.ParseBool(s); err != nil {
                return reflect.Value{}, err
            }
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            v = from.Int() != 0
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            v = from.Uint() != 0
        case reflect.Float32, reflect.Float64:
            v = from.Float() != 0.0
        case reflect.Bool:
            v = from.Bool()
        default:
            return reflect.Value{}, unsupportedType(from)
        }

        to.SetBool(v)

    default:
        return reflect.Value{}, unsupportedType(to)
    }

    return to, nil
}
