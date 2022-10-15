package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"federation-gateway/backend/graph/generated"
	"federation-gateway/backend/graph/model"
	"federation-gateway/backend/models"

	"github.com/99designs/gqlgen/graphql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	todo := models.Todo{
		UserID:  input.UserID,
		Content: input.Content,
	}
	todo.Insert(ctx, boil.GetContextDB(), boil.Infer())
	return models.Todos(models.TodoWhere.ID.EQ(todo.ID)).One(ctx, boil.GetContextDB())
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, count *int) ([]*models.User, error) {
	return models.Users(qm.Load("Todos"), qm.Limit(*count)).All(ctx, boil.GetContextDB())
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, count *int) ([]*models.Todo, error) {
	return models.Todos(qm.Load("User"), qm.Limit(*count)).All(ctx, boil.GetContextDB())
}

// User is the resolver for the User field.
func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	return obj.R.User, nil
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	return obj.R.Todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}
func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}
func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}
