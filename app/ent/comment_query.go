



// Code generated by ent, DO NOT EDIT.



package ent



	
import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
		"placio-app/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
			"database/sql/driver"
			"entgo.io/ent/dialect/sql"
			"entgo.io/ent/dialect/sql/sqlgraph"
			"entgo.io/ent/dialect/sql/sqljson"
			"entgo.io/ent/schema/field"
			 "placio-app/ent/comment"
			 "placio-app/ent/user"
			 "placio-app/ent/post"
			 "placio-app/ent/notification"

)






// CommentQuery is the builder for querying Comment entities.
type CommentQuery struct {
	config
	ctx			*QueryContext
	order		[]comment.OrderOption
	inters		[]Interceptor
	predicates 	[]predicate.Comment
		withUser *UserQuery
		withPost *PostQuery
		withParentComment *CommentQuery
		withReplies *CommentQuery
		withNotifications *NotificationQuery
		withFKs bool
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommentQuery builder.
func (cq *CommentQuery) Where(ps ...predicate.Comment) *CommentQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CommentQuery) Limit(limit int) *CommentQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CommentQuery) Offset(offset int) *CommentQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CommentQuery) Unique(unique bool) *CommentQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CommentQuery) Order(o ...comment.OrderOption) *CommentQuery {
	cq.order = append(cq.order, o...)
	return cq
}



	
	// QueryUser chains the current query on the "user" edge.
	func (cq *CommentQuery) QueryUser() *UserQuery {
		query := (&UserClient{config: cq.config}).Query()
		query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
			if err := cq.prepareQuery(ctx); err != nil {
				return nil, err
			}  
	selector := cq.sqlQuery(ctx)
	if err := selector.Err(); err != nil {
		return nil, err
	}
	step := sqlgraph.NewStep(
		sqlgraph.From(comment.Table, comment.FieldID, selector),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, comment.UserTable,comment.UserColumn),
	)
	fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
return fromU, nil
		}
		return query
	}

	
	// QueryPost chains the current query on the "post" edge.
	func (cq *CommentQuery) QueryPost() *PostQuery {
		query := (&PostClient{config: cq.config}).Query()
		query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
			if err := cq.prepareQuery(ctx); err != nil {
				return nil, err
			}  
	selector := cq.sqlQuery(ctx)
	if err := selector.Err(); err != nil {
		return nil, err
	}
	step := sqlgraph.NewStep(
		sqlgraph.From(comment.Table, comment.FieldID, selector),
		sqlgraph.To(post.Table, post.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, comment.PostTable,comment.PostColumn),
	)
	fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
return fromU, nil
		}
		return query
	}

	
	// QueryParentComment chains the current query on the "parentComment" edge.
	func (cq *CommentQuery) QueryParentComment() *CommentQuery {
		query := (&CommentClient{config: cq.config}).Query()
		query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
			if err := cq.prepareQuery(ctx); err != nil {
				return nil, err
			}  
	selector := cq.sqlQuery(ctx)
	if err := selector.Err(); err != nil {
		return nil, err
	}
	step := sqlgraph.NewStep(
		sqlgraph.From(comment.Table, comment.FieldID, selector),
		sqlgraph.To(comment.Table, comment.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, comment.ParentCommentTable,comment.ParentCommentColumn),
	)
	fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
return fromU, nil
		}
		return query
	}

	
	// QueryReplies chains the current query on the "replies" edge.
	func (cq *CommentQuery) QueryReplies() *CommentQuery {
		query := (&CommentClient{config: cq.config}).Query()
		query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
			if err := cq.prepareQuery(ctx); err != nil {
				return nil, err
			}  
	selector := cq.sqlQuery(ctx)
	if err := selector.Err(); err != nil {
		return nil, err
	}
	step := sqlgraph.NewStep(
		sqlgraph.From(comment.Table, comment.FieldID, selector),
		sqlgraph.To(comment.Table, comment.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, comment.RepliesTable,comment.RepliesColumn),
	)
	fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
return fromU, nil
		}
		return query
	}

	
	// QueryNotifications chains the current query on the "notifications" edge.
	func (cq *CommentQuery) QueryNotifications() *NotificationQuery {
		query := (&NotificationClient{config: cq.config}).Query()
		query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
			if err := cq.prepareQuery(ctx); err != nil {
				return nil, err
			}  
	selector := cq.sqlQuery(ctx)
	if err := selector.Err(); err != nil {
		return nil, err
	}
	step := sqlgraph.NewStep(
		sqlgraph.From(comment.Table, comment.FieldID, selector),
		sqlgraph.To(notification.Table, notification.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, comment.NotificationsTable,comment.NotificationsPrimaryKey...),
	)
	fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
return fromU, nil
		}
		return query
	}


// First returns the first Comment entity from the query. 
// Returns a *NotFoundError when no Comment was found.
func (cq *CommentQuery) First(ctx context.Context) (*Comment, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ comment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CommentQuery) FirstX(ctx context.Context) *Comment {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}


	// FirstID returns the first Comment ID from the query.
	// Returns a *NotFoundError when no Comment ID was found.
	func (cq *CommentQuery) FirstID(ctx context.Context) (id string, err error) {
		var ids []string
		if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
			return
		}
		if len(ids) == 0 {
			err = &NotFoundError{ comment.Label}
			return
		}
		return ids[0], nil
	}

	// FirstIDX is like FirstID, but panics if an error occurs.
	func (cq *CommentQuery) FirstIDX(ctx context.Context) string {
		id, err := cq.FirstID(ctx)
		if err != nil && !IsNotFound(err) {
			panic(err)
		}
		return id
	}


// Only returns a single Comment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Comment entity is found.
// Returns a *NotFoundError when no Comment entities are found.
func (cq *CommentQuery) Only(ctx context.Context) (*Comment, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ comment.Label}
	default:
		return nil, &NotSingularError{ comment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CommentQuery) OnlyX(ctx context.Context) *Comment {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}


	// OnlyID is like Only, but returns the only Comment ID in the query.
	// Returns a *NotSingularError when more than one Comment ID is found.
	// Returns a *NotFoundError when no entities are found.
	func (cq *CommentQuery) OnlyID(ctx context.Context) (id string, err error) {
		var ids []string
		if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
			return
		}
		switch len(ids) {
		case 1:
			id = ids[0]
		case 0:
			err = &NotFoundError{ comment.Label}
		default:
			err = &NotSingularError{ comment.Label}
		}
		return
	}

	// OnlyIDX is like OnlyID, but panics if an error occurs.
	func (cq *CommentQuery) OnlyIDX(ctx context.Context) string {
		id, err := cq.OnlyID(ctx)
		if err != nil {
			panic(err)
		}
		return id
	}


// All executes the query and returns a list of Comments.
func (cq *CommentQuery) All(ctx context.Context) ([]*Comment, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Comment, *CommentQuery]()
	return withInterceptors[[]*Comment](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CommentQuery) AllX(ctx context.Context) []*Comment {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}


	// IDs executes the query and returns a list of Comment IDs.
	func (cq *CommentQuery) IDs(ctx context.Context) (ids []string, err error) {
		if cq.ctx.Unique == nil && cq.path != nil {
			cq.Unique(true)
		}
		ctx = setContextOp(ctx, cq.ctx, "IDs")
		if err = cq.Select(comment.FieldID).Scan(ctx, &ids); err != nil {
			return nil, err
		}
		return ids, nil
	}

	// IDsX is like IDs, but panics if an error occurs.
	func (cq *CommentQuery) IDsX(ctx context.Context) []string {
		ids, err := cq.IDs(ctx)
		if err != nil {
			panic(err)
		}
		return ids
	}


// Count returns the count of the given query.
func (cq *CommentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CommentQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CommentQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CommentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx);{
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CommentQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CommentQuery) Clone() *CommentQuery {
	if cq == nil {
		return nil
	}
	return &CommentQuery{
		config: 	cq.config,
		ctx: 		cq.ctx.Clone(),
		order: 		append([]comment.OrderOption{}, cq.order...),
		inters: 	append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Comment{}, cq.predicates...),
			withUser: cq.withUser.Clone(),
			withPost: cq.withPost.Clone(),
			withParentComment: cq.withParentComment.Clone(),
			withReplies: cq.withReplies.Clone(),
			withNotifications: cq.withNotifications.Clone(),
		// clone intermediate query.
		sql: cq.sql.Clone(),
		path: cq.path,
	}
}
	
	
	// WithUser tells the query-builder to eager-load the nodes that are connected to
	// the "user" edge. The optional arguments are used to configure the query builder of the edge.
	func (cq *CommentQuery) WithUser(opts ...func(*UserQuery)) *CommentQuery {
		query := (&UserClient{config: cq.config}).Query()
		for _, opt := range opts {
			opt(query)
		}
		cq.withUser = query
		return cq
	}
	
	
	// WithPost tells the query-builder to eager-load the nodes that are connected to
	// the "post" edge. The optional arguments are used to configure the query builder of the edge.
	func (cq *CommentQuery) WithPost(opts ...func(*PostQuery)) *CommentQuery {
		query := (&PostClient{config: cq.config}).Query()
		for _, opt := range opts {
			opt(query)
		}
		cq.withPost = query
		return cq
	}
	
	
	// WithParentComment tells the query-builder to eager-load the nodes that are connected to
	// the "parentComment" edge. The optional arguments are used to configure the query builder of the edge.
	func (cq *CommentQuery) WithParentComment(opts ...func(*CommentQuery)) *CommentQuery {
		query := (&CommentClient{config: cq.config}).Query()
		for _, opt := range opts {
			opt(query)
		}
		cq.withParentComment = query
		return cq
	}
	
	
	// WithReplies tells the query-builder to eager-load the nodes that are connected to
	// the "replies" edge. The optional arguments are used to configure the query builder of the edge.
	func (cq *CommentQuery) WithReplies(opts ...func(*CommentQuery)) *CommentQuery {
		query := (&CommentClient{config: cq.config}).Query()
		for _, opt := range opts {
			opt(query)
		}
		cq.withReplies = query
		return cq
	}
	
	
	// WithNotifications tells the query-builder to eager-load the nodes that are connected to
	// the "notifications" edge. The optional arguments are used to configure the query builder of the edge.
	func (cq *CommentQuery) WithNotifications(opts ...func(*NotificationQuery)) *CommentQuery {
		query := (&NotificationClient{config: cq.config}).Query()
		for _, opt := range opts {
			opt(query)
		}
		cq.withNotifications = query
		return cq
	}



// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Content string `json:"Content,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Comment.Query().
//		GroupBy(comment.FieldContent).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *CommentQuery) GroupBy(field string, fields ...string) *CommentGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommentGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = comment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}



// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Content string `json:"Content,omitempty"`
//	}
//
//	client.Comment.Query().
//		Select(comment.FieldContent).
//		Scan(ctx, &v)
//
func (cq *CommentQuery) Select(fields ...string) *CommentSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CommentSelect{ CommentQuery: cq }
	sbuild.label = comment.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommentSelect configured with the given aggregations.
func (cq *CommentQuery) Aggregate(fns ...AggregateFunc) *CommentSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CommentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !comment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}


	
	




func (cq *CommentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Comment, error) {
	var (
		nodes = []*Comment{}
			withFKs = cq.withFKs
		_spec = cq.querySpec()
			loadedTypes = [5]bool{
					cq.withUser != nil,
					cq.withPost != nil,
					cq.withParentComment != nil,
					cq.withReplies != nil,
					cq.withNotifications != nil,
			}
	)
			if cq.withUser != nil || cq.withPost != nil {
				withFKs = true
			}
		if withFKs {
			_spec.Node.Columns = append(_spec.Node.Columns, comment.ForeignKeys...)
		}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Comment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Comment{config: cq.config}
		nodes = append(nodes, node)
			node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
		if query := cq.withUser; query != nil {
			if err := cq.loadUser(ctx, query, nodes, nil,
				func(n *Comment, e *User){ n.Edges.User = e }); err != nil {
				return nil, err
			}
		}
		if query := cq.withPost; query != nil {
			if err := cq.loadPost(ctx, query, nodes, nil,
				func(n *Comment, e *Post){ n.Edges.Post = e }); err != nil {
				return nil, err
			}
		}
		if query := cq.withParentComment; query != nil {
			if err := cq.loadParentComment(ctx, query, nodes, nil,
				func(n *Comment, e *Comment){ n.Edges.ParentComment = e }); err != nil {
				return nil, err
			}
		}
		if query := cq.withReplies; query != nil {
			if err := cq.loadReplies(ctx, query, nodes, 
				func(n *Comment){ n.Edges.Replies = []*Comment{} },
				func(n *Comment, e *Comment){ n.Edges.Replies = append(n.Edges.Replies, e) }); err != nil {
				return nil, err
			}
		}
		if query := cq.withNotifications; query != nil {
			if err := cq.loadNotifications(ctx, query, nodes, 
				func(n *Comment){ n.Edges.Notifications = []*Notification{} },
				func(n *Comment, e *Notification){ n.Edges.Notifications = append(n.Edges.Notifications, e) }); err != nil {
				return nil, err
			}
		}
	return nodes, nil
}


	func (cq *CommentQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Comment, init func(*Comment), assign func(*Comment, *User)) error {
			ids := make([]string, 0, len(nodes))
			nodeids := make(map[string][]*Comment)
			for i := range nodes {
					if nodes[i].user_comments == nil {
						continue
					}
				fk := *nodes[i].user_comments
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
					return fmt.Errorf(`unexpected foreign-key "user_comments" returned %v`, n.ID)
				}
				for i := range nodes {
					assign(nodes[i], n)
				}
			}
		return nil
	}
	func (cq *CommentQuery) loadPost(ctx context.Context, query *PostQuery, nodes []*Comment, init func(*Comment), assign func(*Comment, *Post)) error {
			ids := make([]string, 0, len(nodes))
			nodeids := make(map[string][]*Comment)
			for i := range nodes {
					if nodes[i].post_comments == nil {
						continue
					}
				fk := *nodes[i].post_comments
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
					return fmt.Errorf(`unexpected foreign-key "post_comments" returned %v`, n.ID)
				}
				for i := range nodes {
					assign(nodes[i], n)
				}
			}
		return nil
	}
	func (cq *CommentQuery) loadParentComment(ctx context.Context, query *CommentQuery, nodes []*Comment, init func(*Comment), assign func(*Comment, *Comment)) error {
			ids := make([]string, 0, len(nodes))
			nodeids := make(map[string][]*Comment)
			for i := range nodes {
					if nodes[i].ParentCommentID == nil {
						continue
					}
				fk := *nodes[i].ParentCommentID
				if _, ok := nodeids[fk]; !ok {
					ids = append(ids, fk)
				}
				nodeids[fk] = append(nodeids[fk], nodes[i])
			}
			if len(ids) == 0 {
				return nil
			}
			query.Where(comment.IDIn(ids...))
			neighbors, err := query.All(ctx)
			if err != nil {
				return err
			}
			for _, n := range neighbors {
				nodes, ok := nodeids[n.ID]
				if !ok {
					return fmt.Errorf(`unexpected foreign-key "parentCommentID" returned %v`, n.ID)
				}
				for i := range nodes {
					assign(nodes[i], n)
				}
			}
		return nil
	}
	func (cq *CommentQuery) loadReplies(ctx context.Context, query *CommentQuery, nodes []*Comment, init func(*Comment), assign func(*Comment, *Comment)) error {
			fks := make([]driver.Value, 0, len(nodes))
			nodeids := make(map[string]*Comment)
			for i := range nodes {
				fks = append(fks, nodes[i].ID)
				nodeids[nodes[i].ID] = nodes[i]
					if init != nil {
						init(nodes[i])
					}
			}
				query.withFKs = true
				if len(query.ctx.Fields) > 0 {
					query.ctx.AppendFieldOnce(comment.FieldParentCommentID)
				}
			query.Where(predicate.Comment(func(s *sql.Selector) {
				s.Where(sql.InValues(s.C(comment.RepliesColumn), fks...))
			}))
			neighbors, err := query.All(ctx)
			if err != nil {
				return err
			}
			for _, n := range neighbors {
				fk := n.ParentCommentID
					if fk == nil {
						return fmt.Errorf(`foreign-key "parentCommentID" is nil for node %v`, n.ID)
					}
				node, ok := nodeids[*fk]
				if !ok {
					return fmt.Errorf(`unexpected referenced foreign-key "parentCommentID" returned %v for node %v`, *fk, n.ID)
				}
				assign(node, n)
			}
		return nil
	}
	func (cq *CommentQuery) loadNotifications(ctx context.Context, query *NotificationQuery, nodes []*Comment, init func(*Comment), assign func(*Comment, *Notification)) error {
			edgeIDs := make([]driver.Value, len(nodes))
			byID := make(map[string]*Comment)
			nids := make(map[string]map[*Comment]struct{})
			for i, node := range nodes {
				edgeIDs[i] = node.ID
				byID[node.ID] = node
				if init != nil {
					init(node)
				}
			}
			query.Where(func(s *sql.Selector) {
				joinT := sql.Table(comment.NotificationsTable)
				s.Join(joinT).On(s.C(notification.FieldID), joinT.C(comment.NotificationsPrimaryKey[1]))
				s.Where(sql.InValues(joinT.C(comment.NotificationsPrimaryKey[0]), edgeIDs...))
				columns := s.SelectedColumns()
				s.Select(joinT.C(comment.NotificationsPrimaryKey[0]))
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
							nids[inValue] = map[*Comment]struct{}{byID[outValue]: {}}
							return assign(columns[1:], values[1:])
						}
						nids[inValue][byID[outValue]] = struct{}{}
						return nil
					}
				})
			})
			neighbors, err :=  withInterceptors[[]*Notification](ctx, query, qr, query.inters)
			if err != nil {
				return err
			}
			for _, n := range neighbors {
				nodes, ok := nids[n.ID]
				if !ok {
					return fmt.Errorf(`unexpected "notifications" node returned %v`, n.ID)
				}
				for kn := range nodes {
					assign(kn, n)
				}
			}
		return nil
	}

func (cq *CommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
		_spec.Node.Columns = cq.ctx.Fields
		if len(cq.ctx.Fields) > 0 {
			_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
		}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
			_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
			for i := range fields {
				if fields[i] != comment.FieldID {
					_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
				}
			}
			if cq.withParentComment != nil {
				_spec.Node.AddColumnOnce(comment.FieldParentCommentID)
			}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}





func (cq *CommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(comment.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = comment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

    

    











// CommentGroupBy is the group-by builder for Comment entities.
type CommentGroupBy struct {
	selector
	build *CommentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CommentGroupBy) Aggregate(fns ...AggregateFunc) *CommentGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CommentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentQuery, *CommentGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}


	
	



func (cgb *CommentGroupBy) sqlScan(ctx context.Context, root *CommentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds) + len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}







// CommentSelect is the builder for selecting fields of Comment entities.
type CommentSelect struct {
	*CommentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CommentSelect) Aggregate(fns ...AggregateFunc) *CommentSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CommentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentQuery, *CommentSelect](ctx, cs.CommentQuery, cs, cs.inters, v)
}


	
	



func (cs *CommentSelect) sqlScan(ctx context.Context, root *CommentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}



    






