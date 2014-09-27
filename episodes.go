package sr

import (
	"errors"
	"fmt"
)

// Episode represent a Radio episode
type Episode struct {
	ID               int        `json:"id"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	URL              string     `json:"url"`
	Program          *Program   `json:"program"`
	PublishDateUTC   string     `json:"publishdateutc"`
	ImageURL         string     `json:"imageurl"`
	ImageURLTemplate string     `json:"imageurltemplate"`
	Broadcast        *Broadcast `json:"broadcast"`
	ListenPodfile    *Podfile   `json:"listenpodfile"`
	DownloadPodfile  *Podfile   `json:"downloadpodfile"`
}

// Podfile represent a podfile
type Podfile struct {
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	FileSizeInBytes int      `json:"filesizeinbytes"`
	Program         *Program `json:"program"`
	Duration        int      `json:"duration"`
	PublishDateUTC  string   `json:"publishdateutc"`
	ID              int      `json:"id"`
	URL             string   `json:"url"`
	StatKey         string   `json:"statkey"`
}

// EpisodesService communicates with the episodes
// related endpoints in the Sveriges Radio API
type EpisodesService interface {
	GetLatest(id int) (*Episode, error)
	ByProgramID(id int) ([]*Episode, error)
}

// episodesService implements EpisodesService.
type episodesService struct {
	client *Client
}

func (s *episodesService) GetLatest(id int) (*Episode, error) {
	path := fmt.Sprintf("episodes/getlatest?programid=%v&format=json", id)

	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, err
	}

	var value struct {
		Episode *Episode `json:"episode"`
	}

	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	if value.Episode == nil {
		return &Episode{}, errors.New("missing episode key in JSON")
	}

	return value.Episode, nil
}

// ByProgramID retrieves all episodes by program id
func (s *episodesService) ByProgramID(id int) ([]*Episode, error) {
	pathStr := "episodes/index?programid=%v&urltemplateid=3&audioquality=hi&pagination=false&format=json"
	path := fmt.Sprintf(pathStr, id)

	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, err
	}

	var value struct {
		Episodes []*Episode `json:"episodes"`
	}

	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	if value.Episodes == nil {
		return []*Episode{}, errors.New("missing episodes key in JSON")
	}

	return value.Episodes, nil
}
