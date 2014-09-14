package api

import (
	"fmt"
	"net/http"
)

const (
	// APIBaseURL is the base url for Sveriges Radio API
	APIBaseURL = "http://api.sr.se/api/v2"
)

// Get generates a URL and performs a HTTP GET
func Get(s string, a ...interface{}) (*http.Response, error) {
	return http.Get(URL(s, a...))
}

// URL generates a URL to an API endpoint
func URL(s string, a ...interface{}) string {
	return fmt.Sprintf("%s/%s", APIBaseURL, fmt.Sprintf(s, a...))
}
