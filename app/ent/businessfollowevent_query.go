// Code generated by ent, DO NOT EDIT.

package placio_api

import (
	"context"
	"fmt"
	"math"
	"placio_api/business"
	"placio_api/businessfollowevent"
	"placio_api/event"
	"placio_api/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BusinessFollowEventQuery is the builder for querying BusinessFollowEvent entities.
type BusinessFollowEventQuery struct {
	config
	ctx          *QueryContext
	order        []businessfollowevent.OrderOption
	inters       []Interceptor
	predicates   []predicate.BusinessFollowEvent
	withBusiness *BusinessQuery
	withEvent    *EventQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BusinessFollowEventQuery builder.
func (bfeq *BusinessFollowEventQuery) Where(ps ...predicate.BusinessFollowEvent) *BusinessFollowEventQuery {
	bfeq.predicates = append(bfeq.predicates, ps...)
	return bfeq
}

// Limit the number of records to be returned by this query.
func (bfeq *BusinessFollowEventQuery) Limit(limit int) *BusinessFollowEventQuery {
	bfeq.ctx.Limit = &limit
	return bfeq
}

// Offset to start from.
func (bfeq *BusinessFollowEventQuery) Offset(offset int) *BusinessFollowEventQuery {
	bfeq.ctx.Offset = &offset
	return bfeq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bfeq *BusinessFollowEventQuery) Unique(unique bool) *BusinessFollowEventQuery {
	bfeq.ctx.Unique = &unique
	return bfeq
}

// Order specifies how the records should be ordered.
func (bfeq *BusinessFollowEventQuery) Order(o ...businessfollowevent.OrderOption) *BusinessFollowEventQuery {
	bfeq.order = append(bfeq.order, o...)
	return bfeq
}

// QueryBusiness chains the current query on the "business" edge.
func (bfeq *BusinessFollowEventQuery) QueryBusiness() *BusinessQuery {
	query := (&BusinessClient{config: bfeq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bfeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bfeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(businessfollowevent.Table, businessfollowevent.FieldID, selector),
			sqlgraph.To(business.Table, business.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, businessfollowevent.BusinessTable, businessfollowevent.BusinessColumn),
		)
		fromU = sqlgraph.SetNeighbors(bfeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEvent chains the current query on the "event" edge.
func (bfeq *BusinessFollowEventQuery) QueryEvent() *EventQuery {
	query := (&EventClient{config: bfeq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bfeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bfeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(businessfollowevent.Table, businessfollowevent.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, businessfollowevent.EventTable, businessfollowevent.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(bfeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BusinessFollowEvent entity from the query.
// Returns a *NotFoundError when no BusinessFollowEvent was found.
func (bfeq *BusinessFollowEventQuery) First(ctx context.Context) (*BusinessFollowEvent, error) {
	nodes, err := bfeq.Limit(1).All(setContextOp(ctx, bfeq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{businessfollowevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bfeq *BusinessFollowEventQuery) FirstX(ctx context.Context) *BusinessFollowEvent {
	node, err := bfeq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BusinessFollowEvent ID from the query.
// Returns a *NotFoundError when no BusinessFollowEvent ID was found.
func (bfeq *BusinessFollowEventQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bfeq.Limit(1).IDs(setContextOp(ctx, bfeq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{businessfollowevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bfeq *BusinessFollowEventQuery) FirstIDX(ctx context.Context) string {
	id, err := bfeq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BusinessFollowEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BusinessFollowEvent entity is found.
// Returns a *NotFoundError when no BusinessFollowEvent entities are found.
func (bfeq *BusinessFollowEventQuery) Only(ctx context.Context) (*BusinessFollowEvent, error) {
	nodes, err := bfeq.Limit(2).All(setContextOp(ctx, bfeq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{businessfollowevent.Label}
	default:
		return nil, &NotSingularError{businessfollowevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bfeq *BusinessFollowEventQuery) OnlyX(ctx context.Context) *BusinessFollowEvent {
	node, err := bfeq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BusinessFollowEvent ID in the query.
// Returns a *NotSingularError when more than one BusinessFollowEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (bfeq *BusinessFollowEventQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bfeq.Limit(2).IDs(setContextOp(ctx, bfeq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{businessfollowevent.Label}
	default:
		err = &NotSingularError{businessfollowevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bfeq *BusinessFollowEventQuery) OnlyIDX(ctx context.Context) string {
	id, err := bfeq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BusinessFollowEvents.
func (bfeq *BusinessFollowEventQuery) All(ctx context.Context) ([]*BusinessFollowEvent, error) {
	ctx = setContextOp(ctx, bfeq.ctx, "All")
	if err := bfeq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BusinessFollowEvent, *BusinessFollowEventQuery]()
	return withInterceptors[[]*BusinessFollowEvent](ctx, bfeq, qr, bfeq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bfeq *BusinessFollowEventQuery) AllX(ctx context.Context) []*BusinessFollowEvent {
	nodes, err := bfeq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BusinessFollowEvent IDs.
func (bfeq *BusinessFollowEventQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bfeq.ctx.Unique == nil && bfeq.path != nil {
		bfeq.Unique(true)
	}
	ctx = setContextOp(ctx, bfeq.ctx, "IDs")
	if err = bfeq.Select(businessfollowevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bfeq *BusinessFollowEventQuery) IDsX(ctx context.Context) []string {
	ids, err := bfeq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bfeq *BusinessFollowEventQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bfeq.ctx, "Count")
	if err := bfeq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bfeq, querierCount[*BusinessFollowEventQuery](), bfeq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bfeq *BusinessFollowEventQuery) CountX(ctx context.Context) int {
	count, err := bfeq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bfeq *BusinessFollowEventQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bfeq.ctx, "Exist")
	switch _, err := bfeq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("placio_api: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bfeq *BusinessFollowEventQuery) ExistX(ctx context.Context) bool {
	exist, err := bfeq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BusinessFollowEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bfeq *BusinessFollowEventQuery) Clone() *BusinessFollowEventQuery {
	if bfeq == nil {
		return nil
	}
	return &BusinessFollowEventQuery{
		config:       bfeq.config,
		ctx:          bfeq.ctx.Clone(),
		order:        append([]businessfollowevent.OrderOption{}, bfeq.order...),
		inters:       append([]Interceptor{}, bfeq.inters...),
		predicates:   append([]predicate.BusinessFollowEvent{}, bfeq.predicates...),
		withBusiness: bfeq.withBusiness.Clone(),
		withEvent:    bfeq.withEvent.Clone(),
		// clone intermediate query.
		sql:  bfeq.sql.Clone(),
		path: bfeq.path,
	}
}

// WithBusiness tells the query-builder to eager-load the nodes that are connected to
// the "business" edge. The optional arguments are used to configure the query builder of the edge.
func (bfeq *BusinessFollowEventQuery) WithBusiness(opts ...func(*BusinessQuery)) *BusinessFollowEventQuery {
	query := (&BusinessClient{config: bfeq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bfeq.withBusiness = query
	return bfeq
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (bfeq *BusinessFollowEventQuery) WithEvent(opts ...func(*EventQuery)) *BusinessFollowEventQuery {
	query := (&EventClient{config: bfeq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bfeq.withEvent = query
	return bfeq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BusinessFollowEvent.Query().
//		GroupBy(businessfollowevent.FieldCreatedAt).
//		Aggregate(placio_api.Count()).
//		Scan(ctx, &v)
func (bfeq *BusinessFollowEventQuery) GroupBy(field string, fields ...string) *BusinessFollowEventGroupBy {
	bfeq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BusinessFollowEventGroupBy{build: bfeq}
	grbuild.flds = &bfeq.ctx.Fields
	grbuild.label = businessfollowevent.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//	}
//
//	client.BusinessFollowEvent.Query().
//		Select(businessfollowevent.FieldCreatedAt).
//		Scan(ctx, &v)
func (bfeq *BusinessFollowEventQuery) Select(fields ...string) *BusinessFollowEventSelect {
	bfeq.ctx.Fields = append(bfeq.ctx.Fields, fields...)
	sbuild := &BusinessFollowEventSelect{BusinessFollowEventQuery: bfeq}
	sbuild.label = businessfollowevent.Label
	sbuild.flds, sbuild.scan = &bfeq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BusinessFollowEventSelect configured with the given aggregations.
func (bfeq *BusinessFollowEventQuery) Aggregate(fns ...AggregateFunc) *BusinessFollowEventSelect {
	return bfeq.Select().Aggregate(fns...)
}

func (bfeq *BusinessFollowEventQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bfeq.inters {
		if inter == nil {
			return fmt.Errorf("placio_api: uninitialized interceptor (forgotten import placio_api/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bfeq); err != nil {
				return err
			}
		}
	}
	for _, f := range bfeq.ctx.Fields {
		if !businessfollowevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
		}
	}
	if bfeq.path != nil {
		prev, err := bfeq.path(ctx)
		if err != nil {
			return err
		}
		bfeq.sql = prev
	}
	return nil
}

func (bfeq *BusinessFollowEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BusinessFollowEvent, error) {
	var (
		nodes       = []*BusinessFollowEvent{}
		withFKs     = bfeq.withFKs
		_spec       = bfeq.querySpec()
		loadedTypes = [2]bool{
			bfeq.withBusiness != nil,
			bfeq.withEvent != nil,
		}
	)
	if bfeq.withBusiness != nil || bfeq.withEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, businessfollowevent.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BusinessFollowEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BusinessFollowEvent{config: bfeq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bfeq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bfeq.withBusiness; query != nil {
		if err := bfeq.loadBusiness(ctx, query, nodes, nil,
			func(n *BusinessFollowEvent, e *Business) { n.Edges.Business = e }); err != nil {
			return nil, err
		}
	}
	if query := bfeq.withEvent; query != nil {
		if err := bfeq.loadEvent(ctx, query, nodes, nil,
			func(n *BusinessFollowEvent, e *Event) { n.Edges.Event = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bfeq *BusinessFollowEventQuery) loadBusiness(ctx context.Context, query *BusinessQuery, nodes []*BusinessFollowEvent, init func(*BusinessFollowEvent), assign func(*BusinessFollowEvent, *Business)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*BusinessFollowEvent)
	for i := range nodes {
		if nodes[i].business_business_follow_events == nil {
			continue
		}
		fk := *nodes[i].business_business_follow_events
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
			return fmt.Errorf(`unexpected foreign-key "business_business_follow_events" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bfeq *BusinessFollowEventQuery) loadEvent(ctx context.Context, query *EventQuery, nodes []*BusinessFollowEvent, init func(*BusinessFollowEvent), assign func(*BusinessFollowEvent, *Event)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*BusinessFollowEvent)
	for i := range nodes {
		if nodes[i].business_follow_event_event == nil {
			continue
		}
		fk := *nodes[i].business_follow_event_event
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
			return fmt.Errorf(`unexpected foreign-key "business_follow_event_event" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (bfeq *BusinessFollowEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bfeq.querySpec()
	_spec.Node.Columns = bfeq.ctx.Fields
	if len(bfeq.ctx.Fields) > 0 {
		_spec.Unique = bfeq.ctx.Unique != nil && *bfeq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bfeq.driver, _spec)
}

func (bfeq *BusinessFollowEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(businessfollowevent.Table, businessfollowevent.Columns, sqlgraph.NewFieldSpec(businessfollowevent.FieldID, field.TypeString))
	_spec.From = bfeq.sql
	if unique := bfeq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bfeq.path != nil {
		_spec.Unique = true
	}
	if fields := bfeq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, businessfollowevent.FieldID)
		for i := range fields {
			if fields[i] != businessfollowevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bfeq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bfeq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bfeq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bfeq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bfeq *BusinessFollowEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bfeq.driver.Dialect())
	t1 := builder.Table(businessfollowevent.Table)
	columns := bfeq.ctx.Fields
	if len(columns) == 0 {
		columns = businessfollowevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bfeq.sql != nil {
		selector = bfeq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bfeq.ctx.Unique != nil && *bfeq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bfeq.predicates {
		p(selector)
	}
	for _, p := range bfeq.order {
		p(selector)
	}
	if offset := bfeq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bfeq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BusinessFollowEventGroupBy is the group-by builder for BusinessFollowEvent entities.
type BusinessFollowEventGroupBy struct {
	selector
	build *BusinessFollowEventQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bfegb *BusinessFollowEventGroupBy) Aggregate(fns ...AggregateFunc) *BusinessFollowEventGroupBy {
	bfegb.fns = append(bfegb.fns, fns...)
	return bfegb
}

// Scan applies the selector query and scans the result into the given value.
func (bfegb *BusinessFollowEventGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bfegb.build.ctx, "GroupBy")
	if err := bfegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BusinessFollowEventQuery, *BusinessFollowEventGroupBy](ctx, bfegb.build, bfegb, bfegb.build.inters, v)
}

func (bfegb *BusinessFollowEventGroupBy) sqlScan(ctx context.Context, root *BusinessFollowEventQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bfegb.fns))
	for _, fn := range bfegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bfegb.flds)+len(bfegb.fns))
		for _, f := range *bfegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bfegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bfegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BusinessFollowEventSelect is the builder for selecting fields of BusinessFollowEvent entities.
type BusinessFollowEventSelect struct {
	*BusinessFollowEventQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bfes *BusinessFollowEventSelect) Aggregate(fns ...AggregateFunc) *BusinessFollowEventSelect {
	bfes.fns = append(bfes.fns, fns...)
	return bfes
}

// Scan applies the selector query and scans the result into the given value.
func (bfes *BusinessFollowEventSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bfes.ctx, "Select")
	if err := bfes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BusinessFollowEventQuery, *BusinessFollowEventSelect](ctx, bfes.BusinessFollowEventQuery, bfes, bfes.inters, v)
}

func (bfes *BusinessFollowEventSelect) sqlScan(ctx context.Context, root *BusinessFollowEventQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bfes.fns))
	for _, fn := range bfes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bfes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bfes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
