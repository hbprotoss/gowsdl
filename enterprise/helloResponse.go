package enterprise

import "encoding/xml"

type HelloResponse struct {
	XMLName xml.Name `xml:"helloResponse"`

	Return HelloResponseData `xml:"return"`
}

type HelloResponseData struct {
	Id      int64  `xml:"id"`
	Message string `xml:"message"`
}

func NewHelloResponse() *HelloResponse {
	return &HelloResponse{}
}
