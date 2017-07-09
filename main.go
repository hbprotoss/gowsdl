package main

import (
	"encoding/xml"
	"gowsdl/soap"
	soapbase "gowsdl/soap/base"
)

func main() {
	var envelope = soapbase.NewEnvelopeWithSecurity("client", "GT666lucknumber")
	//var request = soap.NewGetEnterpriseRequest()
	//request.UserId = 35910
	//request.Local = "EN"
	var request = soap.NewHelloRequest()

	envelope.Body.Content = request
	xmlText, err := xml.MarshalIndent(envelope, "", "\t")
	if err != nil {
		println(err)
		return
	}
	println(string(xmlText))

	var client = soapbase.NewSOAPClientWithWsse(
		"http://127.0.0.1:8080/ws/hello",
		&soapbase.SecurityAuth{
			Username: "client",
			Password: "GT666lucknumber",
			Type: "PasswordText",
		},
	)
	var response = soap.NewHelloResponse()
	xmlText, err = xml.MarshalIndent(response, "", "\t")
	if err != nil {
		println(err)
		return
	}
	println(string(xmlText))

	client.Call("hello", request, response)
	println(response)
}

