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
	if c.DoReg || c.DoHist || c.DoScatter {
		parts := make([]string, 0, 3)
		if c.DoReg {
			parts = append(parts, "regular expression calculations")
		}
		if c.DoHist {
			parts = append(parts, "histogram plots")
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

func ScatterPng(c Config, vX, vY *maths.Var, cf float64) error {
	now := time.Now()
	timeStr := now.Format("060102_1504_")
	f, err := os.Create(fmt.Sprintf("%simg/%sscatter_%s_%s.png", c.DataDir, timeStr, vX.Name, vY.Name))
	if err != nil {
		return err
	}
	return pl.MakeScatter(f, vX, vY, cf)
}

func HistPng(c Config, vX *maths.Var) error {
	now := time.Now()
	timeStr := now.Format("060102_1504_")
	f, err := os.Create(fmt.Sprintf("%simg/%shist_%s.png", c.DataDir, timeStr, vX.Name))
	if err != nil {
		return err
	}
	return pl.MakeHist(f, vX)
}
