// cities.json 2014 data credit:
// https://gist.github.com/Miserlou/c5cd8364bf9b2420bb29
package cities

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"strconv"
)

type City struct {
	Name string  `json:"city"`
	Lat  float64 `json:"latitude"`
	Long float64 `json:"longitude"`
	Rank int     `json:"-"`
	Pop  int     `json:"-"`
}
type Cities []City

// Closest returns the closest city, and the distance in km
func (cities Cities) Closest(pt [2]float64) (City, float64) {
	var best City
	var dist float64
	for i, c := range cities {
		loc := [2]float64{c.Lat, c.Long}
		if hv := Haversine(pt, loc); i == 0 || hv < dist {
			best = c
			dist = hv
		}
	}
	return best, dist
}

// FromFile expects to be called on cities.json, wherever that is located
// 1000 cities possible, probably don't want to use all 1000
func FromFile(fileName string, max int) (Cities, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var cities Cities
	err = json.Unmarshal(bytes, &cities)
	if err != nil {
		return nil, err
	}
	if len(cities) == 0 {
		return Cities([]City{}), nil
	}
	if len(cities) <= max {
		return cities, nil
	}
	c2 := make([]City, max)
	copy(c2, cities)
	return Cities(c2), nil
}

// Credit http://www.movable-type.co.uk/scripts/latlong.html
// And presumably some guy named Haversine
func Haversine(p1, p2 [2]float64) float64 {
	const R = 6371 // kilometers
	lat1, lat2 := ToRad(p1[0]), ToRad(p2[0])
	dLat := ToRad(p2[0] - p1[0])
	dLong := ToRad(p2[1] - p1[1])
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLong/2)*math.Sin(dLong/2)
	return R * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

func ToRad(deg float64) float64 {
	return (deg / 180.0) * math.Pi
}

func (c *City) UnmarshalJSON(bytes []byte) error {
	type Alias City
	type Temp struct {
		Alias
		Rank string `json:"rank"`
		Pop  string `json:"population"`
	}
	a := Temp{}
	err := json.Unmarshal(bytes, &a)
	if err != nil {
		return err
	}
	a.Alias.Rank, err = strconv.Atoi(a.Rank)
	if err != nil {
		return err
	}
	a.Alias.Pop, err = strconv.Atoi(a.Pop)
	if err != nil {
		return err
	}
	*c = City(a.Alias)
	return nil
}
