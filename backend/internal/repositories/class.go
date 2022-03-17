package repositories

import (
	"context"
	"time"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
)

type ClassRepository struct {
	client *ent.Client
}

func (r *ClassRepository) Find(ctx context.Context, id int) (*ent.Class, error) {
	return r.client.Class.Get(ctx, id)
}

func (r *ClassRepository) Create(ctx context.Context, input model.CreateClassInput) (*ent.Class, error) {
	creation := r.client.Class.Create().
		SetName(input.Name).
		SetLevel(input.Level).
		SetTuition(input.Tuition).
		SetStartDate(time.Now()).
		SetInstructorID(input.InstructorID).
		SetScheduleID(input.ScheduleID)
	if input.EndDate != nil {
		creation = creation.SetEndDate(*input.EndDate)
	}
	return creation.Save(ctx)
}

func (r *ClassRepository) Update(ctx context.Context, input model.UpdateClassInput) (*ent.Class, error) {
	updation := r.client.Class.UpdateOneID(input.ID)
	if input.Name != nil {
		updation.SetName(*input.Name)
	}
	if input.Level != nil {
		updation.SetLevel(*input.Level)
	}
	if input.Tuition != nil {
		updation.SetTuition(*input.Tuition)
	}
	if input.StartDate != nil {
		updation.SetStartDate(*input.StartDate)
	}
	if input.EndDate != nil {
		updation.SetEndDate(*input.EndDate)
	}
	if input.InstructorID != nil {
		updation.SetInstructorID(*input.InstructorID)
	}
	if input.ScheduleID != nil {
		updation.SetScheduleID(*input.ScheduleID)
	}
	return updation.Save(ctx)
}

func (r *ClassRepository) Delete(ctx context.Context, id int) (*ent.Class, error) {
	class, err := r.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := r.client.Class.DeleteOne(class).Exec(ctx); err != nil {
		return nil, err
	}
	return class, nil
}

func NewClassRepository(client *ent.Client) *ClassRepository {
	return &ClassRepository{client: client}
}
