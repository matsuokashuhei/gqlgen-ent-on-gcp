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

func (r *mutationResolver) CreateRoom(ctx context.Context, input model.CreateRoomInput) (*models.Room, error) {
	repository := repositories.NewRoomRepository(r.DB)
	room := &models.Room{
		Name:     input.Name,
		StudioID: input.StudioID,
	}
	_, err := repository.Create(room)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *mutationResolver) UpdateRoom(ctx context.Context, input model.UpdateRoomInput) (*models.Room, error) {
	repository := repositories.NewRoomRepository(r.DB)
	var room, err = repository.Find(input.ID)
	if err != nil {
		return nil, err
	}
	if input.Name == nil && input.StudioID == nil {
		return room, nil
	}
	if input.Name != nil {
		room.Name = *input.Name
	}
	if input.StudioID != nil {
		room.StudioID = *input.StudioID
	}
	_, err = repository.Update(room)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *mutationResolver) DeleteRoom(ctx context.Context, id uint) (*models.Room, error) {
	repository := repositories.NewRoomRepository(r.DB)
	room, err := repository.Delete(id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *queryResolver) Room(ctx context.Context, id uint) (*models.Room, error) {
	panic(fmt.Errorf("Room was not implemented"))
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*models.Room, error) {
	panic(fmt.Errorf("Rooms was not implemented"))
}

func (r *roomResolver) Studio(ctx context.Context, obj *models.Room) (*models.Studio, error) {
	repository := repositories.NewStudioRepository(r.DB)
	studio, err := repository.Find(obj.StudioID)
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (r *roomResolver) Schedules(ctx context.Context, obj *models.Room) ([]*models.Schedule, error) {
	repository := repositories.NewScheduleRepository(r.DB)
	schedules, err := repository.FindAll(obj.ID)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

type roomResolver struct{ *Resolver }
