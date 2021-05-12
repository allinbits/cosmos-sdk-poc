package module

import "github.com/fdymylja/tmos/runtime/authentication"

type AuthenticationExtension interface {
	Initialize(builder *AuthenticationExtensionBuilder)
}

type authenticationAdmissionController struct {
	Handler authentication.AdmissionHandler
}

type authenticationTransitionController struct {
	Handler authentication.PostAuthenticationHandler
}

// AuthenticationExtensionDescriptor describes an AuthenticationExtension
type AuthenticationExtensionDescriptor struct {
	AdmissionControllers  []authenticationAdmissionController
	TransitionControllers []authenticationTransitionController
}

func NewAuthenticationExtensionBuilder() *AuthenticationExtensionBuilder {
	return &AuthenticationExtensionBuilder{descriptor: new(AuthenticationExtensionDescriptor)}
}

// AuthenticationExtensionBuilder is a structure that can be used to extend authentication
type AuthenticationExtensionBuilder struct {
	descriptor *AuthenticationExtensionDescriptor
}

func (a *AuthenticationExtensionBuilder) WithAdmissionController(ctrl authentication.AdmissionHandler) *AuthenticationExtensionBuilder {
	a.descriptor.AdmissionControllers = append(a.descriptor.AdmissionControllers, authenticationAdmissionController{Handler: ctrl})
	return a
}

func (a *AuthenticationExtensionBuilder) WithTransitionController(ctrl authentication.PostAuthenticationHandler) *AuthenticationExtensionBuilder {
	a.descriptor.TransitionControllers = append(a.descriptor.TransitionControllers, authenticationTransitionController{Handler: ctrl})
	return a
}
