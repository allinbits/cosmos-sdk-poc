package tx

import (
	"fmt"

	"github.com/fdymylja/tmos/module/x/authn/v1alpha1"
	"github.com/fdymylja/tmos/runtime/meta"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

type errUnknownField struct {
	Number int
	Type   string
}

func (e errUnknownField) Error() string {
	return fmt.Sprintf("Tx: unkown field in provided protobuf message of type %s number %d", e.Type, e.Number)
}

func unmarshalAndRejectUnknowns(b []byte, pb proto.Message) error {
	err := proto.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	if len(pb.ProtoReflect().GetUnknown()) != 0 {
		n, t, i := protowire.ConsumeTag(pb.ProtoReflect().GetUnknown())
		if i <= 0 {
			return fmt.Errorf("Tx: unable to consume tags for unknown fields reporting: %w", protowire.ParseError(i))
		}
		return errUnknownField{
			Number: (int)(n),
			Type:   fmt.Sprintf("%d", t), // TODO convert to stringified type
		}
	}
	return nil
}

func getTransitions(body *v1alpha1.TxBody) ([]meta.StateTransition, error) {
	if len(body.Messages) == 0 {
		return nil, fmt.Errorf("Tx: no messages in Tx")
	}
	transitions := make([]meta.StateTransition, len(body.Messages))
	for i, msg := range body.Messages {
		// unmarshal from any
		rawPb, err := msg.UnmarshalNew()
		if err != nil {
			return nil, fmt.Errorf("Tx: unable to unmarshal message %s to known type", msg.TypeUrl)
		}
		transition, ok := rawPb.(meta.StateTransition)
		if !ok {
			return nil, fmt.Errorf("Tx: %s is not a meta.StateTransition", msg.TypeUrl)
		}
		transitions[i] = transition
	}
	return transitions, nil
}
