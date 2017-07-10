package resp

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`

	Body *Body
}

func NewEnvelope() *Envelope {
	return &Envelope{
		Body: &Body{},
	}
}
