package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dhanusaputra/k8s-exercises/main-project/backend/pkg/graph"
	"github.com/dhanusaputra/k8s-exercises/main-project/backend/pkg/graph/generated"
	"github.com/dhanusaputra/k8s-exercises/main-project/backend/pkg/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := util.InitDB()
	defer db.Close()

	nc := util.InitNats()
	defer nc.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: graph.NewResolver(db)}))

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if err := util.PingDB(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, "ok")
	})

	log.Printf("connect to :%s for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
