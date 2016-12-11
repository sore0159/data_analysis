package main

import (
	"math"

	"mule/data_analysis/maths"
	tw "mule/data_analysis/twitter"
)

func (d *Data) ProcessTweets(tws []tw.TweetData) (maths.Vars, error) {
	vars := maths.CollectVars(
		maths.NewVar("LnFollowers"),
		maths.NewVar("Links"),
		maths.NewVar("Words"),
		maths.NewVar("LnTweetCount"),
		//maths.NewVar("LnDist"),
		maths.NewVar("LnPopulation"),
		maths.NewVar("Age(days)"),
	)
	for _, v := range vars {
		v.Data = make([]float64, 0, len(tws))
	}
	data := make([]float64, len(vars))
	for _, t := range tws {
		// I guess twitter coords are (Long, Lat) instead of (Lat, Long)
		ct, dist := d.C.Closest([2]float64{t.Location[1], t.Location[0]})
		_ = dist
		days := math.Floor(t.TweetDate.Sub(t.UserSinceDate).Hours() / 24.0)

		data = []float64{
			math.Log(float64(t.Followers + 1)),
			float64(t.Links),
			float64(len(t.Words)),
			math.Log(float64(t.TweetCount + 1)),
			//math.Log(dist + 1),
			math.Log(float64(ct.Pop)),
			days,
		}

		for i, x := range data {
			vars[i].Data = append(vars[i].Data, x)
		}
	}
	vars.Normalize()
	return vars, nil
}
