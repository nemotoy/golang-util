package ping

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ping(t *testing.T) {
	want := "pong"
	ts := httptest.NewServer(http.HandlerFunc(ping))
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	if v, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fatal(err)
	} else if string(v) != want {
		t.Fatalf("expected %q, got %q", want, v)
	}
}
