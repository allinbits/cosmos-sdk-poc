package authn

import (
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/x/authn/tx"
)

func newTxDecoder() txDecoder {
	return txDecoder{
		decoder: tx.NewDecoder("test"),
	}
}

type txDecoder struct {
	decoder *tx.Decoder
}

func (a txDecoder) DecodeTx(txBytes []byte) (tx authentication.Tx, err error) {
	return a.decoder.Decode(txBytes)
}
