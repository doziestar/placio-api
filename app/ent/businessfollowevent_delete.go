// Code generated by ent, DO NOT EDIT.

package placio_api

import (
	"context"
	"placio_api/businessfollowevent"
	"placio_api/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BusinessFollowEventDelete is the builder for deleting a BusinessFollowEvent entity.
type BusinessFollowEventDelete struct {
	config
	hooks    []Hook
	mutation *BusinessFollowEventMutation
}

// Where appends a list predicates to the BusinessFollowEventDelete builder.
func (bfed *BusinessFollowEventDelete) Where(ps ...predicate.BusinessFollowEvent) *BusinessFollowEventDelete {
	bfed.mutation.Where(ps...)
	return bfed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bfed *BusinessFollowEventDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bfed.sqlExec, bfed.mutation, bfed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bfed *BusinessFollowEventDelete) ExecX(ctx context.Context) int {
	n, err := bfed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bfed *BusinessFollowEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(businessfollowevent.Table, sqlgraph.NewFieldSpec(businessfollowevent.FieldID, field.TypeString))
	if ps := bfed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bfed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bfed.mutation.done = true
	return affected, err
}

// BusinessFollowEventDeleteOne is the builder for deleting a single BusinessFollowEvent entity.
type BusinessFollowEventDeleteOne struct {
	bfed *BusinessFollowEventDelete
}

// Where appends a list predicates to the BusinessFollowEventDelete builder.
func (bfedo *BusinessFollowEventDeleteOne) Where(ps ...predicate.BusinessFollowEvent) *BusinessFollowEventDeleteOne {
	bfedo.bfed.mutation.Where(ps...)
	return bfedo
}

// Exec executes the deletion query.
func (bfedo *BusinessFollowEventDeleteOne) Exec(ctx context.Context) error {
	n, err := bfedo.bfed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{businessfollowevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bfedo *BusinessFollowEventDeleteOne) ExecX(ctx context.Context) {
	if err := bfedo.Exec(ctx); err != nil {
		panic(err)
	}
}
