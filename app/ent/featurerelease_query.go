// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"placio-app/ent/featurerelease"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeatureReleaseQuery is the builder for querying FeatureRelease entities.
type FeatureReleaseQuery struct {
	config
	ctx        *QueryContext
	order      []featurerelease.OrderOption
	inters     []Interceptor
	predicates []predicate.FeatureRelease
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FeatureReleaseQuery builder.
func (frq *FeatureReleaseQuery) Where(ps ...predicate.FeatureRelease) *FeatureReleaseQuery {
	frq.predicates = append(frq.predicates, ps...)
	return frq
}

// Limit the number of records to be returned by this query.
func (frq *FeatureReleaseQuery) Limit(limit int) *FeatureReleaseQuery {
	frq.ctx.Limit = &limit
	return frq
}

// Offset to start from.
func (frq *FeatureReleaseQuery) Offset(offset int) *FeatureReleaseQuery {
	frq.ctx.Offset = &offset
	return frq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (frq *FeatureReleaseQuery) Unique(unique bool) *FeatureReleaseQuery {
	frq.ctx.Unique = &unique
	return frq
}

// Order specifies how the records should be ordered.
func (frq *FeatureReleaseQuery) Order(o ...featurerelease.OrderOption) *FeatureReleaseQuery {
	frq.order = append(frq.order, o...)
	return frq
}

// First returns the first FeatureRelease entity from the query.
// Returns a *NotFoundError when no FeatureRelease was found.
func (frq *FeatureReleaseQuery) First(ctx context.Context) (*FeatureRelease, error) {
	nodes, err := frq.Limit(1).All(setContextOp(ctx, frq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{featurerelease.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (frq *FeatureReleaseQuery) FirstX(ctx context.Context) *FeatureRelease {
	node, err := frq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FeatureRelease ID from the query.
// Returns a *NotFoundError when no FeatureRelease ID was found.
func (frq *FeatureReleaseQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = frq.Limit(1).IDs(setContextOp(ctx, frq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{featurerelease.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (frq *FeatureReleaseQuery) FirstIDX(ctx context.Context) string {
	id, err := frq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FeatureRelease entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FeatureRelease entity is found.
// Returns a *NotFoundError when no FeatureRelease entities are found.
func (frq *FeatureReleaseQuery) Only(ctx context.Context) (*FeatureRelease, error) {
	nodes, err := frq.Limit(2).All(setContextOp(ctx, frq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{featurerelease.Label}
	default:
		return nil, &NotSingularError{featurerelease.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (frq *FeatureReleaseQuery) OnlyX(ctx context.Context) *FeatureRelease {
	node, err := frq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FeatureRelease ID in the query.
// Returns a *NotSingularError when more than one FeatureRelease ID is found.
// Returns a *NotFoundError when no entities are found.
func (frq *FeatureReleaseQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = frq.Limit(2).IDs(setContextOp(ctx, frq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{featurerelease.Label}
	default:
		err = &NotSingularError{featurerelease.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (frq *FeatureReleaseQuery) OnlyIDX(ctx context.Context) string {
	id, err := frq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FeatureReleases.
func (frq *FeatureReleaseQuery) All(ctx context.Context) ([]*FeatureRelease, error) {
	ctx = setContextOp(ctx, frq.ctx, "All")
	if err := frq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FeatureRelease, *FeatureReleaseQuery]()
	return withInterceptors[[]*FeatureRelease](ctx, frq, qr, frq.inters)
}

// AllX is like All, but panics if an error occurs.
func (frq *FeatureReleaseQuery) AllX(ctx context.Context) []*FeatureRelease {
	nodes, err := frq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FeatureRelease IDs.
func (frq *FeatureReleaseQuery) IDs(ctx context.Context) (ids []string, err error) {
	if frq.ctx.Unique == nil && frq.path != nil {
		frq.Unique(true)
	}
	ctx = setContextOp(ctx, frq.ctx, "IDs")
	if err = frq.Select(featurerelease.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (frq *FeatureReleaseQuery) IDsX(ctx context.Context) []string {
	ids, err := frq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (frq *FeatureReleaseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, frq.ctx, "Count")
	if err := frq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, frq, querierCount[*FeatureReleaseQuery](), frq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (frq *FeatureReleaseQuery) CountX(ctx context.Context) int {
	count, err := frq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (frq *FeatureReleaseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, frq.ctx, "Exist")
	switch _, err := frq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (frq *FeatureReleaseQuery) ExistX(ctx context.Context) bool {
	exist, err := frq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FeatureReleaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (frq *FeatureReleaseQuery) Clone() *FeatureReleaseQuery {
	if frq == nil {
		return nil
	}
	return &FeatureReleaseQuery{
		config:     frq.config,
		ctx:        frq.ctx.Clone(),
		order:      append([]featurerelease.OrderOption{}, frq.order...),
		inters:     append([]Interceptor{}, frq.inters...),
		predicates: append([]predicate.FeatureRelease{}, frq.predicates...),
		// clone intermediate query.
		sql:  frq.sql.Clone(),
		path: frq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FeatureName string `json:"feature_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FeatureRelease.Query().
//		GroupBy(featurerelease.FieldFeatureName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (frq *FeatureReleaseQuery) GroupBy(field string, fields ...string) *FeatureReleaseGroupBy {
	frq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FeatureReleaseGroupBy{build: frq}
	grbuild.flds = &frq.ctx.Fields
	grbuild.label = featurerelease.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FeatureName string `json:"feature_name,omitempty"`
//	}
//
//	client.FeatureRelease.Query().
//		Select(featurerelease.FieldFeatureName).
//		Scan(ctx, &v)
func (frq *FeatureReleaseQuery) Select(fields ...string) *FeatureReleaseSelect {
	frq.ctx.Fields = append(frq.ctx.Fields, fields...)
	sbuild := &FeatureReleaseSelect{FeatureReleaseQuery: frq}
	sbuild.label = featurerelease.Label
	sbuild.flds, sbuild.scan = &frq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FeatureReleaseSelect configured with the given aggregations.
func (frq *FeatureReleaseQuery) Aggregate(fns ...AggregateFunc) *FeatureReleaseSelect {
	return frq.Select().Aggregate(fns...)
}

func (frq *FeatureReleaseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range frq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, frq); err != nil {
				return err
			}
		}
	}
	for _, f := range frq.ctx.Fields {
		if !featurerelease.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if frq.path != nil {
		prev, err := frq.path(ctx)
		if err != nil {
			return err
		}
		frq.sql = prev
	}
	return nil
}

func (frq *FeatureReleaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FeatureRelease, error) {
	var (
		nodes = []*FeatureRelease{}
		_spec = frq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FeatureRelease).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FeatureRelease{config: frq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, frq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (frq *FeatureReleaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := frq.querySpec()
	_spec.Node.Columns = frq.ctx.Fields
	if len(frq.ctx.Fields) > 0 {
		_spec.Unique = frq.ctx.Unique != nil && *frq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, frq.driver, _spec)
}

func (frq *FeatureReleaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(featurerelease.Table, featurerelease.Columns, sqlgraph.NewFieldSpec(featurerelease.FieldID, field.TypeString))
	_spec.From = frq.sql
	if unique := frq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if frq.path != nil {
		_spec.Unique = true
	}
	if fields := frq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, featurerelease.FieldID)
		for i := range fields {
			if fields[i] != featurerelease.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := frq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := frq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := frq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := frq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (frq *FeatureReleaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(frq.driver.Dialect())
	t1 := builder.Table(featurerelease.Table)
	columns := frq.ctx.Fields
	if len(columns) == 0 {
		columns = featurerelease.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if frq.sql != nil {
		selector = frq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if frq.ctx.Unique != nil && *frq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range frq.predicates {
		p(selector)
	}
	for _, p := range frq.order {
		p(selector)
	}
	if offset := frq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := frq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FeatureReleaseGroupBy is the group-by builder for FeatureRelease entities.
type FeatureReleaseGroupBy struct {
	selector
	build *FeatureReleaseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (frgb *FeatureReleaseGroupBy) Aggregate(fns ...AggregateFunc) *FeatureReleaseGroupBy {
	frgb.fns = append(frgb.fns, fns...)
	return frgb
}

// Scan applies the selector query and scans the result into the given value.
func (frgb *FeatureReleaseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, frgb.build.ctx, "GroupBy")
	if err := frgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeatureReleaseQuery, *FeatureReleaseGroupBy](ctx, frgb.build, frgb, frgb.build.inters, v)
}

func (frgb *FeatureReleaseGroupBy) sqlScan(ctx context.Context, root *FeatureReleaseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(frgb.fns))
	for _, fn := range frgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*frgb.flds)+len(frgb.fns))
		for _, f := range *frgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*frgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := frgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FeatureReleaseSelect is the builder for selecting fields of FeatureRelease entities.
type FeatureReleaseSelect struct {
	*FeatureReleaseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (frs *FeatureReleaseSelect) Aggregate(fns ...AggregateFunc) *FeatureReleaseSelect {
	frs.fns = append(frs.fns, fns...)
	return frs
}

// Scan applies the selector query and scans the result into the given value.
func (frs *FeatureReleaseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, frs.ctx, "Select")
	if err := frs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeatureReleaseQuery, *FeatureReleaseSelect](ctx, frs.FeatureReleaseQuery, frs, frs.inters, v)
}

func (frs *FeatureReleaseSelect) sqlScan(ctx context.Context, root *FeatureReleaseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(frs.fns))
	for _, fn := range frs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*frs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := frs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
