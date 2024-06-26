// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"placio-app/ent/placeinventoryattribute"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaceInventoryAttributeDelete is the builder for deleting a PlaceInventoryAttribute entity.
type PlaceInventoryAttributeDelete struct {
	config
	hooks    []Hook
	mutation *PlaceInventoryAttributeMutation
}

// Where appends a list predicates to the PlaceInventoryAttributeDelete builder.
func (piad *PlaceInventoryAttributeDelete) Where(ps ...predicate.PlaceInventoryAttribute) *PlaceInventoryAttributeDelete {
	piad.mutation.Where(ps...)
	return piad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (piad *PlaceInventoryAttributeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, piad.sqlExec, piad.mutation, piad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (piad *PlaceInventoryAttributeDelete) ExecX(ctx context.Context) int {
	n, err := piad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (piad *PlaceInventoryAttributeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(placeinventoryattribute.Table, sqlgraph.NewFieldSpec(placeinventoryattribute.FieldID, field.TypeString))
	if ps := piad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, piad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	piad.mutation.done = true
	return affected, err
}

// PlaceInventoryAttributeDeleteOne is the builder for deleting a single PlaceInventoryAttribute entity.
type PlaceInventoryAttributeDeleteOne struct {
	piad *PlaceInventoryAttributeDelete
}

// Where appends a list predicates to the PlaceInventoryAttributeDelete builder.
func (piado *PlaceInventoryAttributeDeleteOne) Where(ps ...predicate.PlaceInventoryAttribute) *PlaceInventoryAttributeDeleteOne {
	piado.piad.mutation.Where(ps...)
	return piado
}

// Exec executes the deletion query.
func (piado *PlaceInventoryAttributeDeleteOne) Exec(ctx context.Context) error {
	n, err := piado.piad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{placeinventoryattribute.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (piado *PlaceInventoryAttributeDeleteOne) ExecX(ctx context.Context) {
	if err := piado.Exec(ctx); err != nil {
		panic(err)
	}
}
