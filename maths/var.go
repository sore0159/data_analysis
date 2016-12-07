package maths

import (
	"github.com/gonum/matrix/mat64"
	"github.com/gonum/stat"
	"github.com/sajari/regression"
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

func (v *Var) Normalize() {
	v.Mean, v.STD = stat.MeanStdDev(v.Data, nil)
	for i, x := range v.Data {
		v.Data[i] = (x - v.Mean) / v.STD
	}
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
	// for _, v := range vs {
	// v.Data = nil  // Memory freedom?
	// }
	return mat64.NewDense(r, c, mat)
}

func (vs Vars) Regression(depI int) *regression.Regression {
	if depI < 0 || len(vs) < 2 || len(vs) < depI+1 || len(vs[0].Data) < len(vs) {
		return nil
	}
	indV := make([]*Var, 0, len(vs)-1)
	for i, v := range vs {
		if i != depI {
			indV = append(indV, v)
		}
	}
	r := new(regression.Regression)
	r.SetObserved(vs[depI].Name)
	for i, v := range indV {
		r.SetVar(i, v.Name)
	}
	for i, dep := range vs[0].Data {
		ind := make([]float64, len(indV))
		for j, v := range indV {
			ind[j] = v.Data[i]
		}
		r.Train(regression.DataPoint(dep, ind))
	}
	return r
}
