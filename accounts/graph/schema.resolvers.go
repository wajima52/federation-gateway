package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"accounts/graph/generated"
	"accounts/models"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Accounts(ctx context.Context, count *int) ([]*models.Account, error) {
	return models.Accounts(qm.Limit(*count)).All(ctx, boil.GetContextDB())
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
