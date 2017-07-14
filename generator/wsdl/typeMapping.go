package wsdl

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
