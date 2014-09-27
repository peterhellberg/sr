package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaylistsGet(t *testing.T) {
	c := NewClient(nil)

	playlist, err := c.Playlists.Get(132)

	assert.Nil(t, err)
	assert.Equal(t, "P1", playlist.Channel.Name)
}
