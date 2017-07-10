package req

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	SoapEnv string `xml:"xmlns:soap,attr"`
	Namespace string `xml:"xmlns:ns2,attr"`

	Header *Header
	Body *Body
}

func NewEnvelope() *Envelope {
	return &Envelope{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: &Body{},
	}
}

func NewEnvelopeWithSecurity(username, password string) *Envelope {
	return &Envelope{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Header: NewHeaderWithSecurity(username, password),
		Body: &Body{},
	}
}
