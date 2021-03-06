package wsdl

const (
	// Entity Template
	EntityTplText = `package {{.package}}
{{if .isRequestEntity -}}
import "encoding/xml"
{{end -}}

type {{.structName}} struct {
	{{if .isRequestEntity -}}
	XMLName xml.Name ` + "`" + `xml:"ns2:{{.name}}"` + "`" + `
	{{end -}}
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
	{{.Name}}({{.ParamsString}}) ({{.ReturnsString}}{{if .ReturnsString}}, {{end}}err error)
	{{end}}
}`

	// Implementation Template
	ImplementationTplText = `package {{.package}}
import (
	"fmt"
	"wsdl2go/soap/client"
	"wsdl2go/soap/req"
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
func (s *Default{{$.serviceName}}) {{.Name}}({{.ParamsString}}) ({{.ReturnsString}}{{if .ReturnsString}}, {{end}}err error) {
	var envelope = req.NewEnvelope()

	var request = New{{.Name}}(s.Namespace)
	{{range $pIndex, $Param := .ParamNames -}}
	request.{{$Param | FirstLetterToUpper}} = {{$Param}}
	{{end -}}
	envelope.Body.Content = request

	var response = New{{.Name}}Response(s.Namespace)

	err = s.soapClient.Call("{{.Name | FirstLetterToLower}}", request, response)
	if err != nil {
		fmt.Println(err)
		return {{if .ReturnsString}}nil, {{end}}err
	}
	return {{if .ReturnsString}}response.Return, {{end}}nil
}
{{end}}
`
)
