package rpc

import (
	"bytes"
	"html/template"
	"log"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func init() {
	generator.RegisterPlugin(new(netRpcPlugin))
}

type netRpcPlugin struct {
	*generator.Generator
}

func (p *netRpcPlugin) Name() string {
	return "netRpcPlugin"
}

func (p *netRpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *netRpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

func (p *netRpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *netRpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P("//TODO: import code")
}

func (p *netRpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	spec := p.buildServiceSpec(svc)

	var buf bytes.Buffer
	if err := template.Must(template.New("").Parse("")).Execute(&buf, spec); err != nil {
		log.Fatal(err)
	}

	p.P(buf.String())
}

func (p *netRpcPlugin) buildServiceSpec(svc *descriptor.ServiceDescriptorProto) *ServiceSpec {
	spec := &ServiceSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}

	for _, m := range svc.Method {
		spec.Methods = append(spec.Methods, ServiceMethodSpec{
			MethodName:     generator.CamelCase(m.GetName()),
			InputTypeName:  p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName: p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
	}
}

type ServiceSpec struct {
	ServiceName string
	Methods     []ServiceMethodSpec
}

type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
}
