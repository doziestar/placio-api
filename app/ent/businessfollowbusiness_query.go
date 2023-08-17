// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"placio-app/ent/business"
	"placio-app/ent/businessfollowbusiness"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BusinessFollowBusinessQuery is the builder for querying BusinessFollowBusiness entities.
type BusinessFollowBusinessQuery struct {
	config
	ctx          *QueryContext
	order        []businessfollowbusiness.OrderOption
	inters       []Interceptor
	predicates   []predicate.BusinessFollowBusiness
	withFollower *BusinessQuery
	withFollowed *BusinessQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BusinessFollowBusinessQuery builder.
func (bfbq *BusinessFollowBusinessQuery) Where(ps ...predicate.BusinessFollowBusiness) *BusinessFollowBusinessQuery {
	bfbq.predicates = append(bfbq.predicates, ps...)
	return bfbq
}

// Limit the number of records to be returned by this query.
func (bfbq *BusinessFollowBusinessQuery) Limit(limit int) *BusinessFollowBusinessQuery {
	bfbq.ctx.Limit = &limit
	return bfbq
}

// Offset to start from.
func (bfbq *BusinessFollowBusinessQuery) Offset(offset int) *BusinessFollowBusinessQuery {
	bfbq.ctx.Offset = &offset
	return bfbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bfbq *BusinessFollowBusinessQuery) Unique(unique bool) *BusinessFollowBusinessQuery {
	bfbq.ctx.Unique = &unique
	return bfbq
}

// Order specifies how the records should be ordered.
func (bfbq *BusinessFollowBusinessQuery) Order(o ...businessfollowbusiness.OrderOption) *BusinessFollowBusinessQuery {
	bfbq.order = append(bfbq.order, o...)
	return bfbq
}

// QueryFollower chains the current query on the "follower" edge.
func (bfbq *BusinessFollowBusinessQuery) QueryFollower() *BusinessQuery {
	query := (&BusinessClient{config: bfbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bfbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bfbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(businessfollowbusiness.Table, businessfollowbusiness.FieldID, selector),
			sqlgraph.To(business.Table, business.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, businessfollowbusiness.FollowerTable, businessfollowbusiness.FollowerColumn),
		)
		fromU = sqlgraph.SetNeighbors(bfbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFollowed chains the current query on the "followed" edge.
func (bfbq *BusinessFollowBusinessQuery) QueryFollowed() *BusinessQuery {
	query := (&BusinessClient{config: bfbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bfbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bfbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(businessfollowbusiness.Table, businessfollowbusiness.FieldID, selector),
			sqlgraph.To(business.Table, business.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, businessfollowbusiness.FollowedTable, businessfollowbusiness.FollowedColumn),
		)
		fromU = sqlgraph.SetNeighbors(bfbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BusinessFollowBusiness entity from the query.
// Returns a *NotFoundError when no BusinessFollowBusiness was found.
func (bfbq *BusinessFollowBusinessQuery) First(ctx context.Context) (*BusinessFollowBusiness, error) {
	nodes, err := bfbq.Limit(1).All(setContextOp(ctx, bfbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{businessfollowbusiness.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bfbq *BusinessFollowBusinessQuery) FirstX(ctx context.Context) *BusinessFollowBusiness {
	node, err := bfbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BusinessFollowBusiness ID from the query.
// Returns a *NotFoundError when no BusinessFollowBusiness ID was found.
func (bfbq *BusinessFollowBusinessQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bfbq.Limit(1).IDs(setContextOp(ctx, bfbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{businessfollowbusiness.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bfbq *BusinessFollowBusinessQuery) FirstIDX(ctx context.Context) string {
	id, err := bfbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BusinessFollowBusiness entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BusinessFollowBusiness entity is found.
// Returns a *NotFoundError when no BusinessFollowBusiness entities are found.
func (bfbq *BusinessFollowBusinessQuery) Only(ctx context.Context) (*BusinessFollowBusiness, error) {
	nodes, err := bfbq.Limit(2).All(setContextOp(ctx, bfbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{businessfollowbusiness.Label}
	default:
		return nil, &NotSingularError{businessfollowbusiness.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bfbq *BusinessFollowBusinessQuery) OnlyX(ctx context.Context) *BusinessFollowBusiness {
	node, err := bfbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BusinessFollowBusiness ID in the query.
// Returns a *NotSingularError when more than one BusinessFollowBusiness ID is found.
// Returns a *NotFoundError when no entities are found.
func (bfbq *BusinessFollowBusinessQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bfbq.Limit(2).IDs(setContextOp(ctx, bfbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{businessfollowbusiness.Label}
	default:
		err = &NotSingularError{businessfollowbusiness.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bfbq *BusinessFollowBusinessQuery) OnlyIDX(ctx context.Context) string {
	id, err := bfbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BusinessFollowBusinesses.
func (bfbq *BusinessFollowBusinessQuery) All(ctx context.Context) ([]*BusinessFollowBusiness, error) {
	ctx = setContextOp(ctx, bfbq.ctx, "All")
	if err := bfbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BusinessFollowBusiness, *BusinessFollowBusinessQuery]()
	return withInterceptors[[]*BusinessFollowBusiness](ctx, bfbq, qr, bfbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bfbq *BusinessFollowBusinessQuery) AllX(ctx context.Context) []*BusinessFollowBusiness {
	nodes, err := bfbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BusinessFollowBusiness IDs.
func (bfbq *BusinessFollowBusinessQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bfbq.ctx.Unique == nil && bfbq.path != nil {
		bfbq.Unique(true)
	}
	ctx = setContextOp(ctx, bfbq.ctx, "IDs")
	if err = bfbq.Select(businessfollowbusiness.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bfbq *BusinessFollowBusinessQuery) IDsX(ctx context.Context) []string {
	ids, err := bfbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bfbq *BusinessFollowBusinessQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bfbq.ctx, "Count")
	if err := bfbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bfbq, querierCount[*BusinessFollowBusinessQuery](), bfbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bfbq *BusinessFollowBusinessQuery) CountX(ctx context.Context) int {
	count, err := bfbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bfbq *BusinessFollowBusinessQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bfbq.ctx, "Exist")
	switch _, err := bfbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bfbq *BusinessFollowBusinessQuery) ExistX(ctx context.Context) bool {
	exist, err := bfbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BusinessFollowBusinessQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bfbq *BusinessFollowBusinessQuery) Clone() *BusinessFollowBusinessQuery {
	if bfbq == nil {
		return nil
	}
	return &BusinessFollowBusinessQuery{
		config:       bfbq.config,
		ctx:          bfbq.ctx.Clone(),
		order:        append([]businessfollowbusiness.OrderOption{}, bfbq.order...),
		inters:       append([]Interceptor{}, bfbq.inters...),
		predicates:   append([]predicate.BusinessFollowBusiness{}, bfbq.predicates...),
		withFollower: bfbq.withFollower.Clone(),
		withFollowed: bfbq.withFollowed.Clone(),
		// clone intermediate query.
		sql:  bfbq.sql.Clone(),
		path: bfbq.path,
	}
}

// WithFollower tells the query-builder to eager-load the nodes that are connected to
// the "follower" edge. The optional arguments are used to configure the query builder of the edge.
func (bfbq *BusinessFollowBusinessQuery) WithFollower(opts ...func(*BusinessQuery)) *BusinessFollowBusinessQuery {
	query := (&BusinessClient{config: bfbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bfbq.withFollower = query
	return bfbq
}

// WithFollowed tells the query-builder to eager-load the nodes that are connected to
// the "followed" edge. The optional arguments are used to configure the query builder of the edge.
func (bfbq *BusinessFollowBusinessQuery) WithFollowed(opts ...func(*BusinessQuery)) *BusinessFollowBusinessQuery {
	query := (&BusinessClient{config: bfbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bfbq.withFollowed = query
	return bfbq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"CreatedAt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BusinessFollowBusiness.Query().
//		GroupBy(businessfollowbusiness.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bfbq *BusinessFollowBusinessQuery) GroupBy(field string, fields ...string) *BusinessFollowBusinessGroupBy {
	bfbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BusinessFollowBusinessGroupBy{build: bfbq}
	grbuild.flds = &bfbq.ctx.Fields
	grbuild.label = businessfollowbusiness.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"CreatedAt,omitempty"`
//	}
//
//	client.BusinessFollowBusiness.Query().
//		Select(businessfollowbusiness.FieldCreatedAt).
//		Scan(ctx, &v)
func (bfbq *BusinessFollowBusinessQuery) Select(fields ...string) *BusinessFollowBusinessSelect {
	bfbq.ctx.Fields = append(bfbq.ctx.Fields, fields...)
	sbuild := &BusinessFollowBusinessSelect{BusinessFollowBusinessQuery: bfbq}
	sbuild.label = businessfollowbusiness.Label
	sbuild.flds, sbuild.scan = &bfbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BusinessFollowBusinessSelect configured with the given aggregations.
func (bfbq *BusinessFollowBusinessQuery) Aggregate(fns ...AggregateFunc) *BusinessFollowBusinessSelect {
	return bfbq.Select().Aggregate(fns...)
}

func (bfbq *BusinessFollowBusinessQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bfbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bfbq); err != nil {
				return err
			}
		}
	}
	for _, f := range bfbq.ctx.Fields {
		if !businessfollowbusiness.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bfbq.path != nil {
		prev, err := bfbq.path(ctx)
		if err != nil {
			return err
		}
		bfbq.sql = prev
	}
	return nil
}

func (bfbq *BusinessFollowBusinessQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BusinessFollowBusiness, error) {
	var (
		nodes       = []*BusinessFollowBusiness{}
		withFKs     = bfbq.withFKs
		_spec       = bfbq.querySpec()
		loadedTypes = [2]bool{
			bfbq.withFollower != nil,
			bfbq.withFollowed != nil,
		}
	)
	if bfbq.withFollower != nil || bfbq.withFollowed != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, businessfollowbusiness.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BusinessFollowBusiness).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BusinessFollowBusiness{config: bfbq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bfbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bfbq.withFollower; query != nil {
		if err := bfbq.loadFollower(ctx, query, nodes, nil,
			func(n *BusinessFollowBusiness, e *Business) { n.Edges.Follower = e }); err != nil {
			return nil, err
		}
	}
	if query := bfbq.withFollowed; query != nil {
		if err := bfbq.loadFollowed(ctx, query, nodes, nil,
			func(n *BusinessFollowBusiness, e *Business) { n.Edges.Followed = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bfbq *BusinessFollowBusinessQuery) loadFollower(ctx context.Context, query *BusinessQuery, nodes []*BusinessFollowBusiness, init func(*BusinessFollowBusiness), assign func(*BusinessFollowBusiness, *Business)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*BusinessFollowBusiness)
	for i := range nodes {
		if nodes[i].business_followed_businesses == nil {
			continue
		}
		fk := *nodes[i].business_followed_businesses
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(business.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "business_followed_businesses" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bfbq *BusinessFollowBusinessQuery) loadFollowed(ctx context.Context, query *BusinessQuery, nodes []*BusinessFollowBusiness, init func(*BusinessFollowBusiness), assign func(*BusinessFollowBusiness, *Business)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*BusinessFollowBusiness)
	for i := range nodes {
		if nodes[i].business_follower_businesses == nil {
			continue
		}
		fk := *nodes[i].business_follower_businesses
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(business.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "business_follower_businesses" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (bfbq *BusinessFollowBusinessQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bfbq.querySpec()
	_spec.Node.Columns = bfbq.ctx.Fields
	if len(bfbq.ctx.Fields) > 0 {
		_spec.Unique = bfbq.ctx.Unique != nil && *bfbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bfbq.driver, _spec)
}

func (bfbq *BusinessFollowBusinessQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(businessfollowbusiness.Table, businessfollowbusiness.Columns, sqlgraph.NewFieldSpec(businessfollowbusiness.FieldID, field.TypeString))
	_spec.From = bfbq.sql
	if unique := bfbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bfbq.path != nil {
		_spec.Unique = true
	}
	if fields := bfbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, businessfollowbusiness.FieldID)
		for i := range fields {
			if fields[i] != businessfollowbusiness.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bfbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bfbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bfbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bfbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bfbq *BusinessFollowBusinessQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bfbq.driver.Dialect())
	t1 := builder.Table(businessfollowbusiness.Table)
	columns := bfbq.ctx.Fields
	if len(columns) == 0 {
		columns = businessfollowbusiness.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bfbq.sql != nil {
		selector = bfbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bfbq.ctx.Unique != nil && *bfbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bfbq.predicates {
		p(selector)
	}
	for _, p := range bfbq.order {
		p(selector)
	}
	if offset := bfbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bfbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BusinessFollowBusinessGroupBy is the group-by builder for BusinessFollowBusiness entities.
type BusinessFollowBusinessGroupBy struct {
	selector
	build *BusinessFollowBusinessQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bfbgb *BusinessFollowBusinessGroupBy) Aggregate(fns ...AggregateFunc) *BusinessFollowBusinessGroupBy {
	bfbgb.fns = append(bfbgb.fns, fns...)
	return bfbgb
}

// Scan applies the selector query and scans the result into the given value.
func (bfbgb *BusinessFollowBusinessGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bfbgb.build.ctx, "GroupBy")
	if err := bfbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BusinessFollowBusinessQuery, *BusinessFollowBusinessGroupBy](ctx, bfbgb.build, bfbgb, bfbgb.build.inters, v)
}

func (bfbgb *BusinessFollowBusinessGroupBy) sqlScan(ctx context.Context, root *BusinessFollowBusinessQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bfbgb.fns))
	for _, fn := range bfbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bfbgb.flds)+len(bfbgb.fns))
		for _, f := range *bfbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bfbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bfbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BusinessFollowBusinessSelect is the builder for selecting fields of BusinessFollowBusiness entities.
type BusinessFollowBusinessSelect struct {
	*BusinessFollowBusinessQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bfbs *BusinessFollowBusinessSelect) Aggregate(fns ...AggregateFunc) *BusinessFollowBusinessSelect {
	bfbs.fns = append(bfbs.fns, fns...)
	return bfbs
}

// Scan applies the selector query and scans the result into the given value.
func (bfbs *BusinessFollowBusinessSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bfbs.ctx, "Select")
	if err := bfbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BusinessFollowBusinessQuery, *BusinessFollowBusinessSelect](ctx, bfbs.BusinessFollowBusinessQuery, bfbs, bfbs.inters, v)
}

func (bfbs *BusinessFollowBusinessSelect) sqlScan(ctx context.Context, root *BusinessFollowBusinessQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bfbs.fns))
	for _, fn := range bfbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bfbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bfbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
