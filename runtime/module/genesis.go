package module

import "encoding/json"

// GenesisHandler defines the module Genesis handler
type GenesisHandler interface {
	Default(c Client) error
	Import(state json.RawMessage) error
	Export() (state json.RawMessage, err error)
}
