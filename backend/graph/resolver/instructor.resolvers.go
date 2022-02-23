package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *mutationResolver) CreateInstructor(ctx context.Context, input model.CreateInstructorInput) (*ent.Instructor, error) {
	repository := repositories.NewInstructorRepository(r.client)
	return repository.Create(ctx, input)
}

func (r *mutationResolver) UpdateInstructor(ctx context.Context, input model.UpdateInstructorInput) (*ent.Instructor, error) {
	repository := repositories.NewInstructorRepository(r.client)
	return repository.Update(ctx, input)
}

func (r *mutationResolver) DeleteInstructor(ctx context.Context, input model.DeleteInstructorInput) (*ent.Instructor, error) {
	repository := repositories.NewInstructorRepository(r.client)
	return repository.Delete(ctx, input)
}

func (r *queryResolver) Instructor(ctx context.Context, id int) (*ent.Instructor, error) {
	repository := repositories.NewInstructorRepository(r.client)
	return repository.Find(ctx, id)
}

func (r *queryResolver) Instructors(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.InstructorOrder) (*ent.InstructorConnection, error) {
	return r.client.Instructor.Query().Paginate(ctx, after, first, before, last, ent.WithInstructorOrder(orderBy))
}
