package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	random  string
	counter int
)

func main() {
	random = fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String())))

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/random", randomHandler)
	http.HandleFunc("/pingpong", pingpongHandler)

	log.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "pong")
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", time.Now().Format(time.RFC3339), random)
}

func pingpongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "pong %d", counter)
	counter++
}
