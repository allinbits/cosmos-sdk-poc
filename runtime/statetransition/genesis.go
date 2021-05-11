package statetransition

import "encoding/json"

// Genesis defines the genesis controller
type Genesis interface {
	SetDefault() error
	Import(state json.RawMessage) error
	Export(state json.RawMessage) error
}
