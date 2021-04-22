package accounting

import (
	"testing"

	meta "github.com/fdymylja/tmos/apis/meta/v1alpha1"
	"github.com/fdymylja/tmos/apis/x/accounting/v1alpha1"
	"github.com/fdymylja/tmos/pkg/runtime"
)

func TestMsgSendHandler(t *testing.T) {
	builder := runtime.NewBuilder()
	app := &Module{}
	builder.MountApplication(app)
	rt := builder.Build()
	// initialization done
	// set some mock state
	app.client.Set(&v1alpha1.Balance{
		ObjectMeta: &meta.ObjectMeta{Id: "frojdi"},
		Amount:     1000,
	})

	err := rt.Deliver(nil, &v1alpha1.MsgSend{
		Sender:   "frojdi",
		Receiver: "jonathan",
		Amount:   100,
	})
	if err != nil {
		t.Fatal(err)
	}
	balance := &v1alpha1.Balance{ObjectMeta: &meta.ObjectMeta{Id: "frojdi"}}
	app.client.Get(balance)
	t.Logf("%s", balance)
}
