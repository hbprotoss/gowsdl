package example
type HelloMultiResponse struct {
	ns      string   `xml:"-"`

	
}

func NewHelloMultiResponse(namespace string) *HelloMultiResponse {
	return &HelloMultiResponse{
		ns: namespace,
	}
}

func (req *HelloMultiResponse) Namespace() string {
	return req.ns
}