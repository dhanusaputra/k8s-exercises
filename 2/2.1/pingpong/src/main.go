package main

import (
	"fmt"
	"log"
	"net/http"
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
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "pong %d", counter)
	counter++
}

