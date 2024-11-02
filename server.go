package main

import (
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rsp84/ride-with-me-backend/graph/database"
	graph "github.com/rsp84/ride-with-me-backend/graph/generated"
	"github.com/rsp84/ride-with-me-backend/graph/resolvers"
)

const defaultPort = "8080"

func init() {
	time.Local = time.UTC
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{DB: db}}))

	http.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
