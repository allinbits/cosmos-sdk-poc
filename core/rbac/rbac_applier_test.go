package rbac

import (
	"testing"

	"github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestInitRoleCreator(t *testing.T) {
	noXt := (&v1alpha1.MsgBindRole{}).ProtoReflect().Descriptor().Options().(*descriptorpb.MessageOptions)
	withXt := (&v1alpha1.MsgCreateRole{}).ProtoReflect().Descriptor().Options().(*descriptorpb.MessageOptions)

	t.Logf("%#v", noXt)
	t.Logf("%#v", withXt)

	notExist := proto.GetExtension(noXt, v1alpha1.E_StateTransitionAcl)
	xtExists := proto.GetExtension(withXt, v1alpha1.E_StateTransitionAcl)

	t.Logf("%#v %v", notExist, notExist.(*v1alpha1.StateTransitionAccessControl) == nil)
	t.Logf("%#v", xtExists)
}
