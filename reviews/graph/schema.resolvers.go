package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"reviews/graph/generated"
	"reviews/graph/model"
	"reviews/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Reviews is the resolver for the reviews field.
func (r *queryResolver) Reviews(ctx context.Context, count *int) ([]*models.Review, error) {
	return models.Reviews(qm.Limit(*count)).All(ctx, boil.GetContextDB())
}

// Text is the resolver for the text field.
func (r *reviewResolver) Text(ctx context.Context, obj *models.Review) (*string, error) {
	return &obj.Text.String, nil
}

// Product is the resolver for the product field.
func (r *reviewResolver) Product(ctx context.Context, obj *models.Review) (*model.Product, error) {
	return &model.Product{
		ID: obj.ProductID,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

type queryResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
