// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"placio_api/inventoryattribute"
	"placio_api/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InventoryAttributeDelete is the builder for deleting a InventoryAttribute entity.
type InventoryAttributeDelete struct {
	config
	hooks    []Hook
	mutation *InventoryAttributeMutation
}

// Where appends a list predicates to the InventoryAttributeDelete builder.
func (iad *InventoryAttributeDelete) Where(ps ...predicate.InventoryAttribute) *InventoryAttributeDelete {
	iad.mutation.Where(ps...)
	return iad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (iad *InventoryAttributeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, iad.sqlExec, iad.mutation, iad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (iad *InventoryAttributeDelete) ExecX(ctx context.Context) int {
	n, err := iad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (iad *InventoryAttributeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(inventoryattribute.Table, sqlgraph.NewFieldSpec(inventoryattribute.FieldID, field.TypeString))
	if ps := iad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, iad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	iad.mutation.done = true
	return affected, err
}

// InventoryAttributeDeleteOne is the builder for deleting a single InventoryAttribute entity.
type InventoryAttributeDeleteOne struct {
	iad *InventoryAttributeDelete
}

// Where appends a list predicates to the InventoryAttributeDelete builder.
func (iado *InventoryAttributeDeleteOne) Where(ps ...predicate.InventoryAttribute) *InventoryAttributeDeleteOne {
	iado.iad.mutation.Where(ps...)
	return iado
}

// Exec executes the deletion query.
func (iado *InventoryAttributeDeleteOne) Exec(ctx context.Context) error {
	n, err := iado.iad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{inventoryattribute.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (iado *InventoryAttributeDeleteOne) ExecX(ctx context.Context) {
	if err := iado.Exec(ctx); err != nil {
		panic(err)
	}
}
