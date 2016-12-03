package main

import (
	"bufio"
	"log"
	"mule/data_analysis/twitter/record"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: decode FILENAME")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Println("Error opening file: ", err)
		return
	}
	scanner := bufio.NewScanner(f)
	var count int
	var lastTD record.TweetData
	for scanner.Scan() {
		td, ok := record.FromCSV(scanner.Text())
		if !ok {
			continue
		}
		lastTD = td
		count += 1
	}
	if err := scanner.Err(); err != nil {
		log.Println("Scanner error!  ", err)
	}
	log.Printf("Last scanned TD:\n %+v\n", lastTD)
	t, err := lastTD.UserSinceDate()
	if err == nil {
		log.Printf("Sample UserSinceDate: %s", t)
	} else {
		log.Println("Error parsing ", lastTD.UserSince, ":", err)
	}

	log.Printf("Exiting: %d successful data reads!\n", count)
}
