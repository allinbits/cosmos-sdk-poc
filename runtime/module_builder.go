package runtime

import (
	"github.com/fdymylja/tmos/runtime/controller"
	"github.com/fdymylja/tmos/runtime/meta"
)

// ModuleDescriptor describes the full functionality set of a Module
type ModuleDescriptor struct {
	Name                           string
	AdmissionControllers           []admissionController
	MutatingAdmissionControllers   []mutatingAdmissionController
	StateTransitionControllers     []stateTransitionController
	PostStateTransitionControllers []postStateTransitionController
	StateObjects                   []stateObject
	Needs                          []meta.StateTransition
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

func NewModuleBuilder() *ModuleBuilder {
	return &ModuleBuilder{Descriptor: new(ModuleDescriptor)}
}

// ModuleBuilder is used to build a module
type ModuleBuilder struct {
	Descriptor *ModuleDescriptor
}

func (b *ModuleBuilder) Named(name string) *ModuleBuilder {
	b.Descriptor.Name = name
	return b
}

func (b *ModuleBuilder) HandlesStateTransition(transition meta.StateTransition, ctrl controller.StateTransition) *ModuleBuilder {
	b.Descriptor.StateTransitionControllers = append(b.Descriptor.StateTransitionControllers, stateTransitionController{transition, ctrl})
	return b
}

func (b *ModuleBuilder) HandlesAdmission(transition meta.StateTransition, ctrl controller.Admission) *ModuleBuilder {
	b.Descriptor.AdmissionControllers = append(b.Descriptor.AdmissionControllers, admissionController{transition, ctrl})
	return b
}

func (b *ModuleBuilder) OwnsStateObject(object meta.StateObject) *ModuleBuilder {
	b.Descriptor.StateObjects = append(b.Descriptor.StateObjects, stateObject{object})
	return b
}
