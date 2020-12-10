package application

import "github.com/fdymylja/cosmos-os/pkg/codec"

// Application defines an application running on the runtime
type Application interface {
	// ID is the applications unique ID
	ID() ID
	// RegisterDeliverers registers handlers that do state transitions
	RegisterDeliverers(register RegisterDelivererFn)
	// RegisterCheckers registers functions that do check state
	// transitions messages in a stateless way (ADMISSION POLICY)
	RegisterCheckers(register RegisterCheckerFn)
	// RegisterQueriers registers functions that handle queries
	// coming from outside and within modules
	RegisterQueriers(register RegisterQuerierFn)
	// TODO
	InitGenesis()
	ExportGenesis()
	RegisterInvariant()
}

type RegisterDelivererFn func(message codec.Object, handler DeliverFunc)
type RegisterCheckerFn func(message codec.Object, handler CheckFunc)
type RegisterQuerierFn func(request codec.Object, response codec.Object, handler QueryFunc)

// CheckFunc is the function applications use to check requests in a stateless way
type CheckFunc func(CheckRequest) error

// DeliverFunc is the function applications use to deliver state changing requests
type DeliverFunc func(DeliverRequest) (DeliverResponse, error)

// QueryFunc is the function used by applications to query state
type QueryFunc func(QueryRequest) (QueryResponse, error)

// CheckRequest
type CheckRequest struct {
	Request []byte
}

type DeliverRequest struct {
	Request []byte
	Client  Client
	Store   DB
}

type DeliverResponse struct {
}

type QueryRequest struct {
	Request []byte
	Client  QueryClient
	DB      QueryDB
}

type QueryResponse struct {
	Response []byte
}

// Client is the client used to interact with other applications
type Client interface {
	QueryClient
	Deliver(request codec.Object) (DeliverResponse, error)
}

type QueryClient interface {
	Query(request codec.Object, response codec.Object) error
}

type DB interface {
	QueryDB
	Set(key []byte, object codec.Object) error
}

type QueryDB interface {
	Get(key []byte, object codec.Object) error
}

type ID string
