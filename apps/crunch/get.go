package main

import (
	"errors"
	"os"
	"strings"

	"mule/data_analysis/cities"
	tw "mule/data_analysis/twitter"
)

func GetTweets(c Config) ([]tw.TweetData, error) {
	fName := c.DataDir + "super_data.txt"
	if c.Tiny {
		fName = c.DataDir + "little_data.txt"
	} else {
		for i, str := range os.Args {
			if i != 0 && !strings.HasPrefix(str, "-") {
				fName = str
				break
			}
		}
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

func GetData(c Config) (*Data, error) {
	cts, err := cities.FromFile(c.DataDir+"cities.json", 100)
	if err != nil {
		return nil, err
	}
	return &Data{
		C: cts,
	}, nil
}
