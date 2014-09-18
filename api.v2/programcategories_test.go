package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProgramCategories(t *testing.T) {
	assert := assert.New(t)

	pc, err := GetProgramCategories()

	assert.Nil(err)

	assert.Equal(len(pc), 11)
	assert.Equal(pc[10].Name, "Vetenskap/Miljö")
}
