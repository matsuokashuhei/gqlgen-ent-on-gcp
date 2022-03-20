package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/matsuokashuhei/landin/ent"
	"github.com/matsuokashuhei/landin/ent/class"
	"github.com/matsuokashuhei/landin/graph/generated"
	"github.com/matsuokashuhei/landin/internal/repositories"
)

func (r *queryResolver) Schedule(ctx context.Context, id int) (*ent.Schedule, error) {
	repository := repositories.NewScheduleRepository(r.client)
	return repository.Find(ctx, id)
}

func (r *scheduleResolver) Class(ctx context.Context, obj *ent.Schedule, date *time.Time) (*ent.Class, error) {
	// TODO: Move to repository
	var start time.Time
	if date != nil {
		start = *date
	} else {
		start = time.Now()
	}
	classes, err := obj.
		QueryClasses().
		Where(class.StartDateLTE(start)).
		Where(class.Or(class.EndDateGTE(start), class.EndDateIsNil())).
		All(ctx)
	if err != nil {
		return nil, err
	}
	if len(classes) > 0 {
		return classes[0], nil
	} else {
		return nil, nil
	}
}

// Schedule returns generated.ScheduleResolver implementation.
func (r *Resolver) Schedule() generated.ScheduleResolver { return &scheduleResolver{r} }

type scheduleResolver struct{ *Resolver }
