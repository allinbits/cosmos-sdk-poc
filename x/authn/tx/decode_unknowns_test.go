package tx

import (
	"errors"
	"testing"

	testdata2 "github.com/fdymylja/tmos/x/authn/tx/testdata"
	"google.golang.org/protobuf/proto"
)

func Test_rejectUnknowns(t *testing.T) {
	// we generate some proto with unknowns
	x := &testdata2.WithoutUnknowns{
		A: "1",
		B: "2",
	}
	b, err := proto.Marshal(x)
	if err != nil {
		t.Fatal(err)
	}
	hasUnknowns := new(testdata2.HasUnknowns)
	err = unmarshalAndRejectUnknowns(b, hasUnknowns)
	if !errors.As(err, new(errUnknownField)) {
		t.Fatalf("unexpected error: %s", err)
	}
}
