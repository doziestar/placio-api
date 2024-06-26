// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/category"
	"placio-app/ent/comment"
	"placio-app/ent/like"
	"placio-app/ent/media"
	"placio-app/ent/notification"
	"placio-app/ent/post"
	"placio-app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
}

// SetContent sets the "Content" field.
func (pc *PostCreate) SetContent(s string) *PostCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetNillableContent sets the "Content" field if the given value is not nil.
func (pc *PostCreate) SetNillableContent(s *string) *PostCreate {
	if s != nil {
		pc.SetContent(*s)
	}
	return pc
}

// SetCreatedAt sets the "CreatedAt" field.
func (pc *PostCreate) SetCreatedAt(t time.Time) *PostCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (pc *PostCreate) SetUpdatedAt(t time.Time) *PostCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetPrivacy sets the "Privacy" field.
func (pc *PostCreate) SetPrivacy(po post.Privacy) *PostCreate {
	pc.mutation.SetPrivacy(po)
	return pc
}

// SetNillablePrivacy sets the "Privacy" field if the given value is not nil.
func (pc *PostCreate) SetNillablePrivacy(po *post.Privacy) *PostCreate {
	if po != nil {
		pc.SetPrivacy(*po)
	}
	return pc
}

// SetLikedByMe sets the "LikedByMe" field.
func (pc *PostCreate) SetLikedByMe(b bool) *PostCreate {
	pc.mutation.SetLikedByMe(b)
	return pc
}

// SetNillableLikedByMe sets the "LikedByMe" field if the given value is not nil.
func (pc *PostCreate) SetNillableLikedByMe(b *bool) *PostCreate {
	if b != nil {
		pc.SetLikedByMe(*b)
	}
	return pc
}

// SetLikeCount sets the "LikeCount" field.
func (pc *PostCreate) SetLikeCount(i int) *PostCreate {
	pc.mutation.SetLikeCount(i)
	return pc
}

// SetNillableLikeCount sets the "LikeCount" field if the given value is not nil.
func (pc *PostCreate) SetNillableLikeCount(i *int) *PostCreate {
	if i != nil {
		pc.SetLikeCount(*i)
	}
	return pc
}

// SetCommentCount sets the "CommentCount" field.
func (pc *PostCreate) SetCommentCount(i int) *PostCreate {
	pc.mutation.SetCommentCount(i)
	return pc
}

// SetNillableCommentCount sets the "CommentCount" field if the given value is not nil.
func (pc *PostCreate) SetNillableCommentCount(i *int) *PostCreate {
	if i != nil {
		pc.SetCommentCount(*i)
	}
	return pc
}

// SetShareCount sets the "ShareCount" field.
func (pc *PostCreate) SetShareCount(i int) *PostCreate {
	pc.mutation.SetShareCount(i)
	return pc
}

// SetNillableShareCount sets the "ShareCount" field if the given value is not nil.
func (pc *PostCreate) SetNillableShareCount(i *int) *PostCreate {
	if i != nil {
		pc.SetShareCount(*i)
	}
	return pc
}

// SetViewCount sets the "ViewCount" field.
func (pc *PostCreate) SetViewCount(i int) *PostCreate {
	pc.mutation.SetViewCount(i)
	return pc
}

// SetNillableViewCount sets the "ViewCount" field if the given value is not nil.
func (pc *PostCreate) SetNillableViewCount(i *int) *PostCreate {
	if i != nil {
		pc.SetViewCount(*i)
	}
	return pc
}

// SetIsSponsored sets the "IsSponsored" field.
func (pc *PostCreate) SetIsSponsored(b bool) *PostCreate {
	pc.mutation.SetIsSponsored(b)
	return pc
}

// SetNillableIsSponsored sets the "IsSponsored" field if the given value is not nil.
func (pc *PostCreate) SetNillableIsSponsored(b *bool) *PostCreate {
	if b != nil {
		pc.SetIsSponsored(*b)
	}
	return pc
}

// SetIsPromoted sets the "IsPromoted" field.
func (pc *PostCreate) SetIsPromoted(b bool) *PostCreate {
	pc.mutation.SetIsPromoted(b)
	return pc
}

// SetNillableIsPromoted sets the "IsPromoted" field if the given value is not nil.
func (pc *PostCreate) SetNillableIsPromoted(b *bool) *PostCreate {
	if b != nil {
		pc.SetIsPromoted(*b)
	}
	return pc
}

// SetIsBoosted sets the "IsBoosted" field.
func (pc *PostCreate) SetIsBoosted(b bool) *PostCreate {
	pc.mutation.SetIsBoosted(b)
	return pc
}

// SetNillableIsBoosted sets the "IsBoosted" field if the given value is not nil.
func (pc *PostCreate) SetNillableIsBoosted(b *bool) *PostCreate {
	if b != nil {
		pc.SetIsBoosted(*b)
	}
	return pc
}

// SetIsPinned sets the "IsPinned" field.
func (pc *PostCreate) SetIsPinned(b bool) *PostCreate {
	pc.mutation.SetIsPinned(b)
	return pc
}

// SetNillableIsPinned sets the "IsPinned" field if the given value is not nil.
func (pc *PostCreate) SetNillableIsPinned(b *bool) *PostCreate {
	if b != nil {
		pc.SetIsPinned(*b)
	}
	return pc
}

// SetIsHidden sets the "IsHidden" field.
func (pc *PostCreate) SetIsHidden(b bool) *PostCreate {
	pc.mutation.SetIsHidden(b)
	return pc
}

// SetNillableIsHidden sets the "IsHidden" field if the given value is not nil.
func (pc *PostCreate) SetNillableIsHidden(b *bool) *PostCreate {
	if b != nil {
		pc.SetIsHidden(*b)
	}
	return pc
}

// SetRepostCount sets the "RepostCount" field.
func (pc *PostCreate) SetRepostCount(i int) *PostCreate {
	pc.mutation.SetRepostCount(i)
	return pc
}

// SetNillableRepostCount sets the "RepostCount" field if the given value is not nil.
func (pc *PostCreate) SetNillableRepostCount(i *int) *PostCreate {
	if i != nil {
		pc.SetRepostCount(*i)
	}
	return pc
}

// SetIsRepost sets the "IsRepost" field.
func (pc *PostCreate) SetIsRepost(b bool) *PostCreate {
	pc.mutation.SetIsRepost(b)
	return pc
}

// SetNillableIsRepost sets the "IsRepost" field if the given value is not nil.
func (pc *PostCreate) SetNillableIsRepost(b *bool) *PostCreate {
	if b != nil {
		pc.SetIsRepost(*b)
	}
	return pc
}

// SetRelevanceScore sets the "RelevanceScore" field.
func (pc *PostCreate) SetRelevanceScore(i int) *PostCreate {
	pc.mutation.SetRelevanceScore(i)
	return pc
}

// SetNillableRelevanceScore sets the "RelevanceScore" field if the given value is not nil.
func (pc *PostCreate) SetNillableRelevanceScore(i *int) *PostCreate {
	if i != nil {
		pc.SetRelevanceScore(*i)
	}
	return pc
}

// SetSearchText sets the "SearchText" field.
func (pc *PostCreate) SetSearchText(s string) *PostCreate {
	pc.mutation.SetSearchText(s)
	return pc
}

// SetNillableSearchText sets the "SearchText" field if the given value is not nil.
func (pc *PostCreate) SetNillableSearchText(s *string) *PostCreate {
	if s != nil {
		pc.SetSearchText(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PostCreate) SetID(s string) *PostCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *PostCreate) SetUserID(id string) *PostCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableUserID(id *string) *PostCreate {
	if id != nil {
		pc = pc.SetUserID(*id)
	}
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PostCreate) SetUser(u *User) *PostCreate {
	return pc.SetUserID(u.ID)
}

// SetBusinessAccountID sets the "business_account" edge to the Business entity by ID.
func (pc *PostCreate) SetBusinessAccountID(id string) *PostCreate {
	pc.mutation.SetBusinessAccountID(id)
	return pc
}

// SetNillableBusinessAccountID sets the "business_account" edge to the Business entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableBusinessAccountID(id *string) *PostCreate {
	if id != nil {
		pc = pc.SetBusinessAccountID(*id)
	}
	return pc
}

// SetBusinessAccount sets the "business_account" edge to the Business entity.
func (pc *PostCreate) SetBusinessAccount(b *Business) *PostCreate {
	return pc.SetBusinessAccountID(b.ID)
}

// AddMediaIDs adds the "medias" edge to the Media entity by IDs.
func (pc *PostCreate) AddMediaIDs(ids ...string) *PostCreate {
	pc.mutation.AddMediaIDs(ids...)
	return pc
}

// AddMedias adds the "medias" edges to the Media entity.
func (pc *PostCreate) AddMedias(m ...*Media) *PostCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pc.AddMediaIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (pc *PostCreate) AddCommentIDs(ids ...string) *PostCreate {
	pc.mutation.AddCommentIDs(ids...)
	return pc
}

// AddComments adds the "comments" edges to the Comment entity.
func (pc *PostCreate) AddComments(c ...*Comment) *PostCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCommentIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (pc *PostCreate) AddLikeIDs(ids ...string) *PostCreate {
	pc.mutation.AddLikeIDs(ids...)
	return pc
}

// AddLikes adds the "likes" edges to the Like entity.
func (pc *PostCreate) AddLikes(l ...*Like) *PostCreate {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pc.AddLikeIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (pc *PostCreate) AddCategoryIDs(ids ...string) *PostCreate {
	pc.mutation.AddCategoryIDs(ids...)
	return pc
}

// AddCategories adds the "categories" edges to the Category entity.
func (pc *PostCreate) AddCategories(c ...*Category) *PostCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCategoryIDs(ids...)
}

// AddNotificationIDs adds the "notifications" edge to the Notification entity by IDs.
func (pc *PostCreate) AddNotificationIDs(ids ...string) *PostCreate {
	pc.mutation.AddNotificationIDs(ids...)
	return pc
}

// AddNotifications adds the "notifications" edges to the Notification entity.
func (pc *PostCreate) AddNotifications(n ...*Notification) *PostCreate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return pc.AddNotificationIDs(ids...)
}

// SetRepostsID sets the "reposts" edge to the Post entity by ID.
func (pc *PostCreate) SetRepostsID(id string) *PostCreate {
	pc.mutation.SetRepostsID(id)
	return pc
}

// SetNillableRepostsID sets the "reposts" edge to the Post entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableRepostsID(id *string) *PostCreate {
	if id != nil {
		pc = pc.SetRepostsID(*id)
	}
	return pc
}

// SetReposts sets the "reposts" edge to the Post entity.
func (pc *PostCreate) SetReposts(p *Post) *PostCreate {
	return pc.SetRepostsID(p.ID)
}

// AddOriginalPostIDs adds the "original_post" edge to the Post entity by IDs.
func (pc *PostCreate) AddOriginalPostIDs(ids ...string) *PostCreate {
	pc.mutation.AddOriginalPostIDs(ids...)
	return pc
}

// AddOriginalPost adds the "original_post" edges to the Post entity.
func (pc *PostCreate) AddOriginalPost(p ...*Post) *PostCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddOriginalPostIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := post.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.Privacy(); !ok {
		v := post.DefaultPrivacy
		pc.mutation.SetPrivacy(v)
	}
	if _, ok := pc.mutation.LikedByMe(); !ok {
		v := post.DefaultLikedByMe
		pc.mutation.SetLikedByMe(v)
	}
	if _, ok := pc.mutation.LikeCount(); !ok {
		v := post.DefaultLikeCount
		pc.mutation.SetLikeCount(v)
	}
	if _, ok := pc.mutation.CommentCount(); !ok {
		v := post.DefaultCommentCount
		pc.mutation.SetCommentCount(v)
	}
	if _, ok := pc.mutation.ShareCount(); !ok {
		v := post.DefaultShareCount
		pc.mutation.SetShareCount(v)
	}
	if _, ok := pc.mutation.ViewCount(); !ok {
		v := post.DefaultViewCount
		pc.mutation.SetViewCount(v)
	}
	if _, ok := pc.mutation.IsSponsored(); !ok {
		v := post.DefaultIsSponsored
		pc.mutation.SetIsSponsored(v)
	}
	if _, ok := pc.mutation.IsPromoted(); !ok {
		v := post.DefaultIsPromoted
		pc.mutation.SetIsPromoted(v)
	}
	if _, ok := pc.mutation.IsBoosted(); !ok {
		v := post.DefaultIsBoosted
		pc.mutation.SetIsBoosted(v)
	}
	if _, ok := pc.mutation.IsPinned(); !ok {
		v := post.DefaultIsPinned
		pc.mutation.SetIsPinned(v)
	}
	if _, ok := pc.mutation.IsHidden(); !ok {
		v := post.DefaultIsHidden
		pc.mutation.SetIsHidden(v)
	}
	if _, ok := pc.mutation.RepostCount(); !ok {
		v := post.DefaultRepostCount
		pc.mutation.SetRepostCount(v)
	}
	if _, ok := pc.mutation.IsRepost(); !ok {
		v := post.DefaultIsRepost
		pc.mutation.SetIsRepost(v)
	}
	if _, ok := pc.mutation.RelevanceScore(); !ok {
		v := post.DefaultRelevanceScore
		pc.mutation.SetRelevanceScore(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if v, ok := pc.mutation.Content(); ok {
		if err := post.ContentValidator(v); err != nil {
			return &ValidationError{Name: "Content", err: fmt.Errorf(`ent: validator failed for field "Post.Content": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "Post.CreatedAt"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "Post.UpdatedAt"`)}
	}
	if _, ok := pc.mutation.Privacy(); !ok {
		return &ValidationError{Name: "Privacy", err: errors.New(`ent: missing required field "Post.Privacy"`)}
	}
	if v, ok := pc.mutation.Privacy(); ok {
		if err := post.PrivacyValidator(v); err != nil {
			return &ValidationError{Name: "Privacy", err: fmt.Errorf(`ent: validator failed for field "Post.Privacy": %w`, err)}
		}
	}
	if _, ok := pc.mutation.LikedByMe(); !ok {
		return &ValidationError{Name: "LikedByMe", err: errors.New(`ent: missing required field "Post.LikedByMe"`)}
	}
	if _, ok := pc.mutation.LikeCount(); !ok {
		return &ValidationError{Name: "LikeCount", err: errors.New(`ent: missing required field "Post.LikeCount"`)}
	}
	if _, ok := pc.mutation.CommentCount(); !ok {
		return &ValidationError{Name: "CommentCount", err: errors.New(`ent: missing required field "Post.CommentCount"`)}
	}
	if _, ok := pc.mutation.ShareCount(); !ok {
		return &ValidationError{Name: "ShareCount", err: errors.New(`ent: missing required field "Post.ShareCount"`)}
	}
	if _, ok := pc.mutation.ViewCount(); !ok {
		return &ValidationError{Name: "ViewCount", err: errors.New(`ent: missing required field "Post.ViewCount"`)}
	}
	if _, ok := pc.mutation.IsSponsored(); !ok {
		return &ValidationError{Name: "IsSponsored", err: errors.New(`ent: missing required field "Post.IsSponsored"`)}
	}
	if _, ok := pc.mutation.IsPromoted(); !ok {
		return &ValidationError{Name: "IsPromoted", err: errors.New(`ent: missing required field "Post.IsPromoted"`)}
	}
	if _, ok := pc.mutation.IsBoosted(); !ok {
		return &ValidationError{Name: "IsBoosted", err: errors.New(`ent: missing required field "Post.IsBoosted"`)}
	}
	if _, ok := pc.mutation.IsPinned(); !ok {
		return &ValidationError{Name: "IsPinned", err: errors.New(`ent: missing required field "Post.IsPinned"`)}
	}
	if _, ok := pc.mutation.IsHidden(); !ok {
		return &ValidationError{Name: "IsHidden", err: errors.New(`ent: missing required field "Post.IsHidden"`)}
	}
	if _, ok := pc.mutation.RepostCount(); !ok {
		return &ValidationError{Name: "RepostCount", err: errors.New(`ent: missing required field "Post.RepostCount"`)}
	}
	if _, ok := pc.mutation.IsRepost(); !ok {
		return &ValidationError{Name: "IsRepost", err: errors.New(`ent: missing required field "Post.IsRepost"`)}
	}
	if _, ok := pc.mutation.RelevanceScore(); !ok {
		return &ValidationError{Name: "RelevanceScore", err: errors.New(`ent: missing required field "Post.RelevanceScore"`)}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := post.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Post.id": %w`, err)}
		}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Post.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(post.Table, sqlgraph.NewFieldSpec(post.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Privacy(); ok {
		_spec.SetField(post.FieldPrivacy, field.TypeEnum, value)
		_node.Privacy = value
	}
	if value, ok := pc.mutation.LikedByMe(); ok {
		_spec.SetField(post.FieldLikedByMe, field.TypeBool, value)
		_node.LikedByMe = value
	}
	if value, ok := pc.mutation.LikeCount(); ok {
		_spec.SetField(post.FieldLikeCount, field.TypeInt, value)
		_node.LikeCount = value
	}
	if value, ok := pc.mutation.CommentCount(); ok {
		_spec.SetField(post.FieldCommentCount, field.TypeInt, value)
		_node.CommentCount = value
	}
	if value, ok := pc.mutation.ShareCount(); ok {
		_spec.SetField(post.FieldShareCount, field.TypeInt, value)
		_node.ShareCount = value
	}
	if value, ok := pc.mutation.ViewCount(); ok {
		_spec.SetField(post.FieldViewCount, field.TypeInt, value)
		_node.ViewCount = value
	}
	if value, ok := pc.mutation.IsSponsored(); ok {
		_spec.SetField(post.FieldIsSponsored, field.TypeBool, value)
		_node.IsSponsored = value
	}
	if value, ok := pc.mutation.IsPromoted(); ok {
		_spec.SetField(post.FieldIsPromoted, field.TypeBool, value)
		_node.IsPromoted = value
	}
	if value, ok := pc.mutation.IsBoosted(); ok {
		_spec.SetField(post.FieldIsBoosted, field.TypeBool, value)
		_node.IsBoosted = value
	}
	if value, ok := pc.mutation.IsPinned(); ok {
		_spec.SetField(post.FieldIsPinned, field.TypeBool, value)
		_node.IsPinned = value
	}
	if value, ok := pc.mutation.IsHidden(); ok {
		_spec.SetField(post.FieldIsHidden, field.TypeBool, value)
		_node.IsHidden = value
	}
	if value, ok := pc.mutation.RepostCount(); ok {
		_spec.SetField(post.FieldRepostCount, field.TypeInt, value)
		_node.RepostCount = value
	}
	if value, ok := pc.mutation.IsRepost(); ok {
		_spec.SetField(post.FieldIsRepost, field.TypeBool, value)
		_node.IsRepost = value
	}
	if value, ok := pc.mutation.RelevanceScore(); ok {
		_spec.SetField(post.FieldRelevanceScore, field.TypeInt, value)
		_node.RelevanceScore = value
	}
	if value, ok := pc.mutation.SearchText(); ok {
		_spec.SetField(post.FieldSearchText, field.TypeString, value)
		_node.SearchText = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.UserTable,
			Columns: []string{post.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_posts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BusinessAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.BusinessAccountTable,
			Columns: []string{post.BusinessAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_posts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MediasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.MediasTable,
			Columns: []string{post.MediasColumn},
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
	if nodes := pc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CommentsTable,
			Columns: []string{post.CommentsColumn},
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
	if nodes := pc.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.LikesTable,
			Columns: []string{post.LikesColumn},
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
	if nodes := pc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.CategoriesTable,
			Columns: []string{post.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.NotificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.NotificationsTable,
			Columns: post.NotificationsPrimaryKey,
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
	if nodes := pc.mutation.RepostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.RepostsTable,
			Columns: []string{post.RepostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.post_original_post = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OriginalPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.OriginalPostTable,
			Columns: []string{post.OriginalPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	err      error
	builders []*PostCreate
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
