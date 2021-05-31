package module

import (
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/orm/schema"
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
	AdmissionControllers                 []authenticationAdmissionController
	TransitionControllers                []authenticationTransitionController
}

type stateObject struct {
	StateObject meta.StateObject
	Options     schema.Definition
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

func (b *DescriptorBuilder) OwnsStateObject(object meta.StateObject, options schema.Definition) *DescriptorBuilder {
	b.descriptor.StateObjects = append(b.descriptor.StateObjects, stateObject{
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

func (b *DescriptorBuilder) WithAdmissionController(ctrl authentication.AdmissionHandler) *DescriptorBuilder {
	b.descriptor.AdmissionControllers = append(b.descriptor.AdmissionControllers, authenticationAdmissionController{Handler: ctrl})
	return b
}

func (b *DescriptorBuilder) WithTransitionController(ctrl authentication.PostAuthenticationHandler) *DescriptorBuilder {
	b.descriptor.TransitionControllers = append(b.descriptor.TransitionControllers, authenticationTransitionController{Handler: ctrl})
	return b
}

func (b *DescriptorBuilder) Build() Descriptor {
	return b.descriptor
}
