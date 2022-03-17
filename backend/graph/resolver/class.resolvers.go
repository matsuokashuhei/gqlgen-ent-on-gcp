package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *mutationResolver) CreateClass(ctx context.Context, input model.CreateClassInput) (*ent.Class, error) {
	repository := repositories.NewClassRepository(r.client)
	return repository.Create(ctx, input)
}

func (r *mutationResolver) UpdateClass(ctx context.Context, input model.UpdateClassInput) (*ent.Class, error) {
	repository := repositories.NewClassRepository(r.client)
	return repository.Update(ctx, input)
}

func (r *mutationResolver) DeleteClass(ctx context.Context, id int) (*ent.Class, error) {
	repository := repositories.NewClassRepository(r.client)
	return repository.Delete(ctx, id)
}

func (r *queryResolver) Class(ctx context.Context, id int) (*ent.Class, error) {
	repository := repositories.NewClassRepository(r.client)
	return repository.Find(ctx, id)
}
