package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpInput) (*ent.User, error) {
	user, err := r.resolveAuth(ctx, input)
	if err != nil {
		return nil, err
	}
	return r.resolveUser(ctx, user)
}

func (r *queryResolver) CurrentUser(ctx context.Context) (*ent.User, error) {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	user, exists := gc.Get("user")
	log.Printf("user: %v, exists: %t", user, exists)
	if exists == false {
		return nil, fmt.Errorf("access defnied")
	}
	return user.(*ent.User), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) resolveAuth(ctx context.Context, input model.SignUpInput) (*auth.UserRecord, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	repository := repositories.NewAuthRepository(auth)
	return repository.Create(ctx, input)
}
func (r *mutationResolver) resolveUser(ctx context.Context, input *auth.UserRecord) (*ent.User, error) {
	return r.client.User.Create().
		SetName(strings.Split(input.Email, "@")[0]).
		SetFirebaseUID(input.UID).
		Save(ctx)
}
