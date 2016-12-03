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
	for scanner.Scan() {
		td, ok := record.FromCSV(scanner.Text())
		if !ok {
			continue
		}
		count += 1
		if count == 1 {
			log.Printf("Sample scanned TD:\n %+v\n", td)
			t, err := td.UserSinceDate()
			if err == nil {
				log.Printf("Sample UserSinceDate: %s", t)
			} else {
				log.Println("Error parsing ", td.UserSince, ":", err)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println("Scanner error!  ", err)
	}
	log.Printf("Exiting: %d successful data reads!\n", count)
}
