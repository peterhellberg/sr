package api

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
	Haspod               bool             `json:"haspod"`
	ID                   int              `json:"id"`
	Name                 string           `json:"name"`
}
