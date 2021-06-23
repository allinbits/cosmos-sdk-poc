package module

import "github.com/fdymylja/tmos/runtime/client"

// ExtensionService defines a module that extends runtime with a secondary service
type ExtensionService interface {
	// Name identifies the service with a name
	Name() string
	// SetClient is called to set the client used to interact
	// with runtime.
	SetClient(client client.Client)
	// Start is called when the service starts
	Start() error
	// Stop is called when the runtime is stopping
	Stop() error
}

var _ ExtensionService = (*BaseService)(nil)

// BaseService defines a runtime ExtensionService
type BaseService struct {
	client.Client
}

func (b *BaseService) SetClient(client client.Client) {
	b.Client = client
}

func (b *BaseService) Start() error {
	return nil
}

func (b *BaseService) Stop() error {
	return nil
}

func (b *BaseService) Name() string {
	panic("services must have a unique name that identifies them")
}
