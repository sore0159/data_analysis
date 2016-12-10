package main

import (
	"fmt"
	"strings"

	"github.com/gonum/matrix/mat64"

	"mule/data_analysis/maths"
)

func DispCov(cfg Config, vs maths.Vars, mat *mat64.SymDense) {
	for _, v := range vs {
		fmt.Fprintf(cfg.Output, "%s Mean: %f, STD: %f\n", v.Name, v.OldMean, v.OldSTD)
	}
	fmt.Fprintln(cfg.Output, "")
	for _, v := range vs {
		var chars int
		for i, c := range v.Name {
			if i < 8 {
				chars += 1
				fmt.Fprint(cfg.Output, string(c))
			}
		}
		for ; chars < 10; chars += 1 {
			fmt.Fprint(cfg.Output, " ")
		}
	}
	fmt.Fprintln(cfg.Output, "")
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
		fmt.Fprintf(cfg.Output, "%s  %s\n", strings.Join(parts, "  "), vs[i].Name)
	}

}
