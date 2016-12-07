package pl

import (
	"fmt"
	"image/color"
	"io"

	"mule/data_analysis/maths"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	//"github.com/gonum/plot/vg/draw"
)

func MakeScatter(w io.Writer, vX, vY *maths.Var, cf float64) error { //(io.WriterTo, error) {
	pts := make(plotter.XYs, len(vX.Data))
	for i, x := range vX.Data {
		pts[i].X = x
		pts[i].Y = vY.Data[i]
	}
	p, err := plot.New()
	if err != nil {
		return err
	}
	p.Title.Text = "Normalized Points Test"
	p.X.Label.Text = fmt.Sprintf("%s (m: %.2f, std: %.2f", vX.Name, vX.Mean, vX.STD)
	p.Y.Label.Text = fmt.Sprintf("%s (m: %.2f std: %.2f", vY.Name, vY.Mean, vY.STD)
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(1)
	p.Add(s)
	f := plotter.NewFunction(func(x float64) float64 {
		return x * cf
	})
	f.Samples = 2
	f.LineStyle.Width = vg.Points(1)
	f.LineStyle.Color = color.RGBA{B: 255, A: 255}
	p.Add(f)
	//
	wr, err := p.WriterTo(500, 500, "png")
	if err != nil {
		return err
	}
	_, err = wr.WriteTo(w)
	return err
}
