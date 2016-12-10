package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Config struct {
	DataDir string
}

func GetConfig() (c Config) {
	defaultDir := "DATA/"
	if _, err := os.Stat(defaultDir); err == nil {
		c.DataDir = defaultDir
	} else if _, err = os.Stat("../" + defaultDir); err == nil {
		c.DataDir = "../" + defaultDir
	}
	return
}

func MakeLogger(c Config) (func(...interface{}), error) {
	logFile, err := os.OpenFile(c.DataDir+"logs.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Println("Cannot create/open logfile: ", err)
		return nil, err
	}
	fLog := func(args ...interface{}) {
		log.Println(args...)
		fmt.Fprintf(logFile, "%s %s\n", time.Now(), fmt.Sprint(args...))
	}
	return fLog, nil
}
