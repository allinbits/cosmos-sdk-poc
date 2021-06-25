package rbac

import (
	"testing"

	"github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/core/rbac/xt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestInitRoleCreator(t *testing.T) {
	noXt := (&v1alpha1.MsgBindRole{}).ProtoReflect().Descriptor().Options().(*descriptorpb.MessageOptions)
	withXt := (&v1alpha1.MsgCreateRole{}).ProtoReflect().Descriptor().Options().(*descriptorpb.MessageOptions)

	t.Logf("%#v", noXt)
	t.Logf("%#v", withXt)

	notExist := proto.GetExtension(noXt, xt.E_StateTransitionAcl)
	xtExists := proto.GetExtension(withXt, xt.E_StateTransitionAcl)

	t.Logf("%#v %v", notExist, notExist.(*xt.StateTransitionAccessControl) == nil)
	t.Logf("%#v", xtExists)
}
