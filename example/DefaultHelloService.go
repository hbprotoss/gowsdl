package example
import (
	"fmt"
	"wsdl2go/soap/client"
	"wsdl2go/soap/req"
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

func (s *DefaultHelloService) HelloMulti(message string, entity *Entity, ids []int32, list []int32) (err error) {
	var envelope = req.NewEnvelope()

	var request = NewHelloMulti(s.Namespace)
	request.Message = message
	request.Entity = entity
	request.Ids = ids
	request.List = list
	envelope.Body.Content = request

	var response = NewHelloMultiResponse(s.Namespace)

	err = s.soapClient.Call("helloMulti", request, response)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *DefaultHelloService) HelloList(messages []string) (entity []*Entity, err error) {
	var envelope = req.NewEnvelope()

	var request = NewHelloList(s.Namespace)
	request.Messages = messages
	envelope.Body.Content = request

	var response = NewHelloListResponse(s.Namespace)

	err = s.soapClient.Call("helloList", request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return response.Return, nil
}

func (s *DefaultHelloService) Hello(message string) (entity *Entity, err error) {
	var envelope = req.NewEnvelope()

	var request = NewHello(s.Namespace)
	request.Message = message
	envelope.Body.Content = request

	var response = NewHelloResponse(s.Namespace)

	err = s.soapClient.Call("hello", request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return response.Return, nil
}

