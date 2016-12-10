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
	fName := c.DataDir + "img/" + FileName("scatter", "png", vX, vY)
	f, err := os.Create(fName)
	if err != nil {
		return err
	}
	return pl.MakeScatter(f, vX, vY, lns)
}

func HistPng(c Config, vX *maths.Var) error {
	fName := c.DataDir + "img/" + FileName("hist", "png", vX)
	f, err := os.Create(fName)
	if err != nil {
		return err
	}
	return pl.MakeHist(f, vX)
}
