package author

import (
	"github.com/fdymylja/cosmos-os/pkg/application"
	"github.com/fdymylja/cosmos-os/pkg/codec"
)

func handleRegisterAuthor() application.DeliverFunc {
	return func(req application.DeliverRequest) (resp application.DeliverResponse, err error) {
		msg := new(MsgRegisterAuthor)
		err = codec.Unmarshal(req.Request, msg)
		if err != nil {
			return application.DeliverResponse{}, err
		}

		author := &Author{
			Name:        msg.Name,
			DateOfBirth: msg.DateOfBirth,
		}

		err = req.Store.Set([]byte(author.Name), author)
		if err != nil {
			return application.DeliverResponse{}, err
		}

		return application.DeliverResponse{}, nil
	}
}
