package wsdl

const (
	// Entity Template
	EntityTplText = `package {{.package}}

import "encoding/xml"

type {{.structName}} struct {
	XMLName xml.Name ` + "`" + `xml:"ns2:{{.name}}"` + "`" + `
	ns      string   ` + "`" + `xml:"-"` + "`" + `

	{{range $index, $element := .members -}}
	{{.Name}} {{.Type}} ` + "`xml:\"{{.XmlName}}\"`" + `
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

	// Interface Template
	InterfaceTplText = `package {{.package}}

type {{.serviceName}} interface {
	{{range $index, $element := .methods -}}
	{{.Name}}({{.Params}}) ({{.Returns}}, error)
	{{end}}
}`
)
