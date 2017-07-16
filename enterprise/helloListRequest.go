package enterprise

import "encoding/xml"

type HelloListRequest struct {
	XMLName xml.Name `xml:"ns2:hello"`
	ns      string   `xml:"-"`

	Message []string `xml:"message"`
}

func NewHelloListRequest(namespace string) *HelloListRequest {
	return &HelloListRequest{
		ns: namespace,
	}
}

func (req *HelloListRequest) Namespace() string {
	return req.ns
}
