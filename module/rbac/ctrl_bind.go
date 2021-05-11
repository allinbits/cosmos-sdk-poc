package rbac

import (
	"github.com/fdymylja/tmos/module/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime/admission"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

func NewBindRoleController(client module.Client) statetransition.Handler {
	return BindRoleController{client: client}
}

type BindRoleController struct {
	client module.Client
}

func (b BindRoleController) Deliver(req statetransition.Request) (statetransition.Response, error) {
	msg := req.Transition.(*v1alpha1.MsgBindRole)
	return statetransition.Response{}, b.client.Create(&v1alpha1.RoleBinding{
		Subject: msg.Subject,
		RoleRef: msg.RoleId,
	})
}

func NewBindRoleAdmission(client module.Client) admission.Handler {
	return BindRoleAdmission{client: client}
}

type BindRoleAdmission struct {
	client module.Client
}

func (b BindRoleAdmission) Validate(request admission.Request) error {
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
