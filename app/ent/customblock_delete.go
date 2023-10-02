// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"placio_api/customblock"
	"placio_api/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomBlockDelete is the builder for deleting a CustomBlock entity.
type CustomBlockDelete struct {
	config
	hooks    []Hook
	mutation *CustomBlockMutation
}

// Where appends a list predicates to the CustomBlockDelete builder.
func (cbd *CustomBlockDelete) Where(ps ...predicate.CustomBlock) *CustomBlockDelete {
	cbd.mutation.Where(ps...)
	return cbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cbd *CustomBlockDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cbd.sqlExec, cbd.mutation, cbd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cbd *CustomBlockDelete) ExecX(ctx context.Context) int {
	n, err := cbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cbd *CustomBlockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(customblock.Table, sqlgraph.NewFieldSpec(customblock.FieldID, field.TypeString))
	if ps := cbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cbd.mutation.done = true
	return affected, err
}

// CustomBlockDeleteOne is the builder for deleting a single CustomBlock entity.
type CustomBlockDeleteOne struct {
	cbd *CustomBlockDelete
}

// Where appends a list predicates to the CustomBlockDelete builder.
func (cbdo *CustomBlockDeleteOne) Where(ps ...predicate.CustomBlock) *CustomBlockDeleteOne {
	cbdo.cbd.mutation.Where(ps...)
	return cbdo
}

// Exec executes the deletion query.
func (cbdo *CustomBlockDeleteOne) Exec(ctx context.Context) error {
	n, err := cbdo.cbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{customblock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cbdo *CustomBlockDeleteOne) ExecX(ctx context.Context) {
	if err := cbdo.Exec(ctx); err != nil {
		panic(err)
	}
}
