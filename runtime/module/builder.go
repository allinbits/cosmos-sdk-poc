package module

import (
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

// Descriptor describes the full functionality set of a Module
type Descriptor struct {
	Name                                 string
	GenesisHandler                       GenesisHandler
	StateTransitionAdmissionHandlers     []stateTransitionAdmissionHandlers
	StateTransitionPreExecHandlers       []stateTransitionPreExecutionHandler
	StateTransitionExecutionHandlers     []stateTransitionExecutionHandler
	StateTransitionPostExecutionHandlers []stateTransitionPostExecutionHandler
	StateObjects                         []stateObject
	Needs                                []meta.StateTransition
	AuthenticationExtension              *AuthenticationExtensionDescriptor
}

type stateTransitionAdmissionHandlers struct {
	StateTransition meta.StateTransition
	Controller      statetransition.AdmissionHandler
}

type stateTransitionPreExecutionHandler struct {
	StateTransition meta.StateTransition
	Handler         statetransition.PreExecutionHandler
}

type stateTransitionExecutionHandler struct {
	StateTransition meta.StateTransition
	Controller      statetransition.ExecutionHandler
	External        bool
}

type stateTransitionPostExecutionHandler struct {
	StateTransition statetransition.StateTransition
	Handler         statetransition.PostExecutionHandler
}

type stateObject struct {
	StateObject meta.StateObject
}

func NewDescriptorBuilder() *DescriptorBuilder {
	return &DescriptorBuilder{descriptor: Descriptor{}}
}

// DescriptorBuilder is used to build a module descriptor
type DescriptorBuilder struct {
	descriptor Descriptor
}

func (b *DescriptorBuilder) Named(name string) *DescriptorBuilder {
	b.descriptor.Name = name
	return b
}

func (b *DescriptorBuilder) HandlesStateTransition(transition meta.StateTransition, ctrl statetransition.ExecutionHandler, external bool) *DescriptorBuilder {
	b.descriptor.StateTransitionExecutionHandlers = append(b.descriptor.StateTransitionExecutionHandlers, stateTransitionExecutionHandler{
		StateTransition: transition,
		Controller:      ctrl,
		External:        external,
	})
	return b
}

func (b *DescriptorBuilder) HandlesAdmission(transition meta.StateTransition, ctrl statetransition.AdmissionHandler) *DescriptorBuilder {
	b.descriptor.StateTransitionAdmissionHandlers = append(b.descriptor.StateTransitionAdmissionHandlers, stateTransitionAdmissionHandlers{transition, ctrl})
	return b
}

func (b *DescriptorBuilder) OwnsStateObject(object meta.StateObject) *DescriptorBuilder {
	b.descriptor.StateObjects = append(b.descriptor.StateObjects, stateObject{object})
	return b
}

func (b *DescriptorBuilder) WithGenesis(ctrl GenesisHandler) *DescriptorBuilder {
	b.descriptor.GenesisHandler = ctrl

	return b
}

func (b *DescriptorBuilder) ExtendsAuthentication(xt AuthenticationExtension) *DescriptorBuilder {
	authXtB := NewAuthenticationExtensionBuilder()
	xt.Initialize(authXtB)
	b.descriptor.AuthenticationExtension = authXtB.descriptor
	return b
}

func (b *DescriptorBuilder) NeedsStateTransition(transition meta.StateTransition) *DescriptorBuilder {
	b.descriptor.Needs = append(b.descriptor.Needs, transition)
	return b
}

func (b *DescriptorBuilder) Build() Descriptor {
	return b.descriptor
}
