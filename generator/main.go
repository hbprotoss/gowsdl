package main

import (
	"fmt"
	"gowsdl/generator/util"
	"gowsdl/generator/wsdl"
	"os"
	"path/filepath"
	"text/template"
)

var (
	entityTpl *template.Template
)

func Init() (err error) {
	entityTpl, err = template.New("entityTpl").Parse(wsdl.EntityTplText)
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

	generateEntity(definitions.Types.Schema.ComplexType, sourcePath)
}

func generateEntity(complexTypes []wsdl.ComplexType, sourcePath string) {
	var packageName = filepath.Base(sourcePath)
	for _, complexType := range complexTypes {

		fmt.Printf("name: %s\n", complexType.Name)
		file, err := os.Create(fmt.Sprintf("%s%s%s.go", sourcePath, string(os.PathSeparator), complexType.Name))
		if err != nil {
			fmt.Printf("Init() error: %v\n", err)
			return
		}

		data := make(map[string]interface{})
		data["package"] = packageName
		data["name"] = complexType.Name
		data["structName"] = util.FirstLetterToUpper(complexType.Name)
		data["members"] = generateMembers(complexType.Sequence)
		err = entityTpl.Execute(file, data)
		file.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func generateMembers(sequence *wsdl.Sequence) (fields []wsdl.StructField) {
	fields = make([]wsdl.StructField, len(sequence.Element))
	for index, element := range sequence.Element {
		if element.Name == "" {
			fmt.Println("Element name is empty")
			continue
		}
		var field = wsdl.StructField{
			FieldName: util.FirstLetterToUpper(element.Name),
			FieldType: wsdl.GetType(element.Type),
			XmlName:   element.Name,
		}
		fields[index] = field
	}
	return
}
