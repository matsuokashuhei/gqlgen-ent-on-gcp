package repositories

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/matsuokashuhei/landin/graph/model"
)

type AuthRepository struct {
	client *auth.Client
}

func (r *AuthRepository) Create(ctx context.Context, input model.SignUpInput) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).Email(input.Email).EmailVerified(false).Password(input.Password)
	return r.client.CreateUser(ctx, params)
}

func NewAuthRepository(client *auth.Client) *AuthRepository {
	return &AuthRepository{client: client}
}
