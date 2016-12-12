package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"mule/data_analysis/maths"
	pl "mule/data_analysis/plot"
)

func FileName(base, ext string, vs ...*maths.Var) string {
	timeStr := time.Now().Format("060102_1504_")
	parts := make([]string, len(vs))
	for i, v := range vs {
		parts[i] = "_" + strings.Join(strings.Fields(v.Name), "")
	}
	return fmt.Sprintf("%s%s%s.%s", timeStr, base, strings.Join(parts, ""), ext)
}

func ScatterPng(c Config, vX, vY *maths.Var, lns [][2]float64) error {
	var alpha uint8
	if n := len(vX.Data); n < 100000 {
		alpha = 160
	} else if n < 500000 {
		alpha = 80
	} else if n < 1000000 {
		alpha = 10
	} else {
		alpha = 3
	}
	fName := c.DataDir + "img/" + FileName("scatter", "png", vX, vY)
	f, err := os.Create(fName)
	if err != nil {
		return err
	}
	return pl.MakeScatter(f, vX, vY, alpha, lns)
}

func HistPng(c Config, vX *maths.Var) error {
	fName := c.DataDir + "img/" + FileName("hist", "png", vX)
	f, err := os.Create(fName)
	if err != nil {
		return err
	}
	return pl.MakeHist(f, vX)
}
