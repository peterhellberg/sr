package sr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	assert.Equal(t, "http://api.sr.se/api/v2/", c.BaseURL.String())
	assert.Equal(t, "sr.go/0.0.1", c.UserAgent)
}

func TestNewRequest(t *testing.T) {
	r, err := NewClient(nil).NewRequest(fmt.Sprintf("foo?bar=%v", 123))

	assert.Nil(t, err)
	assert.Equal(t, "http://api.sr.se/api/v2/foo?bar=123", r.URL.String())
}

func testServerAndClientByFixture(fn string) (*httptest.Server, *Client) {
	body, _ := ioutil.ReadFile("_fixtures/" + fn + ".json")

	ts := testServer(body)

	c := NewClient(nil)
	c.BaseURL, _ = url.Parse(ts.URL)

	return ts, c
}

func testServer(body []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(body)
		}))
}
