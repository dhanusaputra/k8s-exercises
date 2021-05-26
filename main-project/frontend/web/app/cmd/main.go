package main

import (
	"log"
	"net/http"

	"github.com/dhanusaputra/k8s-exercises/web/app/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.HandleFunc("/", handler.View)
	r.HandleFunc("/create", handler.CreateTodo)
	r.HandleFunc("/update/{id}", handler.UpdateTodo)

	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Starting server at port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
