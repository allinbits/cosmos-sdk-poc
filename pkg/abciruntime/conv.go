package abciruntime

import (
	abcialias "github.com/fdymylja/tmos/apis/abci/tendermint/abci"
	gogoproto "github.com/gogo/protobuf/proto"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/proto"
)

// TODO(fdymylja): im a lazy programmer PLEASE DO THIS DECENTLY

func convBeginBlock(block abcitypes.RequestBeginBlock) *abcialias.RequestBeginBlock {
	b, err := gogoproto.Marshal(&block)
	if err != nil {
		panic(err)
	}
	converted := new(abcialias.RequestBeginBlock)
	err = proto.Unmarshal(b, converted)
	if err != nil {
		panic(err)
	}
	return converted
}

/*
func convBeginBlock(block abcitypes.RequestBeginBlock)  *abcialias.RequestBeginBlock {
	return &abcialias.RequestBeginBlock{
		Hash:                block.Hash,
		Header:              convHeader(block.Header),
		LastCommitInfo:      &abcialias.LastCommitInfo{
			Round: block.LastCommitInfo.Round,
			Votes: ,
		},
		ByzantineValidators: nil,
	}
}

func convHeader(header tmtypes.Header) *typesalias.Header {
	return &typesalias.Header{
		Version:            &version.Consensus{
			Block: header.Version.Block,
			App:   header.Version.App,
		},
		ChainId:            header.ChainID,
		Height:             header.Height,
		Time:               timestamppb.New(header.Time),
		LastBlockId:        nil,
		LastCommitHash:     nil,
		DataHash:           nil,
		ValidatorsHash:     nil,
		NextValidatorsHash: nil,
		ConsensusHash:      nil,
		AppHash:            nil,
		LastResultsHash:    nil,
		EvidenceHash:       nil,
		ProposerAddress:    nil,
	}
}

*/
