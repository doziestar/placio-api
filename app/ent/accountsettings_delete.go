// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"placio-app/ent/accountsettings"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountSettingsDelete is the builder for deleting a AccountSettings entity.
type AccountSettingsDelete struct {
	config
	hooks    []Hook
	mutation *AccountSettingsMutation
}

// Where appends a list predicates to the AccountSettingsDelete builder.
func (asd *AccountSettingsDelete) Where(ps ...predicate.AccountSettings) *AccountSettingsDelete {
	asd.mutation.Where(ps...)
	return asd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (asd *AccountSettingsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, asd.sqlExec, asd.mutation, asd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (asd *AccountSettingsDelete) ExecX(ctx context.Context) int {
	n, err := asd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (asd *AccountSettingsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(accountsettings.Table, sqlgraph.NewFieldSpec(accountsettings.FieldID, field.TypeString))
	if ps := asd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, asd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	asd.mutation.done = true
	return affected, err
}

// AccountSettingsDeleteOne is the builder for deleting a single AccountSettings entity.
type AccountSettingsDeleteOne struct {
	asd *AccountSettingsDelete
}

// Where appends a list predicates to the AccountSettingsDelete builder.
func (asdo *AccountSettingsDeleteOne) Where(ps ...predicate.AccountSettings) *AccountSettingsDeleteOne {
	asdo.asd.mutation.Where(ps...)
	return asdo
}

// Exec executes the deletion query.
func (asdo *AccountSettingsDeleteOne) Exec(ctx context.Context) error {
	n, err := asdo.asd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accountsettings.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (asdo *AccountSettingsDeleteOne) ExecX(ctx context.Context) {
	if err := asdo.Exec(ctx); err != nil {
		panic(err)
	}
}
