package sr

import "github.com/peterhellberg/sr/wcf"

// ScheduledEpisodesService communicates with the scheduledepisodes
// related endpoints in the Sveriges Radio API
type ScheduledEpisodesService interface {
}

// scheduledEpisodesService implements ScheduledEpisodesService.
type scheduledEpisodesService struct {
	client *Client
}

// ScheduledEpisode represents a scheduled Radio episode
type ScheduledEpisode struct {
	EpisodeID int      `json:"episodeid"`
	Title     string   `json:"title"`
	StartTime wcf.Time `json:"starttimeutc"`
	EndTime   wcf.Time `json:"endtimeutc"`
	Program   *Program `json:"program"`
	Channel   *Channel `json:"channel"`
}
