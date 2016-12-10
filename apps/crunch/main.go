package main

import (
	"fmt"
	"log"
	"time"

	"mule/data_analysis/maths"
)

func main() {
	log.Println("Starting up...")
	cfg := GetConfig()
	DispCfg(cfg)

	fmt.Fprintf(cfg.Output, "(%s) Starting crunch...\n", time.Now())
	d, err := GetData(cfg)
	if err != nil {
		log.Println("Data load failure: ", err)
		return
	}
	list, err := GetTweets(cfg)
	if err != nil {
		log.Println("Tweet load failure: ", err)
		return
	}
	log.Println("Data loaded!")

	log.Println("Processing tweets...")
	vars, err := d.ProcessTweets2(list)
	if err != nil {
		log.Println("Error processing tweets: ", err)
		return
	}
	log.Printf("Tweets processed! (%d data points)\n", len(vars[0].Data))

	log.Println("Normalizing...")
	vars.Normalize()

	log.Println("Calculating covariance matrix...")
	mat := vars.Matrix()
	cov := maths.Cov(mat)
	DispCov(cfg.Output, vars, cov)

	//TestReg(cfg, vars)

	if cfg.DoReg {
		log.Println("Running regression...")
		vIs := maths.CollectVars(vars[1], vars[2])
		vD := vars[0]
		coef, err := vIs.Regress(vD)
		if err != nil {
			log.Println("Regression error:", err)
		} else {
			log.Println("Regression complete!")
			DispReg(cfg.Output, vIs, vD, coef)
		}

	}

	if cfg.DoScatter {
		log.Println("Making scatterplots...")
		for i, vX := range vars {
			for j, vY := range vars {
				if i >= j {
					continue
				}
				log.Println("Plotting", vX.Name, "and", vY.Name+"...")
				err = ScatterPng(cfg, vX, vY, [2]float64{0, 0})
				if err != nil {
					log.Println(vX.Name, " ", vY.Name, " plot error: ", err)
				}
			}
		}
		log.Println("Scatterplots complete!")
	}

	if cfg.DoHist {
		log.Println("Making histograms...")
		for _, vX := range vars {
			log.Println("Plotting", vX.Name+"...")
			err = HistPng(cfg, vX)
			if err != nil {
				log.Println(vX.Name, " histogram error: ", err)
			}
		}
		log.Println("Histograms complete!")
	}

	fmt.Fprintf(cfg.Output, "\n(%s) Crunch complete!\n", time.Now())
}
