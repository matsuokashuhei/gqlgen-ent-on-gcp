package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/ent/school"
	"github.com/matsuokashuhei/landin/graph/model"
)

func (r *mutationResolver) CreateSchool(ctx context.Context, input model.CreateSchoolInput) (*ent.School, error) {
	school, err := r.client.School.Create().SetName(input.Name).Save(ctx)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (r *mutationResolver) UpdateSchool(ctx context.Context, input model.UpdateSchoolInput) (*ent.School, error) {
	school, err := r.client.School.UpdateOneID(input.ID).SetName(input.Name).Save(ctx)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (r *mutationResolver) DeleteSchool(ctx context.Context, id int) (*ent.School, error) {
	school, err := r.client.School.Query().Where(school.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	err = r.client.School.DeleteOne(school).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (r *queryResolver) School(ctx context.Context, id int) (*ent.School, error) {
	school, err := r.client.School.Query().Where(school.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return school, err
}

func (r *queryResolver) Schools(ctx context.Context) ([]*ent.School, error) {
	schools, err := r.client.School.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return schools, err
}
