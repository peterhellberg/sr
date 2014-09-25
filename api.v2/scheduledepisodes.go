package api

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
	EpisodeID    int      `json:"episodeid"`
	Title        string   `json:"title"`
	StartTimeUTC string   `json:"starttimeutc"`
	EndTimeUTC   string   `json:"endtimeutc"`
	Program      *Program `json:"program"`
	Channel      *Channel `json:"channel"`
}
