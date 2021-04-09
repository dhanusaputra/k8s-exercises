package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var (
	counter int
)

func main() {
	http.HandleFunc("/pingpong", pingpongHandler)

	log.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pingpongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "pong %d", counter)
	counter++
	writeToVolume("./shared/test", counter)
}

func writeToVolume(p string, c int) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(strconv.Itoa(c))
	if err != nil {
		log.Fatal(err)
	}
	f.Sync()
}
