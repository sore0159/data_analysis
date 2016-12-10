package main

import (
	"github.com/dghubble/go-twitter/twitter"
)

// Handler creates async calls to Parser for each tweet and provides a
// dataPipe to re-sync Aggregator's proccessing of parsed tweet data
type Handler struct {
	Parser
	Aggregator
	dataPipe chan TweetData
}

// HandleDemux sets up the flow of information.  demux operates in
// sequence: HandleDemux creates goroutines to allow async parsing
// of tweets, with one goroutine monitering the parsed data
func (h *Handler) HandleDemux(demux *twitter.SwitchDemux) {
	go func() {
		for {
			h.Aggregator.AggregateData(<-h.dataPipe)
		}
	}()
	demux.Tweet = func(tweet *twitter.Tweet) {
		go func() {
			td, ok := h.Parser.ParseTweet(tweet)
			if ok {
				h.dataPipe <- td
			}
		}()
	}
}

func MakeHandler(c Config) (*Handler, error) {
	p, err := MakeParser()
	if err != nil {
		return nil, err
	}
	a, err := MakeAggregator(c)
	if err != nil {
		return nil, err
	}
	return &Handler{
		Parser:     p,
		Aggregator: a,
		dataPipe:   make(chan TweetData),
	}, nil
}

func (h *Handler) Close() {
	h.Parser.Close()
	h.Aggregator.Close()
}
