// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"placio-app/ent/placeinventory"
	"placio-app/ent/predicate"
	"placio-app/ent/transactionhistory"
	"placio-app/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransactionHistoryQuery is the builder for querying TransactionHistory entities.
type TransactionHistoryQuery struct {
	config
	ctx                *QueryContext
	order              []transactionhistory.OrderOption
	inters             []Interceptor
	predicates         []predicate.TransactionHistory
	withPlaceInventory *PlaceInventoryQuery
	withUser           *UserQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TransactionHistoryQuery builder.
func (thq *TransactionHistoryQuery) Where(ps ...predicate.TransactionHistory) *TransactionHistoryQuery {
	thq.predicates = append(thq.predicates, ps...)
	return thq
}

// Limit the number of records to be returned by this query.
func (thq *TransactionHistoryQuery) Limit(limit int) *TransactionHistoryQuery {
	thq.ctx.Limit = &limit
	return thq
}

// Offset to start from.
func (thq *TransactionHistoryQuery) Offset(offset int) *TransactionHistoryQuery {
	thq.ctx.Offset = &offset
	return thq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (thq *TransactionHistoryQuery) Unique(unique bool) *TransactionHistoryQuery {
	thq.ctx.Unique = &unique
	return thq
}

// Order specifies how the records should be ordered.
func (thq *TransactionHistoryQuery) Order(o ...transactionhistory.OrderOption) *TransactionHistoryQuery {
	thq.order = append(thq.order, o...)
	return thq
}

// QueryPlaceInventory chains the current query on the "place_inventory" edge.
func (thq *TransactionHistoryQuery) QueryPlaceInventory() *PlaceInventoryQuery {
	query := (&PlaceInventoryClient{config: thq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := thq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := thq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transactionhistory.Table, transactionhistory.FieldID, selector),
			sqlgraph.To(placeinventory.Table, placeinventory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transactionhistory.PlaceInventoryTable, transactionhistory.PlaceInventoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(thq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (thq *TransactionHistoryQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: thq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := thq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := thq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(transactionhistory.Table, transactionhistory.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, transactionhistory.UserTable, transactionhistory.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(thq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TransactionHistory entity from the query.
// Returns a *NotFoundError when no TransactionHistory was found.
func (thq *TransactionHistoryQuery) First(ctx context.Context) (*TransactionHistory, error) {
	nodes, err := thq.Limit(1).All(setContextOp(ctx, thq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{transactionhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (thq *TransactionHistoryQuery) FirstX(ctx context.Context) *TransactionHistory {
	node, err := thq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TransactionHistory ID from the query.
// Returns a *NotFoundError when no TransactionHistory ID was found.
func (thq *TransactionHistoryQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = thq.Limit(1).IDs(setContextOp(ctx, thq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{transactionhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (thq *TransactionHistoryQuery) FirstIDX(ctx context.Context) string {
	id, err := thq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TransactionHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TransactionHistory entity is found.
// Returns a *NotFoundError when no TransactionHistory entities are found.
func (thq *TransactionHistoryQuery) Only(ctx context.Context) (*TransactionHistory, error) {
	nodes, err := thq.Limit(2).All(setContextOp(ctx, thq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{transactionhistory.Label}
	default:
		return nil, &NotSingularError{transactionhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (thq *TransactionHistoryQuery) OnlyX(ctx context.Context) *TransactionHistory {
	node, err := thq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TransactionHistory ID in the query.
// Returns a *NotSingularError when more than one TransactionHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (thq *TransactionHistoryQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = thq.Limit(2).IDs(setContextOp(ctx, thq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{transactionhistory.Label}
	default:
		err = &NotSingularError{transactionhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (thq *TransactionHistoryQuery) OnlyIDX(ctx context.Context) string {
	id, err := thq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TransactionHistories.
func (thq *TransactionHistoryQuery) All(ctx context.Context) ([]*TransactionHistory, error) {
	ctx = setContextOp(ctx, thq.ctx, "All")
	if err := thq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TransactionHistory, *TransactionHistoryQuery]()
	return withInterceptors[[]*TransactionHistory](ctx, thq, qr, thq.inters)
}

// AllX is like All, but panics if an error occurs.
func (thq *TransactionHistoryQuery) AllX(ctx context.Context) []*TransactionHistory {
	nodes, err := thq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TransactionHistory IDs.
func (thq *TransactionHistoryQuery) IDs(ctx context.Context) (ids []string, err error) {
	if thq.ctx.Unique == nil && thq.path != nil {
		thq.Unique(true)
	}
	ctx = setContextOp(ctx, thq.ctx, "IDs")
	if err = thq.Select(transactionhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (thq *TransactionHistoryQuery) IDsX(ctx context.Context) []string {
	ids, err := thq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (thq *TransactionHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, thq.ctx, "Count")
	if err := thq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, thq, querierCount[*TransactionHistoryQuery](), thq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (thq *TransactionHistoryQuery) CountX(ctx context.Context) int {
	count, err := thq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (thq *TransactionHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, thq.ctx, "Exist")
	switch _, err := thq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (thq *TransactionHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := thq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TransactionHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (thq *TransactionHistoryQuery) Clone() *TransactionHistoryQuery {
	if thq == nil {
		return nil
	}
	return &TransactionHistoryQuery{
		config:             thq.config,
		ctx:                thq.ctx.Clone(),
		order:              append([]transactionhistory.OrderOption{}, thq.order...),
		inters:             append([]Interceptor{}, thq.inters...),
		predicates:         append([]predicate.TransactionHistory{}, thq.predicates...),
		withPlaceInventory: thq.withPlaceInventory.Clone(),
		withUser:           thq.withUser.Clone(),
		// clone intermediate query.
		sql:  thq.sql.Clone(),
		path: thq.path,
	}
}

// WithPlaceInventory tells the query-builder to eager-load the nodes that are connected to
// the "place_inventory" edge. The optional arguments are used to configure the query builder of the edge.
func (thq *TransactionHistoryQuery) WithPlaceInventory(opts ...func(*PlaceInventoryQuery)) *TransactionHistoryQuery {
	query := (&PlaceInventoryClient{config: thq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	thq.withPlaceInventory = query
	return thq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (thq *TransactionHistoryQuery) WithUser(opts ...func(*UserQuery)) *TransactionHistoryQuery {
	query := (&UserClient{config: thq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	thq.withUser = query
	return thq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TransactionType transactionhistory.TransactionType `json:"transaction_type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TransactionHistory.Query().
//		GroupBy(transactionhistory.FieldTransactionType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (thq *TransactionHistoryQuery) GroupBy(field string, fields ...string) *TransactionHistoryGroupBy {
	thq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TransactionHistoryGroupBy{build: thq}
	grbuild.flds = &thq.ctx.Fields
	grbuild.label = transactionhistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TransactionType transactionhistory.TransactionType `json:"transaction_type,omitempty"`
//	}
//
//	client.TransactionHistory.Query().
//		Select(transactionhistory.FieldTransactionType).
//		Scan(ctx, &v)
func (thq *TransactionHistoryQuery) Select(fields ...string) *TransactionHistorySelect {
	thq.ctx.Fields = append(thq.ctx.Fields, fields...)
	sbuild := &TransactionHistorySelect{TransactionHistoryQuery: thq}
	sbuild.label = transactionhistory.Label
	sbuild.flds, sbuild.scan = &thq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TransactionHistorySelect configured with the given aggregations.
func (thq *TransactionHistoryQuery) Aggregate(fns ...AggregateFunc) *TransactionHistorySelect {
	return thq.Select().Aggregate(fns...)
}

func (thq *TransactionHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range thq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, thq); err != nil {
				return err
			}
		}
	}
	for _, f := range thq.ctx.Fields {
		if !transactionhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if thq.path != nil {
		prev, err := thq.path(ctx)
		if err != nil {
			return err
		}
		thq.sql = prev
	}
	return nil
}

func (thq *TransactionHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TransactionHistory, error) {
	var (
		nodes       = []*TransactionHistory{}
		withFKs     = thq.withFKs
		_spec       = thq.querySpec()
		loadedTypes = [2]bool{
			thq.withPlaceInventory != nil,
			thq.withUser != nil,
		}
	)
	if thq.withPlaceInventory != nil || thq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, transactionhistory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TransactionHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TransactionHistory{config: thq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, thq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := thq.withPlaceInventory; query != nil {
		if err := thq.loadPlaceInventory(ctx, query, nodes, nil,
			func(n *TransactionHistory, e *PlaceInventory) { n.Edges.PlaceInventory = e }); err != nil {
			return nil, err
		}
	}
	if query := thq.withUser; query != nil {
		if err := thq.loadUser(ctx, query, nodes, nil,
			func(n *TransactionHistory, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (thq *TransactionHistoryQuery) loadPlaceInventory(ctx context.Context, query *PlaceInventoryQuery, nodes []*TransactionHistory, init func(*TransactionHistory), assign func(*TransactionHistory, *PlaceInventory)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*TransactionHistory)
	for i := range nodes {
		if nodes[i].place_inventory_transaction_histories == nil {
			continue
		}
		fk := *nodes[i].place_inventory_transaction_histories
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(placeinventory.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "place_inventory_transaction_histories" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (thq *TransactionHistoryQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*TransactionHistory, init func(*TransactionHistory), assign func(*TransactionHistory, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*TransactionHistory)
	for i := range nodes {
		if nodes[i].user_transaction_histories == nil {
			continue
		}
		fk := *nodes[i].user_transaction_histories
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
			return fmt.Errorf(`unexpected foreign-key "user_transaction_histories" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (thq *TransactionHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := thq.querySpec()
	_spec.Node.Columns = thq.ctx.Fields
	if len(thq.ctx.Fields) > 0 {
		_spec.Unique = thq.ctx.Unique != nil && *thq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, thq.driver, _spec)
}

func (thq *TransactionHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(transactionhistory.Table, transactionhistory.Columns, sqlgraph.NewFieldSpec(transactionhistory.FieldID, field.TypeString))
	_spec.From = thq.sql
	if unique := thq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if thq.path != nil {
		_spec.Unique = true
	}
	if fields := thq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transactionhistory.FieldID)
		for i := range fields {
			if fields[i] != transactionhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := thq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := thq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := thq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := thq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (thq *TransactionHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(thq.driver.Dialect())
	t1 := builder.Table(transactionhistory.Table)
	columns := thq.ctx.Fields
	if len(columns) == 0 {
		columns = transactionhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if thq.sql != nil {
		selector = thq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if thq.ctx.Unique != nil && *thq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range thq.predicates {
		p(selector)
	}
	for _, p := range thq.order {
		p(selector)
	}
	if offset := thq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := thq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TransactionHistoryGroupBy is the group-by builder for TransactionHistory entities.
type TransactionHistoryGroupBy struct {
	selector
	build *TransactionHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (thgb *TransactionHistoryGroupBy) Aggregate(fns ...AggregateFunc) *TransactionHistoryGroupBy {
	thgb.fns = append(thgb.fns, fns...)
	return thgb
}

// Scan applies the selector query and scans the result into the given value.
func (thgb *TransactionHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, thgb.build.ctx, "GroupBy")
	if err := thgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TransactionHistoryQuery, *TransactionHistoryGroupBy](ctx, thgb.build, thgb, thgb.build.inters, v)
}

func (thgb *TransactionHistoryGroupBy) sqlScan(ctx context.Context, root *TransactionHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(thgb.fns))
	for _, fn := range thgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*thgb.flds)+len(thgb.fns))
		for _, f := range *thgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*thgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := thgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TransactionHistorySelect is the builder for selecting fields of TransactionHistory entities.
type TransactionHistorySelect struct {
	*TransactionHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ths *TransactionHistorySelect) Aggregate(fns ...AggregateFunc) *TransactionHistorySelect {
	ths.fns = append(ths.fns, fns...)
	return ths
}

// Scan applies the selector query and scans the result into the given value.
func (ths *TransactionHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ths.ctx, "Select")
	if err := ths.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TransactionHistoryQuery, *TransactionHistorySelect](ctx, ths.TransactionHistoryQuery, ths, ths.inters, v)
}

func (ths *TransactionHistorySelect) sqlScan(ctx context.Context, root *TransactionHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ths.fns))
	for _, fn := range ths.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ths.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ths.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
