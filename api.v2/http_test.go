package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIBaseURL(t *testing.T) {
	expectedURL := "http://api.sr.se/api/v2"
	assert.Equal(t, expectedURL, APIBaseURL)
}

func TestURL(t *testing.T) {
	expectedURL := "http://api.sr.se/api/v2/foo?id=123"
	assert.Equal(t, expectedURL, URL("foo?id=%v", 123))
}
