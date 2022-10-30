package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"subgraph/graph/generated"
	"subgraph/graph/model"
)

// FindAccountByID is the resolver for the findAccountByID field.
func (r *entityResolver) FindAccountByID(ctx context.Context, id string) (*model.Account, error) {
	panic(fmt.Errorf("not implemented: FindAccountByID - findAccountByID"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
