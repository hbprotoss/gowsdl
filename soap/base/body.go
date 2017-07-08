package base

import "encoding/xml"

type Body struct {
	XMLName xml.Name `xml:"soapenv:Body"`

	Fault   *Fault      `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}
