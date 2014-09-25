package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelsAll(t *testing.T) {
	c := NewClient(nil)

	channels, _ := c.Channels.All()

	assert.Equal(t, 51, len(channels))
}

func TestChannelsGet(t *testing.T) {
	c := NewClient(nil)

	channel, err := c.Channels.Get(132)

	assert.Nil(t, err)
	assert.Equal(t, "P1", channel.Name)
}
