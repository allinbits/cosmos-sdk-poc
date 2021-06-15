package testmodule

import (
	"encoding/json"
	"fmt"

	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
	v1 "github.com/fdymylja/tmos/testdata/testmodule/v1"
)

type Module struct{}

func (m Module) Initialize(client module.Client) module.Descriptor {
	return module.NewDescriptorBuilder().
		Named("test").
		OwnsStateObject(&v1.Post{}, v1.PostSchema).
		OwnsStateObject(&v1.Params{}, v1.ParamsSchema).
		HandlesStateTransition(&v1.MsgCreatePost{}, newMsgCreatePost(client), true).
		HandlesAdmission(&v1.MsgCreatePost{}, newMsgCreatePostAdmission(client)).
		WithGenesis(genesis{client: v1.NewClientSet(client)}).
		Build()
}

func newMsgCreatePost(c module.Client) statetransition.ExecutionHandlerFunc {
	client := v1.NewClientSet(c)
	return func(req statetransition.ExecutionRequest) (resp statetransition.ExecutionResponse, err error) {
		msg := req.Transition.(*v1.MsgCreatePost)
		params, err := client.Params().Get()
		if err != nil {
			return
		}
		err = client.Posts().Create(&v1.Post{
			Id:      fmt.Sprintf("%d", params.LastPostNumber),
			Creator: msg.Creator,
			Title:   msg.Title,
			Text:    msg.Text,
		})

		if err != nil {
			return statetransition.ExecutionResponse{}, err
		}
		return resp, nil
	}
}

func newMsgCreatePostAdmission(_ module.Client) statetransition.AdmissionHandlerFunc {
	return func(req statetransition.AdmissionRequest) error {
		msg := req.Transition.(*v1.MsgCreatePost)
		if !req.Users.Has(msg.Creator) {
			return fmt.Errorf("unauthorized")
		}
		if msg.Title == "" {
			return fmt.Errorf("empty title")
		}
		if msg.Text == "" {
			return fmt.Errorf("empty text")
		}
		return nil
	}
}

type genesis struct {
	client v1.ClientSet
}

func (g genesis) Default() error {
	return g.client.Params().Create(&v1.Params{LastPostNumber: 1})
}

func (g genesis) Import(state json.RawMessage) error {
	panic("implement me")
}

func (g genesis) Export() (state json.RawMessage, err error) {
	panic("implement me")
}
