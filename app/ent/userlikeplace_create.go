// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/place"
	"placio-app/ent/user"
	"placio-app/ent/userlikeplace"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserLikePlaceCreate is the builder for creating a UserLikePlace entity.
type UserLikePlaceCreate struct {
	config
	mutation *UserLikePlaceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "CreatedAt" field.
func (ulpc *UserLikePlaceCreate) SetCreatedAt(t time.Time) *UserLikePlaceCreate {
	ulpc.mutation.SetCreatedAt(t)
	return ulpc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (ulpc *UserLikePlaceCreate) SetNillableCreatedAt(t *time.Time) *UserLikePlaceCreate {
	if t != nil {
		ulpc.SetCreatedAt(*t)
	}
	return ulpc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (ulpc *UserLikePlaceCreate) SetUpdatedAt(t time.Time) *UserLikePlaceCreate {
	ulpc.mutation.SetUpdatedAt(t)
	return ulpc
}

// SetID sets the "id" field.
func (ulpc *UserLikePlaceCreate) SetID(s string) *UserLikePlaceCreate {
	ulpc.mutation.SetID(s)
	return ulpc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ulpc *UserLikePlaceCreate) SetUserID(id string) *UserLikePlaceCreate {
	ulpc.mutation.SetUserID(id)
	return ulpc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ulpc *UserLikePlaceCreate) SetNillableUserID(id *string) *UserLikePlaceCreate {
	if id != nil {
		ulpc = ulpc.SetUserID(*id)
	}
	return ulpc
}

// SetUser sets the "user" edge to the User entity.
func (ulpc *UserLikePlaceCreate) SetUser(u *User) *UserLikePlaceCreate {
	return ulpc.SetUserID(u.ID)
}

// SetPlaceID sets the "place" edge to the Place entity by ID.
func (ulpc *UserLikePlaceCreate) SetPlaceID(id string) *UserLikePlaceCreate {
	ulpc.mutation.SetPlaceID(id)
	return ulpc
}

// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
func (ulpc *UserLikePlaceCreate) SetNillablePlaceID(id *string) *UserLikePlaceCreate {
	if id != nil {
		ulpc = ulpc.SetPlaceID(*id)
	}
	return ulpc
}

// SetPlace sets the "place" edge to the Place entity.
func (ulpc *UserLikePlaceCreate) SetPlace(p *Place) *UserLikePlaceCreate {
	return ulpc.SetPlaceID(p.ID)
}

// Mutation returns the UserLikePlaceMutation object of the builder.
func (ulpc *UserLikePlaceCreate) Mutation() *UserLikePlaceMutation {
	return ulpc.mutation
}

// Save creates the UserLikePlace in the database.
func (ulpc *UserLikePlaceCreate) Save(ctx context.Context) (*UserLikePlace, error) {
	ulpc.defaults()
	return withHooks(ctx, ulpc.sqlSave, ulpc.mutation, ulpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ulpc *UserLikePlaceCreate) SaveX(ctx context.Context) *UserLikePlace {
	v, err := ulpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ulpc *UserLikePlaceCreate) Exec(ctx context.Context) error {
	_, err := ulpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulpc *UserLikePlaceCreate) ExecX(ctx context.Context) {
	if err := ulpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ulpc *UserLikePlaceCreate) defaults() {
	if _, ok := ulpc.mutation.CreatedAt(); !ok {
		v := userlikeplace.DefaultCreatedAt()
		ulpc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ulpc *UserLikePlaceCreate) check() error {
	if _, ok := ulpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "UserLikePlace.CreatedAt"`)}
	}
	if _, ok := ulpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "UserLikePlace.UpdatedAt"`)}
	}
	return nil
}

func (ulpc *UserLikePlaceCreate) sqlSave(ctx context.Context) (*UserLikePlace, error) {
	if err := ulpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ulpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ulpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserLikePlace.ID type: %T", _spec.ID.Value)
		}
	}
	ulpc.mutation.id = &_node.ID
	ulpc.mutation.done = true
	return _node, nil
}

func (ulpc *UserLikePlaceCreate) createSpec() (*UserLikePlace, *sqlgraph.CreateSpec) {
	var (
		_node = &UserLikePlace{config: ulpc.config}
		_spec = sqlgraph.NewCreateSpec(userlikeplace.Table, sqlgraph.NewFieldSpec(userlikeplace.FieldID, field.TypeString))
	)
	if id, ok := ulpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ulpc.mutation.CreatedAt(); ok {
		_spec.SetField(userlikeplace.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ulpc.mutation.UpdatedAt(); ok {
		_spec.SetField(userlikeplace.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ulpc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userlikeplace.UserTable,
			Columns: []string{userlikeplace.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_liked_places = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ulpc.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userlikeplace.PlaceTable,
			Columns: []string{userlikeplace.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_like_place_place = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserLikePlaceCreateBulk is the builder for creating many UserLikePlace entities in bulk.
type UserLikePlaceCreateBulk struct {
	config
	err      error
	builders []*UserLikePlaceCreate
}

// Save creates the UserLikePlace entities in the database.
func (ulpcb *UserLikePlaceCreateBulk) Save(ctx context.Context) ([]*UserLikePlace, error) {
	if ulpcb.err != nil {
		return nil, ulpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ulpcb.builders))
	nodes := make([]*UserLikePlace, len(ulpcb.builders))
	mutators := make([]Mutator, len(ulpcb.builders))
	for i := range ulpcb.builders {
		func(i int, root context.Context) {
			builder := ulpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserLikePlaceMutation)
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
					_, err = mutators[i+1].Mutate(root, ulpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ulpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ulpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ulpcb *UserLikePlaceCreateBulk) SaveX(ctx context.Context) []*UserLikePlace {
	v, err := ulpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ulpcb *UserLikePlaceCreateBulk) Exec(ctx context.Context) error {
	_, err := ulpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulpcb *UserLikePlaceCreateBulk) ExecX(ctx context.Context) {
	if err := ulpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
