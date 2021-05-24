package indexes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_indexList(t *testing.T) {
	x := indexList{[]byte("osu-king"), []byte("brain-power"), []byte("noma")}
	b := x.marshal()
	u := new(indexList)
	err := u.unmarshal(b)
	require.NoError(t, err)
	require.Equal(t, x, *u)
}

func Test_indexerKey(t *testing.T) {
	x := &indexObjectWithSecondaryKey{
		objectPrefix:      []byte("some-proto-object"),
		indexPrefix:       []byte("some-index-name"),
		secondaryKeyValue: []byte("some-index-value"),
		primaryKey:        []byte("some-primary-key"),
	}
	b := x.marshal()

	t.Logf("%s", b)
	y := &indexObjectWithSecondaryKey{}
	err := y.unmarshal(b)
	require.NoError(t, err)
	require.Equal(t, x, y)
}

func Test_typePrefixedKey_bytes(t *testing.T) {
	x := typePrefixedKey{
		primaryKey: []byte("primary_key"),
		typePrefix: []byte("osu.v1alpha1.King"),
	}
	b := x.bytes()
	t.Logf("%s", b)
}
