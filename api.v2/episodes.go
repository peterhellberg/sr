package api

// EpisodesService communicates with the episodes
// related endpoints in the Sveriges Radio API
type EpisodesService interface {
}

// episodesService implements EpisodesService.
type episodesService struct {
	client *Client
}

// Episode represent a Radio episode
type Episode struct {
	ID               int      `json:"id"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	URL              string   `json:"url"`
	Program          *Program `json:"program"`
	PublishDateUTC   string   `json:"publishdateutc"`
	ImageURL         string   `json:"imageurl"`
	ImageURLTemplate string   `json:"imageurltemplate"`
	ListenPodfile    *Podfile `json:"listenpodfile"`
	DownloadPodfile  *Podfile `json:"downloadpodfile"`
}

// Podfile represent a podfile
type Podfile struct {
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	FileSizeInBytes int      `json:"filesizeinbytes"`
	Program         *Program `json:"program"`
	Duration        int      `json:"duration"`
	PublishDateUTC  string   `json:"publishdateutc"`
	ID              int      `json:"id"`
	URL             string   `json:"url"`
	StatKey         string   `json:"statkey"`
}
