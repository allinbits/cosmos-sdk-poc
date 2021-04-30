package crypto

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdkcrypto "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	gogoproto "github.com/gogo/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type PubKey interface {
	sdkcrypto.PubKey
}

func NewDefaultPubKeyResolver() *PubKeyResolver {
	s256k1PbName := gogoproto.MessageName(&secp256k1.PubKey{})
	return &PubKeyResolver{
		knownPubKeys: map[string]struct{}{
			"/" + s256k1PbName: {},
		},
		newPubKey: map[string]func(b []byte) (PubKey, error){
			"/" + s256k1PbName: func(b []byte) (PubKey, error) {
				k := new(secp256k1.PubKey)
				err := gogoproto.Unmarshal(b, k)
				if err != nil {
					return nil, err
				}
				// check size
				if len(k.Key) != secp256k1.PubKeySize {
					return nil, fmt.Errorf("invalid secp2561k key size: %d", len(k.Key))
				}
				return k, nil
			},
		},
	}
}

// PubKeyResolver resolves types and provides pubkeys
type PubKeyResolver struct {
	knownPubKeys map[string]struct{}
	newPubKey    map[string]func(b []byte) (PubKey, error)
}

func (r *PubKeyResolver) New(any *anypb.Any) (PubKey, error) {
	if _, exists := r.knownPubKeys[any.TypeUrl]; !exists {
		return nil, fmt.Errorf("unknown pub key: %s", any.TypeUrl)
	}
	pk, err := r.newPubKey[any.TypeUrl](any.Value)
	return pk, err
}

func (r *PubKeyResolver) Address(bech32Prefix string, any *anypb.Any) (string, error) {
	pk, err := r.New(any)
	if err != nil {
		return "", err
	}
	return bech32.ConvertAndEncode(bech32Prefix, pk.Address())
}
