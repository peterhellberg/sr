package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// APIBaseURL is the base url for Sveriges Radio API
	APIBaseURL = "http://api.sr.se/api/v2"
)

// URL generates a URL to an API endpoint
func URL(s string, a ...interface{}) string {
	return fmt.Sprintf("%s/%s", APIBaseURL, fmt.Sprintf(s, a...))
}

type Fetcher interface {
	Fetch(url string) ([]byte, error)
}

type HTTPFetcher struct{}

func (f HTTPFetcher) Fetch(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
