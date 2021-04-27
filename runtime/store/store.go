package store

import (
	"github.com/fdymylja/tmos/module/meta"
	"github.com/fdymylja/tmos/module/meta/v1alpha1"
)

type ID interface {
	Bytes()
}

type Client interface {
	Get(ID) error
	Create(ID, StateObject) error
	Update(ID, StateObject) error
	List(object StateObject) error
}

// StateObject defines an object which is saved in the state
type StateObject interface {
	meta.Type
	GetObjectMeta() *v1alpha1.ObjectMeta
}

