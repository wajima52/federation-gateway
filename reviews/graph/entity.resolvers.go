package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"subgraph/graph/generated"
	"subgraph/graph/model"
	"subgraph/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// FindAccountByID is the resolver for the findAccountByID field.
func (r *entityResolver) FindAccountByID(ctx context.Context, id int) (*model.Account, error) {
	return &model.Account{
		ID: id,
	}, nil
}

// FindProductByID is the resolver for the findProductByID field.
func (r *entityResolver) FindProductByID(ctx context.Context, id int) (*model.Product, error) {
	return &model.Product{
		ID: id,
	}, nil
}

// FindReviewByID is the resolver for the findReviewByID field.
func (r *entityResolver) FindReviewByID(ctx context.Context, id int) (*models.Review, error) {
	return models.FindReview(ctx, boil.GetContextDB(), id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
