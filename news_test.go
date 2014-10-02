package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewsAll(t *testing.T) {
	ts, c := testServerAndClientByFixture("news")
	defer ts.Close()

	news, err := c.News.All()

	assert.Nil(t, err)
	assert.Equal(t, 33, len(news))

	n := news[11]

	assert.Equal(t, "Nyheter P4 Göteborg", n.Name)
	assert.Equal(t, "P4 Göteborg", n.Channel.Name)
}
