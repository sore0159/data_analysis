package main

import (
	"fmt"
	"strings"

	"mule/data_analysis/maths"
)

func TestReg(cfg Config, vars maths.Vars) error {
	vD, vIs := vars.Pop(0)
	parts := make([]string, len(vIs))
	for i, v := range vIs {
		parts[i] = v.Name
	}
	fmt.Fprintf(cfg.Output, "Testing %s regression over %s (N %d)...\n",
		vD.Name, strings.Join(parts, ", "), len(vars[0].Data))
	coef, err := vIs.Regress(vD)
	if err != nil {
		return err
	}
	DispReg(cfg, vIs, vD, coef)
	p := vIs.Predictions(coef)
	p.Add(vD, -1)
	p.Transform(func(x float64) float64 { return -1 * x })
	p.CalcMeanStd()
	p.Name = "Residuals From Prediction"

	resid2 := p.Copy()
	resid2.Transform(func(x float64) float64 { return x * x })
	old2 := vD.Copy()
	old2.Transform(func(x float64) float64 { return x * x })
	r2 := 1 - (resid2.Sum() / old2.Sum()) // old is normalized so mean = 0
	fmt.Fprintf(cfg.Output, "R2: %f\n", r2)
	if !cfg.DoScatter {
		return nil
	}
	if err = ScatterPng(cfg, vD, p, nil); err != nil {
		return err
	}
	for i, v := range vIs {
		c := coef[0]
		for j, v2 := range vIs {
			mn, _ := v2.CalcMeanStd() // should be just zero
			if i != j {
				c += coef[j+1] * mn
			}
		}
		fmt.Fprintln(cfg.Output, "Plotting", v.Name, "and", vD.Name+"...")
		if err = ScatterPng(cfg, v, vD, [][2]float64{[2]float64{c, coef[i+1]}}); err != nil {
			return err
		}
		fmt.Fprintln(cfg.Output, "Plotting", v.Name, "and", p.Name+"...")
		if err = ScatterPng(cfg, v, p, nil); err != nil {
			return err
		}
	}
	return nil
}

func DispReg(cfg Config, vIs maths.Vars, vD *maths.Var, coef []float64) {
	parts := make([]string, len(coef))
	parts[0] = fmt.Sprintf("Regression formula:\n %s = (%f)", vD.Name, coef[0])
	for i, v := range vIs {
		parts[i+1] = fmt.Sprintf("(%f)[%s]", coef[i+1], v.Name)
	}
	fmt.Fprintln(cfg.Output, strings.Join(parts, " + "))
}
