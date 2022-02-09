package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *mutationResolver) CreateSchool(ctx context.Context, input model.CreateSchoolInput) (*ent.School, error) {
	repository := repositories.NewSchoolRepository(r.client)
	return repository.Create(ctx, input)
}

func (r *mutationResolver) UpdateSchool(ctx context.Context, input model.UpdateSchoolInput) (*ent.School, error) {
	repository := repositories.NewSchoolRepository(r.client)
	return repository.Update(ctx, input)
}

func (r *mutationResolver) DeleteSchool(ctx context.Context, id int) (*ent.School, error) {
	repository := repositories.NewSchoolRepository(r.client)
	return repository.Delete(ctx, id)
}

func (r *queryResolver) School(ctx context.Context, id int) (*ent.School, error) {
	repository := repositories.NewSchoolRepository(r.client)
	return repository.Find(ctx, id)
}

func (r *queryResolver) Schools(ctx context.Context) ([]*ent.School, error) {
	repository := repositories.NewSchoolRepository(r.client)
	return repository.All(ctx)
}
