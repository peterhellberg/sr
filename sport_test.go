package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSportAll(t *testing.T) {
	ts, c := testServerAndClientByFixture("sport_broadcasts")
	defer ts.Close()

	broadcasts, err := c.Sport.All()

	assert.Nil(t, err)
	assert.Equal(t, 108, len(broadcasts))

	b := broadcasts[11]

	assert.Equal(t, "Ishockey: Björklöven - AIK", b.Name)
	assert.Equal(t, "Ishockey", b.Sport.Name)
}
