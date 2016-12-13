package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"os"
	"strconv"
	"strings"
)

// R code for loading this result:
// A <- matrix(scan("matrix.dat", n = 1644590*6), 1644590, 6, byrow = TRUE)
// dat <- data.frame(A)
// y <- lm(V1 ~ . -1, data=dat)
func StoreProcessed(cfg Config, mat *mat64.Dense) error {
	f, err := os.Create(cfg.DataDir + "matrix.dat")
	if err != nil {
		return err
	}
	r, c := mat.Dims()
	parts := make([]string, c)
	for i := 0; i < r; i += 1 {
		row := mat.RawRowView(i)
		for j, x := range row {
			parts[j] = strconv.FormatFloat(x, 'f', -1, 64)
		}
		fmt.Fprintln(f, strings.Join(parts, " "))
	}
	return nil
}
