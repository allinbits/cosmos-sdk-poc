package v1alpha1

import (
	"reflect"
	"testing"

	"github.com/tendermint/tendermint/abci/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/fdymylja/tmos/apis/core/runtime/v1alpha1"
	tx "github.com/fdymylja/tmos/apis/core/tx/v1alpha1"
	"github.com/fdymylja/tmos/pkg/runtime"
)

func BenchmarkReflectNew(b *testing.B) {
	c := &Balance{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		typ := reflect.TypeOf(c)
		v := reflect.New(typ)
		_, _ = v.Interface().(proto.Message)
	}
}

func BenchmarkProtoreflectNew(b *testing.B) {
	typ, err := protoregistry.GlobalTypes.FindMessageByURL("/accounting.v1alpha1.Balance")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := typ.New()
		v.Interface()
	}
}

func TestApplication(t *testing.T) {
	rtb := runtime.NewBuilder()
	err := rtb.Mount(&Application{})
	if err != nil {
		t.Fatal(err)
	}
	rt, err := rtb.Build()
	if err != nil {
		t.Fatal(err)
	}

	msg := &MsgSend{
		Currency:  "money",
		Amount:    100,
		Sender:    "me",
		Recipient: "you",
	}

	anyMsg, err := anypb.New(msg)
	if err != nil {
		t.Fatal(err)
	}
	// yikes
	tx := &tx.Tx{
		StateTransitions: []*anypb.Any{anyMsg},
	}
	anyTx, err := anypb.New(tx)
	if err != nil {
		t.Fatal(err)
	}
	// p3
	rawTx := &v1alpha1.RawTx{
		Body: anyTx,
	}
	// marshal
	b, err := proto.Marshal(rawTx)
	if err != nil {
		t.Fatal(err)
	}
	rt.DeliverTx(types.RequestDeliverTx{Tx: b})
}
