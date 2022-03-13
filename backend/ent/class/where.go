// Code generated by entc, DO NOT EDIT.

package class

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/matsuokashuhei/landin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Tuition applies equality check predicate on the "tuition" field. It's identical to TuitionEQ.
func Tuition(v int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTuition), v))
	})
}

// StartDate applies equality check predicate on the "start_date" field. It's identical to StartDateEQ.
func StartDate(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// EndDate applies equality check predicate on the "end_date" field. It's identical to EndDateEQ.
func EndDate(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndDate), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// TuitionEQ applies the EQ predicate on the "tuition" field.
func TuitionEQ(v int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTuition), v))
	})
}

// TuitionNEQ applies the NEQ predicate on the "tuition" field.
func TuitionNEQ(v int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTuition), v))
	})
}

// TuitionIn applies the In predicate on the "tuition" field.
func TuitionIn(vs ...int) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTuition), v...))
	})
}

// TuitionNotIn applies the NotIn predicate on the "tuition" field.
func TuitionNotIn(vs ...int) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTuition), v...))
	})
}

// TuitionGT applies the GT predicate on the "tuition" field.
func TuitionGT(v int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTuition), v))
	})
}

// TuitionGTE applies the GTE predicate on the "tuition" field.
func TuitionGTE(v int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTuition), v))
	})
}

// TuitionLT applies the LT predicate on the "tuition" field.
func TuitionLT(v int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTuition), v))
	})
}

// TuitionLTE applies the LTE predicate on the "tuition" field.
func TuitionLTE(v int) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTuition), v))
	})
}

// StartDateEQ applies the EQ predicate on the "start_date" field.
func StartDateEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartDate), v))
	})
}

// StartDateNEQ applies the NEQ predicate on the "start_date" field.
func StartDateNEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartDate), v))
	})
}

// StartDateIn applies the In predicate on the "start_date" field.
func StartDateIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartDate), v...))
	})
}

// StartDateNotIn applies the NotIn predicate on the "start_date" field.
func StartDateNotIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartDate), v...))
	})
}

// StartDateGT applies the GT predicate on the "start_date" field.
func StartDateGT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartDate), v))
	})
}

// StartDateGTE applies the GTE predicate on the "start_date" field.
func StartDateGTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartDate), v))
	})
}

// StartDateLT applies the LT predicate on the "start_date" field.
func StartDateLT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartDate), v))
	})
}

// StartDateLTE applies the LTE predicate on the "start_date" field.
func StartDateLTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartDate), v))
	})
}

// EndDateEQ applies the EQ predicate on the "end_date" field.
func EndDateEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndDate), v))
	})
}

// EndDateNEQ applies the NEQ predicate on the "end_date" field.
func EndDateNEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndDate), v))
	})
}

// EndDateIn applies the In predicate on the "end_date" field.
func EndDateIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndDate), v...))
	})
}

// EndDateNotIn applies the NotIn predicate on the "end_date" field.
func EndDateNotIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndDate), v...))
	})
}

// EndDateGT applies the GT predicate on the "end_date" field.
func EndDateGT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndDate), v))
	})
}

// EndDateGTE applies the GTE predicate on the "end_date" field.
func EndDateGTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndDate), v))
	})
}

// EndDateLT applies the LT predicate on the "end_date" field.
func EndDateLT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndDate), v))
	})
}

// EndDateLTE applies the LTE predicate on the "end_date" field.
func EndDateLTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndDate), v))
	})
}

// EndDateIsNil applies the IsNil predicate on the "end_date" field.
func EndDateIsNil() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndDate)))
	})
}

// EndDateNotNil applies the NotNil predicate on the "end_date" field.
func EndDateNotNil() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndDate)))
	})
}

// HasSchedule applies the HasEdge predicate on the "schedule" edge.
func HasSchedule() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ScheduleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ScheduleTable, ScheduleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScheduleWith applies the HasEdge predicate on the "schedule" edge with a given conditions (other predicates).
func HasScheduleWith(preds ...predicate.Schedule) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ScheduleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ScheduleTable, ScheduleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstructor applies the HasEdge predicate on the "instructor" edge.
func HasInstructor() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstructorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstructorTable, InstructorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstructorWith applies the HasEdge predicate on the "instructor" edge with a given conditions (other predicates).
func HasInstructorWith(preds ...predicate.Instructor) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InstructorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstructorTable, InstructorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		p(s.Not())
	})
}
