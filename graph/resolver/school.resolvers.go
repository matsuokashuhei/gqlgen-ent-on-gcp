package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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

func (r *mutationResolver) UpdateSchool(ctx context.Context, id int, input model.SchoolInput) (*models.School, error) {
	repository := repositories.NewSchoolRepository(r.DB)
	school, err := repository.Find(id)
	if err != nil {
		return nil, err
	}
	school.Name = input.Name
	repository.Update(school)
	return school, nil
}

func (r *mutationResolver) DeleteSchool(ctx context.Context, id int) (*models.School, error) {
	repository := repositories.NewSchoolRepository(r.DB)
	var school *models.School
	var err error
	school, err = repository.Find(id)
	if err != nil {
		return nil, err
	}
	school, err = repository.Delete(id)
	if err != nil {
		return nil, err
	} else {
		return school, nil
	}
}

func (r *queryResolver) School(ctx context.Context, id int) (*models.School, error) {
	repository := repositories.NewSchoolRepository(r.DB)
	school, err := repository.Find(id)
	if err != nil {
		return nil, err
	}
	return school, err
}

func (r *queryResolver) Schools(ctx context.Context) ([]*models.School, error) {
	repository := repositories.NewSchoolRepository(r.DB)
	schools, err := repository.FindAll()
	if err != nil {
		return nil, err
	}
	return schools, err
}

func (r *schoolResolver) ID(ctx context.Context, obj *models.School) (int, error) {
	return int(obj.ID), nil
}

func (r *schoolResolver) Studios(ctx context.Context, obj *models.School) ([]*models.Studio, error) {
	repository := repositories.NewStudioRepository(r.DB)
	studios, err := repository.FindAll(int(obj.ID))
	if err != nil {
		return nil, err
	}
	return studios, nil
}

func (r *schoolResolver) CreatedAt(ctx context.Context, obj *models.School) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *schoolResolver) UpdatedAt(ctx context.Context, obj *models.School) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}

// School returns generated.SchoolResolver implementation.
func (r *Resolver) School() generated.SchoolResolver { return &schoolResolver{r} }

type schoolResolver struct{ *Resolver }
