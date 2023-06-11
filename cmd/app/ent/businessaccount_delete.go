// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"placio-app/ent/businessaccount"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BusinessAccountDelete is the builder for deleting a BusinessAccount entity.
type BusinessAccountDelete struct {
	config
	hooks    []Hook
	mutation *BusinessAccountMutation
}

// Where appends a list predicates to the BusinessAccountDelete builder.
func (bad *BusinessAccountDelete) Where(ps ...predicate.BusinessAccount) *BusinessAccountDelete {
	bad.mutation.Where(ps...)
	return bad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bad *BusinessAccountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bad.sqlExec, bad.mutation, bad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bad *BusinessAccountDelete) ExecX(ctx context.Context) int {
	n, err := bad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bad *BusinessAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(businessaccount.Table, sqlgraph.NewFieldSpec(businessaccount.FieldID, field.TypeInt))
	if ps := bad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bad.mutation.done = true
	return affected, err
}

// BusinessAccountDeleteOne is the builder for deleting a single BusinessAccount entity.
type BusinessAccountDeleteOne struct {
	bad *BusinessAccountDelete
}

// Where appends a list predicates to the BusinessAccountDelete builder.
func (bado *BusinessAccountDeleteOne) Where(ps ...predicate.BusinessAccount) *BusinessAccountDeleteOne {
	bado.bad.mutation.Where(ps...)
	return bado
}

// Exec executes the deletion query.
func (bado *BusinessAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := bado.bad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{businessaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bado *BusinessAccountDeleteOne) ExecX(ctx context.Context) {
	if err := bado.Exec(ctx); err != nil {
		panic(err)
	}
}
