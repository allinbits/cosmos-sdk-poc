package crypto

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
)

func NewDefaultPubKeyResolver(supportedPubKeys []PubKey) *PubKeyResolver {
	known := make(map[string]struct{})
	for _, pk := range supportedPubKeys {
		known[fmt.Sprintf("/%s", pk.ProtoReflect().Descriptor().FullName())] = struct{}{}
	}
	known["/cosmos.crypto.secp256k1.PubKey"] = struct{}{}
	return &PubKeyResolver{knownPubKeys: known}
}

// PubKeyResolver resolves types and provides pubkeys
type PubKeyResolver struct {
	knownPubKeys map[string]struct{}
}

func (r *PubKeyResolver) New(any *anypb.Any) (PubKey, error) {

	// TODO(fdymylja): must fetch abs.Path
	if _, exists := r.knownPubKeys[any.TypeUrl]; !exists {
		return nil, fmt.Errorf("unknown pub key: %s", any.TypeUrl)
	}
	rawPK, err := anypb.UnmarshalNew(any, proto.UnmarshalOptions{Resolver: pubKeyAliasResolver{}})
	if err != nil {
		return nil, err
	}
	pk, ok := rawPK.(PubKey) // attempt to cast to real thin
	if !ok {
		panic(fmt.Errorf("%T does not implement PubKey", pk))
	}
	return pk, nil
}

func (r *PubKeyResolver) Address(bech32Prefix string, any *anypb.Any) (string, error) {
	pk, err := r.New(any)
	if err != nil {
		return "", err
	}
	return bech32.ConvertAndEncode(bech32Prefix, pk.Address())
}

type pubKeyAliasResolver struct {
}

func (s pubKeyAliasResolver) FindMessageByName(message protoreflect.FullName) (protoreflect.MessageType, error) {
	return nil, fmt.Errorf("tx: unsupported")
}

func (s pubKeyAliasResolver) FindMessageByURL(url string) (protoreflect.MessageType, error) {
	switch url {
	case "/cosmos.crypto.secp256k1.PubKey":
		return protoregistry.GlobalTypes.FindMessageByURL("/tmos.x.authn.crypto.v1alpha1.PubKey")
	default:
		return protoregistry.GlobalTypes.FindMessageByURL(url)
	}
}

func (s pubKeyAliasResolver) FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error) {
	return nil, fmt.Errorf("not implemented")
}
func (s pubKeyAliasResolver) FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error) {
	return nil, fmt.Errorf("not implemented")
}
