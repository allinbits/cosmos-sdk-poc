package authn

import (
	"fmt"

	abciv1alpha1 "github.com/fdymylja/tmos/module/abci/v1alpha1"
	"github.com/fdymylja/tmos/module/x/authn/tx"
	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/module"
)

func newAuthenticator(c module.Client) authenticator {
	return authenticator{
		abci: abciv1alpha1.NewClient(c),
		auth: v1alpha1.NewClient(c),
	}
}

type authenticator struct {
	abci *abciv1alpha1.Client
	auth *v1alpha1.Client
}

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
	// deduct fee

	// set pub key if not exists
	// validate number of sigs
	// consume gas for signatures
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
