package wsdl

import (
	"gowsdl/generator/util"
	"strings"
)

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
	var typeName = ""
	if len(types) == 1 {
		typeName = types[0]
	} else {
		typeName = types[1]
	}
	if ret := typeMapping[typeName]; ret != "" {
		return ret
	} else {
		return util.FirstLetterToUpper(typeName)
	}
}
