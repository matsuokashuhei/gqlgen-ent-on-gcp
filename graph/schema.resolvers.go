package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"
	"time"

	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/internal/models"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *queryResolver) Schools(ctx context.Context) ([]*models.School, error) {
	repository := repositories.NewSchoolRepository(r.DB)
	var schools []*models.School = repository.FindAll()
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// School returns generated.SchoolResolver implementation.
func (r *Resolver) School() generated.SchoolResolver { return &schoolResolver{r} }

type queryResolver struct{ *Resolver }
type schoolResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
