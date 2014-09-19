package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProgramCategories(t *testing.T) {
	assert := assert.New(t)

	pc, err := GetProgramCategories()

	assert.Nil(err)

	assert.Equal(11, len(pc))
	assert.Equal("Vetenskap/Miljö", pc[10].Name)
}
