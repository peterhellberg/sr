package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNews(t *testing.T) {
	assert := assert.New(t)

	news, err := GetNews()

	assert.Nil(err)

	assert.Equal(33, len(news))
	assert.Equal(news[3].Name, "P3 Nyheter")
}
