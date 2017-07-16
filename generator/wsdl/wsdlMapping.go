package wsdl

type ElementMapping struct {
	ComplexType map[string]*ComplexType
	Operation   map[string]*Operation
}

func NewElementMapping() *ElementMapping {
	return &ElementMapping{
		ComplexType: make(map[string]*ComplexType),
		Operation:   make(map[string]*Operation),
	}
}

func NewElementMappingFromDefinitions(definitions *Definitions) *ElementMapping {
	var mapping = NewElementMapping()
	for _, complexType := range definitions.Types.Schema.ComplexType {
		mapping.ComplexType[complexType.Name] = &complexType
	}
	for _, operation := range definitions.PortType.Operation {
		mapping.Operation[operation.Name] = &operation
	}
	return mapping
}
