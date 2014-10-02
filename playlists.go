package sr

import (
	"errors"
	"fmt"
)

// Playlist represents a playlist for a channel
type Playlist struct {
	Song     *Song    `json:"song"`
	NextSong *Song    `json:"nextsong"`
	Channel  *Channel `json:"channel"`
}

// Song represents a song
type Song struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Artist      string `json:"artist"`
	Composer    string `json:"composer"`
	Conductor   string `json:"conductor"`
}

// PlaylistsService communicates with the playlists
// related endpoints in the Sveriges Radio API
type PlaylistsService interface {
	Get(id int) (*Playlist, error)
}

// playlistsService implements PlaylistsService.
type playlistsService struct {
	client *Client
}

// Get retrieves a playlist with the given id
func (s *playlistsService) Get(id int) (*Playlist, error) {
	req, err := s.client.NewRequest(s.getPath(id))
	if err != nil {
		return nil, err
	}

	var value struct {
		Playlist *Playlist `json:"playlist"`
	}

	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	if value.Playlist == nil {
		return nil, errors.New("missing playlist key in JSON")
	}

	return value.Playlist, nil
}

func (s *playlistsService) getPath(id int) string {
	return fmt.Sprintf("playlists/rightnow?format=json&channelid=%v", id)
}
