package store

import "github.com/fdymylja/tmos/runtime/meta"

type Store interface {
	Get(meta.ID, meta.StateObject) error
	List(object meta.StateObject) error
	Create(meta.StateObject) error
	Update(meta.StateObject) error
	Delete(meta.ID, meta.StateObject) error
}
