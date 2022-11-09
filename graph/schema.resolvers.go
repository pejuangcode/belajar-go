package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
)

// User is the resolver for the user field.
func (r *mutationResolver) User(ctx context.Context) (*model.UserOps, error) {
	return &model.UserOps{}, nil
}

// User is the resolver for the user field.
func (r *mutationResolver) Product(ctx context.Context) (*model.ProductMutation, error) {
	return &model.ProductMutation{}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.UserQuery, error) {
	return &model.UserQuery{}, nil
}

// User is the resolver for the user field.
func (r *queryResolver) Product(ctx context.Context) (*model.ProductQuery, error) {
	return &model.ProductQuery{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
