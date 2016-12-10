package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gonum/matrix/mat64"
	"github.com/sajari/regression"

	"mule/data_analysis/maths"
	pl "mule/data_analysis/plot"
)

func DispCfg(c Config) {
	fmt.Printf("Using data dir %s\n", c.DataDir)
	if c.DoHeat || c.DoReg || c.DoHist || c.DoScatter {
		parts := make([]string, 0, 4)
		if c.DoReg {
			parts = append(parts, "regular expression calculations")
		}
		if c.DoHist {
			parts = append(parts, "histogram plots")
		}
		if c.DoHeat {
			parts = append(parts, "heatmap plot")
		}
		if c.DoScatter {
			parts = append(parts, "scatter plots")
		}
		fmt.Printf("Performing: %s\n", strings.Join(parts, ", "))
	} else {
		fmt.Println("No heavy operations requested")
	}
	if c.Log {
		fmt.Println("Results being logged.")
	}
}

func DispCov(w io.Writer, vs maths.Vars, mat *mat64.SymDense) {
	for _, v := range vs {
		fmt.Fprintf(w, "%s Mean: %f, STD: %f\n", v.Name, v.Mean, v.STD)
	}
	fmt.Fprintln(w, "")
	for _, v := range vs {
		var chars int
		for i, c := range v.Name {
			if i < 8 {
				chars += 1
				fmt.Fprint(w, string(c))
			}
		}
		for ; chars < 10; chars += 1 {
			fmt.Fprint(w, " ")
		}
	}
	fmt.Fprintln(w, "")
	r, c := mat.Dims()
	for i := 0; i < r; i += 1 {
		parts := make([]string, c)
		for j := 0; j < c; j += 1 {
			x := mat.At(i, j)
			if i > j {
				parts[j] = "        "
			} else {
				parts[j] = fmt.Sprintf("%+7.5f", x)
			}
		}
		fmt.Fprintf(w, "%s  %s\n", strings.Join(parts, "  "), vs[i].Name)
	}

}

func DispReg(w io.Writer, vs maths.Vars, i int, r *regression.Regression) {
	parts := make([]string, 1, len(vs))
	parts[0] = fmt.Sprintf("Regression formula:\n %s = %v", vs[i].Name, r.Coeff(0))

	var count int
	for j, v := range vs {
		if i == j {
			continue
		}
		count += 1
		parts = append(parts, fmt.Sprintf("%v * %s", r.Coeff(count), v.Name))
	}
	fmt.Fprintln(w, strings.Join(parts, " + "))
	fmt.Fprintf(w, "R2: %v\n\n", r.R2)
}

func FileName(base, ext string, vs ...*maths.Var) string {
	timeStr := time.Now().Format("060102_1504_")
	parts := make([]string, len(vs))
	for i, v := range vs {
		parts[i] = "_" + strings.Join(strings.Fields(v.Name), "")
	}
	return fmt.Sprintf("%s%s%s.%s", timeStr, base, strings.Join(parts, ""), ext)
}

func ScatterPng(c Config, vX, vY *maths.Var, cf float64) error {
	fName := c.DataDir + "img/" + FileName("scatter", "png", vX, vY)
	f, err := os.Create(fName)
	if err != nil {
		return err
	}
	return pl.MakeScatter(f, vX, vY, cf)
}

func HistPng(c Config, vX *maths.Var) error {
	fName := c.DataDir + "img/" + FileName("hist", "png", vX)
	f, err := os.Create(fName)
	if err != nil {
		return err
	}
	return pl.MakeHist(f, vX)
}

func HeatPng(c Config, vX, vY, vZ *maths.Var) error {
	fName := c.DataDir + "img/" + FileName("heat", "png", vX, vY, vZ)
	f, err := os.Create(fName)
	if err != nil {
		return err
	}
	return pl.MakeHeat(f, vX, vY, vZ)
}
