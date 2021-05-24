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
			wantErr:   ErrBadOptions,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := NewIndexer(tc.Object, tc.JsonField)
			require.ErrorIs(t, err, tc.wantErr)
		})
	}
}
