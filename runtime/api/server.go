package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/encoding/protojson"
	"k8s.io/klog/v2"
)

func NewServer(store orm.Store) *Server {
	return &Server{
		server:  nil,
		store:   store,
		mux:     mux.NewRouter(),
		modules: map[string]module.Descriptor{},
	}
}

type Server struct {
	server     *http.Server
	store      orm.Store
	mux        *mux.Router
	modules    map[string]module.Descriptor
	knownPaths map[string]string // maps known paths of objects
}

// RegisterModuleAPI registers the API for the server
func (s *Server) RegisterModuleAPI(module module.Descriptor) error {
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

func (s *Server) Start() {
	klog.Infof("starting api server...")
	go func() {
		s.server = &http.Server{
			Addr:    ":8080", // TODO configurable
			Handler: s.mux,
		}
		panic(s.server.ListenAndServe()) // TODO me better
	}()
}

func (s *Server) loadStore(writer http.ResponseWriter, request *http.Request) (orm.Store, bool) {
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

func (s *Server) registerStateObjectHandlers(descriptor module.Descriptor, obj module.StateObject) error {
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
	case false:
		// create get
		def := obj.StateObject.APIDefinition()
		s.mux.
			Methods(http.MethodGet).
			Path(
				strings.ToLower(
					fmt.Sprintf("/%s/%s/{%s}", def.Group, def.Kind, obj.Options.PrimaryKey),
				)).
			HandlerFunc(newGetHandler(sch, obj.Options, s.loadStore))
		// create list
		s.mux.Methods(http.MethodGet).
			Path(
				strings.ToLower(
					fmt.Sprintf("/%s/%ss", def.Group, def.Kind), // TODO the plural name should be in the state object schema options
				),
			).
			HandlerFunc(newListHandler(sch, obj.Options, s.loadStore))
	}

	return nil
}

func writeObject(w io.Writer, obj meta.StateObject) bool {
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

// loadStoreFunc returns the versioned store given an http request and returns if the loading succeeded or not
// if the loading fails loadStoreFunc will write the error to the http.ResponseWriter
type loadStoreFunc func(w http.ResponseWriter, r *http.Request) (store orm.Store, failed bool)

// newListHandler creates an http.HandlerFunc that can be used to fetch a list of state objects of the same kind
func newListHandler(schema *schema.Schema, definition schema.Definition, loadStore loadStoreFunc) http.HandlerFunc {
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
		var opts []orm.ListOption
		opts = append(opts, orm.ListRange{
			Start: listOptions.Start,
			End:   listOptions.End,
		})

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
			opts = append(opts, orm.ListMatchFieldString{
				Field: sp[0],
				Value: sp[1],
			})
		}

		iter, err := store.List(schema.NewStateObject(), opts...)
		if err != nil {
			badRequest(w, "unable to list any object: %s", err)
			return
		}
		defer iter.Close()

		var objects []meta.StateObject

		for iter.Valid() {
			obj := schema.NewStateObject()
			err := iter.Get(obj)
			if err != nil {
				badRequest(w, "unable to list object: %s", err)
				return
			}
			objects = append(objects, obj)
			iter.Next()
		}

		writeObjectList(w, objects)
	}
}

func writeObjectList(w io.Writer, objects []meta.StateObject) {
	_, _ = w.Write([]byte("{\"items\":["))
	max := len(objects)
	for i, o := range objects {
		writeObject(w, o)
		if i != max-1 {
			_, _ = w.Write([]byte(","))
		}
	}
	_, _ = w.Write([]byte("]}"))
}

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
