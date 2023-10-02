// Code generated by ent, DO NOT EDIT.

package placio_api

import (
	"context"
	"errors"
	"fmt"
	"placio_api/business"
	"placio_api/businessfollowbusiness"
	"placio_api/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BusinessFollowBusinessUpdate is the builder for updating BusinessFollowBusiness entities.
type BusinessFollowBusinessUpdate struct {
	config
	hooks    []Hook
	mutation *BusinessFollowBusinessMutation
}

// Where appends a list predicates to the BusinessFollowBusinessUpdate builder.
func (bfbu *BusinessFollowBusinessUpdate) Where(ps ...predicate.BusinessFollowBusiness) *BusinessFollowBusinessUpdate {
	bfbu.mutation.Where(ps...)
	return bfbu
}

// SetCreatedAt sets the "CreatedAt" field.
func (bfbu *BusinessFollowBusinessUpdate) SetCreatedAt(t time.Time) *BusinessFollowBusinessUpdate {
	bfbu.mutation.SetCreatedAt(t)
	return bfbu
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (bfbu *BusinessFollowBusinessUpdate) SetNillableCreatedAt(t *time.Time) *BusinessFollowBusinessUpdate {
	if t != nil {
		bfbu.SetCreatedAt(*t)
	}
	return bfbu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (bfbu *BusinessFollowBusinessUpdate) SetUpdatedAt(t time.Time) *BusinessFollowBusinessUpdate {
	bfbu.mutation.SetUpdatedAt(t)
	return bfbu
}

// SetFollowerID sets the "follower" edge to the Business entity by ID.
func (bfbu *BusinessFollowBusinessUpdate) SetFollowerID(id string) *BusinessFollowBusinessUpdate {
	bfbu.mutation.SetFollowerID(id)
	return bfbu
}

// SetNillableFollowerID sets the "follower" edge to the Business entity by ID if the given value is not nil.
func (bfbu *BusinessFollowBusinessUpdate) SetNillableFollowerID(id *string) *BusinessFollowBusinessUpdate {
	if id != nil {
		bfbu = bfbu.SetFollowerID(*id)
	}
	return bfbu
}

// SetFollower sets the "follower" edge to the Business entity.
func (bfbu *BusinessFollowBusinessUpdate) SetFollower(b *Business) *BusinessFollowBusinessUpdate {
	return bfbu.SetFollowerID(b.ID)
}

// SetFollowedID sets the "followed" edge to the Business entity by ID.
func (bfbu *BusinessFollowBusinessUpdate) SetFollowedID(id string) *BusinessFollowBusinessUpdate {
	bfbu.mutation.SetFollowedID(id)
	return bfbu
}

// SetNillableFollowedID sets the "followed" edge to the Business entity by ID if the given value is not nil.
func (bfbu *BusinessFollowBusinessUpdate) SetNillableFollowedID(id *string) *BusinessFollowBusinessUpdate {
	if id != nil {
		bfbu = bfbu.SetFollowedID(*id)
	}
	return bfbu
}

// SetFollowed sets the "followed" edge to the Business entity.
func (bfbu *BusinessFollowBusinessUpdate) SetFollowed(b *Business) *BusinessFollowBusinessUpdate {
	return bfbu.SetFollowedID(b.ID)
}

// Mutation returns the BusinessFollowBusinessMutation object of the builder.
func (bfbu *BusinessFollowBusinessUpdate) Mutation() *BusinessFollowBusinessMutation {
	return bfbu.mutation
}

// ClearFollower clears the "follower" edge to the Business entity.
func (bfbu *BusinessFollowBusinessUpdate) ClearFollower() *BusinessFollowBusinessUpdate {
	bfbu.mutation.ClearFollower()
	return bfbu
}

// ClearFollowed clears the "followed" edge to the Business entity.
func (bfbu *BusinessFollowBusinessUpdate) ClearFollowed() *BusinessFollowBusinessUpdate {
	bfbu.mutation.ClearFollowed()
	return bfbu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bfbu *BusinessFollowBusinessUpdate) Save(ctx context.Context) (int, error) {
	bfbu.defaults()
	return withHooks(ctx, bfbu.sqlSave, bfbu.mutation, bfbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bfbu *BusinessFollowBusinessUpdate) SaveX(ctx context.Context) int {
	affected, err := bfbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bfbu *BusinessFollowBusinessUpdate) Exec(ctx context.Context) error {
	_, err := bfbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bfbu *BusinessFollowBusinessUpdate) ExecX(ctx context.Context) {
	if err := bfbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bfbu *BusinessFollowBusinessUpdate) defaults() {
	if _, ok := bfbu.mutation.UpdatedAt(); !ok {
		v := businessfollowbusiness.UpdateDefaultUpdatedAt()
		bfbu.mutation.SetUpdatedAt(v)
	}
}

func (bfbu *BusinessFollowBusinessUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(businessfollowbusiness.Table, businessfollowbusiness.Columns, sqlgraph.NewFieldSpec(businessfollowbusiness.FieldID, field.TypeString))
	if ps := bfbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bfbu.mutation.CreatedAt(); ok {
		_spec.SetField(businessfollowbusiness.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bfbu.mutation.UpdatedAt(); ok {
		_spec.SetField(businessfollowbusiness.FieldUpdatedAt, field.TypeTime, value)
	}
	if bfbu.mutation.FollowerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowbusiness.FollowerTable,
			Columns: []string{businessfollowbusiness.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bfbu.mutation.FollowerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowbusiness.FollowerTable,
			Columns: []string{businessfollowbusiness.FollowerColumn},
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
	if bfbu.mutation.FollowedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowbusiness.FollowedTable,
			Columns: []string{businessfollowbusiness.FollowedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bfbu.mutation.FollowedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowbusiness.FollowedTable,
			Columns: []string{businessfollowbusiness.FollowedColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, bfbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{businessfollowbusiness.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bfbu.mutation.done = true
	return n, nil
}

// BusinessFollowBusinessUpdateOne is the builder for updating a single BusinessFollowBusiness entity.
type BusinessFollowBusinessUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BusinessFollowBusinessMutation
}

// SetCreatedAt sets the "CreatedAt" field.
func (bfbuo *BusinessFollowBusinessUpdateOne) SetCreatedAt(t time.Time) *BusinessFollowBusinessUpdateOne {
	bfbuo.mutation.SetCreatedAt(t)
	return bfbuo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (bfbuo *BusinessFollowBusinessUpdateOne) SetNillableCreatedAt(t *time.Time) *BusinessFollowBusinessUpdateOne {
	if t != nil {
		bfbuo.SetCreatedAt(*t)
	}
	return bfbuo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (bfbuo *BusinessFollowBusinessUpdateOne) SetUpdatedAt(t time.Time) *BusinessFollowBusinessUpdateOne {
	bfbuo.mutation.SetUpdatedAt(t)
	return bfbuo
}

// SetFollowerID sets the "follower" edge to the Business entity by ID.
func (bfbuo *BusinessFollowBusinessUpdateOne) SetFollowerID(id string) *BusinessFollowBusinessUpdateOne {
	bfbuo.mutation.SetFollowerID(id)
	return bfbuo
}

// SetNillableFollowerID sets the "follower" edge to the Business entity by ID if the given value is not nil.
func (bfbuo *BusinessFollowBusinessUpdateOne) SetNillableFollowerID(id *string) *BusinessFollowBusinessUpdateOne {
	if id != nil {
		bfbuo = bfbuo.SetFollowerID(*id)
	}
	return bfbuo
}

// SetFollower sets the "follower" edge to the Business entity.
func (bfbuo *BusinessFollowBusinessUpdateOne) SetFollower(b *Business) *BusinessFollowBusinessUpdateOne {
	return bfbuo.SetFollowerID(b.ID)
}

// SetFollowedID sets the "followed" edge to the Business entity by ID.
func (bfbuo *BusinessFollowBusinessUpdateOne) SetFollowedID(id string) *BusinessFollowBusinessUpdateOne {
	bfbuo.mutation.SetFollowedID(id)
	return bfbuo
}

// SetNillableFollowedID sets the "followed" edge to the Business entity by ID if the given value is not nil.
func (bfbuo *BusinessFollowBusinessUpdateOne) SetNillableFollowedID(id *string) *BusinessFollowBusinessUpdateOne {
	if id != nil {
		bfbuo = bfbuo.SetFollowedID(*id)
	}
	return bfbuo
}

// SetFollowed sets the "followed" edge to the Business entity.
func (bfbuo *BusinessFollowBusinessUpdateOne) SetFollowed(b *Business) *BusinessFollowBusinessUpdateOne {
	return bfbuo.SetFollowedID(b.ID)
}

// Mutation returns the BusinessFollowBusinessMutation object of the builder.
func (bfbuo *BusinessFollowBusinessUpdateOne) Mutation() *BusinessFollowBusinessMutation {
	return bfbuo.mutation
}

// ClearFollower clears the "follower" edge to the Business entity.
func (bfbuo *BusinessFollowBusinessUpdateOne) ClearFollower() *BusinessFollowBusinessUpdateOne {
	bfbuo.mutation.ClearFollower()
	return bfbuo
}

// ClearFollowed clears the "followed" edge to the Business entity.
func (bfbuo *BusinessFollowBusinessUpdateOne) ClearFollowed() *BusinessFollowBusinessUpdateOne {
	bfbuo.mutation.ClearFollowed()
	return bfbuo
}

// Where appends a list predicates to the BusinessFollowBusinessUpdate builder.
func (bfbuo *BusinessFollowBusinessUpdateOne) Where(ps ...predicate.BusinessFollowBusiness) *BusinessFollowBusinessUpdateOne {
	bfbuo.mutation.Where(ps...)
	return bfbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bfbuo *BusinessFollowBusinessUpdateOne) Select(field string, fields ...string) *BusinessFollowBusinessUpdateOne {
	bfbuo.fields = append([]string{field}, fields...)
	return bfbuo
}

// Save executes the query and returns the updated BusinessFollowBusiness entity.
func (bfbuo *BusinessFollowBusinessUpdateOne) Save(ctx context.Context) (*BusinessFollowBusiness, error) {
	bfbuo.defaults()
	return withHooks(ctx, bfbuo.sqlSave, bfbuo.mutation, bfbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bfbuo *BusinessFollowBusinessUpdateOne) SaveX(ctx context.Context) *BusinessFollowBusiness {
	node, err := bfbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bfbuo *BusinessFollowBusinessUpdateOne) Exec(ctx context.Context) error {
	_, err := bfbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bfbuo *BusinessFollowBusinessUpdateOne) ExecX(ctx context.Context) {
	if err := bfbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bfbuo *BusinessFollowBusinessUpdateOne) defaults() {
	if _, ok := bfbuo.mutation.UpdatedAt(); !ok {
		v := businessfollowbusiness.UpdateDefaultUpdatedAt()
		bfbuo.mutation.SetUpdatedAt(v)
	}
}

func (bfbuo *BusinessFollowBusinessUpdateOne) sqlSave(ctx context.Context) (_node *BusinessFollowBusiness, err error) {
	_spec := sqlgraph.NewUpdateSpec(businessfollowbusiness.Table, businessfollowbusiness.Columns, sqlgraph.NewFieldSpec(businessfollowbusiness.FieldID, field.TypeString))
	id, ok := bfbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`placio_api: missing "BusinessFollowBusiness.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bfbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, businessfollowbusiness.FieldID)
		for _, f := range fields {
			if !businessfollowbusiness.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
			}
			if f != businessfollowbusiness.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bfbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bfbuo.mutation.CreatedAt(); ok {
		_spec.SetField(businessfollowbusiness.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bfbuo.mutation.UpdatedAt(); ok {
		_spec.SetField(businessfollowbusiness.FieldUpdatedAt, field.TypeTime, value)
	}
	if bfbuo.mutation.FollowerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowbusiness.FollowerTable,
			Columns: []string{businessfollowbusiness.FollowerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bfbuo.mutation.FollowerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowbusiness.FollowerTable,
			Columns: []string{businessfollowbusiness.FollowerColumn},
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
	if bfbuo.mutation.FollowedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowbusiness.FollowedTable,
			Columns: []string{businessfollowbusiness.FollowedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bfbuo.mutation.FollowedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   businessfollowbusiness.FollowedTable,
			Columns: []string{businessfollowbusiness.FollowedColumn},
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
	_node = &BusinessFollowBusiness{config: bfbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bfbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{businessfollowbusiness.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bfbuo.mutation.done = true
	return _node, nil
}
