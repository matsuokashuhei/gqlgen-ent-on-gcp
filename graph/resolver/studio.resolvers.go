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

func (r *mutationResolver) UpdateStudio(ctx context.Context, id int, input model.StudioInput) (*models.Studio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteStudio(ctx context.Context, id int) (*models.Studio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Studio(ctx context.Context, id int) (*models.Studio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Studios(ctx context.Context) ([]*models.Studio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *studioResolver) ID(ctx context.Context, obj *models.Studio) (int, error) {
	return int(obj.ID), nil
}

func (r *studioResolver) School(ctx context.Context, obj *models.Studio) (*models.School, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *studioResolver) Rooms(ctx context.Context, obj *models.Studio) ([]*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *studioResolver) CreatedAt(ctx context.Context, obj *models.Studio) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

func (r *studioResolver) UpdatedAt(ctx context.Context, obj *models.Studio) (string, error) {
	return obj.UpdatedAt.Format(time.RFC3339), nil
}

// Studio returns generated.StudioResolver implementation.
func (r *Resolver) Studio() generated.StudioResolver { return &studioResolver{r} }

type studioResolver struct{ *Resolver }
