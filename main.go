package main

import (
	"encoding/xml"
	"xmltest/soap"
	soapbase "xmltest/soap/base"
)

func main() {
	var envelope = soapbase.NewEnvelopeWithSecurity("client", "GT666lucknumber")
	var request = soap.NewGetEnterpriseRequest()
	request.UserId = 35910
	request.Local = "EN"

	envelope.Body.Content = request
	xmlText, err := xml.MarshalIndent(envelope, "", "\t")
	if err != nil {
		println(err)
		return
	}
	println(string(xmlText))

	var client = soapbase.NewSOAPClientWithWsse(
		"http://192.168.2.41:9040/gttown-enterprise-soa/ws/enterprise",
		&soapbase.SecurityAuth{
			Username: "client",
			Password: "GT666lucknumber",
			Type: "PasswordText",
		},
	)
	var response = &soap.GetEnterpriseResponse{}
	client.Call("getEnterprise", request, response)
	println(response)
}

