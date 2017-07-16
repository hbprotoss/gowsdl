package main

const (
	requestTplText = `package enterprise

import "encoding/xml"

type {{.name}} struct {
	XMLName xml.Name ` + "`" + `xml:"ns2:{{.name}}"` + "`" + `
	ns      string   ` + "`" + `xml:"-"` + "`" + `

{{.members}}
}

func New{{.name}}(namespace string) *{{.name}} {
	return &{{.name}}{
		ns: namespace,
	}
}

func (req *{{.name}}) Namespace() string {
	return req.ns
}`
)
