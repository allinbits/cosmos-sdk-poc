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

func (x *postClient) Create(post *Post, opts ...client.CreateOption) error {
	return x.client.Create(post, opts...)
}

func (x *postClient) Delete(post *Post, opts ...client.DeleteOption) error {
	return x.client.Delete(post, opts...)
}

func (x *postClient) Update(post *Post, opts ...client.UpdateOption) error {
	return x.client.Update(post, opts...)
}

var PostSchema = schema.Definition{
	PrimaryKey:    "id",
	SecondaryKeys: []string{"creator"},
}

type ClientSet interface {
	Posts() PostClient
	ExecMsgCreatePost(msg *MsgCreatePost) error
}

func NewClientSet(client client.RuntimeClient) ClientSet {
	return &clientSet{
		client:     client,
		postClient: &postClient{client: client},
	}
}

type clientSet struct {
	client client.RuntimeClient
	// postClient is the client used to interact with Post
	postClient PostClient
}

func (x *clientSet) Posts() PostClient {
	return x.postClient
}

func (x *clientSet) ExecMsgCreatePost(msg *MsgCreatePost) error {
	return x.client.Deliver(msg)
}
