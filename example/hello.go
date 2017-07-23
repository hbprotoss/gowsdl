package example
import "encoding/xml"
type Hello struct {
	XMLName xml.Name `xml:"ns2:hello"`
	ns      string   `xml:"-"`

	Message string `xml:"message"`
	
}

func NewHello(namespace string) *Hello {
	return &Hello{
		ns: namespace,
	}
}

func (req *Hello) Namespace() string {
	return req.ns
}