// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/comment"
	"placio-app/ent/event"
	"placio-app/ent/like"
	"placio-app/ent/media"
	"placio-app/ent/place"
	"placio-app/ent/review"
	"placio-app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReviewCreate is the builder for creating a Review entity.
type ReviewCreate struct {
	config
	mutation *ReviewMutation
	hooks    []Hook
}

// SetScore sets the "score" field.
func (rc *ReviewCreate) SetScore(f float64) *ReviewCreate {
	rc.mutation.SetScore(f)
	return rc
}

// SetContent sets the "content" field.
func (rc *ReviewCreate) SetContent(s string) *ReviewCreate {
	rc.mutation.SetContent(s)
	return rc
}

// SetCreatedAt sets the "createdAt" field.
func (rc *ReviewCreate) SetCreatedAt(t time.Time) *ReviewCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableCreatedAt(t *time.Time) *ReviewCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetLikeCount sets the "likeCount" field.
func (rc *ReviewCreate) SetLikeCount(i int) *ReviewCreate {
	rc.mutation.SetLikeCount(i)
	return rc
}

// SetNillableLikeCount sets the "likeCount" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableLikeCount(i *int) *ReviewCreate {
	if i != nil {
		rc.SetLikeCount(*i)
	}
	return rc
}

// SetDislikeCount sets the "dislikeCount" field.
func (rc *ReviewCreate) SetDislikeCount(i int) *ReviewCreate {
	rc.mutation.SetDislikeCount(i)
	return rc
}

// SetNillableDislikeCount sets the "dislikeCount" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableDislikeCount(i *int) *ReviewCreate {
	if i != nil {
		rc.SetDislikeCount(*i)
	}
	return rc
}

// SetFlag sets the "flag" field.
func (rc *ReviewCreate) SetFlag(s string) *ReviewCreate {
	rc.mutation.SetFlag(s)
	return rc
}

// SetNillableFlag sets the "flag" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableFlag(s *string) *ReviewCreate {
	if s != nil {
		rc.SetFlag(*s)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *ReviewCreate) SetID(s string) *ReviewCreate {
	rc.mutation.SetID(s)
	return rc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rc *ReviewCreate) SetUserID(id string) *ReviewCreate {
	rc.mutation.SetUserID(id)
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *ReviewCreate) SetUser(u *User) *ReviewCreate {
	return rc.SetUserID(u.ID)
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (rc *ReviewCreate) SetBusinessID(id string) *ReviewCreate {
	rc.mutation.SetBusinessID(id)
	return rc
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (rc *ReviewCreate) SetNillableBusinessID(id *string) *ReviewCreate {
	if id != nil {
		rc = rc.SetBusinessID(*id)
	}
	return rc
}

// SetBusiness sets the "business" edge to the Business entity.
func (rc *ReviewCreate) SetBusiness(b *Business) *ReviewCreate {
	return rc.SetBusinessID(b.ID)
}

// SetPlaceID sets the "place" edge to the Place entity by ID.
func (rc *ReviewCreate) SetPlaceID(id string) *ReviewCreate {
	rc.mutation.SetPlaceID(id)
	return rc
}

// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
func (rc *ReviewCreate) SetNillablePlaceID(id *string) *ReviewCreate {
	if id != nil {
		rc = rc.SetPlaceID(*id)
	}
	return rc
}

// SetPlace sets the "place" edge to the Place entity.
func (rc *ReviewCreate) SetPlace(p *Place) *ReviewCreate {
	return rc.SetPlaceID(p.ID)
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (rc *ReviewCreate) SetEventID(id string) *ReviewCreate {
	rc.mutation.SetEventID(id)
	return rc
}

// SetNillableEventID sets the "event" edge to the Event entity by ID if the given value is not nil.
func (rc *ReviewCreate) SetNillableEventID(id *string) *ReviewCreate {
	if id != nil {
		rc = rc.SetEventID(*id)
	}
	return rc
}

// SetEvent sets the "event" edge to the Event entity.
func (rc *ReviewCreate) SetEvent(e *Event) *ReviewCreate {
	return rc.SetEventID(e.ID)
}

// AddMediaIDs adds the "medias" edge to the Media entity by IDs.
func (rc *ReviewCreate) AddMediaIDs(ids ...string) *ReviewCreate {
	rc.mutation.AddMediaIDs(ids...)
	return rc
}

// AddMedias adds the "medias" edges to the Media entity.
func (rc *ReviewCreate) AddMedias(m ...*Media) *ReviewCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return rc.AddMediaIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (rc *ReviewCreate) AddCommentIDs(ids ...string) *ReviewCreate {
	rc.mutation.AddCommentIDs(ids...)
	return rc
}

// AddComments adds the "comments" edges to the Comment entity.
func (rc *ReviewCreate) AddComments(c ...*Comment) *ReviewCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rc.AddCommentIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (rc *ReviewCreate) AddLikeIDs(ids ...string) *ReviewCreate {
	rc.mutation.AddLikeIDs(ids...)
	return rc
}

// AddLikes adds the "likes" edges to the Like entity.
func (rc *ReviewCreate) AddLikes(l ...*Like) *ReviewCreate {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return rc.AddLikeIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (rc *ReviewCreate) Mutation() *ReviewMutation {
	return rc.mutation
}

// Save creates the Review in the database.
func (rc *ReviewCreate) Save(ctx context.Context) (*Review, error) {
	if err := rc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReviewCreate) SaveX(ctx context.Context) *Review {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReviewCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReviewCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReviewCreate) defaults() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		if review.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized review.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := review.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.LikeCount(); !ok {
		v := review.DefaultLikeCount
		rc.mutation.SetLikeCount(v)
	}
	if _, ok := rc.mutation.DislikeCount(); !ok {
		v := review.DefaultDislikeCount
		rc.mutation.SetDislikeCount(v)
	}
	if _, ok := rc.mutation.Flag(); !ok {
		v := review.DefaultFlag
		rc.mutation.SetFlag(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReviewCreate) check() error {
	if _, ok := rc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "Review.score"`)}
	}
	if v, ok := rc.mutation.Score(); ok {
		if err := review.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Review.score": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Review.content"`)}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Review.createdAt"`)}
	}
	if _, ok := rc.mutation.LikeCount(); !ok {
		return &ValidationError{Name: "likeCount", err: errors.New(`ent: missing required field "Review.likeCount"`)}
	}
	if _, ok := rc.mutation.DislikeCount(); !ok {
		return &ValidationError{Name: "dislikeCount", err: errors.New(`ent: missing required field "Review.dislikeCount"`)}
	}
	if _, ok := rc.mutation.Flag(); !ok {
		return &ValidationError{Name: "flag", err: errors.New(`ent: missing required field "Review.flag"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Review.user"`)}
	}
	return nil
}

func (rc *ReviewCreate) sqlSave(ctx context.Context) (*Review, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Review.ID type: %T", _spec.ID.Value)
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReviewCreate) createSpec() (*Review, *sqlgraph.CreateSpec) {
	var (
		_node = &Review{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(review.Table, sqlgraph.NewFieldSpec(review.FieldID, field.TypeString))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Score(); ok {
		_spec.SetField(review.FieldScore, field.TypeFloat64, value)
		_node.Score = value
	}
	if value, ok := rc.mutation.Content(); ok {
		_spec.SetField(review.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(review.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.LikeCount(); ok {
		_spec.SetField(review.FieldLikeCount, field.TypeInt, value)
		_node.LikeCount = value
	}
	if value, ok := rc.mutation.DislikeCount(); ok {
		_spec.SetField(review.FieldDislikeCount, field.TypeInt, value)
		_node.DislikeCount = value
	}
	if value, ok := rc.mutation.Flag(); ok {
		_spec.SetField(review.FieldFlag, field.TypeString, value)
		_node.Flag = value
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   review.UserTable,
			Columns: []string{review.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_reviews = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   review.BusinessTable,
			Columns: []string{review.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.review_business = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   review.PlaceTable,
			Columns: []string{review.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.review_place = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   review.EventTable,
			Columns: []string{review.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.review_event = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.MediasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.MediasTable,
			Columns: []string{review.MediasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.CommentsTable,
			Columns: []string{review.CommentsColumn},
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
	if nodes := rc.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   review.LikesTable,
			Columns: []string{review.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReviewCreateBulk is the builder for creating many Review entities in bulk.
type ReviewCreateBulk struct {
	config
	builders []*ReviewCreate
}

// Save creates the Review entities in the database.
func (rcb *ReviewCreateBulk) Save(ctx context.Context) ([]*Review, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Review, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReviewMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReviewCreateBulk) SaveX(ctx context.Context) []*Review {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReviewCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReviewCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
