package main

import (
	"fmt"
	"gowsdl/generator/wsdl"
	"os"
	"text/template"
	"bytes"
	"unicode"
	"path/filepath"
)

var (
	requestTpl *template.Template
)


func Init() (err error) {
	requestTpl, err = template.New("requestTpl").Parse(wsdl.RequestTplText)
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
		data["fieldName"] = firstLetterToUpper(complexType.Name)
		data["members"] = generateMembers(complexType.Sequence)
		err = requestTpl.Execute(file, data)
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
		var fieldName = firstLetterToUpper(element.Name)
		var member = fmt.Sprintf(
			"\t%s %s `xml:\"%s\"`\n",
			fieldName,
			firstLetterToUpper(wsdl.GetType(element.Type)),
			element.Name,
		)
		buffer.WriteString(member)
	}
	return buffer.String()
}

func firstLetterToUpper(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}