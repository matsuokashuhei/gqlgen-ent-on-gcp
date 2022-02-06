package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/ent/instructor"
	"github.com/matsuokashuhei/landin/graph/model"
)

func (r *mutationResolver) CreateInstructor(ctx context.Context, input model.CreateInstructorInput) (*ent.Instructor, error) {
	instructor, err := r.client.Instructor.Create().
		SetName(input.Name).
		SetSyllabicCharacters(input.SyllabicCharacters).
		SetBiography(*input.Biography).
		SetPhoneNumber(*input.PhoneNumber).
		SetEmail(*input.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return instructor, nil
}

func (r *queryResolver) Instructor(ctx context.Context, id int) (*ent.Instructor, error) {
	instructor, err := r.client.Instructor.Query().Where(instructor.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return instructor, nil
}

func (r *queryResolver) InstructorsConnection(ctx context.Context, offset *int, limit *int) (*model.InstructorsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}
