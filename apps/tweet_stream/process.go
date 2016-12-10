package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	tw "mule/data_analysis/twitter"
	"os"
	"time"
)

type TweetData struct {
	tw.TweetData
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
	fileName := fmt.Sprintf("%s%stweets.txt", DATA_DIR, time.Now().Format("060102_1504_"))
	_, err := os.Stat(DATA_DIR)
	if os.IsNotExist(err) {
		fileName = "../" + fileName
	}
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
	if td, ok := tw.FromTweet(tweet); ok {
		return TweetData{td}, true
	}
	return TweetData{}, false
}
