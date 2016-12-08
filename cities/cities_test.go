package cities

import (
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}

func TestTwo(t *testing.T) {
	c, err := FromFile("../DATA/cities.json", 100)
	if err != nil {
		t.Fatal(err)
	}
	if len(c) != 100 {
		t.Fatal("Len C: ", len(c))
	}
	pt := [2]float64{75, 100}
	ct, d := c.Closest(pt)
	log.Println("Closest to ", pt, ":\n", ct, "(", d, ")")
}
