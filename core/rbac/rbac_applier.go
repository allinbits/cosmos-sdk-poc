package rbac

import (
	"fmt"

	abciv1alpha1 "github.com/fdymylja/tmos/core/abci/v1alpha1"
	"github.com/fdymylja/tmos/core/module"
	"github.com/fdymylja/tmos/core/rbac/v1alpha1"
	runtimev1alpha1 "github.com/fdymylja/tmos/core/runtime/v1alpha1"
	"github.com/fdymylja/tmos/runtime/client"
	"github.com/fdymylja/tmos/runtime/statetransition"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

func NewInitRoleCreator(client client.Client) statetransition.PostExecutionHandler {
	return initGenesisRoleCreator{
		client: v1alpha1.NewClientSet(client),
		abci:   abciv1alpha1.NewClientSet(client),
	}
}

type initGenesisRoleCreator struct {
	client v1alpha1.ClientSet
	abci   abciv1alpha1.ClientSet
}

func (i initGenesisRoleCreator) PostExec(req statetransition.PostExecutionRequest) error {
	// this operation runs only at init genesis level
	// if this is not init genesis we skip.
	stage, err := i.abci.Stage().Get()
	if err != nil {
		return err
	}

	if stage.Stage != abciv1alpha1.ABCIStage_InitChain {
		return nil
	}

	msg := req.Transition.(*runtimev1alpha1.CreateModuleDescriptors)

	for _, m := range msg.Modules {
		for _, st := range m.StateTransitions {
			acl, err := i.isExternal(st)
			if err != nil {
				return err
			}
			panic(acl)
		}
	}

	return nil
}

func (i initGenesisRoleCreator) handleModule(m *module.Descriptor) error {
	moduleRole := new(v1alpha1.Role)
	userRole := new(v1alpha1.Role)

	// handle state transitions
	for _, st := range m.StateTransitions {
		acl, err := i.isExternal(st)
		if err != nil {
			return err
		}
		err = moduleRole.ExtendRaw(runtimev1alpha1.Verb_Deliver, st.ApiDefinition)
		if err != nil {
			return err
		}

		if !acl.External {
			continue
		}

		err = userRole.ExtendRaw(runtimev1alpha1.Verb_Deliver, st.ApiDefinition)
		if err != nil {
			return err
		}
	}
	// handle state objects role
	return nil
}

func (i initGenesisRoleCreator) isExternal(std *module.StateTransition) (*v1alpha1.StateTransitionAccessControl, error) {
	// we search for the file descriptor which contains the given state transition
	fileRegistry := new(protoregistry.Files)

	// runtime provides us the files in order so every file
	// we add we are sure it already has the required dependencies.
	var fds []protoreflect.FileDescriptor // NOTE(fdymylja): we cannot use protoregistry.Files.Range because the order is non deterministic
	for _, rawFD := range std.ProtoDependencies {
		fd, err := protodesc.NewFile(rawFD, fileRegistry)
		if err != nil {
			return nil, err
		}
		fds = append(fds, fd)
	}

	var md protoreflect.MessageDescriptor
	for _, fd := range fds {
		got := fd.Messages().ByName(protoreflect.FullName(std.ProtobufFullname).Name())
		if got != nil {
			md = got
			break
		}
	}

	// critical error, there's some data corruption uh
	if md == nil {
		panic(fmt.Errorf("protobuf file %s was not found in the provided runtime registry", std.ProtobufFullname))
	}

	opt := md.Options().(*descriptorpb.MessageOptions)
	xt := proto.GetExtension(opt, v1alpha1.E_StateTransitionAcl).(*v1alpha1.StateTransitionAccessControl)
	if xt == nil {
		return nil, fmt.Errorf("state transition %s has no RBAC options set", std.ProtobufFullname)
	}

	return xt, nil
}
