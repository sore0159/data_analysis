package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("Making handler...")
	handler, err := MakeHandler()
	if err != nil {
		log.Println("HANDLER CREATION FAILURE: ", err)
	}
	log.Println("Starting stream...")
	stream, err := MakeStream(handler)
	if err != nil {
		handler.Close()
		log.Fatal("STEAM CREATION FAILURE: ", err)
	}

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.Tick(time.Minute)
	for {
		select {
		case <-ch:
			stream.Stop()
			handler.Close()
			log.Println("Signal Recieved: stopping stream...")
			return
		case <-ticker:
			if stream.Messages == nil {
				stream.Stop()
				handler.Close()
				log.Println("Early stream termination detected: stopping program...")
				return
			}
		}
	}
}
