package main

import (
	"log"
	"net/http"

	todo "github.com/norbux/gqlgen/example/config"
	"github.com/norbux/gqlgen/graphql/handler"
	"github.com/norbux/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Todo", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
