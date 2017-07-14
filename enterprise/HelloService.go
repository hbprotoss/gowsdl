package enterprise

type HelloService interface {
	Hello(message string) (*HelloResponse, error)
	HelloList(messages []string)
}
