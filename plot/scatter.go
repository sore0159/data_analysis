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

func MakeScatter(w io.Writer, vX, vY *maths.Var, ln [2]float64) error {
	pts := make(plotter.XYs, len(vX.Data))
	for i, x := range vX.Data {
		pts[i].X = x
		pts[i].Y = vY.Data[i]
	}
	p, err := plot.New()
	if err != nil {
		return err
	}
	p.X.Label.Text = fmt.Sprintf("%s (m: %.2f, std: %.2f)", vX.Name, vX.Mean, vX.STD)
	p.Y.Label.Text = fmt.Sprintf("%s (m: %.2f std: %.2f)", vY.Name, vY.Mean, vY.STD)
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(1)
	p.Add(s)
	if ln[1] != 0 {
		f := plotter.NewFunction(func(x float64) float64 {
			return ln[0] + ln[1]*x
		})
		f.Samples = 2
		f.LineStyle.Width = vg.Points(1)
		f.LineStyle.Color = color.RGBA{B: 255, A: 255}
		p.Add(f)
		p.Title.Text = fmt.Sprintf("Normalized Data (N %d) With Regression (B %3.3f)", len(vX.Data), ln[1])
	} else {
		p.Title.Text = fmt.Sprintf("Normalized Data (N %d)", len(vX.Data))
	}
	wr, err := p.WriterTo(375, 375, "png")
	// not in pixels!  "vg.Length" units
	// 375vg == 500px
	if err != nil {
		return err
	}
	_, err = wr.WriteTo(w)
	return err
}
