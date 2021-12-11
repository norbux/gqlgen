package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/norbux/gqlgen/example/starwars"
	"github.com/norbux/gqlgen/example/starwars/generated"
	"github.com/norbux/gqlgen/graphql"
	"github.com/norbux/gqlgen/graphql/handler"
	"github.com/norbux/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(starwars.NewResolver()))
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		rc := graphql.GetFieldContext(ctx)
		fmt.Println("Entered", rc.Object, rc.Field.Name)
		res, err = next(ctx)
		fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
		return res, err
	})

	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
