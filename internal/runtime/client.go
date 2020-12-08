package runtime

import (
	"github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/fdymylja/cosmos-os/pkg/application"
	"github.com/fdymylja/cosmos-os/pkg/codec"
)

func newAppClient(cdc codec.Codec, db *iavl.Store, deliver map[string]deliverer, query map[string]querier) appClient {
	return appClient{
		db:            db,
		queryRouter:   query,
		deliverRouter: deliver,
		cdc:           cdc,
	}
}

type appClient struct {
	db *iavl.Store
	queryRouter map[string]querier
	deliverRouter map[string]deliverer
	cdc codec.Codec
}

func (q appClient) copy() appClient {
	return appClient{
		db:          q.db,
		queryRouter: q.queryRouter,
		cdc:         q.cdc,
	}
}

func (q appClient) Query(request codec.Object, response codec.Object) error {
	name := q.cdc.Name(request)
	querier, exists := q.queryRouter[name]
	if !exists {
		panic("not exists")
	}
	reqBytes, err := q.cdc.Marshal(request)
	if err != nil {
		panic(err)
	}
	resp, err := querier.do(application.QueryRequest{
		Request: reqBytes,
		Client:  q.copy(),
		DB:      appStore{
			db:     q.db,
			prefix: []byte(querier.applicationID),
			cdc:    q.cdc,
		},
	})
	if err != nil {
		panic(err)
	}
	return q.cdc.Unmarshal(resp.Response, response)
}

func (q appClient) Deliver(request codec.Object) (application.DeliverResponse, error) {
	name := q.cdc.Name(request)
	deliverer, exists := q.deliverRouter[name]
	if !exists {
		panic("not exists")
	}
	reqBytes, err := q.cdc.Marshal(request)
	if err != nil {
		panic(err)
	}
	resp, err := deliverer.do(application.DeliverRequest{
		Request: reqBytes,
		Client:  newAppClient(q.cdc, q.db, q.deliverRouter, q.queryRouter),
		Store:   newAppStore(q.cdc, q.db, deliverer.applicationID),
	})
	return resp, err
}
