package pl

import (
	"fmt"
	"io"

	"mule/data_analysis/maths"

	"github.com/gonum/plot"
	"github.com/gonum/plot/palette"
	"github.com/gonum/plot/plotter"
	//"github.com/gonum/plot/vg/draw"
)

func MakeHeat(w io.Writer, vX, vY, vZ *maths.Var) error {

	p, err := plot.New()
	if err != nil {
		return err
	}
	p.X.Label.Text = fmt.Sprintf("%s (m: %.2f, std: %.2f)", vX.Name, vX.Mean, vX.STD)
	p.Y.Label.Text = fmt.Sprintf("%s (m: %.2f std: %.2f)", vY.Name, vY.Mean, vY.STD)
	p.Title.Text = fmt.Sprintf("Normalized Data (N %d)", len(vX.Data))
	p.Add(plotter.NewGrid())

	h := plotter.NewHeatMap(NewHeater(vX, vY, vZ), palette.Heat(12, 1))
	p.Add(h)

	wr, err := p.WriterTo(375, 375, "png")
	// not in pixels!  "vg.Length" units
	// 375vg == 500px
	if err != nil {
		return err
	}
	_, err = wr.WriteTo(w)
	return err
}

// Heater satisfies plotter.GridXYZ
type Heater struct {
	Vars [3]*maths.Var
}

func (h Heater) Dims() (int, int) {
	return len(h.Vars[0].Data), len(h.Vars[0].Data)
}
func (h Heater) Z(c, r int) float64 {
	return h.Vars[2].Data[c]
}
func (h Heater) X(c int) float64 {
	return h.Vars[0].Data[c]
}
func (h Heater) Y(r int) float64 {
	return h.Vars[1].Data[r]
}

func NewHeater(vX, vY, vZ *maths.Var) Heater {
	return Heater{Vars: [3]*maths.Var{vX, vY, vZ}}
}
