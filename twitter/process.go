package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
)

// ParseTweet is async
func (p *Parser) ParseTweet(tweet *twitter.Tweet) (TweetData, bool) {
	if tweet.RetweetedStatus != nil {
		return TweetData{}, false
	}
	return TweetData{
		Text: fmt.Sprintf("(%s) %s\n%s\n\n", tweet.User.Name, tweet.CreatedAt, tweet.Text),
	}, true

}

// AggregateData is channel-sequenced
func (a *Aggregator) AggregateData(td TweetData) {
	fmt.Fprint(a.File, fmt.Sprintf("%d %s", a.Count, td.Text))
	a.Count += 1
}
