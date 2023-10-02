// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio_api/business"
	"placio_api/businessfollowbusiness"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BusinessFollowBusinessCreate is the builder for creating a BusinessFollowBusiness entity.
type BusinessFollowBusinessCreate struct {
	config
	mutation *BusinessFollowBusinessMutation
	hooks    []Hook
}

// SetCreatedAt sets the "CreatedAt" field.
func (bfbc *BusinessFollowBusinessCreate) SetCreatedAt(t time.Time) *BusinessFollowBusinessCreate {
	bfbc.mutation.SetCreatedAt(t)
	return bfbc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (bfbc *BusinessFollowBusinessCreate) SetNillableCreatedAt(t *time.Time) *BusinessFollowBusinessCreate {
	if t != nil {
		bfbc.SetCreatedAt(*t)
	}
	return bfbc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (bfbc *BusinessFollowBusinessCreate) SetUpdatedAt(t time.Time) *BusinessFollowBusinessCreate {
	bfbc.mutation.SetUpdatedAt(t)
	return bfbc
}

// SetID sets the "id" field.
func (bfbc *BusinessFollowBusinessCreate) SetID(s string) *BusinessFollowBusinessCreate {
	bfbc.mutation.SetID(s)
	return bfbc
}

// SetFollowerID sets the "follower" edge to the Business entity by ID.
func (bfbc *BusinessFollowBusinessCreate) SetFollowerID(id string) *BusinessFollowBusinessCreate {
	bfbc.mutation.SetFollowerID(id)
	return bfbc
}

// SetNillableFollowerID sets the "follower" edge to the Business entity by ID if the given value is not nil.
func (bfbc *BusinessFollowBusinessCreate) SetNillableFollowerID(id *string) *BusinessFollowBusinessCreate {
	if id != nil {
		bfbc = bfbc.SetFollowerID(*id)
	}
	return bfbc
}

// SetFollower sets the "follower" edge to the Business entity.
func (bfbc *BusinessFollowBusinessCreate) SetFollower(b *Business) *BusinessFollowBusinessCreate {
	return bfbc.SetFollowerID(b.ID)
}

// SetFollowedID sets the "followed" edge to the Business entity by ID.
func (bfbc *BusinessFollowBusinessCreate) SetFollowedID(id string) *BusinessFollowBusinessCreate {
	bfbc.mutation.SetFollowedID(id)
	return bfbc
}

// SetNillableFollowedID sets the "followed" edge to the Business entity by ID if the given value is not nil.
func (bfbc *BusinessFollowBusinessCreate) SetNillableFollowedID(id *string) *BusinessFollowBusinessCreate {
	if id != nil {
		bfbc = bfbc.SetFollowedID(*id)
	}
	return bfbc
}

// SetFollowed sets the "followed" edge to the Business entity.
func (bfbc *BusinessFollowBusinessCreate) SetFollowed(b *Business) *BusinessFollowBusinessCreate {
	return bfbc.SetFollowedID(b.ID)
}

// Mutation returns the BusinessFollowBusinessMutation object of the builder.
func (bfbc *BusinessFollowBusinessCreate) Mutation() *BusinessFollowBusinessMutation {
	return bfbc.mutation
}

// Save creates the BusinessFollowBusiness in the database.
func (bfbc *BusinessFollowBusinessCreate) Save(ctx context.Context) (*BusinessFollowBusiness, error) {
	bfbc.defaults()
	return withHooks(ctx, bfbc.sqlSave, bfbc.mutation, bfbc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bfbc *BusinessFollowBusinessCreate) SaveX(ctx context.Context) *BusinessFollowBusiness {
	v, err := bfbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bfbc *BusinessFollowBusinessCreate) Exec(ctx context.Context) error {
	_, err := bfbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bfbc *BusinessFollowBusinessCreate) ExecX(ctx context.Context) {
	if err := bfbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bfbc *BusinessFollowBusinessCreate) defaults() {
	if _, ok := bfbc.mutation.CreatedAt(); !ok {
		v := businessfollowbusiness.DefaultCreatedAt()
		bfbc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bfbc *BusinessFollowBusinessCreate) check() error {
	if _, ok := bfbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`placio_api: missing required field "BusinessFollowBusiness.CreatedAt"`)}
	}
	if _, ok := bfbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`placio_api: missing required field "BusinessFollowBusiness.UpdatedAt"`)}
	}
	return nil
}

func (bfbc *BusinessFollowBusinessCreate) sqlSave(ctx context.Context) (*BusinessFollowBusiness, error) {
	if err := bfbc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bfbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bfbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected BusinessFollowBusiness.ID type: %T", _spec.ID.Value)
		}
	}
	bfbc.mutation.id = &_node.ID
	bfbc.mutation.done = true
	return _node, nil
}

func (bfbc *BusinessFollowBusinessCreate) createSpec() (*BusinessFollowBusiness, *sqlgraph.CreateSpec) {
	var (
		_node = &BusinessFollowBusiness{config: bfbc.config}
		_spec = sqlgraph.NewCreateSpec(businessfollowbusiness.Table, sqlgraph.NewFieldSpec(businessfollowbusiness.FieldID, field.TypeString))
	)
	if id, ok := bfbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bfbc.mutation.CreatedAt(); ok {
		_spec.SetField(businessfollowbusiness.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bfbc.mutation.UpdatedAt(); ok {
		_spec.SetField(businessfollowbusiness.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := bfbc.mutation.FollowerIDs(); len(nodes) > 0 {
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
		_node.business_followed_businesses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bfbc.mutation.FollowedIDs(); len(nodes) > 0 {
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
		_node.business_follower_businesses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BusinessFollowBusinessCreateBulk is the builder for creating many BusinessFollowBusiness entities in bulk.
type BusinessFollowBusinessCreateBulk struct {
	config
	err      error
	builders []*BusinessFollowBusinessCreate
}

// Save creates the BusinessFollowBusiness entities in the database.
func (bfbcb *BusinessFollowBusinessCreateBulk) Save(ctx context.Context) ([]*BusinessFollowBusiness, error) {
	if bfbcb.err != nil {
		return nil, bfbcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bfbcb.builders))
	nodes := make([]*BusinessFollowBusiness, len(bfbcb.builders))
	mutators := make([]Mutator, len(bfbcb.builders))
	for i := range bfbcb.builders {
		func(i int, root context.Context) {
			builder := bfbcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BusinessFollowBusinessMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bfbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bfbcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bfbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bfbcb *BusinessFollowBusinessCreateBulk) SaveX(ctx context.Context) []*BusinessFollowBusiness {
	v, err := bfbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bfbcb *BusinessFollowBusinessCreateBulk) Exec(ctx context.Context) error {
	_, err := bfbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bfbcb *BusinessFollowBusinessCreateBulk) ExecX(ctx context.Context) {
	if err := bfbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
