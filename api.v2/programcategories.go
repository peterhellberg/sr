package api

import "errors"

// ProgramCategory represents a Radio program category
type ProgramCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ProgramCategoriesService communicates with the program categories
// related endpoints in the Sveriges Radio API
type ProgramCategoriesService interface {
	All() ([]*ProgramCategory, error)
}

// programCategoriesService implements ProgramCategoriesService.
type programCategoriesService struct {
	client *Client
}

// All retrieves all program categories
func (s *programCategoriesService) All() ([]*ProgramCategory, error) {
	path := "programcategories?format=json&pagination=false"

	req, err := s.client.NewRequest(path)
	if err != nil {
		return nil, err
	}

	var value struct {
		ProgramCategories []*ProgramCategory `json:"programcategories"`
	}

	_, err = s.client.Do(req, &value)
	if err != nil {
		return nil, err
	}

	if value.ProgramCategories == nil {
		return []*ProgramCategory{}, errors.New("missing programcategories key in JSON")
	}

	return value.ProgramCategories, nil
}
