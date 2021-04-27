package module

import (
	"github.com/fdymylja/tmos/apis/meta"
	"github.com/fdymylja/tmos/pkg/controller"
)

// Descriptor describes the full functionality set of a Module
type Descriptor struct {
	Name                           string
	AdmissionControllers           []admissionController
	MutatingAdmissionControllers   []mutatingAdmissionController
	StateTransitionControllers     []stateTransitionController
	PostStateTransitionControllers []postStateTransitionController
	StateObjects                   []stateObject
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
}

type postStateTransitionController struct {
}

type stateObject struct {
	StateObject meta.StateObject
}

func NewBuilder() *Builder {
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

func (b *Builder) HandleStateTransition(transition meta.StateTransition, ctrl controller.StateTransition) *Builder {
	b.Descriptor.StateTransitionControllers = append(b.Descriptor.StateTransitionControllers, stateTransitionController{transition, ctrl})
	return b
}

func (b *Builder) OwnsStateObject(object meta.StateObject) *Builder {
	b.Descriptor.StateObjects = append(b.Descriptor.StateObjects, stateObject{object})
	return b
}
