package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
)

var (
	httpClient = &http.Client{Timeout: time.Second * 10}
)

// GraphqlResponse ...
type GraphqlResponse struct {
	Data   interface{}    `json:"data"`
	Errors []GraphqlError `json:"errors,omitempty"`
}

// Todos ...
type Todos struct {
	Todos []Todo `json:"todos"`
}

// Todo ...
type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

// GraphqlError ...
type GraphqlError struct {
	Message string        `json:"message"`
	Path    []interface{} `json:"path,omitempty"`
}

func main() {
	r := chi.NewRouter()

	r.HandleFunc("/", handler)
	r.HandleFunc("/submit", submitHandler)
	r.HandleFunc("/update/{id}", updateHandler)

	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Starting server at port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusBadRequest)
		return
	}

	if err := downloadImageToVolume("https://picsum.photos/1200", "./static/test"); err != nil {
		http.Error(w, fmt.Sprintf("failed when download image, err: %v", err), http.StatusInternalServerError)
		return
	}

	_, err := readImageFromVolume("./static/test")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed when read image, err: %v", err), http.StatusInternalServerError)
		return
	}

	query := fmt.Sprintf("{\"query\":\"query listTodo {\\n    todos {\\n      id\\n      text\\n     done\\n     }\\n}\",\"variables\":{}}")

	resp, statusCode, err := reqBackend(query)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	t := template.Must(template.New("view.html").ParseFiles("view.html"))
	if err := t.Execute(w, resp.Data); err != nil {
		http.Error(w, fmt.Sprintf("failed when execute template, err: %v", err), http.StatusInternalServerError)
		return
	}
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

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todo := r.FormValue("todo")

	query := fmt.Sprintf("{\"query\":\"mutation createTodo ($todo:String!,$done:Boolean!) {\\n  createTodo(input:{text:$todo,done:$done}) {\\n    text\\n    done\\n  }\\n}\",\"variables\":{\"todo\":\"%s\",\"done\":false}}", todo)

	_, statusCode, err := reqBackend(query)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "not yet implemented", http.StatusNotImplemented)
		return
	}
}

func reqBackend(query string) (*GraphqlResponse, int, error) {
	req, err := http.NewRequest(http.MethodPost, "http://backend-svc:8080/query", bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if !isStatusCode2xx(resp.StatusCode) {
		return nil, resp.StatusCode, fmt.Errorf("failed with statusCode: %d, body: %s", resp.StatusCode, string(body))
	}

	graphqlResp := &GraphqlResponse{Data: &Todos{}}

	if err := json.Unmarshal(body, graphqlResp); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(graphqlResp.Errors) > 0 {
		return nil, http.StatusInternalServerError, fmt.Errorf("failed with graphql, err: %s", graphqlResp.Errors[0].Message)
	}

	return graphqlResp, resp.StatusCode, nil
}

func isStatusCode2xx(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
