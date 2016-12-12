package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Config struct {
	DataDir string
	Output  io.Writer

	DoReg       bool
	DoHist      bool
	DoScatter   bool
	Log         bool
	Tiny        bool
	MatrixStore bool
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
			case 's':
				c.DoScatter = true
			case 'l':
				c.Log = true
			case 't':
				c.Tiny = true
			case 'm':
				c.MatrixStore = true
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
			c.Output = io.MultiWriter(f, os.Stdout)
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

func DispCfg(c Config) {
	fmt.Printf("Using data dir %s\n", c.DataDir)
	if c.DoReg || c.DoHist || c.DoScatter || c.MatrixStore {
		parts := make([]string, 0, 2)
		if c.DoReg {
			parts = append(parts, "linear regression calculations")
		}
		if c.DoHist {
			parts = append(parts, "histogram plots")
		}
		if c.DoScatter {
			parts = append(parts, "scatter plots")
		}
		if c.MatrixStore {
			parts = append(parts, "matrix storing")
		}
		fmt.Printf("Performing: %s\n", strings.Join(parts, ", "))
	} else {
		fmt.Println("No heavy operations requested")
	}
	if c.Log {
		fmt.Println("Results being logged.")
	}
}
