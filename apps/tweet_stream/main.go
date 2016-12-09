package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const DATA_DIR = "DATA/"

func main() {
	logFile, err := os.OpenFile(DATA_DIR+"logs.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Println("Cannot create/open logfile: ", err)
		return
	}
	defer logFile.Close()
	fLog := func(args ...interface{}) {
		log.Println(args...)
		fmt.Fprintf(logFile, "%s %s\n", time.Now(), fmt.Sprint(args...))
	}
	fLog("Making handler...")
	handler, err := MakeHandler()
	if err != nil {
		fLog("HANDLER CREATION FAILURE: ", err)
		return
	}
	defer handler.Close()
	fLog("Starting stream...")
	stream, err := MakeStream(handler)
	if err != nil {
		fLog("STEAM CREATION FAILURE: ", err)
		return
	}
	defer stream.Stop()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	ticker := time.Tick(time.Minute)
	for {
		select {
		case <-ch:
			fLog("Signal Recieved: stopping stream...")
			return
		case <-ticker:
			if stream.Messages == nil {
				fLog("Early stream termination detected: stopping program...")
				return
			}
		}
	}
}
