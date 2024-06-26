// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"placio-app/ent/place"
	"placio-app/ent/predicate"
	"placio-app/ent/user"
	"placio-app/ent/userfollowplace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFollowPlaceQuery is the builder for querying UserFollowPlace entities.
type UserFollowPlaceQuery struct {
	config
	ctx        *QueryContext
	order      []userfollowplace.OrderOption
	inters     []Interceptor
	predicates []predicate.UserFollowPlace
	withUser   *UserQuery
	withPlace  *PlaceQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserFollowPlaceQuery builder.
func (ufpq *UserFollowPlaceQuery) Where(ps ...predicate.UserFollowPlace) *UserFollowPlaceQuery {
	ufpq.predicates = append(ufpq.predicates, ps...)
	return ufpq
}

// Limit the number of records to be returned by this query.
func (ufpq *UserFollowPlaceQuery) Limit(limit int) *UserFollowPlaceQuery {
	ufpq.ctx.Limit = &limit
	return ufpq
}

// Offset to start from.
func (ufpq *UserFollowPlaceQuery) Offset(offset int) *UserFollowPlaceQuery {
	ufpq.ctx.Offset = &offset
	return ufpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ufpq *UserFollowPlaceQuery) Unique(unique bool) *UserFollowPlaceQuery {
	ufpq.ctx.Unique = &unique
	return ufpq
}

// Order specifies how the records should be ordered.
func (ufpq *UserFollowPlaceQuery) Order(o ...userfollowplace.OrderOption) *UserFollowPlaceQuery {
	ufpq.order = append(ufpq.order, o...)
	return ufpq
}

// QueryUser chains the current query on the "user" edge.
func (ufpq *UserFollowPlaceQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ufpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userfollowplace.Table, userfollowplace.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userfollowplace.UserTable, userfollowplace.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlace chains the current query on the "place" edge.
func (ufpq *UserFollowPlaceQuery) QueryPlace() *PlaceQuery {
	query := (&PlaceClient{config: ufpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userfollowplace.Table, userfollowplace.FieldID, selector),
			sqlgraph.To(place.Table, place.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, userfollowplace.PlaceTable, userfollowplace.PlaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserFollowPlace entity from the query.
// Returns a *NotFoundError when no UserFollowPlace was found.
func (ufpq *UserFollowPlaceQuery) First(ctx context.Context) (*UserFollowPlace, error) {
	nodes, err := ufpq.Limit(1).All(setContextOp(ctx, ufpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userfollowplace.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ufpq *UserFollowPlaceQuery) FirstX(ctx context.Context) *UserFollowPlace {
	node, err := ufpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserFollowPlace ID from the query.
// Returns a *NotFoundError when no UserFollowPlace ID was found.
func (ufpq *UserFollowPlaceQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ufpq.Limit(1).IDs(setContextOp(ctx, ufpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userfollowplace.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ufpq *UserFollowPlaceQuery) FirstIDX(ctx context.Context) string {
	id, err := ufpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserFollowPlace entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserFollowPlace entity is found.
// Returns a *NotFoundError when no UserFollowPlace entities are found.
func (ufpq *UserFollowPlaceQuery) Only(ctx context.Context) (*UserFollowPlace, error) {
	nodes, err := ufpq.Limit(2).All(setContextOp(ctx, ufpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userfollowplace.Label}
	default:
		return nil, &NotSingularError{userfollowplace.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ufpq *UserFollowPlaceQuery) OnlyX(ctx context.Context) *UserFollowPlace {
	node, err := ufpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserFollowPlace ID in the query.
// Returns a *NotSingularError when more than one UserFollowPlace ID is found.
// Returns a *NotFoundError when no entities are found.
func (ufpq *UserFollowPlaceQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ufpq.Limit(2).IDs(setContextOp(ctx, ufpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userfollowplace.Label}
	default:
		err = &NotSingularError{userfollowplace.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ufpq *UserFollowPlaceQuery) OnlyIDX(ctx context.Context) string {
	id, err := ufpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserFollowPlaces.
func (ufpq *UserFollowPlaceQuery) All(ctx context.Context) ([]*UserFollowPlace, error) {
	ctx = setContextOp(ctx, ufpq.ctx, "All")
	if err := ufpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserFollowPlace, *UserFollowPlaceQuery]()
	return withInterceptors[[]*UserFollowPlace](ctx, ufpq, qr, ufpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ufpq *UserFollowPlaceQuery) AllX(ctx context.Context) []*UserFollowPlace {
	nodes, err := ufpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserFollowPlace IDs.
func (ufpq *UserFollowPlaceQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ufpq.ctx.Unique == nil && ufpq.path != nil {
		ufpq.Unique(true)
	}
	ctx = setContextOp(ctx, ufpq.ctx, "IDs")
	if err = ufpq.Select(userfollowplace.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ufpq *UserFollowPlaceQuery) IDsX(ctx context.Context) []string {
	ids, err := ufpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ufpq *UserFollowPlaceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ufpq.ctx, "Count")
	if err := ufpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ufpq, querierCount[*UserFollowPlaceQuery](), ufpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ufpq *UserFollowPlaceQuery) CountX(ctx context.Context) int {
	count, err := ufpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ufpq *UserFollowPlaceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ufpq.ctx, "Exist")
	switch _, err := ufpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ufpq *UserFollowPlaceQuery) ExistX(ctx context.Context) bool {
	exist, err := ufpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserFollowPlaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ufpq *UserFollowPlaceQuery) Clone() *UserFollowPlaceQuery {
	if ufpq == nil {
		return nil
	}
	return &UserFollowPlaceQuery{
		config:     ufpq.config,
		ctx:        ufpq.ctx.Clone(),
		order:      append([]userfollowplace.OrderOption{}, ufpq.order...),
		inters:     append([]Interceptor{}, ufpq.inters...),
		predicates: append([]predicate.UserFollowPlace{}, ufpq.predicates...),
		withUser:   ufpq.withUser.Clone(),
		withPlace:  ufpq.withPlace.Clone(),
		// clone intermediate query.
		sql:  ufpq.sql.Clone(),
		path: ufpq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ufpq *UserFollowPlaceQuery) WithUser(opts ...func(*UserQuery)) *UserFollowPlaceQuery {
	query := (&UserClient{config: ufpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufpq.withUser = query
	return ufpq
}

// WithPlace tells the query-builder to eager-load the nodes that are connected to
// the "place" edge. The optional arguments are used to configure the query builder of the edge.
func (ufpq *UserFollowPlaceQuery) WithPlace(opts ...func(*PlaceQuery)) *UserFollowPlaceQuery {
	query := (&PlaceClient{config: ufpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufpq.withPlace = query
	return ufpq
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
//	client.UserFollowPlace.Query().
//		GroupBy(userfollowplace.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ufpq *UserFollowPlaceQuery) GroupBy(field string, fields ...string) *UserFollowPlaceGroupBy {
	ufpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserFollowPlaceGroupBy{build: ufpq}
	grbuild.flds = &ufpq.ctx.Fields
	grbuild.label = userfollowplace.Label
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
//	client.UserFollowPlace.Query().
//		Select(userfollowplace.FieldCreatedAt).
//		Scan(ctx, &v)
func (ufpq *UserFollowPlaceQuery) Select(fields ...string) *UserFollowPlaceSelect {
	ufpq.ctx.Fields = append(ufpq.ctx.Fields, fields...)
	sbuild := &UserFollowPlaceSelect{UserFollowPlaceQuery: ufpq}
	sbuild.label = userfollowplace.Label
	sbuild.flds, sbuild.scan = &ufpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserFollowPlaceSelect configured with the given aggregations.
func (ufpq *UserFollowPlaceQuery) Aggregate(fns ...AggregateFunc) *UserFollowPlaceSelect {
	return ufpq.Select().Aggregate(fns...)
}

func (ufpq *UserFollowPlaceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ufpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ufpq); err != nil {
				return err
			}
		}
	}
	for _, f := range ufpq.ctx.Fields {
		if !userfollowplace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ufpq.path != nil {
		prev, err := ufpq.path(ctx)
		if err != nil {
			return err
		}
		ufpq.sql = prev
	}
	return nil
}

func (ufpq *UserFollowPlaceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserFollowPlace, error) {
	var (
		nodes       = []*UserFollowPlace{}
		withFKs     = ufpq.withFKs
		_spec       = ufpq.querySpec()
		loadedTypes = [2]bool{
			ufpq.withUser != nil,
			ufpq.withPlace != nil,
		}
	)
	if ufpq.withUser != nil || ufpq.withPlace != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userfollowplace.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserFollowPlace).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserFollowPlace{config: ufpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ufpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ufpq.withUser; query != nil {
		if err := ufpq.loadUser(ctx, query, nodes, nil,
			func(n *UserFollowPlace, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ufpq.withPlace; query != nil {
		if err := ufpq.loadPlace(ctx, query, nodes, nil,
			func(n *UserFollowPlace, e *Place) { n.Edges.Place = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ufpq *UserFollowPlaceQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserFollowPlace, init func(*UserFollowPlace), assign func(*UserFollowPlace, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserFollowPlace)
	for i := range nodes {
		if nodes[i].user_followed_places == nil {
			continue
		}
		fk := *nodes[i].user_followed_places
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
			return fmt.Errorf(`unexpected foreign-key "user_followed_places" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ufpq *UserFollowPlaceQuery) loadPlace(ctx context.Context, query *PlaceQuery, nodes []*UserFollowPlace, init func(*UserFollowPlace), assign func(*UserFollowPlace, *Place)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserFollowPlace)
	for i := range nodes {
		if nodes[i].user_follow_place_place == nil {
			continue
		}
		fk := *nodes[i].user_follow_place_place
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
			return fmt.Errorf(`unexpected foreign-key "user_follow_place_place" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ufpq *UserFollowPlaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ufpq.querySpec()
	_spec.Node.Columns = ufpq.ctx.Fields
	if len(ufpq.ctx.Fields) > 0 {
		_spec.Unique = ufpq.ctx.Unique != nil && *ufpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ufpq.driver, _spec)
}

func (ufpq *UserFollowPlaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userfollowplace.Table, userfollowplace.Columns, sqlgraph.NewFieldSpec(userfollowplace.FieldID, field.TypeString))
	_spec.From = ufpq.sql
	if unique := ufpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ufpq.path != nil {
		_spec.Unique = true
	}
	if fields := ufpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userfollowplace.FieldID)
		for i := range fields {
			if fields[i] != userfollowplace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ufpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ufpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ufpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ufpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ufpq *UserFollowPlaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ufpq.driver.Dialect())
	t1 := builder.Table(userfollowplace.Table)
	columns := ufpq.ctx.Fields
	if len(columns) == 0 {
		columns = userfollowplace.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ufpq.sql != nil {
		selector = ufpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ufpq.ctx.Unique != nil && *ufpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ufpq.predicates {
		p(selector)
	}
	for _, p := range ufpq.order {
		p(selector)
	}
	if offset := ufpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ufpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserFollowPlaceGroupBy is the group-by builder for UserFollowPlace entities.
type UserFollowPlaceGroupBy struct {
	selector
	build *UserFollowPlaceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ufpgb *UserFollowPlaceGroupBy) Aggregate(fns ...AggregateFunc) *UserFollowPlaceGroupBy {
	ufpgb.fns = append(ufpgb.fns, fns...)
	return ufpgb
}

// Scan applies the selector query and scans the result into the given value.
func (ufpgb *UserFollowPlaceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufpgb.build.ctx, "GroupBy")
	if err := ufpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserFollowPlaceQuery, *UserFollowPlaceGroupBy](ctx, ufpgb.build, ufpgb, ufpgb.build.inters, v)
}

func (ufpgb *UserFollowPlaceGroupBy) sqlScan(ctx context.Context, root *UserFollowPlaceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ufpgb.fns))
	for _, fn := range ufpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ufpgb.flds)+len(ufpgb.fns))
		for _, f := range *ufpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ufpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserFollowPlaceSelect is the builder for selecting fields of UserFollowPlace entities.
type UserFollowPlaceSelect struct {
	*UserFollowPlaceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ufps *UserFollowPlaceSelect) Aggregate(fns ...AggregateFunc) *UserFollowPlaceSelect {
	ufps.fns = append(ufps.fns, fns...)
	return ufps
}

// Scan applies the selector query and scans the result into the given value.
func (ufps *UserFollowPlaceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufps.ctx, "Select")
	if err := ufps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserFollowPlaceQuery, *UserFollowPlaceSelect](ctx, ufps.UserFollowPlaceQuery, ufps, ufps.inters, v)
}

func (ufps *UserFollowPlaceSelect) sqlScan(ctx context.Context, root *UserFollowPlaceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ufps.fns))
	for _, fn := range ufps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ufps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
