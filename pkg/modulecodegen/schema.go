package modulecodegen

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

const schemaImport = protogen.GoImportPath("github.com/fdymylja/tmos/runtime/orm/schema")

var schemaDefinition = schemaImport.Ident("Definition")

// genSchema generates the schema.Schema for the object
func genSchema(g *protogen.GeneratedFile, object *protogen.Message) error {
	g.Import(schemaImport)
	// get options
	stateObjectDesc, err := getStateObjectDesc(object.Desc)
	if err != nil {
		return err
	}

	if stateObjectDesc.PrimaryKey == "" && !stateObjectDesc.Singleton {
		return fmt.Errorf("invalid protobuf message at %s identifies itself as state transition but has not a primary key or singleton", object.Location.SourceFile)
	}

	// write the schema
	g.P("var ", object.GoIdent, "Schema = &", schemaDefinition, "{")
	if stateObjectDesc.PrimaryKey != "" {
		g.P("PrimaryKey: \"", stateObjectDesc.PrimaryKey, "\"", ",")
	}
	if stateObjectDesc.Singleton {
		g.P("Singleton: true,")
	}
	if len(stateObjectDesc.SecondaryKeys) != 0 {
		g.P("SecondaryKeys: []string{\"", strings.Join(stateObjectDesc.SecondaryKeys, "\",\""), "\"},")
	}
	g.P("}")
	g.P()
	return nil
}
