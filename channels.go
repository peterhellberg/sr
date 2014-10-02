package sr

import (
	"errors"
	"fmt"
)

// LiveAudio represents live audio
type LiveAudio struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	StatKey string `json:"statkey"`
}

// Channel represent a Radio channel
type Channel struct {
	Image         string     `json:"image"`
	ImageTemplate string     `json:"imagetemplate"`
	Color         string     `json:"color"`
	SiteURL       string     `json:"siteurl"`
	LiveAudio     *LiveAudio `json:"liveaudio"`
	ScheduleURL   string     `json:"scheduleurl"`
	ChannelType   string     `json:"channeltype"`
	XMLTvID       string     `json:"xmltvid"`
	ID            int        `json:"id"`
	Name          string     `json:"name"`
}

// ChannelsService communicates with the channels
// related endpoints in the Sveriges Radio API
type ChannelsService interface {
	All() ([]*Channel, error)
	Get(id int) (*Channel, error)
}

// channelsService implements ChannelsService.
type channelsService struct {
	client *Client
}

// All retrieves all channels
func (s *channelsService) All() ([]*Channel, error) {
	req, err := s.client.NewRequest(s.allPath())
	if err != nil {
		return nil, err
	}

	var value struct {
		Channels []*Channel `json:"channels"`
	}

	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	if value.Channels == nil {
		return []*Channel{}, errors.New("missing channels key in JSON")
	}

	return value.Channels, nil
}

// Get retrieves a channel with the given id
func (s *channelsService) Get(id int) (*Channel, error) {
	req, err := s.client.NewRequest(s.getPath(id))
	if err != nil {
		return nil, err
	}

	var value struct {
		Channel *Channel `json:"channel"`
	}

	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	if value.Channel == nil {
		return nil, errors.New("missing channels key in JSON")
	}

	return value.Channel, nil
}

func (s *channelsService) allPath() string {
	return "channels?format=json&pagination=false"
}

func (s *channelsService) getPath(id int) string {
	return fmt.Sprintf("channels/%v?format=json", id)
}
