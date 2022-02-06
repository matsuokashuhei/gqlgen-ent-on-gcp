package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/ent/studio"
	"github.com/matsuokashuhei/landin/graph/model"
)

func (r *mutationResolver) CreateStudio(ctx context.Context, input model.CreateStudioInput) (*ent.Studio, error) {
	studio, err := r.client.Studio.Create().SetName(input.Name).SetLocation(input.Location).Save(ctx)
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (r *mutationResolver) UpdateStudio(ctx context.Context, input model.UpdateStudioInput) (*ent.Studio, error) {
	studio, err := r.client.Studio.Query().Where(studio.ID(input.ID)).First(ctx)
	if err != nil {
		return nil, err
	}
	update := studio.Update()
	if input.Name != nil {
		update = update.SetName(*input.Name)
	}
	if input.Location != nil {
		update = update.SetLocation(*input.Location)
	}
	studio, err = update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (r *mutationResolver) DeleteStudio(ctx context.Context, id int) (*ent.Studio, error) {
	studio, err := r.client.Studio.Query().Where(studio.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	if err := r.client.Studio.DeleteOne(studio).Exec(ctx); err != nil {
		return nil, err
	}
	return studio, nil
}

func (r *queryResolver) Studio(ctx context.Context, id int) (*ent.Studio, error) {
	studio, err := r.client.Studio.Query().Where(studio.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (r *queryResolver) Studios(ctx context.Context) ([]*ent.Studio, error) {
	studios, err := r.client.Studio.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return studios, nil
}
