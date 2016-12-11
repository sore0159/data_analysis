package pl

import (
	"fmt"
	"image/color"
	"io"

	"mule/data_analysis/maths"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
)

func MakeScatter(w io.Writer, vX, vY *maths.Var, alpha uint8, lns [][2]float64) error {
	pts := make(plotter.XYs, len(vX.Data))
	for i, x := range vX.Data {
		pts[i].X = x
		pts[i].Y = vY.Data[i]
	}
	p, err := plot.New()
	if err != nil {
		return err
	}
	n := len(vX.Data)
	if vX.OldSTD == 0 {
		p.X.Label.Text = vX.Name
		p.Title.Text = fmt.Sprintf("Scatterplot (N %d)", n)
	} else {
		p.X.Label.Text = fmt.Sprintf("%s (m: %.2f, std: %.2f)", vX.Name, vX.OldMean, vX.OldSTD)
		p.Title.Text = fmt.Sprintf("Normalized Data (N %d)", n)
	}
	if vX.OldSTD == 0 {
		p.Y.Label.Text = vY.Name
	} else {
		p.Y.Label.Text = fmt.Sprintf("%s (m: %.2f std: %.2f)", vY.Name, vY.OldMean, vY.OldSTD)
	}
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		return err
	}
	s.GlyphStyle.Color = color.NRGBA{R: 255, A: alpha}
	//s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(0.25)
	s.GlyphStyle.Shape = draw.SquareGlyph{}
	p.Add(s)
	for i, ln := range lns {
		f := plotter.NewFunction(func(x float64) float64 {
			return ln[0] + ln[1]*x
		})
		f.Samples = 2
		f.LineStyle.Width = vg.Points(1)
		switch i % 3 {
		case 1:
			f.LineStyle.Color = color.RGBA{G: 255, A: 255}
		case 2:
			f.LineStyle.Color = color.RGBA{R: 100, B: 100, A: 255}
		default:
			f.LineStyle.Color = color.RGBA{B: 255, A: 255}
		}
		p.Add(f)
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
