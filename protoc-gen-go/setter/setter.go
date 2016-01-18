package setter

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func init() {
	generator.RegisterPlugin(&SetterPlugin{})
}

type SetterPlugin struct {
	g *generator.Generator
}

// Name identifies the plugin.
func (s *SetterPlugin) Name() string {
	return "setter"
}

// Init is called once after data structures are built but before
// code generation begins.
func (s *SetterPlugin) Init(g *generator.Generator) {
	s.g = g
}

// Generate produces the code generated by the plugin for this file,
// except for the imports, by calling the generator's methods P, In, and Out.
func (s *SetterPlugin) Generate(file *generator.FileDescriptor) {
	for _, msg := range file.GetMessageType() {
		ccName := generator.CamelCase(msg.GetName())
		for _, f := range msg.GetField() {
			ccFieldName := generator.CamelCase(f.GetName())
			typ, _ := s.g.GoType(nil, f)
			s.g.P("func (m *", ccName, ") Set", ccFieldName, "(v ", typ, ") {")
			s.g.P("  m.", ccFieldName, "= v")
			s.g.P("}")
			s.g.P()
		}
	}

}

func (s *SetterPlugin) GenerateImports(file *generator.FileDescriptor) {}
