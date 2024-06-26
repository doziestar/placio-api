// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/businessfollowuser"
	"placio-app/ent/predicate"
	"placio-app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BusinessFollowUserUpdate is the builder for updating BusinessFollowUser entities.
type BusinessFollowUserUpdate struct {
	config
	hooks    []Hook
	mutation *BusinessFollowUserMutation
}

// Where appends a list predicates to the BusinessFollowUserUpdate builder.
func (bfuu *BusinessFollowUserUpdate) Where(ps ...predicate.BusinessFollowUser) *BusinessFollowUserUpdate {
	bfuu.mutation.Where(ps...)
	return bfuu
}

// SetCreatedAt sets the "CreatedAt" field.
func (bfuu *BusinessFollowUserUpdate) SetCreatedAt(t time.Time) *BusinessFollowUserUpdate {
	bfuu.mutation.SetCreatedAt(t)
	return bfuu
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (bfuu *BusinessFollowUserUpdate) SetNillableCreatedAt(t *time.Time) *BusinessFollowUserUpdate {
	if t != nil {
		bfuu.SetCreatedAt(*t)
	}
	return bfuu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (bfuu *BusinessFollowUserUpdate) SetUpdatedAt(t time.Time) *BusinessFollowUserUpdate {
	bfuu.mutation.SetUpdatedAt(t)
	return bfuu
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (bfuu *BusinessFollowUserUpdate) SetBusinessID(id string) *BusinessFollowUserUpdate {
	bfuu.mutation.SetBusinessID(id)
	return bfuu
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (bfuu *BusinessFollowUserUpdate) SetNillableBusinessID(id *string) *BusinessFollowUserUpdate {
	if id != nil {
		bfuu = bfuu.SetBusinessID(*id)
	}
	return bfuu
}

// SetBusiness sets the "business" edge to the Business entity.
func (bfuu *BusinessFollowUserUpdate) SetBusiness(b *Business) *BusinessFollowUserUpdate {
	return bfuu.SetBusinessID(b.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bfuu *BusinessFollowUserUpdate) SetUserID(id string) *BusinessFollowUserUpdate {
	bfuu.mutation.SetUserID(id)
	return bfuu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (bfuu *BusinessFollowUserUpdate) SetNillableUserID(id *string) *BusinessFollowUserUpdate {
	if id != nil {
		bfuu = bfuu.SetUserID(*id)
	}
	return bfuu
}

// SetUser sets the "user" edge to the User entity.
func (bfuu *BusinessFollowUserUpdate) SetUser(u *User) *BusinessFollowUserUpdate {
	return bfuu.SetUserID(u.ID)
}

// Mutation returns the BusinessFollowUserMutation object of the builder.
func (bfuu *BusinessFollowUserUpdate) Mutation() *BusinessFollowUserMutation {
	return bfuu.mutation
}

// ClearBusiness clears the "business" edge to the Business entity.
func (bfuu *BusinessFollowUserUpdate) ClearBusiness() *BusinessFollowUserUpdate {
	bfuu.mutation.ClearBusiness()
	return bfuu
}

// ClearUser clears the "user" edge to the User entity.
func (bfuu *BusinessFollowUserUpdate) ClearUser() *BusinessFollowUserUpdate {
	bfuu.mutation.ClearUser()
	return bfuu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bfuu *BusinessFollowUserUpdate) Save(ctx context.Context) (int, error) {
	bfuu.defaults()
	return withHooks(ctx, bfuu.sqlSave, bfuu.mutation, bfuu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bfuu *BusinessFollowUserUpdate) SaveX(ctx context.Context) int {
	affected, err := bfuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bfuu *BusinessFollowUserUpdate) Exec(ctx context.Context) error {
	_, err := bfuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bfuu *BusinessFollowUserUpdate) ExecX(ctx context.Context) {
	if err := bfuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bfuu *BusinessFollowUserUpdate) defaults() {
	if _, ok := bfuu.mutation.UpdatedAt(); !ok {
		v := businessfollowuser.UpdateDefaultUpdatedAt()
		bfuu.mutation.SetUpdatedAt(v)
	}
}

func (bfuu *BusinessFollowUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(businessfollowuser.Table, businessfollowuser.Columns, sqlgraph.NewFieldSpec(businessfollowuser.FieldID, field.TypeString))
	if ps := bfuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bfuu.mutation.CreatedAt(); ok {
		_spec.SetField(businessfollowuser.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bfuu.mutation.UpdatedAt(); ok {
		_spec.SetField(businessfollowuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if bfuu.mutation.BusinessCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowuser.BusinessTable,
			Columns: []string{businessfollowuser.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bfuu.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowuser.BusinessTable,
			Columns: []string{businessfollowuser.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bfuu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowuser.UserTable,
			Columns: []string{businessfollowuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bfuu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowuser.UserTable,
			Columns: []string{businessfollowuser.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, bfuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{businessfollowuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bfuu.mutation.done = true
	return n, nil
}

// BusinessFollowUserUpdateOne is the builder for updating a single BusinessFollowUser entity.
type BusinessFollowUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BusinessFollowUserMutation
}

// SetCreatedAt sets the "CreatedAt" field.
func (bfuuo *BusinessFollowUserUpdateOne) SetCreatedAt(t time.Time) *BusinessFollowUserUpdateOne {
	bfuuo.mutation.SetCreatedAt(t)
	return bfuuo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (bfuuo *BusinessFollowUserUpdateOne) SetNillableCreatedAt(t *time.Time) *BusinessFollowUserUpdateOne {
	if t != nil {
		bfuuo.SetCreatedAt(*t)
	}
	return bfuuo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (bfuuo *BusinessFollowUserUpdateOne) SetUpdatedAt(t time.Time) *BusinessFollowUserUpdateOne {
	bfuuo.mutation.SetUpdatedAt(t)
	return bfuuo
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (bfuuo *BusinessFollowUserUpdateOne) SetBusinessID(id string) *BusinessFollowUserUpdateOne {
	bfuuo.mutation.SetBusinessID(id)
	return bfuuo
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (bfuuo *BusinessFollowUserUpdateOne) SetNillableBusinessID(id *string) *BusinessFollowUserUpdateOne {
	if id != nil {
		bfuuo = bfuuo.SetBusinessID(*id)
	}
	return bfuuo
}

// SetBusiness sets the "business" edge to the Business entity.
func (bfuuo *BusinessFollowUserUpdateOne) SetBusiness(b *Business) *BusinessFollowUserUpdateOne {
	return bfuuo.SetBusinessID(b.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bfuuo *BusinessFollowUserUpdateOne) SetUserID(id string) *BusinessFollowUserUpdateOne {
	bfuuo.mutation.SetUserID(id)
	return bfuuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (bfuuo *BusinessFollowUserUpdateOne) SetNillableUserID(id *string) *BusinessFollowUserUpdateOne {
	if id != nil {
		bfuuo = bfuuo.SetUserID(*id)
	}
	return bfuuo
}

// SetUser sets the "user" edge to the User entity.
func (bfuuo *BusinessFollowUserUpdateOne) SetUser(u *User) *BusinessFollowUserUpdateOne {
	return bfuuo.SetUserID(u.ID)
}

// Mutation returns the BusinessFollowUserMutation object of the builder.
func (bfuuo *BusinessFollowUserUpdateOne) Mutation() *BusinessFollowUserMutation {
	return bfuuo.mutation
}

// ClearBusiness clears the "business" edge to the Business entity.
func (bfuuo *BusinessFollowUserUpdateOne) ClearBusiness() *BusinessFollowUserUpdateOne {
	bfuuo.mutation.ClearBusiness()
	return bfuuo
}

// ClearUser clears the "user" edge to the User entity.
func (bfuuo *BusinessFollowUserUpdateOne) ClearUser() *BusinessFollowUserUpdateOne {
	bfuuo.mutation.ClearUser()
	return bfuuo
}

// Where appends a list predicates to the BusinessFollowUserUpdate builder.
func (bfuuo *BusinessFollowUserUpdateOne) Where(ps ...predicate.BusinessFollowUser) *BusinessFollowUserUpdateOne {
	bfuuo.mutation.Where(ps...)
	return bfuuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bfuuo *BusinessFollowUserUpdateOne) Select(field string, fields ...string) *BusinessFollowUserUpdateOne {
	bfuuo.fields = append([]string{field}, fields...)
	return bfuuo
}

// Save executes the query and returns the updated BusinessFollowUser entity.
func (bfuuo *BusinessFollowUserUpdateOne) Save(ctx context.Context) (*BusinessFollowUser, error) {
	bfuuo.defaults()
	return withHooks(ctx, bfuuo.sqlSave, bfuuo.mutation, bfuuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bfuuo *BusinessFollowUserUpdateOne) SaveX(ctx context.Context) *BusinessFollowUser {
	node, err := bfuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bfuuo *BusinessFollowUserUpdateOne) Exec(ctx context.Context) error {
	_, err := bfuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bfuuo *BusinessFollowUserUpdateOne) ExecX(ctx context.Context) {
	if err := bfuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bfuuo *BusinessFollowUserUpdateOne) defaults() {
	if _, ok := bfuuo.mutation.UpdatedAt(); !ok {
		v := businessfollowuser.UpdateDefaultUpdatedAt()
		bfuuo.mutation.SetUpdatedAt(v)
	}
}

func (bfuuo *BusinessFollowUserUpdateOne) sqlSave(ctx context.Context) (_node *BusinessFollowUser, err error) {
	_spec := sqlgraph.NewUpdateSpec(businessfollowuser.Table, businessfollowuser.Columns, sqlgraph.NewFieldSpec(businessfollowuser.FieldID, field.TypeString))
	id, ok := bfuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BusinessFollowUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bfuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, businessfollowuser.FieldID)
		for _, f := range fields {
			if !businessfollowuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != businessfollowuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bfuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bfuuo.mutation.CreatedAt(); ok {
		_spec.SetField(businessfollowuser.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bfuuo.mutation.UpdatedAt(); ok {
		_spec.SetField(businessfollowuser.FieldUpdatedAt, field.TypeTime, value)
	}
	if bfuuo.mutation.BusinessCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowuser.BusinessTable,
			Columns: []string{businessfollowuser.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bfuuo.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowuser.BusinessTable,
			Columns: []string{businessfollowuser.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bfuuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowuser.UserTable,
			Columns: []string{businessfollowuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bfuuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowuser.UserTable,
			Columns: []string{businessfollowuser.UserColumn},
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
	_node = &BusinessFollowUser{config: bfuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bfuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{businessfollowuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bfuuo.mutation.done = true
	return _node, nil
}
