package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"products/graph/generated"
	"products/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, count *int) ([]*models.Product, error) {
	return models.Products(qm.Limit(*count)).All(ctx, boil.GetContextDB())
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
