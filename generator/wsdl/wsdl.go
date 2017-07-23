package wsdl

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Definitions struct {
	Import   *Import   `xml:"import"`
	Types    *Types    `xml:"types"`
	PortType *PortType `xml:"portType"`
	Binding  *Binding  `xml:"binding"`
	Service  *Service  `xml:"service"`

	TargetNamespace string `xml:"targetNamespace,attr"`
}

type Import struct {
	Location string `xml:"location,attr"`
}

type Types struct {
	Schema *Schema `xml:"schema"`
}

type Schema struct {
	Element     []*Element     `xml:"element"`
	ComplexType []*ComplexType `xml:"complexType"`
}

type ComplexType struct {
	Name     string    `xml:"name,attr"`
	Sequence *Sequence `xml:"sequence"`

	IsRequestType bool
}

type Sequence struct {
	Element []*Element `xml:"element"`
}

type Element struct {
	Name      string `xml:"name,attr"`
	Type      string `xml:"type,attr"`
	MinOccurs string `xml:"minOccurs,attr"`
	MaxOccurs string `xml:"maxOccurs,attr"`
}

type Binding struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

type Operation struct {
	Name   string  `xml:"name,attr"`
	Input  *Input  `xml:"input"`
	Output *Output `xml:"output"`
}

type Input struct {
	Message string `xml:"message,attr"`
	Name    string `xml:"name,attr"`
}

type Output struct {
	Message string `xml:"message,attr"`
	Name    string `xml:"name,attr"`
}

type Service struct {
	Name string `xml:"name,attr"`
	Port *Port  `xml:"port"`
}

type Port struct {
	Binding string `xml:"binding"`
	Name    string `xml:"name,attr"`
}

type PortType struct {
	Name      string       `xml:"name,attr"`
	Operation []*Operation `xml:"operation"`
}

func NewDefinitionsFromUrl(url string) (*Definitions, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	rawbody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var definitions = &Definitions{}
	err = xml.Unmarshal([]byte(rawbody), definitions)
	if err != nil {
		return nil, err
	}

	if definitions.Import != nil && definitions.Import.Location != "" {
		importDef, err := NewDefinitionsFromUrl(definitions.Import.Location)
		if err != nil {
			return nil, err
		}
		if importDef.Service != nil {
			definitions.Service = importDef.Service
		}
		if importDef.Binding != nil {
			definitions.Binding = importDef.Binding
		}
		if importDef.Types != nil {
			definitions.Types = importDef.Types
		}
		if importDef.PortType != nil {
			definitions.PortType = importDef.PortType
		}
		if importDef.TargetNamespace != "" {
			definitions.TargetNamespace = importDef.TargetNamespace
		}
	}
	return definitions, nil
}
