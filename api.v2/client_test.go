package api

import (
	"fmt"
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

// Old HTTP tests

func TestAPIBaseURL(t *testing.T) {
	expectedURL := "http://api.sr.se/api/v2"
	assert.Equal(t, expectedURL, APIBaseURL)
}

func TestURL(t *testing.T) {
	expectedURL := "http://api.sr.se/api/v2/foo?id=123"
	assert.Equal(t, expectedURL, URL("foo?id=%v", 123))
}
