// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"placio-app/ent/event"
	"placio-app/ent/media"
	"placio-app/ent/predicate"
	"placio-app/ent/ticket"
	"placio-app/ent/ticketoption"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TicketOptionQuery is the builder for querying TicketOption entities.
type TicketOptionQuery struct {
	config
	ctx         *QueryContext
	order       []ticketoption.OrderOption
	inters      []Interceptor
	predicates  []predicate.TicketOption
	withEvent   *EventQuery
	withTickets *TicketQuery
	withMedia   *MediaQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TicketOptionQuery builder.
func (toq *TicketOptionQuery) Where(ps ...predicate.TicketOption) *TicketOptionQuery {
	toq.predicates = append(toq.predicates, ps...)
	return toq
}

// Limit the number of records to be returned by this query.
func (toq *TicketOptionQuery) Limit(limit int) *TicketOptionQuery {
	toq.ctx.Limit = &limit
	return toq
}

// Offset to start from.
func (toq *TicketOptionQuery) Offset(offset int) *TicketOptionQuery {
	toq.ctx.Offset = &offset
	return toq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (toq *TicketOptionQuery) Unique(unique bool) *TicketOptionQuery {
	toq.ctx.Unique = &unique
	return toq
}

// Order specifies how the records should be ordered.
func (toq *TicketOptionQuery) Order(o ...ticketoption.OrderOption) *TicketOptionQuery {
	toq.order = append(toq.order, o...)
	return toq
}

// QueryEvent chains the current query on the "event" edge.
func (toq *TicketOptionQuery) QueryEvent() *EventQuery {
	query := (&EventClient{config: toq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := toq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := toq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ticketoption.Table, ticketoption.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticketoption.EventTable, ticketoption.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(toq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTickets chains the current query on the "tickets" edge.
func (toq *TicketOptionQuery) QueryTickets() *TicketQuery {
	query := (&TicketClient{config: toq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := toq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := toq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ticketoption.Table, ticketoption.FieldID, selector),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ticketoption.TicketsTable, ticketoption.TicketsColumn),
		)
		fromU = sqlgraph.SetNeighbors(toq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMedia chains the current query on the "media" edge.
func (toq *TicketOptionQuery) QueryMedia() *MediaQuery {
	query := (&MediaClient{config: toq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := toq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := toq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ticketoption.Table, ticketoption.FieldID, selector),
			sqlgraph.To(media.Table, media.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ticketoption.MediaTable, ticketoption.MediaColumn),
		)
		fromU = sqlgraph.SetNeighbors(toq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TicketOption entity from the query.
// Returns a *NotFoundError when no TicketOption was found.
func (toq *TicketOptionQuery) First(ctx context.Context) (*TicketOption, error) {
	nodes, err := toq.Limit(1).All(setContextOp(ctx, toq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ticketoption.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (toq *TicketOptionQuery) FirstX(ctx context.Context) *TicketOption {
	node, err := toq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TicketOption ID from the query.
// Returns a *NotFoundError when no TicketOption ID was found.
func (toq *TicketOptionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = toq.Limit(1).IDs(setContextOp(ctx, toq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ticketoption.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (toq *TicketOptionQuery) FirstIDX(ctx context.Context) string {
	id, err := toq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TicketOption entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TicketOption entity is found.
// Returns a *NotFoundError when no TicketOption entities are found.
func (toq *TicketOptionQuery) Only(ctx context.Context) (*TicketOption, error) {
	nodes, err := toq.Limit(2).All(setContextOp(ctx, toq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ticketoption.Label}
	default:
		return nil, &NotSingularError{ticketoption.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (toq *TicketOptionQuery) OnlyX(ctx context.Context) *TicketOption {
	node, err := toq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TicketOption ID in the query.
// Returns a *NotSingularError when more than one TicketOption ID is found.
// Returns a *NotFoundError when no entities are found.
func (toq *TicketOptionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = toq.Limit(2).IDs(setContextOp(ctx, toq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ticketoption.Label}
	default:
		err = &NotSingularError{ticketoption.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (toq *TicketOptionQuery) OnlyIDX(ctx context.Context) string {
	id, err := toq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TicketOptions.
func (toq *TicketOptionQuery) All(ctx context.Context) ([]*TicketOption, error) {
	ctx = setContextOp(ctx, toq.ctx, "All")
	if err := toq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TicketOption, *TicketOptionQuery]()
	return withInterceptors[[]*TicketOption](ctx, toq, qr, toq.inters)
}

// AllX is like All, but panics if an error occurs.
func (toq *TicketOptionQuery) AllX(ctx context.Context) []*TicketOption {
	nodes, err := toq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TicketOption IDs.
func (toq *TicketOptionQuery) IDs(ctx context.Context) (ids []string, err error) {
	if toq.ctx.Unique == nil && toq.path != nil {
		toq.Unique(true)
	}
	ctx = setContextOp(ctx, toq.ctx, "IDs")
	if err = toq.Select(ticketoption.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (toq *TicketOptionQuery) IDsX(ctx context.Context) []string {
	ids, err := toq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (toq *TicketOptionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, toq.ctx, "Count")
	if err := toq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, toq, querierCount[*TicketOptionQuery](), toq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (toq *TicketOptionQuery) CountX(ctx context.Context) int {
	count, err := toq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (toq *TicketOptionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, toq.ctx, "Exist")
	switch _, err := toq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (toq *TicketOptionQuery) ExistX(ctx context.Context) bool {
	exist, err := toq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TicketOptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (toq *TicketOptionQuery) Clone() *TicketOptionQuery {
	if toq == nil {
		return nil
	}
	return &TicketOptionQuery{
		config:      toq.config,
		ctx:         toq.ctx.Clone(),
		order:       append([]ticketoption.OrderOption{}, toq.order...),
		inters:      append([]Interceptor{}, toq.inters...),
		predicates:  append([]predicate.TicketOption{}, toq.predicates...),
		withEvent:   toq.withEvent.Clone(),
		withTickets: toq.withTickets.Clone(),
		withMedia:   toq.withMedia.Clone(),
		// clone intermediate query.
		sql:  toq.sql.Clone(),
		path: toq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (toq *TicketOptionQuery) WithEvent(opts ...func(*EventQuery)) *TicketOptionQuery {
	query := (&EventClient{config: toq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	toq.withEvent = query
	return toq
}

// WithTickets tells the query-builder to eager-load the nodes that are connected to
// the "tickets" edge. The optional arguments are used to configure the query builder of the edge.
func (toq *TicketOptionQuery) WithTickets(opts ...func(*TicketQuery)) *TicketOptionQuery {
	query := (&TicketClient{config: toq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	toq.withTickets = query
	return toq
}

// WithMedia tells the query-builder to eager-load the nodes that are connected to
// the "media" edge. The optional arguments are used to configure the query builder of the edge.
func (toq *TicketOptionQuery) WithMedia(opts ...func(*MediaQuery)) *TicketOptionQuery {
	query := (&MediaClient{config: toq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	toq.withMedia = query
	return toq
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
//	client.TicketOption.Query().
//		GroupBy(ticketoption.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (toq *TicketOptionQuery) GroupBy(field string, fields ...string) *TicketOptionGroupBy {
	toq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TicketOptionGroupBy{build: toq}
	grbuild.flds = &toq.ctx.Fields
	grbuild.label = ticketoption.Label
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
//	client.TicketOption.Query().
//		Select(ticketoption.FieldName).
//		Scan(ctx, &v)
func (toq *TicketOptionQuery) Select(fields ...string) *TicketOptionSelect {
	toq.ctx.Fields = append(toq.ctx.Fields, fields...)
	sbuild := &TicketOptionSelect{TicketOptionQuery: toq}
	sbuild.label = ticketoption.Label
	sbuild.flds, sbuild.scan = &toq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TicketOptionSelect configured with the given aggregations.
func (toq *TicketOptionQuery) Aggregate(fns ...AggregateFunc) *TicketOptionSelect {
	return toq.Select().Aggregate(fns...)
}

func (toq *TicketOptionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range toq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, toq); err != nil {
				return err
			}
		}
	}
	for _, f := range toq.ctx.Fields {
		if !ticketoption.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if toq.path != nil {
		prev, err := toq.path(ctx)
		if err != nil {
			return err
		}
		toq.sql = prev
	}
	return nil
}

func (toq *TicketOptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TicketOption, error) {
	var (
		nodes       = []*TicketOption{}
		withFKs     = toq.withFKs
		_spec       = toq.querySpec()
		loadedTypes = [3]bool{
			toq.withEvent != nil,
			toq.withTickets != nil,
			toq.withMedia != nil,
		}
	)
	if toq.withEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, ticketoption.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TicketOption).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TicketOption{config: toq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, toq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := toq.withEvent; query != nil {
		if err := toq.loadEvent(ctx, query, nodes, nil,
			func(n *TicketOption, e *Event) { n.Edges.Event = e }); err != nil {
			return nil, err
		}
	}
	if query := toq.withTickets; query != nil {
		if err := toq.loadTickets(ctx, query, nodes,
			func(n *TicketOption) { n.Edges.Tickets = []*Ticket{} },
			func(n *TicketOption, e *Ticket) { n.Edges.Tickets = append(n.Edges.Tickets, e) }); err != nil {
			return nil, err
		}
	}
	if query := toq.withMedia; query != nil {
		if err := toq.loadMedia(ctx, query, nodes,
			func(n *TicketOption) { n.Edges.Media = []*Media{} },
			func(n *TicketOption, e *Media) { n.Edges.Media = append(n.Edges.Media, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (toq *TicketOptionQuery) loadEvent(ctx context.Context, query *EventQuery, nodes []*TicketOption, init func(*TicketOption), assign func(*TicketOption, *Event)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*TicketOption)
	for i := range nodes {
		if nodes[i].event_ticket_options == nil {
			continue
		}
		fk := *nodes[i].event_ticket_options
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
			return fmt.Errorf(`unexpected foreign-key "event_ticket_options" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (toq *TicketOptionQuery) loadTickets(ctx context.Context, query *TicketQuery, nodes []*TicketOption, init func(*TicketOption), assign func(*TicketOption, *Ticket)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*TicketOption)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Ticket(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(ticketoption.TicketsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ticket_option_tickets
		if fk == nil {
			return fmt.Errorf(`foreign-key "ticket_option_tickets" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "ticket_option_tickets" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (toq *TicketOptionQuery) loadMedia(ctx context.Context, query *MediaQuery, nodes []*TicketOption, init func(*TicketOption), assign func(*TicketOption, *Media)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*TicketOption)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Media(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(ticketoption.MediaColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ticket_option_media
		if fk == nil {
			return fmt.Errorf(`foreign-key "ticket_option_media" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "ticket_option_media" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (toq *TicketOptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := toq.querySpec()
	_spec.Node.Columns = toq.ctx.Fields
	if len(toq.ctx.Fields) > 0 {
		_spec.Unique = toq.ctx.Unique != nil && *toq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, toq.driver, _spec)
}

func (toq *TicketOptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ticketoption.Table, ticketoption.Columns, sqlgraph.NewFieldSpec(ticketoption.FieldID, field.TypeString))
	_spec.From = toq.sql
	if unique := toq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if toq.path != nil {
		_spec.Unique = true
	}
	if fields := toq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ticketoption.FieldID)
		for i := range fields {
			if fields[i] != ticketoption.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := toq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := toq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := toq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := toq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (toq *TicketOptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(toq.driver.Dialect())
	t1 := builder.Table(ticketoption.Table)
	columns := toq.ctx.Fields
	if len(columns) == 0 {
		columns = ticketoption.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if toq.sql != nil {
		selector = toq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if toq.ctx.Unique != nil && *toq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range toq.predicates {
		p(selector)
	}
	for _, p := range toq.order {
		p(selector)
	}
	if offset := toq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := toq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TicketOptionGroupBy is the group-by builder for TicketOption entities.
type TicketOptionGroupBy struct {
	selector
	build *TicketOptionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (togb *TicketOptionGroupBy) Aggregate(fns ...AggregateFunc) *TicketOptionGroupBy {
	togb.fns = append(togb.fns, fns...)
	return togb
}

// Scan applies the selector query and scans the result into the given value.
func (togb *TicketOptionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, togb.build.ctx, "GroupBy")
	if err := togb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TicketOptionQuery, *TicketOptionGroupBy](ctx, togb.build, togb, togb.build.inters, v)
}

func (togb *TicketOptionGroupBy) sqlScan(ctx context.Context, root *TicketOptionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(togb.fns))
	for _, fn := range togb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*togb.flds)+len(togb.fns))
		for _, f := range *togb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*togb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := togb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TicketOptionSelect is the builder for selecting fields of TicketOption entities.
type TicketOptionSelect struct {
	*TicketOptionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tos *TicketOptionSelect) Aggregate(fns ...AggregateFunc) *TicketOptionSelect {
	tos.fns = append(tos.fns, fns...)
	return tos
}

// Scan applies the selector query and scans the result into the given value.
func (tos *TicketOptionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tos.ctx, "Select")
	if err := tos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TicketOptionQuery, *TicketOptionSelect](ctx, tos.TicketOptionQuery, tos, tos.inters, v)
}

func (tos *TicketOptionSelect) sqlScan(ctx context.Context, root *TicketOptionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tos.fns))
	for _, fn := range tos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
