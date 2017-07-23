package main

import (
	"fmt"
	"wsdl2go/util"
	"wsdl2go/wsdl"
	"os"
	"text/template"
)

var (
	entityTpl    *template.Template
	interfaceTpl *template.Template
	implTpl      *template.Template
)

func Init() (err error) {
	funcMap := template.FuncMap{
		"FirstLetterToUpper": util.FirstLetterToUpper,
		"FirstLetterToLower": util.FirstLetterToLower,
	}
	entityTpl, err = template.New("entityTpl").Funcs(funcMap).Parse(wsdl.EntityTplText)
	if err != nil {
		return
	}
	interfaceTpl, err = template.New("interfaceTpl").Funcs(funcMap).Parse(wsdl.InterfaceTplText)
	if err != nil {
		return
	}
	implTpl, err = template.New("implTpl").Funcs(funcMap).Parse(wsdl.ImplementationTplText)
	if err != nil {
		return
	}
	return nil
}

func main() {
	if err := Init(); err != nil {
		fmt.Printf("Init() error: %v\n", err)
		return
	}

	var sourcePath = os.Args[2]
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		os.Mkdir(sourcePath, os.ModePerm)
	}

	var url = os.Args[1]
	definitions, err := wsdl.NewDefinitionsFromUrl(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	var elementMapping = wsdl.NewElementMappingFromDefinitions(definitions)
	fmt.Println(elementMapping)

	if err = generateEntity(definitions, sourcePath); err != nil {
		fmt.Println("Failed to generateEntity")
		return
	}
	if err = generateInterface(definitions, elementMapping, sourcePath); err != nil {
		fmt.Println("Failed to generateInterface")
		return
	}
	if err = generateInterfaceImpl(definitions, elementMapping, sourcePath); err != nil {
		fmt.Println("Failed to generateInterfaceImpl")
		return
	}
}


