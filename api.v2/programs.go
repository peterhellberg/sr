package api

// ProgramsService communicates with the programs
// related endpoints in the Sveriges Radio API
type ProgramsService interface {
}

// programsService implements ProgramsService.
type programsService struct {
	client *Client
}

// Program represents a Radio program
type Program struct {
	Description          string           `json:"description"`
	ProgramCategory      *ProgramCategory `json:"programcategory"`
	Payoff               string           `json:"payoff"`
	BroadcastInfo        string           `json:"broadcastinfo"`
	Email                string           `json:"email"`
	Phone                string           `json:"phone"`
	ProgramURL           string           `json:"programurl"`
	ProgramImage         string           `json:"programimage"`
	ProgramImageTemplate string           `json:"programimagetemplate"`
	SocialImage          string           `json:"socialimage"`
	SocialImageTemplate  string           `json:"socialimagetemplate"`
	Channel              *Channel         `json:"channel"`
	Archived             bool             `json:"archived"`
	HasOndemand          bool             `json:"hasondemand"`
	HasPod               bool             `json:"haspod"`
	ID                   int              `json:"id"`
	Name                 string           `json:"name"`
}
