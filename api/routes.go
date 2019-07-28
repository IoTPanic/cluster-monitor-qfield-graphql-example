package api

import (
	"github.com/gorilla/mux"
	"github.com/iotpanic/cluster-monitor-qfield-graphql-example/ql"
	"github.com/iotpanic/gqlgen/handler"
)

func applyRoutes(r *mux.Router) {
	r.HandleFunc("/", rootHandler)
	r.Handle("/query/playground", handler.Playground("GALIoT Systems GraphQL API", "/query"))
	r.Handle("/query", (handler.GraphQL(ql.NewExecutableSchema(ql.Config{Resolvers: &ql.Resolver{}}))))
}
