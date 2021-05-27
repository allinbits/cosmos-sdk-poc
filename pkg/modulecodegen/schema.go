package modulecodegen

import (
	"fmt"
	"strings"

	"github.com/fdymylja/tmos/core/modulegen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

const schemaImport = protogen.GoImportPath("github.com/fdymylja/tmos/runtime/orm/schema")
const metaImport = protogen.GoImportPath("github.com/fdymylja/tmos/runtime/meta")

var schemaDefinition = schemaImport.Ident("Definition")
var metaMeta = metaImport.Ident("Meta")

// genSchema generates the schema.Schema for the object
func genSchema(g *protogen.GeneratedFile, object *protogen.Message) error {
	g.Import(schemaImport)
	// get options
	opts := object.Desc.Options().(*descriptorpb.MessageOptions)
	isSingleTon := proto.GetExtension(opts, modulegen.E_Singleton).(bool)
	primaryKey := proto.GetExtension(opts, modulegen.E_PrimaryKey).(string)
	secondaryKeys := proto.GetExtension(opts, modulegen.E_SecondaryKey).([]string)

	if primaryKey == "" && !isSingleTon {
		return fmt.Errorf("invalid protobuf message at %s identifies itself as state transition but has not a primary key or singleton", object.Location.SourceFile)
	}

	// write the schema
	g.P("var ", object.GoIdent, "Schema = ", schemaDefinition, "{")
	// write meta
	g.P("Meta: ", metaMeta, "{")
	g.P("APIGroup: \"", object.Desc.FullName().Parent(), "\",")
	g.P("APIKind: \"", object.Desc.Name(), "\",")
	g.P("},")
	if primaryKey != "" {
		g.P("PrimaryKey: \"", primaryKey, "\"", ",")
	}
	if isSingleTon {
		g.P("Singleton: true,")
	}
	if len(secondaryKeys) != 0 {
		g.P("SecondaryKeys: []string{\"", strings.Join(secondaryKeys, "\",\""), "\"},")
	}
	g.P("}")
	g.P()
	return nil
}
