package client


import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
	"gowsdl/soap/req"
	"errors"
	"strings"
	"gowsdl/soap/resp"
	"fmt"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type BasicAuth struct {
	Login    string
	Password string
}

type SecurityAuth struct {
	Username string
	Password string
	Type string
}

type SOAPClient struct {
	url  string
	tls  bool
	auth *BasicAuth
	securityAuth *SecurityAuth
}

func NewSOAPClientWithWsse(url string, auth *SecurityAuth) *SOAPClient  {
	return &SOAPClient{
		url: url,
		tls: false,
		securityAuth: auth,
	}
}

func (s *SOAPClient) Call(soapAction string, request req.Request, response interface{}) error {
	var envelope = &req.Envelope{}
	if s.securityAuth != nil {
		envelope = req.NewEnvelopeWithSecurity(s.securityAuth.Username, s.securityAuth.Password)
	}
	envelope.Namespace = request.Namespace()

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	log.Println(buffer.String())

	httpReq, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		httpReq.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	httpReq.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	//if soapAction != "" {
	//	req.Header.Add("SOAPAction", soapAction)
	//}

	httpReq.Header.Set("User-Agent", "gowsdl/0.1")
	httpReq.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	httpClient := &http.Client{Transport: tr}
	res, err := httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	log.Println(string(rawbody))
	respEnvelope := resp.NewEnvelope()
	respEnvelope.Body = &resp.Body{Content: response}
	err = xml.Unmarshal([]byte(parseSoapData(string(rawbody))), respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return errors.New(fmt.Sprintf("Soap Error: code(%s), message(%s)", fault.Code, fault.String))
	}

	return nil
}

func parseSoapData(rawbody string) string {
	var begin = strings.Index(rawbody, "<soap")
	if begin == -1 {
		return rawbody
	}

	var end = strings.LastIndex(rawbody, "--uuid")
	if end == -1 {
		return rawbody
	}

	return rawbody[begin:end]
}
