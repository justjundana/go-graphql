package main

import (
	"log"
	"net/http"

	_playground "github.com/99designs/gqlgen/graphql/playground"
	_config "github.com/justjundana/go-graphql/config"
	_graph "github.com/justjundana/go-graphql/graph"
	_generated "github.com/justjundana/go-graphql/graph/generated"
	_bookRepository "github.com/justjundana/go-graphql/repository/book"
	_userRepository "github.com/justjundana/go-graphql/repository/user"

	handler "github.com/99designs/gqlgen/graphql/handler"
)

func main() {
	db := _config.FetchConnection()

	userRepo := _userRepository.New(db)
	bookRepo := _bookRepository.New(db)

	client := _graph.NewResolver(*userRepo, *bookRepo)
	srv := handler.NewDefaultServer(_generated.NewExecutableSchema(_generated.Config{Resolvers: client}))

	http.Handle("/", _playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
