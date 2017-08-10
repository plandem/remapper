// Copyright 2017 Andrey Gayvorosnky. All rights reserved.
// Use of this source code is governed by the MIT
// license which can be found in the LICENSE file.

//
// Package remapper provides some convenience functions for converting to/from one type of data to another.
//
package remapper

import (
    "fmt"
    "errors"
    "reflect"
)

type option func(m *Mapper)(error)

// Creates a new Mapper object that can be used for mapping from one type of data to another. E.g.: slice -> struct, struct -> slice, struct -> map, ...
func New(args ...interface{})(*Mapper, error) {
    m := &Mapper{}
    options := []option{}

    //resolve type1 mapper
    if len(args) > 0 {
        if typeMapper, err := resolveTypeMapper(args[0]); err == nil {
            options = append(options, typeMapper)
        } else {
            return nil, err
        }
    }

    //resolve type2 mapper
    if len(args) > 1 {
        if typeMapper, err := resolveTypeMapper(args[1]); err == nil {
            options = append(options, typeMapper)
        } else {
            return nil, err
        }
    }

    //resolve mapping between types
    var mapping interface{}
    if len(args) > 2 {
        mapping = args[2]
    }

    if typeMapping, err := resolveTypeMapping(mapping); err == nil {
        options = append(options, typeMapping)
    } else {
        return nil, err
    }

    // process options to setup mapper
    if len(options) < 3 {
        return nil, errors.New("Not enough information to setup mapper between types. You must provide at least two types and mapping if required.")
    }

    for _, op := range options {
        err := op(m)
        if err != nil {
            return nil, err
        }
    }

    return m, nil
}

func Slice(t interface{}, options interface{})(option) {
    return func(m *Mapper)(error) {
        normalizedType, err := resolveType(t, reflect.Slice)

        if err == nil {
            sliceType := reflect.TypeOf(t)
            sliceMapper := newSliceMapper(sliceType, normalizedType, options)
            err = m.setType(&sliceMapper)
        }

        return err
    }
}

//Map returns option to setup map mapper
func Map(t interface{}, names []string)(option) {
    return func(m *Mapper)(error) {
        normalizedType, err := resolveType(t, reflect.Map)

        if err == nil {
            mapType := reflect.TypeOf(t)
            mapMapper := newMapMapper(mapType, normalizedType, names)
            err = m.setType(&mapMapper)
        }

        return err
    }
}

//Struct returns option to setup struct mapper
func Struct(t interface{})(option) {
    return func(m *Mapper) (error) {
        normalizedType, err := resolveType(t, reflect.Struct)

        if err == nil {
            structType := reflect.TypeOf(t)
            structMapper := newStructMapper(structType, normalizedType)
            err = m.setType(&structMapper)
        }

        return err
    }
}

//resolveTypeMapper returns option to setup mapper for type
func resolveTypeMapper(t interface{})(option, error) {
    normalizedType, err := resolveType(t, reflect.Struct, reflect.Slice, reflect.Map, reflect.Func)
    if err != nil {
        return nil, err
    }

    var invalidFuncType option
    invalidFuncOption := errors.New(fmt.Sprintf("Unknown option. It must be type of %+s", reflect.TypeOf(invalidFuncType)))

    switch normalizedType.Kind() {
    case reflect.Slice:
        return Slice(t, nil), nil
    case reflect.Struct:
        return Struct(t), nil
    case reflect.Map:
        return Map(t, nil), nil
    case reflect.Func:
        if type1Option, ok := t.(option); !ok {
            return nil, invalidFuncOption
        } else {
            return type1Option, nil
        }
    }

    return nil, invalidFuncOption
}

//resolveTypeMapping returns option to setup mapping between types
func resolveTypeMapping(m interface{})(option, error) {
    if m == nil {
        m = TagName
    }

    normalizedMapping, err := resolveType(m, reflect.Map, reflect.String)
    if err != nil {
        return nil, err
    }

    switch normalizedMapping.Kind() {
    case reflect.String:
        m := m.(string)
        return tagMapping(m), nil
    case reflect.Map:
        return fieldMapping(m), nil
    }

    return nil, errors.New("You must provide mapping between types.")
}

//tagMapping returns option to setup mapping via tags of struct
func tagMapping(tag string)(option) {
    return func(m *Mapper) (error) {
        var mapping map[string]string
        var err error

        if m.types[0].normalizedType.Kind() == reflect.Struct {
            mapping, err = getTagMapping(m.types[0], tag, false)
        } else if m.types[1].normalizedType.Kind() == reflect.Struct{
            mapping, err = getTagMapping(m.types[1], tag, true)
        } else {
            err = errors.New("Only struct supports mapping via tags. You must provide manual mapping via 'mapping' argument for other types.")
        }

        if err != nil {
            return err
        }

        return fieldMapping(mapping)(m)
    }
}

//fieldMapping returns option to setup mapping via map with names or indexes
func fieldMapping(mapping interface{})(option) {
    return func(m *Mapper) (error) {
        return resolveMapping(m.types[0], m.types[1], mapping)
    }
}

// getTagMapping returns a mapping extracted from tags tagName of struct t that can be used to link fields.
func getTagMapping(t *mapperType, tagName string, reverse bool) (map[string]string, error) {
    if t.normalizedType.Kind() != reflect.Struct {
        return nil, errors.New("Only struct supports mapping via tags. You must provide manual mapping for other types.")
    }

    tagMapping := make(map[string]string)
    for fieldName, field := range t.fields {
        f := t.normalizedType.Field(field.id)

        if fromTag := f.Tag.Get(tagName); len(fromTag) > 0 {
            if reverse {
                tagMapping[fromTag] = fieldName
            } else {
                tagMapping[fieldName] = fromTag
            }
        }
    }

    return tagMapping, nil
}
