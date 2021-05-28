package runtime

import (
	"fmt"
	"strings"

	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	ResourcesPath        = "api_resources"
	GetStateObjectsPath  = "get"
	ListStateObjectsPath = "list"
)

type Querier struct {
	reg   *schema.Registry
	store orm.Store
}

func NewQuerier(store orm.Store) *Querier {
	return &Querier{
		reg:   store.SchemaRegistry(),
		store: store,
	}
}

func (q *Querier) Handle(query types.RequestQuery) types.ResponseQuery {
	splitPath := strings.Split(query.Path, "/")
	var response meta.Type
	var err error
	switch splitPath[0] {
	case ResourcesPath:
		// we need to list all APIs in the given group
		response, err = q.apiResources(query)
	case GetStateObjectsPath:
		response, err = q.get(query)
	case ListStateObjectsPath:
	default:
		return types.ResponseQuery{Code: 0x1, Log: "unknown verb"}
	}
	if err != nil {
		return types.ResponseQuery{Code: 0x1, Log: err.Error()}
	}
	b, err := protojson.Marshal(response)
	if err != nil {
		panic(err)
	}
	return types.ResponseQuery{
		Value: b,
	}
}

func (q *Querier) apiResources(_ types.RequestQuery) (meta.Type, error) {
	// TODO we shouldn't do it in this way... we should enable runtime to hold this info
	// so it can be dynamically queried based on height.
	groups := q.reg.ListAPIGroups()
	m := &runtimev1alpha1.Resources{
		ApiGroupResources: make([]*runtimev1alpha1.APIGroupResources, len(groups)),
	}

	for i, g := range groups {
		stateObjects, err := q.reg.ListKindsInGroup(g)
		if err != nil {
			panic(err)
		}

		stateObjectsStr := make([]string, len(stateObjects))
		for j, so := range stateObjects {
			stateObjectsStr[j] = so.String()
		}
		res := &runtimev1alpha1.APIGroupResources{
			ApiGroup:         g.String(),
			StateObjects:     stateObjectsStr,
			StateTransitions: nil,
		}

		m.ApiGroupResources[i] = res
	}

	return m, nil
}

func (q *Querier) get(req types.RequestQuery) (meta.Type, error) {
	split := strings.Split(req.Path, "/")
	if len(split) < 4 {
		return nil, fmt.Errorf("invalid path for get")
	}
	// TODO we should check this in state
	sch, err := q.reg.GetByMeta(meta.Meta{
		APIGroup: meta.APIGroup(split[1]),
		APIKind:  meta.APIKind(split[2]),
	})
	if err != nil {
		return nil, err
	}
	obj := sch.NewStateObject()
	err = q.store.Get(meta.NewStringID(split[3]), obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
