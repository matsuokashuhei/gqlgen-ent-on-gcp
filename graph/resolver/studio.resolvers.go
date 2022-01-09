package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/models"
)

func (r *mutationResolver) CreateStudio(ctx context.Context, input model.StudioInput) (*models.Studio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStudio(ctx context.Context, id uint, input model.StudioInput) (*models.Studio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteStudio(ctx context.Context, id uint) (*models.Studio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Studio(ctx context.Context, id uint) (*models.Studio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Studios(ctx context.Context) ([]*models.Studio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *studioResolver) School(ctx context.Context, obj *models.Studio) (*models.School, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *studioResolver) Rooms(ctx context.Context, obj *models.Studio) ([]*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

// Studio returns generated.StudioResolver implementation.
func (r *Resolver) Studio() generated.StudioResolver { return &studioResolver{r} }

type studioResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *studioResolver) ID(ctx context.Context, obj *models.Studio) (int, error) {
	return int(obj.ID), nil
}
func (r *studioResolver) CreatedAt(ctx context.Context, obj *models.Studio) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}
func (r *studioResolver) UpdatedAt(ctx context.Context, obj *models.Studio) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}
