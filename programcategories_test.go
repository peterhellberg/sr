package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgramCategoriesAll(t *testing.T) {
	c := NewClient(nil)

	pc, err := c.ProgramCategories.All()

	assert.Nil(t, err)
	assert.Equal(t, 11, len(pc))
}
