package wsdl

import (
	"wsdl2go/util"
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
	return typeMapping[rawType]
}

func GetTypeWithUpperEntity(rawType string) string {
	var typeName = util.GetEntityName(rawType)
	if ret := GetType(typeName); ret != "" {
		return ret
	} else {
		return "*" + util.FirstLetterToUpper(typeName)
	}
}
