package repositories

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
)

type RoomRepository struct {
	client *ent.Client
}

func (r *RoomRepository) Find(ctx context.Context, id int) (*ent.Room, error) {
	return r.client.Room.Get(ctx, id)
}

func (r *RoomRepository) Create(ctx context.Context, input model.CreateRoomInput) (*ent.Room, error) {
	return r.client.Room.Create().
		SetName(input.Name).
		SetCapacity(input.Capacity).
		SetStudioID(input.StudioID).
		Save(ctx)
}

func (r *RoomRepository) Update(ctx context.Context, input model.UpdateRoomInput) (*ent.Room, error) {
	return r.client.Room.UpdateOneID(input.ID).
		SetName(*input.Name).
		SetCapacity(*input.Capacity).
		SetStudioID(*input.StudioID).
		Save(ctx)
}

func (r *RoomRepository) Delete(ctx context.Context, id int) (*ent.Room, error) {
	room, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := r.client.Room.DeleteOne(room).Exec(ctx); err != nil {
		return nil, err
	}
	return room, nil
}

func NewRoomRepository(client *ent.Client) *RoomRepository {
	return &RoomRepository{client: client}
}
