package main

import (
	"fmt"
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

	l := len(list)
	if l == 0 {
		log.Println("No values scanned!")
		return
	} else if l < 5 {
		log.Println("Not enough values scanned for analysis")
		return
	}
	log.Println("File loaded!")

	log.Println("Parsing data...")
	vars, err := d.ProcessTweets(list)
	if err != nil {
		log.Println("Error processing tweets: ", err)
		return
	}
	if len(vars) == 0 {
		log.Println("Empty vars result!")
		return
	}
	log.Printf("Data parsed! (%d data points)\n", len(vars[0].Data))

	log.Println("Normalizing...")
	vars.Normalize()
	r := vars.Regression(0)
	log.Println("Running regression...")
	r.Run()
	log.Println("Regression complete!")

	fmt.Printf("Regression formula (N:%d):\n%v\n", l, r.Formula)
	fmt.Printf("COEF: %v, %v, %v, %v, %v, %v\n", r.Coeff(0), r.Coeff(1), r.Coeff(2), r.Coeff(3), r.Coeff(4), r.Coeff(5))
	fmt.Printf("R2: %v\n", r.R2)
	//	fmt.Printf("Regression:\n%s\n", d.R)
	fmt.Println("\n")
	log.Println("Calculating matrix...")
	mat := vars.Matrix()
	cov := maths.Cov(mat)
	log.Println("Matrix calculated!")
	DispCov(vars, cov)

}
