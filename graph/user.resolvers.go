package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

// Products is the resolver for the products field.
func (r *userResolver) Products(ctx context.Context, obj *model.User) ([]*model.Product, error) {
	return service.GetProductByUserId(ctx)
}

// Create is the resolver for the create field.
func (r *userOpsResolver) Create(ctx context.Context, obj *model.UserOps, input model.NewUser) (*model.User, error) {
	return service.CreateUser(ctx, input)
}

// Update is the resolver for the update field.
func (r *userOpsResolver) Update(ctx context.Context, obj *model.UserOps, id int, input model.UpdateUser) (string, error) {
	return service.UpdateUser(ctx, id, input)
}

// Delete is the resolver for the delete field.
func (r *userOpsResolver) Delete(ctx context.Context, obj *model.UserOps, id int) (string, error) {
	return service.DelteUser(ctx, id)
}

// User is the resolver for the user field.
func (r *userQueryResolver) User(ctx context.Context, obj *model.UserQuery, id int) (*model.User, error) {
	return service.GetUserById(ctx, id)
}

// Users is the resolver for the users field.
func (r *userQueryResolver) Users(ctx context.Context, obj *model.UserQuery) ([]*model.User, error) {
	return service.GetAllUser(ctx)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// UserOps returns generated.UserOpsResolver implementation.
func (r *Resolver) UserOps() generated.UserOpsResolver { return &userOpsResolver{r} }

// UserQuery returns generated.UserQueryResolver implementation.
func (r *Resolver) UserQuery() generated.UserQueryResolver { return &userQueryResolver{r} }

type userResolver struct{ *Resolver }
type userOpsResolver struct{ *Resolver }
type userQueryResolver struct{ *Resolver }
