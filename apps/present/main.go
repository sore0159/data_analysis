package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"strings"
	"syscall"
)

const (
	CONTENT_DIR    = "FILES/"
	REQUEST_PREFIX = ""
)

func main() {

	dn := make(chan byte)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Starting server on port :8000")
	go func() {
		if err := http.ListenAndServe(":8000", server{}); err != nil {
			log.Println("Listen and Serve Error:", err)
			dn <- 0
		}
	}()
	go func() {
		log.Println("Starting firefox...")
		cmd := exec.Command("/Applications/Firefox.app/Contents/MacOS/firefox", "127.0.0.1:8000/present.html")
		if err := cmd.Start(); err != nil {
			log.Println("Firefox start error: ", err)
			fmt.Println("\nPlease start a browser manually and go to http://127.0.0.1:8000/present.html")
		}
	}()
	select {
	case <-ch:
		fmt.Println("")
		log.Println("Termination signal recieved, stopping server...")
	case <-dn:
		fmt.Println("")
		log.Println("Exiting program...")
	}
}

type server struct{}

func (server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rIP := r.Header.Get("x-forwarded-for")
	if rIP == "" {
		rIP = r.RemoteAddr
	}
	if !strings.HasPrefix(rIP, "192.168.1.") && !strings.HasPrefix(rIP, "127.0.0.1") {
		log.Println("Unauth access attempt from ", rIP)
		http.Error(w, "Does not support nonlocal connections", 500)
		return
	}
	if strings.Index(r.URL.Path, "..") != -1 {
		http.Error(w, "Does not support .. paths", 500)
		return
	}
	http.ServeFile(w, r, TranslateRequest(r.URL.Path))
}
func TranslateRequest(requestPath string) (filePath string) {
	return CONTENT_DIR + strings.TrimPrefix(path.Clean(requestPath), REQUEST_PREFIX)
}
