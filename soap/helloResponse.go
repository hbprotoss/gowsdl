package soap

import "encoding/xml"

type HelloResponse struct {
	XMLName xml.Name `xml:"ns2:helloResponse"`
	Namespace string `xml:"xmlns:ns2,attr"`

	Return *HelloResponseData `xml:"return"`
}

type HelloResponseData struct {
	Id int32
	Message string
}

func NewHelloResponse() *HelloResponse {
	return &HelloResponse{
		Namespace: "http://service.server.soa.demo.hbprotoss.io/",
		Return: &HelloResponseData{},
	}
}
