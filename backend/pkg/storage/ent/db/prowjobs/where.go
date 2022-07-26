// Code generated by entc, DO NOT EDIT.

package prowjobs

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// JobID applies equality check predicate on the "job_id" field. It's identical to JobIDEQ.
func JobID(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJobID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v float64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDuration), v))
	})
}

// TestsCount applies equality check predicate on the "tests_count" field. It's identical to TestsCountEQ.
func TestsCount(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestsCount), v))
	})
}

// FailedCount applies equality check predicate on the "failed_count" field. It's identical to FailedCountEQ.
func FailedCount(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFailedCount), v))
	})
}

// SkippedCount applies equality check predicate on the "skipped_count" field. It's identical to SkippedCountEQ.
func SkippedCount(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkippedCount), v))
	})
}

// JobIDEQ applies the EQ predicate on the "job_id" field.
func JobIDEQ(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJobID), v))
	})
}

// JobIDNEQ applies the NEQ predicate on the "job_id" field.
func JobIDNEQ(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldJobID), v))
	})
}

// JobIDIn applies the In predicate on the "job_id" field.
func JobIDIn(vs ...string) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldJobID), v...))
	})
}

// JobIDNotIn applies the NotIn predicate on the "job_id" field.
func JobIDNotIn(vs ...string) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldJobID), v...))
	})
}

// JobIDGT applies the GT predicate on the "job_id" field.
func JobIDGT(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldJobID), v))
	})
}

// JobIDGTE applies the GTE predicate on the "job_id" field.
func JobIDGTE(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldJobID), v))
	})
}

// JobIDLT applies the LT predicate on the "job_id" field.
func JobIDLT(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldJobID), v))
	})
}

// JobIDLTE applies the LTE predicate on the "job_id" field.
func JobIDLTE(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldJobID), v))
	})
}

// JobIDContains applies the Contains predicate on the "job_id" field.
func JobIDContains(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldJobID), v))
	})
}

// JobIDHasPrefix applies the HasPrefix predicate on the "job_id" field.
func JobIDHasPrefix(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldJobID), v))
	})
}

// JobIDHasSuffix applies the HasSuffix predicate on the "job_id" field.
func JobIDHasSuffix(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldJobID), v))
	})
}

// JobIDEqualFold applies the EqualFold predicate on the "job_id" field.
func JobIDEqualFold(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldJobID), v))
	})
}

// JobIDContainsFold applies the ContainsFold predicate on the "job_id" field.
func JobIDContainsFold(v string) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldJobID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v float64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDuration), v))
	})
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v float64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDuration), v))
	})
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...float64) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDuration), v...))
	})
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...float64) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDuration), v...))
	})
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v float64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDuration), v))
	})
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v float64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDuration), v))
	})
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v float64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDuration), v))
	})
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v float64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDuration), v))
	})
}

// TestsCountEQ applies the EQ predicate on the "tests_count" field.
func TestsCountEQ(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestsCount), v))
	})
}

// TestsCountNEQ applies the NEQ predicate on the "tests_count" field.
func TestsCountNEQ(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTestsCount), v))
	})
}

// TestsCountIn applies the In predicate on the "tests_count" field.
func TestsCountIn(vs ...int64) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTestsCount), v...))
	})
}

// TestsCountNotIn applies the NotIn predicate on the "tests_count" field.
func TestsCountNotIn(vs ...int64) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTestsCount), v...))
	})
}

// TestsCountGT applies the GT predicate on the "tests_count" field.
func TestsCountGT(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTestsCount), v))
	})
}

// TestsCountGTE applies the GTE predicate on the "tests_count" field.
func TestsCountGTE(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTestsCount), v))
	})
}

// TestsCountLT applies the LT predicate on the "tests_count" field.
func TestsCountLT(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTestsCount), v))
	})
}

// TestsCountLTE applies the LTE predicate on the "tests_count" field.
func TestsCountLTE(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTestsCount), v))
	})
}

// FailedCountEQ applies the EQ predicate on the "failed_count" field.
func FailedCountEQ(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFailedCount), v))
	})
}

// FailedCountNEQ applies the NEQ predicate on the "failed_count" field.
func FailedCountNEQ(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFailedCount), v))
	})
}

// FailedCountIn applies the In predicate on the "failed_count" field.
func FailedCountIn(vs ...int64) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFailedCount), v...))
	})
}

// FailedCountNotIn applies the NotIn predicate on the "failed_count" field.
func FailedCountNotIn(vs ...int64) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFailedCount), v...))
	})
}

// FailedCountGT applies the GT predicate on the "failed_count" field.
func FailedCountGT(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFailedCount), v))
	})
}

// FailedCountGTE applies the GTE predicate on the "failed_count" field.
func FailedCountGTE(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFailedCount), v))
	})
}

// FailedCountLT applies the LT predicate on the "failed_count" field.
func FailedCountLT(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFailedCount), v))
	})
}

// FailedCountLTE applies the LTE predicate on the "failed_count" field.
func FailedCountLTE(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFailedCount), v))
	})
}

// SkippedCountEQ applies the EQ predicate on the "skipped_count" field.
func SkippedCountEQ(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSkippedCount), v))
	})
}

// SkippedCountNEQ applies the NEQ predicate on the "skipped_count" field.
func SkippedCountNEQ(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSkippedCount), v))
	})
}

// SkippedCountIn applies the In predicate on the "skipped_count" field.
func SkippedCountIn(vs ...int64) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSkippedCount), v...))
	})
}

// SkippedCountNotIn applies the NotIn predicate on the "skipped_count" field.
func SkippedCountNotIn(vs ...int64) predicate.ProwJobs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProwJobs(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSkippedCount), v...))
	})
}

// SkippedCountGT applies the GT predicate on the "skipped_count" field.
func SkippedCountGT(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSkippedCount), v))
	})
}

// SkippedCountGTE applies the GTE predicate on the "skipped_count" field.
func SkippedCountGTE(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSkippedCount), v))
	})
}

// SkippedCountLT applies the LT predicate on the "skipped_count" field.
func SkippedCountLT(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSkippedCount), v))
	})
}

// SkippedCountLTE applies the LTE predicate on the "skipped_count" field.
func SkippedCountLTE(v int64) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSkippedCount), v))
	})
}

// HasProwJobs applies the HasEdge predicate on the "prow_jobs" edge.
func HasProwJobs() predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProwJobsTable, RepositoryFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProwJobsTable, ProwJobsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProwJobsWith applies the HasEdge predicate on the "prow_jobs" edge with a given conditions (other predicates).
func HasProwJobsWith(preds ...predicate.Repository) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProwJobsInverseTable, RepositoryFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProwJobsTable, ProwJobsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProwJobs) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProwJobs) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
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
func Not(p predicate.ProwJobs) predicate.ProwJobs {
	return predicate.ProwJobs(func(s *sql.Selector) {
		p(s.Not())
	})
}
