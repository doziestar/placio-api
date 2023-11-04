// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"placio-app/ent/predicate"
	"placio-app/ent/staff"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StaffDelete is the builder for deleting a Staff entity.
type StaffDelete struct {
	config
	hooks    []Hook
	mutation *StaffMutation
}

// Where appends a list predicates to the StaffDelete builder.
func (sd *StaffDelete) Where(ps ...predicate.Staff) *StaffDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *StaffDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *StaffDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *StaffDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(staff.Table, sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// StaffDeleteOne is the builder for deleting a single Staff entity.
type StaffDeleteOne struct {
	sd *StaffDelete
}

// Where appends a list predicates to the StaffDelete builder.
func (sdo *StaffDeleteOne) Where(ps ...predicate.Staff) *StaffDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *StaffDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{staff.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *StaffDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}