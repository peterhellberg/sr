package api

type ScheduledEpisode struct {
	EpisodeID    int      `json:"episodeid"`
	Title        string   `json:"title"`
	StartTimeUTC string   `json:"starttimeutc"`
	EndTimeUTC   string   `json:"endtimeutc"`
	Program      *Program `json:"program"`
	Channel      *Channel `json:"channel"`
}
