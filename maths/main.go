package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sajari/regression"
	"mule/data_analysis/cities"
	"mule/data_analysis/twitter/record"
)

func main() {
	log.Println("Starting up...")
	if len(os.Args) < 2 {
		log.Println("Usage: maths FILENAME")
		return
	}
	log.Println("Loading CSV File...")
	list, err := record.FromCSVFile(os.Args[1])
	if err != nil {
		log.Println("FILE READ ERROR:", err)
		return
	}
	l := len(list)
	if l == 0 {
		log.Println("No values scanned!")
		return
	} else if l < 5 {
		log.Println("Not enough values scanned for analysis")
		return
	}
	log.Println("File loaded!")

	log.Println("Parsing data...")
	d, err := NewData()
	if err != nil {
		log.Println("Data loading error: ", err)
		return
	}

	for _, td := range list {
		d.AddTweet(td)
	}
	log.Println("Data parsed!")
	log.Println("Running regression...")
	d.R.Run()
	log.Println("Regression complete!")

	fmt.Printf("Regression formula (N:%d):\n%v\n", l, d.R.Formula)
	fmt.Printf("COEF: %v, %v, %v, %v, %v, %v\n", d.R.Coeff(0), d.R.Coeff(1), d.R.Coeff(2), d.R.Coeff(3), d.R.Coeff(4), d.R.Coeff(5))
	fmt.Printf("R2: %v\n", d.R.R2)
	//	fmt.Printf("Regression:\n%s\n", d.R)
}

type Data struct {
	R *regression.Regression
	C cities.Cities
}

func (d *Data) AddTweet(td record.TweetData) {
	d.R.Train(regression.DataPoint(d.TweetDep(td), d.TweetInd(td)))
}

func (d *Data) TweetDep(td record.TweetData) float64 {
	return float64(td.Followers)
}
func (d *Data) TweetInd(td record.TweetData) []float64 {
	ct, dist := d.C.Closest(td.Location)
	return []float64{
		float64(td.Links),
		float64(td.TweetCount),
		float64(len(td.Words)),
		float64(dist),
		float64(ct.Pop),
	}
}

func NewData() (*Data, error) {
	r := new(regression.Regression)
	r.SetObserved("Number of Twitter Followers Who Will See This Tweet")
	r.SetVar(0, "How many links are in this Tweet")
	r.SetVar(1, "How much the creator has Tweeted before")
	r.SetVar(2, "How many words are in this Tweet")
	r.SetVar(3, "How far in km is this tweet to a major urban center (out of top 100)")
	r.SetVar(4, "How populous is closest major urban center (out of top 100)")
	c, err := cities.FromFile("../cities/cities.json", 100)
	if err != nil {
		return nil, err
	}
	return &Data{
		R: r,
		C: c,
	}, nil
}
