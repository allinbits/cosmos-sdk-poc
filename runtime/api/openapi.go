package api

import (
	"fmt"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/pkg/protoutils/oas3schema"
	v3 "github.com/googleapis/gnostic/openapiv3"
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
	// TODO listPath

	err = o.gen.AddRequiredMessage(obj.ProtoReflect().Descriptor())
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
	return doc, nil
}
