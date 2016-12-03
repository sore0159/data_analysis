package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func MakeStream(h *Handler) (*twitter.Stream, error) {
	config := oauth1.NewConfig(CONSUMER_KEY, CONSUMER_SECRET)
	token := oauth1.NewToken(ACCESS_TOKEN, ACCESS_SECRET)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	demux := twitter.NewSwitchDemux()
	h.HandleDemux(&demux)

	filterParams := &twitter.StreamFilterParams{
		Language: []string{"en"},
		//Track:         []string{"cat"},
		Locations:     []string{"-124.85,24.39,-66.88,49.38"},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		return nil, err
	}

	go demux.HandleChan(stream.Messages)
	return stream, nil
}
