package remapper

// mapperField hold minimal information to map between two fields
type mapperField struct {
    // ID of field
    id int

    // Reverse ID of field
    reverseId int

    // Reverse Name of field
    reverseName string

    // Ignore or not field. It's a normal to have linked fields that must not be mapped.
    // E.g.: you have field that must not be mapped, but you would like to get value manually via 'GetByName'
    //
    // Default: false
    omit bool

    // Function that will be using to convert value for this field. Default: Convert
    convert ConvertFunc
}

// resolveOptions configures a field with provided options
func (f *mapperField) resolveOptions(options mappingOptions) {
    f.omit = options.Contains("omit") || options.Contains("-")
}
