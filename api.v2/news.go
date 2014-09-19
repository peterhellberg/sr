package api

import "encoding/json"

// GetNews retrieves all news programs
func GetNews() ([]*Program, error) {
	resp, err := Get("news?format=json")
	if err != nil {
		return nil, err
	}

	var value struct {
		Programs []*Program `json:"programs"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&value); err != nil {
		return nil, err
	}

	return value.Programs, nil
}
