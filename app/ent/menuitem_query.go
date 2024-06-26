// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"placio-app/ent/media"
	"placio-app/ent/menu"
	"placio-app/ent/menuitem"
	"placio-app/ent/orderitem"
	"placio-app/ent/placeinventory"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuItemQuery is the builder for querying MenuItem entities.
type MenuItemQuery struct {
	config
	ctx            *QueryContext
	order          []menuitem.OrderOption
	inters         []Interceptor
	predicates     []predicate.MenuItem
	withMenu       *MenuQuery
	withInventory  *PlaceInventoryQuery
	withMedia      *MediaQuery
	withOrderItems *OrderItemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MenuItemQuery builder.
func (miq *MenuItemQuery) Where(ps ...predicate.MenuItem) *MenuItemQuery {
	miq.predicates = append(miq.predicates, ps...)
	return miq
}

// Limit the number of records to be returned by this query.
func (miq *MenuItemQuery) Limit(limit int) *MenuItemQuery {
	miq.ctx.Limit = &limit
	return miq
}

// Offset to start from.
func (miq *MenuItemQuery) Offset(offset int) *MenuItemQuery {
	miq.ctx.Offset = &offset
	return miq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (miq *MenuItemQuery) Unique(unique bool) *MenuItemQuery {
	miq.ctx.Unique = &unique
	return miq
}

// Order specifies how the records should be ordered.
func (miq *MenuItemQuery) Order(o ...menuitem.OrderOption) *MenuItemQuery {
	miq.order = append(miq.order, o...)
	return miq
}

// QueryMenu chains the current query on the "menu" edge.
func (miq *MenuItemQuery) QueryMenu() *MenuQuery {
	query := (&MenuClient{config: miq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := miq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(menuitem.Table, menuitem.FieldID, selector),
			sqlgraph.To(menu.Table, menu.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, menuitem.MenuTable, menuitem.MenuPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(miq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInventory chains the current query on the "inventory" edge.
func (miq *MenuItemQuery) QueryInventory() *PlaceInventoryQuery {
	query := (&PlaceInventoryClient{config: miq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := miq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(menuitem.Table, menuitem.FieldID, selector),
			sqlgraph.To(placeinventory.Table, placeinventory.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, menuitem.InventoryTable, menuitem.InventoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(miq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMedia chains the current query on the "media" edge.
func (miq *MenuItemQuery) QueryMedia() *MediaQuery {
	query := (&MediaClient{config: miq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := miq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(menuitem.Table, menuitem.FieldID, selector),
			sqlgraph.To(media.Table, media.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, menuitem.MediaTable, menuitem.MediaColumn),
		)
		fromU = sqlgraph.SetNeighbors(miq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderItems chains the current query on the "order_items" edge.
func (miq *MenuItemQuery) QueryOrderItems() *OrderItemQuery {
	query := (&OrderItemClient{config: miq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := miq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(menuitem.Table, menuitem.FieldID, selector),
			sqlgraph.To(orderitem.Table, orderitem.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, menuitem.OrderItemsTable, menuitem.OrderItemsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(miq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MenuItem entity from the query.
// Returns a *NotFoundError when no MenuItem was found.
func (miq *MenuItemQuery) First(ctx context.Context) (*MenuItem, error) {
	nodes, err := miq.Limit(1).All(setContextOp(ctx, miq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{menuitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (miq *MenuItemQuery) FirstX(ctx context.Context) *MenuItem {
	node, err := miq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MenuItem ID from the query.
// Returns a *NotFoundError when no MenuItem ID was found.
func (miq *MenuItemQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = miq.Limit(1).IDs(setContextOp(ctx, miq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{menuitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (miq *MenuItemQuery) FirstIDX(ctx context.Context) string {
	id, err := miq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MenuItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MenuItem entity is found.
// Returns a *NotFoundError when no MenuItem entities are found.
func (miq *MenuItemQuery) Only(ctx context.Context) (*MenuItem, error) {
	nodes, err := miq.Limit(2).All(setContextOp(ctx, miq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{menuitem.Label}
	default:
		return nil, &NotSingularError{menuitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (miq *MenuItemQuery) OnlyX(ctx context.Context) *MenuItem {
	node, err := miq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MenuItem ID in the query.
// Returns a *NotSingularError when more than one MenuItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (miq *MenuItemQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = miq.Limit(2).IDs(setContextOp(ctx, miq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{menuitem.Label}
	default:
		err = &NotSingularError{menuitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (miq *MenuItemQuery) OnlyIDX(ctx context.Context) string {
	id, err := miq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MenuItems.
func (miq *MenuItemQuery) All(ctx context.Context) ([]*MenuItem, error) {
	ctx = setContextOp(ctx, miq.ctx, "All")
	if err := miq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MenuItem, *MenuItemQuery]()
	return withInterceptors[[]*MenuItem](ctx, miq, qr, miq.inters)
}

// AllX is like All, but panics if an error occurs.
func (miq *MenuItemQuery) AllX(ctx context.Context) []*MenuItem {
	nodes, err := miq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MenuItem IDs.
func (miq *MenuItemQuery) IDs(ctx context.Context) (ids []string, err error) {
	if miq.ctx.Unique == nil && miq.path != nil {
		miq.Unique(true)
	}
	ctx = setContextOp(ctx, miq.ctx, "IDs")
	if err = miq.Select(menuitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (miq *MenuItemQuery) IDsX(ctx context.Context) []string {
	ids, err := miq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (miq *MenuItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, miq.ctx, "Count")
	if err := miq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, miq, querierCount[*MenuItemQuery](), miq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (miq *MenuItemQuery) CountX(ctx context.Context) int {
	count, err := miq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (miq *MenuItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, miq.ctx, "Exist")
	switch _, err := miq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (miq *MenuItemQuery) ExistX(ctx context.Context) bool {
	exist, err := miq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MenuItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (miq *MenuItemQuery) Clone() *MenuItemQuery {
	if miq == nil {
		return nil
	}
	return &MenuItemQuery{
		config:         miq.config,
		ctx:            miq.ctx.Clone(),
		order:          append([]menuitem.OrderOption{}, miq.order...),
		inters:         append([]Interceptor{}, miq.inters...),
		predicates:     append([]predicate.MenuItem{}, miq.predicates...),
		withMenu:       miq.withMenu.Clone(),
		withInventory:  miq.withInventory.Clone(),
		withMedia:      miq.withMedia.Clone(),
		withOrderItems: miq.withOrderItems.Clone(),
		// clone intermediate query.
		sql:  miq.sql.Clone(),
		path: miq.path,
	}
}

// WithMenu tells the query-builder to eager-load the nodes that are connected to
// the "menu" edge. The optional arguments are used to configure the query builder of the edge.
func (miq *MenuItemQuery) WithMenu(opts ...func(*MenuQuery)) *MenuItemQuery {
	query := (&MenuClient{config: miq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	miq.withMenu = query
	return miq
}

// WithInventory tells the query-builder to eager-load the nodes that are connected to
// the "inventory" edge. The optional arguments are used to configure the query builder of the edge.
func (miq *MenuItemQuery) WithInventory(opts ...func(*PlaceInventoryQuery)) *MenuItemQuery {
	query := (&PlaceInventoryClient{config: miq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	miq.withInventory = query
	return miq
}

// WithMedia tells the query-builder to eager-load the nodes that are connected to
// the "media" edge. The optional arguments are used to configure the query builder of the edge.
func (miq *MenuItemQuery) WithMedia(opts ...func(*MediaQuery)) *MenuItemQuery {
	query := (&MediaClient{config: miq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	miq.withMedia = query
	return miq
}

// WithOrderItems tells the query-builder to eager-load the nodes that are connected to
// the "order_items" edge. The optional arguments are used to configure the query builder of the edge.
func (miq *MenuItemQuery) WithOrderItems(opts ...func(*OrderItemQuery)) *MenuItemQuery {
	query := (&OrderItemClient{config: miq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	miq.withOrderItems = query
	return miq
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
//	client.MenuItem.Query().
//		GroupBy(menuitem.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (miq *MenuItemQuery) GroupBy(field string, fields ...string) *MenuItemGroupBy {
	miq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MenuItemGroupBy{build: miq}
	grbuild.flds = &miq.ctx.Fields
	grbuild.label = menuitem.Label
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
//	client.MenuItem.Query().
//		Select(menuitem.FieldName).
//		Scan(ctx, &v)
func (miq *MenuItemQuery) Select(fields ...string) *MenuItemSelect {
	miq.ctx.Fields = append(miq.ctx.Fields, fields...)
	sbuild := &MenuItemSelect{MenuItemQuery: miq}
	sbuild.label = menuitem.Label
	sbuild.flds, sbuild.scan = &miq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MenuItemSelect configured with the given aggregations.
func (miq *MenuItemQuery) Aggregate(fns ...AggregateFunc) *MenuItemSelect {
	return miq.Select().Aggregate(fns...)
}

func (miq *MenuItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range miq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, miq); err != nil {
				return err
			}
		}
	}
	for _, f := range miq.ctx.Fields {
		if !menuitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if miq.path != nil {
		prev, err := miq.path(ctx)
		if err != nil {
			return err
		}
		miq.sql = prev
	}
	return nil
}

func (miq *MenuItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MenuItem, error) {
	var (
		nodes       = []*MenuItem{}
		_spec       = miq.querySpec()
		loadedTypes = [4]bool{
			miq.withMenu != nil,
			miq.withInventory != nil,
			miq.withMedia != nil,
			miq.withOrderItems != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MenuItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MenuItem{config: miq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, miq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := miq.withMenu; query != nil {
		if err := miq.loadMenu(ctx, query, nodes,
			func(n *MenuItem) { n.Edges.Menu = []*Menu{} },
			func(n *MenuItem, e *Menu) { n.Edges.Menu = append(n.Edges.Menu, e) }); err != nil {
			return nil, err
		}
	}
	if query := miq.withInventory; query != nil {
		if err := miq.loadInventory(ctx, query, nodes, nil,
			func(n *MenuItem, e *PlaceInventory) { n.Edges.Inventory = e }); err != nil {
			return nil, err
		}
	}
	if query := miq.withMedia; query != nil {
		if err := miq.loadMedia(ctx, query, nodes,
			func(n *MenuItem) { n.Edges.Media = []*Media{} },
			func(n *MenuItem, e *Media) { n.Edges.Media = append(n.Edges.Media, e) }); err != nil {
			return nil, err
		}
	}
	if query := miq.withOrderItems; query != nil {
		if err := miq.loadOrderItems(ctx, query, nodes,
			func(n *MenuItem) { n.Edges.OrderItems = []*OrderItem{} },
			func(n *MenuItem, e *OrderItem) { n.Edges.OrderItems = append(n.Edges.OrderItems, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (miq *MenuItemQuery) loadMenu(ctx context.Context, query *MenuQuery, nodes []*MenuItem, init func(*MenuItem), assign func(*MenuItem, *Menu)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*MenuItem)
	nids := make(map[string]map[*MenuItem]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(menuitem.MenuTable)
		s.Join(joinT).On(s.C(menu.FieldID), joinT.C(menuitem.MenuPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(menuitem.MenuPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(menuitem.MenuPrimaryKey[1]))
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
					nids[inValue] = map[*MenuItem]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Menu](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "menu" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (miq *MenuItemQuery) loadInventory(ctx context.Context, query *PlaceInventoryQuery, nodes []*MenuItem, init func(*MenuItem), assign func(*MenuItem, *PlaceInventory)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*MenuItem)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.PlaceInventory(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(menuitem.InventoryColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.menu_item_inventory
		if fk == nil {
			return fmt.Errorf(`foreign-key "menu_item_inventory" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "menu_item_inventory" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (miq *MenuItemQuery) loadMedia(ctx context.Context, query *MediaQuery, nodes []*MenuItem, init func(*MenuItem), assign func(*MenuItem, *Media)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*MenuItem)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Media(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(menuitem.MediaColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.menu_item_media
		if fk == nil {
			return fmt.Errorf(`foreign-key "menu_item_media" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "menu_item_media" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (miq *MenuItemQuery) loadOrderItems(ctx context.Context, query *OrderItemQuery, nodes []*MenuItem, init func(*MenuItem), assign func(*MenuItem, *OrderItem)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*MenuItem)
	nids := make(map[string]map[*MenuItem]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(menuitem.OrderItemsTable)
		s.Join(joinT).On(s.C(orderitem.FieldID), joinT.C(menuitem.OrderItemsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(menuitem.OrderItemsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(menuitem.OrderItemsPrimaryKey[0]))
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
					nids[inValue] = map[*MenuItem]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*OrderItem](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "order_items" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (miq *MenuItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := miq.querySpec()
	_spec.Node.Columns = miq.ctx.Fields
	if len(miq.ctx.Fields) > 0 {
		_spec.Unique = miq.ctx.Unique != nil && *miq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, miq.driver, _spec)
}

func (miq *MenuItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(menuitem.Table, menuitem.Columns, sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeString))
	_spec.From = miq.sql
	if unique := miq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if miq.path != nil {
		_spec.Unique = true
	}
	if fields := miq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menuitem.FieldID)
		for i := range fields {
			if fields[i] != menuitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := miq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := miq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := miq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := miq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (miq *MenuItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(miq.driver.Dialect())
	t1 := builder.Table(menuitem.Table)
	columns := miq.ctx.Fields
	if len(columns) == 0 {
		columns = menuitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if miq.sql != nil {
		selector = miq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if miq.ctx.Unique != nil && *miq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range miq.predicates {
		p(selector)
	}
	for _, p := range miq.order {
		p(selector)
	}
	if offset := miq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := miq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MenuItemGroupBy is the group-by builder for MenuItem entities.
type MenuItemGroupBy struct {
	selector
	build *MenuItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (migb *MenuItemGroupBy) Aggregate(fns ...AggregateFunc) *MenuItemGroupBy {
	migb.fns = append(migb.fns, fns...)
	return migb
}

// Scan applies the selector query and scans the result into the given value.
func (migb *MenuItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, migb.build.ctx, "GroupBy")
	if err := migb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MenuItemQuery, *MenuItemGroupBy](ctx, migb.build, migb, migb.build.inters, v)
}

func (migb *MenuItemGroupBy) sqlScan(ctx context.Context, root *MenuItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(migb.fns))
	for _, fn := range migb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*migb.flds)+len(migb.fns))
		for _, f := range *migb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*migb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := migb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MenuItemSelect is the builder for selecting fields of MenuItem entities.
type MenuItemSelect struct {
	*MenuItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mis *MenuItemSelect) Aggregate(fns ...AggregateFunc) *MenuItemSelect {
	mis.fns = append(mis.fns, fns...)
	return mis
}

// Scan applies the selector query and scans the result into the given value.
func (mis *MenuItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mis.ctx, "Select")
	if err := mis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MenuItemQuery, *MenuItemSelect](ctx, mis.MenuItemQuery, mis, mis.inters, v)
}

func (mis *MenuItemSelect) sqlScan(ctx context.Context, root *MenuItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mis.fns))
	for _, fn := range mis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
