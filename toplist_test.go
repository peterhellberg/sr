package sr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToplistGetDay(t *testing.T) {
	ts, c := testServerAndClientByFixture("toplist")
	defer ts.Close()

	tl, err := c.Toplist.GetDay()

	assert.Nil(t, err)
	assert.Equal(t, 10, len(tl))
	assert.Equal(t, "Luncheko 20140927 12:30", tl[6].Title)
}

func TestToplistGetWeek(t *testing.T) {
	ts, c := testServerAndClientByFixture("toplist")
	defer ts.Close()

	tl, err := c.Toplist.GetMonth()

	assert.Nil(t, err)
	assert.Equal(t, "P3 Dokumentär 20140921 18:03", tl[3].Title)
}

func TestToplistGetMonth(t *testing.T) {
	ts, c := testServerAndClientByFixture("toplist")
	defer ts.Close()

	tl, err := c.Toplist.GetMonth()

	assert.Nil(t, err)
	assert.Equal(t, "Ring så spelar vi 20140927 07:03", tl[8].Title)
}
