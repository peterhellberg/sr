package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaylistsGet(t *testing.T) {
	ts, c := testServerAndClientByFixture("playlists_rightnow_2576")
	defer ts.Close()

	playlist, err := c.Playlists.Get(132)

	assert.Nil(t, err)
	assert.Equal(t, "Din gata", playlist.Channel.Name)
}
