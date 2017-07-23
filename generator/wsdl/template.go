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
	{{.Name}}({{.ParamsString}}) ({{.ReturnsString}}, error)
	{{end}}
}`

	// Implementation Template
	ImplementationTplText = `package {{.package}}
import (
	"fmt"
	"gowsdl/soap/client"
	"gowsdl/soap/req"
)

type Default{{.serviceName}} struct {
	Namespace  string
	soapClient *client.SOAPClient
}

func New{{.serviceName}}(url string, auth *client.SecurityAuth) *Default{{.serviceName}} {
	return &Default{{.serviceName}}{
		Namespace:  "{{.namespace}}",
		soapClient: client.NewSOAPClientWithWsse(url, auth),
	}
}
{{range $index, $element := .methods }}
func (s *Default{{$.serviceName}}) {{.Name}}({{.ParamsString}}) ({{.ReturnsString}}, error) {
	var envelope = req.NewEnvelope()

	var request = New{{.Name}}(s.Namespace)
	{{range $pIndex, $Param := .ParamNames -}}
	request.{{$Param | FirstLetterToUpper}} = {{$Param}}
	{{end -}}
	envelope.Body.Content = request

	var response = New{{.Name}}Response()

	err := s.soapClient.Call("{{.Name | FirstLetterToLower}}", request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return response.Return, nil
}
{{end}}
`
)
