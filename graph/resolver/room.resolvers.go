package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/models"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, input model.RoomInput) (*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateRoom(ctx context.Context, id uint, input model.RoomInput) (*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRoom(ctx context.Context, id uint) (*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Room(ctx context.Context, id uint) (*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *roomResolver) Studio(ctx context.Context, obj *models.Room) (*models.Studio, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type roomResolver struct{ *Resolver }
