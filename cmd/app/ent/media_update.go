// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/category"
	"placio-app/ent/media"
	"placio-app/ent/post"
	"placio-app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MediaUpdate is the builder for updating Media entities.
type MediaUpdate struct {
	config
	hooks    []Hook
	mutation *MediaMutation
}

// Where appends a list predicates to the MediaUpdate builder.
func (mu *MediaUpdate) Where(ps ...predicate.Media) *MediaUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetURL sets the "URL" field.
func (mu *MediaUpdate) SetURL(s string) *MediaUpdate {
	mu.mutation.SetURL(s)
	return mu
}

// SetMediaType sets the "MediaType" field.
func (mu *MediaUpdate) SetMediaType(s string) *MediaUpdate {
	mu.mutation.SetMediaType(s)
	return mu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (mu *MediaUpdate) SetUpdatedAt(t time.Time) *MediaUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (mu *MediaUpdate) SetPostID(id string) *MediaUpdate {
	mu.mutation.SetPostID(id)
	return mu
}

// SetNillablePostID sets the "post" edge to the Post entity by ID if the given value is not nil.
func (mu *MediaUpdate) SetNillablePostID(id *string) *MediaUpdate {
	if id != nil {
		mu = mu.SetPostID(*id)
	}
	return mu
}

// SetPost sets the "post" edge to the Post entity.
func (mu *MediaUpdate) SetPost(p *Post) *MediaUpdate {
	return mu.SetPostID(p.ID)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (mu *MediaUpdate) AddCategoryIDs(ids ...string) *MediaUpdate {
	mu.mutation.AddCategoryIDs(ids...)
	return mu
}

// AddCategories adds the "categories" edges to the Category entity.
func (mu *MediaUpdate) AddCategories(c ...*Category) *MediaUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mu.AddCategoryIDs(ids...)
}

// Mutation returns the MediaMutation object of the builder.
func (mu *MediaUpdate) Mutation() *MediaMutation {
	return mu.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (mu *MediaUpdate) ClearPost() *MediaUpdate {
	mu.mutation.ClearPost()
	return mu
}

// ClearCategories clears all "categories" edges to the Category entity.
func (mu *MediaUpdate) ClearCategories() *MediaUpdate {
	mu.mutation.ClearCategories()
	return mu
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (mu *MediaUpdate) RemoveCategoryIDs(ids ...string) *MediaUpdate {
	mu.mutation.RemoveCategoryIDs(ids...)
	return mu
}

// RemoveCategories removes "categories" edges to Category entities.
func (mu *MediaUpdate) RemoveCategories(c ...*Category) *MediaUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mu.RemoveCategoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MediaUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MediaUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MediaUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MediaUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MediaUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := media.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

func (mu *MediaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(media.Table, media.Columns, sqlgraph.NewFieldSpec(media.FieldID, field.TypeString))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.URL(); ok {
		_spec.SetField(media.FieldURL, field.TypeString, value)
	}
	if value, ok := mu.mutation.MediaType(); ok {
		_spec.SetField(media.FieldMediaType, field.TypeString, value)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(media.FieldUpdatedAt, field.TypeTime, value)
	}
	if mu.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.PostTable,
			Columns: []string{media.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.PostTable,
			Columns: []string{media.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.CategoriesTable,
			Columns: []string{media.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !mu.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.CategoriesTable,
			Columns: []string{media.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.CategoriesTable,
			Columns: []string{media.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MediaUpdateOne is the builder for updating a single Media entity.
type MediaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MediaMutation
}

// SetURL sets the "URL" field.
func (muo *MediaUpdateOne) SetURL(s string) *MediaUpdateOne {
	muo.mutation.SetURL(s)
	return muo
}

// SetMediaType sets the "MediaType" field.
func (muo *MediaUpdateOne) SetMediaType(s string) *MediaUpdateOne {
	muo.mutation.SetMediaType(s)
	return muo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (muo *MediaUpdateOne) SetUpdatedAt(t time.Time) *MediaUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (muo *MediaUpdateOne) SetPostID(id string) *MediaUpdateOne {
	muo.mutation.SetPostID(id)
	return muo
}

// SetNillablePostID sets the "post" edge to the Post entity by ID if the given value is not nil.
func (muo *MediaUpdateOne) SetNillablePostID(id *string) *MediaUpdateOne {
	if id != nil {
		muo = muo.SetPostID(*id)
	}
	return muo
}

// SetPost sets the "post" edge to the Post entity.
func (muo *MediaUpdateOne) SetPost(p *Post) *MediaUpdateOne {
	return muo.SetPostID(p.ID)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (muo *MediaUpdateOne) AddCategoryIDs(ids ...string) *MediaUpdateOne {
	muo.mutation.AddCategoryIDs(ids...)
	return muo
}

// AddCategories adds the "categories" edges to the Category entity.
func (muo *MediaUpdateOne) AddCategories(c ...*Category) *MediaUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return muo.AddCategoryIDs(ids...)
}

// Mutation returns the MediaMutation object of the builder.
func (muo *MediaUpdateOne) Mutation() *MediaMutation {
	return muo.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (muo *MediaUpdateOne) ClearPost() *MediaUpdateOne {
	muo.mutation.ClearPost()
	return muo
}

// ClearCategories clears all "categories" edges to the Category entity.
func (muo *MediaUpdateOne) ClearCategories() *MediaUpdateOne {
	muo.mutation.ClearCategories()
	return muo
}

// RemoveCategoryIDs removes the "categories" edge to Category entities by IDs.
func (muo *MediaUpdateOne) RemoveCategoryIDs(ids ...string) *MediaUpdateOne {
	muo.mutation.RemoveCategoryIDs(ids...)
	return muo
}

// RemoveCategories removes "categories" edges to Category entities.
func (muo *MediaUpdateOne) RemoveCategories(c ...*Category) *MediaUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return muo.RemoveCategoryIDs(ids...)
}

// Where appends a list predicates to the MediaUpdate builder.
func (muo *MediaUpdateOne) Where(ps ...predicate.Media) *MediaUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MediaUpdateOne) Select(field string, fields ...string) *MediaUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Media entity.
func (muo *MediaUpdateOne) Save(ctx context.Context) (*Media, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MediaUpdateOne) SaveX(ctx context.Context) *Media {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MediaUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MediaUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MediaUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := media.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

func (muo *MediaUpdateOne) sqlSave(ctx context.Context) (_node *Media, err error) {
	_spec := sqlgraph.NewUpdateSpec(media.Table, media.Columns, sqlgraph.NewFieldSpec(media.FieldID, field.TypeString))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Media.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, media.FieldID)
		for _, f := range fields {
			if !media.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != media.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.URL(); ok {
		_spec.SetField(media.FieldURL, field.TypeString, value)
	}
	if value, ok := muo.mutation.MediaType(); ok {
		_spec.SetField(media.FieldMediaType, field.TypeString, value)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(media.FieldUpdatedAt, field.TypeTime, value)
	}
	if muo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.PostTable,
			Columns: []string{media.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   media.PostTable,
			Columns: []string{media.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.CategoriesTable,
			Columns: []string{media.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedCategoriesIDs(); len(nodes) > 0 && !muo.mutation.CategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.CategoriesTable,
			Columns: []string{media.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   media.CategoriesTable,
			Columns: []string{media.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Media{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
