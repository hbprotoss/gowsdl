package enterprise

import (
	"fmt"
	"gowsdl/soap/client"
	"gowsdl/soap/req"
)

type DefaultHelloService struct {
	Namespace  string
	soapClient *client.SOAPClient
}

func NewHelloService(url string, auth *client.SecurityAuth) *DefaultHelloService {
	return &DefaultHelloService{
		Namespace:  "http://service.server.soa.demo.hbprotoss.io/",
		soapClient: client.NewSOAPClientWithWsse(url, auth),
	}
}

func (s *DefaultHelloService) Hello(message string) (*HelloResponseData, error) {
	var envelope = req.NewEnvelope()

	var request = NewHelloRequest(s.Namespace)
	request.Message = message
	envelope.Body.Content = request

	var response = NewHelloResponse()

	err := s.soapClient.Call("hello", request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return response.Return, nil
}
