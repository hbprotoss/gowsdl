package wsdl

import "strings"

var typeMapping = map[string]string{
	// string
	"string":           "string",
	"normalizedString": "string",
	"token":            "string",
	// date
	"date":     "time",
	"time":     "time",
	"dateTime": "time",
	// number
	"byte":    "byte",
	"short":   "int16",
	"int":     "int32",
	"integer": "int32",
	"long":    "int64",
	"decimal": "??",
	// misc
	"boolean": "bool",
	"any":     "interface{}",
}

func GetType(rawType string) string {
	types := strings.Split(rawType, ":")
	if len(types) == 1 {
		return types[0]
	} else {
		return typeMapping[types[1]]
	}
}
