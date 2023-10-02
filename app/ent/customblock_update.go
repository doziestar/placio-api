// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio_api/customblock"
	"placio_api/predicate"
	"placio_api/website"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomBlockUpdate is the builder for updating CustomBlock entities.
type CustomBlockUpdate struct {
	config
	hooks    []Hook
	mutation *CustomBlockMutation
}

// Where appends a list predicates to the CustomBlockUpdate builder.
func (cbu *CustomBlockUpdate) Where(ps ...predicate.CustomBlock) *CustomBlockUpdate {
	cbu.mutation.Where(ps...)
	return cbu
}

// SetContent sets the "content" field.
func (cbu *CustomBlockUpdate) SetContent(s string) *CustomBlockUpdate {
	cbu.mutation.SetContent(s)
	return cbu
}

// SetWebsiteID sets the "website" edge to the Website entity by ID.
func (cbu *CustomBlockUpdate) SetWebsiteID(id string) *CustomBlockUpdate {
	cbu.mutation.SetWebsiteID(id)
	return cbu
}

// SetWebsite sets the "website" edge to the Website entity.
func (cbu *CustomBlockUpdate) SetWebsite(w *Website) *CustomBlockUpdate {
	return cbu.SetWebsiteID(w.ID)
}

// Mutation returns the CustomBlockMutation object of the builder.
func (cbu *CustomBlockUpdate) Mutation() *CustomBlockMutation {
	return cbu.mutation
}

// ClearWebsite clears the "website" edge to the Website entity.
func (cbu *CustomBlockUpdate) ClearWebsite() *CustomBlockUpdate {
	cbu.mutation.ClearWebsite()
	return cbu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cbu *CustomBlockUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cbu.sqlSave, cbu.mutation, cbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cbu *CustomBlockUpdate) SaveX(ctx context.Context) int {
	affected, err := cbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cbu *CustomBlockUpdate) Exec(ctx context.Context) error {
	_, err := cbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbu *CustomBlockUpdate) ExecX(ctx context.Context) {
	if err := cbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbu *CustomBlockUpdate) check() error {
	if _, ok := cbu.mutation.WebsiteID(); cbu.mutation.WebsiteCleared() && !ok {
		return errors.New(`placio_api: clearing a required unique edge "CustomBlock.website"`)
	}
	return nil
}

func (cbu *CustomBlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cbu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(customblock.Table, customblock.Columns, sqlgraph.NewFieldSpec(customblock.FieldID, field.TypeString))
	if ps := cbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbu.mutation.Content(); ok {
		_spec.SetField(customblock.FieldContent, field.TypeString, value)
	}
	if cbu.mutation.WebsiteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customblock.WebsiteTable,
			Columns: []string{customblock.WebsiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbu.mutation.WebsiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customblock.WebsiteTable,
			Columns: []string{customblock.WebsiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customblock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cbu.mutation.done = true
	return n, nil
}

// CustomBlockUpdateOne is the builder for updating a single CustomBlock entity.
type CustomBlockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomBlockMutation
}

// SetContent sets the "content" field.
func (cbuo *CustomBlockUpdateOne) SetContent(s string) *CustomBlockUpdateOne {
	cbuo.mutation.SetContent(s)
	return cbuo
}

// SetWebsiteID sets the "website" edge to the Website entity by ID.
func (cbuo *CustomBlockUpdateOne) SetWebsiteID(id string) *CustomBlockUpdateOne {
	cbuo.mutation.SetWebsiteID(id)
	return cbuo
}

// SetWebsite sets the "website" edge to the Website entity.
func (cbuo *CustomBlockUpdateOne) SetWebsite(w *Website) *CustomBlockUpdateOne {
	return cbuo.SetWebsiteID(w.ID)
}

// Mutation returns the CustomBlockMutation object of the builder.
func (cbuo *CustomBlockUpdateOne) Mutation() *CustomBlockMutation {
	return cbuo.mutation
}

// ClearWebsite clears the "website" edge to the Website entity.
func (cbuo *CustomBlockUpdateOne) ClearWebsite() *CustomBlockUpdateOne {
	cbuo.mutation.ClearWebsite()
	return cbuo
}

// Where appends a list predicates to the CustomBlockUpdate builder.
func (cbuo *CustomBlockUpdateOne) Where(ps ...predicate.CustomBlock) *CustomBlockUpdateOne {
	cbuo.mutation.Where(ps...)
	return cbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cbuo *CustomBlockUpdateOne) Select(field string, fields ...string) *CustomBlockUpdateOne {
	cbuo.fields = append([]string{field}, fields...)
	return cbuo
}

// Save executes the query and returns the updated CustomBlock entity.
func (cbuo *CustomBlockUpdateOne) Save(ctx context.Context) (*CustomBlock, error) {
	return withHooks(ctx, cbuo.sqlSave, cbuo.mutation, cbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cbuo *CustomBlockUpdateOne) SaveX(ctx context.Context) *CustomBlock {
	node, err := cbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cbuo *CustomBlockUpdateOne) Exec(ctx context.Context) error {
	_, err := cbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbuo *CustomBlockUpdateOne) ExecX(ctx context.Context) {
	if err := cbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbuo *CustomBlockUpdateOne) check() error {
	if _, ok := cbuo.mutation.WebsiteID(); cbuo.mutation.WebsiteCleared() && !ok {
		return errors.New(`placio_api: clearing a required unique edge "CustomBlock.website"`)
	}
	return nil
}

func (cbuo *CustomBlockUpdateOne) sqlSave(ctx context.Context) (_node *CustomBlock, err error) {
	if err := cbuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(customblock.Table, customblock.Columns, sqlgraph.NewFieldSpec(customblock.FieldID, field.TypeString))
	id, ok := cbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`placio_api: missing "CustomBlock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customblock.FieldID)
		for _, f := range fields {
			if !customblock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
			}
			if f != customblock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbuo.mutation.Content(); ok {
		_spec.SetField(customblock.FieldContent, field.TypeString, value)
	}
	if cbuo.mutation.WebsiteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customblock.WebsiteTable,
			Columns: []string{customblock.WebsiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbuo.mutation.WebsiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   customblock.WebsiteTable,
			Columns: []string{customblock.WebsiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CustomBlock{config: cbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customblock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cbuo.mutation.done = true
	return _node, nil
}
