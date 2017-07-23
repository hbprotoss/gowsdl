package example
import "encoding/xml"
type HelloList struct {
	XMLName xml.Name `xml:"ns2:helloList"`
	ns      string   `xml:"-"`

	Messages []string `xml:"messages"`
	
}

func NewHelloList(namespace string) *HelloList {
	return &HelloList{
		ns: namespace,
	}
}

func (req *HelloList) Namespace() string {
	return req.ns
}