package remapper

import (
    "strings"
)

// mappingOptions is comma separated list of additional options for mapping between two fields.
type mappingOptions string

// parseFieldMapping returns a 'reverse' name of field and additional options
func parseFieldMapping(fieldMapping string) (string, mappingOptions) {
    if idx := strings.Index(fieldMapping, ","); idx != -1 {
        return fieldMapping[:idx], mappingOptions(fieldMapping[idx+1:])
    }

    return fieldMapping, mappingOptions("")
}

// Contains returns true/false if options with name was set
func (o mappingOptions) Contains(name string) bool {
    if len(o) > 0 {
        s := string(o)

        for s != "" {
            var next string
            i := strings.Index(s, ",")

            if i >= 0 {
                s, next = s[:i], s[i+1:]
            }

            if s == name {
                return true
            }

            s = next
        }
    }

    return false
}
