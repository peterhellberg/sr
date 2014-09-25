package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "sr.go/" + libraryVersion
)

// A Client communicates with the Sveriges Radio API.
type Client struct {
	Channels          ChannelsService
	Episodes          EpisodesService
	News              NewsService
	Playlists         PlaylistsService
	ProgramCategories ProgramCategoriesService
	Programs          ProgramsService
	ScheduledEpisodes ScheduledEpisodesService
	Sport             SportService
	Toplist           ToplistService

	// BaseURL is the base url for Sveriges Radio API.
	BaseURL *url.URL

	// User agent used for HTTP requests to Sveriges Radio API.
	UserAgent string

	// HTTP client used to communicate with the Sveriges Radio API.
	httpClient *http.Client
}

// NewClient returns a new Sveriges Radio API client.
// If httpClient is nil, http.DefaultClient is used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		cloned := *http.DefaultClient
		httpClient = &cloned
	}

	c := &Client{
		BaseURL: &url.URL{
			Scheme: "http",
			Host:   "api.sr.se",
			Path:   "/api/v2/",
		},
		UserAgent:  userAgent,
		httpClient: httpClient,
	}

	c.Channels = &channelsService{c}
	c.Episodes = &episodesService{c}
	c.News = &newsService{c}
	c.Playlists = &playlistsService{c}
	c.ProgramCategories = &programCategoriesService{c}
	c.Programs = &programsService{c}
	c.ScheduledEpisodes = &scheduledEpisodesService{c}
	c.Sport = &sportService{c}
	c.Toplist = &toplistService{c}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(s string) (*http.Request, error) {
	rel, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	if err != nil {
		return nil, fmt.Errorf("error reading response from %s %s: %s", req.Method, req.URL.RequestURI(), err)
	}

	return resp, nil
}

// OLD HTTP

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
