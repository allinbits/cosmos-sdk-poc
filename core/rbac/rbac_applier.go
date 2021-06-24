package rbac

import (
	"fmt"

	abciv1alpha1 "github.com/fdymylja/tmos/core/abci/v1alpha1"
	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/core/module"
	"github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/core/rbac/xt"
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
	/*
		NOTE(fdymylja): this check is currently disabled for different reasons:
		- requires ABCI to initialize before runtime (we're getting ABCI stage which is unset yet)
		- in theory since the message we're executing after is CreateModuleDescriptors, it will be executed once.

		// this operation runs only at init genesis level
		// if this is not init genesis we skip.
		stage, err := i.abci.Stage().Get()
		if err != nil {
			return err
		}

		if stage.Stage != abciv1alpha1.ABCIStage_InitChain {
			return nil
		}
	*/
	msg := req.Transition.(*runtimev1alpha1.CreateModuleDescriptors)

	userRole := v1alpha1.NewExternalAccountRole()
	for _, m := range msg.Modules {
		externalTransitions, err := i.handleModule(m)
		if err != nil {
			return fmt.Errorf("failed processing module %s: %w", m.Name, err)
		}

		for _, st := range externalTransitions {
			err = userRole.ExtendRaw(runtimev1alpha1.Verb_Deliver, st)
			if err != nil {
				return err
			}
		}
	}

	// create external account role
	err := i.client.ExecMsgCreateRole(&v1alpha1.MsgCreateRole{NewRole: userRole})
	if err != nil {
		return err
	}

	return nil
}

func (i initGenesisRoleCreator) handleModule(m *module.Descriptor) ([]*meta.APIDefinition, error) {
	moduleRole := v1alpha1.NewRoleNameForModule(m.Name)

	externalTransitions := make([]*meta.APIDefinition, 0)
	// handle state transitions
	for _, st := range m.StateTransitions {
		acl, err := i.getACL(st)
		if err != nil {
			return nil, err
		}
		err = moduleRole.ExtendRaw(runtimev1alpha1.Verb_Deliver, st.ApiDefinition)
		if err != nil {
			return nil, err
		}

		if !acl.External {
			continue
		}

		externalTransitions = append(externalTransitions, st.ApiDefinition)
	}
	// handle state objects role
	for _, so := range m.StateObjects {
		err := extendRoleForStateObject(moduleRole, so.ApiDefinition)
		if err != nil {
			return nil, err
		}
	}
	// now we go and create the role for the module
	err := i.client.ExecMsgCreateRole(&v1alpha1.MsgCreateRole{NewRole: moduleRole})
	if err != nil {
		return nil, err
	}
	err = i.client.ExecMsgBindRole(&v1alpha1.MsgBindRole{
		RoleId:  moduleRole.Id,
		Subject: m.Name,
	})

	return externalTransitions, nil
}

func (i initGenesisRoleCreator) getACL(std *module.StateTransition) (*xt.StateTransitionAccessControl, error) {
	// we search for the file descriptor which contains the given state transition
	fileRegistry := new(protoregistry.Files)

	// runtime provides us the files in order so every file
	// we add we are sure it already has the required dependencies.
	var fds []protoreflect.FileDescriptor // NOTE(fdymylja): we cannot use protoregistry.Files.Range because the order is non deterministic
	for _, rawFD := range std.ProtoDependencies {
		fd, err := protodesc.NewFile(rawFD, fileRegistry)
		if err != nil {
			return nil, fmt.Errorf("unable to create file descriptor %s for state transition %s: %w", *rawFD.Name, std.ApiDefinition.Name(), err)
		}
		fds = append(fds, fd)
		err = fileRegistry.RegisterFile(fd)
		if err != nil {
			return nil, fmt.Errorf("unable to register file %s for state transition %s: %w", *rawFD.Name, std.ApiDefinition.Name(), err)
		}
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
	xt := proto.GetExtension(opt, xt.E_StateTransitionAcl).(*xt.StateTransitionAccessControl)
	if xt == nil {
		return nil, fmt.Errorf("state transition %s has no RBAC options set", std.ProtobufFullname)
	}

	return xt, nil
}

func extendRoleForStateObject(role *v1alpha1.Role, def *meta.APIDefinition) (err error) {
	err = role.ExtendRaw(runtimev1alpha1.Verb_Create, def)
	if err != nil {
		return err
	}
	err = role.ExtendRaw(runtimev1alpha1.Verb_Delete, def)
	if err != nil {
		return err
	}
	err = role.ExtendRaw(runtimev1alpha1.Verb_Update, def)
	if err != nil {
		return err
	}
	err = role.ExtendRaw(runtimev1alpha1.Verb_Get, def)
	if err != nil {
		return err
	}
	err = role.ExtendRaw(runtimev1alpha1.Verb_List, def)
	if err != nil {
		return err
	}
	return nil
}
