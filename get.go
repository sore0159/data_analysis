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
	list, err := record.FromCSVFile(os.Args[1])
	if err != nil {
		return nil, err
	}
	if l := len(list); l == 0 {
		return nil, errors.New("No values scanned!")
	} else if l < 5 {
		return nil, errors.New("Not enough values scanned for analysis")
	}
	return list, nil
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
