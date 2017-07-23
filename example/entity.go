package example
type Entity struct {
	ns      string   `xml:"-"`

	Id int64 `xml:"id"`
	Message string `xml:"message"`
	
}

func NewEntity(namespace string) *Entity {
	return &Entity{
		ns: namespace,
	}
}

func (req *Entity) Namespace() string {
	return req.ns
}