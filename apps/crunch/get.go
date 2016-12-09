package main

import (
	"errors"
	"os"

	"mule/data_analysis/cities"
	tw "mule/data_analysis/twitter"
)

func GetTweets() ([]tw.TweetData, error) {
	var fName string
	if len(os.Args) < 2 {
		fName = DATA_DIR + "little_data.txt"
	} else {
		fName = os.Args[1]
	}
	list, err := tw.FromCSVFile(fName)
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
	cities, err := cities.FromFile(DATA_DIR+"cities.json", 100)
	if err != nil {
		return nil, err
	}
	return &Data{
		C: cities,
	}, nil
}
