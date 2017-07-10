package enterprise

import "encoding/xml"

type HelloResponse struct {
	XMLName   xml.Name `xml:"helloResponse"`

	Return struct {
		Id      int64        `xml:"id"`
		Message string    `xml:"message"`
	} `xml:"return"`
}

type HelloResponseData struct {
	XMLName xml.Name `xml:"return"`
}

func NewHelloResponse() *HelloResponse {
	return &HelloResponse{
	}
}
