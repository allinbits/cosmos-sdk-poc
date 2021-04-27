package v1alpha1

import (
	"testing"

	"github.com/fdymylja/tmos/module/abci/tendermint/abci"
	"github.com/fdymylja/tmos/module/abci/tendermint/types"
	"github.com/fdymylja/tmos/module/abci/tendermint/version"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestT(t *testing.T) {
	c := &MsgSetBeginBlockState{BeginBlock: &abci.RequestBeginBlock{
		Hash: nil,
		Header: &types.Header{
			Version:            &version.Consensus{},
			ChainId:            "xd",
			Height:             0,
			Time:               &timestamppb.Timestamp{},
			LastBlockId:        &types.BlockID{},
			LastCommitHash:     nil,
			DataHash:           nil,
			ValidatorsHash:     nil,
			NextValidatorsHash: nil,
			ConsensusHash:      nil,
			AppHash:            nil,
			LastResultsHash:    nil,
			EvidenceHash:       nil,
			ProposerAddress:    nil,
		},
		LastCommitInfo:      &abci.LastCommitInfo{},
		ByzantineValidators: nil,
	}}
	t.Logf("%s", c)
	b, err := protojson.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", b)
}
