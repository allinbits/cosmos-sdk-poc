package tx

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	coin "github.com/fdymylja/tmos/module/core/coin/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/crypto"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
)

var _ authentication.Tx = (*Wrapper)(nil)

type Signer struct {
	Address       string
	PubKey        crypto.PubKey
	SignatureData signing.SignatureData
}

// Wrapper wraps the raw *v1alpha1.Tx but contains parsed information
// regarding pub keys and such.
type Wrapper struct {
	raw         *v1alpha1.Tx
	bytes       []byte
	transitions []meta.StateTransition
	signers     *authentication.Subjects
	pubKeys     []Signer
	payer       string
}

func (t *Wrapper) StateTransitions() []meta.StateTransition {
	return t.transitions
}

func (t *Wrapper) Subjects() *authentication.Subjects {
	return t.signers
}

func (t *Wrapper) Fee() []*coin.Coin {
	return t.raw.AuthInfo.Fee.Amount
}

func (t *Wrapper) Payer() string {
	return t.payer
}

func (t *Wrapper) Raw() interface{} {
	return t.raw
}

func (t *Wrapper) RawBytes() []byte {
	return t.bytes
}

// Signers returns a map containing the account identifier (address) and the public key the user used to sign.
func (t *Wrapper) Signers() []Signer {
	return t.pubKeys
}

func (t *Wrapper) Signatures() ([]signing.SignatureV2, error) {
	signerInfos := t.raw.AuthInfo.SignerInfos
	sigs := t.raw.Signatures
	signers := t.Signers()
	n := len(signerInfos)
	res := make([]signing.SignatureV2, n)

	for i, si := range signerInfos {
		// handle nil signatures (in case of simulation)
		var err error
		sigData, err := ModeInfoAndSigToSignatureData(&v1alpha1.ModeInfo{
			Sum: si.ModeInfo.Sum,
		}, sigs[i])
		if err != nil {
			return nil, err
		}
		res[i] = signing.SignatureV2{
			PubKey:   signers[i].PubKey,
			Data:     sigData,
			Sequence: si.GetSequence(),
		}

	}
	return res, nil
}

// ModeInfoAndSigToSignatureData converts a ModeInfo and raw bytes signature to a SignatureData or returns
// an error
func ModeInfoAndSigToSignatureData(modeInfo *v1alpha1.ModeInfo, sig []byte) (signing.SignatureData, error) {
	switch modeInfo := modeInfo.Sum.(type) {
	case *v1alpha1.ModeInfo_Single_:
		return &signing.SingleSignatureData{
			SignMode:  signing.SignMode(modeInfo.Single.Mode),
			Signature: sig,
		}, nil
	/*
		case *v1alpha1.ModeInfo_Multi_:
			multi := modeInfo.Multi

			sigs, err := decodeMultisignatures(sig)
			if err != nil {
				return nil, err
			}

			sigv2s := make([]signing.SignatureData, len(sigs))
			for i, mi := range multi.ModeInfos {
				sigv2s[i], err = ModeInfoAndSigToSignatureData(mi, sigs[i])
				if err != nil {
					return nil, err
				}
			}

			return &signing.MultiSignatureData{
				BitArray:   multi.Bitarray,
				Signatures: sigv2s,
			}, nil
	*/
	default:
		return nil, fmt.Errorf("unsupported signing mode: %s", modeInfo)
	}
}
