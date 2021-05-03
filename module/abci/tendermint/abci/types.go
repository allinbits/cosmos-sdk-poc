package abci

import (
	"fmt"

	crypto2 "github.com/fdymylja/tmos/module/abci/tendermint/crypto"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/proto/tendermint/crypto"
	"google.golang.org/protobuf/proto"
)

// FromLegacyProto fills *RequestBeginBlock from the legacy gogoproto.Message
func (x *RequestBeginBlock) FromLegacyProto(block *types.RequestBeginBlock) {
	convert(block, x)
}

func (x *ValidatorUpdate) ToLegacy() types.ValidatorUpdate {
	pk := crypto.PublicKey{}
	switch t := x.PubKey.Sum.(type) {
	case *crypto2.PublicKey_Ed25519:
		pk = crypto.PublicKey{Sum: &crypto.PublicKey_Ed25519{Ed25519: t.Ed25519}}
	case *crypto2.PublicKey_Secp256K1:
		pk = crypto.PublicKey{Sum: &crypto.PublicKey_Secp256K1{Secp256K1: t.Secp256K1}}
	default:
		panic(fmt.Sprintf("unsupported key type: %T", t))
	}
	return types.ValidatorUpdate{
		PubKey: pk,
		Power:  x.Power,
	}
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
