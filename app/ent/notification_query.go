// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"placio_api/business"
	"placio_api/notification"
	"placio_api/predicate"
	"placio_api/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationQuery is the builder for querying Notification entities.
type NotificationQuery struct {
	config
	ctx                 *QueryContext
	order               []notification.OrderOption
	inters              []Interceptor
	predicates          []predicate.Notification
	withUser            *UserQuery
	withBusinessAccount *BusinessQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotificationQuery builder.
func (nq *NotificationQuery) Where(ps ...predicate.Notification) *NotificationQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NotificationQuery) Limit(limit int) *NotificationQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NotificationQuery) Offset(offset int) *NotificationQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NotificationQuery) Unique(unique bool) *NotificationQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NotificationQuery) Order(o ...notification.OrderOption) *NotificationQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryUser chains the current query on the "user" edge.
func (nq *NotificationQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, notification.UserTable, notification.UserPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusinessAccount chains the current query on the "business_account" edge.
func (nq *NotificationQuery) QueryBusinessAccount() *BusinessQuery {
	query := (&BusinessClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notification.Table, notification.FieldID, selector),
			sqlgraph.To(business.Table, business.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, notification.BusinessAccountTable, notification.BusinessAccountPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Notification entity from the query.
// Returns a *NotFoundError when no Notification was found.
func (nq *NotificationQuery) First(ctx context.Context) (*Notification, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notification.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NotificationQuery) FirstX(ctx context.Context) *Notification {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Notification ID from the query.
// Returns a *NotFoundError when no Notification ID was found.
func (nq *NotificationQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notification.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NotificationQuery) FirstIDX(ctx context.Context) string {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Notification entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Notification entity is found.
// Returns a *NotFoundError when no Notification entities are found.
func (nq *NotificationQuery) Only(ctx context.Context) (*Notification, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notification.Label}
	default:
		return nil, &NotSingularError{notification.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NotificationQuery) OnlyX(ctx context.Context) *Notification {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Notification ID in the query.
// Returns a *NotSingularError when more than one Notification ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NotificationQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notification.Label}
	default:
		err = &NotSingularError{notification.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NotificationQuery) OnlyIDX(ctx context.Context) string {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Notifications.
func (nq *NotificationQuery) All(ctx context.Context) ([]*Notification, error) {
	ctx = setContextOp(ctx, nq.ctx, "All")
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Notification, *NotificationQuery]()
	return withInterceptors[[]*Notification](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NotificationQuery) AllX(ctx context.Context) []*Notification {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Notification IDs.
func (nq *NotificationQuery) IDs(ctx context.Context) (ids []string, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, "IDs")
	if err = nq.Select(notification.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NotificationQuery) IDsX(ctx context.Context) []string {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NotificationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, "Count")
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NotificationQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NotificationQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NotificationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, "Exist")
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("placio_api: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NotificationQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotificationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NotificationQuery) Clone() *NotificationQuery {
	if nq == nil {
		return nil
	}
	return &NotificationQuery{
		config:              nq.config,
		ctx:                 nq.ctx.Clone(),
		order:               append([]notification.OrderOption{}, nq.order...),
		inters:              append([]Interceptor{}, nq.inters...),
		predicates:          append([]predicate.Notification{}, nq.predicates...),
		withUser:            nq.withUser.Clone(),
		withBusinessAccount: nq.withBusinessAccount.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotificationQuery) WithUser(opts ...func(*UserQuery)) *NotificationQuery {
	query := (&UserClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withUser = query
	return nq
}

// WithBusinessAccount tells the query-builder to eager-load the nodes that are connected to
// the "business_account" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotificationQuery) WithBusinessAccount(opts ...func(*BusinessQuery)) *NotificationQuery {
	query := (&BusinessClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withBusinessAccount = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Notification.Query().
//		GroupBy(notification.FieldTitle).
//		Aggregate(placio_api.Count()).
//		Scan(ctx, &v)
func (nq *NotificationQuery) GroupBy(field string, fields ...string) *NotificationGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NotificationGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = notification.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Notification.Query().
//		Select(notification.FieldTitle).
//		Scan(ctx, &v)
func (nq *NotificationQuery) Select(fields ...string) *NotificationSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NotificationSelect{NotificationQuery: nq}
	sbuild.label = notification.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NotificationSelect configured with the given aggregations.
func (nq *NotificationQuery) Aggregate(fns ...AggregateFunc) *NotificationSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NotificationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("placio_api: uninitialized interceptor (forgotten import placio_api/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !notification.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NotificationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Notification, error) {
	var (
		nodes       = []*Notification{}
		_spec       = nq.querySpec()
		loadedTypes = [2]bool{
			nq.withUser != nil,
			nq.withBusinessAccount != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Notification).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Notification{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withUser; query != nil {
		if err := nq.loadUser(ctx, query, nodes,
			func(n *Notification) { n.Edges.User = []*User{} },
			func(n *Notification, e *User) { n.Edges.User = append(n.Edges.User, e) }); err != nil {
			return nil, err
		}
	}
	if query := nq.withBusinessAccount; query != nil {
		if err := nq.loadBusinessAccount(ctx, query, nodes,
			func(n *Notification) { n.Edges.BusinessAccount = []*Business{} },
			func(n *Notification, e *Business) { n.Edges.BusinessAccount = append(n.Edges.BusinessAccount, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NotificationQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Notification, init func(*Notification), assign func(*Notification, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Notification)
	nids := make(map[string]map[*Notification]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(notification.UserTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(notification.UserPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(notification.UserPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(notification.UserPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*Notification]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "user" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (nq *NotificationQuery) loadBusinessAccount(ctx context.Context, query *BusinessQuery, nodes []*Notification, init func(*Notification), assign func(*Notification, *Business)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Notification)
	nids := make(map[string]map[*Notification]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(notification.BusinessAccountTable)
		s.Join(joinT).On(s.C(business.FieldID), joinT.C(notification.BusinessAccountPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(notification.BusinessAccountPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(notification.BusinessAccountPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*Notification]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Business](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "business_account" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (nq *NotificationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NotificationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(notification.Table, notification.Columns, sqlgraph.NewFieldSpec(notification.FieldID, field.TypeString))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notification.FieldID)
		for i := range fields {
			if fields[i] != notification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NotificationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(notification.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = notification.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NotificationGroupBy is the group-by builder for Notification entities.
type NotificationGroupBy struct {
	selector
	build *NotificationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NotificationGroupBy) Aggregate(fns ...AggregateFunc) *NotificationGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NotificationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, "GroupBy")
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotificationQuery, *NotificationGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NotificationGroupBy) sqlScan(ctx context.Context, root *NotificationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NotificationSelect is the builder for selecting fields of Notification entities.
type NotificationSelect struct {
	*NotificationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NotificationSelect) Aggregate(fns ...AggregateFunc) *NotificationSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NotificationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, "Select")
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotificationQuery, *NotificationSelect](ctx, ns.NotificationQuery, ns, ns.inters, v)
}

func (ns *NotificationSelect) sqlScan(ctx context.Context, root *NotificationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
