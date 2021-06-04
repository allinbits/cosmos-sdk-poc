package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewServer(store orm.Store) *Server {
	return &Server{
		server:  nil,
		store:   store,
		mux:     nil,
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
	apiDef := obj.StateObject.APIDefinition()
	path := strings.ToLower(fmt.Sprintf("/%s/%s", apiDef.Group, apiDef.Kind))
	// we check if it's singleton or not, if it's singleton
	// the path remains like it currently is
	// ex: /bank.v1/params
	// because only one instance of this object can exist
	// if it's not a singleton, then we add a parameter which represents the primary key
	// ex: /bank.v1/balance/{address}
	if !obj.Options.Singleton {
		path = fmt.Sprintf("%s/{%s}", path, obj.Options.PrimaryKey)
	}
	_, exists := s.knownPaths[path]
	if exists {
		return fmt.Errorf("path already registered: %s", path)
	}
	// create handler based to, if singleton or not
	switch obj.Options.Singleton {
	case true:
		s.mux.
			Methods(http.MethodGet).
			Path(path).
			HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				store, ok := s.loadStore(writer, request)
				if !ok {
					return
				}
				newObj := obj.StateObject.NewStateObject()
				err := store.Get(meta.SingletonID, newObj)
				if err != nil {
					notFound(writer, "not found")
					return
				}
				writeResponse(writer, newObj)
			})
	case false:
		// create get
		s.mux.
			Methods(http.MethodGet).
			Path(path).
			HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			})
		// create list
		s.mux.Methods(http.MethodGet).
			Path(path + "s"). // TODO the plural name should be in the state object schema options
			HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			})
	}

	return nil
}

func writeResponse(writer http.ResponseWriter, obj meta.StateObject) {
	b, err := protojson.Marshal(obj)
	if err != nil {
		panic(err)
	}
	_, _ = writer.Write(b)
}

func notFound(writer http.ResponseWriter, format string, values ...interface{}) {

}

func badRequest(writer http.ResponseWriter, format string, values ...interface{}) {

}
