package store

import (
	"github.com/fdymylja/tmos/runtime"
)

type Store interface {
	Get(runtime.ID) error
	Create(runtime.ID, runtime.StateObject) error
	Update(runtime.ID, runtime.StateObject) error
	List(object runtime.StateObject) error
}
