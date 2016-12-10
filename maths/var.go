package maths

import (
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/stat"
)

type Var struct {
	Name string
	Data []float64
	Mean float64
	STD  float64
}

func NewVar(name string) *Var {
	return &Var{Name: name}
}

func (v *Var) CalcMeanStd() {
	v.Mean, v.STD = stat.MeanStdDev(v.Data, nil)
}
func (v *Var) Normalize() {
	v.CalcMeanStd()
	for i, x := range v.Data {
		v.Data[i] = (x - v.Mean) / v.STD
	}
}

func (v *Var) Add(v2 *Var, s float64) {
	for i, x := range v.Data {
		v.Data[i] = x + s*v2.Data[i]
	}
}
func (v *Var) Transform(f func(float64) float64) {
	for i, x := range v.Data {
		v.Data[i] = f(x)
	}
}
func (v *Var) CopyLen(v2 *Var) {
	v.Data = make([]float64, len(v2.Data))
}

type Vars []*Var

func CollectVars(vs ...*Var) Vars {
	return Vars(vs)
}
func (vs Vars) Normalize() {
	for _, v := range vs {
		v.Normalize()
	}
}

// i, j is vs[j].Data[i]
func (vs Vars) Matrix() *mat64.Dense {
	if len(vs) == 0 || len(vs[0].Data) == 0 {
		return nil
	}
	c := len(vs)
	r := len(vs[0].Data)
	mat := make([]float64, r*c)
	for i := 0; i < r; i += 1 {
		for j := 0; j < c; j += 1 {
			mat[i*c+j] = vs[j].Data[i]
		}
	}
	return mat64.NewDense(r, c, mat)
}

// i, j is 1 if j == 0, else vs[j-1].Data[i]
func (vs Vars) RegressionMatrix() *mat64.Dense {
	if len(vs) == 0 || len(vs[0].Data) == 0 {
		return nil
	}
	c := len(vs) + 1
	r := len(vs[0].Data)
	mat := make([]float64, r*c)
	for i := 0; i < r; i += 1 {
		mat[i*c] = 1
		for j := 1; j < c; j += 1 {
			mat[i*c+j] = vs[j-1].Data[i]
		}
	}
	return mat64.NewDense(r, c, mat)
}
