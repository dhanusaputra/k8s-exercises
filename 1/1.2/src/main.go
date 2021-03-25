package main

import (
  "fmt"
  "net/http"
  "log"
)

func main() {
  http.HandleFunc("/ping", pingHandler)

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
