// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"placio_api/category"
	"placio_api/media"
	"placio_api/place"
	"placio_api/placeinventory"
	"placio_api/post"
	"placio_api/predicate"
	"placio_api/review"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MediaQuery is the builder for querying Media entities.
type MediaQuery struct {
	config
	ctx                *QueryContext
	order              []media.OrderOption
	inters             []Interceptor
	predicates         []predicate.Media
	withPost           *PostQuery
	withReview         *ReviewQuery
	withCategories     *CategoryQuery
	withPlace          *PlaceQuery
	withPlaceInventory *PlaceInventoryQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MediaQuery builder.
func (mq *MediaQuery) Where(ps ...predicate.Media) *MediaQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MediaQuery) Limit(limit int) *MediaQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MediaQuery) Offset(offset int) *MediaQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MediaQuery) Unique(unique bool) *MediaQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MediaQuery) Order(o ...media.OrderOption) *MediaQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryPost chains the current query on the "post" edge.
func (mq *MediaQuery) QueryPost() *PostQuery {
	query := (&PostClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(media.Table, media.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, media.PostTable, media.PostColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReview chains the current query on the "review" edge.
func (mq *MediaQuery) QueryReview() *ReviewQuery {
	query := (&ReviewClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(media.Table, media.FieldID, selector),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, media.ReviewTable, media.ReviewColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCategories chains the current query on the "categories" edge.
func (mq *MediaQuery) QueryCategories() *CategoryQuery {
	query := (&CategoryClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(media.Table, media.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, media.CategoriesTable, media.CategoriesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlace chains the current query on the "place" edge.
func (mq *MediaQuery) QueryPlace() *PlaceQuery {
	query := (&PlaceClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(media.Table, media.FieldID, selector),
			sqlgraph.To(place.Table, place.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, media.PlaceTable, media.PlacePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlaceInventory chains the current query on the "place_inventory" edge.
func (mq *MediaQuery) QueryPlaceInventory() *PlaceInventoryQuery {
	query := (&PlaceInventoryClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(media.Table, media.FieldID, selector),
			sqlgraph.To(placeinventory.Table, placeinventory.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, media.PlaceInventoryTable, media.PlaceInventoryPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Media entity from the query.
// Returns a *NotFoundError when no Media was found.
func (mq *MediaQuery) First(ctx context.Context) (*Media, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{media.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MediaQuery) FirstX(ctx context.Context) *Media {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Media ID from the query.
// Returns a *NotFoundError when no Media ID was found.
func (mq *MediaQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{media.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MediaQuery) FirstIDX(ctx context.Context) string {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Media entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Media entity is found.
// Returns a *NotFoundError when no Media entities are found.
func (mq *MediaQuery) Only(ctx context.Context) (*Media, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{media.Label}
	default:
		return nil, &NotSingularError{media.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MediaQuery) OnlyX(ctx context.Context) *Media {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Media ID in the query.
// Returns a *NotSingularError when more than one Media ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MediaQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{media.Label}
	default:
		err = &NotSingularError{media.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MediaQuery) OnlyIDX(ctx context.Context) string {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MediaSlice.
func (mq *MediaQuery) All(ctx context.Context) ([]*Media, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Media, *MediaQuery]()
	return withInterceptors[[]*Media](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MediaQuery) AllX(ctx context.Context) []*Media {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Media IDs.
func (mq *MediaQuery) IDs(ctx context.Context) (ids []string, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(media.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MediaQuery) IDsX(ctx context.Context) []string {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MediaQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MediaQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MediaQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MediaQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("placio_api: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MediaQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MediaQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MediaQuery) Clone() *MediaQuery {
	if mq == nil {
		return nil
	}
	return &MediaQuery{
		config:             mq.config,
		ctx:                mq.ctx.Clone(),
		order:              append([]media.OrderOption{}, mq.order...),
		inters:             append([]Interceptor{}, mq.inters...),
		predicates:         append([]predicate.Media{}, mq.predicates...),
		withPost:           mq.withPost.Clone(),
		withReview:         mq.withReview.Clone(),
		withCategories:     mq.withCategories.Clone(),
		withPlace:          mq.withPlace.Clone(),
		withPlaceInventory: mq.withPlaceInventory.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithPost tells the query-builder to eager-load the nodes that are connected to
// the "post" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MediaQuery) WithPost(opts ...func(*PostQuery)) *MediaQuery {
	query := (&PostClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withPost = query
	return mq
}

// WithReview tells the query-builder to eager-load the nodes that are connected to
// the "review" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MediaQuery) WithReview(opts ...func(*ReviewQuery)) *MediaQuery {
	query := (&ReviewClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withReview = query
	return mq
}

// WithCategories tells the query-builder to eager-load the nodes that are connected to
// the "categories" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MediaQuery) WithCategories(opts ...func(*CategoryQuery)) *MediaQuery {
	query := (&CategoryClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withCategories = query
	return mq
}

// WithPlace tells the query-builder to eager-load the nodes that are connected to
// the "place" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MediaQuery) WithPlace(opts ...func(*PlaceQuery)) *MediaQuery {
	query := (&PlaceClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withPlace = query
	return mq
}

// WithPlaceInventory tells the query-builder to eager-load the nodes that are connected to
// the "place_inventory" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MediaQuery) WithPlaceInventory(opts ...func(*PlaceInventoryQuery)) *MediaQuery {
	query := (&PlaceInventoryClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withPlaceInventory = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		URL string `json:"URL,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Media.Query().
//		GroupBy(media.FieldURL).
//		Aggregate(placio_api.Count()).
//		Scan(ctx, &v)
func (mq *MediaQuery) GroupBy(field string, fields ...string) *MediaGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MediaGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = media.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		URL string `json:"URL,omitempty"`
//	}
//
//	client.Media.Query().
//		Select(media.FieldURL).
//		Scan(ctx, &v)
func (mq *MediaQuery) Select(fields ...string) *MediaSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MediaSelect{MediaQuery: mq}
	sbuild.label = media.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MediaSelect configured with the given aggregations.
func (mq *MediaQuery) Aggregate(fns ...AggregateFunc) *MediaSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MediaQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("placio_api: uninitialized interceptor (forgotten import placio_api/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !media.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MediaQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Media, error) {
	var (
		nodes       = []*Media{}
		withFKs     = mq.withFKs
		_spec       = mq.querySpec()
		loadedTypes = [5]bool{
			mq.withPost != nil,
			mq.withReview != nil,
			mq.withCategories != nil,
			mq.withPlace != nil,
			mq.withPlaceInventory != nil,
		}
	)
	if mq.withPost != nil || mq.withReview != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, media.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Media).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Media{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withPost; query != nil {
		if err := mq.loadPost(ctx, query, nodes, nil,
			func(n *Media, e *Post) { n.Edges.Post = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withReview; query != nil {
		if err := mq.loadReview(ctx, query, nodes, nil,
			func(n *Media, e *Review) { n.Edges.Review = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withCategories; query != nil {
		if err := mq.loadCategories(ctx, query, nodes,
			func(n *Media) { n.Edges.Categories = []*Category{} },
			func(n *Media, e *Category) { n.Edges.Categories = append(n.Edges.Categories, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withPlace; query != nil {
		if err := mq.loadPlace(ctx, query, nodes,
			func(n *Media) { n.Edges.Place = []*Place{} },
			func(n *Media, e *Place) { n.Edges.Place = append(n.Edges.Place, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withPlaceInventory; query != nil {
		if err := mq.loadPlaceInventory(ctx, query, nodes,
			func(n *Media) { n.Edges.PlaceInventory = []*PlaceInventory{} },
			func(n *Media, e *PlaceInventory) { n.Edges.PlaceInventory = append(n.Edges.PlaceInventory, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MediaQuery) loadPost(ctx context.Context, query *PostQuery, nodes []*Media, init func(*Media), assign func(*Media, *Post)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Media)
	for i := range nodes {
		if nodes[i].post_medias == nil {
			continue
		}
		fk := *nodes[i].post_medias
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(post.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "post_medias" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MediaQuery) loadReview(ctx context.Context, query *ReviewQuery, nodes []*Media, init func(*Media), assign func(*Media, *Review)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Media)
	for i := range nodes {
		if nodes[i].review_medias == nil {
			continue
		}
		fk := *nodes[i].review_medias
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(review.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "review_medias" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MediaQuery) loadCategories(ctx context.Context, query *CategoryQuery, nodes []*Media, init func(*Media), assign func(*Media, *Category)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Media)
	nids := make(map[string]map[*Media]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(media.CategoriesTable)
		s.Join(joinT).On(s.C(category.FieldID), joinT.C(media.CategoriesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(media.CategoriesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(media.CategoriesPrimaryKey[1]))
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
					nids[inValue] = map[*Media]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Category](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "categories" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (mq *MediaQuery) loadPlace(ctx context.Context, query *PlaceQuery, nodes []*Media, init func(*Media), assign func(*Media, *Place)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Media)
	nids := make(map[string]map[*Media]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(media.PlaceTable)
		s.Join(joinT).On(s.C(place.FieldID), joinT.C(media.PlacePrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(media.PlacePrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(media.PlacePrimaryKey[1]))
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
					nids[inValue] = map[*Media]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Place](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "place" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (mq *MediaQuery) loadPlaceInventory(ctx context.Context, query *PlaceInventoryQuery, nodes []*Media, init func(*Media), assign func(*Media, *PlaceInventory)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Media)
	nids := make(map[string]map[*Media]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(media.PlaceInventoryTable)
		s.Join(joinT).On(s.C(placeinventory.FieldID), joinT.C(media.PlaceInventoryPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(media.PlaceInventoryPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(media.PlaceInventoryPrimaryKey[1]))
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
					nids[inValue] = map[*Media]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*PlaceInventory](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "place_inventory" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (mq *MediaQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MediaQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(media.Table, media.Columns, sqlgraph.NewFieldSpec(media.FieldID, field.TypeString))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, media.FieldID)
		for i := range fields {
			if fields[i] != media.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MediaQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(media.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = media.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MediaGroupBy is the group-by builder for Media entities.
type MediaGroupBy struct {
	selector
	build *MediaQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MediaGroupBy) Aggregate(fns ...AggregateFunc) *MediaGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MediaGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MediaQuery, *MediaGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MediaGroupBy) sqlScan(ctx context.Context, root *MediaQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MediaSelect is the builder for selecting fields of Media entities.
type MediaSelect struct {
	*MediaQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MediaSelect) Aggregate(fns ...AggregateFunc) *MediaSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MediaSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MediaQuery, *MediaSelect](ctx, ms.MediaQuery, ms, ms.inters, v)
}

func (ms *MediaSelect) sqlScan(ctx context.Context, root *MediaQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
