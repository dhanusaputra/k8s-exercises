package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	random string
)

func main() {
	port := os.Getenv("APP_PORT")
	if len(port) == 0 {
		port = "5000"
	}

	random = fmt.Sprintf("%x", md5.Sum([]byte(time.Now().String())))

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/random", randomHandler)

	log.Println("Starting server at port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "pong")
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	m := os.Getenv("MESSAGE")
	fmt.Fprintln(w, m)

	fmt.Fprintf(w, "%s: %s\n", time.Now().Format(time.RFC3339), random)
	pong, err := getPong("http://pingpong-app:8080/pingpong")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Ping / Pongs: %s", pong)
}

func getPong(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodies := strings.Fields(string(body))
	if len(bodies) != 2 {
		return "", errors.New("invalid response format")
	}
	return bodies[1], nil
}
