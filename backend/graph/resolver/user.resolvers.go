package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/auth"
	"github.com/matsuokashuhei/landin/internal/models"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	if !IsAuthenticated(ctx) {
		return nil, auth.Error()
	}
	firebaseUser, err := auth.CreateFirebaseUser(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Name:        strings.Split(firebaseUser.Email, "@")[0],
		FirebaseUID: firebaseUser.UID,
	}
	repository := repositories.NewUserRepository(r.DB)
	_, err = repository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) User(ctx context.Context, id uint) (*models.User, error) {
	if !IsAuthenticated(ctx) {
		return nil, auth.Error()
	}
	repository := repositories.NewUserRepository(r.DB)
	user, err := repository.Find(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	if !IsAuthenticated(ctx) {
		return nil, auth.Error()
	}
	repository := repositories.NewUserRepository(r.DB)
	users, err := repository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
