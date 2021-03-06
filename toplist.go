package sr

import (
	"errors"
	"fmt"

	"github.com/peterhellberg/sr/wcf"
)

// Show represents a radio show
type Show struct {
	ID               int        `json:"id"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	Date             wcf.Time   `json:"dateutc"`
	Type             string     `json:"type"`
	Program          *Program   `json:"program"`
	ImageURL         string     `json:"imageurl"`
	ImageURLTemplate string     `json:"imageurltemplate"`
	Broadcast        *Broadcast `json:"broadcast"`
}

// Asset represents a radio asset
type Asset struct {
	Duration    int      `json:"duration"`
	PublishDate wcf.Time `json:"publishdateutc"`
	ID          int      `json:"id"`
	URL         string   `json:"url"`
	StatKey     string   `json:"statkey"`
}

// Broadcast represents a radio broadcast
type Broadcast struct {
	AvailableStop  wcf.Time `json:"availablestoputc"`
	Playlist       Asset    `json:"playlist"`
	Broadcastfiles []Asset  `json:"broadcastfiles"`
}

// ToplistService communicates with the toplist
// related endpoints in the Sveriges Radio API
type ToplistService interface {
	Get(days int) ([]*Show, error)
	GetDay() ([]*Show, error)
	GetWeek() ([]*Show, error)
	GetMonth() ([]*Show, error)
}

// toplistService implements ToplistService.
type toplistService struct {
	client *Client
}

// Get retrieves a toplist for the given number of days
func (s *toplistService) Get(days int) ([]*Show, error) {
	req, err := s.client.NewRequest(s.getPath(days))
	if err != nil {
		return nil, err
	}

	var value struct {
		Shows []*Show `json:"shows"`
	}

	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	if value.Shows == nil {
		return nil, errors.New("missing shows key in JSON")
	}

	return value.Shows, nil
}

// GetDay returns the toplist for the last day
func (s *toplistService) GetDay() ([]*Show, error) {
	return s.Get(1)
}

// GetWeek returns the toplist for the last week
func (s *toplistService) GetWeek() ([]*Show, error) {
	return s.Get(7)
}

// GetMonth returns the toplist for the last month
func (s *toplistService) GetMonth() ([]*Show, error) {
	return s.Get(30)
}

func (s *toplistService) getPath(days int) string {
	return fmt.Sprintf("toplist?format=json&pagination=false&interval=%v", days)
}
