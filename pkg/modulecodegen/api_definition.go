package modulecodegen

import (
	"fmt"

	"github.com/fdymylja/tmos/core/modulegen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

const metaPkg = protogen.GoImportPath("github.com/fdymylja/tmos/core/meta")

func genAPIDefinition(g *protogen.GeneratedFile, message *protogen.Message) error {
	apiDefinitionIdent := metaPkg.Ident("APIDefinition")
	apiTypeStateObject := metaPkg.Ident("APIType_StateObject")
	apiTypeStateTransition := metaPkg.Ident("APIType_StateTransition")
	// check if object is something that we own
	opts := message.Desc.Options().(*descriptorpb.MessageOptions)
	stateObject := proto.GetExtension(opts, modulegen.E_StateObject).(*modulegen.StateObjectDescriptor)
	stateTransition := proto.GetExtension(opts, modulegen.E_StateTransition).(*modulegen.StateTransitionDescriptor)
	if stateObject == nil && stateTransition == nil {
		return nil
	}
	if stateObject != nil && stateTransition != nil {
		return fmt.Errorf("%s in file %s is both state transition and state object", message.GoIdent, message.Location.SourceFile)
	}

	g.Import(metaPkg)
	g.P("func (x *", message.GoIdent, ") APIDefinition() *", apiDefinitionIdent, " {")
	g.P("return ", apiDefinitionIdent, "{")
	g.P("Group: ", message.Desc.FullName().Parent(), ",")
	g.P("Kind: ", message.Desc.Name(), ",")
	switch {
	case stateObject != nil:
		g.P("ApiType: ", apiTypeStateObject, ",")
	case stateTransition != nil:
		g.P("ApiType: ", apiTypeStateTransition, ",")
	}
	g.P("}")
	g.P("}")
	g.P()

	return nil
}
