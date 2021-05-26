package objects

import (
	"testing"

	"github.com/fdymylja/tmos/pkg/prototest"
	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/schema"
	"github.com/fdymylja/tmos/testdata/testpb"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	sch, err := schema.NewSchema(&testpb.SimpleMessage{}, schema.Definition{PrimaryKey: "a"})
	require.NoError(t, err)
	s := kv.NewBadger()
	store := NewStore(s)

	// test creation
	o1 := &testpb.SimpleMessage{A: "1", B: 2}
	require.NoError(t, store.Create(sch, o1))

	// test get
	got := &testpb.SimpleMessage{}
	require.NoError(t, store.Get(sch, meta.NewStringID(o1.A), got))
	prototest.Equal(t, o1, got)

	// test create same object
	require.ErrorIs(t, store.Create(sch, o1), orm.ErrAlreadyExists)

	// test get object does not exist
	require.ErrorIs(t, store.Get(sch, meta.NewStringID("does not exist for sure"), got), orm.ErrNotFound)

	// test update
	o2 := &testpb.SimpleMessage{A: o1.A, B: 3}
	require.NoError(t, store.Update(sch, o2))
	got.Reset()
	require.NoError(t, store.Get(sch, meta.NewStringID(o1.A), got))
	prototest.Equal(t, o2, got)

	// test update non existing object
	require.ErrorIs(t, store.Update(sch, &testpb.SimpleMessage{A: "not exists"}), orm.ErrNotFound)

	// test delete
	require.NoError(t, store.Delete(sch, o2))
	require.ErrorIs(t, store.Get(sch, meta.NewStringID(o2.A), got), orm.ErrNotFound)

	// test delete non existing object
	require.ErrorIs(t, store.Delete(sch, &testpb.SimpleMessage{A: "not exists"}), orm.ErrNotFound)
}
