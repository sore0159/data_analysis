package main

import (
	"fmt"
	"os"
	"time"
)

type TweetData struct {
	Text string
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
