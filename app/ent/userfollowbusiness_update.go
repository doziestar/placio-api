// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio_api/business"
	"placio_api/predicate"
	"placio_api/user"
	"placio_api/userfollowbusiness"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFollowBusinessUpdate is the builder for updating UserFollowBusiness entities.
type UserFollowBusinessUpdate struct {
	config
	hooks    []Hook
	mutation *UserFollowBusinessMutation
}

// Where appends a list predicates to the UserFollowBusinessUpdate builder.
func (ufbu *UserFollowBusinessUpdate) Where(ps ...predicate.UserFollowBusiness) *UserFollowBusinessUpdate {
	ufbu.mutation.Where(ps...)
	return ufbu
}

// SetCreatedAt sets the "CreatedAt" field.
func (ufbu *UserFollowBusinessUpdate) SetCreatedAt(t time.Time) *UserFollowBusinessUpdate {
	ufbu.mutation.SetCreatedAt(t)
	return ufbu
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (ufbu *UserFollowBusinessUpdate) SetNillableCreatedAt(t *time.Time) *UserFollowBusinessUpdate {
	if t != nil {
		ufbu.SetCreatedAt(*t)
	}
	return ufbu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (ufbu *UserFollowBusinessUpdate) SetUpdatedAt(t time.Time) *UserFollowBusinessUpdate {
	ufbu.mutation.SetUpdatedAt(t)
	return ufbu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ufbu *UserFollowBusinessUpdate) SetUserID(id string) *UserFollowBusinessUpdate {
	ufbu.mutation.SetUserID(id)
	return ufbu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ufbu *UserFollowBusinessUpdate) SetNillableUserID(id *string) *UserFollowBusinessUpdate {
	if id != nil {
		ufbu = ufbu.SetUserID(*id)
	}
	return ufbu
}

// SetUser sets the "user" edge to the User entity.
func (ufbu *UserFollowBusinessUpdate) SetUser(u *User) *UserFollowBusinessUpdate {
	return ufbu.SetUserID(u.ID)
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (ufbu *UserFollowBusinessUpdate) SetBusinessID(id string) *UserFollowBusinessUpdate {
	ufbu.mutation.SetBusinessID(id)
	return ufbu
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (ufbu *UserFollowBusinessUpdate) SetNillableBusinessID(id *string) *UserFollowBusinessUpdate {
	if id != nil {
		ufbu = ufbu.SetBusinessID(*id)
	}
	return ufbu
}

// SetBusiness sets the "business" edge to the Business entity.
func (ufbu *UserFollowBusinessUpdate) SetBusiness(b *Business) *UserFollowBusinessUpdate {
	return ufbu.SetBusinessID(b.ID)
}

// Mutation returns the UserFollowBusinessMutation object of the builder.
func (ufbu *UserFollowBusinessUpdate) Mutation() *UserFollowBusinessMutation {
	return ufbu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ufbu *UserFollowBusinessUpdate) ClearUser() *UserFollowBusinessUpdate {
	ufbu.mutation.ClearUser()
	return ufbu
}

// ClearBusiness clears the "business" edge to the Business entity.
func (ufbu *UserFollowBusinessUpdate) ClearBusiness() *UserFollowBusinessUpdate {
	ufbu.mutation.ClearBusiness()
	return ufbu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ufbu *UserFollowBusinessUpdate) Save(ctx context.Context) (int, error) {
	ufbu.defaults()
	return withHooks(ctx, ufbu.sqlSave, ufbu.mutation, ufbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ufbu *UserFollowBusinessUpdate) SaveX(ctx context.Context) int {
	affected, err := ufbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ufbu *UserFollowBusinessUpdate) Exec(ctx context.Context) error {
	_, err := ufbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufbu *UserFollowBusinessUpdate) ExecX(ctx context.Context) {
	if err := ufbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufbu *UserFollowBusinessUpdate) defaults() {
	if _, ok := ufbu.mutation.UpdatedAt(); !ok {
		v := userfollowbusiness.UpdateDefaultUpdatedAt()
		ufbu.mutation.SetUpdatedAt(v)
	}
}

func (ufbu *UserFollowBusinessUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userfollowbusiness.Table, userfollowbusiness.Columns, sqlgraph.NewFieldSpec(userfollowbusiness.FieldID, field.TypeString))
	if ps := ufbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ufbu.mutation.CreatedAt(); ok {
		_spec.SetField(userfollowbusiness.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ufbu.mutation.UpdatedAt(); ok {
		_spec.SetField(userfollowbusiness.FieldUpdatedAt, field.TypeTime, value)
	}
	if ufbu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowbusiness.UserTable,
			Columns: []string{userfollowbusiness.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufbu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowbusiness.UserTable,
			Columns: []string{userfollowbusiness.UserColumn},
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
	if ufbu.mutation.BusinessCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowbusiness.BusinessTable,
			Columns: []string{userfollowbusiness.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufbu.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowbusiness.BusinessTable,
			Columns: []string{userfollowbusiness.BusinessColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ufbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userfollowbusiness.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ufbu.mutation.done = true
	return n, nil
}

// UserFollowBusinessUpdateOne is the builder for updating a single UserFollowBusiness entity.
type UserFollowBusinessUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserFollowBusinessMutation
}

// SetCreatedAt sets the "CreatedAt" field.
func (ufbuo *UserFollowBusinessUpdateOne) SetCreatedAt(t time.Time) *UserFollowBusinessUpdateOne {
	ufbuo.mutation.SetCreatedAt(t)
	return ufbuo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (ufbuo *UserFollowBusinessUpdateOne) SetNillableCreatedAt(t *time.Time) *UserFollowBusinessUpdateOne {
	if t != nil {
		ufbuo.SetCreatedAt(*t)
	}
	return ufbuo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (ufbuo *UserFollowBusinessUpdateOne) SetUpdatedAt(t time.Time) *UserFollowBusinessUpdateOne {
	ufbuo.mutation.SetUpdatedAt(t)
	return ufbuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ufbuo *UserFollowBusinessUpdateOne) SetUserID(id string) *UserFollowBusinessUpdateOne {
	ufbuo.mutation.SetUserID(id)
	return ufbuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ufbuo *UserFollowBusinessUpdateOne) SetNillableUserID(id *string) *UserFollowBusinessUpdateOne {
	if id != nil {
		ufbuo = ufbuo.SetUserID(*id)
	}
	return ufbuo
}

// SetUser sets the "user" edge to the User entity.
func (ufbuo *UserFollowBusinessUpdateOne) SetUser(u *User) *UserFollowBusinessUpdateOne {
	return ufbuo.SetUserID(u.ID)
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (ufbuo *UserFollowBusinessUpdateOne) SetBusinessID(id string) *UserFollowBusinessUpdateOne {
	ufbuo.mutation.SetBusinessID(id)
	return ufbuo
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (ufbuo *UserFollowBusinessUpdateOne) SetNillableBusinessID(id *string) *UserFollowBusinessUpdateOne {
	if id != nil {
		ufbuo = ufbuo.SetBusinessID(*id)
	}
	return ufbuo
}

// SetBusiness sets the "business" edge to the Business entity.
func (ufbuo *UserFollowBusinessUpdateOne) SetBusiness(b *Business) *UserFollowBusinessUpdateOne {
	return ufbuo.SetBusinessID(b.ID)
}

// Mutation returns the UserFollowBusinessMutation object of the builder.
func (ufbuo *UserFollowBusinessUpdateOne) Mutation() *UserFollowBusinessMutation {
	return ufbuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ufbuo *UserFollowBusinessUpdateOne) ClearUser() *UserFollowBusinessUpdateOne {
	ufbuo.mutation.ClearUser()
	return ufbuo
}

// ClearBusiness clears the "business" edge to the Business entity.
func (ufbuo *UserFollowBusinessUpdateOne) ClearBusiness() *UserFollowBusinessUpdateOne {
	ufbuo.mutation.ClearBusiness()
	return ufbuo
}

// Where appends a list predicates to the UserFollowBusinessUpdate builder.
func (ufbuo *UserFollowBusinessUpdateOne) Where(ps ...predicate.UserFollowBusiness) *UserFollowBusinessUpdateOne {
	ufbuo.mutation.Where(ps...)
	return ufbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ufbuo *UserFollowBusinessUpdateOne) Select(field string, fields ...string) *UserFollowBusinessUpdateOne {
	ufbuo.fields = append([]string{field}, fields...)
	return ufbuo
}

// Save executes the query and returns the updated UserFollowBusiness entity.
func (ufbuo *UserFollowBusinessUpdateOne) Save(ctx context.Context) (*UserFollowBusiness, error) {
	ufbuo.defaults()
	return withHooks(ctx, ufbuo.sqlSave, ufbuo.mutation, ufbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ufbuo *UserFollowBusinessUpdateOne) SaveX(ctx context.Context) *UserFollowBusiness {
	node, err := ufbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ufbuo *UserFollowBusinessUpdateOne) Exec(ctx context.Context) error {
	_, err := ufbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufbuo *UserFollowBusinessUpdateOne) ExecX(ctx context.Context) {
	if err := ufbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufbuo *UserFollowBusinessUpdateOne) defaults() {
	if _, ok := ufbuo.mutation.UpdatedAt(); !ok {
		v := userfollowbusiness.UpdateDefaultUpdatedAt()
		ufbuo.mutation.SetUpdatedAt(v)
	}
}

func (ufbuo *UserFollowBusinessUpdateOne) sqlSave(ctx context.Context) (_node *UserFollowBusiness, err error) {
	_spec := sqlgraph.NewUpdateSpec(userfollowbusiness.Table, userfollowbusiness.Columns, sqlgraph.NewFieldSpec(userfollowbusiness.FieldID, field.TypeString))
	id, ok := ufbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`placio_api: missing "UserFollowBusiness.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ufbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userfollowbusiness.FieldID)
		for _, f := range fields {
			if !userfollowbusiness.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
			}
			if f != userfollowbusiness.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ufbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ufbuo.mutation.CreatedAt(); ok {
		_spec.SetField(userfollowbusiness.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ufbuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userfollowbusiness.FieldUpdatedAt, field.TypeTime, value)
	}
	if ufbuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowbusiness.UserTable,
			Columns: []string{userfollowbusiness.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufbuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowbusiness.UserTable,
			Columns: []string{userfollowbusiness.UserColumn},
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
	if ufbuo.mutation.BusinessCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowbusiness.BusinessTable,
			Columns: []string{userfollowbusiness.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ufbuo.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowbusiness.BusinessTable,
			Columns: []string{userfollowbusiness.BusinessColumn},
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
	_node = &UserFollowBusiness{config: ufbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ufbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userfollowbusiness.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ufbuo.mutation.done = true
	return _node, nil
}
