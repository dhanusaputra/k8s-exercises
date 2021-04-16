package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/assets/", http.FileServer(http.Dir("./shared")))

	log.Println("Starting server at port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	if err := downloadImageToVolume("https://picsum.photos/1200", "./shared/test"); err != nil {
		http.Error(w, fmt.Sprintf("failed when download image, err: %v", err), http.StatusInternalServerError)
		return
	}

	_, err := readImageFromVolume("./shared/test")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed when read image, err: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "<html><body>")
	fmt.Fprint(w, "<img src='assets/test' alt='test' style='width:300px;height:300px;'></br></br>")
	fmt.Fprint(w, "<input type='text' id='todo' name='todo'> <button type='button'>Create TODO</button></br>")
	fmt.Fprint(w, "<ul><li>TODO 1</li><li>TODO 2</li></ul>")
	fmt.Fprint(w, "</html></body>")
}

func downloadImageToVolume(url string, path string) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	f.Sync()

	return nil
}

func readImageFromVolume(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
