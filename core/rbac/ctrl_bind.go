package rbac

import (
	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

func NewBindRoleHandler(client module.Client) statetransition.ExecutionHandler {
	return BindRoleHandler{client: client}
}

type BindRoleHandler struct {
	client module.Client
}

func (b BindRoleHandler) Exec(req statetransition.ExecutionRequest) (statetransition.ExecutionResponse, error) {
	msg := req.Transition.(*v1alpha1.MsgBindRole)
	return statetransition.ExecutionResponse{}, b.client.Create(&v1alpha1.RoleBinding{
		Subject: msg.Subject,
		RoleRef: msg.RoleId,
	})
}

func NewBindRoleAdmission(client module.Client) statetransition.AdmissionHandler {
	return BindRoleAdmission{client: client}
}

type BindRoleAdmission struct {
	client module.Client
}

func (b BindRoleAdmission) Validate(request statetransition.AdmissionRequest) error {
	msg := request.Transition.(*v1alpha1.MsgBindRole)
	if err := b.roleExists(msg.RoleId); err != nil {
		return err
	}
	return nil
}

func (b BindRoleAdmission) roleExists(id string) error {
	err := b.client.Get(meta.NewStringID(id), new(v1alpha1.Role))
	if err != nil {
		return err
	}
	return nil
}
