package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{client: client}})
}
