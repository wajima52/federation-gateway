package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"subgraph/graph/generated"
	"subgraph/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// FindAccountByID is the resolver for the findAccountByID field.
func (r *entityResolver) FindAccountByID(ctx context.Context, id int) (*models.Account, error) {
	return models.FindAccount(ctx, boil.GetContextDB(), id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
