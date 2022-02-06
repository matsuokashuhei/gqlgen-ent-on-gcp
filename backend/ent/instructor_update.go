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
	"github.com/matsuokashuhei/landin/ent/class"
	"github.com/matsuokashuhei/landin/ent/instructor"
	"github.com/matsuokashuhei/landin/ent/predicate"
)

// InstructorUpdate is the builder for updating Instructor entities.
type InstructorUpdate struct {
	config
	hooks    []Hook
	mutation *InstructorMutation
}

// Where appends a list predicates to the InstructorUpdate builder.
func (iu *InstructorUpdate) Where(ps ...predicate.Instructor) *InstructorUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdateTime sets the "update_time" field.
func (iu *InstructorUpdate) SetUpdateTime(t time.Time) *InstructorUpdate {
	iu.mutation.SetUpdateTime(t)
	return iu
}

// SetName sets the "name" field.
func (iu *InstructorUpdate) SetName(s string) *InstructorUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetSyllabicCharacters sets the "syllabic_characters" field.
func (iu *InstructorUpdate) SetSyllabicCharacters(s string) *InstructorUpdate {
	iu.mutation.SetSyllabicCharacters(s)
	return iu
}

// SetBiography sets the "biography" field.
func (iu *InstructorUpdate) SetBiography(s string) *InstructorUpdate {
	iu.mutation.SetBiography(s)
	return iu
}

// SetNillableBiography sets the "biography" field if the given value is not nil.
func (iu *InstructorUpdate) SetNillableBiography(s *string) *InstructorUpdate {
	if s != nil {
		iu.SetBiography(*s)
	}
	return iu
}

// ClearBiography clears the value of the "biography" field.
func (iu *InstructorUpdate) ClearBiography() *InstructorUpdate {
	iu.mutation.ClearBiography()
	return iu
}

// SetPhoneNumber sets the "phone_number" field.
func (iu *InstructorUpdate) SetPhoneNumber(s string) *InstructorUpdate {
	iu.mutation.SetPhoneNumber(s)
	return iu
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (iu *InstructorUpdate) SetNillablePhoneNumber(s *string) *InstructorUpdate {
	if s != nil {
		iu.SetPhoneNumber(*s)
	}
	return iu
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (iu *InstructorUpdate) ClearPhoneNumber() *InstructorUpdate {
	iu.mutation.ClearPhoneNumber()
	return iu
}

// SetEmail sets the "email" field.
func (iu *InstructorUpdate) SetEmail(s string) *InstructorUpdate {
	iu.mutation.SetEmail(s)
	return iu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (iu *InstructorUpdate) SetNillableEmail(s *string) *InstructorUpdate {
	if s != nil {
		iu.SetEmail(*s)
	}
	return iu
}

// ClearEmail clears the value of the "email" field.
func (iu *InstructorUpdate) ClearEmail() *InstructorUpdate {
	iu.mutation.ClearEmail()
	return iu
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (iu *InstructorUpdate) AddClassIDs(ids ...int) *InstructorUpdate {
	iu.mutation.AddClassIDs(ids...)
	return iu
}

// AddClasses adds the "classes" edges to the Class entity.
func (iu *InstructorUpdate) AddClasses(c ...*Class) *InstructorUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iu.AddClassIDs(ids...)
}

// Mutation returns the InstructorMutation object of the builder.
func (iu *InstructorUpdate) Mutation() *InstructorMutation {
	return iu.mutation
}

// ClearClasses clears all "classes" edges to the Class entity.
func (iu *InstructorUpdate) ClearClasses() *InstructorUpdate {
	iu.mutation.ClearClasses()
	return iu
}

// RemoveClassIDs removes the "classes" edge to Class entities by IDs.
func (iu *InstructorUpdate) RemoveClassIDs(ids ...int) *InstructorUpdate {
	iu.mutation.RemoveClassIDs(ids...)
	return iu
}

// RemoveClasses removes "classes" edges to Class entities.
func (iu *InstructorUpdate) RemoveClasses(c ...*Class) *InstructorUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iu.RemoveClassIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InstructorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstructorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InstructorUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InstructorUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InstructorUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *InstructorUpdate) defaults() {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := instructor.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
}

func (iu *InstructorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instructor.Table,
			Columns: instructor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instructor.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instructor.FieldUpdateTime,
		})
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructor.FieldName,
		})
	}
	if value, ok := iu.mutation.SyllabicCharacters(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructor.FieldSyllabicCharacters,
		})
	}
	if value, ok := iu.mutation.Biography(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructor.FieldBiography,
		})
	}
	if iu.mutation.BiographyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instructor.FieldBiography,
		})
	}
	if value, ok := iu.mutation.PhoneNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructor.FieldPhoneNumber,
		})
	}
	if iu.mutation.PhoneNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instructor.FieldPhoneNumber,
		})
	}
	if value, ok := iu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructor.FieldEmail,
		})
	}
	if iu.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instructor.FieldEmail,
		})
	}
	if iu.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instructor.ClassesTable,
			Columns: []string{instructor.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedClassesIDs(); len(nodes) > 0 && !iu.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instructor.ClassesTable,
			Columns: []string{instructor.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instructor.ClassesTable,
			Columns: []string{instructor.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instructor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// InstructorUpdateOne is the builder for updating a single Instructor entity.
type InstructorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InstructorMutation
}

// SetUpdateTime sets the "update_time" field.
func (iuo *InstructorUpdateOne) SetUpdateTime(t time.Time) *InstructorUpdateOne {
	iuo.mutation.SetUpdateTime(t)
	return iuo
}

// SetName sets the "name" field.
func (iuo *InstructorUpdateOne) SetName(s string) *InstructorUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetSyllabicCharacters sets the "syllabic_characters" field.
func (iuo *InstructorUpdateOne) SetSyllabicCharacters(s string) *InstructorUpdateOne {
	iuo.mutation.SetSyllabicCharacters(s)
	return iuo
}

// SetBiography sets the "biography" field.
func (iuo *InstructorUpdateOne) SetBiography(s string) *InstructorUpdateOne {
	iuo.mutation.SetBiography(s)
	return iuo
}

// SetNillableBiography sets the "biography" field if the given value is not nil.
func (iuo *InstructorUpdateOne) SetNillableBiography(s *string) *InstructorUpdateOne {
	if s != nil {
		iuo.SetBiography(*s)
	}
	return iuo
}

// ClearBiography clears the value of the "biography" field.
func (iuo *InstructorUpdateOne) ClearBiography() *InstructorUpdateOne {
	iuo.mutation.ClearBiography()
	return iuo
}

// SetPhoneNumber sets the "phone_number" field.
func (iuo *InstructorUpdateOne) SetPhoneNumber(s string) *InstructorUpdateOne {
	iuo.mutation.SetPhoneNumber(s)
	return iuo
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (iuo *InstructorUpdateOne) SetNillablePhoneNumber(s *string) *InstructorUpdateOne {
	if s != nil {
		iuo.SetPhoneNumber(*s)
	}
	return iuo
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (iuo *InstructorUpdateOne) ClearPhoneNumber() *InstructorUpdateOne {
	iuo.mutation.ClearPhoneNumber()
	return iuo
}

// SetEmail sets the "email" field.
func (iuo *InstructorUpdateOne) SetEmail(s string) *InstructorUpdateOne {
	iuo.mutation.SetEmail(s)
	return iuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (iuo *InstructorUpdateOne) SetNillableEmail(s *string) *InstructorUpdateOne {
	if s != nil {
		iuo.SetEmail(*s)
	}
	return iuo
}

// ClearEmail clears the value of the "email" field.
func (iuo *InstructorUpdateOne) ClearEmail() *InstructorUpdateOne {
	iuo.mutation.ClearEmail()
	return iuo
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (iuo *InstructorUpdateOne) AddClassIDs(ids ...int) *InstructorUpdateOne {
	iuo.mutation.AddClassIDs(ids...)
	return iuo
}

// AddClasses adds the "classes" edges to the Class entity.
func (iuo *InstructorUpdateOne) AddClasses(c ...*Class) *InstructorUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iuo.AddClassIDs(ids...)
}

// Mutation returns the InstructorMutation object of the builder.
func (iuo *InstructorUpdateOne) Mutation() *InstructorMutation {
	return iuo.mutation
}

// ClearClasses clears all "classes" edges to the Class entity.
func (iuo *InstructorUpdateOne) ClearClasses() *InstructorUpdateOne {
	iuo.mutation.ClearClasses()
	return iuo
}

// RemoveClassIDs removes the "classes" edge to Class entities by IDs.
func (iuo *InstructorUpdateOne) RemoveClassIDs(ids ...int) *InstructorUpdateOne {
	iuo.mutation.RemoveClassIDs(ids...)
	return iuo
}

// RemoveClasses removes "classes" edges to Class entities.
func (iuo *InstructorUpdateOne) RemoveClasses(c ...*Class) *InstructorUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iuo.RemoveClassIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InstructorUpdateOne) Select(field string, fields ...string) *InstructorUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Instructor entity.
func (iuo *InstructorUpdateOne) Save(ctx context.Context) (*Instructor, error) {
	var (
		err  error
		node *Instructor
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstructorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InstructorUpdateOne) SaveX(ctx context.Context) *Instructor {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InstructorUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InstructorUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *InstructorUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := instructor.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
}

func (iuo *InstructorUpdateOne) sqlSave(ctx context.Context) (_node *Instructor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instructor.Table,
			Columns: instructor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: instructor.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Instructor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, instructor.FieldID)
		for _, f := range fields {
			if !instructor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != instructor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instructor.FieldUpdateTime,
		})
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructor.FieldName,
		})
	}
	if value, ok := iuo.mutation.SyllabicCharacters(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructor.FieldSyllabicCharacters,
		})
	}
	if value, ok := iuo.mutation.Biography(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructor.FieldBiography,
		})
	}
	if iuo.mutation.BiographyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instructor.FieldBiography,
		})
	}
	if value, ok := iuo.mutation.PhoneNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructor.FieldPhoneNumber,
		})
	}
	if iuo.mutation.PhoneNumberCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instructor.FieldPhoneNumber,
		})
	}
	if value, ok := iuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instructor.FieldEmail,
		})
	}
	if iuo.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: instructor.FieldEmail,
		})
	}
	if iuo.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instructor.ClassesTable,
			Columns: []string{instructor.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedClassesIDs(); len(nodes) > 0 && !iuo.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instructor.ClassesTable,
			Columns: []string{instructor.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instructor.ClassesTable,
			Columns: []string{instructor.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Instructor{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instructor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
