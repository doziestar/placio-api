// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"placio-app/ent/place"
	"placio-app/ent/user"
	"placio-app/ent/userfollowplace"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFollowPlaceCreate is the builder for creating a UserFollowPlace entity.
type UserFollowPlaceCreate struct {
	config
	mutation *UserFollowPlaceMutation
	hooks    []Hook
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ufpc *UserFollowPlaceCreate) SetUserID(id string) *UserFollowPlaceCreate {
	ufpc.mutation.SetUserID(id)
	return ufpc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ufpc *UserFollowPlaceCreate) SetNillableUserID(id *string) *UserFollowPlaceCreate {
	if id != nil {
		ufpc = ufpc.SetUserID(*id)
	}
	return ufpc
}

// SetUser sets the "user" edge to the User entity.
func (ufpc *UserFollowPlaceCreate) SetUser(u *User) *UserFollowPlaceCreate {
	return ufpc.SetUserID(u.ID)
}

// SetPlaceID sets the "place" edge to the Place entity by ID.
func (ufpc *UserFollowPlaceCreate) SetPlaceID(id string) *UserFollowPlaceCreate {
	ufpc.mutation.SetPlaceID(id)
	return ufpc
}

// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
func (ufpc *UserFollowPlaceCreate) SetNillablePlaceID(id *string) *UserFollowPlaceCreate {
	if id != nil {
		ufpc = ufpc.SetPlaceID(*id)
	}
	return ufpc
}

// SetPlace sets the "place" edge to the Place entity.
func (ufpc *UserFollowPlaceCreate) SetPlace(p *Place) *UserFollowPlaceCreate {
	return ufpc.SetPlaceID(p.ID)
}

// Mutation returns the UserFollowPlaceMutation object of the builder.
func (ufpc *UserFollowPlaceCreate) Mutation() *UserFollowPlaceMutation {
	return ufpc.mutation
}

// Save creates the UserFollowPlace in the database.
func (ufpc *UserFollowPlaceCreate) Save(ctx context.Context) (*UserFollowPlace, error) {
	return withHooks(ctx, ufpc.sqlSave, ufpc.mutation, ufpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ufpc *UserFollowPlaceCreate) SaveX(ctx context.Context) *UserFollowPlace {
	v, err := ufpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufpc *UserFollowPlaceCreate) Exec(ctx context.Context) error {
	_, err := ufpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufpc *UserFollowPlaceCreate) ExecX(ctx context.Context) {
	if err := ufpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ufpc *UserFollowPlaceCreate) check() error {
	return nil
}

func (ufpc *UserFollowPlaceCreate) sqlSave(ctx context.Context) (*UserFollowPlace, error) {
	if err := ufpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ufpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ufpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserFollowPlace.ID type: %T", _spec.ID.Value)
		}
	}
	ufpc.mutation.id = &_node.ID
	ufpc.mutation.done = true
	return _node, nil
}

func (ufpc *UserFollowPlaceCreate) createSpec() (*UserFollowPlace, *sqlgraph.CreateSpec) {
	var (
		_node = &UserFollowPlace{config: ufpc.config}
		_spec = sqlgraph.NewCreateSpec(userfollowplace.Table, sqlgraph.NewFieldSpec(userfollowplace.FieldID, field.TypeString))
	)
	if nodes := ufpc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowplace.UserTable,
			Columns: []string{userfollowplace.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_followed_places = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ufpc.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userfollowplace.PlaceTable,
			Columns: []string{userfollowplace.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.place_follower_users = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserFollowPlaceCreateBulk is the builder for creating many UserFollowPlace entities in bulk.
type UserFollowPlaceCreateBulk struct {
	config
	builders []*UserFollowPlaceCreate
}

// Save creates the UserFollowPlace entities in the database.
func (ufpcb *UserFollowPlaceCreateBulk) Save(ctx context.Context) ([]*UserFollowPlace, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ufpcb.builders))
	nodes := make([]*UserFollowPlace, len(ufpcb.builders))
	mutators := make([]Mutator, len(ufpcb.builders))
	for i := range ufpcb.builders {
		func(i int, root context.Context) {
			builder := ufpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserFollowPlaceMutation)
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
					_, err = mutators[i+1].Mutate(root, ufpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ufpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ufpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ufpcb *UserFollowPlaceCreateBulk) SaveX(ctx context.Context) []*UserFollowPlace {
	v, err := ufpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufpcb *UserFollowPlaceCreateBulk) Exec(ctx context.Context) error {
	_, err := ufpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufpcb *UserFollowPlaceCreateBulk) ExecX(ctx context.Context) {
	if err := ufpcb.Exec(ctx); err != nil {
		panic(err)
	}
}