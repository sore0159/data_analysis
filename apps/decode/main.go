package main

import (
	"log"
	tw "mule/data_analysis/twitter"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: decode FILENAME")
		return
	}
	list, err := tw.FromCSVFile(os.Args[1])
	if err != nil {
		log.Println("FILE READ ERROR:", err)
		return
	}
	l := len(list)
	if l == 0 {
		log.Println("No values scanned!")
		return
	}
	log.Println(l, "values scanned.")
	lastTD := list[l-1]
	log.Printf("Last scanned TD:\n %+v\n", lastTD)
	t := lastTD.UserSinceDate
	log.Printf("Sample UserSinceDate: %s", t)
}
