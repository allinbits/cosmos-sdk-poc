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
	x := &indexerKey{
		objectPrefix: []byte("some-proto-object"),
		indexName:    []byte("some-index-name"),
		indexValue:   []byte("some-index-value"),
		primaryKey:   []byte("some-primary-key"),
	}
	b := x.marshal()

	t.Logf("%s", b)
	y := &indexerKey{}
	err := y.unmarshal(b)
	require.NoError(t, err)
	require.Equal(t, x, y)
}
