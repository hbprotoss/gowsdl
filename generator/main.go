package main

import (
	"bytes"
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

	var packageName = filepath.Base(sourcePath)

	var url = os.Args[1]
	definitions, err := wsdl.NewDefinitionsFromUrl(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	var elementMapping = wsdl.NewElementMappingFromDefinitions(definitions)
	fmt.Println(elementMapping)

	for _, complexType := range definitions.Types.Schema.ComplexType {

		fmt.Printf("name: %s\n", complexType.Name)
		file, err := os.Create(fmt.Sprintf("%s%s%s.go", sourcePath, string(os.PathSeparator), complexType.Name))
		if err != nil {
			fmt.Printf("Init() error: %v\n", err)
			return
		}

		data := make(map[string]string)
		data["package"] = packageName
		data["name"] = complexType.Name
		data["fieldName"] = util.FirstLetterToUpper(complexType.Name)
		data["members"] = generateMembers(complexType.Sequence)
		err = entityTpl.Execute(file, data)
		file.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func generateMembers(sequence *wsdl.Sequence) string {
	var buffer bytes.Buffer
	for _, element := range sequence.Element {
		if element.Name == "" {
			fmt.Println("Element name is empty")
			continue
		}
		var fieldName = util.FirstLetterToUpper(element.Name)
		var member = fmt.Sprintf(
			"\t%s %s `xml:\"%s\"`\n",
			fieldName,
			wsdl.GetType(element.Type),
			element.Name,
		)
		buffer.WriteString(member)
	}
	return buffer.String()
}
