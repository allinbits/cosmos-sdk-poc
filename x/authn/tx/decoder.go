package tx

import (
	"fmt"
	"github.com/fdymylja/tmos/x/authn/crypto/secp256k1"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/authentication/user"
	crypto2 "github.com/fdymylja/tmos/x/authn/crypto"
	v1alpha12 "github.com/fdymylja/tmos/x/authn/v1alpha1"
)

// NewDecoder instantiates a new *Decoder instance
// with a default crypto.PubKey set
func NewDecoder(bech32Prefix string) *Decoder {
	return &Decoder{
		pubKeyResolver: crypto2.NewDefaultPubKeyResolver([]crypto2.PubKey{&secp256k1.PubKey{}}),
		bech32pfx:      bech32Prefix,
	}
}

// Decoder is used to decode transactions
type Decoder struct {
	pubKeyResolver *crypto2.PubKeyResolver
	bech32pfx      string
}

// Decode decodes the transaction bytes into an authentication.Tx
func (d *Decoder) Decode(txBytes []byte) (authentication.Tx, error) {
	rawTx := new(v1alpha12.TxRaw)
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
	txBody := new(v1alpha12.TxBody)
	err = unmarshalAndRejectUnknowns(rawTx.BodyBytes, txBody)
	if err != nil {
		return nil, err
	}
	authInfo := new(v1alpha12.AuthInfo)
	err = unmarshalAndRejectUnknowns(rawTx.AuthInfoBytes, authInfo)
	if err != nil {
		return nil, err
	}
	// check that signatures match signer infos
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
	raw := &v1alpha12.Tx{
		Body:       txBody,
		AuthInfo:   authInfo,
		Signatures: rawTx.Signatures,
	}
	// get signers
	payer, signers, pubKeys, err := d.authInfo(raw.AuthInfo, raw.Signatures)
	if err != nil {
		return nil, err
	}
	return &Wrapper{
		txRaw:       rawTx,
		raw:         raw,
		bytes:       txBytes,
		transitions: transitions,
		signers:     user.NewUsersFromString(signers...),
		pubKeys:     pubKeys,
		payer:       payer,
	}, nil
}

func (d *Decoder) authInfo(info *v1alpha12.AuthInfo, signatures [][]byte) (string, []string, []Signer, error) {
	if len(info.SignerInfos) == 0 {
		return "", nil, nil, fmt.Errorf("tx: no signer provided")
	}
	if len(signatures) != len(info.SignerInfos) {
		return "", nil, nil, fmt.Errorf("tx: signers and signatures number mimsatch")
	}
	var usersStr []string

	var signers []Signer
	feePayer := ""
	for i, sig := range info.SignerInfos {
		// check if pk set
		if sig.PublicKey == nil {
			return "", nil, nil, fmt.Errorf("tx: pubkey at index %d is nil", i)
		}
		// check if sig mode set
		if sig.ModeInfo == nil {
			return "", nil, nil, fmt.Errorf("tx: no mode info provided at index %d", i)
		}
		// check if sig mode is direct TODO support all else
		mode := sig.ModeInfo.GetSingle()
		if mode == nil {
			return "", nil, nil, fmt.Errorf("tx: unsupported sign mode %s", sig.ModeInfo.GetSum())
		}
		if mode.Mode != v1alpha12.SignMode_SIGN_MODE_DIRECT {
			return "", nil, nil, fmt.Errorf("tx: unsupported sign mode %s", mode.Mode)
		}
		pk, err := d.pubKeyResolver.New(sig.PublicKey)
		if err != nil {
			return "", nil, nil, fmt.Errorf("tx: unable to resolve public key at index %d: %w", i, err)
		}
		addr, err := bech32.ConvertAndEncode(d.bech32pfx, pk.Address())
		if err != nil {
			return "", nil, nil, fmt.Errorf("tx: unable to bechify address of public key at index %d: %w", i, err)
		}
		usersStr = append(usersStr, addr)
		signers = append(signers, Signer{
			Address:   addr,
			PubKey:    pk,
			Signature: signatures[i],
			Sequence:  sig.Sequence,
		})
		// set fee payer as first signer
		if i == 0 {
			feePayer = addr
		}
	}
	// override fee payer if provided
	if info.Fee == nil {
		return feePayer, usersStr, signers, nil
	}
	if info.Fee.Payer != "" {
		feePayer = info.Fee.Payer
	}
	return feePayer, usersStr, signers, nil
}
