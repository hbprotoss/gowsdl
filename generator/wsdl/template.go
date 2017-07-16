package wsdl

const (
	// Entity Template
	EntityTplText = `package {{.package}}

import "encoding/xml"

type {{.structName}} struct {
	XMLName xml.Name ` + "`" + `xml:"ns2:{{.name}}"` + "`" + `
	ns      string   ` + "`" + `xml:"-"` + "`" + `

	{{range $index, $element := .members -}}
	{{.FieldName}} {{.FieldType}} ` + "`xml:\"{{.XmlName}}\"`" + `
	{{end}}
}

func New{{.structName}}(namespace string) *{{.structName}} {
	return &{{.structName}}{
		ns: namespace,
	}
}

func (req *{{.structName}}) Namespace() string {
	return req.ns
}`
)
