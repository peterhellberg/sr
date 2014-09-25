package api

// SportService communicates with the sport
// related endpoints in the Sveriges Radio API
type SportService interface {
}

// sportService implements SportService.
type sportService struct {
	client *Client
}

// SportBroadcast represents a sport broadcast
type SportBroadcast struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	LocalStartTime string `json:"localstarttime"`
	LocalStopTime  string `json:"localstoptime"`
	Hometeam       struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"hometeam"`
	Awayteam struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"awayteam"`
	League struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"league"`
	Season struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"season"`
	Arena struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"arena"`
	Sport struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"sport"`
	Publisher struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"publisher"`
	Channel struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
	LiveAudio       *LiveAudio `json:"liveaudio"`
	MobileLiveAudio *LiveAudio `json:"mobileliveaudio"`
}
