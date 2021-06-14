package forge

import (
	"testing"

	"github.com/fdymylja/tmos/pkg/protoutils/prototest"
	"github.com/fdymylja/tmos/testdata/testpb"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
)

func TestList(t *testing.T) {
	forgedListType, err := List(&testpb.SimpleMessage{}, protoregistry.GlobalFiles)
	require.NoError(t, err)

	fd := forgedListType.Descriptor().Fields().Get(0)
	require.Equal(t, protoreflect.FullName("dynamic.testdata.SimpleMessageList"), forgedListType.Descriptor().FullName())
	require.Equal(t, protoreflect.FullName("dynamic.testdata.SimpleMessageList.list"), fd.FullName())
	require.True(t, fd.IsList())

	// now we test if it's compatible
	m := &testpb.SimpleMessageList{List: []*testpb.SimpleMessage{
		{
			A: "1",
			B: 2,
			C: [][]byte{[]byte("hello")},
		},
		{
			A: "2",
			B: 3,
			C: [][]byte{[]byte("test")},
		},
	}}

	jsonBytes, err := protojson.Marshal(m)
	require.NoError(t, err)

	newDynamicList := dynamicpb.NewMessage(forgedListType.Descriptor())
	err = protojson.Unmarshal(jsonBytes, newDynamicList)
	require.NoError(t, err)

	v := newDynamicList.Get(forgedListType.Descriptor().Fields().Get(0))
	require.NotNil(t, v)
	list := v.List()
	require.Equal(t, list.Len(), len(m.List))

	// assert equality
	for i := 0; i < list.Len(); i++ {
		got := list.Get(i)
		expected := m.List[i]
		prototest.Equal(t, got.Interface().(*dynamicpb.Message), expected)
	}
}
