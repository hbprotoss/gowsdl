package enterprise

import "encoding/xml"

type HelloRequest struct {
	XMLName xml.Name `xml:"ns2:hello"`
	ns      string   `xml:"-"`

	Message string `xml:"message"`
}

func NewHelloRequest(namespace string) *HelloRequest {
	return &HelloRequest{
		ns: namespace,
	}
}

func (req *HelloRequest) Namespace() string {
	return req.ns
}
