package req

import "encoding/xml"

type Header struct {
	XMLName xml.Name `xml:"soap:Header"`

	SecurityHeader *SecurityHeader
}

func NewHeader() *Header {
	return &Header{}
}

func NewHeaderWithSecurity(username, password string) *Header {
	return &Header{
		SecurityHeader: NewSecurityHeader(username, password, "PasswordText"),
	}
}

type SecurityHeader struct {
	XMLName        xml.Name `xml:"wsse:Security"`
	Wsse           string   `xml:"xmlns:wsse,attr"`
	Wsu            string   `xml:"xmlns:wsu,attr"`
	MustUnderstand string   `xml:"soap:mustUnderstand,attr"`

	UsernameToken UsernameToken
}

type UsernameToken struct {
	XMLName xml.Name `xml:"wsse:UsernameToken"`

	Username string   `xml:"wsse:Username"`
	Password Password `xml:"wsse:Password"`
}

type Password struct {
	Type  string `xml:"Type,attr"`
	Value string `xml:",chardata"`
}

func NewSecurityHeader(username, password, headerType string) *SecurityHeader {
	return &SecurityHeader{
		Wsse:           "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
		Wsu:            "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
		MustUnderstand: "1",
		UsernameToken: UsernameToken{
			Username: username,
			Password: Password{
				Type:  "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#" + headerType,
				Value: password,
			},
		},
	}
}
