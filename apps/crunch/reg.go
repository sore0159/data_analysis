package main

import (
	"log"
	"mule/data_analysis/maths"
)

func TestReg(cfg Config, vars maths.Vars) {
	log.Printf("Testing regression (N %d)...\n", len(vars[0].Data))
	count, age, follows := vars[0], vars[1], vars[2]
	_ = follows
	line := age.Regress(count)
	log.Printf("Regression line: %f + %f *X\n", line[0], line[1])
	test := func(vX, vY *maths.Var, ln [2]float64) {
		err := ScatterPng(cfg, vX, vY, ln)
		if err != nil {
			log.Printf("%s %s regress test plot error: %s\n", vX.Name, vY.Name, err)
		} else {
			log.Printf("%s %s Plot complete!\n", vX.Name, vY.Name)
		}
	}
	test(age, count, line)
	resids := age.Residuals(count, line)
	test(age, resids, [2]float64{})
}
