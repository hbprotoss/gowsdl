package main

import (
	"fmt"
	"strings"
	"wsdl2go/wsdl"
	"path/filepath"
	"os"
	"text/template"
	"wsdl2go/util"
)

func generateInterface(definitions *wsdl.Definitions, mapping *wsdl.ElementMapping, sourceRoot string) (err error) {
	var portType = definitions.PortType
	var sourcePath = fmt.Sprintf("%s%s%s.go", sourceRoot, string(os.PathSeparator), portType.Name)
	if err = gatherInterfaceInfo(definitions, mapping, sourcePath, interfaceTpl); err != nil {
		fmt.Printf("%v\n", err)
	}
	return
}

func generateInterfaceImpl(definitions *wsdl.Definitions, mapping *wsdl.ElementMapping, sourceRoot string) (err error) {
	var portType = definitions.PortType
	var sourcePath = fmt.Sprintf("%s%sDefault%s.go", sourceRoot, string(os.PathSeparator), portType.Name)
	if err := gatherInterfaceInfo(definitions, mapping, sourcePath, implTpl); err != nil {
		fmt.Printf("%v\n", err)
	}
	return
}

func gatherInterfaceInfo(definitions *wsdl.Definitions, mapping *wsdl.ElementMapping, sourcePath string, tpl *template.Template) (err error) {
	var portType = definitions.PortType
	var packageName = filepath.Base(filepath.Dir(sourcePath))
	fmt.Printf("Interface name: %s\n", portType.Name)
	file, err := os.Create(sourcePath)
	if err != nil {
		fmt.Printf("generateEntity() error: %v\n", err)
		return
	}
	data := make(map[string]interface{})
	data["package"] = packageName
	data["namespace"] = definitions.TargetNamespace
	data["serviceName"] = portType.Name
	var methods = make([]*wsdl.ServiceMethod, len(portType.Operation))
	for index, operation := range portType.Operation {
		methods[index] = generateInterfaceMethod(operation, mapping)
	}
	data["methods"] = methods
	return tpl.Execute(file, data)
}

func generateInterfaceMethod(operation *wsdl.Operation, mapping *wsdl.ElementMapping) (method *wsdl.ServiceMethod) {
	method = &wsdl.ServiceMethod{
		Name: util.FirstLetterToUpper(operation.Name),
	}
	var inputTypeName = util.GetEntityName(operation.Input.Message)
	var inputType = wsdl.GetType(operation.Input.Message)
	if inputType == "" {
		var complexType = mapping.ComplexType[inputTypeName]
		if complexType == nil {
			return nil
		}
		method.ParamsString = generateTypeDefs(complexType.Sequence)
		method.ParamNames = generateParams(complexType.Sequence)
	}

	var outputTypeName = util.GetEntityName(operation.Output.Message)
	var outputType = wsdl.GetType(operation.Output.Message)
	if outputType == "" {
		var complexType = mapping.ComplexType[outputTypeName]
		if complexType == nil {
			return nil
		}
		method.ReturnsString = generateTypeDefs(complexType.Sequence)
	}
	return
}

func generateParams(sequence *wsdl.Sequence) []string {
	params := make([]string, len(sequence.Element))
	for index, element := range sequence.Element {
		params[index] = element.Name
	}
	return params
}

func generateTypeDefs(sequence *wsdl.Sequence) string {
	params := make([]string, len(sequence.Element))
	for index, element := range sequence.Element {
		var elementType = util.GetEntityName(element.Type)
		paramType := wsdl.GetType(elementType)
		if paramType == "" {
			paramType = util.FirstLetterToUpper(elementType)
			if element.MaxOccurs == "unbounded" {
				paramType = "[]*" + paramType
			} else {
				paramType = "*" + paramType
			}
		} else {
			if element.MaxOccurs == "unbounded" {
				paramType = "[]" + paramType
			}
		}
		var name = ""
		if strings.Compare(element.Name, "return") == 0 {
			name = elementType
		} else {
			name = element.Name
		}
		params[index] = fmt.Sprintf("%s %s", name, paramType)
	}
	return strings.Join(params, ", ")
}

func generateEntity(definitions *wsdl.Definitions, sourceRoot string) (error) {
	var complexTypes = definitions.Types.Schema.ComplexType
	var packageName = filepath.Base(sourceRoot)
	for _, complexType := range complexTypes {

		fmt.Printf("Entity name: %s\n", complexType.Name)
		file, err := os.Create(fmt.Sprintf("%s%s%s.go", sourceRoot, string(os.PathSeparator), complexType.Name))
		if err != nil {
			fmt.Printf("generateEntity() error: %v\n", err)
			return err
		}

		data := make(map[string]interface{})
		data["package"] = packageName
		data["name"] = complexType.Name
		data["structName"] = util.FirstLetterToUpper(complexType.Name)
		data["members"] = generateEntityMembers(complexType.Sequence)
		data["isRequestEntity"] = complexType.IsRequestType
		err = entityTpl.Execute(file, data)
		file.Close()
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func generateEntityMembers(sequence *wsdl.Sequence) (members []wsdl.EntityMember) {
	members = make([]wsdl.EntityMember, len(sequence.Element))
	for index, element := range sequence.Element {
		if element.Name == "" {
			fmt.Println("Element name is empty")
			continue
		}
		var member = wsdl.EntityMember{
			Name:    util.FirstLetterToUpper(element.Name),
			Type:    wsdl.GetTypeWithUpperEntity(element.Type),
			XmlName: element.Name,
		}
		if element.MaxOccurs == "unbounded" {
			member.Type = "[]" + member.Type
		}
		members[index] = member
	}
	return
}
