package module

import (
	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

// Descriptor describes the full functionality set of a Module
type Descriptor struct {
	Name                                 string
	GenesisHandler                       GenesisHandler
	StateTransitionAdmissionHandlers     []stateTransitionAdmissionHandler
	StateTransitionPreExecHandlers       []stateTransitionPreExecutionHandler
	StateTransitionExecutionHandlers     []stateTransitionExecutionHandler
	StateTransitionPostExecutionHandlers []stateTransitionPostExecutionHandler
	StateObjects                         []StateObject
	Needs                                []meta.StateTransition
	AuthAdmissionHandlers                []authAdmissionHandler
	PostAuthenticationHandler            []postAuthenticationHandler
}

type StateObject struct {
	StateObject meta.StateObject
	Options     schema.Definition
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
		Handler:         ctrl,
		External:        external,
	})
	return b
}

func (b *DescriptorBuilder) HandlesAdmission(transition meta.StateTransition, ctrl statetransition.AdmissionHandler) *DescriptorBuilder {
	b.descriptor.StateTransitionAdmissionHandlers = append(b.descriptor.StateTransitionAdmissionHandlers, stateTransitionAdmissionHandler{transition, ctrl})
	return b
}

func (b *DescriptorBuilder) OwnsStateObject(object meta.StateObject, options schema.Definition) *DescriptorBuilder {
	b.descriptor.StateObjects = append(b.descriptor.StateObjects, StateObject{
		StateObject: object,
		Options:     options,
	})
	return b
}

func (b *DescriptorBuilder) WithGenesis(ctrl GenesisHandler) *DescriptorBuilder {
	b.descriptor.GenesisHandler = ctrl

	return b
}

func (b *DescriptorBuilder) NeedsStateTransition(transition meta.StateTransition) *DescriptorBuilder {
	b.descriptor.Needs = append(b.descriptor.Needs, transition)
	return b
}

func (b *DescriptorBuilder) WithAuthAdmissionHandler(ctrl authentication.AdmissionHandler) *DescriptorBuilder {
	b.descriptor.AuthAdmissionHandlers = append(b.descriptor.AuthAdmissionHandlers, authAdmissionHandler{Handler: ctrl})
	return b
}

func (b *DescriptorBuilder) WithPostAuthenticationHandler(ctrl authentication.PostAuthenticationHandler) *DescriptorBuilder {
	b.descriptor.PostAuthenticationHandler = append(b.descriptor.PostAuthenticationHandler, postAuthenticationHandler{Handler: ctrl})
	return b
}

func (b *DescriptorBuilder) Build() Descriptor {
	return b.descriptor
}
