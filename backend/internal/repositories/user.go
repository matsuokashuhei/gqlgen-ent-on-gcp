package repositories

import (
	"context"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/ent/user"
	"github.com/matsuokashuhei/landin/graph/model"
)

type UserRepository struct {
	client *ent.Client
}

func (r *UserRepository) All(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

func (r *UserRepository) Find(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Get(ctx, id)
}

func (r *UserRepository) FindByAuth(ctx context.Context, uid string) (*ent.User, error) {
	return r.client.User.Query().
		Where(user.FirebaseUID(uid)).
		First(ctx)
}

func (r *UserRepository) Create(ctx context.Context, input model.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().
		SetName(input.Email).
		Save(ctx)
}

func (r *UserRepository) Update(ctx context.Context, input model.UpdateUserInput) (*ent.User, error) {
	return r.client.User.UpdateOneID(input.ID).SetName(input.Name).Save(ctx)
}

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{client: client}
}
