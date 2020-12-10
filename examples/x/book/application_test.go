package book

import (
	"github.com/fdymylja/cosmos-os/examples/x/author"
	"github.com/fdymylja/cosmos-os/internal/runtime"
	"github.com/fdymylja/cosmos-os/pkg/client"
	"testing"
)

func TestApplication(t *testing.T) {
	rt := runtime.NewRuntime()
	rt.LoadApplication(Application{})
	rt.LoadApplication(author.Application{})

	cl := client.NewClient(rt)

	req := &author.MsgRegisterAuthor{
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
	query := &author.QueryAuthorRequest{
		Name: "hello",
	}
	resp := &author.QueryAuthorResponse{}
	err = cl.Query(1, query, resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp.Author)

	// deliver new book tx
	registerBook := &MsgRegisterBook{
		Author: query.Name,
		Name:   "anewbook",
	}
	err = cl.Tx(registerBook)
	if err != nil {
		t.Fatal(err)
	}

	rt.Commit()

	bookResp := new(QueryBookResponse)
	err = cl.Query(2, &QueryBookRequest{
		Name: registerBook.Name,
	}, bookResp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", bookResp.Book)
}
