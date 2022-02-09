package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/graph/generated"
)

func (r *scheduleResolver) CreatedAt(ctx context.Context, obj *ent.Schedule) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *scheduleResolver) UpdatedAt(ctx context.Context, obj *ent.Schedule) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Schedule returns generated.ScheduleResolver implementation.
func (r *Resolver) Schedule() generated.ScheduleResolver { return &scheduleResolver{r} }

type scheduleResolver struct{ *Resolver }
