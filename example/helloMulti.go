package example
import "encoding/xml"
type HelloMulti struct {
	XMLName xml.Name `xml:"ns2:helloMulti"`
	ns      string   `xml:"-"`

	Message string `xml:"message"`
	Entity *Entity `xml:"entity"`
	Ids []int32 `xml:"ids"`
	List []int32 `xml:"list"`
	
}

func NewHelloMulti(namespace string) *HelloMulti {
	return &HelloMulti{
		ns: namespace,
	}
}

func (req *HelloMulti) Namespace() string {
	return req.ns
}