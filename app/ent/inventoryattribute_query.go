// Code generated by ent, DO NOT EDIT.

package placio_api

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"placio_api/inventoryattribute"
	"placio_api/inventorytype"
	"placio_api/placeinventoryattribute"
	"placio_api/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InventoryAttributeQuery is the builder for querying InventoryAttribute entities.
type InventoryAttributeQuery struct {
	config
	ctx                          *QueryContext
	order                        []inventoryattribute.OrderOption
	inters                       []Interceptor
	predicates                   []predicate.InventoryAttribute
	withInventoryType            *InventoryTypeQuery
	withPlaceInventoryAttributes *PlaceInventoryAttributeQuery
	withFKs                      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InventoryAttributeQuery builder.
func (iaq *InventoryAttributeQuery) Where(ps ...predicate.InventoryAttribute) *InventoryAttributeQuery {
	iaq.predicates = append(iaq.predicates, ps...)
	return iaq
}

// Limit the number of records to be returned by this query.
func (iaq *InventoryAttributeQuery) Limit(limit int) *InventoryAttributeQuery {
	iaq.ctx.Limit = &limit
	return iaq
}

// Offset to start from.
func (iaq *InventoryAttributeQuery) Offset(offset int) *InventoryAttributeQuery {
	iaq.ctx.Offset = &offset
	return iaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iaq *InventoryAttributeQuery) Unique(unique bool) *InventoryAttributeQuery {
	iaq.ctx.Unique = &unique
	return iaq
}

// Order specifies how the records should be ordered.
func (iaq *InventoryAttributeQuery) Order(o ...inventoryattribute.OrderOption) *InventoryAttributeQuery {
	iaq.order = append(iaq.order, o...)
	return iaq
}

// QueryInventoryType chains the current query on the "inventory_type" edge.
func (iaq *InventoryAttributeQuery) QueryInventoryType() *InventoryTypeQuery {
	query := (&InventoryTypeClient{config: iaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(inventoryattribute.Table, inventoryattribute.FieldID, selector),
			sqlgraph.To(inventorytype.Table, inventorytype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, inventoryattribute.InventoryTypeTable, inventoryattribute.InventoryTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(iaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlaceInventoryAttributes chains the current query on the "place_inventory_attributes" edge.
func (iaq *InventoryAttributeQuery) QueryPlaceInventoryAttributes() *PlaceInventoryAttributeQuery {
	query := (&PlaceInventoryAttributeClient{config: iaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(inventoryattribute.Table, inventoryattribute.FieldID, selector),
			sqlgraph.To(placeinventoryattribute.Table, placeinventoryattribute.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, inventoryattribute.PlaceInventoryAttributesTable, inventoryattribute.PlaceInventoryAttributesColumn),
		)
		fromU = sqlgraph.SetNeighbors(iaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first InventoryAttribute entity from the query.
// Returns a *NotFoundError when no InventoryAttribute was found.
func (iaq *InventoryAttributeQuery) First(ctx context.Context) (*InventoryAttribute, error) {
	nodes, err := iaq.Limit(1).All(setContextOp(ctx, iaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{inventoryattribute.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iaq *InventoryAttributeQuery) FirstX(ctx context.Context) *InventoryAttribute {
	node, err := iaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first InventoryAttribute ID from the query.
// Returns a *NotFoundError when no InventoryAttribute ID was found.
func (iaq *InventoryAttributeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = iaq.Limit(1).IDs(setContextOp(ctx, iaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{inventoryattribute.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iaq *InventoryAttributeQuery) FirstIDX(ctx context.Context) string {
	id, err := iaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single InventoryAttribute entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one InventoryAttribute entity is found.
// Returns a *NotFoundError when no InventoryAttribute entities are found.
func (iaq *InventoryAttributeQuery) Only(ctx context.Context) (*InventoryAttribute, error) {
	nodes, err := iaq.Limit(2).All(setContextOp(ctx, iaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{inventoryattribute.Label}
	default:
		return nil, &NotSingularError{inventoryattribute.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iaq *InventoryAttributeQuery) OnlyX(ctx context.Context) *InventoryAttribute {
	node, err := iaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only InventoryAttribute ID in the query.
// Returns a *NotSingularError when more than one InventoryAttribute ID is found.
// Returns a *NotFoundError when no entities are found.
func (iaq *InventoryAttributeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = iaq.Limit(2).IDs(setContextOp(ctx, iaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{inventoryattribute.Label}
	default:
		err = &NotSingularError{inventoryattribute.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iaq *InventoryAttributeQuery) OnlyIDX(ctx context.Context) string {
	id, err := iaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of InventoryAttributes.
func (iaq *InventoryAttributeQuery) All(ctx context.Context) ([]*InventoryAttribute, error) {
	ctx = setContextOp(ctx, iaq.ctx, "All")
	if err := iaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*InventoryAttribute, *InventoryAttributeQuery]()
	return withInterceptors[[]*InventoryAttribute](ctx, iaq, qr, iaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iaq *InventoryAttributeQuery) AllX(ctx context.Context) []*InventoryAttribute {
	nodes, err := iaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of InventoryAttribute IDs.
func (iaq *InventoryAttributeQuery) IDs(ctx context.Context) (ids []string, err error) {
	if iaq.ctx.Unique == nil && iaq.path != nil {
		iaq.Unique(true)
	}
	ctx = setContextOp(ctx, iaq.ctx, "IDs")
	if err = iaq.Select(inventoryattribute.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iaq *InventoryAttributeQuery) IDsX(ctx context.Context) []string {
	ids, err := iaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iaq *InventoryAttributeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iaq.ctx, "Count")
	if err := iaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iaq, querierCount[*InventoryAttributeQuery](), iaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iaq *InventoryAttributeQuery) CountX(ctx context.Context) int {
	count, err := iaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iaq *InventoryAttributeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iaq.ctx, "Exist")
	switch _, err := iaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("placio_api: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iaq *InventoryAttributeQuery) ExistX(ctx context.Context) bool {
	exist, err := iaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InventoryAttributeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iaq *InventoryAttributeQuery) Clone() *InventoryAttributeQuery {
	if iaq == nil {
		return nil
	}
	return &InventoryAttributeQuery{
		config:                       iaq.config,
		ctx:                          iaq.ctx.Clone(),
		order:                        append([]inventoryattribute.OrderOption{}, iaq.order...),
		inters:                       append([]Interceptor{}, iaq.inters...),
		predicates:                   append([]predicate.InventoryAttribute{}, iaq.predicates...),
		withInventoryType:            iaq.withInventoryType.Clone(),
		withPlaceInventoryAttributes: iaq.withPlaceInventoryAttributes.Clone(),
		// clone intermediate query.
		sql:  iaq.sql.Clone(),
		path: iaq.path,
	}
}

// WithInventoryType tells the query-builder to eager-load the nodes that are connected to
// the "inventory_type" edge. The optional arguments are used to configure the query builder of the edge.
func (iaq *InventoryAttributeQuery) WithInventoryType(opts ...func(*InventoryTypeQuery)) *InventoryAttributeQuery {
	query := (&InventoryTypeClient{config: iaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iaq.withInventoryType = query
	return iaq
}

// WithPlaceInventoryAttributes tells the query-builder to eager-load the nodes that are connected to
// the "place_inventory_attributes" edge. The optional arguments are used to configure the query builder of the edge.
func (iaq *InventoryAttributeQuery) WithPlaceInventoryAttributes(opts ...func(*PlaceInventoryAttributeQuery)) *InventoryAttributeQuery {
	query := (&PlaceInventoryAttributeClient{config: iaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iaq.withPlaceInventoryAttributes = query
	return iaq
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
//	client.InventoryAttribute.Query().
//		GroupBy(inventoryattribute.FieldName).
//		Aggregate(placio_api.Count()).
//		Scan(ctx, &v)
func (iaq *InventoryAttributeQuery) GroupBy(field string, fields ...string) *InventoryAttributeGroupBy {
	iaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InventoryAttributeGroupBy{build: iaq}
	grbuild.flds = &iaq.ctx.Fields
	grbuild.label = inventoryattribute.Label
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
//	client.InventoryAttribute.Query().
//		Select(inventoryattribute.FieldName).
//		Scan(ctx, &v)
func (iaq *InventoryAttributeQuery) Select(fields ...string) *InventoryAttributeSelect {
	iaq.ctx.Fields = append(iaq.ctx.Fields, fields...)
	sbuild := &InventoryAttributeSelect{InventoryAttributeQuery: iaq}
	sbuild.label = inventoryattribute.Label
	sbuild.flds, sbuild.scan = &iaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InventoryAttributeSelect configured with the given aggregations.
func (iaq *InventoryAttributeQuery) Aggregate(fns ...AggregateFunc) *InventoryAttributeSelect {
	return iaq.Select().Aggregate(fns...)
}

func (iaq *InventoryAttributeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iaq.inters {
		if inter == nil {
			return fmt.Errorf("placio_api: uninitialized interceptor (forgotten import placio_api/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iaq); err != nil {
				return err
			}
		}
	}
	for _, f := range iaq.ctx.Fields {
		if !inventoryattribute.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
		}
	}
	if iaq.path != nil {
		prev, err := iaq.path(ctx)
		if err != nil {
			return err
		}
		iaq.sql = prev
	}
	return nil
}

func (iaq *InventoryAttributeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*InventoryAttribute, error) {
	var (
		nodes       = []*InventoryAttribute{}
		withFKs     = iaq.withFKs
		_spec       = iaq.querySpec()
		loadedTypes = [2]bool{
			iaq.withInventoryType != nil,
			iaq.withPlaceInventoryAttributes != nil,
		}
	)
	if iaq.withInventoryType != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, inventoryattribute.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*InventoryAttribute).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &InventoryAttribute{config: iaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iaq.withInventoryType; query != nil {
		if err := iaq.loadInventoryType(ctx, query, nodes, nil,
			func(n *InventoryAttribute, e *InventoryType) { n.Edges.InventoryType = e }); err != nil {
			return nil, err
		}
	}
	if query := iaq.withPlaceInventoryAttributes; query != nil {
		if err := iaq.loadPlaceInventoryAttributes(ctx, query, nodes,
			func(n *InventoryAttribute) { n.Edges.PlaceInventoryAttributes = []*PlaceInventoryAttribute{} },
			func(n *InventoryAttribute, e *PlaceInventoryAttribute) {
				n.Edges.PlaceInventoryAttributes = append(n.Edges.PlaceInventoryAttributes, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iaq *InventoryAttributeQuery) loadInventoryType(ctx context.Context, query *InventoryTypeQuery, nodes []*InventoryAttribute, init func(*InventoryAttribute), assign func(*InventoryAttribute, *InventoryType)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*InventoryAttribute)
	for i := range nodes {
		if nodes[i].inventory_type_attributes == nil {
			continue
		}
		fk := *nodes[i].inventory_type_attributes
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(inventorytype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "inventory_type_attributes" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iaq *InventoryAttributeQuery) loadPlaceInventoryAttributes(ctx context.Context, query *PlaceInventoryAttributeQuery, nodes []*InventoryAttribute, init func(*InventoryAttribute), assign func(*InventoryAttribute, *PlaceInventoryAttribute)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*InventoryAttribute)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.PlaceInventoryAttribute(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(inventoryattribute.PlaceInventoryAttributesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.inventory_attribute_place_inventory_attributes
		if fk == nil {
			return fmt.Errorf(`foreign-key "inventory_attribute_place_inventory_attributes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "inventory_attribute_place_inventory_attributes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (iaq *InventoryAttributeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iaq.querySpec()
	_spec.Node.Columns = iaq.ctx.Fields
	if len(iaq.ctx.Fields) > 0 {
		_spec.Unique = iaq.ctx.Unique != nil && *iaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iaq.driver, _spec)
}

func (iaq *InventoryAttributeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(inventoryattribute.Table, inventoryattribute.Columns, sqlgraph.NewFieldSpec(inventoryattribute.FieldID, field.TypeString))
	_spec.From = iaq.sql
	if unique := iaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iaq.path != nil {
		_spec.Unique = true
	}
	if fields := iaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inventoryattribute.FieldID)
		for i := range fields {
			if fields[i] != inventoryattribute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iaq *InventoryAttributeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iaq.driver.Dialect())
	t1 := builder.Table(inventoryattribute.Table)
	columns := iaq.ctx.Fields
	if len(columns) == 0 {
		columns = inventoryattribute.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iaq.sql != nil {
		selector = iaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iaq.ctx.Unique != nil && *iaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iaq.predicates {
		p(selector)
	}
	for _, p := range iaq.order {
		p(selector)
	}
	if offset := iaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InventoryAttributeGroupBy is the group-by builder for InventoryAttribute entities.
type InventoryAttributeGroupBy struct {
	selector
	build *InventoryAttributeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iagb *InventoryAttributeGroupBy) Aggregate(fns ...AggregateFunc) *InventoryAttributeGroupBy {
	iagb.fns = append(iagb.fns, fns...)
	return iagb
}

// Scan applies the selector query and scans the result into the given value.
func (iagb *InventoryAttributeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iagb.build.ctx, "GroupBy")
	if err := iagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InventoryAttributeQuery, *InventoryAttributeGroupBy](ctx, iagb.build, iagb, iagb.build.inters, v)
}

func (iagb *InventoryAttributeGroupBy) sqlScan(ctx context.Context, root *InventoryAttributeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(iagb.fns))
	for _, fn := range iagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*iagb.flds)+len(iagb.fns))
		for _, f := range *iagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*iagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InventoryAttributeSelect is the builder for selecting fields of InventoryAttribute entities.
type InventoryAttributeSelect struct {
	*InventoryAttributeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ias *InventoryAttributeSelect) Aggregate(fns ...AggregateFunc) *InventoryAttributeSelect {
	ias.fns = append(ias.fns, fns...)
	return ias
}

// Scan applies the selector query and scans the result into the given value.
func (ias *InventoryAttributeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ias.ctx, "Select")
	if err := ias.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InventoryAttributeQuery, *InventoryAttributeSelect](ctx, ias.InventoryAttributeQuery, ias, ias.inters, v)
}

func (ias *InventoryAttributeSelect) sqlScan(ctx context.Context, root *InventoryAttributeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ias.fns))
	for _, fn := range ias.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ias.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ias.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
