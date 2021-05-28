package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/dhanusaputra/k8s-exercises/main-project/frontend/web/app/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handler.View)
	r.Post("/create", handler.CreateTodo)
	r.Post("/toggle/{id}", handler.ToggleTodo)

	fileServer(r, "/static/", http.Dir("./static"))

	log.Println("Starting server at port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		log.Fatal("fileServer does not permit any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
