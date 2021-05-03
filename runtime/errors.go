package runtime

import (
	"errors"
	"net/http"

	abci "github.com/tendermint/tendermint/abci/types"
	"k8s.io/klog/v2"
)

const (
	Codespace = "runtime"
)

var (
	ErrBadRequest      = errors.New("runtime: bad request")
	ErrNotFound        = errors.New("runtime: not found")
	ErrAlreadyExists   = errors.New("runtime: already exists")
	ErrConditionNotMet = errors.New("runtime: conditions were not met")
	ErrEmptyModuleName = errors.New("runtime: empty module name")
)

var (
	CodeUnknown          uint32 = 1
	CodeNotFound         uint32 = http.StatusNotFound
	CodeBadRequest       uint32 = http.StatusBadRequest
	CodeAlreadyExists    uint32 = http.StatusConflict
	CodeConditionsNotMet uint32 = http.StatusPreconditionFailed
)

func ToABCIResponse(gasWanted, gasUsed uint64, err error) abci.ResponseDeliverTx {
	var code uint32
	switch {
	case errors.Is(err, ErrBadRequest):
		code = CodeBadRequest
	case errors.Is(err, ErrNotFound):
		code = CodeNotFound
	case errors.Is(err, ErrAlreadyExists):
		code = CodeAlreadyExists
	case errors.Is(err, ErrConditionNotMet):
		code = CodeConditionsNotMet
	default:
		klog.Warningf("unregistered error of type %T: %s", err, err)
		code = CodeUnknown
	}
	return abci.ResponseDeliverTx{
		Code:      code,
		Log:       err.Error(),
		Info:      "",
		Codespace: Codespace,
	}
}
