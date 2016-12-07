package maths

import (
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}
func TestTwo(t *testing.T) {
	v1 := &Var{Name: "v1", Data: []float64{0.5, 1.0, 1.5}}
	v2 := &Var{Name: "v2", Data: []float64{1.0, 1.5, 3.5}}
	v1.Normalize()
	v2.Normalize()
	log.Println("v1, v2:", v1, v2)
	vs := Vars([]*Var{v1, v2})
	mat := vs.Matrix()
	log.Println("vs, mat:", vs, mat)
	log.Println(mat.At(0, 0), mat.At(1, 0), mat.At(2, 0))
	log.Println(mat.At(1, 0), mat.At(1, 1))
}
