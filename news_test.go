package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewsAll(t *testing.T) {
	c := NewClient(nil)

	news, err := c.News.All()

	assert.Nil(t, err)
	assert.Equal(t, 33, len(news))
}
