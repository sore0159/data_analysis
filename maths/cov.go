package maths

import (
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/stat"
)

func Cov(mat *mat64.Dense) *mat64.SymDense {
	return stat.CovarianceMatrix(nil, mat, nil)
}
