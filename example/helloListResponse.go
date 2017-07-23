package example
type HelloListResponse struct {
	ns      string   `xml:"-"`

	Return []*Entity `xml:"return"`
	
}

func NewHelloListResponse(namespace string) *HelloListResponse {
	return &HelloListResponse{
		ns: namespace,
	}
}

func (req *HelloListResponse) Namespace() string {
	return req.ns
}