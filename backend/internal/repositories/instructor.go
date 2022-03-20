package repositories

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
)

type InstructorRepository struct {
	client *ent.Client
}

func (r *InstructorRepository) Create(ctx context.Context, input model.CreateInstructorInput) (*ent.Instructor, error) {
	create := r.client.Instructor.Create().
		SetName(input.Name).
		SetKana(input.Kana)
	if input.Biography != nil {
		create = create.SetBiography(*input.Biography)
	}
	if input.PhoneNumber != nil {
		create = create.SetPhoneNumber(*input.PhoneNumber)
	}
	if input.Email != nil {
		create = create.SetEmail(*input.Email)
	}
	return create.Save(ctx)
}

func (r *InstructorRepository) Update(ctx context.Context, input model.UpdateInstructorInput) (*ent.Instructor, error) {
	update := r.client.Instructor.UpdateOneID(input.ID)
	if input.Name != nil {
		update = update.SetName(*input.Name)
	}
	if input.Kana != nil {
		update = update.SetKana(*input.Kana)
	}
	if input.Biography != nil {
		update = update.SetBiography(*input.Biography)
	}
	if input.PhoneNumber != nil {
		update = update.SetPhoneNumber(*input.PhoneNumber)
	}
	if input.Email != nil {
		update = update.SetEmail(*input.Email)
	}
	return update.Save(ctx)
}

func (r *InstructorRepository) Delete(ctx context.Context, input model.DeleteInstructorInput) (*ent.Instructor, error) {
	instructor, err := r.Find(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	if err := r.client.Instructor.DeleteOne(instructor).Exec(ctx); err != nil {
		return nil, err
	}
	return instructor, nil
}

func (r *InstructorRepository) Find(ctx context.Context, id int) (*ent.Instructor, error) {
	return r.client.Instructor.Get(ctx, id)
}

func NewInstructorRepository(client *ent.Client) *InstructorRepository {
	return &InstructorRepository{client: client}
}
