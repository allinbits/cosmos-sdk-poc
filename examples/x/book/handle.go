package book

import (
	"github.com/fdymylja/cosmos-os/examples/x/author"
	"github.com/fdymylja/cosmos-os/pkg/application"
	"github.com/fdymylja/cosmos-os/pkg/codec"
)

func handleRegisterBook() application.DeliverFunc {
	return func(req application.DeliverRequest) (application.DeliverResponse, error) {
		msg := new(MsgRegisterBook)
		err := codec.Unmarshal(req.Request, msg)
		if err != nil {
			return application.DeliverResponse{}, err
		}

		// check if author exists
		writer := new(author.QueryAuthorResponse)
		err = req.Client.Query(&author.QueryAuthorRequest{Name: msg.Author}, writer)
		if err != nil {
			return application.DeliverResponse{}, err
		}
		// check if the book exists TODO missing has() from store
		// register book
		book := &Book{
			Author: msg.Author,
			Name:   msg.Name,
		}
		err = req.Store.Set([]byte(book.Name), book)
		if err != nil {
			return application.DeliverResponse{}, err
		}
		// done!
		return application.DeliverResponse{}, nil
	}
}
