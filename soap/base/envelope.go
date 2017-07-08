package base

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	SoapEnv string `xml:"xmlns:soapenv,attr"`
	Ser string `xml:"xmlns:ser,attr"`

	Header *Header
	Body *Body
}

func NewEnvelope() *Envelope {
	return &Envelope{
		Body: &Body{},
	}
}

func NewEnvelopeWithSecurity(username, password string) *Envelope {
	return &Envelope{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Ser: "http://service.enterprise.soa.gttown.com/",
		Header: NewHeaderWithSecurity(username, password),
		Body: &Body{},
	}
}
