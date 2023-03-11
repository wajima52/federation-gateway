package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"accounts/graph/generated"
	"accounts/graph/model"
	"accounts/models"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Accounts(ctx context.Context, count *int) ([]*models.Account, error) {
	return models.Accounts(qm.Limit(*count), qm.Load("AccountProfiles")).All(ctx, boil.GetContextDB())
}

// Author is the resolver for the author field.
func (r *reviewResolver) Author(ctx context.Context, obj *model.Review) (*models.Account, error) {
	return models.FindAccount(ctx, boil.GetContextDB(), obj.UserID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

type queryResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
