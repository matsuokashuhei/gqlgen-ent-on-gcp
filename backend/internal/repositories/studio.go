package repositories

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
)

type StudioRepository struct {
	client *ent.Client
}

func (r *StudioRepository) Find(ctx context.Context, id int) (*ent.Studio, error) {
	return r.client.Studio.Get(ctx, id)
}

func (r *StudioRepository) Create(ctx context.Context, input model.CreateStudioInput) (*ent.Studio, error) {
	return r.client.Studio.Create().
		SetName(input.Name).
		SetLocation(input.Location).
		Save(ctx)
}

func (r *StudioRepository) Update(ctx context.Context, input model.UpdateStudioInput) (*ent.Studio, error) {
	return r.client.Studio.UpdateOneID(input.ID).
		SetName(*input.Name).
		SetLocation(*input.Location).
		Save(ctx)
}

func (r *StudioRepository) Delete(ctx context.Context, id int) (*ent.Studio, error) {
	studio, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := r.client.Studio.DeleteOne(studio).Exec(ctx); err != nil {
		return nil, err
	}
	return studio, nil
}

func NewStudioRepository(client *ent.Client) *StudioRepository {
	return &StudioRepository{client: client}
}
