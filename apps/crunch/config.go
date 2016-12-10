package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type Config struct {
	DataDir string
	Output  io.Writer

	DoReg     bool
	DoHist    bool
	DoHeat    bool
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
		var unused []string
		for _, cr := range str {
			switch cr {
			case 'r':
				c.DoReg = true
			case 'h':
				c.DoHist = true
			// Heatmap nonfunctional
			//case 't':
			//c.DoHeat = true
			case 's':
				c.DoScatter = true
			case 'l':
				c.Log = true
			case '-':
			default:
				unused = append(unused, string(cr))
			}
		}
		if len(unused) > 0 {
			log.Printf("Unknown flags: %s\n", strings.Join(unused, ", "))
		}
	}
	if c.Log {
		f, err := os.OpenFile(c.DataDir+"crunch_logs.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		if err == nil {
			c.Output = f
		} else {
			log.Println("Unable to open crunch logfile: using stdout. Err:", err)
			c.Log = false
			c.Output = os.Stdout
		}
	} else {
		c.Output = os.Stdout
	}
	return c
}
