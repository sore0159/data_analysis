package main

import (
	"fmt"
	//	"math"
	"strings"

	"mule/data_analysis/maths"
)

func TestReg(cfg Config, vars maths.Vars) error {
	vD, vIs := vars.Pop(0)
	n := float64(len(vars[0].Data))
	fmt.Fprintf(cfg.Output, "Testing %s regression over %s (N %1f)...\n",
		vD, vIs, n)
	rg, err := maths.FullRegression(vD, vIs)
	//coef, err := vIs.Regress(vD)
	if err != nil {
		return err
	}
	DispReg(cfg, vIs, vD, rg)
	if !cfg.DoReg {
		return nil
	}
	if cfg.DoHist {
		fmt.Fprintln(cfg.Output, "Plotting histogram for", rg.Residuals.Name+"...")
		if err = HistPng(cfg, rg.Residuals); err != nil {
			return err
		}

	}
	if !cfg.DoScatter {
		return nil
	}
	if err = ScatterPng(cfg, vD, rg.Residuals, nil); err != nil {
		return err
	}
	for i, v := range vIs {
		fmt.Fprintln(cfg.Output, "Plotting", v.Name, "and", vD.Name+"...")
		if err = ScatterPng(cfg, v, vD, [][2]float64{rg.LineFor(i)}); err != nil {
			return err
		}
		fmt.Fprintln(cfg.Output, "Plotting", v.Name, "and", rg.Residuals.Name+"...")
		if err = ScatterPng(cfg, v, rg.Residuals, nil); err != nil {
			return err
		}
	}
	return nil
}

func DispReg(cfg Config, vIs maths.Vars, vD *maths.Var, rg *maths.Regression) {
	parts := make([]string, len(rg.Coef))
	parts[0] = fmt.Sprintf("Regression formula:\n %s = (%f)", vD.Name, rg.Coef[0])
	for i, v := range vIs {
		parts[i+1] = fmt.Sprintf("(%f)[%s]", rg.Coef[i+1], v.Name)
	}
	fmt.Fprintln(cfg.Output, strings.Join(parts, " + "))
	fmt.Fprintf(cfg.Output, "MeanSqError: %f, MeanSquareResiduals: %f, R^2: %f\n", rg.MeanSqError(), rg.MeanSqResid(), rg.RSq())
}
