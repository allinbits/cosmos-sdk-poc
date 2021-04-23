package tx

import (
	"errors"
	"fmt"

	"github.com/fdymylja/tmos/apis/x/authn/v1alpha1"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

type errUnknownField struct {
	Number int
	Type   string
}

func (e errUnknownField) Error() string {
	return fmt.Sprintf("tx: unkown field in provided protobuf message of type %s number %d", e.Type, e.Number)
}

var errUnknownFields = errors.New("tx: unknown fields")

func DecodeTx(b []byte) (*v1alpha1.Tx, error) {
	rawTx := new(v1alpha1.TxRaw)
	err := unmarshalAndRejectUnknowns(b, rawTx)
	if err != nil {
		return nil, err
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
	return &v1alpha1.Tx{
		Body:       txBody,
		AuthInfo:   authInfo,
		Signatures: rawTx.Signatures,
	}, nil
}

func unmarshalAndRejectUnknowns(b []byte, pb proto.Message) error {
	err := proto.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	if len(pb.ProtoReflect().GetUnknown()) != 0 {
		n, t, i := protowire.ConsumeTag(pb.ProtoReflect().GetUnknown())
		if i <= 0 {
			return fmt.Errorf("tx: unable to consume tags for unknown fields reporting: %w", protowire.ParseError(i))
		}
		return errUnknownField{
			Number: (int)(n),
			Type:   fmt.Sprintf("%d", t), // TODO convert to stringified type
		}
	}
	return nil
}
