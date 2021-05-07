package module

import (
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/meta"
)

// Descriptor describes the full functionality set of a Module
type Descriptor struct {
	Name                           string
	GenesisHandler                 controller.Genesis
	AdmissionControllers           []admissionController
	MutatingAdmissionControllers   []mutatingAdmissionController
	StateTransitionControllers     []stateTransitionController
	PostStateTransitionControllers []postStateTransitionController
	StateObjects                   []stateObject
	Needs                          []meta.StateTransition
	AuthenticationExtension        *AuthenticationExtensionDescriptor
}

type admissionController struct {
	StateTransition meta.StateTransition
	Controller      controller.Admission
}

type mutatingAdmissionController struct {
}

type stateTransitionController struct {
	StateTransition meta.StateTransition
	Controller      controller.StateTransition
	Internal        bool
}

type postStateTransitionController struct {
}

type stateObject struct {
	StateObject meta.StateObject
}

func NewModuleBuilder() *Builder {
	return &Builder{Descriptor: new(Descriptor)}
}

// Builder is used to build a module
type Builder struct {
	Descriptor *Descriptor
}

func (b *Builder) Named(name string) *Builder {
	b.Descriptor.Name = name
	return b
}

func (b *Builder) HandlesStateTransition(transition meta.StateTransition, ctrl controller.StateTransition) *Builder {
	b.Descriptor.StateTransitionControllers = append(b.Descriptor.StateTransitionControllers, stateTransitionController{
		StateTransition: transition,
		Controller:      ctrl,
		Internal:        false,
	})
	return b
}

func (b *Builder) HandlesAdmission(transition meta.StateTransition, ctrl controller.Admission) *Builder {
	b.Descriptor.AdmissionControllers = append(b.Descriptor.AdmissionControllers, admissionController{transition, ctrl})
	return b
}

func (b *Builder) OwnsStateObject(object meta.StateObject) *Builder {
	b.Descriptor.StateObjects = append(b.Descriptor.StateObjects, stateObject{object})
	return b
}

func (b *Builder) WithGenesis(ctrl controller.Genesis) *Builder {
	b.Descriptor.GenesisHandler = ctrl

	return b
}

func (b *Builder) ExtendsAuthentication(xt AuthenticationExtension) *Builder {
	authXtB := NewAuthenticationExtensionBuilder()
	xt.Initialize(authXtB)
	b.Descriptor.AuthenticationExtension = authXtB.descriptor
	return b
}

func (b *Builder) NeedsStateTransition(transition meta.StateTransition) *Builder {
	b.Descriptor.Needs = append(b.Descriptor.Needs, transition)
	return b
}
