package author

import (
	"github.com/fdymylja/cosmos-os/pkg/application"
	"github.com/fdymylja/cosmos-os/pkg/codec"
)

func queryAuthor() application.QueryFunc {
	return func(req application.QueryRequest) (resp application.QueryResponse, err error) {
		query := new(QueryAuthorRequest)
		err = codec.Unmarshal(req.Request, query)
		if err != nil {
			return application.QueryResponse{}, err
		}

		// query author
		author := new(Author)
		err = req.Store.Get([]byte(query.Name), author)
		if err != nil {
			return application.QueryResponse{}, err
		}
		// finalize resp
		queryResp := &QueryAuthorResponse{
			Author: author,
		}
		respBytes, err := codec.Marshal(queryResp)
		return application.QueryResponse{Response: respBytes}, err
	}
}
