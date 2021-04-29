package abci

import (
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/proto"
)

// FromLegacyProto fills *RequestBeginBlock from the legacy gogoproto.Message
func (x *RequestBeginBlock) FromLegacyProto(block *types.RequestBeginBlock) {
	convert(block, x)
}

func convert(legacyMessage gogoproto.Message, targetMessage proto.Message) {
	b, err := gogoproto.Marshal(legacyMessage)
	if err != nil {
		panic(err)
	}
	if err := proto.Unmarshal(b, targetMessage); err != nil {
		panic(err)
	}
}
