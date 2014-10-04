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

	assert.Equal(t, "broadcasts?format=json&pagination=false", c.Sport.allPath())
}

func TestSportAllByTeamIDs(t *testing.T) {
	ts, c := testServerAndClientByFixture("sport_broadcasts_HV71")
	defer ts.Close()

	broadcasts, err := c.Sport.AllByTeamIDs("2")

	assert.Nil(t, err)
	assert.Equal(t, 8, len(broadcasts))

	b := broadcasts[4]

	assert.Equal(t, "Ishockey: HV71 - Växjö Lakers", b.Name)
	assert.Equal(t, "Kinnarps Arena", b.Arena.Name)

	ids := []string{"1", "2"}
	expectedPath := "broadcasts?format=json&pagination=false&teamIds=1,2"
	assert.Equal(t, expectedPath, c.Sport.allByTeamIDsPath(ids))
}
