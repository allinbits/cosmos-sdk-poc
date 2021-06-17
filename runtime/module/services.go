package module

import "github.com/fdymylja/tmos/runtime/client"

// ExtensionService defines a module that extends runtime with a secondary service
type ExtensionService interface {
	// SetClient is called to set the client used to interact
	// with runtime.
	SetClient(client client.ReadOnlyClient)
	// Start is called when the service starts
	Start() error
	// Stop is called when the runtime is stopping
	Stop() error
}

var _ ExtensionService = (*BaseService)(nil)

// BaseService defines a runtime ExtensionService
type BaseService struct {
	client.ReadOnlyClient
}

func (b *BaseService) SetClient(client client.ReadOnlyClient) {
	b.ReadOnlyClient = client
}

func (b *BaseService) Start() error {
	return nil
}

func (b *BaseService) Stop() error {
	return nil
}
