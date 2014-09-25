package api

import "encoding/json"

// NewsService communicates with the news
// related endpoints in the Sveriges Radio API
type NewsService interface {
}

// newsService implements NewsService.
type newsService struct {
	client *Client
}

// GetNews retrieves all news programs
func GetNews() ([]*Program, error) {
	return FetchNews(HTTPFetcher{})
}

// FetchNames retrieves all news programs using a fetcher
func FetchNews(f Fetcher) ([]*Program, error) {
	body, err := f.Fetch(URL("news?format=json"))
	if err != nil {
		return nil, err
	}

	var value struct {
		Programs []*Program `json:"programs"`
	}

	if err = json.Unmarshal(body, &value); err != nil {
		return nil, err
	}

	return value.Programs, nil
}
