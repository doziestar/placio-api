// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio_api/predicate"
	"placio_api/resourse"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ResourseUpdate is the builder for updating Resourse entities.
type ResourseUpdate struct {
	config
	hooks    []Hook
	mutation *ResourseMutation
}

// Where appends a list predicates to the ResourseUpdate builder.
func (ru *ResourseUpdate) Where(ps ...predicate.Resourse) *ResourseUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// Mutation returns the ResourseMutation object of the builder.
func (ru *ResourseUpdate) Mutation() *ResourseMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ResourseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResourseUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResourseUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResourseUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ResourseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(resourse.Table, resourse.Columns, sqlgraph.NewFieldSpec(resourse.FieldID, field.TypeString))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ResourseUpdateOne is the builder for updating a single Resourse entity.
type ResourseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResourseMutation
}

// Mutation returns the ResourseMutation object of the builder.
func (ruo *ResourseUpdateOne) Mutation() *ResourseMutation {
	return ruo.mutation
}

// Where appends a list predicates to the ResourseUpdate builder.
func (ruo *ResourseUpdateOne) Where(ps ...predicate.Resourse) *ResourseUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ResourseUpdateOne) Select(field string, fields ...string) *ResourseUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Resourse entity.
func (ruo *ResourseUpdateOne) Save(ctx context.Context) (*Resourse, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResourseUpdateOne) SaveX(ctx context.Context) *Resourse {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ResourseUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResourseUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ResourseUpdateOne) sqlSave(ctx context.Context) (_node *Resourse, err error) {
	_spec := sqlgraph.NewUpdateSpec(resourse.Table, resourse.Columns, sqlgraph.NewFieldSpec(resourse.FieldID, field.TypeString))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`placio_api: missing "Resourse.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resourse.FieldID)
		for _, f := range fields {
			if !resourse.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
			}
			if f != resourse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &Resourse{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resourse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
