package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchNews(t *testing.T) {
	assert := assert.New(t)

	news, err := FetchNews(stubNewsFetcher{t})

	assert.Nil(err)
	assert.Equal(2, len(news))
	assert.Equal(news[1].Name, "Vetenskapsradions nyheter")
}

type stubNewsFetcher struct{ t *testing.T }

func (fetcher stubNewsFetcher) Fetch(url string) ([]byte, error) {
	if url != "http://api.sr.se/api/v2/news?format=json" {
		fetcher.t.Errorf(url)
	}

	return []byte(`{
		"programs": [
			{
			  "description": "Lyssna på de senaste nyhetssändningarna och arkivet med Kvart-i-fem-Ekot",
			  "programcategory": {
			    "id": 68,
			    "name": "Nyheter"
			  },
			  "payoff": "",
			  "broadcastinfo": "",
			  "email": "ekot@sverigesradio.se",
			  "phone": "",
			  "programurl": "http://sverigesradio.se/sida/default.aspx?programid=4540",
			  "programimage": "http://sverigesradio.se/sida/images/83/2874631_512_512.jpg?preset=api-default-square",
			  "programimagetemplate": "http://sverigesradio.se/sida/images/83/2874631_512_512.jpg",
			  "socialimage": "http://sverigesradio.se/sida/images/83/2874631_512_512.jpg?preset=api-default-square",
			  "socialimagetemplate": "http://sverigesradio.se/sida/images/83/2874631_512_512.jpg",
			  "channel": {
			    "id": 83,
			    "name": "[No channel]"
			  },
			  "archived": false,
			  "hasondemand": true,
			  "haspod": true,
			  "id": 4540,
			  "name": "Ekot"
			},
			{
				"description": "Vetenskapsnyheter från alla tänkbara forskningsområden. Här får du som lyssnare ofta höra nyheten innan den blir uppmärksammad av övriga media. Sänds i P1.",
				"programcategory": {
					"id": 12,
					"name": "Vetenskap/Miljö"
				},
				"payoff": "",
				"broadcastinfo": "",
				"email": "vet@sverigesradio.se",
				"phone": "",
				"programurl": "http://sverigesradio.se/sida/default.aspx?programid=406",
				"programimage": "http://sverigesradio.se/sida/images/406/3282832_1400_1400.jpg?preset=api-default-square",
				"programimagetemplate": "http://sverigesradio.se/sida/images/406/3282832_1400_1400.jpg",
				"socialimage": "http://sverigesradio.se/sida/images/406/3282832_1400_1400.jpg?preset=api-default-square",
				"socialimagetemplate": "http://sverigesradio.se/sida/images/406/3282832_1400_1400.jpg",
				"channel": {
					"id": 406,
					"name": "[No channel]"
				},
				"archived": false,
				"hasondemand": true,
				"haspod": true,
				"id": 406,
				"name": "Vetenskapsradions nyheter"
			}
		]
	}`), nil
}
