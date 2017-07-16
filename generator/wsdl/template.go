package wsdl

const (
	// Request Template
	RequestTplText =
`package {{.package}}

import "encoding/xml"

type {{.fieldName}} struct {
	XMLName xml.Name ` + "`" + `xml:"ns2:{{.name}}"` + "`" + `
	ns      string   ` + "`" + `xml:"-"` + "`" + `

{{.members}}
}

func New{{.fieldName}}(namespace string) *{{.fieldName}} {
	return &{{.fieldName}}{
		ns: namespace,
	}
}

func (req *{{.fieldName}}) Namespace() string {
	return req.ns
}`

)
