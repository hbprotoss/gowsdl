package enterprise

import (
	"fmt"
	"gowsdl/soap/client"
	"gowsdl/soap/req"
)

type DefaultHelloService struct {
	Namespace string
}

func NewHelloService() *DefaultHelloService {
	return &DefaultHelloService{
		Namespace: "http://service.enterprise.soa.gttown.com/",
	}
}

func (s *DefaultHelloService) Hello(message string) (*HelloResponse, error) {
	var envelope = req.NewEnvelopeWithSecurity("client", "GT666lucknumber")
	var request = NewHelloRequest("http://service.enterprise.soa.gttown.com/")
	request.Message = message

	envelope.Body.Content = request

	var soapClient = client.NewSOAPClientWithWsse(
		"http://kf.egtcp.com:8080/gttown-enterprise-soa/ws/hello",
		&client.SecurityAuth{
			Username: "client",
			Password: "GT666lucknumber",
			Type:     "PasswordText",
		},
	)
	var response = NewHelloResponse()

	err := soapClient.Call("hello", request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return response, nil
}
