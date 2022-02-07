package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/ent/user"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/auth"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	if !IsAuthenticated(ctx) {
		return nil, auth.Error()
	}
	firebaseUser, err := auth.CreateFirebaseUser(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	user, err := r.client.User.Create().SetName(strings.Split(firebaseUser.Email, "@")[0]).SetFirebaseUID(firebaseUser.UID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id int) (*ent.User, error) {
	if !IsAuthenticated(ctx) {
		return nil, auth.Error()
	}
	user, err := r.client.User.Query().Where(user.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	if !IsAuthenticated(ctx) {
		return nil, auth.Error()
	}
	users, err := r.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
