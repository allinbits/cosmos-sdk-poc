package module

import "github.com/fdymylja/tmos/runtime/authentication"

type authAdmissionHandler struct {
	Handler authentication.AdmissionHandler
}

type postAuthenticationHandler struct {
	Handler authentication.PostAuthenticationHandler
}
