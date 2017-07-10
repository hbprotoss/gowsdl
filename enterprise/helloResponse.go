package enterprise

import "encoding/xml"

type HelloResponse struct {
	XMLName xml.Name `xml:"helloResponse"`
	Namespace string `xml:"xmlns:ns2,attr"`

	Return *HelloResponseData `xml:"return"`
}

type HelloResponseData struct {
	XMLName xml.Name `xml:"return"`
	Id int64		`xml:"id"`
	Message string	`xml:"message"`
}

func NewHelloResponse() *HelloResponse {
	return &HelloResponse{
		Namespace: "http://service.server.soa.demo.hbprotoss.io/",
		Return: &HelloResponseData{},
	}
}
