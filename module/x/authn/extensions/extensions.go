package extensions

import (
	"fmt"

	kmultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	abciv1alpha1 "github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/crypto"
	"github.com/fdymylja/tmos/module/x/authn/tx"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
)

func New(c module.Client) module.AuthenticationExtension {
	return authExtension{c: c}
}

type authExtension struct {
	c module.Client
}

func (a authExtension) Initialize(builder *module.AuthenticationExtensionBuilder) {
	// add admission controllers, they will only read state
	// never modify it.
	builder.
		WithAdmissionController(mempoolFee{}).                  // verifies if fee matches the minimum
		WithAdmissionController(newTimeoutBlockExtension(a.c)). // verifies if tx is not timed-out compared to block
		WithAdmissionController(newValidateMemoExtension(a.c)). // validates memo length
		WithAdmissionController(newValidateSigCount(a.c))       // validates number of signatures
	// add transition controllers for tx
	builder.
		WithTransitionController(newConsumeGasForTxSize(a.c)) // consumes gas for tx size

}

func newTimeoutBlockExtension(c module.Client) timeoutBlockExtension {
	return timeoutBlockExtension{abci: abciv1alpha1.NewClient(c)}
}

type timeoutBlockExtension struct {
	abci *abciv1alpha1.Client
}

func (t timeoutBlockExtension) Validate(request authentication.ValidateRequest) (authentication.ValidateResponse, error) {
	tx := request.Tx.Raw().(*v1alpha1.Tx)
	currentBlock, err := t.abci.GetCurrentBlock()
	if err != nil {
		return authentication.ValidateResponse{}, err
	}
	if currentBlock.BlockNumber > tx.Body.TimeoutHeight {
		return authentication.ValidateResponse{}, fmt.Errorf("invalid block height")
	}
	return authentication.ValidateResponse{}, nil
}

func newValidateMemoExtension(c module.Client) validateMemoExtension {
	return validateMemoExtension{c: v1alpha1.NewClient(c)}
}

type validateMemoExtension struct {
	c *v1alpha1.Client
}

func (e validateMemoExtension) Validate(request authentication.ValidateRequest) (authentication.ValidateResponse, error) {
	tx := request.Tx.Raw().(*v1alpha1.Tx)
	memo := tx.Body.Memo
	memoLength := len(memo)
	params, err := e.c.GetParams()
	if err != nil {
		return authentication.ValidateResponse{}, nil
	}
	if uint64(memoLength) > params.MaxMemoCharacters {
		return authentication.ValidateResponse{}, fmt.Errorf("invalid memo length")
	}
	return authentication.ValidateResponse{}, nil
}

// mempoolFee is used to check if the transaction meets the minimum mempool requirements
type mempoolFee struct {
}

func (e mempoolFee) Validate(request authentication.ValidateRequest) (authentication.ValidateResponse, error) {
	return authentication.ValidateResponse{}, nil
}

func newValidateSigCount(c module.Client) sigCount {
	return sigCount{c: v1alpha1.NewClient(c)}
}

type sigCount struct {
	c *v1alpha1.Client
}

func (s sigCount) Validate(request authentication.ValidateRequest) (authentication.ValidateResponse, error) {
	wrapper := request.Tx.(*tx.Wrapper) // TODO: should we return an error? or simply skip validation?
	params, err := s.c.GetParams()
	if err != nil {
		return authentication.ValidateResponse{}, err
	}
	pubKeys := wrapper.PubKeys()
	sigs := 0
	for _, pk := range pubKeys {
		subKeys := s.countSubKeys(pk)
		sigs += subKeys
		if uint64(sigs) > params.TxSigLimit {
			return authentication.ValidateResponse{}, fmt.Errorf("number of maximum signatures is %d got %d", params.TxSigLimit, sigs)
		}
	}
	return authentication.ValidateResponse{}, nil
}

func (s sigCount) countSubKeys(pk crypto.PubKey) int {
	v, ok := pk.(*kmultisig.LegacyAminoPubKey)
	if !ok {
		return 1
	}

	numKeys := 0
	for _, subkey := range v.GetPubKeys() {
		numKeys += s.countSubKeys(subkey)
	}
	return numKeys
}

func newConsumeGasForTxSize(c module.Client) consumeGasForTxSize {
	return consumeGasForTxSize{}
}

type consumeGasForTxSize struct {
}

func (s consumeGasForTxSize) Deliver(req authentication.DeliverRequest) (authentication.DeliverResponse, error) {
	// TODO
	return authentication.DeliverResponse{}, nil
}

type setPubKeys struct {
	c module.Client
}

func (s setPubKeys) Deliver(req authentication.DeliverRequest) (authentication.DeliverResponse, error) {
	for _, sig := range req.Tx.Subjects().List() {
		acc := new(v1alpha1.Account)
		err := s.c.Get(meta.NewStringID(sig), acc)
		if err != nil {
			return authentication.DeliverResponse{}, err
		}
		// if pub key is set skip
		if acc.PubKey != nil {
			continue
		}
		// otherwise set the pub key for the given subject

	}
}
