package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannel(t *testing.T) {
	assert := assert.New(t)

	p3, err := Channel(164)

	assert.Nil(err)
	assert.Equal("P3", p3.Name)
	assert.Equal("19a972", p3.Color)
}

func TestChannels(t *testing.T) {
	assert := assert.New(t)

	channels, err := Channels()

	assert.Nil(err)
	assert.Equal(51, len(channels))

	assert.Equal("P4 Blekinge", channels[3].Name)
	assert.Equal("http://sverigesradio.se/uppland/", channels[20].SiteURL)
}
