package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := GetConfig()
	log, err := MakeLogger(cfg)
	if err != nil {
		return
	}
	log("Making handler...")
	handler, err := MakeHandler(cfg)
	if err != nil {
		log("HANDLER CREATION FAILURE: ", err)
		return
	}
	defer handler.Close()
	log("Starting stream...")
	stream, err := MakeStream(handler)
	if err != nil {
		log("STEAM CREATION FAILURE: ", err)
		return
	}
	defer stream.Stop()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.Tick(time.Minute)
	for {
		select {
		case <-ch:
			log("Signal Recieved: stopping stream...")
			return
		case <-ticker:
			if stream.Messages == nil {
				log("Early stream termination detected: stopping program...")
				return
			}
		}
	}
}
