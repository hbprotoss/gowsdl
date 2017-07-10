package main

import (
	"gowsdl/enterprise"
	"gowsdl/soap/req"
	"gowsdl/soap/client"
	"fmt"
)

func main() {
	var envelope = req.NewEnvelopeWithSecurity("client", "GT666lucknumber")
	//var req = soap.NewGetEnterpriseRequest()
	//req.UserId = 35910
	//req.Local = "EN"
	var request = enterprise.NewHelloRequest()
	request.Message = "hahaha"

	envelope.Body.Content = request

	var soapClient = client.NewSOAPClientWithWsse(
		"http://kf.egtcp.com:8080/gttown-enterprise-soa/ws/hello",
		&client.SecurityAuth{
			Username: "client",
			Password: "GT666lucknumber",
			Type: "PasswordText",
		},
	)
	var response = enterprise.NewHelloResponse()

	err := soapClient.Call("hello", request, response)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(response)
}
