// Code generated by ent, DO NOT EDIT.

package placio_api

import (
	"context"
	"placio_api/media"
	"placio_api/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MediaDelete is the builder for deleting a Media entity.
type MediaDelete struct {
	config
	hooks    []Hook
	mutation *MediaMutation
}

// Where appends a list predicates to the MediaDelete builder.
func (md *MediaDelete) Where(ps ...predicate.Media) *MediaDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MediaDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MediaDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MediaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(media.Table, sqlgraph.NewFieldSpec(media.FieldID, field.TypeString))
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MediaDeleteOne is the builder for deleting a single Media entity.
type MediaDeleteOne struct {
	md *MediaDelete
}

// Where appends a list predicates to the MediaDelete builder.
func (mdo *MediaDeleteOne) Where(ps ...predicate.Media) *MediaDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MediaDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{media.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MediaDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}
