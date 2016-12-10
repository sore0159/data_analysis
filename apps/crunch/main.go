package main

import (
	"fmt"
	"log"
	"os/exec"
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
	vars, err := d.ProcessTweets(list)
	if err != nil {
		log.Println("Error processing tweets: ", err)
		return
	}
	log.Printf("Tweets processed! (%d data points)\n", len(vars[0].Data))

	log.Println("Calculating covariance matrix...")
	mat := vars.Matrix()
	cov := maths.Cov(mat)
	DispCov(cfg, vars, cov)

	if cfg.DoReg {
		err = TestReg(cfg, vars)
		if err != nil {
			log.Println("Error testing regression: ", err)
		}
	} else if cfg.DoScatter {
		log.Println("Making scatterplots...")
		for i, vX := range vars {
			for j, vY := range vars {
				if i >= j {
					continue
				}
				log.Println("Plotting", vX.Name, "and", vY.Name+"...")
				err = ScatterPng(cfg, vX, vY, nil)
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

	exec.Command("say", "Your program is complete!").Start()
	fmt.Fprintf(cfg.Output, "\n(%s) Crunch complete!\n", time.Now())
}
