package api

import "encoding/json"

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
	resp, err := Get("channels/%v?format=json", id)
	if err != nil {
		return nil, err
	}

	var value struct {
		Copyright string   `json:"copyright"`
		Channel   *Channel `json:"channel"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&value); err != nil {
		return nil, err
	}

	return value.Channel, nil
}

// GetChannels retrieves all channels
func GetChannels() ([]*Channel, error) {
	resp, err := Get("channels?format=json&pagination=false")
	if err != nil {
		return nil, err
	}

	var value struct {
		Copyright string     `json:"copyright"`
		Channels  []*Channel `json:"channels"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&value); err != nil {
		return nil, err
	}

	return value.Channels, nil
}
