package maths

import (
	"log"
	"math"
	"testing"

	"github.com/gonum/stat"
	"github.com/gonum/stat/distuv"
)

func TestDOne(t *testing.T) {
	log.Println("TEST D ONE")
}

func TestDTwo(t *testing.T) {
	c := Conf(29.2, 7.5, 0.95, 126)
	log.Println("TEST CONF:", c)
	beef := []float64{118, 115, 125, 110, 112, 130, 117, 112,
		115, 120, 113, 118, 119, 122, 123, 126}
	sM, sSTD := stat.MeanStdDev(beef, nil)
	log.Println("Beef Mean, STD:", sM, sSTD)
	c = T_Conf(sM, sSTD, 0.95, len(beef))
	log.Println("Beef Conf:", c)
	log.Println("----------------------------")
}

func Conf(sM, pSTD, conf float64, n int) (sol [2]float64) {
	z := distuv.Normal{Mu: 0, Sigma: 1}
	zA := z.Quantile((1 - conf) * 0.5)
	w := zA * pSTD / math.Sqrt(float64(n))
	return [2]float64{sM + w, sM - w}
}

func T_Conf(sM, sSTD, conf float64, n int) (sol [2]float64) {
	t := distuv.StudentsT{Mu: 0, Sigma: 1, Nu: float64(n - 1)}
	tA := t.Quantile((1 - conf) * 0.5)
	w := tA * sSTD / math.Sqrt(float64(n))
	return [2]float64{sM + w, sM - w}
}
