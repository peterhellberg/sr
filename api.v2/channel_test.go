package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchChannel(t *testing.T) {
	assert := assert.New(t)

	channel, err := FetchChannel(stubChannelFetcher{t}, 132)

	assert.Nil(err)
	assert.Equal("P3", channel.Name)
}

type stubChannelFetcher struct{ t *testing.T }

func (fetcher stubChannelFetcher) Fetch(url string) ([]byte, error) {
	if url != "http://api.sr.se/api/v2/channels/132?format=json" {
		fetcher.t.Errorf(url)
	}

	return []byte(`{
    "channel": {
      "image": "http://sverigesradio.se/sida/images/164/2186756_512_512.jpg?preset=api-default-square",
      "imagetemplate": "http://sverigesradio.se/sida/images/164/2186756_512_512.jpg",
      "color": "19a972",
      "siteurl": "http://sverigesradio.se/p3",
      "liveaudio": {
        "id": 164,
        "url": "http://sverigesradio.se/topsy/direkt/164.mp3",
        "statkey": "/app/direkt/p3[k(164)]"
      },
      "scheduleurl": "http://sverigesradio.se/api/v2/scheduledepisodes?channelid=164",
      "channeltype": "Rikskanal",
      "xmltvid": "p3.sr.se",
      "id": 164,
      "name": "P3"
    }
	}`), nil
}

func TestFetchChannels(t *testing.T) {
	assert := assert.New(t)

	channels, err := FetchChannels(stubChannelsFetcher{t})

	assert.Nil(err)

	p1 := channels[0]
	assert.Equal("P1", p1.Name)
	assert.Equal("31a1bd", p1.Color)

	p2 := channels[1]
	assert.Equal("ff5a00", p2.Color)
	assert.Equal("http://sverigesradio.se/p2", p2.SiteURL)
}

type stubChannelsFetcher struct{ t *testing.T }

func (fetcher stubChannelsFetcher) Fetch(url string) ([]byte, error) {
	if url != "http://api.sr.se/api/v2/channels?format=json&pagination=false" {
		fetcher.t.Errorf(url)
	}

	return []byte(`{
  	"channels": [
  	  {
  	    "image": "http://sverigesradio.se/sida/images/132/2186745_512_512.jpg?preset=api-default-square",
  	    "imagetemplate": "http://sverigesradio.se/sida/images/132/2186745_512_512.jpg",
  	    "color": "31a1bd",
  	    "siteurl": "http://sverigesradio.se/p1",
  	    "liveaudio": {
  	      "id": 132,
  	      "url": "http://sverigesradio.se/topsy/direkt/132.mp3",
  	      "statkey": "/app/direkt/p1[k(132)]"
  	    },
  	    "scheduleurl": "http://sverigesradio.se/api/v2/scheduledepisodes?channelid=132",
  	    "channeltype": "Rikskanal",
  	    "xmltvid": "p1.sr.se",
  	    "id": 132,
  	    "name": "P1"
  	  },
  	  {
  	    "image": "http://sverigesradio.se/sida/images/163/2186754_512_512.jpg?preset=api-default-square",
  	    "imagetemplate": "http://sverigesradio.se/sida/images/163/2186754_512_512.jpg",
  	    "color": "ff5a00",
  	    "siteurl": "http://sverigesradio.se/p2",
  	    "liveaudio": {
  	      "id": 163,
  	      "url": "http://sverigesradio.se/topsy/direkt/163.mp3",
  	      "statkey": "/app/direkt/p2[k(163)]"
  	    },
  	    "scheduleurl": "http://sverigesradio.se/api/v2/scheduledepisodes?channelid=163",
  	    "channeltype": "Rikskanal",
  	    "xmltvid": "p2.sr.se",
  	    "id": 163,
  	    "name": "P2"
  	  }
		]
	}`), nil
}
