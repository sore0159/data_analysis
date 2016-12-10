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

	if cfg.DoReg {
		r := vars.Regression(2)
		log.Println("Running regression...")
		r.Run()
		log.Println("Regression complete!")
		DispReg(cfg.Output, vars, 2, r)
	}

	if cfg.DoScatter {
		log.Println("Making scatterplots...")
		for i, vX := range vars {
			for j, vY := range vars {
				if i >= j {
					continue
				}
				log.Println("Plotting", vX.Name, "and", vY.Name+"...")
				err = ScatterPng(cfg, vX, vY, 0)
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

	if cfg.DoHeat && len(vars) == 3 {
		log.Println("Making heatmaps...")
		for i, vX := range vars {
			var vY, vZ *maths.Var
			switch i {
			case 0:
				vY, vZ = vars[1], vars[2]
			case 1:
				vY, vZ = vars[2], vars[0]
			case 2:
				vY, vZ = vars[0], vars[1]
			}
			log.Println("Plotting", vX.Name+", "+vY.Name+", "+vZ.Name+"...")
			err = HeatPng(cfg, vX, vY, vZ)
			if err != nil {
				log.Println(vX.Name, " heatmap error: ", err)
			}
		}
		log.Println("Histograms complete!")
	}
	fmt.Fprintf(cfg.Output, "\n(%s) Crunch complete!\n", time.Now())
}
