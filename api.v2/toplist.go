package api

import "encoding/json"

// Show represents a radio show
type Show struct {
	ID               int      `json:"id"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	Dateutc          string   `json:"dateutc"`
	Type             string   `json:"type"`
	Program          *Program `json:"program"`
	ImageURL         string   `json:"imageurl"`
	ImageURLTemplate string   `json:"imageurltemplate"`
	Broadcast        struct {
		AvailableStopUTC string `json:"availablestoputc"`
		Playlist         struct {
			Duration       int    `json:"duration"`
			PublishDateUTC string `json:"publishdateutc"`
			ID             int    `json:"id"`
			URL            string `json:"url"`
			StatKey        string `json:"statkey"`
		} `json:"playlist"`
		Broadcastfiles []struct {
			Duration       int    `json:"duration"`
			PublishDateUTC string `json:"publishdateutc"`
			ID             int    `json:"id"`
			URL            string `json:"url"`
			StatKey        string `json:"statkey"`
		} `json:"broadcastfiles"`
	} `json:"broadcast"`
}

// GetDayToplist returns the toplist for the last day
func GetDayToplist() ([]*Show, error) {
	return GetToplist(1)
}

// GetWeekToplist returns the toplist for the last week
func GetWeekToplist() ([]*Show, error) {
	return GetToplist(7)
}

// GetMonthToplist returns the toplist for the last month
func GetMonthToplist() ([]*Show, error) {
	return GetToplist(30)
}

// GetToplist retrieves the toplist for the given number of days
func GetToplist(days int) ([]*Show, error) {
	resp, err := Get("toplist?format=json&pagination=false&interval=%v", days)
	if err != nil {
		return nil, err
	}

	var value struct {
		Description string  `json:"description"`
		Copyright   string  `json:"copyright"`
		Shows       []*Show `json:"shows"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&value); err != nil {
		return nil, err
	}

	return value.Shows, nil
}
