package api

import "encoding/json"

type ProgramCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// GetProgramCategories retrieves all channels
func GetProgramCategories() ([]*ProgramCategory, error) {
	resp, err := Get("programcategories?format=json&pagination=false")
	if err != nil {
		return nil, err
	}

	var value struct {
		Copyright         string             `json:"copyright"`
		ProgramCategories []*ProgramCategory `json:"programcategories"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&value); err != nil {
		return nil, err
	}

	return value.ProgramCategories, nil
}
