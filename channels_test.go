package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelsAll(t *testing.T) {
	ts, c := testServerAndClientByFixture("channels")
	defer ts.Close()

	channels, _ := c.Channels.All()

	assert.Equal(t, 51, len(channels))
}

func TestChannelsGet(t *testing.T) {
	ts, c := testServerAndClientByFixture("channels_132")
	defer ts.Close()

	channel, err := c.Channels.Get(132)

	assert.Nil(t, err)
	assert.Equal(t, "P1", channel.Name)
}
