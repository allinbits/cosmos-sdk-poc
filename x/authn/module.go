package authn

import (
	rbacv1alpha1 "github.com/fdymylja/tmos/core/rbac/v1alpha1"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
	"github.com/fdymylja/tmos/x/authn/v1alpha1"
)

// Module implements the authentication.Module
type Module struct {
	txDecoder authentication.TxDecoder
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Initialize(c module.Client) module.Descriptor {
	m.txDecoder = newTxDecoder()

	builder := module.NewDescriptorBuilder()

	builder.Named("authn").
		HandlesStateTransition(&v1alpha1.MsgCreateAccount{}, NewCreateAccountController(c), true).
		HandlesAdmission(&v1alpha1.MsgCreateAccount{}, NewCreateAccountAdmissionController()).
		OwnsStateObject(&v1alpha1.Account{}, v1alpha1.AccountSchema).
		OwnsStateObject(&v1alpha1.Params{}, v1alpha1.ParamsSchema).
		OwnsStateObject(&v1alpha1.CurrentAccountNumber{}, v1alpha1.CurrentAccountNumberSchema).
		WithGenesis(genesis{c: c}).
		NeedsStateTransition(&rbacv1alpha1.MsgBindRole{})

	// add admission controllers, they will only read state
	// never modify it.
	builder.
		WithAuthAdmissionHandler(mempoolFee{}).                // verifies if fee matches the minimum
		WithAuthAdmissionHandler(newAccountExists(c)).         // verifies that all signer accounts exist
		WithAuthAdmissionHandler(newTimeoutBlockExtension(c)). // verifies if tx is not timed-out compared to block
		WithAuthAdmissionHandler(newValidateMemoExtension(c)). // validates memo length
		WithAuthAdmissionHandler(newValidateSigCount(c)).      // validates number of signatures
		WithAuthAdmissionHandler(newSigVerifier(c))            // validate signatures

	// add transition controllers for tx, they CAN modify state after
	// a tx is authenticated
	builder.
		WithPostAuthenticationHandler(newConsumeGasForTxSize(c)). // consumes gas for tx size
		WithPostAuthenticationHandler(newSetPubKeys(c)).          // sets pub keys
		WithPostAuthenticationHandler(newIncreaseSequence(c))     // increases sequence

	return builder.Build()
}

func (m *Module) GetTxDecoder() authentication.TxDecoder {
	return m.txDecoder
}
