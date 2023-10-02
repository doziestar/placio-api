// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"placio_api/place"
	"placio_api/predicate"
	"placio_api/user"
	"placio_api/userlikeplace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserLikePlaceQuery is the builder for querying UserLikePlace entities.
type UserLikePlaceQuery struct {
	config
	ctx        *QueryContext
	order      []userlikeplace.OrderOption
	inters     []Interceptor
	predicates []predicate.UserLikePlace
	withUser   *UserQuery
	withPlace  *PlaceQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserLikePlaceQuery builder.
func (ulpq *UserLikePlaceQuery) Where(ps ...predicate.UserLikePlace) *UserLikePlaceQuery {
	ulpq.predicates = append(ulpq.predicates, ps...)
	return ulpq
}

// Limit the number of records to be returned by this query.
func (ulpq *UserLikePlaceQuery) Limit(limit int) *UserLikePlaceQuery {
	ulpq.ctx.Limit = &limit
	return ulpq
}

// Offset to start from.
func (ulpq *UserLikePlaceQuery) Offset(offset int) *UserLikePlaceQuery {
	ulpq.ctx.Offset = &offset
	return ulpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ulpq *UserLikePlaceQuery) Unique(unique bool) *UserLikePlaceQuery {
	ulpq.ctx.Unique = &unique
	return ulpq
}

// Order specifies how the records should be ordered.
func (ulpq *UserLikePlaceQuery) Order(o ...userlikeplace.OrderOption) *UserLikePlaceQuery {
	ulpq.order = append(ulpq.order, o...)
	return ulpq
}

// QueryUser chains the current query on the "user" edge.
func (ulpq *UserLikePlaceQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ulpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ulpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ulpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userlikeplace.Table, userlikeplace.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userlikeplace.UserTable, userlikeplace.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ulpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlace chains the current query on the "place" edge.
func (ulpq *UserLikePlaceQuery) QueryPlace() *PlaceQuery {
	query := (&PlaceClient{config: ulpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ulpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ulpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userlikeplace.Table, userlikeplace.FieldID, selector),
			sqlgraph.To(place.Table, place.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userlikeplace.PlaceTable, userlikeplace.PlaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(ulpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserLikePlace entity from the query.
// Returns a *NotFoundError when no UserLikePlace was found.
func (ulpq *UserLikePlaceQuery) First(ctx context.Context) (*UserLikePlace, error) {
	nodes, err := ulpq.Limit(1).All(setContextOp(ctx, ulpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userlikeplace.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ulpq *UserLikePlaceQuery) FirstX(ctx context.Context) *UserLikePlace {
	node, err := ulpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserLikePlace ID from the query.
// Returns a *NotFoundError when no UserLikePlace ID was found.
func (ulpq *UserLikePlaceQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ulpq.Limit(1).IDs(setContextOp(ctx, ulpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userlikeplace.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ulpq *UserLikePlaceQuery) FirstIDX(ctx context.Context) string {
	id, err := ulpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserLikePlace entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserLikePlace entity is found.
// Returns a *NotFoundError when no UserLikePlace entities are found.
func (ulpq *UserLikePlaceQuery) Only(ctx context.Context) (*UserLikePlace, error) {
	nodes, err := ulpq.Limit(2).All(setContextOp(ctx, ulpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userlikeplace.Label}
	default:
		return nil, &NotSingularError{userlikeplace.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ulpq *UserLikePlaceQuery) OnlyX(ctx context.Context) *UserLikePlace {
	node, err := ulpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserLikePlace ID in the query.
// Returns a *NotSingularError when more than one UserLikePlace ID is found.
// Returns a *NotFoundError when no entities are found.
func (ulpq *UserLikePlaceQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ulpq.Limit(2).IDs(setContextOp(ctx, ulpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userlikeplace.Label}
	default:
		err = &NotSingularError{userlikeplace.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ulpq *UserLikePlaceQuery) OnlyIDX(ctx context.Context) string {
	id, err := ulpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserLikePlaces.
func (ulpq *UserLikePlaceQuery) All(ctx context.Context) ([]*UserLikePlace, error) {
	ctx = setContextOp(ctx, ulpq.ctx, "All")
	if err := ulpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserLikePlace, *UserLikePlaceQuery]()
	return withInterceptors[[]*UserLikePlace](ctx, ulpq, qr, ulpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ulpq *UserLikePlaceQuery) AllX(ctx context.Context) []*UserLikePlace {
	nodes, err := ulpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserLikePlace IDs.
func (ulpq *UserLikePlaceQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ulpq.ctx.Unique == nil && ulpq.path != nil {
		ulpq.Unique(true)
	}
	ctx = setContextOp(ctx, ulpq.ctx, "IDs")
	if err = ulpq.Select(userlikeplace.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ulpq *UserLikePlaceQuery) IDsX(ctx context.Context) []string {
	ids, err := ulpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ulpq *UserLikePlaceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ulpq.ctx, "Count")
	if err := ulpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ulpq, querierCount[*UserLikePlaceQuery](), ulpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ulpq *UserLikePlaceQuery) CountX(ctx context.Context) int {
	count, err := ulpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ulpq *UserLikePlaceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ulpq.ctx, "Exist")
	switch _, err := ulpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("placio_api: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ulpq *UserLikePlaceQuery) ExistX(ctx context.Context) bool {
	exist, err := ulpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserLikePlaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ulpq *UserLikePlaceQuery) Clone() *UserLikePlaceQuery {
	if ulpq == nil {
		return nil
	}
	return &UserLikePlaceQuery{
		config:     ulpq.config,
		ctx:        ulpq.ctx.Clone(),
		order:      append([]userlikeplace.OrderOption{}, ulpq.order...),
		inters:     append([]Interceptor{}, ulpq.inters...),
		predicates: append([]predicate.UserLikePlace{}, ulpq.predicates...),
		withUser:   ulpq.withUser.Clone(),
		withPlace:  ulpq.withPlace.Clone(),
		// clone intermediate query.
		sql:  ulpq.sql.Clone(),
		path: ulpq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ulpq *UserLikePlaceQuery) WithUser(opts ...func(*UserQuery)) *UserLikePlaceQuery {
	query := (&UserClient{config: ulpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ulpq.withUser = query
	return ulpq
}

// WithPlace tells the query-builder to eager-load the nodes that are connected to
// the "place" edge. The optional arguments are used to configure the query builder of the edge.
func (ulpq *UserLikePlaceQuery) WithPlace(opts ...func(*PlaceQuery)) *UserLikePlaceQuery {
	query := (&PlaceClient{config: ulpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ulpq.withPlace = query
	return ulpq
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
//	client.UserLikePlace.Query().
//		GroupBy(userlikeplace.FieldCreatedAt).
//		Aggregate(placio_api.Count()).
//		Scan(ctx, &v)
func (ulpq *UserLikePlaceQuery) GroupBy(field string, fields ...string) *UserLikePlaceGroupBy {
	ulpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserLikePlaceGroupBy{build: ulpq}
	grbuild.flds = &ulpq.ctx.Fields
	grbuild.label = userlikeplace.Label
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
//	client.UserLikePlace.Query().
//		Select(userlikeplace.FieldCreatedAt).
//		Scan(ctx, &v)
func (ulpq *UserLikePlaceQuery) Select(fields ...string) *UserLikePlaceSelect {
	ulpq.ctx.Fields = append(ulpq.ctx.Fields, fields...)
	sbuild := &UserLikePlaceSelect{UserLikePlaceQuery: ulpq}
	sbuild.label = userlikeplace.Label
	sbuild.flds, sbuild.scan = &ulpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserLikePlaceSelect configured with the given aggregations.
func (ulpq *UserLikePlaceQuery) Aggregate(fns ...AggregateFunc) *UserLikePlaceSelect {
	return ulpq.Select().Aggregate(fns...)
}

func (ulpq *UserLikePlaceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ulpq.inters {
		if inter == nil {
			return fmt.Errorf("placio_api: uninitialized interceptor (forgotten import placio_api/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ulpq); err != nil {
				return err
			}
		}
	}
	for _, f := range ulpq.ctx.Fields {
		if !userlikeplace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
		}
	}
	if ulpq.path != nil {
		prev, err := ulpq.path(ctx)
		if err != nil {
			return err
		}
		ulpq.sql = prev
	}
	return nil
}

func (ulpq *UserLikePlaceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserLikePlace, error) {
	var (
		nodes       = []*UserLikePlace{}
		withFKs     = ulpq.withFKs
		_spec       = ulpq.querySpec()
		loadedTypes = [2]bool{
			ulpq.withUser != nil,
			ulpq.withPlace != nil,
		}
	)
	if ulpq.withUser != nil || ulpq.withPlace != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userlikeplace.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserLikePlace).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserLikePlace{config: ulpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ulpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ulpq.withUser; query != nil {
		if err := ulpq.loadUser(ctx, query, nodes, nil,
			func(n *UserLikePlace, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ulpq.withPlace; query != nil {
		if err := ulpq.loadPlace(ctx, query, nodes, nil,
			func(n *UserLikePlace, e *Place) { n.Edges.Place = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ulpq *UserLikePlaceQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserLikePlace, init func(*UserLikePlace), assign func(*UserLikePlace, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserLikePlace)
	for i := range nodes {
		if nodes[i].user_liked_places == nil {
			continue
		}
		fk := *nodes[i].user_liked_places
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_liked_places" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ulpq *UserLikePlaceQuery) loadPlace(ctx context.Context, query *PlaceQuery, nodes []*UserLikePlace, init func(*UserLikePlace), assign func(*UserLikePlace, *Place)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserLikePlace)
	for i := range nodes {
		if nodes[i].user_like_place_place == nil {
			continue
		}
		fk := *nodes[i].user_like_place_place
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(place.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_like_place_place" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ulpq *UserLikePlaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ulpq.querySpec()
	_spec.Node.Columns = ulpq.ctx.Fields
	if len(ulpq.ctx.Fields) > 0 {
		_spec.Unique = ulpq.ctx.Unique != nil && *ulpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ulpq.driver, _spec)
}

func (ulpq *UserLikePlaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userlikeplace.Table, userlikeplace.Columns, sqlgraph.NewFieldSpec(userlikeplace.FieldID, field.TypeString))
	_spec.From = ulpq.sql
	if unique := ulpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ulpq.path != nil {
		_spec.Unique = true
	}
	if fields := ulpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userlikeplace.FieldID)
		for i := range fields {
			if fields[i] != userlikeplace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ulpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ulpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ulpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ulpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ulpq *UserLikePlaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ulpq.driver.Dialect())
	t1 := builder.Table(userlikeplace.Table)
	columns := ulpq.ctx.Fields
	if len(columns) == 0 {
		columns = userlikeplace.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ulpq.sql != nil {
		selector = ulpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ulpq.ctx.Unique != nil && *ulpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ulpq.predicates {
		p(selector)
	}
	for _, p := range ulpq.order {
		p(selector)
	}
	if offset := ulpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ulpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserLikePlaceGroupBy is the group-by builder for UserLikePlace entities.
type UserLikePlaceGroupBy struct {
	selector
	build *UserLikePlaceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ulpgb *UserLikePlaceGroupBy) Aggregate(fns ...AggregateFunc) *UserLikePlaceGroupBy {
	ulpgb.fns = append(ulpgb.fns, fns...)
	return ulpgb
}

// Scan applies the selector query and scans the result into the given value.
func (ulpgb *UserLikePlaceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ulpgb.build.ctx, "GroupBy")
	if err := ulpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserLikePlaceQuery, *UserLikePlaceGroupBy](ctx, ulpgb.build, ulpgb, ulpgb.build.inters, v)
}

func (ulpgb *UserLikePlaceGroupBy) sqlScan(ctx context.Context, root *UserLikePlaceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ulpgb.fns))
	for _, fn := range ulpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ulpgb.flds)+len(ulpgb.fns))
		for _, f := range *ulpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ulpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ulpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserLikePlaceSelect is the builder for selecting fields of UserLikePlace entities.
type UserLikePlaceSelect struct {
	*UserLikePlaceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ulps *UserLikePlaceSelect) Aggregate(fns ...AggregateFunc) *UserLikePlaceSelect {
	ulps.fns = append(ulps.fns, fns...)
	return ulps
}

// Scan applies the selector query and scans the result into the given value.
func (ulps *UserLikePlaceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ulps.ctx, "Select")
	if err := ulps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserLikePlaceQuery, *UserLikePlaceSelect](ctx, ulps.UserLikePlaceQuery, ulps, ulps.inters, v)
}

func (ulps *UserLikePlaceSelect) sqlScan(ctx context.Context, root *UserLikePlaceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ulps.fns))
	for _, fn := range ulps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ulps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ulps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
