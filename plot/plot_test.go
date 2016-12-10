// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pl

import (
	"log"
	"testing"

	"mule/data_analysis/maths"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}

func TestTwo(t *testing.T) {
	v1 := maths.NewVar("v1")
	v2 := maths.NewVar("v2")
	v1.Data = []float64{0, 0, 0, 1, 4}
	v2.Data = []float64{1, 2, 3, 4, 5}
	slope := 1.0
	f, err := os.Create("test_scatter.png")
	if err != nil {
		t.Fatal(err)
	}
	log.Println("MakeScatter: ", MakeScatter(f, v1, v2, slope))
}
