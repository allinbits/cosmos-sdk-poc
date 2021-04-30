package tx

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/fdymylja/tmos/module/x/authn/crypto"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
)

// NewDecoder instantiates a new *Decoder instance
// with a default crypto.PubKey set
func NewDecoder(bech32Prefix string) *Decoder {
	return &Decoder{
		pubKeyResolver: crypto.NewDefaultPubKeyResolver(),
		bech32pfx:      bech32Prefix,
	}
}

// Decoder is used to decode transactions
type Decoder struct {
	pubKeyResolver *crypto.PubKeyResolver
	bech32pfx      string
}

// Decode decodes the transaction bytes into an authentication.Tx
func (d *Decoder) Decode(txBytes []byte) (authentication.Tx, error) {
	rawTx := new(v1alpha1.TxRaw)
	err := unmarshalAndRejectUnknowns(txBytes, rawTx)
	if err != nil {
		return nil, err
	}
	// check for 0 lenghts
	switch {
	case len(rawTx.BodyBytes) == 0:
		return nil, fmt.Errorf("tx: empty body bytes")
	case len(rawTx.AuthInfoBytes) == 0:
		return nil, fmt.Errorf("tx: empty auth info bytes")
	case len(rawTx.Signatures) == 0:
		return nil, fmt.Errorf("tx: empty signatures")
	}
	txBody := new(v1alpha1.TxBody)
	err = unmarshalAndRejectUnknowns(rawTx.BodyBytes, txBody)
	if err != nil {
		return nil, err
	}
	authInfo := new(v1alpha1.AuthInfo)
	err = unmarshalAndRejectUnknowns(rawTx.AuthInfoBytes, authInfo)
	if err != nil {
		return nil, err
	}
	// check if fees are set
	if authInfo.Fee == nil {
		return nil, fmt.Errorf("tx: missing fees")
	}
	if authInfo.Fee.Amount == nil {
		return nil, fmt.Errorf("tx: missing fee amount")
	}
	// get transitions from body
	transitions, err := getTransitions(txBody)
	if err != nil {
		return nil, err
	}
	raw := &v1alpha1.Tx{
		Body:       txBody,
		AuthInfo:   authInfo,
		Signatures: rawTx.Signatures,
	}
	// get signers
	payer, signers, pubKeys, err := d.authInfo(raw.AuthInfo)
	if err != nil {
		return nil, err
	}
	return &Tx{
		raw:         raw,
		bytes:       txBytes,
		transitions: transitions,
		signers:     signers,
		pubKeys:     pubKeys,
		payer:       payer,
	}, nil
}

func (d *Decoder) authInfo(info *v1alpha1.AuthInfo) (string, *authentication.Subjects, []crypto.PubKey, error) {
	if len(info.SignerInfos) == 0 {
		return "", nil, nil, fmt.Errorf("tx: no signer provided")
	}

	subjects := authentication.NewSubjects()

	pubKeys := make([]crypto.PubKey, len(info.SignerInfos))
	feePayer := ""
	for i, sig := range info.SignerInfos {
		if sig.PublicKey == nil {
			return "", nil, nil, fmt.Errorf("tx: pubkey at index %d is nil", i)
		}
		pk, err := d.pubKeyResolver.New(sig.PublicKey)
		if err != nil {
			return "", nil, nil, fmt.Errorf("tx: unable to resolve public key at index %d: %w", i, err)
		}
		addr, err := bech32.ConvertAndEncode(d.bech32pfx, pk.Address())
		if err != nil {
			return "", nil, nil, fmt.Errorf("tx: unable to bechify address of public key at index %d: %w", i, err)
		}
		subjects.Add(addr)
		pubKeys[i] = pk
		// set fee payer as first signer
		if i == 0 {
			feePayer = addr
		}
	}
	// override fee payer if provided
	if info.Fee == nil {
		return feePayer, subjects, pubKeys, nil
	}
	if info.Fee.Payer != "" {
		feePayer = info.Fee.Payer
	}
	return feePayer, subjects, pubKeys, nil
}
