package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	random string
)

func main() {
	random = fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String())))

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/random", randomHandler)

	log.Println("Starting server at port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
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
	b, err := ioutil.ReadFile("./shared/test")
	if err != nil {
		log.Println(err)
		b = []byte("0")
	}
	fmt.Fprintf(w, "%s: %s\n", time.Now().Format(time.RFC3339), random)
	fmt.Fprintf(w, "Ping / Pongs: %s", string(b))
}
