package main

import (
	"fmt"
	"log"

	"mule/data_analysis/maths"
	tw "mule/data_analysis/twitter"
)

func (d *Data) ProcessTweets(tws []tw.TweetData) (maths.Vars, error) {
	dV := maths.NewVar("Followers")
	iV1 := maths.NewVar("Links")
	iV2 := maths.NewVar("TweetCount")
	iV3 := maths.NewVar("Words")
	iV4 := maths.NewVar("Dist")
	iV5 := maths.NewVar("Population")
	vars := maths.Vars([]*maths.Var{dV, iV1, iV2, iV3, iV4, iV5})
	for _, v := range vars {
		v.Data = make([]float64, 0, len(tws))
	}
	//var robots int
	ctCount := map[string]int{}
	for _, t := range tws {
		/*
			if !tw.Human(t) {
				robots += 1
				log.Printf("(%d)NOT HUMAN:  %+v\n", robots, t)
				continue
			}
		*/

		// I guess twitter coords are (Long, Lat) instead of (Lat, Long)
		ct, dist := d.C.Closest([2]float64{t.Location[1], t.Location[0]})
		fmt.Println("LOC:", t.Location, ct.Name, dist)
		ctCount[ct.Name] += 1

		dV.Data = append(dV.Data, float64(t.Followers))
		iV1.Data = append(iV1.Data, float64(t.Links))
		iV2.Data = append(iV2.Data, float64(t.TweetCount))
		iV3.Data = append(iV3.Data, float64(len(t.Words)))
		iV4.Data = append(iV4.Data, float64(dist))
		iV5.Data = append(iV5.Data, float64(ct.Pop))
	}
	log.Println("DEBUG -- CT COUNT: ", ctCount)
	return vars, nil
}
