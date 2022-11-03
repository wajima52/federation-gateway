package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"accounts/graph/generated"
	"accounts/graph/model"
	"accounts/models"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// FindAccountByID is the resolver for the findAccountByID field.
func (r *entityResolver) FindAccountByID(ctx context.Context, id int) (*models.Account, error) {
	return models.FindAccount(ctx, boil.GetContextDB(), id)
}

// FindReviewByID is the resolver for the findReviewByID field.
func (r *entityResolver) FindReviewByID(ctx context.Context, id int) (*model.Review, error) {
	return &model.Review{
		ID: id,
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
