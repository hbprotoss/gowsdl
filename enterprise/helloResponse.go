package enterprise

import "encoding/xml"

type HelloResponse struct {
	XMLName   xml.Name `xml:"helloResponse"`

	Return struct {
		Id      int64        `xml:"id"`
		Message string    `xml:"message"`
	} `xml:"return"`
}

func NewHelloResponse() *HelloResponse {
	return &HelloResponse{
	}
}