package api

import "encoding/json"

// ProgramCategory represents a Radio program category
type ProgramCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetProgramCategories retrieves all program categories
func GetProgramCategories() ([]*ProgramCategory, error) {
	return FetchProgramCategories(HTTPFetcher{})
}

// FetchProgramCategories retrieves all program categories using a Fetcher
func FetchProgramCategories(f Fetcher) ([]*ProgramCategory, error) {
	body, err := f.Fetch(URL("programcategories?format=json&pagination=false"))
	if err != nil {
		return nil, err
	}

	var value struct {
		ProgramCategories []*ProgramCategory `json:"programcategories"`
	}

	if err = json.Unmarshal(body, &value); err != nil {
		return nil, err
	}

	return value.ProgramCategories, nil
}
