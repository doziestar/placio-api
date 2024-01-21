// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"placio-app/ent/fitness"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FitnessQuery is the builder for querying Fitness entities.
type FitnessQuery struct {
	config
	ctx        *QueryContext
	order      []fitness.OrderOption
	inters     []Interceptor
	predicates []predicate.Fitness
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FitnessQuery builder.
func (fq *FitnessQuery) Where(ps ...predicate.Fitness) *FitnessQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FitnessQuery) Limit(limit int) *FitnessQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FitnessQuery) Offset(offset int) *FitnessQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FitnessQuery) Unique(unique bool) *FitnessQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FitnessQuery) Order(o ...fitness.OrderOption) *FitnessQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// First returns the first Fitness entity from the query.
// Returns a *NotFoundError when no Fitness was found.
func (fq *FitnessQuery) First(ctx context.Context) (*Fitness, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fitness.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FitnessQuery) FirstX(ctx context.Context) *Fitness {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Fitness ID from the query.
// Returns a *NotFoundError when no Fitness ID was found.
func (fq *FitnessQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fitness.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FitnessQuery) FirstIDX(ctx context.Context) string {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Fitness entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Fitness entity is found.
// Returns a *NotFoundError when no Fitness entities are found.
func (fq *FitnessQuery) Only(ctx context.Context) (*Fitness, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fitness.Label}
	default:
		return nil, &NotSingularError{fitness.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FitnessQuery) OnlyX(ctx context.Context) *Fitness {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Fitness ID in the query.
// Returns a *NotSingularError when more than one Fitness ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FitnessQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fitness.Label}
	default:
		err = &NotSingularError{fitness.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FitnessQuery) OnlyIDX(ctx context.Context) string {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Fitnesses.
func (fq *FitnessQuery) All(ctx context.Context) ([]*Fitness, error) {
	ctx = setContextOp(ctx, fq.ctx, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Fitness, *FitnessQuery]()
	return withInterceptors[[]*Fitness](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FitnessQuery) AllX(ctx context.Context) []*Fitness {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Fitness IDs.
func (fq *FitnessQuery) IDs(ctx context.Context) (ids []string, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, "IDs")
	if err = fq.Select(fitness.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FitnessQuery) IDsX(ctx context.Context) []string {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FitnessQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FitnessQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FitnessQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FitnessQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, "Exist")
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FitnessQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FitnessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FitnessQuery) Clone() *FitnessQuery {
	if fq == nil {
		return nil
	}
	return &FitnessQuery{
		config:     fq.config,
		ctx:        fq.ctx.Clone(),
		order:      append([]fitness.OrderOption{}, fq.order...),
		inters:     append([]Interceptor{}, fq.inters...),
		predicates: append([]predicate.Fitness{}, fq.predicates...),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Fitness.Query().
//		GroupBy(fitness.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FitnessQuery) GroupBy(field string, fields ...string) *FitnessGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FitnessGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = fitness.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Fitness.Query().
//		Select(fitness.FieldName).
//		Scan(ctx, &v)
func (fq *FitnessQuery) Select(fields ...string) *FitnessSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FitnessSelect{FitnessQuery: fq}
	sbuild.label = fitness.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FitnessSelect configured with the given aggregations.
func (fq *FitnessQuery) Aggregate(fns ...AggregateFunc) *FitnessSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FitnessQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !fitness.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FitnessQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Fitness, error) {
	var (
		nodes = []*Fitness{}
		_spec = fq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Fitness).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Fitness{config: fq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fq *FitnessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FitnessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(fitness.Table, fitness.Columns, sqlgraph.NewFieldSpec(fitness.FieldID, field.TypeString))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fitness.FieldID)
		for i := range fields {
			if fields[i] != fitness.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FitnessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(fitness.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = fitness.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FitnessGroupBy is the group-by builder for Fitness entities.
type FitnessGroupBy struct {
	selector
	build *FitnessQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FitnessGroupBy) Aggregate(fns ...AggregateFunc) *FitnessGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FitnessGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FitnessQuery, *FitnessGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FitnessGroupBy) sqlScan(ctx context.Context, root *FitnessQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FitnessSelect is the builder for selecting fields of Fitness entities.
type FitnessSelect struct {
	*FitnessQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FitnessSelect) Aggregate(fns ...AggregateFunc) *FitnessSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FitnessSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FitnessQuery, *FitnessSelect](ctx, fs.FitnessQuery, fs, fs.inters, v)
}

func (fs *FitnessSelect) sqlScan(ctx context.Context, root *FitnessQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
