package v1

import (
	meta "github.com/fdymylja/tmos/core/meta"
	client "github.com/fdymylja/tmos/runtime/client"
	schema "github.com/fdymylja/tmos/runtime/orm/schema"
)

func (x *MsgCreatePost) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "testmodule.v1",
		Kind:    "MsgCreatePost",
		ApiType: meta.APIType_StateTransition,
	}
}

func (x *MsgCreatePost) NewStateTransition() meta.StateTransition {
	return new(MsgCreatePost)
}

func (x *Post) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "testmodule.v1",
		Kind:    "Post",
		ApiType: meta.APIType_StateObject,
	}
}

func (x *Post) NewStateObject() meta.StateObject {
	return new(Post)
}

type PostClient interface {
	Get(id string, opts ...client.GetOption) (*Post, error)
	List(opts ...client.ListOption) (PostIterator, error)
	Create(post *Post, opts ...client.CreateOption) error
	Delete(post *Post, opts ...client.DeleteOption) error
	Update(post *Post, opts ...client.UpdateOption) error
}

type postClient struct {
	client client.RuntimeClient
}

func (x *postClient) Get(id string, opts ...client.GetOption) (*Post, error) {
	_spfGenO := new(Post)
	_spfGenID := meta.NewStringID(id)
	_spfGenErr := x.client.Get(_spfGenID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *postClient) List(opts ...client.ListOption) (PostIterator, error) {
	iter, err := x.client.List(new(Post), opts...)
	if err != nil {
		return nil, err
	}
	return &postIterator{iter: iter}, nil
}

func (x *postClient) Create(post *Post, opts ...client.CreateOption) error {
	return x.client.Create(post, opts...)
}

func (x *postClient) Delete(post *Post, opts ...client.DeleteOption) error {
	return x.client.Delete(post, opts...)
}

func (x *postClient) Update(post *Post, opts ...client.UpdateOption) error {
	return x.client.Update(post, opts...)
}

type PostIterator interface {
	Get() (*Post, error)
	Valid() bool
	Next()
}

type postIterator struct {
	iter client.ObjectIterator
}

func (x *postIterator) Get() (*Post, error) {
	obj := new(Post)
	err := x.iter.Get(obj)
	return obj, err
}
func (x *postIterator) Valid() bool {
	return x.iter.Valid()
}

func (x *postIterator) Next() {
	x.iter.Next()
}

func (x *Params) APIDefinition() *meta.APIDefinition {
	return &meta.APIDefinition{
		Group:   "testmodule.v1",
		Kind:    "Params",
		ApiType: meta.APIType_StateObject,
	}
}

func (x *Params) NewStateObject() meta.StateObject {
	return new(Params)
}

type ParamsClient interface {
	Get(opts ...client.GetOption) (*Params, error)
	Create(params *Params, opts ...client.CreateOption) error
	Delete(params *Params, opts ...client.DeleteOption) error
	Update(params *Params, opts ...client.UpdateOption) error
}

type paramsClient struct {
	client client.RuntimeClient
}

func (x *paramsClient) Get(opts ...client.GetOption) (*Params, error) {
	_spfGenO := new(Params)
	_spfGenErr := x.client.Get(meta.SingletonID, _spfGenO, opts...)
	if _spfGenErr != nil {
		return nil, _spfGenErr
	}
	return _spfGenO, nil
}

func (x *paramsClient) Create(params *Params, opts ...client.CreateOption) error {
	return x.client.Create(params, opts...)
}

func (x *paramsClient) Delete(params *Params, opts ...client.DeleteOption) error {
	return x.client.Delete(params, opts...)
}

func (x *paramsClient) Update(params *Params, opts ...client.UpdateOption) error {
	return x.client.Update(params, opts...)
}

var PostSchema = &schema.Definition{
	PrimaryKey:    "id",
	SecondaryKeys: []string{"creator"},
}

var ParamsSchema = &schema.Definition{
	Singleton: true,
}

type ClientSet interface {
	Posts() PostClient
	Params() ParamsClient
	ExecMsgCreatePost(msg *MsgCreatePost) error
}

func NewClientSet(client client.RuntimeClient) ClientSet {
	return &clientSet{
		client:       client,
		postClient:   &postClient{client: client},
		paramsClient: &paramsClient{client: client},
	}
}

type clientSet struct {
	client client.RuntimeClient
	// postClient is the client used to interact with Post
	postClient PostClient
	// paramsClient is the client used to interact with Params
	paramsClient ParamsClient
}

func (x *clientSet) Posts() PostClient {
	return x.postClient
}

func (x *clientSet) Params() ParamsClient {
	return x.paramsClient
}

func (x *clientSet) ExecMsgCreatePost(msg *MsgCreatePost) error {
	return x.client.Deliver(msg)
}
