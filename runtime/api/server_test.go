package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/fdymylja/tmos/core/meta"
	"github.com/fdymylja/tmos/runtime/kv"
	"github.com/fdymylja/tmos/runtime/orm"
	"github.com/fdymylja/tmos/runtime/orm/indexes"
	"github.com/fdymylja/tmos/runtime/orm/objects"
	"github.com/fdymylja/tmos/testdata/testmodule"
	v1 "github.com/fdymylja/tmos/testdata/testmodule/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestServer(t *testing.T) {
	kv1, kv2 := kv.NewBadger(), kv.NewBadger()

	obj := objects.NewStore(kv1)
	idx := indexes.NewStore(kv2)
	store := orm.NewStore(obj, idx)
	desc := (testmodule.Module{}).Initialize(nil)
	for _, stateObject := range desc.StateObjects {
		require.NoError(t, store.RegisterObject(stateObject.StateObject, stateObject.Options))
	}

	srvBuilder := NewServer(store)
	err := srvBuilder.RegisterModuleAPI(desc) // TODO replace nil with an always failing client implementation
	require.NoError(t, err)

	createBoilerplateState(t, store)

	srv := httptest.NewServer(srvBuilder.mux)
	defer srv.Close()
	c := srv.Client()
	u, err := url.Parse(srv.URL)
	require.NoError(t, err)

	t.Run("test state object by id", func(t *testing.T) {
		u := &(*u)
		u.Path = "/testmodule.v1/post/1"
		resp, err := c.Get(u.String())
		require.NoError(t, err)
		require.Equal(t, resp.StatusCode, http.StatusOK)

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		obj := new(v1.Post)
		require.NoError(t, protojson.Unmarshal(b, obj))
		require.NotEmpty(t, obj.Creator)
		t.Logf("%s", obj.String())
	})

	t.Run("test singleton", func(t *testing.T) {
		u := &(*u)
		u.Path = "/testmodule.v1/params"
		resp, err := c.Get(u.String())
		require.NoError(t, err)
		require.Equal(t, resp.StatusCode, http.StatusOK)

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		obj := new(v1.Params)
		require.NoError(t, protojson.Unmarshal(b, obj))
		require.NotEmpty(t, obj.LastPostNumber)
		t.Logf("%s", obj.String())
	})

	t.Run("test list", func(t *testing.T) {
		u := &(*u)
		u.Path = "/testmodule.v1/posts"
		values := make(url.Values)
		values.Set(QueryParamSelectField, "creator=Johnny")
		u.RawQuery = values.Encode()
		resp, err := c.Get(u.String())
		require.NoError(t, err)
		require.Equal(t, resp.StatusCode, http.StatusOK)

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		t.Logf("%s", b)

		var items struct {
			Items []*v1.Post `json:"items"`
		}

		require.NoError(t, json.Unmarshal(b, &items)) // FIXME(fdymylja): when we implement dynamic proto array creation

		require.NotEmpty(t, items.Items)
	})

}

func createBoilerplateState(t *testing.T, store orm.Store) {
	require.NoError(t, store.Create(&v1.Post{
		Id:      "0",
		Creator: "Johnny",
		Title:   "The Beatles are bad",
		Text:    "I hate the beatles",
	}))
	require.NoError(t, store.Create(&v1.Post{
		Id:      "1",
		Creator: "Johnny",
		Title:   "The Queens are bad",
		Text:    "I hate The Queens",
	}))
	require.NoError(t, store.Create(&v1.Post{
		Id:      "2",
		Creator: "Frojdi",
		Title:   "Apex Predator",
		Text:    "dear diary I reached Apex Predator Rank again on Apex Legends, I must be built different",
	}))
	require.NoError(t, store.Create(&v1.Params{LastPostNumber: 1}))
}

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