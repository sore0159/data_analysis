package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"mule/data_analysis/twitter/record"
	"os"
	"time"
)

type TweetData struct {
	record.TweetData
}

type Parser struct{}

func MakeParser() (Parser, error) {
	return Parser{}, nil
}
func (p *Parser) Close() {
}

type Aggregator struct {
	File *os.File
}

func MakeAggregator() (Aggregator, error) {
	now := time.Now()
	fileName := fmt.Sprintf(DATA_DIR+"%d%02d%02d_tweets.txt", now.Year(), now.Month(), now.Day())
	f, err := os.Create(fileName)
	if err != nil {
		return Aggregator{}, err
	}
	return Aggregator{
		File: f,
	}, nil
}

func (a *Aggregator) Close() {
	a.File.Close()
}

func (a *Aggregator) AggregateData(td TweetData) {
	fmt.Fprint(a.File, td.ToCSV())
}
func (p *Parser) ParseTweet(tweet *twitter.Tweet) (TweetData, bool) {
	if td, ok := record.FromTweet(tweet); ok {
		return TweetData{td}, true
	}
	return TweetData{}, false
}
