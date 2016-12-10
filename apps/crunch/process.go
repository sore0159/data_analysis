package main

import (
	//"fmt"
	//"log"

	"math"
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

		dV.Data = append(dV.Data, float64(t.Followers))
		iV1.Data = append(iV1.Data, float64(t.Links))
		iV2.Data = append(iV2.Data, float64(t.TweetCount))
		iV3.Data = append(iV3.Data, float64(len(t.Words)))
		iV4.Data = append(iV4.Data, float64(dist))
		iV5.Data = append(iV5.Data, float64(ct.Pop))
	}
	return vars, nil
}

func (d *Data) ProcessTweets2(tws []tw.TweetData) (maths.Vars, error) {
	v1 := maths.NewVar("Ln TweetCount")
	v2 := maths.NewVar("Account Age in Days")
	v3 := maths.NewVar("Ln Followers")
	vars := maths.Vars([]*maths.Var{v1, v2, v3})
	for _, v := range vars {
		v.Data = make([]float64, 0, len(tws))
	}
	seen := make(map[int64]bool, len(tws))
	for _, t := range tws {
		//if t.TweetCount > 400000 || t.Followers > 1000000 {
		//if t.TweetCount > 200000 || t.Followers > 60000 {
		//if t.TweetCount > 70000 || t.Followers > 10000 {
		//if t.TweetCount > 20000 || t.Followers > 2000 {
		//continue
		//}
		if seen[t.UserID] {
			continue
		}
		seen[t.UserID] = true
		days := math.Floor(t.TweetDate.Sub(t.UserSinceDate).Hours() / 24.0)
		v1.Data = append(v1.Data, math.Log(float64(t.TweetCount+1)))
		v2.Data = append(v2.Data, days)
		v3.Data = append(v3.Data, math.Log(float64(t.Followers+1)))
	}
	return vars, nil
}
