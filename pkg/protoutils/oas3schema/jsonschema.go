package oas3schema

import (
	"github.com/getkin/kin-openapi/openapi3"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func FromMessageDescriptor(md protoreflect.MessageDescriptor) (*openapi3.Schema, error) {
	schema := openapi3.NewSchema()
	schema.Type = (string)(md.FullName())
	panic("implement me")
}
