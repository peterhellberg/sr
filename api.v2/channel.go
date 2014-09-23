package api

import (
	"encoding/json"
	"errors"
)

// Channel represent a Radio channel
type Channel struct {
	Image         string `json:"image"`
	Imagetemplate string `json:"imagetemplate"`
	Color         string `json:"color"`
	SiteURL       string `json:"siteurl"`
	LiveAudio     struct {
		ID      int    `json:"id"`
		URL     string `json:"url"`
		Statkey string `json:"statkey"`
	} `json:"liveaudio"`
	ScheduleURL string `json:"scheduleurl"`
	ChannelType string `json:"channeltype"`
	XMLTvID     string `json:"xmltvid"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
}

// GetChannel retrieves a channel with the given id
func GetChannel(id int) (*Channel, error) {
	return FetchChannel(HTTPFetcher{}, id)
}

// FetchChannel retrieves a channel with the given id using a Fetcher
func FetchChannel(f Fetcher, id int) (*Channel, error) {
	body, err := f.Fetch(URL("channels/%v?format=json", id))
	if err != nil {
		return nil, err
	}

	var value struct {
		Channel *Channel `json:"channel"`
	}

	if err = json.Unmarshal(body, &value); err != nil {
		return nil, err
	}

	if value.Channel == nil {
		return nil, errors.New("missing channels key in JSON")
	}

	return value.Channel, nil
}

// GetChannels retrieves all channels
func GetChannels() ([]*Channel, error) {
	return FetchChannels(HTTPFetcher{})
}

// FetchChannels retrieves all channels using a Fetcher
func FetchChannels(f Fetcher) ([]*Channel, error) {
	body, err := f.Fetch(URL("channels?format=json&pagination=false"))
	if err != nil {
		return nil, err
	}

	var value struct {
		Channels []*Channel `json:"channels"`
	}

	if err = json.Unmarshal(body, &value); err != nil {
		return nil, err
	}

	if value.Channels == nil {
		return []*Channel{}, errors.New("missing channels key in JSON")
	}

	return value.Channels, nil
}
