package router

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fdymylja/tmos/core/meta"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/pkg/protoutils/forge"
	"github.com/fdymylja/tmos/runtime/client"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func NewBuilder(c client.RuntimeClient) *Builder {
	return &Builder{
		client:     c,
		mux:        mux.NewRouter(),
		modules:    map[string]module.Descriptor{},
		knownPaths: nil,
		openAPI:    NewOpenAPIBuilder(),
	}
}

type Builder struct {
	client     client.RuntimeClient
	mux        *mux.Router
	modules    map[string]module.Descriptor
	knownPaths map[string]string // maps known paths of objects
	openAPI    *openAPI
}

// CreateModuleHandlers creates the handler given a module.Descriptor
func (s *Builder) CreateModuleHandlers(module module.Descriptor) error {
	_, exists := s.modules[module.Name]
	if exists {
		return fmt.Errorf("api: module already registered: %s", module.Name)
	}
	// register object handler
	for _, obj := range module.StateObjects {
		err := s.CreateStateObjectHandler(obj.StateObject, &runtimev1alpha1.SchemaDefinition{
			Singleton:     obj.Options.Singleton,
			PrimaryKey:    obj.Options.PrimaryKey,
			SecondaryKeys: obj.Options.SecondaryKeys,
		})
		if err != nil {
			return fmt.Errorf("api: unable to register %s for module %s", meta.Name(obj.StateObject), module.Name)
		}
	}
	return nil
}

// CreateStateObjectHandler creates a state object handler given the state object and its definition
func (s *Builder) CreateStateObjectHandler(object meta.StateObject, definition *runtimev1alpha1.SchemaDefinition) error {
	// get schema for the object
	sch, err := schema.NewSchema(object, schema.Definition{
		Singleton:     definition.Singleton,
		PrimaryKey:    definition.PrimaryKey,
		SecondaryKeys: definition.SecondaryKeys,
	})
	if err != nil {
		return err
	}
	// create handler based to, if singleton or not
	switch definition.Singleton {
	case true:
		def := object.APIDefinition()
		path := strings.ToLower(
			fmt.Sprintf("/%s/%s", def.Group, def.Kind),
		)
		s.mux.
			Methods(http.MethodGet).
			Path(path).
			HandlerFunc(newSingletonGetHandler(s.client, sch))

		// add to open API spec.
		err = s.openAPI.AddSingleton(object, path)
		if err != nil {
			return err
		}
	case false:
		// create get
		def := object.APIDefinition()
		singleInstancePath := strings.ToLower(
			fmt.Sprintf("/%s/%s/{%s}", def.Group, def.Kind, definition.PrimaryKey),
		)
		listInstancePath := strings.ToLower(
			fmt.Sprintf("/%s/%ss", def.Group, def.Kind), // TODO the plural name should be in the state object schema options
		)
		s.mux.
			Methods(http.MethodGet).
			Path(singleInstancePath).
			HandlerFunc(newGetHandler(s.client, sch, definition))
		// create list
		s.mux.Methods(http.MethodGet).
			Path(listInstancePath).
			HandlerFunc(newListHandler(s.client, sch))

		// add to open API spec.
		err = s.openAPI.AddObject(object, singleInstancePath, listInstancePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Builder) Build() (*mux.Router, error) {
	doc, err := s.openAPI.Build()
	if err != nil {
		return nil, err
	}
	b, err := doc.YAMLValue("test")
	if err != nil {
		return nil, err
	}
	s.mux.HandleFunc("/spec", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(b)
	})

	return s.mux, nil
}

func newSingletonGetHandler(c client.RuntimeClient, schema *schema.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newObj := schema.NewStateObject()
		err := c.Get(meta.SingletonID, newObj)
		if err != nil {
			notFound(w, "not found")
			return
		}
		writeObject(w, newObj)
	}
}

// newGetHandler creates an http.HandlerFunc that can be used to fetch a state object
func newGetHandler(c client.RuntimeClient, schema *schema.Schema, definition *runtimev1alpha1.SchemaDefinition) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		primaryKeyValue, exists := vars[definition.PrimaryKey]
		// if the user didn't set the primary key in the url variables
		if !exists {
			badRequest(w, "missing object key in url path")
			return
		}
		// encode primary key
		pkBytes, err := schema.EncodePrimaryKeyString(primaryKeyValue)
		if err != nil {
			badRequest(w, "bad primary key format in url path: %s", err)
			return
		}
		newObj := schema.NewStateObject()
		err = c.Get(meta.NewBytesID(pkBytes), newObj)
		if err != nil {
			notFound(w, err.Error())
			return
		}
		writeObject(w, newObj)
	}
}

// newListHandler creates an http.HandlerFunc that can be used to fetch a list of state objects of the same kind
func newListHandler(c client.RuntimeClient, schema *schema.Schema) http.HandlerFunc {
	listObject, err := forge.List(schema.NewStateObject(), protoregistry.GlobalFiles)
	if err != nil {
		panic(err)
	}
	listFd := listObject.Descriptor().Fields().Get(0)
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		listOptions := new(ListQueryParams)
		err := listOptions.UnmarshalURLValues(q)
		if err != nil {
			badRequest(w, "bad query: %s", err)
			return
		}
		opts := []client.ListOption{
			client.ListRange(listOptions.Start, listOptions.End),
		}

		for _, selection := range listOptions.SelectFields {
			sp := strings.SplitN(selection, "=", 2)
			if len(sp) != 2 {
				badRequest(w, "bad fieldSelection in query %s", selection)
				return
			}
			// check that index exists
			_, err := schema.Indexer(sp[0])
			if err != nil {
				badRequest(w, "bad fieldSelection in query: %s", err)
				return
			}
			opts = append(opts, client.ListMatchFieldString(sp[0], sp[1]))
		}

		iter, err := c.List(schema.NewStateObject(), opts...)
		if err != nil {
			badRequest(w, "unable to list any object: %s", err)
			return
		}
		defer iter.Close()

		list := listObject.New()
		listValue := list.NewField(listFd).List()

		for iter.Valid() {
			obj := schema.NewStateObject()
			err := iter.Get(obj)
			if err != nil {
				badRequest(w, "unable to list object: %s", err)
				return
			}
			listValue.Append(protoreflect.ValueOfMessage(obj.ProtoReflect()))
			iter.Next()
		}
		list.Set(listFd, protoreflect.ValueOfList(listValue))
		writeObject(w, list.Interface())
	}
}

func writeObject(w io.Writer, obj proto.Message) bool {
	b, err := protojson.Marshal(obj)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(b)
	if err != nil {
		return false
	}
	return true
}

func notFound(writer http.ResponseWriter, format string, values ...interface{}) {
	writer.WriteHeader(http.StatusNotFound)
	_, _ = writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func badRequest(writer http.ResponseWriter, format string, values ...interface{}) {
	writer.WriteHeader(http.StatusBadRequest)
	_, _ = writer.Write([]byte(fmt.Sprintf(format, values...)))
}
