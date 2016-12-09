package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gonum/matrix/mat64"
	"github.com/sajari/regression"

	"mule/data_analysis/maths"
	pl "mule/data_analysis/plot"
)

const DATA_DIR = "DATA/"

func DispCov(vs maths.Vars, mat *mat64.SymDense) {
	for _, v := range vs {
		fmt.Printf("%s Mean: %f, STD: %f\n", v.Name, v.Mean, v.STD)
	}
	fmt.Println("")
	for _, v := range vs {
		var chars int
		for i, c := range v.Name {
			if i < 8 {
				chars += 1
				fmt.Print(string(c))
			}
		}
		for ; chars < 10; chars += 1 {
			fmt.Print(" ")
		}
	}
	fmt.Println("")
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
		fmt.Printf("%s  %s\n", strings.Join(parts, "  "), vs[i].Name)
	}

}

func DispReg(r *regression.Regression) {
	fmt.Printf("Regression formula:\n%v\n", r.Formula)
	fmt.Printf("COEF: %v, %v, %v, %v, %v, %v\n", r.Coeff(0), r.Coeff(1), r.Coeff(2), r.Coeff(3), r.Coeff(4), r.Coeff(5))
	fmt.Printf("R2: %v\n", r.R2)
	//	fmt.Printf("Regression:\n%s\n", d.R)
	fmt.Println("\n")
}

func ScatterPng(vX, vY *maths.Var, cf float64) error {
	now := time.Now()
	timeStr := now.Format("060102_1504_")
	f, err := os.Create(fmt.Sprintf("%simg/%sscatter_%s_%s.png", DATA_DIR, timeStr, vX.Name, vY.Name))
	if err != nil {
		return err
	}
	return pl.MakeScatter(f, vX, vY, cf)
}

func HistPng(vX *maths.Var) error {
	now := time.Now()
	timeStr := now.Format("060102_1504_")
	f, err := os.Create(fmt.Sprintf("%simg/%shist_%s.png", DATA_DIR, timeStr, vX.Name))
	if err != nil {
		return err
	}
	return pl.MakeHist(f, vX)
}
