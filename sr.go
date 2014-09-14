package sr

import "github.com/peterhellberg/sr/api.v2"

// Channel retrieves a channel with the given id from the API
func Channel(id int) (*api.Channel, error) {
	return api.GetChannel(id)
}

// Channels retrieves all channels from the API
func Channels() ([]*api.Channel, error) {
	return api.GetChannels()
}

// Playlist retrieves the current playlist for the given channel id from the API
func Playlist(id int) (*api.Playlist, error) {
	return api.GetPlaylist(id)
}
