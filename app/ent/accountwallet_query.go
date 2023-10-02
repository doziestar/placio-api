// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"placio_api/accountwallet"
	"placio_api/business"
	"placio_api/predicate"
	"placio_api/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountWalletQuery is the builder for querying AccountWallet entities.
type AccountWalletQuery struct {
	config
	ctx          *QueryContext
	order        []accountwallet.OrderOption
	inters       []Interceptor
	predicates   []predicate.AccountWallet
	withUser     *UserQuery
	withBusiness *BusinessQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountWalletQuery builder.
func (awq *AccountWalletQuery) Where(ps ...predicate.AccountWallet) *AccountWalletQuery {
	awq.predicates = append(awq.predicates, ps...)
	return awq
}

// Limit the number of records to be returned by this query.
func (awq *AccountWalletQuery) Limit(limit int) *AccountWalletQuery {
	awq.ctx.Limit = &limit
	return awq
}

// Offset to start from.
func (awq *AccountWalletQuery) Offset(offset int) *AccountWalletQuery {
	awq.ctx.Offset = &offset
	return awq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (awq *AccountWalletQuery) Unique(unique bool) *AccountWalletQuery {
	awq.ctx.Unique = &unique
	return awq
}

// Order specifies how the records should be ordered.
func (awq *AccountWalletQuery) Order(o ...accountwallet.OrderOption) *AccountWalletQuery {
	awq.order = append(awq.order, o...)
	return awq
}

// QueryUser chains the current query on the "user" edge.
func (awq *AccountWalletQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: awq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := awq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := awq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accountwallet.Table, accountwallet.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, accountwallet.UserTable, accountwallet.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(awq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusiness chains the current query on the "business" edge.
func (awq *AccountWalletQuery) QueryBusiness() *BusinessQuery {
	query := (&BusinessClient{config: awq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := awq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := awq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accountwallet.Table, accountwallet.FieldID, selector),
			sqlgraph.To(business.Table, business.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, accountwallet.BusinessTable, accountwallet.BusinessColumn),
		)
		fromU = sqlgraph.SetNeighbors(awq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AccountWallet entity from the query.
// Returns a *NotFoundError when no AccountWallet was found.
func (awq *AccountWalletQuery) First(ctx context.Context) (*AccountWallet, error) {
	nodes, err := awq.Limit(1).All(setContextOp(ctx, awq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accountwallet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (awq *AccountWalletQuery) FirstX(ctx context.Context) *AccountWallet {
	node, err := awq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccountWallet ID from the query.
// Returns a *NotFoundError when no AccountWallet ID was found.
func (awq *AccountWalletQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = awq.Limit(1).IDs(setContextOp(ctx, awq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accountwallet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (awq *AccountWalletQuery) FirstIDX(ctx context.Context) string {
	id, err := awq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccountWallet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AccountWallet entity is found.
// Returns a *NotFoundError when no AccountWallet entities are found.
func (awq *AccountWalletQuery) Only(ctx context.Context) (*AccountWallet, error) {
	nodes, err := awq.Limit(2).All(setContextOp(ctx, awq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accountwallet.Label}
	default:
		return nil, &NotSingularError{accountwallet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (awq *AccountWalletQuery) OnlyX(ctx context.Context) *AccountWallet {
	node, err := awq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccountWallet ID in the query.
// Returns a *NotSingularError when more than one AccountWallet ID is found.
// Returns a *NotFoundError when no entities are found.
func (awq *AccountWalletQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = awq.Limit(2).IDs(setContextOp(ctx, awq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accountwallet.Label}
	default:
		err = &NotSingularError{accountwallet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (awq *AccountWalletQuery) OnlyIDX(ctx context.Context) string {
	id, err := awq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccountWallets.
func (awq *AccountWalletQuery) All(ctx context.Context) ([]*AccountWallet, error) {
	ctx = setContextOp(ctx, awq.ctx, "All")
	if err := awq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AccountWallet, *AccountWalletQuery]()
	return withInterceptors[[]*AccountWallet](ctx, awq, qr, awq.inters)
}

// AllX is like All, but panics if an error occurs.
func (awq *AccountWalletQuery) AllX(ctx context.Context) []*AccountWallet {
	nodes, err := awq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccountWallet IDs.
func (awq *AccountWalletQuery) IDs(ctx context.Context) (ids []string, err error) {
	if awq.ctx.Unique == nil && awq.path != nil {
		awq.Unique(true)
	}
	ctx = setContextOp(ctx, awq.ctx, "IDs")
	if err = awq.Select(accountwallet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (awq *AccountWalletQuery) IDsX(ctx context.Context) []string {
	ids, err := awq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (awq *AccountWalletQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, awq.ctx, "Count")
	if err := awq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, awq, querierCount[*AccountWalletQuery](), awq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (awq *AccountWalletQuery) CountX(ctx context.Context) int {
	count, err := awq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (awq *AccountWalletQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, awq.ctx, "Exist")
	switch _, err := awq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("placio_api: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (awq *AccountWalletQuery) ExistX(ctx context.Context) bool {
	exist, err := awq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountWalletQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (awq *AccountWalletQuery) Clone() *AccountWalletQuery {
	if awq == nil {
		return nil
	}
	return &AccountWalletQuery{
		config:       awq.config,
		ctx:          awq.ctx.Clone(),
		order:        append([]accountwallet.OrderOption{}, awq.order...),
		inters:       append([]Interceptor{}, awq.inters...),
		predicates:   append([]predicate.AccountWallet{}, awq.predicates...),
		withUser:     awq.withUser.Clone(),
		withBusiness: awq.withBusiness.Clone(),
		// clone intermediate query.
		sql:  awq.sql.Clone(),
		path: awq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (awq *AccountWalletQuery) WithUser(opts ...func(*UserQuery)) *AccountWalletQuery {
	query := (&UserClient{config: awq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	awq.withUser = query
	return awq
}

// WithBusiness tells the query-builder to eager-load the nodes that are connected to
// the "business" edge. The optional arguments are used to configure the query builder of the edge.
func (awq *AccountWalletQuery) WithBusiness(opts ...func(*BusinessQuery)) *AccountWalletQuery {
	query := (&BusinessClient{config: awq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	awq.withBusiness = query
	return awq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AccountWallet.Query().
//		GroupBy(accountwallet.FieldUserID).
//		Aggregate(placio_api.Count()).
//		Scan(ctx, &v)
func (awq *AccountWalletQuery) GroupBy(field string, fields ...string) *AccountWalletGroupBy {
	awq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccountWalletGroupBy{build: awq}
	grbuild.flds = &awq.ctx.Fields
	grbuild.label = accountwallet.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//	}
//
//	client.AccountWallet.Query().
//		Select(accountwallet.FieldUserID).
//		Scan(ctx, &v)
func (awq *AccountWalletQuery) Select(fields ...string) *AccountWalletSelect {
	awq.ctx.Fields = append(awq.ctx.Fields, fields...)
	sbuild := &AccountWalletSelect{AccountWalletQuery: awq}
	sbuild.label = accountwallet.Label
	sbuild.flds, sbuild.scan = &awq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccountWalletSelect configured with the given aggregations.
func (awq *AccountWalletQuery) Aggregate(fns ...AggregateFunc) *AccountWalletSelect {
	return awq.Select().Aggregate(fns...)
}

func (awq *AccountWalletQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range awq.inters {
		if inter == nil {
			return fmt.Errorf("placio_api: uninitialized interceptor (forgotten import placio_api/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, awq); err != nil {
				return err
			}
		}
	}
	for _, f := range awq.ctx.Fields {
		if !accountwallet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
		}
	}
	if awq.path != nil {
		prev, err := awq.path(ctx)
		if err != nil {
			return err
		}
		awq.sql = prev
	}
	return nil
}

func (awq *AccountWalletQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AccountWallet, error) {
	var (
		nodes       = []*AccountWallet{}
		withFKs     = awq.withFKs
		_spec       = awq.querySpec()
		loadedTypes = [2]bool{
			awq.withUser != nil,
			awq.withBusiness != nil,
		}
	)
	if awq.withUser != nil || awq.withBusiness != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, accountwallet.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AccountWallet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AccountWallet{config: awq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, awq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := awq.withUser; query != nil {
		if err := awq.loadUser(ctx, query, nodes, nil,
			func(n *AccountWallet, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := awq.withBusiness; query != nil {
		if err := awq.loadBusiness(ctx, query, nodes, nil,
			func(n *AccountWallet, e *Business) { n.Edges.Business = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (awq *AccountWalletQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*AccountWallet, init func(*AccountWallet), assign func(*AccountWallet, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*AccountWallet)
	for i := range nodes {
		if nodes[i].user_wallet == nil {
			continue
		}
		fk := *nodes[i].user_wallet
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
			return fmt.Errorf(`unexpected foreign-key "user_wallet" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (awq *AccountWalletQuery) loadBusiness(ctx context.Context, query *BusinessQuery, nodes []*AccountWallet, init func(*AccountWallet), assign func(*AccountWallet, *Business)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*AccountWallet)
	for i := range nodes {
		if nodes[i].business_wallet == nil {
			continue
		}
		fk := *nodes[i].business_wallet
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
			return fmt.Errorf(`unexpected foreign-key "business_wallet" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (awq *AccountWalletQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := awq.querySpec()
	_spec.Node.Columns = awq.ctx.Fields
	if len(awq.ctx.Fields) > 0 {
		_spec.Unique = awq.ctx.Unique != nil && *awq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, awq.driver, _spec)
}

func (awq *AccountWalletQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(accountwallet.Table, accountwallet.Columns, sqlgraph.NewFieldSpec(accountwallet.FieldID, field.TypeString))
	_spec.From = awq.sql
	if unique := awq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if awq.path != nil {
		_spec.Unique = true
	}
	if fields := awq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountwallet.FieldID)
		for i := range fields {
			if fields[i] != accountwallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := awq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := awq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := awq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := awq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (awq *AccountWalletQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(awq.driver.Dialect())
	t1 := builder.Table(accountwallet.Table)
	columns := awq.ctx.Fields
	if len(columns) == 0 {
		columns = accountwallet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if awq.sql != nil {
		selector = awq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if awq.ctx.Unique != nil && *awq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range awq.predicates {
		p(selector)
	}
	for _, p := range awq.order {
		p(selector)
	}
	if offset := awq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := awq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccountWalletGroupBy is the group-by builder for AccountWallet entities.
type AccountWalletGroupBy struct {
	selector
	build *AccountWalletQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (awgb *AccountWalletGroupBy) Aggregate(fns ...AggregateFunc) *AccountWalletGroupBy {
	awgb.fns = append(awgb.fns, fns...)
	return awgb
}

// Scan applies the selector query and scans the result into the given value.
func (awgb *AccountWalletGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, awgb.build.ctx, "GroupBy")
	if err := awgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountWalletQuery, *AccountWalletGroupBy](ctx, awgb.build, awgb, awgb.build.inters, v)
}

func (awgb *AccountWalletGroupBy) sqlScan(ctx context.Context, root *AccountWalletQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(awgb.fns))
	for _, fn := range awgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*awgb.flds)+len(awgb.fns))
		for _, f := range *awgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*awgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := awgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccountWalletSelect is the builder for selecting fields of AccountWallet entities.
type AccountWalletSelect struct {
	*AccountWalletQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aws *AccountWalletSelect) Aggregate(fns ...AggregateFunc) *AccountWalletSelect {
	aws.fns = append(aws.fns, fns...)
	return aws
}

// Scan applies the selector query and scans the result into the given value.
func (aws *AccountWalletSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aws.ctx, "Select")
	if err := aws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountWalletQuery, *AccountWalletSelect](ctx, aws.AccountWalletQuery, aws, aws.inters, v)
}

func (aws *AccountWalletSelect) sqlScan(ctx context.Context, root *AccountWalletQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aws.fns))
	for _, fn := range aws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
