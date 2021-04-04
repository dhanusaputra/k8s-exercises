package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	imageData string
)

func main() {
	http.HandleFunc("/image", imageHandler)
	log.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

  data, err := getImage("https://picsum.photos/1200")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed when get image, err: %v", err), http.StatusInternalServerError)
		return
	}

  imageData = data

	fmt.Fprint(w, imageData)
}

func getImage(url string) (string, error) {
	if len(imageData) > 0 {
		return imageData, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respData), nil
}
