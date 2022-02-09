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
	return r.client.Instructor.Create().
		SetName(input.Name).
		SetSyllabicCharacters(input.SyllabicCharacters).
		SetBiography(*input.Biography).
		SetPhoneNumber(*input.PhoneNumber).
		SetEmail(*input.Email).
		Save(ctx)
}

func (r *InstructorRepository) Find(ctx context.Context, id int) (*ent.Instructor, error) {
	return r.client.Instructor.Get(ctx, id)
}

func NewInstructorRepository(client *ent.Client) *InstructorRepository {
	return &InstructorRepository{client: client}
}
