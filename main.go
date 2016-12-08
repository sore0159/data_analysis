package main

import (
	"log"

	"mule/data_analysis/maths"
)

func main() {
	log.Println("Starting up...")
	d, err := GetData()
	if err != nil {
		log.Println("Data load failure: ", err)
		return
	}
	list, err := GetTweets()
	if err != nil {
		log.Println("Tweet load failure: ", err)
		return
	}
	log.Println("Data loaded!")

	log.Println("Parsing data...")
	vars, err := d.ProcessTweets(list)
	if err != nil {
		log.Println("Error processing tweets: ", err)
		return
	}
	log.Printf("Data parsed! (%d data points)\n", len(vars[0].Data))

	log.Println("Normalizing...")
	vars.Normalize()

	/*
		r := vars.Regression(0)
		log.Println("Running regression...")
		r.Run()
		log.Println("Regression complete!")
		DispReg(r)
	*/

	log.Println("Calculating matrix...")
	mat := vars.Matrix()
	cov := maths.Cov(mat)
	log.Println("Matrix calculated!")
	DispCov(vars, cov)

	log.Println("Making scatterplots...")
	for i, vX := range vars {
		for j, vY := range vars {
			if i >= j {
				continue
			}
			log.Println("Plotting", vX.Name, "and", vY.Name+"...")
			err = ScatterPng(vX, vY, 0)
			if err != nil {
				log.Println(vX.Name, " ", vY.Name, " plot error: ", err)
				return
			}
		}
	}
	log.Println("Scatterplots complete!")
}
