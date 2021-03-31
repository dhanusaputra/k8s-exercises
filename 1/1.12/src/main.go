package main

import (
	"log"
	"net/http"
)

func main() {
  http.Handle("/image", http.FileServer(http.Dir("https://picsum.photos/1200")))
	log.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

