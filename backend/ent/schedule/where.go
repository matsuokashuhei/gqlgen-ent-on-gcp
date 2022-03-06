// Code generated by entc, DO NOT EDIT.

package schedule

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/matsuokashuhei/landin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// DayOfWeek applies equality check predicate on the "day_of_week" field. It's identical to DayOfWeekEQ.
func DayOfWeek(v int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDayOfWeek), v))
	})
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Schedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Schedule(func(s *sql.Selector) {
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
func CreateTimeNotIn(vs ...time.Time) predicate.Schedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Schedule(func(s *sql.Selector) {
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
func CreateTimeGT(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Schedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Schedule(func(s *sql.Selector) {
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
func UpdateTimeNotIn(vs ...time.Time) predicate.Schedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Schedule(func(s *sql.Selector) {
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
func UpdateTimeGT(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// DayOfWeekEQ applies the EQ predicate on the "day_of_week" field.
func DayOfWeekEQ(v int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDayOfWeek), v))
	})
}

// DayOfWeekNEQ applies the NEQ predicate on the "day_of_week" field.
func DayOfWeekNEQ(v int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDayOfWeek), v))
	})
}

// DayOfWeekIn applies the In predicate on the "day_of_week" field.
func DayOfWeekIn(vs ...int) predicate.Schedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Schedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDayOfWeek), v...))
	})
}

// DayOfWeekNotIn applies the NotIn predicate on the "day_of_week" field.
func DayOfWeekNotIn(vs ...int) predicate.Schedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Schedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDayOfWeek), v...))
	})
}

// DayOfWeekGT applies the GT predicate on the "day_of_week" field.
func DayOfWeekGT(v int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDayOfWeek), v))
	})
}

// DayOfWeekGTE applies the GTE predicate on the "day_of_week" field.
func DayOfWeekGTE(v int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDayOfWeek), v))
	})
}

// DayOfWeekLT applies the LT predicate on the "day_of_week" field.
func DayOfWeekLT(v int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDayOfWeek), v))
	})
}

// DayOfWeekLTE applies the LTE predicate on the "day_of_week" field.
func DayOfWeekLTE(v int) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDayOfWeek), v))
	})
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartTime), v))
	})
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...string) predicate.Schedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Schedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartTime), v...))
	})
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...string) predicate.Schedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Schedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartTime), v...))
	})
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartTime), v))
	})
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartTime), v))
	})
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartTime), v))
	})
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartTime), v))
	})
}

// StartTimeContains applies the Contains predicate on the "start_time" field.
func StartTimeContains(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStartTime), v))
	})
}

// StartTimeHasPrefix applies the HasPrefix predicate on the "start_time" field.
func StartTimeHasPrefix(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStartTime), v))
	})
}

// StartTimeHasSuffix applies the HasSuffix predicate on the "start_time" field.
func StartTimeHasSuffix(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStartTime), v))
	})
}

// StartTimeEqualFold applies the EqualFold predicate on the "start_time" field.
func StartTimeEqualFold(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStartTime), v))
	})
}

// StartTimeContainsFold applies the ContainsFold predicate on the "start_time" field.
func StartTimeContainsFold(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStartTime), v))
	})
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndTime), v))
	})
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...string) predicate.Schedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Schedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndTime), v...))
	})
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...string) predicate.Schedule {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Schedule(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndTime), v...))
	})
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndTime), v))
	})
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndTime), v))
	})
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndTime), v))
	})
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndTime), v))
	})
}

// EndTimeContains applies the Contains predicate on the "end_time" field.
func EndTimeContains(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEndTime), v))
	})
}

// EndTimeHasPrefix applies the HasPrefix predicate on the "end_time" field.
func EndTimeHasPrefix(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEndTime), v))
	})
}

// EndTimeHasSuffix applies the HasSuffix predicate on the "end_time" field.
func EndTimeHasSuffix(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEndTime), v))
	})
}

// EndTimeEqualFold applies the EqualFold predicate on the "end_time" field.
func EndTimeEqualFold(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEndTime), v))
	})
}

// EndTimeContainsFold applies the ContainsFold predicate on the "end_time" field.
func EndTimeContainsFold(v string) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEndTime), v))
	})
}

// HasRoom applies the HasEdge predicate on the "room" edge.
func HasRoom() predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomWith applies the HasEdge predicate on the "room" edge with a given conditions (other predicates).
func HasRoomWith(preds ...predicate.Room) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClasses applies the HasEdge predicate on the "classes" edge.
func HasClasses() predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClassesTable, ClassesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassesWith applies the HasEdge predicate on the "classes" edge with a given conditions (other predicates).
func HasClassesWith(preds ...predicate.Class) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ClassesTable, ClassesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClass applies the HasEdge predicate on the "class" edge.
func HasClass() predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassWith applies the HasEdge predicate on the "class" edge with a given conditions (other predicates).
func HasClassWith(preds ...predicate.Class) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ClassInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Schedule) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Schedule) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
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
func Not(p predicate.Schedule) predicate.Schedule {
	return predicate.Schedule(func(s *sql.Selector) {
		p(s.Not())
	})
}
