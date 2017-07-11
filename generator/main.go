package main

import (
	"fmt"
	"gowsdl/generator/wsdl"
	//"encoding/xml"
)

func main() {
	definitions, err := wsdl.NewDefinitionsFromUrl("http://kf.egtcp.com:8080/gttown-enterprise-soa/ws/hello?wsdl")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, element := range definitions.Types.Schema.Element {
		fmt.Printf("name: %s, type: %s\n", element.Name, element.Type)
	}
	println(definitions)
}
