package api

import (
	"net/http"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/pkg/protoutils/oas3schema"
	"github.com/getkin/kin-openapi/openapi3"
)

type openAPI struct {
	document *openapi3.T
}

func NewOpenAPIBuilder() *openAPI {
	document := new(openapi3.T)
	document.Info = &openapi3.Info{
		Title:          "Starport Framework API Server",
		Description:    "",
		TermsOfService: "",
		Contact:        nil,
		License:        nil,
		Version:        "0.0.0", // TODO: version should be fetched automatically from commit
	}
	document.OpenAPI = "3.0.0"
	return &openAPI{document: document}
}

func (o openAPI) AddSingleton(obj meta.StateObject) error {
	op := openapi3.NewOperation()
	objSchema, err := oas3schema.FromMessageDescriptor(obj.ProtoReflect().Descriptor())
	if err != nil {
		return err
	}
	resp := openapi3.NewResponse()
	resp.Content = openapi3.NewContentWithJSONSchema(objSchema)
	op.AddResponse(http.StatusOK, resp)
	o.document.AddOperation("/path", http.MethodGet, op)

	return nil
}
