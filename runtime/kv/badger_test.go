package kv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBadgerIterator(t *testing.T) {
	s := NewBadger()
	s.Set([]byte("hi"), []byte("hello"))
	iter := s.Iterate([]byte("hi"), []byte{})
	require.True(t, iter.Valid())
}
