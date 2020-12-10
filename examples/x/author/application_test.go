package author

import (
	"github.com/fdymylja/cosmos-os/internal/runtime"
	"github.com/fdymylja/cosmos-os/pkg/client"
	"testing"
)

func TestApplication(t *testing.T) {
	rt := runtime.NewRuntime()
	rt.LoadApplication(Application{})

	cl := client.NewClient(rt)

	req := &MsgRegisterAuthor{
		Name:        "hello",
		DateOfBirth: "never",
	}

	err := cl.Tx(req)
	if err != nil {
		t.Fatal(err)
	}

	// commit state
	rt.Commit()

	// query
	query := &QueryAuthorRequest{
		Name: "hello",
	}
	resp := &QueryAuthorResponse{}
	err = cl.Query(1, query, resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp.Author)
}
