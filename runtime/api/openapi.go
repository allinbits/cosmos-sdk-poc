package api

import (
	"fmt"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/pkg/protoutils/forge"
	"github.com/fdymylja/tmos/pkg/protoutils/oas3schema"
	v3 "github.com/googleapis/gnostic/openapiv3"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

type openAPI struct {
	gen *oas3schema.OpenAPIv3Generator
}

func NewOpenAPIBuilder() *openAPI {
	return &openAPI{
		gen: oas3schema.NewOpenAPIv3Generator(),
	}
}

func (o openAPI) AddSingleton(obj meta.StateObject, path string) error {
	opID := fmt.Sprintf("singleton.%s", meta.Name(obj))
	comment := fmt.Sprintf("Returns the unique instance of the %s object if it exists.", meta.Name(obj))
	err := o.gen.AddRawOperation("GET", opID, comment, path, "", obj)
	if err != nil {
		return err
	}

	err = o.gen.AddRequiredMessage(obj.ProtoReflect().Descriptor())
	if err != nil {
		return err
	}

	return nil
}

func (o openAPI) AddObject(obj meta.StateObject, singlePath string, listPath string) error {
	opID := fmt.Sprintf("get.%s", meta.Name(obj))
	comment := fmt.Sprintf("Returns an instance of %s", meta.Name(obj))
	err := o.gen.AddRawOperation("GET", opID, comment, singlePath, "", obj)
	if err != nil {
		return err
	}
	err = o.gen.AddRequiredMessage(obj.ProtoReflect().Descriptor())
	if err != nil {
		return err
	}
	// forge list object
	listObject, err := forge.List(obj, protoregistry.GlobalFiles)
	if err != nil {
		return err
	}
	listOPID := fmt.Sprintf("list.%s", meta.Name(obj))
	listComment := fmt.Sprintf("Returns a list of %s", meta.Name(obj))
	err = o.gen.AddRawOperation("GET", listOPID, listComment, listPath, "", dynamicpb.NewMessage(listObject.Descriptor()))
	if err != nil {
		return err
	}
	err = o.gen.AddRequiredMessage(listObject.Descriptor())
	if err != nil {
		return err
	}
	return nil
}

func (o openAPI) Build() (*v3.Document, error) {
	doc, err := o.gen.Build()
	if err != nil {
		return nil, err
	}
	// after building the document we need to
	return doc, nil
}
