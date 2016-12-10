package main

import (
	"io"
	"os"
	"strings"
)

type Config struct {
	DataDir string
	Output  io.Writer

	DoReg     bool
	DoHist    bool
	DoScatter bool
	Log       bool
}

func GetConfig() (c Config) {
	defaultDir := "DATA/"
	_, err := os.Stat(defaultDir)
	if os.IsNotExist(err) {
		_, err = os.Stat("../" + defaultDir)
		if !os.IsNotExist(err) {
			c.DataDir = "../" + defaultDir
		}
	} else {
		c.DataDir = defaultDir
	}
	for _, str := range os.Args {
		if !strings.HasPrefix(str, "-") {
			continue
		}
		for _, cr := range str {
			switch cr {
			case 'r':
				c.DoReg = true
			case 'h':
				c.DoHist = true
			case 's':
				c.DoScatter = true
			case 'l':
				c.Log = true
			}
		}
	}
	if c.Log {
		f, err := os.OpenFile(c.DataDir+"crunch_logs.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		if err == nil {
			c.Output = f
		} else {
			c.Log = false
			c.Output = os.Stdout
		}
	} else {
		c.Output = os.Stdout
	}
	return c
}
