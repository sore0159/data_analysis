package pl

import (
	"fmt"
	"image/color"
	"io"

	"mule/data_analysis/maths"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func MakeHist(w io.Writer, vX *maths.Var) error {
	pts := plotter.Values(vX.Data)
	p, err := plot.New()
	if err != nil {
		return err
	}
	p.Title.Text = fmt.Sprintf("Normalized Data (N %d)", len(vX.Data))
	p.X.Label.Text = fmt.Sprintf("%s (m: %.2f, std: %.2f)", vX.Name, vX.OldMean, vX.OldSTD)
	p.Y.Label.Text = "Frequency"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewHist(pts, 16)
	if err != nil {
		return err
	}
	s.LineStyle.Width = vg.Points(1)
	s.LineStyle.Color = color.RGBA{B: 200, A: 255}
	p.Add(s)
	wr, err := p.WriterTo(375, 375, "png")
	// not in pixels!  "vg.Length" units
	// 375vg == 500px
	if err != nil {
		return err
	}
	_, err = wr.WriteTo(w)
	return err
}
