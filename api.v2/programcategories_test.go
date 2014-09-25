package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchProgramCategories(t *testing.T) {
	assert := assert.New(t)

	pc, err := FetchProgramCategories(stubProgramCategoriesFetcher{t})

	assert.Nil(err)
	assert.Equal(2, len(pc))
	assert.Equal("Dokumentär", pc[1].Name)
}

type stubProgramCategoriesFetcher struct{ t *testing.T }

func (fetcher stubProgramCategoriesFetcher) Fetch(url string) ([]byte, error) {
	if url != "http://api.sr.se/api/v2/programcategories?format=json&pagination=false" {
		fetcher.t.Errorf(url)
	}

	return []byte(`{
		"programcategories":[
			{"id":2,"name":"Barn"},
			{"id":82,"name":"Dokumentär"}
		]
	}`), nil
}
