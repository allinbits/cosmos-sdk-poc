package schema

import (
	"testing"

	"github.com/fdymylja/tmos/testdata/testpb"
	"github.com/stretchr/testify/require"
)

func Test_parseObjectSchema(t *testing.T) {
	type test struct {
		options Definition
		wantErr error
	}

	cases := map[string]test{
		"success": {
			options: Definition{
				PrimaryKey:    "a",
				SecondaryKeys: []string{"b", "c"},
			},
			wantErr: nil,
		},
		"singleton with primary key": {
			options: Definition{
				Singleton:  true,
				PrimaryKey: "a",
			},
			wantErr: ErrBadOptions,
		},
		"primary key not found": {
			options: Definition{
				PrimaryKey: "not-found",
			},
			wantErr: ErrBadOptions,
		},
		"singleton with secondary key": {
			options: Definition{
				Singleton:     true,
				PrimaryKey:    "",
				SecondaryKeys: []string{"a"},
			},
			wantErr: ErrBadOptions,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := parseObjectSchema(&testpb.SimpleMessage{}, tc.options)
			require.ErrorIs(t, err, tc.wantErr)
		})
	}
}
