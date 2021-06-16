package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/pkg/protoutils/forge"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"k8s.io/klog/v2"
)

func NewServer(store orm.Store) *Builder {
	return &Builder{
		server:  nil,
		store:   store,
		mux:     mux.NewRouter(),
		modules: map[string]module.Descriptor{},
		openAPI: NewOpenAPIBuilder(),
	}
}

type Builder struct {
	server     *http.Server
	store      orm.Store
	mux        *mux.Router
	modules    map[string]module.Descriptor
	knownPaths map[string]string // maps known paths of objects
	openAPI    *openAPI
}

// RegisterModuleAPI registers the API for the server
func (s *Builder) RegisterModuleAPI(module module.Descriptor) error {
	_, exists := s.modules[module.Name]
	if exists {
		return fmt.Errorf("api: module already registered: %s", module.Name)
	}
	// register object handler
	for _, obj := range module.StateObjects {
		err := s.registerStateObjectHandlers(module, obj)
		if err != nil {
			return fmt.Errorf("api: unable to register %s for module %s", meta.Name(obj.StateObject), module.Name)
		}
	}
	return nil
}

func (s *Builder) Start() {
	klog.Infof("starting api server...")
	doc, err := s.openAPI.Build()
	if err != nil {
		panic(err)
	}
	b, err := doc.YAMLValue("test")
	if err != nil {
		panic(err)
	}
	s.mux.HandleFunc("/spec", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(b)
	})
	go func() {
		s.server = &http.Server{
			Addr:    ":8080", // TODO configurable
			Handler: s.mux,
		}
		panic(s.server.ListenAndServe()) // TODO me better
	}()
}

func (s *Builder) loadStore(writer http.ResponseWriter, request *http.Request) (orm.Store, bool) {
	height, err := getHeight(request)
	if err != nil {
		badRequest(writer, "bad height value %d: %s", height, err)
		return orm.Store{}, false
	}
	switch height {
	case 0:
		store := s.store.LatestVersion()
		return store, true
	default:
		store, err := s.store.LoadVersion(height)
		if err != nil {
			notFound(writer, "store version %d not found", height)
			return orm.Store{}, false
		}
		return store, true
	}
}

func (s *Builder) registerStateObjectHandlers(descriptor module.Descriptor, obj module.StateObject) error {
	// get schema for the object
	sch, err := s.store.SchemaRegistry().GetByAPIDefinition(obj.StateObject.APIDefinition())
	if err != nil {
		return err
	}
	// create handler based to, if singleton or not
	switch obj.Options.Singleton {
	case true:
		def := obj.StateObject.APIDefinition()
		path := strings.ToLower(
			fmt.Sprintf("/%s/%s", def.Group, def.Kind),
		)
		s.mux.
			Methods(http.MethodGet).
			Path(path).
			HandlerFunc(newSingletonGetHandler(sch, s.loadStore))

		// add to open API spec.
		err = s.openAPI.AddSingleton(obj.StateObject, path)
		if err != nil {
			return err
		}
	case false:
		// create get
		def := obj.StateObject.APIDefinition()
		singleInstancePath := strings.ToLower(
			fmt.Sprintf("/%s/%s/{%s}", def.Group, def.Kind, obj.Options.PrimaryKey),
		)
		listInstancePath := strings.ToLower(
			fmt.Sprintf("/%s/%ss", def.Group, def.Kind), // TODO the plural name should be in the state object schema options
		)
		s.mux.
			Methods(http.MethodGet).
			Path(singleInstancePath).
			HandlerFunc(newGetHandler(sch, obj.Options, s.loadStore))
		// create list
		s.mux.Methods(http.MethodGet).
			Path(listInstancePath).
			HandlerFunc(newListHandler(sch, s.loadStore))

		// add to open API spec.
		err = s.openAPI.AddObject(obj.StateObject, singleInstancePath, listInstancePath)
		if err != nil {
			return err
		}
	}

	return nil
}

// loadStoreFunc returns the versioned store given an http request and returns if the loading succeeded or not
// if the loading fails loadStoreFunc will write the error to the http.ResponseWriter
type loadStoreFunc func(w http.ResponseWriter, r *http.Request) (store orm.Store, failed bool)

func newSingletonGetHandler(schema *schema.Schema, loadStore loadStoreFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store, ok := loadStore(w, r)
		if !ok {
			return
		}
		newObj := schema.NewStateObject()
		err := store.Get(meta.SingletonID, newObj)
		if err != nil {
			notFound(w, "not found")
			return
		}
		writeObject(w, newObj)
	}
}

// newGetHandler creates an http.HandlerFunc that can be used to fetch a state object
func newGetHandler(schema *schema.Schema, definition schema.Definition, loadStore loadStoreFunc) http.HandlerFunc {
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
		// load store
		store, ok := loadStore(w, req)
		if !ok {
			return
		}
		newObj := schema.NewStateObject()
		err = store.Get(meta.NewBytesID(pkBytes), newObj)
		if err != nil {
			notFound(w, err.Error())
			return
		}
		writeObject(w, newObj)
	}
}

// newListHandler creates an http.HandlerFunc that can be used to fetch a list of state objects of the same kind
func newListHandler(schema *schema.Schema, loadStore loadStoreFunc) http.HandlerFunc {
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
		store, ok := loadStore(w, r)
		if !ok {
			return
		}
		opts := orm.ListOptions{
			MatchFieldInterface: nil,
			MatchFieldString:    nil,
			Start:               listOptions.Start,
			End:                 listOptions.End,
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
			opts.MatchFieldString = append(opts.MatchFieldString, orm.ListMatchFieldString{
				Field: sp[0],
				Value: sp[1],
			})
		}

		iter, err := store.List(schema.NewStateObject(), opts)
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
