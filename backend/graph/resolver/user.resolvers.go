package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/models"
	"github.com/matsuokashuhei/landin/internal/repositories"
	"github.com/matsuokashuhei/landin/pkg/firebase"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	record, err := firebase.CreateUser(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	repository := repositories.NewUserRepository(r.DB)
	user := &models.User{
		Name:        input.Name,
		FirebaseUID: record.UID,
	}
	_, err = repository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
