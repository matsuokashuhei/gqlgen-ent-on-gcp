package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *mutationResolver) CreateMembersClass(ctx context.Context, input model.CreateMembersClassInput) (*ent.MembersClass, error) {
	repository := repositories.NewMembersClassRepository(r.client)
	return repository.Create(ctx, input)
}

func (r *mutationResolver) UpdateMembersClass(ctx context.Context, input model.UpdateMembersClassInput) (*ent.MembersClass, error) {
	repository := repositories.NewMembersClassRepository(r.client)
	return repository.Update(ctx, input)
}

func (r *mutationResolver) DeleteMembersClass(ctx context.Context, input model.DeleteMembersClassInput) (*ent.MembersClass, error) {
	repository := repositories.NewMembersClassRepository(r.client)
	return repository.Delete(ctx, input)
}
