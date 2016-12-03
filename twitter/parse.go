package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"os"
	"time"
)

type TweetData struct {
	Text string
}

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

type Parser struct{}

func MakeParser() (Parser, error) {
	return Parser{}, nil
}
func (p *Parser) Close() {
}

type Aggregator struct {
	File  *os.File
	Count int
}

func MakeAggregator() (Aggregator, error) {
	now := time.Now()
	fileName := fmt.Sprintf(DATA_DIR+"%d%02d%02d_tweets.txt", now.Year(), now.Month(), now.Day())
	f, err := os.Create(fileName)
	if err != nil {
		return Aggregator{}, err
	}
	return Aggregator{
		File:  f,
		Count: 0,
	}, nil
}

func (a *Aggregator) Close() {
	a.File.Close()
}
