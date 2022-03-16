package repositories

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
)

type ClassRepository struct {
	client *ent.Client
}

func (r *ClassRepository) Create(ctx context.Context, input model.CreateClassInput) (*ent.Class, error) {
	creation := r.client.Class.Create().
		SetName(input.Name).
		SetLevel(input.Level).
		SetTuition(input.Tuition).
		SetStartDate(input.StartDate).
		SetInstructorID(input.InstructorID).
		SetScheduleID(input.ScheduleID)
	if input.EndDate != nil {
		creation = creation.SetEndDate(*input.EndDate)
	}
	return creation.Save(ctx)
}

func NewClassRepository(client *ent.Client) *ClassRepository {
	return &ClassRepository{client: client}
}
