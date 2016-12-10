package maths

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
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
	fmt.Println("")
	log.Println("VS0:", vs[0])
	x1, x2 := vs.Pop(0)
	log.Println("Pop Test: ", x1, " AND ", x2)
	log.Println("VS0:", vs[0])

}

func TestThree(t *testing.T) {
	rand.Seed(time.Now().Unix())
	size := 1000000

	v1 := RandVar("v1", size)
	v2 := RandVar("v2", size)
	v3 := RandVar("v3", size)
	v3.Transform(func(x float64) float64 {
		return 10 + x*.1
	})
	v3.Add(v1, .4)
	v3.Add(v2, .2)

	vs := CollectVars(v1, v2)
	b, err := vs.Regress(v3)
	if err != nil {
		log.Println("Regress Error: ", err)
	} else {
		log.Println("Regression coef: ", b)
	}
}

func RandVar(name string, l int) *Var {
	v := NewVar(name)
	v.Data = make([]float64, l)
	for i, _ := range v.Data {
		v.Data[i] = rand.Float64()
	}
	return v
}
