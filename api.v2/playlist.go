package api

import (
	"encoding/json"
	"errors"
)

// PlaylistsService communicates with the playlists
// related endpoints in the Sveriges Radio API
type PlaylistsService interface {
}

// playlistsService implements PlaylistsService.
type playlistsService struct {
	client *Client
}

// Song represents a song
type Song struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Artist      string `json:"artist"`
	Composer    string `json:"composer"`
	Conductor   string `json:"conductor"`
}

// Playlist represents a playlist for a channel
type Playlist struct {
	Song     *Song    `json:"song"`
	NextSong *Song    `json:"nextsong"`
	Channel  *Channel `json:"channel"`
}

// GetPlaylist retrieves the current playlist for the given channel id
func GetPlaylist(id int) (*Playlist, error) {
	return FetchPlaylist(HTTPFetcher{}, id)
}

// FetchPlaylist retrieves the current playlist for the given channel id using a Fetcher
func FetchPlaylist(f Fetcher, id int) (*Playlist, error) {
	body, err := f.Fetch(URL("playlists/rightnow?format=json&channelid=%v", id))
	if err != nil {
		return nil, err
	}

	var value struct {
		Playlist *Playlist `json:"playlist"`
	}

	if err = json.Unmarshal(body, &value); err != nil {
		return nil, err
	}

	if value.Playlist == nil {
		return nil, errors.New("missing playlist key in JSON")
	}

	return value.Playlist, nil
}
