// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/matsuokashuhei/landin/ent/predicate"
	"github.com/matsuokashuhei/landin/ent/room"
	"github.com/matsuokashuhei/landin/ent/schedule"
	"github.com/matsuokashuhei/landin/ent/studio"
)

// RoomUpdate is the builder for updating Room entities.
type RoomUpdate struct {
	config
	hooks    []Hook
	mutation *RoomMutation
}

// Where appends a list predicates to the RoomUpdate builder.
func (ru *RoomUpdate) Where(ps ...predicate.Room) *RoomUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *RoomUpdate) SetUpdateTime(t time.Time) *RoomUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// SetName sets the "name" field.
func (ru *RoomUpdate) SetName(s string) *RoomUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetCapacity sets the "capacity" field.
func (ru *RoomUpdate) SetCapacity(i int) *RoomUpdate {
	ru.mutation.ResetCapacity()
	ru.mutation.SetCapacity(i)
	return ru
}

// AddCapacity adds i to the "capacity" field.
func (ru *RoomUpdate) AddCapacity(i int) *RoomUpdate {
	ru.mutation.AddCapacity(i)
	return ru
}

// SetStudioID sets the "studio" edge to the Studio entity by ID.
func (ru *RoomUpdate) SetStudioID(id int) *RoomUpdate {
	ru.mutation.SetStudioID(id)
	return ru
}

// SetStudio sets the "studio" edge to the Studio entity.
func (ru *RoomUpdate) SetStudio(s *Studio) *RoomUpdate {
	return ru.SetStudioID(s.ID)
}

// AddScheduleIDs adds the "schedules" edge to the Schedule entity by IDs.
func (ru *RoomUpdate) AddScheduleIDs(ids ...int) *RoomUpdate {
	ru.mutation.AddScheduleIDs(ids...)
	return ru
}

// AddSchedules adds the "schedules" edges to the Schedule entity.
func (ru *RoomUpdate) AddSchedules(s ...*Schedule) *RoomUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.AddScheduleIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ru *RoomUpdate) Mutation() *RoomMutation {
	return ru.mutation
}

// ClearStudio clears the "studio" edge to the Studio entity.
func (ru *RoomUpdate) ClearStudio() *RoomUpdate {
	ru.mutation.ClearStudio()
	return ru
}

// ClearSchedules clears all "schedules" edges to the Schedule entity.
func (ru *RoomUpdate) ClearSchedules() *RoomUpdate {
	ru.mutation.ClearSchedules()
	return ru
}

// RemoveScheduleIDs removes the "schedules" edge to Schedule entities by IDs.
func (ru *RoomUpdate) RemoveScheduleIDs(ids ...int) *RoomUpdate {
	ru.mutation.RemoveScheduleIDs(ids...)
	return ru
}

// RemoveSchedules removes "schedules" edges to Schedule entities.
func (ru *RoomUpdate) RemoveSchedules(s ...*Schedule) *RoomUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.RemoveScheduleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoomUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoomUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoomUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoomUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RoomUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := room.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoomUpdate) check() error {
	if _, ok := ru.mutation.StudioID(); ru.mutation.StudioCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Room.studio"`)
	}
	return nil
}

func (ru *RoomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: room.FieldUpdateTime,
		})
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldName,
		})
	}
	if value, ok := ru.mutation.Capacity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldCapacity,
		})
	}
	if value, ok := ru.mutation.AddedCapacity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldCapacity,
		})
	}
	if ru.mutation.StudioCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.StudioTable,
			Columns: []string{room.StudioColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studio.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.StudioIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.StudioTable,
			Columns: []string{room.StudioColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.SchedulesTable,
			Columns: []string{room.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedSchedulesIDs(); len(nodes) > 0 && !ru.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.SchedulesTable,
			Columns: []string{room.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.SchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.SchedulesTable,
			Columns: []string{room.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RoomUpdateOne is the builder for updating a single Room entity.
type RoomUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoomMutation
}

// SetUpdateTime sets the "update_time" field.
func (ruo *RoomUpdateOne) SetUpdateTime(t time.Time) *RoomUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// SetName sets the "name" field.
func (ruo *RoomUpdateOne) SetName(s string) *RoomUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetCapacity sets the "capacity" field.
func (ruo *RoomUpdateOne) SetCapacity(i int) *RoomUpdateOne {
	ruo.mutation.ResetCapacity()
	ruo.mutation.SetCapacity(i)
	return ruo
}

// AddCapacity adds i to the "capacity" field.
func (ruo *RoomUpdateOne) AddCapacity(i int) *RoomUpdateOne {
	ruo.mutation.AddCapacity(i)
	return ruo
}

// SetStudioID sets the "studio" edge to the Studio entity by ID.
func (ruo *RoomUpdateOne) SetStudioID(id int) *RoomUpdateOne {
	ruo.mutation.SetStudioID(id)
	return ruo
}

// SetStudio sets the "studio" edge to the Studio entity.
func (ruo *RoomUpdateOne) SetStudio(s *Studio) *RoomUpdateOne {
	return ruo.SetStudioID(s.ID)
}

// AddScheduleIDs adds the "schedules" edge to the Schedule entity by IDs.
func (ruo *RoomUpdateOne) AddScheduleIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.AddScheduleIDs(ids...)
	return ruo
}

// AddSchedules adds the "schedules" edges to the Schedule entity.
func (ruo *RoomUpdateOne) AddSchedules(s ...*Schedule) *RoomUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.AddScheduleIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (ruo *RoomUpdateOne) Mutation() *RoomMutation {
	return ruo.mutation
}

// ClearStudio clears the "studio" edge to the Studio entity.
func (ruo *RoomUpdateOne) ClearStudio() *RoomUpdateOne {
	ruo.mutation.ClearStudio()
	return ruo
}

// ClearSchedules clears all "schedules" edges to the Schedule entity.
func (ruo *RoomUpdateOne) ClearSchedules() *RoomUpdateOne {
	ruo.mutation.ClearSchedules()
	return ruo
}

// RemoveScheduleIDs removes the "schedules" edge to Schedule entities by IDs.
func (ruo *RoomUpdateOne) RemoveScheduleIDs(ids ...int) *RoomUpdateOne {
	ruo.mutation.RemoveScheduleIDs(ids...)
	return ruo
}

// RemoveSchedules removes "schedules" edges to Schedule entities.
func (ruo *RoomUpdateOne) RemoveSchedules(s ...*Schedule) *RoomUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.RemoveScheduleIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoomUpdateOne) Select(field string, fields ...string) *RoomUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Room entity.
func (ruo *RoomUpdateOne) Save(ctx context.Context) (*Room, error) {
	var (
		err  error
		node *Room
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoomUpdateOne) SaveX(ctx context.Context) *Room {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoomUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoomUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RoomUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := room.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoomUpdateOne) check() error {
	if _, ok := ruo.mutation.StudioID(); ruo.mutation.StudioCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Room.studio"`)
	}
	return nil
}

func (ruo *RoomUpdateOne) sqlSave(ctx context.Context) (_node *Room, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   room.Table,
			Columns: room.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: room.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Room.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, room.FieldID)
		for _, f := range fields {
			if !room.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != room.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: room.FieldUpdateTime,
		})
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldName,
		})
	}
	if value, ok := ruo.mutation.Capacity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldCapacity,
		})
	}
	if value, ok := ruo.mutation.AddedCapacity(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: room.FieldCapacity,
		})
	}
	if ruo.mutation.StudioCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.StudioTable,
			Columns: []string{room.StudioColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studio.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.StudioIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   room.StudioTable,
			Columns: []string{room.StudioColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studio.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.SchedulesTable,
			Columns: []string{room.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedSchedulesIDs(); len(nodes) > 0 && !ruo.mutation.SchedulesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.SchedulesTable,
			Columns: []string{room.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.SchedulesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.SchedulesTable,
			Columns: []string{room.SchedulesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: schedule.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Room{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{room.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
