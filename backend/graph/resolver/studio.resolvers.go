package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/models"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *mutationResolver) CreateStudio(ctx context.Context, input model.CreateStudioInput) (*models.Studio, error) {
	repository := repositories.NewStudioRepository(r.DB)
	studio := &models.Studio{
		Name:     input.Name,
		SchoolID: input.SchoolID,
	}
	_, err := repository.Create(studio)
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (r *mutationResolver) UpdateStudio(ctx context.Context, input model.UpdateStudioInput) (*models.Studio, error) {
	repository := repositories.NewStudioRepository(r.DB)
	var studio, err = repository.Find(input.ID)
	if err != nil {
		return nil, err
	}
	if input.Name != nil || input.SchoolID != nil {
		return studio, nil
	}
	if input.Name != nil {
		studio.Name = *input.Name
	}
	if input.SchoolID != nil {
		studio.SchoolID = *input.SchoolID
	}
	_, err = repository.Update(studio)
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (r *mutationResolver) DeleteStudio(ctx context.Context, id uint) (*models.Studio, error) {
	repository := repositories.NewStudioRepository(r.DB)
	studio, err := repository.Delete(id)
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (r *queryResolver) Studio(ctx context.Context, id uint) (*models.Studio, error) {
	panic(fmt.Errorf("Studio was not implemented"))
}

func (r *queryResolver) Studios(ctx context.Context) ([]*models.Studio, error) {
	panic(fmt.Errorf("Studios was not implemented"))
}

func (r *studioResolver) School(ctx context.Context, obj *models.Studio) (*models.School, error) {
	repository := repositories.NewSchoolRepository(r.DB)
	school, err := repository.Find(obj.SchoolID)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (r *studioResolver) Rooms(ctx context.Context, obj *models.Studio) ([]*models.Room, error) {
	repository := repositories.NewRoomRepository(r.DB)
	rooms, err := repository.FindAll(obj.ID)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

// Studio returns generated.StudioResolver implementation.
func (r *Resolver) Studio() generated.StudioResolver { return &studioResolver{r} }

type studioResolver struct{ *Resolver }
