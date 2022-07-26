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
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/prowsuites"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/repository"
)

// ProwSuitesQuery is the builder for querying ProwSuites entities.
type ProwSuitesQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProwSuites
	// eager-loading edges.
	withProwSuites *RepositoryQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProwSuitesQuery builder.
func (psq *ProwSuitesQuery) Where(ps ...predicate.ProwSuites) *ProwSuitesQuery {
	psq.predicates = append(psq.predicates, ps...)
	return psq
}

// Limit adds a limit step to the query.
func (psq *ProwSuitesQuery) Limit(limit int) *ProwSuitesQuery {
	psq.limit = &limit
	return psq
}

// Offset adds an offset step to the query.
func (psq *ProwSuitesQuery) Offset(offset int) *ProwSuitesQuery {
	psq.offset = &offset
	return psq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psq *ProwSuitesQuery) Unique(unique bool) *ProwSuitesQuery {
	psq.unique = &unique
	return psq
}

// Order adds an order step to the query.
func (psq *ProwSuitesQuery) Order(o ...OrderFunc) *ProwSuitesQuery {
	psq.order = append(psq.order, o...)
	return psq
}

// QueryProwSuites chains the current query on the "prow_suites" edge.
func (psq *ProwSuitesQuery) QueryProwSuites() *RepositoryQuery {
	query := &RepositoryQuery{config: psq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(prowsuites.Table, prowsuites.FieldID, selector),
			sqlgraph.To(repository.Table, repository.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, prowsuites.ProwSuitesTable, prowsuites.ProwSuitesColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProwSuites entity from the query.
// Returns a *NotFoundError when no ProwSuites was found.
func (psq *ProwSuitesQuery) First(ctx context.Context) (*ProwSuites, error) {
	nodes, err := psq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{prowsuites.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psq *ProwSuitesQuery) FirstX(ctx context.Context) *ProwSuites {
	node, err := psq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProwSuites ID from the query.
// Returns a *NotFoundError when no ProwSuites ID was found.
func (psq *ProwSuitesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{prowsuites.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psq *ProwSuitesQuery) FirstIDX(ctx context.Context) int {
	id, err := psq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProwSuites entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ProwSuites entity is not found.
// Returns a *NotFoundError when no ProwSuites entities are found.
func (psq *ProwSuitesQuery) Only(ctx context.Context) (*ProwSuites, error) {
	nodes, err := psq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{prowsuites.Label}
	default:
		return nil, &NotSingularError{prowsuites.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psq *ProwSuitesQuery) OnlyX(ctx context.Context) *ProwSuites {
	node, err := psq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProwSuites ID in the query.
// Returns a *NotSingularError when exactly one ProwSuites ID is not found.
// Returns a *NotFoundError when no entities are found.
func (psq *ProwSuitesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = psq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{prowsuites.Label}
	default:
		err = &NotSingularError{prowsuites.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psq *ProwSuitesQuery) OnlyIDX(ctx context.Context) int {
	id, err := psq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProwSuitesSlice.
func (psq *ProwSuitesQuery) All(ctx context.Context) ([]*ProwSuites, error) {
	if err := psq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return psq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (psq *ProwSuitesQuery) AllX(ctx context.Context) []*ProwSuites {
	nodes, err := psq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProwSuites IDs.
func (psq *ProwSuitesQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := psq.Select(prowsuites.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psq *ProwSuitesQuery) IDsX(ctx context.Context) []int {
	ids, err := psq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psq *ProwSuitesQuery) Count(ctx context.Context) (int, error) {
	if err := psq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return psq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (psq *ProwSuitesQuery) CountX(ctx context.Context) int {
	count, err := psq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psq *ProwSuitesQuery) Exist(ctx context.Context) (bool, error) {
	if err := psq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return psq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (psq *ProwSuitesQuery) ExistX(ctx context.Context) bool {
	exist, err := psq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProwSuitesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psq *ProwSuitesQuery) Clone() *ProwSuitesQuery {
	if psq == nil {
		return nil
	}
	return &ProwSuitesQuery{
		config:         psq.config,
		limit:          psq.limit,
		offset:         psq.offset,
		order:          append([]OrderFunc{}, psq.order...),
		predicates:     append([]predicate.ProwSuites{}, psq.predicates...),
		withProwSuites: psq.withProwSuites.Clone(),
		// clone intermediate query.
		sql:  psq.sql.Clone(),
		path: psq.path,
	}
}

// WithProwSuites tells the query-builder to eager-load the nodes that are connected to
// the "prow_suites" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *ProwSuitesQuery) WithProwSuites(opts ...func(*RepositoryQuery)) *ProwSuitesQuery {
	query := &RepositoryQuery{config: psq.config}
	for _, opt := range opts {
		opt(query)
	}
	psq.withProwSuites = query
	return psq
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
//	client.ProwSuites.Query().
//		GroupBy(prowsuites.FieldJobID).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
//
func (psq *ProwSuitesQuery) GroupBy(field string, fields ...string) *ProwSuitesGroupBy {
	group := &ProwSuitesGroupBy{config: psq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return psq.sqlQuery(ctx), nil
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
//	client.ProwSuites.Query().
//		Select(prowsuites.FieldJobID).
//		Scan(ctx, &v)
//
func (psq *ProwSuitesQuery) Select(fields ...string) *ProwSuitesSelect {
	psq.fields = append(psq.fields, fields...)
	return &ProwSuitesSelect{ProwSuitesQuery: psq}
}

func (psq *ProwSuitesQuery) prepareQuery(ctx context.Context) error {
	for _, f := range psq.fields {
		if !prowsuites.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if psq.path != nil {
		prev, err := psq.path(ctx)
		if err != nil {
			return err
		}
		psq.sql = prev
	}
	return nil
}

func (psq *ProwSuitesQuery) sqlAll(ctx context.Context) ([]*ProwSuites, error) {
	var (
		nodes       = []*ProwSuites{}
		withFKs     = psq.withFKs
		_spec       = psq.querySpec()
		loadedTypes = [1]bool{
			psq.withProwSuites != nil,
		}
	)
	if psq.withProwSuites != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, prowsuites.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ProwSuites{config: psq.config}
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
	if err := sqlgraph.QueryNodes(ctx, psq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := psq.withProwSuites; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*ProwSuites)
		for i := range nodes {
			if nodes[i].repository_prow_suites == nil {
				continue
			}
			fk := *nodes[i].repository_prow_suites
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
				return nil, fmt.Errorf(`unexpected foreign-key "repository_prow_suites" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ProwSuites = n
			}
		}
	}

	return nodes, nil
}

func (psq *ProwSuitesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psq.querySpec()
	return sqlgraph.CountNodes(ctx, psq.driver, _spec)
}

func (psq *ProwSuitesQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := psq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("db: check existence: %w", err)
	}
	return n > 0, nil
}

func (psq *ProwSuitesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prowsuites.Table,
			Columns: prowsuites.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prowsuites.FieldID,
			},
		},
		From:   psq.sql,
		Unique: true,
	}
	if unique := psq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := psq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prowsuites.FieldID)
		for i := range fields {
			if fields[i] != prowsuites.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := psq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psq *ProwSuitesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psq.driver.Dialect())
	t1 := builder.Table(prowsuites.Table)
	columns := psq.fields
	if len(columns) == 0 {
		columns = prowsuites.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psq.sql != nil {
		selector = psq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range psq.predicates {
		p(selector)
	}
	for _, p := range psq.order {
		p(selector)
	}
	if offset := psq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProwSuitesGroupBy is the group-by builder for ProwSuites entities.
type ProwSuitesGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psgb *ProwSuitesGroupBy) Aggregate(fns ...AggregateFunc) *ProwSuitesGroupBy {
	psgb.fns = append(psgb.fns, fns...)
	return psgb
}

// Scan applies the group-by query and scans the result into the given value.
func (psgb *ProwSuitesGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := psgb.path(ctx)
	if err != nil {
		return err
	}
	psgb.sql = query
	return psgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (psgb *ProwSuitesGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := psgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (psgb *ProwSuitesGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(psgb.fields) > 1 {
		return nil, errors.New("db: ProwSuitesGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := psgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (psgb *ProwSuitesGroupBy) StringsX(ctx context.Context) []string {
	v, err := psgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (psgb *ProwSuitesGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = psgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowsuites.Label}
	default:
		err = fmt.Errorf("db: ProwSuitesGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (psgb *ProwSuitesGroupBy) StringX(ctx context.Context) string {
	v, err := psgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (psgb *ProwSuitesGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(psgb.fields) > 1 {
		return nil, errors.New("db: ProwSuitesGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := psgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (psgb *ProwSuitesGroupBy) IntsX(ctx context.Context) []int {
	v, err := psgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (psgb *ProwSuitesGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = psgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowsuites.Label}
	default:
		err = fmt.Errorf("db: ProwSuitesGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (psgb *ProwSuitesGroupBy) IntX(ctx context.Context) int {
	v, err := psgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (psgb *ProwSuitesGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(psgb.fields) > 1 {
		return nil, errors.New("db: ProwSuitesGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := psgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (psgb *ProwSuitesGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := psgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (psgb *ProwSuitesGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = psgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowsuites.Label}
	default:
		err = fmt.Errorf("db: ProwSuitesGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (psgb *ProwSuitesGroupBy) Float64X(ctx context.Context) float64 {
	v, err := psgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (psgb *ProwSuitesGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(psgb.fields) > 1 {
		return nil, errors.New("db: ProwSuitesGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := psgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (psgb *ProwSuitesGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := psgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (psgb *ProwSuitesGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = psgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowsuites.Label}
	default:
		err = fmt.Errorf("db: ProwSuitesGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (psgb *ProwSuitesGroupBy) BoolX(ctx context.Context) bool {
	v, err := psgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (psgb *ProwSuitesGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range psgb.fields {
		if !prowsuites.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := psgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (psgb *ProwSuitesGroupBy) sqlQuery() *sql.Selector {
	selector := psgb.sql.Select()
	aggregation := make([]string, 0, len(psgb.fns))
	for _, fn := range psgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(psgb.fields)+len(psgb.fns))
		for _, f := range psgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(psgb.fields...)...)
}

// ProwSuitesSelect is the builder for selecting fields of ProwSuites entities.
type ProwSuitesSelect struct {
	*ProwSuitesQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pss *ProwSuitesSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pss.prepareQuery(ctx); err != nil {
		return err
	}
	pss.sql = pss.ProwSuitesQuery.sqlQuery(ctx)
	return pss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pss *ProwSuitesSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (pss *ProwSuitesSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pss.fields) > 1 {
		return nil, errors.New("db: ProwSuitesSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pss *ProwSuitesSelect) StringsX(ctx context.Context) []string {
	v, err := pss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (pss *ProwSuitesSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowsuites.Label}
	default:
		err = fmt.Errorf("db: ProwSuitesSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pss *ProwSuitesSelect) StringX(ctx context.Context) string {
	v, err := pss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (pss *ProwSuitesSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pss.fields) > 1 {
		return nil, errors.New("db: ProwSuitesSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pss *ProwSuitesSelect) IntsX(ctx context.Context) []int {
	v, err := pss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (pss *ProwSuitesSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowsuites.Label}
	default:
		err = fmt.Errorf("db: ProwSuitesSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pss *ProwSuitesSelect) IntX(ctx context.Context) int {
	v, err := pss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (pss *ProwSuitesSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pss.fields) > 1 {
		return nil, errors.New("db: ProwSuitesSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pss *ProwSuitesSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (pss *ProwSuitesSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowsuites.Label}
	default:
		err = fmt.Errorf("db: ProwSuitesSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pss *ProwSuitesSelect) Float64X(ctx context.Context) float64 {
	v, err := pss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (pss *ProwSuitesSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pss.fields) > 1 {
		return nil, errors.New("db: ProwSuitesSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pss *ProwSuitesSelect) BoolsX(ctx context.Context) []bool {
	v, err := pss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (pss *ProwSuitesSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{prowsuites.Label}
	default:
		err = fmt.Errorf("db: ProwSuitesSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pss *ProwSuitesSelect) BoolX(ctx context.Context) bool {
	v, err := pss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pss *ProwSuitesSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pss.sql.Query()
	if err := pss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
