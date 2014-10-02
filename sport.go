package sr

// SportService communicates with the sport
// related endpoints in the Sveriges Radio API
type SportService interface {
}

// sportService implements SportService.
type sportService struct {
	client *Client
}

// IDName is a struct with the fields ID and Name
type IDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SportBroadcast represents a sport broadcast
type SportBroadcast struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	LocalStartTime  string     `json:"localstarttime"`
	LocalStopTime   string     `json:"localstoptime"`
	Hometeam        *IDName    `json:"hometeam"`
	Awayteam        *IDName    `json:"awayteam"`
	League          *IDName    `json:"league"`
	Season          *IDName    `json:"season"`
	Arena           *IDName    `json:"arena"`
	Sport           *IDName    `json:"sport"`
	Publisher       *IDName    `json:"publisher"`
	Channel         *IDName    `json:"channel"`
	LiveAudio       *LiveAudio `json:"liveaudio"`
	MobileLiveAudio *LiveAudio `json:"mobileliveaudio"`
}
