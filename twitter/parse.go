package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
)

type TweetData struct {
	Text string
}

// ParseTweet is async
func (h *Handler) ParseTweet(tweet *twitter.Tweet) (TweetData, bool) {
	if tweet.RetweetedStatus != nil {
		return TweetData{}, false
	}
	return TweetData{
		Text: fmt.Sprintf("(%s) %s\n%s\n\n", tweet.User.Name, tweet.CreatedAt, tweet.Text),
	}, true

}

// HandleParsed is channel-sequenced
func (h *Handler) HandleParsed(td TweetData) {
	fmt.Fprint(h.File, fmt.Sprintf("%d %s", h.Count, td.Text))
	h.Count += 1
}

type AggregateData struct {
	Count int
}

func (a *AggregateData) Init() {
	a.Count = 0
}
