package module

import (
	meta "github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"github.com/fdymylja/tmos/runtime/statetransition"
)

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

func (b *DescriptorBuilder) OwnsStateObject(object meta.StateObject, options *schema.Definition) *DescriptorBuilder {
	b.descriptor.StateObjects = append(b.descriptor.StateObjects, stateObject{
		StateObject:      object,
		SchemaDefinition: options,
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
	b.descriptor.AuthAdmissionHandlers = append(b.descriptor.AuthAdmissionHandlers, ctrl)
	return b
}

func (b *DescriptorBuilder) WithPostAuthenticationHandler(ctrl authentication.PostAuthenticationHandler) *DescriptorBuilder {
	b.descriptor.PostAuthenticationHandler = append(b.descriptor.PostAuthenticationHandler, ctrl)
	return b
}

func (b *DescriptorBuilder) WithPostStateTransitionHandler(st statetransition.StateTransition, ctrl statetransition.PostExecutionHandler) *DescriptorBuilder {
	b.descriptor.StateTransitionPostExecutionHandlers = append(b.descriptor.StateTransitionPostExecutionHandlers, stateTransitionPostExecutionHandler{
		StateTransition: st,
		Handler:         ctrl,
	})

	return b
}

func (b *DescriptorBuilder) WithExtensionService(xt ExtensionService) *DescriptorBuilder {
	b.descriptor.Services = append(b.descriptor.Services, xt)
	return b
}

func (b *DescriptorBuilder) Build() Descriptor {
	return b.descriptor
}
