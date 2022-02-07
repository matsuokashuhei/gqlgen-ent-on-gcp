package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *mutationResolver) CreateStudio(ctx context.Context, input model.CreateStudioInput) (*ent.Studio, error) {
	repository := repositories.NewStudioRepository(r.client)
	return repository.Create(ctx, input)
}

func (r *mutationResolver) UpdateStudio(ctx context.Context, input model.UpdateStudioInput) (*ent.Studio, error) {
	repository := repositories.NewStudioRepository(r.client)
	return repository.Update(ctx, input)
}

func (r *mutationResolver) DeleteStudio(ctx context.Context, id int) (*ent.Studio, error) {
	repository := repositories.NewStudioRepository(r.client)
	return repository.Delete(ctx, id)
}

func (r *queryResolver) Studio(ctx context.Context, id int) (*ent.Studio, error) {
	repository := repositories.NewStudioRepository(r.client)
	return repository.Find(ctx, id)
}
