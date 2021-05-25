package modulecodegen

import (
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
)

func toLowerCamelCase(ident protogen.GoIdent) string {
	id := ident.GoName
	lower := unicode.ToLower((rune)(id[0]))
	return string(lower) + id[1:]
}
