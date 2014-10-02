package sr

import "errors"

// NewsService communicates with the news
// related endpoints in the Sveriges Radio API
type NewsService interface {
	All() ([]*Program, error)
}

// newsService implements NewsService.
type newsService struct {
	client *Client
}

// All retrieves all programs
func (s *newsService) All() ([]*Program, error) {
	req, err := s.client.NewRequest(s.allPath())
	if err != nil {
		return nil, err
	}

	var value struct {
		Programs []*Program `json:"programs"`
	}

	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	if value.Programs == nil {
		return nil, errors.New("missing programs key in JSON")
	}

	return value.Programs, nil
}

func (s *newsService) allPath() string {
	return "news?format=json"
}
