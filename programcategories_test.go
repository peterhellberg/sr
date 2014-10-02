package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgramCategoriesAll(t *testing.T) {
	ts, c := testServerAndClientByFixture("programcategories")
	defer ts.Close()

	pc, err := c.ProgramCategories.All()

	assert.Nil(t, err)
	assert.Equal(t, 11, len(pc))
	assert.Equal(t, "Livsåskådning", pc[4].Name)
}
