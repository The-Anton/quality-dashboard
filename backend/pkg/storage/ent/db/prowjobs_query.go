// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/prowjobs"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
)

// ProwJobsQuery is the builder for querying ProwJobs entities.
type ProwJobsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProwJobs
	// eager-loading edges.
	withProwJobs *RepositoryQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProwJobsQuery builder.
func (pjq *ProwJobsQuery) Where(ps ...predicate.ProwJobs) *ProwJobsQuery {
	pjq.predicates = append(pjq.predicates, ps...)
	return pjq
}

// Limit adds a limit step to the query.
func (pjq *ProwJobsQuery) Limit(limit int) *ProwJobsQuery {
	pjq.limit = &limit
	return pjq
}

// Offset adds an offset step to the query.
func (pjq *ProwJobsQuery) Offset(offset int) *ProwJobsQuery {
	pjq.offset = &offset
	return pjq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pjq *ProwJobsQuery) Unique(unique bool) *ProwJobsQuery {
	pjq.unique = &unique
	return pjq
}

// Order adds an order step to the query.
func (pjq *ProwJobsQuery) Order(o ...OrderFunc) *ProwJobsQuery {
	pjq.order = append(pjq.order, o...)
	return pjq
}

// QueryProwJobs chains the current query on the "prow_jobs" edge.
func (pjq *ProwJobsQuery) QueryProwJobs() *RepositoryQuery {
	query := &RepositoryQuery{config: pjq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pjq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pjq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(prowjobs.Table, prowjobs.FieldID, selector),
			sqlgraph.To(repository.Table, repository.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, prowjobs.ProwJobsTable, prowjobs.ProwJobsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pjq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProwJobs entity from the query.
// Returns a *NotFoundError when no ProwJobs was found.
func (pjq *ProwJobsQuery) First(ctx context.Context) (*ProwJobs, error) {
	nodes, err := pjq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{prowjobs.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pjq *ProwJobsQuery) FirstX(ctx context.Context) *ProwJobs {
	node, err := pjq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProwJobs ID from the query.
// Returns a *NotFoundError when no ProwJobs ID was found.
func (pjq *ProwJobsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pjq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{prowjobs.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pjq *ProwJobsQuery) FirstIDX(ctx context.Context) int {
	id, err := pjq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProwJobs entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProwJobs entity is not found.
// Returns a *NotFoundError when no ProwJobs entities are found.
func (pjq *ProwJobsQuery) Only(ctx context.Context) (*ProwJobs, error) {
	nodes, err := pjq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{prowjobs.Label}
	default:
		return nil, &NotSingularError{prowjobs.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pjq *ProwJobsQuery) OnlyX(ctx context.Context) *ProwJobs {
	node, err := pjq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProwJobs ID in the query.
// Returns a *NotSingularError when exactly one ProwJobs ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pjq *ProwJobsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pjq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{prowjobs.Label}
	default:
		err = &NotSingularError{prowjobs.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pjq *ProwJobsQuery) OnlyIDX(ctx context.Context) int {
	id, err := pjq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProwJobsSlice.
func (pjq *ProwJobsQuery) All(ctx context.Context) ([]*ProwJobs, error) {
	if err := pjq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pjq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pjq *ProwJobsQuery) AllX(ctx context.Context) []*ProwJobs {
	nodes, err := pjq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProwJobs IDs.
func (pjq *ProwJobsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pjq.Select(prowjobs.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pjq *ProwJobsQuery) IDsX(ctx context.Context) []int {
	ids, err := pjq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pjq *ProwJobsQuery) Count(ctx context.Context) (int, error) {
	if err := pjq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pjq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pjq *ProwJobsQuery) CountX(ctx context.Context) int {
	count, err := pjq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pjq *ProwJobsQuery) Exist(ctx context.Context) (bool, error) {
	if err := pjq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pjq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pjq *ProwJobsQuery) ExistX(ctx context.Context) bool {
	exist, err := pjq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProwJobsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pjq *ProwJobsQuery) Clone() *ProwJobsQuery {
	if pjq == nil {
		return nil
	}
	return &ProwJobsQuery{
		config:       pjq.config,
		limit:        pjq.limit,
		offset:       pjq.offset,
		order:        append([]OrderFunc{}, pjq.order...),
		predicates:   append([]predicate.ProwJobs{}, pjq.predicates...),
		withProwJobs: pjq.withProwJobs.Clone(),
		// clone intermediate query.
		sql:  pjq.sql.Clone(),
		path: pjq.path,
	}
}

// WithProwJobs tells the query-builder to eager-load the nodes that are connected to
// the "prow_jobs" edge. The optional arguments are used to configure the query builder of the edge.
func (pjq *ProwJobsQuery) WithProwJobs(opts ...func(*RepositoryQuery)) *ProwJobsQuery {
	query := &RepositoryQuery{config: pjq.config}
	for _, opt := range opts {
		opt(query)
	}
	pjq.withProwJobs = query
	return pjq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		JobID string `json:"job_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProwJobs.Query().
//		GroupBy(prowjobs.FieldJobID).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
//
func (pjq *ProwJobsQuery) GroupBy(field string, fields ...string) *ProwJobsGroupBy {
	group := &ProwJobsGroupBy{config: pjq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pjq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pjq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		JobID string `json:"job_id,omitempty"`
//	}
//
//	client.ProwJobs.Query().
//		Select(prowjobs.FieldJobID).
//		Scan(ctx, &v)
//
func (pjq *ProwJobsQuery) Select(fields ...string) *ProwJobsSelect {
	pjq.fields = append(pjq.fields, fields...)
	return &ProwJobsSelect{ProwJobsQuery: pjq}
}

func (pjq *ProwJobsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pjq.fields {
		if !prowjobs.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if pjq.path != nil {
		prev, err := pjq.path(ctx)
		if err != nil {
			return err
		}
		pjq.sql = prev
	}
	return nil
}

func (pjq *ProwJobsQuery) sqlAll(ctx context.Context) ([]*ProwJobs, error) {
	var (
		nodes       = []*ProwJobs{}
		withFKs     = pjq.withFKs
		_spec       = pjq.querySpec()
		loadedTypes = [1]bool{
			pjq.withProwJobs != nil,
		}
	)
	if pjq.withProwJobs != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, prowjobs.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProwJobs{config: pjq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("db: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pjq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pjq.withProwJobs; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*ProwJobs)
		for i := range nodes {
			if nodes[i].repository_prow_jobs == nil {
				continue
			}
			fk := *nodes[i].repository_prow_jobs
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(repository.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "repository_prow_jobs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ProwJobs = n
			}
		}
	}

	return nodes, nil
}

func (pjq *ProwJobsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pjq.querySpec()
	return sqlgraph.CountNodes(ctx, pjq.driver, _spec)
}

func (pjq *ProwJobsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pjq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("db: check existence: %w", err)
	}
	return n > 0, nil
}

func (pjq *ProwJobsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prowjobs.Table,
			Columns: prowjobs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prowjobs.FieldID,
			},
		},
		From:   pjq.sql,
		Unique: true,
	}
	if unique := pjq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pjq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prowjobs.FieldID)
		for i := range fields {
			if fields[i] != prowjobs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pjq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pjq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pjq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pjq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pjq *ProwJobsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pjq.driver.Dialect())
	t1 := builder.Table(prowjobs.Table)
	columns := pjq.fields
	if len(columns) == 0 {
		columns = prowjobs.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pjq.sql != nil {
		selector = pjq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range pjq.predicates {
		p(selector)
	}
	for _, p := range pjq.order {
		p(selector)
	}
	if offset := pjq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pjq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProwJobsGroupBy is the group-by builder for ProwJobs entities.
type ProwJobsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pjgb *ProwJobsGroupBy) Aggregate(fns ...AggregateFunc) *ProwJobsGroupBy {
	pjgb.fns = append(pjgb.fns, fns...)
	return pjgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pjgb *ProwJobsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pjgb.path(ctx)
	if err != nil {
		return err
	}
	pjgb.sql = query
	return pjgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pjgb *ProwJobsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pjgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pjgb *ProwJobsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pjgb.fields) > 1 {
		return nil, errors.New("db: ProwJobsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pjgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pjgb *ProwJobsGroupBy) StringsX(ctx context.Context) []string {
	v, err := pjgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pjgb *ProwJobsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pjgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowjobs.Label}
	default:
		err = fmt.Errorf("db: ProwJobsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pjgb *ProwJobsGroupBy) StringX(ctx context.Context) string {
	v, err := pjgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pjgb *ProwJobsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pjgb.fields) > 1 {
		return nil, errors.New("db: ProwJobsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pjgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pjgb *ProwJobsGroupBy) IntsX(ctx context.Context) []int {
	v, err := pjgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pjgb *ProwJobsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pjgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowjobs.Label}
	default:
		err = fmt.Errorf("db: ProwJobsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pjgb *ProwJobsGroupBy) IntX(ctx context.Context) int {
	v, err := pjgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pjgb *ProwJobsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pjgb.fields) > 1 {
		return nil, errors.New("db: ProwJobsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pjgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pjgb *ProwJobsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pjgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pjgb *ProwJobsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pjgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowjobs.Label}
	default:
		err = fmt.Errorf("db: ProwJobsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pjgb *ProwJobsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pjgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pjgb *ProwJobsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pjgb.fields) > 1 {
		return nil, errors.New("db: ProwJobsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pjgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pjgb *ProwJobsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pjgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pjgb *ProwJobsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pjgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowjobs.Label}
	default:
		err = fmt.Errorf("db: ProwJobsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pjgb *ProwJobsGroupBy) BoolX(ctx context.Context) bool {
	v, err := pjgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pjgb *ProwJobsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pjgb.fields {
		if !prowjobs.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pjgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pjgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pjgb *ProwJobsGroupBy) sqlQuery() *sql.Selector {
	selector := pjgb.sql.Select()
	aggregation := make([]string, 0, len(pjgb.fns))
	for _, fn := range pjgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pjgb.fields)+len(pjgb.fns))
		for _, f := range pjgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pjgb.fields...)...)
}

// ProwJobsSelect is the builder for selecting fields of ProwJobs entities.
type ProwJobsSelect struct {
	*ProwJobsQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pjs *ProwJobsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pjs.prepareQuery(ctx); err != nil {
		return err
	}
	pjs.sql = pjs.ProwJobsQuery.sqlQuery(ctx)
	return pjs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pjs *ProwJobsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pjs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pjs *ProwJobsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pjs.fields) > 1 {
		return nil, errors.New("db: ProwJobsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pjs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pjs *ProwJobsSelect) StringsX(ctx context.Context) []string {
	v, err := pjs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pjs *ProwJobsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pjs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowjobs.Label}
	default:
		err = fmt.Errorf("db: ProwJobsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pjs *ProwJobsSelect) StringX(ctx context.Context) string {
	v, err := pjs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pjs *ProwJobsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pjs.fields) > 1 {
		return nil, errors.New("db: ProwJobsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pjs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pjs *ProwJobsSelect) IntsX(ctx context.Context) []int {
	v, err := pjs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pjs *ProwJobsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pjs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowjobs.Label}
	default:
		err = fmt.Errorf("db: ProwJobsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pjs *ProwJobsSelect) IntX(ctx context.Context) int {
	v, err := pjs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pjs *ProwJobsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pjs.fields) > 1 {
		return nil, errors.New("db: ProwJobsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pjs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pjs *ProwJobsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pjs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pjs *ProwJobsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pjs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowjobs.Label}
	default:
		err = fmt.Errorf("db: ProwJobsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pjs *ProwJobsSelect) Float64X(ctx context.Context) float64 {
	v, err := pjs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pjs *ProwJobsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pjs.fields) > 1 {
		return nil, errors.New("db: ProwJobsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pjs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pjs *ProwJobsSelect) BoolsX(ctx context.Context) []bool {
	v, err := pjs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pjs *ProwJobsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pjs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowjobs.Label}
	default:
		err = fmt.Errorf("db: ProwJobsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pjs *ProwJobsSelect) BoolX(ctx context.Context) bool {
	v, err := pjs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pjs *ProwJobsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pjs.sql.Query()
	if err := pjs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
