// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"placio_api/booking"
	"placio_api/place"
	"placio_api/predicate"
	"placio_api/room"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoomQuery is the builder for querying Room entities.
type RoomQuery struct {
	config
	ctx          *QueryContext
	order        []room.OrderOption
	inters       []Interceptor
	predicates   []predicate.Room
	withPlace    *PlaceQuery
	withBookings *BookingQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoomQuery builder.
func (rq *RoomQuery) Where(ps ...predicate.Room) *RoomQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RoomQuery) Limit(limit int) *RoomQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RoomQuery) Offset(offset int) *RoomQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RoomQuery) Unique(unique bool) *RoomQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RoomQuery) Order(o ...room.OrderOption) *RoomQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryPlace chains the current query on the "place" edge.
func (rq *RoomQuery) QueryPlace() *PlaceQuery {
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
			sqlgraph.From(room.Table, room.FieldID, selector),
			sqlgraph.To(place.Table, place.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, room.PlaceTable, room.PlaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBookings chains the current query on the "bookings" edge.
func (rq *RoomQuery) QueryBookings() *BookingQuery {
	query := (&BookingClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(room.Table, room.FieldID, selector),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, room.BookingsTable, room.BookingsColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Room entity from the query.
// Returns a *NotFoundError when no Room was found.
func (rq *RoomQuery) First(ctx context.Context) (*Room, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{room.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RoomQuery) FirstX(ctx context.Context) *Room {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Room ID from the query.
// Returns a *NotFoundError when no Room ID was found.
func (rq *RoomQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{room.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RoomQuery) FirstIDX(ctx context.Context) string {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Room entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Room entity is found.
// Returns a *NotFoundError when no Room entities are found.
func (rq *RoomQuery) Only(ctx context.Context) (*Room, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{room.Label}
	default:
		return nil, &NotSingularError{room.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RoomQuery) OnlyX(ctx context.Context) *Room {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Room ID in the query.
// Returns a *NotSingularError when more than one Room ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RoomQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{room.Label}
	default:
		err = &NotSingularError{room.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RoomQuery) OnlyIDX(ctx context.Context) string {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Rooms.
func (rq *RoomQuery) All(ctx context.Context) ([]*Room, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Room, *RoomQuery]()
	return withInterceptors[[]*Room](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RoomQuery) AllX(ctx context.Context) []*Room {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Room IDs.
func (rq *RoomQuery) IDs(ctx context.Context) (ids []string, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(room.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RoomQuery) IDsX(ctx context.Context) []string {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RoomQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RoomQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RoomQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RoomQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("placio_api: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RoomQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoomQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RoomQuery) Clone() *RoomQuery {
	if rq == nil {
		return nil
	}
	return &RoomQuery{
		config:       rq.config,
		ctx:          rq.ctx.Clone(),
		order:        append([]room.OrderOption{}, rq.order...),
		inters:       append([]Interceptor{}, rq.inters...),
		predicates:   append([]predicate.Room{}, rq.predicates...),
		withPlace:    rq.withPlace.Clone(),
		withBookings: rq.withBookings.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithPlace tells the query-builder to eager-load the nodes that are connected to
// the "place" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoomQuery) WithPlace(opts ...func(*PlaceQuery)) *RoomQuery {
	query := (&PlaceClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withPlace = query
	return rq
}

// WithBookings tells the query-builder to eager-load the nodes that are connected to
// the "bookings" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoomQuery) WithBookings(opts ...func(*BookingQuery)) *RoomQuery {
	query := (&BookingClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withBookings = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Number string `json:"number,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Room.Query().
//		GroupBy(room.FieldNumber).
//		Aggregate(placio_api.Count()).
//		Scan(ctx, &v)
func (rq *RoomQuery) GroupBy(field string, fields ...string) *RoomGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoomGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = room.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Number string `json:"number,omitempty"`
//	}
//
//	client.Room.Query().
//		Select(room.FieldNumber).
//		Scan(ctx, &v)
func (rq *RoomQuery) Select(fields ...string) *RoomSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RoomSelect{RoomQuery: rq}
	sbuild.label = room.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoomSelect configured with the given aggregations.
func (rq *RoomQuery) Aggregate(fns ...AggregateFunc) *RoomSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RoomQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("placio_api: uninitialized interceptor (forgotten import placio_api/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !room.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
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

func (rq *RoomQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Room, error) {
	var (
		nodes       = []*Room{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withPlace != nil,
			rq.withBookings != nil,
		}
	)
	if rq.withPlace != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, room.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Room).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Room{config: rq.config}
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
	if query := rq.withPlace; query != nil {
		if err := rq.loadPlace(ctx, query, nodes, nil,
			func(n *Room, e *Place) { n.Edges.Place = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withBookings; query != nil {
		if err := rq.loadBookings(ctx, query, nodes,
			func(n *Room) { n.Edges.Bookings = []*Booking{} },
			func(n *Room, e *Booking) { n.Edges.Bookings = append(n.Edges.Bookings, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RoomQuery) loadPlace(ctx context.Context, query *PlaceQuery, nodes []*Room, init func(*Room), assign func(*Room, *Place)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Room)
	for i := range nodes {
		if nodes[i].place_rooms == nil {
			continue
		}
		fk := *nodes[i].place_rooms
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
			return fmt.Errorf(`unexpected foreign-key "place_rooms" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RoomQuery) loadBookings(ctx context.Context, query *BookingQuery, nodes []*Room, init func(*Room), assign func(*Room, *Booking)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Room)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(room.BookingsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.room_bookings
		if fk == nil {
			return fmt.Errorf(`foreign-key "room_bookings" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "room_bookings" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *RoomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RoomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(room.Table, room.Columns, sqlgraph.NewFieldSpec(room.FieldID, field.TypeString))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, room.FieldID)
		for i := range fields {
			if fields[i] != room.FieldID {
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

func (rq *RoomQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(room.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = room.Columns
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

// RoomGroupBy is the group-by builder for Room entities.
type RoomGroupBy struct {
	selector
	build *RoomQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RoomGroupBy) Aggregate(fns ...AggregateFunc) *RoomGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RoomGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoomQuery, *RoomGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RoomGroupBy) sqlScan(ctx context.Context, root *RoomQuery, v any) error {
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

// RoomSelect is the builder for selecting fields of Room entities.
type RoomSelect struct {
	*RoomQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RoomSelect) Aggregate(fns ...AggregateFunc) *RoomSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RoomSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoomQuery, *RoomSelect](ctx, rs.RoomQuery, rs, rs.inters, v)
}

func (rs *RoomSelect) sqlScan(ctx context.Context, root *RoomQuery, v any) error {
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
