package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"strings"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/auth"
)

func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpInput) (*ent.User, error) {
	firebaseUser, err := auth.CreateFirebaseUser(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	user, err := r.client.
		User.
		Create().
		SetName(strings.Split(firebaseUser.Email, "@")[0]).
		SetFirebaseUID(firebaseUser.UID).
		Save(ctx)
	return user, nil
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*ent.User, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, exists := gc.Get(auth.USER_KEY)
	log.Printf("user: %v, exists: %t", user, exists)
	if exists == false {
		return nil, auth.Error()
	}
	return user.(*ent.User), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
