package repositories

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
)

type SchoolRepository struct {
	client *ent.Client
}

func (r *SchoolRepository) All(ctx context.Context) ([]*ent.School, error) {
	return r.client.School.Query().All(ctx)
}

func (r *SchoolRepository) Find(ctx context.Context, id int) (*ent.School, error) {
	return r.client.School.Get(ctx, id)
}

func (r *SchoolRepository) Create(ctx context.Context, input model.CreateSchoolInput) (*ent.School, error) {
	return r.client.School.Create().
		SetName(input.Name).
		Save(ctx)
}

func (r *SchoolRepository) Update(ctx context.Context, input model.UpdateSchoolInput) (*ent.School, error) {
	return r.client.School.UpdateOneID(input.ID).
		SetName(input.Name).
		Save(ctx)
}

func (r *SchoolRepository) Delete(ctx context.Context, id int) (*ent.School, error) {
	school, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := r.client.School.DeleteOne(school).Exec(ctx); err != nil {
		return nil, err
	}
	return school, nil
}

func NewSchoolRepository(client *ent.Client) *SchoolRepository {
	return &SchoolRepository{client: client}
}
