// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"placio-app/ent/business"
	"placio-app/ent/event"
	"placio-app/ent/place"
	"placio-app/ent/predicate"
	"placio-app/ent/rating"
	"placio-app/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RatingQuery is the builder for querying Rating entities.
type RatingQuery struct {
	config
	ctx          *QueryContext
	order        []rating.OrderOption
	inters       []Interceptor
	predicates   []predicate.Rating
	withUser     *UserQuery
	withBusiness *BusinessQuery
	withPlace    *PlaceQuery
	withEvent    *EventQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RatingQuery builder.
func (rq *RatingQuery) Where(ps ...predicate.Rating) *RatingQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RatingQuery) Limit(limit int) *RatingQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RatingQuery) Offset(offset int) *RatingQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RatingQuery) Unique(unique bool) *RatingQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RatingQuery) Order(o ...rating.OrderOption) *RatingQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryUser chains the current query on the "user" edge.
func (rq *RatingQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rating.Table, rating.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, rating.UserTable, rating.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusiness chains the current query on the "business" edge.
func (rq *RatingQuery) QueryBusiness() *BusinessQuery {
	query := (&BusinessClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rating.Table, rating.FieldID, selector),
			sqlgraph.To(business.Table, business.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, rating.BusinessTable, rating.BusinessColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlace chains the current query on the "place" edge.
func (rq *RatingQuery) QueryPlace() *PlaceQuery {
	query := (&PlaceClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rating.Table, rating.FieldID, selector),
			sqlgraph.To(place.Table, place.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, rating.PlaceTable, rating.PlaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvent chains the current query on the "event" edge.
func (rq *RatingQuery) QueryEvent() *EventQuery {
	query := (&EventClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rating.Table, rating.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, rating.EventTable, rating.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Rating entity from the query.
// Returns a *NotFoundError when no Rating was found.
func (rq *RatingQuery) First(ctx context.Context) (*Rating, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rating.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RatingQuery) FirstX(ctx context.Context) *Rating {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Rating ID from the query.
// Returns a *NotFoundError when no Rating ID was found.
func (rq *RatingQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{rating.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RatingQuery) FirstIDX(ctx context.Context) string {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Rating entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Rating entity is found.
// Returns a *NotFoundError when no Rating entities are found.
func (rq *RatingQuery) Only(ctx context.Context) (*Rating, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rating.Label}
	default:
		return nil, &NotSingularError{rating.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RatingQuery) OnlyX(ctx context.Context) *Rating {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Rating ID in the query.
// Returns a *NotSingularError when more than one Rating ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RatingQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{rating.Label}
	default:
		err = &NotSingularError{rating.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RatingQuery) OnlyIDX(ctx context.Context) string {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Ratings.
func (rq *RatingQuery) All(ctx context.Context) ([]*Rating, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Rating, *RatingQuery]()
	return withInterceptors[[]*Rating](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RatingQuery) AllX(ctx context.Context) []*Rating {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Rating IDs.
func (rq *RatingQuery) IDs(ctx context.Context) (ids []string, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(rating.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RatingQuery) IDsX(ctx context.Context) []string {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RatingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RatingQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RatingQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RatingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RatingQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RatingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RatingQuery) Clone() *RatingQuery {
	if rq == nil {
		return nil
	}
	return &RatingQuery{
		config:       rq.config,
		ctx:          rq.ctx.Clone(),
		order:        append([]rating.OrderOption{}, rq.order...),
		inters:       append([]Interceptor{}, rq.inters...),
		predicates:   append([]predicate.Rating{}, rq.predicates...),
		withUser:     rq.withUser.Clone(),
		withBusiness: rq.withBusiness.Clone(),
		withPlace:    rq.withPlace.Clone(),
		withEvent:    rq.withEvent.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RatingQuery) WithUser(opts ...func(*UserQuery)) *RatingQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withUser = query
	return rq
}

// WithBusiness tells the query-builder to eager-load the nodes that are connected to
// the "business" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RatingQuery) WithBusiness(opts ...func(*BusinessQuery)) *RatingQuery {
	query := (&BusinessClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withBusiness = query
	return rq
}

// WithPlace tells the query-builder to eager-load the nodes that are connected to
// the "place" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RatingQuery) WithPlace(opts ...func(*PlaceQuery)) *RatingQuery {
	query := (&PlaceClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withPlace = query
	return rq
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RatingQuery) WithEvent(opts ...func(*EventQuery)) *RatingQuery {
	query := (&EventClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withEvent = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Score int `json:"score,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Rating.Query().
//		GroupBy(rating.FieldScore).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RatingQuery) GroupBy(field string, fields ...string) *RatingGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RatingGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = rating.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Score int `json:"score,omitempty"`
//	}
//
//	client.Rating.Query().
//		Select(rating.FieldScore).
//		Scan(ctx, &v)
func (rq *RatingQuery) Select(fields ...string) *RatingSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RatingSelect{RatingQuery: rq}
	sbuild.label = rating.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RatingSelect configured with the given aggregations.
func (rq *RatingQuery) Aggregate(fns ...AggregateFunc) *RatingSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RatingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !rating.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RatingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Rating, error) {
	var (
		nodes       = []*Rating{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [4]bool{
			rq.withUser != nil,
			rq.withBusiness != nil,
			rq.withPlace != nil,
			rq.withEvent != nil,
		}
	)
	if rq.withUser != nil || rq.withBusiness != nil || rq.withPlace != nil || rq.withEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, rating.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Rating).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Rating{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withUser; query != nil {
		if err := rq.loadUser(ctx, query, nodes, nil,
			func(n *Rating, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withBusiness; query != nil {
		if err := rq.loadBusiness(ctx, query, nodes, nil,
			func(n *Rating, e *Business) { n.Edges.Business = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withPlace; query != nil {
		if err := rq.loadPlace(ctx, query, nodes, nil,
			func(n *Rating, e *Place) { n.Edges.Place = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withEvent; query != nil {
		if err := rq.loadEvent(ctx, query, nodes, nil,
			func(n *Rating, e *Event) { n.Edges.Event = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RatingQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Rating, init func(*Rating), assign func(*Rating, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Rating)
	for i := range nodes {
		if nodes[i].user_ratings == nil {
			continue
		}
		fk := *nodes[i].user_ratings
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
			return fmt.Errorf(`unexpected foreign-key "user_ratings" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RatingQuery) loadBusiness(ctx context.Context, query *BusinessQuery, nodes []*Rating, init func(*Rating), assign func(*Rating, *Business)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Rating)
	for i := range nodes {
		if nodes[i].rating_business == nil {
			continue
		}
		fk := *nodes[i].rating_business
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
			return fmt.Errorf(`unexpected foreign-key "rating_business" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RatingQuery) loadPlace(ctx context.Context, query *PlaceQuery, nodes []*Rating, init func(*Rating), assign func(*Rating, *Place)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Rating)
	for i := range nodes {
		if nodes[i].rating_place == nil {
			continue
		}
		fk := *nodes[i].rating_place
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
			return fmt.Errorf(`unexpected foreign-key "rating_place" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RatingQuery) loadEvent(ctx context.Context, query *EventQuery, nodes []*Rating, init func(*Rating), assign func(*Rating, *Event)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Rating)
	for i := range nodes {
		if nodes[i].rating_event == nil {
			continue
		}
		fk := *nodes[i].rating_event
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(event.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "rating_event" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *RatingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RatingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(rating.Table, rating.Columns, sqlgraph.NewFieldSpec(rating.FieldID, field.TypeString))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rating.FieldID)
		for i := range fields {
			if fields[i] != rating.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RatingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(rating.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = rating.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RatingGroupBy is the group-by builder for Rating entities.
type RatingGroupBy struct {
	selector
	build *RatingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RatingGroupBy) Aggregate(fns ...AggregateFunc) *RatingGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RatingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RatingQuery, *RatingGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RatingGroupBy) sqlScan(ctx context.Context, root *RatingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RatingSelect is the builder for selecting fields of Rating entities.
type RatingSelect struct {
	*RatingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RatingSelect) Aggregate(fns ...AggregateFunc) *RatingSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RatingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RatingQuery, *RatingSelect](ctx, rs.RatingQuery, rs, rs.inters, v)
}

func (rs *RatingSelect) sqlScan(ctx context.Context, root *RatingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
