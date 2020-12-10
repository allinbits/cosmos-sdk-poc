package book

import (
	"github.com/fdymylja/cosmos-os/pkg/application"
	"github.com/golang/protobuf/proto"
)

func queryBook() application.QueryFunc {
	return func(req application.QueryRequest) (application.QueryResponse, error) {
		// unmarshal request
		bookRequest := new(QueryBookRequest)
		err := proto.Unmarshal(req.Request, bookRequest)
		if err != nil {
			return application.QueryResponse{}, err
		}
		// get book
		book := new(Book)
		err = req.Store.Get([]byte(bookRequest.Name), book)
		if err != nil {
			return application.QueryResponse{}, err
		}
		// marshal response
		resp, err := proto.Marshal(&QueryBookResponse{
			Book: book,
		})
		if err != nil {
			return application.QueryResponse{}, err
		}
		// formalize response
		return application.QueryResponse{
			Response: resp,
		}, nil
	}
}
