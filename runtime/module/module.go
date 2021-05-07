package module

import "github.com/fdymylja/tmos/runtime/client"

type Client interface {
	client.RuntimeClient
}

// Module defines a basic module which handles changes
type Module interface {
	Initialize(client Client, builder *Builder)
}
