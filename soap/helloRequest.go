package soap

import "encoding/xml"

type HelloRequest struct {
	XMLName xml.Name `xml:"ns2:hello"`

	Message string `xml:"message"`

}

func NewHelloRequest() *HelloRequest {
	return &HelloRequest{}
}

func (req *HelloRequest) Namespace() string {
	return "http://service.server.soa.demo.hbprotoss.io/"
}