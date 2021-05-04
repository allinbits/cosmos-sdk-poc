package module

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types"
	signing2 "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	legacybank "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func Sign(pkHex string, chainID, send, recv string, coins types.Coins) (signedTxBytes []byte, err error) {
	pkB, err := hex.DecodeString(pkHex)
	if err != nil {
		return
	}
	pk := secp256k1.PrivKey{Key: pkB}

	encoding := simapp.MakeTestEncodingConfig()

	msg := &legacybank.MsgSend{
		FromAddress: send,
		ToAddress:   recv,
		Amount:      coins,
	}

	builder := encoding.TxConfig.NewTxBuilder()
	builder.SetFeeAmount(types.NewCoins(types.NewCoin("test", types.NewInt(50))))
	builder.SetGasLimit(5000)

	err = builder.SetMsgs(msg)

	signMode := encoding.TxConfig.SignModeHandler().DefaultMode()

	signerData := signing.SignerData{
		ChainID:       chainID,
		AccountNumber: 0,
		Sequence:      0,
	}

	sigData := signing2.SingleSignatureData{
		SignMode:  signMode,
		Signature: nil,
	}

	sig := signing2.SignatureV2{
		PubKey:   pk.PubKey(),
		Data:     &sigData,
		Sequence: 0,
	}

	err = builder.SetSignatures(sig)
	if err != nil {
		return
	}

	bytesToSign, err := encoding.TxConfig.SignModeHandler().GetSignBytes(signMode, signerData, builder.GetTx())
	if err != nil {
		return
	}

	signature, err := pk.Sign(bytesToSign)
	if err != nil {
		return
	}

	sigData = signing2.SingleSignatureData{
		SignMode:  signMode,
		Signature: signature,
	}

	sig = signing2.SignatureV2{
		PubKey:   pk.PubKey(),
		Data:     &sigData,
		Sequence: 0,
	}

	err = builder.SetSignatures(sig)
	if err != nil {
		return
	}

	bytes, err := encoding.TxConfig.TxEncoder()(builder.GetTx())
	if err != nil {
		return
	}

	return bytes, nil
}
