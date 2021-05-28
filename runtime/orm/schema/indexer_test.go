package schema

import (
	"testing"

	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/testdata/testpb"
	"github.com/stretchr/testify/require"
)

func TestNewIndexer(t *testing.T) {
	type test struct {
		Object    meta.StateObject
		JsonField string
		wantErr   error
	}

	cases := map[string]test{
		"success": {
			Object:    &testpb.SimpleMessage{},
			JsonField: "a",
			wantErr:   nil,
		},
		"unknown field": {
			Object:    &testpb.SimpleMessage{},
			JsonField: "unknown",
			wantErr:   ErrBadDefinition,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := NewIndexer(tc.Object, tc.JsonField)
			require.ErrorIs(t, err, tc.wantErr)
		})
	}
}

func TestIndexer_EncodeInterface(t *testing.T) {
	type test struct {
		input   interface{}
		result  []byte
		wantErr error
	}
	cases := map[string]test{
		"success": {
			input:   "test",
			result:  []byte("test"),
			wantErr: nil,
		},
		"nil interface": {
			input:   nil,
			result:  nil,
			wantErr: ErrFieldTypeMismatch,
		},
		"type mismatch": {
			input:   0,
			result:  nil,
			wantErr: ErrFieldTypeMismatch,
		},
	}
	indexer, err := NewIndexer(&testpb.SimpleMessage{}, "a")
	require.NoError(t, err)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			result, err := indexer.EncodeInterface(tc.input)
			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
			}
			require.EqualValues(t, tc.result, result)
		})
	}
}
