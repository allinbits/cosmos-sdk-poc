package authn

import (
	"fmt"

	"github.com/btcsuite/btcutil/bech32"
	abciv1alpha1 "github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/crypto"
	"github.com/fdymylja/tmos/module/x/authn/tx"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	bankv1alpha1 "github.com/fdymylja/tmos/module/x/bank/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
)

func newAuthenticator(c module.Client) authenticator {
	return authenticator{
		abci:         abciv1alpha1.NewClient(c),
		auth:         v1alpha1.NewClient(c),
		pkRes:        crypto.NewDefaultPubKeyResolver(),
		bech32Prefix: "test",
		c:            c,
	}
}

type authenticator struct {
	abci         *abciv1alpha1.Client
	auth         *v1alpha1.Client
	decoder      *tx.Decoder
	bech32Prefix string
	c            module.Client
}

// Authenticate authenticates a transaction.
// TODO(fdymylja): this is done thinking how the sdk was originally designed
// what in reality should be done here is that the runtime runs a special type called
// Transaction Admission Controllers - they check (can read state) if the request is valid
// After that the write execution flow is run - for example sending coins to fee collector
// setting pubkeys and all of that.
func (a authenticator) Authenticate(txBytes []byte) (subjects []string, transitions []meta.StateTransition, err error) {
	rawTx, transitions, err := tx.Decode(txBytes)
	if err != nil {
		return nil, nil, err
	}
	// reject extension options
	// check if fee matches the minimum of mempool (TODO only checktx || !simulate)
	// validate basic ? not needed here - we run validate basic b4
	// check timeout height
	currentBlock, err := a.abci.GetCurrentBlock()
	if err != nil {
		return nil, nil, err
	}
	if currentBlock.BlockNumber > rawTx.Body.TimeoutHeight {
		return nil, nil, fmt.Errorf("invalid block height")
	}
	// validate memo
	params, err := a.auth.GetParams()
	if err != nil {
		return nil, nil, err
	}
	memoLength := uint64(len(rawTx.Body.Memo))
	if memoLength > params.MaxMemoCharacters {
		return nil, nil, fmt.Errorf("%w: maximum memo lenght is %d got %d", runtime.ErrBadRequest, params.MaxMemoCharacters, memoLength)
	}
	// consume gas for tx TODO
	// deduct fee:
	// TODO (fdymylja): here we're doing things in the old "sdk way" how this should be handled
	// is that - we have some module intercepting this request and collecting fees and sending them
	// to the collector defined by THAT module. (Not at auth level)
	// for instance - in cosmos-sdk - distribution receives the "fee_collector" name
	// and then uses that.
	// with the new model:
	// - Distribution creates at genesis the "fee_collector"
	// - Distribution declares to intercept Authentication requests
	// - Before authn flow execution: checks if the feepayer has enough $
	// - After authn flow execution: send coins from feepayer to fee_collector.
	fp, err := a.getFeePayer(rawTx)
	if err != nil {
		return nil, nil, err
	}
	err = a.c.Deliver(&bankv1alpha1.MsgSendCoins{
		FromAddress: fp,
		ToAddress:   "collector",
		Amount:      rawTx.AuthInfo.Fee.Amount,
	})
	if err != nil {
		return nil, nil, err
	}
	err = a.setPubKeys(rawTx)
	if err != nil {
		return nil, nil, err
	}
	// validate number of sigs
	err = a.validateSigs()
	if err != nil {
		return nil, nil, err
	}
	// consume gas for signatures TODO
	// verify sig
	// increment sequence
	return nil, nil, nil
}

func (a authenticator) DecodeTx(txBytes []byte) (transitions []meta.StateTransition, err error) {
	_, transitions, err = tx.Decode(txBytes)
	if err != nil {
		return
	}
	return
}

// getFeePayer returns the fee payer given a transaction
func (a authenticator) getFeePayer(tx *v1alpha1.Tx) (string, error) {
	feePayer := tx.AuthInfo.Fee.Payer
	if feePayer != "" {
		return feePayer, nil
	}
	// if fee payer was not set we get the fees from the first signer
	// TODO: feegrant
	// we resolve the first signer public key into address
	signer := tx.AuthInfo.SignerInfos[0].PublicKey
	pk, err := a.pkRes.New(signer)
	if err != nil {
		return "", fmt.Errorf("invalid signer pub key: %w", err)
	}
	bech32Addr, err := bech32.Encode(a.bech32Prefix, pk.Address())
	if err != nil {
		return "", fmt.Errorf("unable to bechify pub key address: %w", err)
	}
	return bech32Addr, nil
}

func (a authenticator) setPubKeys(tx *v1alpha1.Tx) error {
	// we check if every signing account exists
	for _, sigInfo := range tx.AuthInfo.SignerInfos {
		addr, err := a.pkRes.Address(a.bech32Prefix, sigInfo.PublicKey)
		if err != nil {
			return err
		}
		// we get the account and check for existence
		acc, err := a.auth.GetAccount(addr)
		if err != nil {
			return err
		}
		if acc.PubKey != nil {
			continue
		}
		acc.PubKey = sigInfo.PublicKey
		err = a.c.Update(acc)
		if err != nil {
			return err
		}
	}

	return nil
}
