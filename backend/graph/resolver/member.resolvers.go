package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/graph/model"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *memberResolver) Gender(ctx context.Context, obj *ent.Member) (model.Gender, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMember(ctx context.Context, input model.CreateMemberInput) (*ent.Member, error) {
	repository := repositories.NewMemberRepsitory(r.client)
	return repository.Create(ctx, input)
}

func (r *mutationResolver) UpdateMember(ctx context.Context, input model.UpdateMemberInput) (*ent.Member, error) {
	repository := repositories.NewMemberRepsitory(r.client)
	return repository.Update(ctx, input)
}

func (r *mutationResolver) DeleteMember(ctx context.Context, input model.DeleteMemberInput) (*ent.Member, error) {
	repository := repositories.NewMemberRepsitory(r.client)
	return repository.Delete(ctx, input)
}

func (r *queryResolver) Member(ctx context.Context, id int) (*ent.Member, error) {
	repository := repositories.NewMemberRepsitory(r.client)
	return repository.Find(ctx, id)
}

func (r *queryResolver) Members(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MemberOrder) (*ent.MemberConnection, error) {
	return r.client.Member.Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			ent.WithMemberOrder(orderBy),
		)
}

// Member returns generated.MemberResolver implementation.
func (r *Resolver) Member() generated.MemberResolver { return &memberResolver{r} }

type memberResolver struct{ *Resolver }
