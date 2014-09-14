package api

import "encoding/json"

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
	resp, err := Get("playlists/rightnow?format=json&channelid=%v", id)
	if err != nil {
		return nil, err
	}

	var value struct {
		Copyright string    `json:"copyright"`
		Playlist  *Playlist `json:"playlist"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&value); err != nil {
		return nil, err
	}

	return value.Playlist, nil
}
