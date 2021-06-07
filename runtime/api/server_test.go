package api

import (
	"bytes"
	"testing"

	"github.com/fdymylja/tmos/core/meta"
	v1 "github.com/fdymylja/tmos/testdata/testmodule/v1"
)

func Test_writers(t *testing.T) {
	p := &v1.Post{
		Id:      "1",
		Creator: "me",
		Title:   "you",
		Text:    "hello",
	}

	p2 := &v1.Post{
		Id:      "2",
		Creator: "you",
		Title:   "hello",
		Text:    "no",
	}
	b := new(bytes.Buffer)
	writeObjectList(b, []meta.StateObject{p, p2})
	// TODO proper matching which can't be done with expected outputs due to protobuf detrand on json
	t.Logf("%s", b)
}
