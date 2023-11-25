// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/predicate"
	"placio-app/ent/user"
	"placio-app/ent/userfollowuser"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFollowUserUpdate is the builder for updating UserFollowUser entities.
type UserFollowUserUpdate struct {
	config
	hooks    []Hook
	mutation *UserFollowUserMutation
}

// Where appends a list predicates to the UserFollowUserUpdate builder.
func (ufuu *UserFollowUserUpdate) Where(ps ...predicate.UserFollowUser) *UserFollowUserUpdate {
	ufuu.mutation.Where(ps...)
	return ufuu
}

// SetCreatedAt sets the "CreatedAt" field.
func (ufuu *UserFollowUserUpdate) SetCreatedAt(t time.Time) *UserFollowUserUpdate {
	ufuu.mutation.SetCreatedAt(t)
	return ufuu
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (ufuu *UserFollowUserUpdate) SetNillableCreatedAt(t *time.Time) *UserFollowUserUpdate {
	if t != nil {
		ufuu.SetCreatedAt(*t)
	}
	return ufuu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (ufuu *UserFollowUserUpdate) SetUpdatedAt(t time.Time) *UserFollowUserUpdate {
	ufuu.mutation.SetUpdatedAt(t)
	return ufuu
}

// SetFollowerID sets the "follower" edge to the User entity by ID.
func (ufuu *UserFollowUserUpdate) SetFollowerID(id string) *UserFollowUserUpdate {
	ufuu.mutation.SetFollowerID(id)
	return ufuu
}

// SetNillableFollowerID sets the "follower" edge to the User entity by ID if the given value is not nil.
func (ufuu *UserFollowUserUpdate) SetNillableFollowerID(id *string) *UserFollowUserUpdate {
	if id != nil {
		ufuu = ufuu.SetFollowerID(*id)
	}
	return ufuu
}

// SetFollower sets the "follower" edge to the User entity.
func (ufuu *UserFollowUserUpdate) SetFollower(u *User) *UserFollowUserUpdate {
	return ufuu.SetFollowerID(u.ID)
}

// SetFollowedID sets the "followed" edge to the User entity by ID.
func (ufuu *UserFollowUserUpdate) SetFollowedID(id string) *UserFollowUserUpdate {
	ufuu.mutation.SetFollowedID(id)
	return ufuu
}

// SetNillableFollowedID sets the "followed" edge to the User entity by ID if the given value is not nil.
func (ufuu *UserFollowUserUpdate) SetNillableFollowedID(id *string) *UserFollowUserUpdate {
	if id != nil {
		ufuu = ufuu.SetFollowedID(*id)
	}
	return ufuu
}

// SetFollowed sets the "followed" edge to the User entity.
func (ufuu *UserFollowUserUpdate) SetFollowed(u *User) *UserFollowUserUpdate {
	return ufuu.SetFollowedID(u.ID)
}

// Mutation returns the UserFollowUserMutation object of the builder.
func (ufuu *UserFollowUserUpdate) Mutation() *UserFollowUserMutation {
	return ufuu.mutation
}

// ClearFollower clears the "follower" edge to the User entity.
func (ufuu *UserFollowUserUpdate) ClearFollower() *UserFollowUserUpdate {
	ufuu.mutation.ClearFollower()
	return ufuu
}

// ClearFollowed clears the "followed" edge to the User entity.
func (ufuu *UserFollowUserUpdate) ClearFollowed() *UserFollowUserUpdate {
	ufuu.mutation.ClearFollowed()
	return ufuu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ufuu *UserFollowUserUpdate) Save(ctx context.Context) (int, error) {
	ufuu.defaults()
	return withHooks(ctx, ufuu.sqlSave, ufuu.mutation, ufuu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ufuu *UserFollowUserUpdate) SaveX(ctx context.Context) int {
	affected, err := ufuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ufuu *UserFollowUserUpdate) Exec(ctx context.Context) error {
	_, err := ufuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufuu *UserFollowUserUpdate) ExecX(ctx context.Context) {
	if err := ufuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufuu *UserFollowUserUpdate) defaults() {
	if _, ok := ufuu.mutation.UpdatedAt(); !ok {
		v := userfollowuser.UpdateDefaultUpdatedAt()
		ufuu.mutation.SetUpdatedAt(v)
	}
}

func (ufuu *UserFollowUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userfollowuser.Table, userfollowuser.Columns, sqlgraph.NewFieldSpec(userfollowuser.FieldID, field.TypeString))
	if ps := ufuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ufuu.mutation.CreatedAt(); ok {
		_spec.SetField(userfollowuser.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ufuu.mutation.UpdatedAt(); ok {
		_spec.SetField(userfollowuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if ufuu.mutation.FollowerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowuser.FollowerTable,
			Columns: []string{userfollowuser.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufuu.mutation.FollowerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowuser.FollowerTable,
			Columns: []string{userfollowuser.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ufuu.mutation.FollowedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowuser.FollowedTable,
			Columns: []string{userfollowuser.FollowedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufuu.mutation.FollowedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowuser.FollowedTable,
			Columns: []string{userfollowuser.FollowedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ufuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userfollowuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ufuu.mutation.done = true
	return n, nil
}

// UserFollowUserUpdateOne is the builder for updating a single UserFollowUser entity.
type UserFollowUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserFollowUserMutation
}

// SetCreatedAt sets the "CreatedAt" field.
func (ufuuo *UserFollowUserUpdateOne) SetCreatedAt(t time.Time) *UserFollowUserUpdateOne {
	ufuuo.mutation.SetCreatedAt(t)
	return ufuuo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (ufuuo *UserFollowUserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserFollowUserUpdateOne {
	if t != nil {
		ufuuo.SetCreatedAt(*t)
	}
	return ufuuo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (ufuuo *UserFollowUserUpdateOne) SetUpdatedAt(t time.Time) *UserFollowUserUpdateOne {
	ufuuo.mutation.SetUpdatedAt(t)
	return ufuuo
}

// SetFollowerID sets the "follower" edge to the User entity by ID.
func (ufuuo *UserFollowUserUpdateOne) SetFollowerID(id string) *UserFollowUserUpdateOne {
	ufuuo.mutation.SetFollowerID(id)
	return ufuuo
}

// SetNillableFollowerID sets the "follower" edge to the User entity by ID if the given value is not nil.
func (ufuuo *UserFollowUserUpdateOne) SetNillableFollowerID(id *string) *UserFollowUserUpdateOne {
	if id != nil {
		ufuuo = ufuuo.SetFollowerID(*id)
	}
	return ufuuo
}

// SetFollower sets the "follower" edge to the User entity.
func (ufuuo *UserFollowUserUpdateOne) SetFollower(u *User) *UserFollowUserUpdateOne {
	return ufuuo.SetFollowerID(u.ID)
}

// SetFollowedID sets the "followed" edge to the User entity by ID.
func (ufuuo *UserFollowUserUpdateOne) SetFollowedID(id string) *UserFollowUserUpdateOne {
	ufuuo.mutation.SetFollowedID(id)
	return ufuuo
}

// SetNillableFollowedID sets the "followed" edge to the User entity by ID if the given value is not nil.
func (ufuuo *UserFollowUserUpdateOne) SetNillableFollowedID(id *string) *UserFollowUserUpdateOne {
	if id != nil {
		ufuuo = ufuuo.SetFollowedID(*id)
	}
	return ufuuo
}

// SetFollowed sets the "followed" edge to the User entity.
func (ufuuo *UserFollowUserUpdateOne) SetFollowed(u *User) *UserFollowUserUpdateOne {
	return ufuuo.SetFollowedID(u.ID)
}

// Mutation returns the UserFollowUserMutation object of the builder.
func (ufuuo *UserFollowUserUpdateOne) Mutation() *UserFollowUserMutation {
	return ufuuo.mutation
}

// ClearFollower clears the "follower" edge to the User entity.
func (ufuuo *UserFollowUserUpdateOne) ClearFollower() *UserFollowUserUpdateOne {
	ufuuo.mutation.ClearFollower()
	return ufuuo
}

// ClearFollowed clears the "followed" edge to the User entity.
func (ufuuo *UserFollowUserUpdateOne) ClearFollowed() *UserFollowUserUpdateOne {
	ufuuo.mutation.ClearFollowed()
	return ufuuo
}

// Where appends a list predicates to the UserFollowUserUpdate builder.
func (ufuuo *UserFollowUserUpdateOne) Where(ps ...predicate.UserFollowUser) *UserFollowUserUpdateOne {
	ufuuo.mutation.Where(ps...)
	return ufuuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ufuuo *UserFollowUserUpdateOne) Select(field string, fields ...string) *UserFollowUserUpdateOne {
	ufuuo.fields = append([]string{field}, fields...)
	return ufuuo
}

// Save executes the query and returns the updated UserFollowUser entity.
func (ufuuo *UserFollowUserUpdateOne) Save(ctx context.Context) (*UserFollowUser, error) {
	ufuuo.defaults()
	return withHooks(ctx, ufuuo.sqlSave, ufuuo.mutation, ufuuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ufuuo *UserFollowUserUpdateOne) SaveX(ctx context.Context) *UserFollowUser {
	node, err := ufuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ufuuo *UserFollowUserUpdateOne) Exec(ctx context.Context) error {
	_, err := ufuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufuuo *UserFollowUserUpdateOne) ExecX(ctx context.Context) {
	if err := ufuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufuuo *UserFollowUserUpdateOne) defaults() {
	if _, ok := ufuuo.mutation.UpdatedAt(); !ok {
		v := userfollowuser.UpdateDefaultUpdatedAt()
		ufuuo.mutation.SetUpdatedAt(v)
	}
}

func (ufuuo *UserFollowUserUpdateOne) sqlSave(ctx context.Context) (_node *UserFollowUser, err error) {
	_spec := sqlgraph.NewUpdateSpec(userfollowuser.Table, userfollowuser.Columns, sqlgraph.NewFieldSpec(userfollowuser.FieldID, field.TypeString))
	id, ok := ufuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserFollowUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ufuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userfollowuser.FieldID)
		for _, f := range fields {
			if !userfollowuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userfollowuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ufuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ufuuo.mutation.CreatedAt(); ok {
		_spec.SetField(userfollowuser.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ufuuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userfollowuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if ufuuo.mutation.FollowerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowuser.FollowerTable,
			Columns: []string{userfollowuser.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufuuo.mutation.FollowerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowuser.FollowerTable,
			Columns: []string{userfollowuser.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ufuuo.mutation.FollowedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowuser.FollowedTable,
			Columns: []string{userfollowuser.FollowedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufuuo.mutation.FollowedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowuser.FollowedTable,
			Columns: []string{userfollowuser.FollowedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserFollowUser{config: ufuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ufuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userfollowuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ufuuo.mutation.done = true
	return _node, nil
}
