package header

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/gddo/httputil/header"
)

var (
	CacheControlKey = "Cache-Control"
	handler         = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(CacheControlKey, "no-cache, no-store, must-revalidate")
		fmt.Fprintf(w, "hello")
	})
)

func TestParseHttpHeader(t *testing.T) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	c := ts.Client()
	r, err := c.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	got := header.ParseList(r.Header, CacheControlKey)
	want := []string{"no-cache", "no-store", "must-revalidate"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v; want %v", got, want)
	}
}
