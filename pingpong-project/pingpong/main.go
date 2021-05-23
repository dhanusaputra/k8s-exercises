package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	counter int
)

func main() {
	port := os.Getenv("PINGPONG_PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/pingpong", pingpongHandler)

	log.Println("Starting server at port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func pingpongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "pong %d", counter)
	counter++
}
