package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"os"
)

type Handler struct {
	File     *os.File
	dataPipe chan TweetData
	AggregateData
}

func (h *Handler) Close() {
	h.File.Close()
}
func (h *Handler) HandleDemux(demux *twitter.SwitchDemux) {
	demux.Tweet = h.HandleTweet
}

func (h *Handler) HandleTweet(tweet *twitter.Tweet) {
	td, ok := h.ParseTweet(tweet)
	if ok {
		h.dataPipe <- td
	}
}

func MakeHandler() (*Handler, error) {
	h := &Handler{}
	f, err := os.Create("data.txt")
	if err != nil {
		return nil, err
	}
	h.File = f
	h.dataPipe = make(chan TweetData)
	h.AggregateData.Init()
	go func() {
		for {
			h.HandleParsed(<-h.dataPipe)
		}
	}()
	return h, nil
}
