package main

import (
	"errors"
	"os"

	"mule/data_analysis/cities"
	"mule/data_analysis/twitter/record"
)

func GetTweets() ([]record.TweetData, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("Usage: data_analysis FILENAME")
	}
	return record.FromCSVFile(os.Args[1])
}

type Data struct {
	C cities.Cities
}

func GetData() (*Data, error) {
	cities, err := cities.FromFile("cities/cities.json", 100)
	if err != nil {
		return nil, err
	}
	return &Data{
		C: cities,
	}, nil
}
