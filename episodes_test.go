package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEpisodesGetLatest(t *testing.T) {
	ts, c := testServerAndClientByFixture("episodes_getlatest_3718")
	defer ts.Close()

	ep, _ := c.Episodes.GetLatest(3718)

	assert.Equal(t, "PR, opium och ministerdebatt - live från Örebro", ep.Title)
	assert.Equal(t, 3240, ep.Broadcast.Playlist.Duration)

	assert.Equal(t, "2014-09-25 14:06:00 +0000 UTC", ep.PublishDate.UTC().String())
}

func TestEpisodesByProgramID(t *testing.T) {
	ts, c := testServerAndClientByFixture("episodes_3718")
	defer ts.Close()

	eps, _ := c.Episodes.ByProgramID(3718)

	assert.Equal(t, 3, len(eps))
	assert.Equal(t, "Abortlagarna, trängselskatterna och pundarboende", eps[2].Title)
}
