package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strconv"
	"subgraph/graph/generated"
	"subgraph/models"
)

// FindAccountByID is the resolver for the findAccountByID field.
func (r *entityResolver) FindAccountByID(ctx context.Context, id string) (*models.Account, error) {
	num, _ := strconv.Atoi(id)
	return models.FindAccount(ctx, boil.GetContextDB(), num)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
