package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"
	"time"

	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/models"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *mutationResolver) CreateSchool(ctx context.Context, input model.SchoolInput) (*models.School, error) {
	repository := repositories.NewSchoolRepository(r.DB)
	school := &models.School{
		Name: input.Name,
	}
	_, err := repository.Create(school)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (r *mutationResolver) UpdateSchool(ctx context.Context, id string, input model.SchoolInput) (*models.School, error) {
	repository := repositories.NewSchoolRepository(r.DB)
	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	school, err := repository.Find(_id)
	if err != nil {
		return nil, err
	}
	school.Name = input.Name
	repository.Update(school)
	return school, nil
}

func (r *mutationResolver) DeleteSchool(ctx context.Context, id string) (*models.School, error) {
	repository := repositories.NewSchoolRepository(r.DB)
	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	school, err := repository.Find(_id)
	if err != nil {
		return nil, err
	}
	x, err := repository.Delete(school)
	if err != nil {
		return nil, err
	} else {
		_ = x
		return school, nil
	}
}

func (r *queryResolver) Schools(ctx context.Context) ([]*models.School, error) {
	repository := repositories.NewSchoolRepository(r.DB)
	// var schools []*models.School, err = repository.FindAll()
	schools, err := repository.FindAll()
	if err != nil {
		return nil, err
	}
	return schools, nil
}

func (r *schoolResolver) ID(ctx context.Context, obj *models.School) (string, error) {
	return strconv.Itoa((int(obj.ID))), nil
}

func (r *schoolResolver) CreatedAt(ctx context.Context, obj *models.School) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *schoolResolver) UpdatedAt(ctx context.Context, obj *models.School) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// School returns generated.SchoolResolver implementation.
func (r *Resolver) School() generated.SchoolResolver { return &schoolResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type schoolResolver struct{ *Resolver }
