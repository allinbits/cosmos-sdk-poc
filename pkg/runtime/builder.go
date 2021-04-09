package runtime

import "github.com/fdymylja/tmos/pkg/application"

func NewBuilder() *Builder {
	return &Builder{
		runtime: newRuntime(),
	}
}

// Builder takes care of builds the Runtime
type Builder struct {
	runtime *Runtime
}

func (b *Builder) Mount(app application.Application) error {
	return b.runtime.mount(app)
}

func (b *Builder) Build() (*Runtime, error) {
	return b.runtime, nil
}
