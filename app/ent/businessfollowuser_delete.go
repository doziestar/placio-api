



// Code generated by ent, DO NOT EDIT.



package ent



import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
		"placio-app/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
			"database/sql/driver"
			"entgo.io/ent/dialect/sql"
			"entgo.io/ent/dialect/sql/sqlgraph"
			"entgo.io/ent/dialect/sql/sqljson"
			"entgo.io/ent/schema/field"

)


import (
	 "placio-app/ent/businessfollowuser"
)





// BusinessFollowUserDelete is the builder for deleting a BusinessFollowUser entity.
type BusinessFollowUserDelete struct {
	config
	hooks      []Hook
	mutation   *BusinessFollowUserMutation
}

// Where appends a list predicates to the BusinessFollowUserDelete builder.
func (bfud *BusinessFollowUserDelete) Where(ps ...predicate.BusinessFollowUser) *BusinessFollowUserDelete {
	bfud.mutation.Where(ps...)
	return bfud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bfud *BusinessFollowUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bfud.sqlExec, bfud.mutation, bfud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bfud *BusinessFollowUserDelete) ExecX(ctx context.Context) int {
	n, err := bfud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}


	
	




func (bfud *BusinessFollowUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(businessfollowuser.Table, sqlgraph.NewFieldSpec(businessfollowuser.FieldID, field.TypeString))
	if ps := bfud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bfud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bfud.mutation.done = true
	return affected, err
}







// BusinessFollowUserDeleteOne is the builder for deleting a single BusinessFollowUser entity.
type BusinessFollowUserDeleteOne struct {
	bfud *BusinessFollowUserDelete
}

// Where appends a list predicates to the BusinessFollowUserDelete builder.
func (bfudo *BusinessFollowUserDeleteOne) Where(ps ...predicate.BusinessFollowUser) *BusinessFollowUserDeleteOne {
	bfudo.bfud.mutation.Where(ps...)
	return bfudo
}

// Exec executes the deletion query.
func (bfudo *BusinessFollowUserDeleteOne) Exec(ctx context.Context) error {
	n, err := bfudo.bfud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ businessfollowuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bfudo *BusinessFollowUserDeleteOne) ExecX(ctx context.Context) {
	if err := bfudo.Exec(ctx); err != nil {
		panic(err)
	}
}

