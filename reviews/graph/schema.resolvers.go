package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"subgraph/graph/generated"
	"subgraph/graph/model"
	"subgraph/models"
)

// Content is the resolver for the content field.
func (r *reviewResolver) Content(ctx context.Context, obj *models.Review) (string, error) {
	panic(fmt.Errorf("not implemented: Content - content"))
}

// Author is the resolver for the author field.
func (r *reviewResolver) Author(ctx context.Context, obj *models.Review) (*model.Account, error) {
	panic(fmt.Errorf("not implemented: Author - author"))
}

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

type reviewResolver struct{ *Resolver }
