package util

import "strings"

func GetEntityName(xmlName string) string {
	types := strings.Split(xmlName, ":")
	var typeName = ""
	if len(types) == 1 {
		typeName = types[0]
	} else {
		typeName = types[1]
	}
	return typeName
}
