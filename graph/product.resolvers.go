package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

// Create is the resolver for the create field.
func (r *productMutationResolver) Create(ctx context.Context, obj *model.ProductMutation, input model.NewProduct) (*model.Product, error) {
	return service.CreateProduct(ctx, input)
}

// Update is the resolver for the update field.
func (r *productMutationResolver) Update(ctx context.Context, obj *model.ProductMutation, id int, input model.UpdateProduct) (string, error) {
	return service.UpdateProduct(ctx, id, input)
}

// Delete is the resolver for the delete field.
func (r *productMutationResolver) Delete(ctx context.Context, obj *model.ProductMutation, id int) (string, error) {
	return service.DelteProduct(ctx, id)
}

// User is the resolver for the Product field.
func (r *productQueryResolver) Product(ctx context.Context, obj *model.ProductQuery, id int) (*model.Product, error) {
	return service.GetProductById(ctx, id)
}

// products is the resolver for the Product field.
func (r *productQueryResolver) Products(ctx context.Context, obj *model.ProductQuery) ([]*model.Product, error) {
	return service.GetAllProduct(ctx)
}

// ProductMutation returns generated.ProductMutationResolver implementation.
func (r *Resolver) ProductMutation() generated.ProductMutationResolver {
	return &productMutationResolver{r}
}

// ProductQuery returns generated.ProductQueryResolver implementation.
func (r *Resolver) ProductQuery() generated.ProductQueryResolver { return &productQueryResolver{r} }

type productMutationResolver struct{ *Resolver }
type productQueryResolver struct{ *Resolver }
