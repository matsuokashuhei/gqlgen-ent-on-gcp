package repositories

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
)

type ScheduleRepository struct {
	client *ent.Client
}

func (r *ScheduleRepository) Find(ctx context.Context, id int) (*ent.Schedule, error) {
	return r.client.Schedule.Get(ctx, id)
}

func NewScheduleRepository(client *ent.Client) *ScheduleRepository {
	return &ScheduleRepository{client: client}
}
