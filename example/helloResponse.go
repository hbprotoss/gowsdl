package example
type HelloResponse struct {
	ns      string   `xml:"-"`

	Return *Entity `xml:"return"`
	
}

func NewHelloResponse(namespace string) *HelloResponse {
	return &HelloResponse{
		ns: namespace,
	}
}

func (req *HelloResponse) Namespace() string {
	return req.ns
}