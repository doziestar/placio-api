// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/comment"
	"placio-app/ent/notification"
	"placio-app/ent/post"
	"placio-app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetContent sets the "Content" field.
func (cc *CommentCreate) SetContent(s string) *CommentCreate {
	cc.mutation.SetContent(s)
	return cc
}

// SetCreatedAt sets the "CreatedAt" field.
func (cc *CommentCreate) SetCreatedAt(t time.Time) *CommentCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (cc *CommentCreate) SetNillableCreatedAt(t *time.Time) *CommentCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (cc *CommentCreate) SetUpdatedAt(t time.Time) *CommentCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetParentCommentID sets the "parentCommentID" field.
func (cc *CommentCreate) SetParentCommentID(s string) *CommentCreate {
	cc.mutation.SetParentCommentID(s)
	return cc
}

// SetNillableParentCommentID sets the "parentCommentID" field if the given value is not nil.
func (cc *CommentCreate) SetNillableParentCommentID(s *string) *CommentCreate {
	if s != nil {
		cc.SetParentCommentID(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CommentCreate) SetID(s string) *CommentCreate {
	cc.mutation.SetID(s)
	return cc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cc *CommentCreate) SetUserID(id string) *CommentCreate {
	cc.mutation.SetUserID(id)
	return cc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillableUserID(id *string) *CommentCreate {
	if id != nil {
		cc = cc.SetUserID(*id)
	}
	return cc
}

// SetUser sets the "user" edge to the User entity.
func (cc *CommentCreate) SetUser(u *User) *CommentCreate {
	return cc.SetUserID(u.ID)
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (cc *CommentCreate) SetPostID(id string) *CommentCreate {
	cc.mutation.SetPostID(id)
	return cc
}

// SetNillablePostID sets the "post" edge to the Post entity by ID if the given value is not nil.
func (cc *CommentCreate) SetNillablePostID(id *string) *CommentCreate {
	if id != nil {
		cc = cc.SetPostID(*id)
	}
	return cc
}

// SetPost sets the "post" edge to the Post entity.
func (cc *CommentCreate) SetPost(p *Post) *CommentCreate {
	return cc.SetPostID(p.ID)
}

// SetParentComment sets the "parentComment" edge to the Comment entity.
func (cc *CommentCreate) SetParentComment(c *Comment) *CommentCreate {
	return cc.SetParentCommentID(c.ID)
}

// AddReplyIDs adds the "replies" edge to the Comment entity by IDs.
func (cc *CommentCreate) AddReplyIDs(ids ...string) *CommentCreate {
	cc.mutation.AddReplyIDs(ids...)
	return cc
}

// AddReplies adds the "replies" edges to the Comment entity.
func (cc *CommentCreate) AddReplies(c ...*Comment) *CommentCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddReplyIDs(ids...)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (cc *CommentCreate) AddNotificationIDs(ids ...string) *CommentCreate {
	cc.mutation.AddNotificationIDs(ids...)
	return cc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (cc *CommentCreate) AddNotifications(n ...*Notification) *CommentCreate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cc.AddNotificationIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommentCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := comment.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.Content(); !ok {
		return &ValidationError{Name: "Content", err: errors.New(`ent: missing required field "Comment.Content"`)}
	}
	if v, ok := cc.mutation.Content(); ok {
		if err := comment.ContentValidator(v); err != nil {
			return &ValidationError{Name: "Content", err: fmt.Errorf(`ent: validator failed for field "Comment.Content": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "Comment.CreatedAt"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "Comment.UpdatedAt"`)}
	}
	if v, ok := cc.mutation.ID(); ok {
		if err := comment.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Comment.id": %w`, err)}
		}
	}
	return nil
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Comment.ID type: %T", _spec.ID.Value)
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(comment.Table, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := cc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_comments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.post_comments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ParentCommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ParentCommentTable,
			Columns: []string{comment.ParentCommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentCommentID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.RepliesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.RepliesTable,
			Columns: []string{comment.RepliesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   comment.NotificationsTable,
			Columns: comment.NotificationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notification.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	err      error
	builders []*CommentCreate
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
