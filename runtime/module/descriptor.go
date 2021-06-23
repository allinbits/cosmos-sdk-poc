package module

import (
	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/core/module"
	"github.com/fdymylja/tmos/pkg/protoutils/desc"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/authorization"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"github.com/fdymylja/tmos/runtime/statetransition"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// RawDescriptor aliases module.Descriptor
type RawDescriptor = module.Descriptor

// Descriptor describes the full functionality set of a Module
type Descriptor struct {
	Name                                 string
	GenesisHandler                       GenesisHandler
	StateTransitionAdmissionHandlers     []stateTransitionAdmissionHandler
	StateTransitionPreExecHandlers       []stateTransitionPreExecutionHandler
	StateTransitionExecutionHandlers     []stateTransitionExecutionHandler
	StateTransitionPostExecutionHandlers []stateTransitionPostExecutionHandler
	StateObjects                         []stateObject
	Needs                                []meta.StateTransition
	AuthAdmissionHandlers                []authentication.AdmissionHandler
	PostAuthenticationHandler            []authentication.PostAuthenticationHandler
	Services                             []ExtensionService
	Authorizer                           authorization.Authorizer
}

func (d Descriptor) Raw() *module.Descriptor {
	x := new(module.Descriptor)

	x.Name = d.Name
	x.StateObjects = make([]*module.StateObject, len(d.StateObjects))
	for i, so := range d.StateObjects {
		x.StateObjects[i] = &module.StateObject{
			ApiDefinition:     so.StateObject.APIDefinition(),
			SchemaDefinition:  so.SchemaDefinition,
			ProtobufFullname:  (string)(so.StateObject.ProtoReflect().Descriptor().FullName()),
			ProtoDependencies: getDeps(so.StateObject.ProtoReflect().Descriptor().ParentFile()),
		}
	}
	x.StateTransitions = make([]*module.StateTransition, len(d.StateTransitionExecutionHandlers))
	for i, st := range d.StateTransitionExecutionHandlers {
		x.StateTransitions[i] = &module.StateTransition{
			ApiDefinition:     st.StateTransition.APIDefinition(),
			ProtobufFullname:  (string)(st.StateTransition.ProtoReflect().Descriptor().FullName()),
			ProtoDependencies: getDeps(st.StateTransition.ProtoReflect().Descriptor().ParentFile()),
		}
	}

	return x
}

func getDeps(file protoreflect.FileDescriptor) []*descriptorpb.FileDescriptorProto {
	deps := desc.Dependencies(file)
	rawDeps := make([]*descriptorpb.FileDescriptorProto, 0, len(deps)+1)
	for _, x := range deps {
		rawDeps = append(rawDeps, protodesc.ToFileDescriptorProto(x))
	}
	rawDeps = append(rawDeps, protodesc.ToFileDescriptorProto(file))
	return rawDeps
}

type stateObject struct {
	StateObject      meta.StateObject
	SchemaDefinition *schema.Definition
}

type stateTransitionAdmissionHandler struct {
	StateTransition  meta.StateTransition
	AdmissionHandler statetransition.AdmissionHandler
}

type stateTransitionPreExecutionHandler struct {
	StateTransition meta.StateTransition
	Handler         statetransition.PreExecutionHandler
}

type stateTransitionExecutionHandler struct {
	StateTransition meta.StateTransition
	Handler         statetransition.ExecutionHandler
	External        bool
}

type stateTransitionPostExecutionHandler struct {
	StateTransition statetransition.StateTransition
	Handler         statetransition.PostExecutionHandler
}
