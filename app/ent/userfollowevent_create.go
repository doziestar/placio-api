




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
			 "placio-app/ent/userfollowevent"
			 "placio-app/ent/user"
			 "placio-app/ent/event"

)







// UserFollowEventCreate is the builder for creating a UserFollowEvent entity.
type UserFollowEventCreate struct {
	config
	mutation *UserFollowEventMutation
	hooks []Hook
}


	





	


	
	
	// SetCreatedAt sets the "createdAt" field.
	func (ufec *UserFollowEventCreate) SetCreatedAt(t time.Time) *UserFollowEventCreate {
		ufec.mutation.SetCreatedAt(t)
		return ufec
	}

	
	
	
	
	
	
		// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
		func (ufec *UserFollowEventCreate) SetNillableCreatedAt(t *time.Time) *UserFollowEventCreate {
			if t != nil {
				ufec.SetCreatedAt(*t)
			}
			return ufec
		}
	

	

	

	

	
	
	// SetUpdatedAt sets the "updatedAt" field.
	func (ufec *UserFollowEventCreate) SetUpdatedAt(t time.Time) *UserFollowEventCreate {
		ufec.mutation.SetUpdatedAt(t)
		return ufec
	}

	
	
	
	
	
	
		// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
		func (ufec *UserFollowEventCreate) SetNillableUpdatedAt(t *time.Time) *UserFollowEventCreate {
			if t != nil {
				ufec.SetUpdatedAt(*t)
			}
			return ufec
		}
	

	

	

	

	
	
	// SetID sets the "id" field.
	func (ufec *UserFollowEventCreate) SetID(s string) *UserFollowEventCreate {
		ufec.mutation.SetID(s)
		return ufec
	}

	
	
	
	
	
	

	

	

	



	
	
	
	
	
		// SetUserID sets the "user" edge to the User entity by ID.
		func (ufec *UserFollowEventCreate) SetUserID(id string) *UserFollowEventCreate {
			ufec.mutation.SetUserID(id)
			return ufec
		}
	
	
	
	
	
	// SetUser sets the "user" edge to the User entity.
	func (ufec *UserFollowEventCreate) SetUser(u *User) *UserFollowEventCreate {
		return ufec.SetUserID(u.ID)
	}

	
	
	
	
	
		// SetEventID sets the "event" edge to the Event entity by ID.
		func (ufec *UserFollowEventCreate) SetEventID(id string) *UserFollowEventCreate {
			ufec.mutation.SetEventID(id)
			return ufec
		}
	
	
	
	
	
	// SetEvent sets the "event" edge to the Event entity.
	func (ufec *UserFollowEventCreate) SetEvent(e *Event) *UserFollowEventCreate {
		return ufec.SetEventID(e.ID)
	}


// Mutation returns the UserFollowEventMutation object of the builder.
func (ufec *UserFollowEventCreate) Mutation() *UserFollowEventMutation {
	return ufec.mutation
}




// Save creates the UserFollowEvent in the database.
func (ufec *UserFollowEventCreate) Save(ctx context.Context) (*UserFollowEvent, error) {
			ufec.defaults()
	return withHooks(ctx, ufec.sqlSave, ufec.mutation, ufec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ufec *UserFollowEventCreate) SaveX(ctx context.Context) *UserFollowEvent {
	v, err := ufec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufec *UserFollowEventCreate) Exec(ctx context.Context) error {
	_, err := ufec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufec *UserFollowEventCreate) ExecX(ctx context.Context) {
	if err := ufec.Exec(ctx); err != nil {
		panic(err)
	}
}

	// defaults sets the default values of the builder before save.
	func (ufec *UserFollowEventCreate) defaults() {
				if _, ok := ufec.mutation.CreatedAt(); !ok {
					v := userfollowevent.DefaultCreatedAt()
					ufec.mutation.SetCreatedAt(v)
				}
				if _, ok := ufec.mutation.UpdatedAt(); !ok {
					v := userfollowevent.DefaultUpdatedAt()
					ufec.mutation.SetUpdatedAt(v)
				}
	}


// check runs all checks and user-defined validators on the builder.
func (ufec *UserFollowEventCreate) check() error {
					if _, ok := ufec.mutation.CreatedAt(); !ok {
						return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "UserFollowEvent.createdAt"`)}
					}
					if _, ok := ufec.mutation.UpdatedAt(); !ok {
						return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "UserFollowEvent.updatedAt"`)}
					}
				if _, ok := ufec.mutation.UserID(); !ok {
				return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserFollowEvent.user"`)}
			}
				if _, ok := ufec.mutation.EventID(); !ok {
				return &ValidationError{Name: "event", err: errors.New(`ent: missing required edge "UserFollowEvent.event"`)}
			}
	return nil
}


	
	




func (ufec *UserFollowEventCreate) sqlSave(ctx context.Context) (*UserFollowEvent, error) {
	if err := ufec.check(); err != nil {
		return nil, err
	}
	_node, _spec  := ufec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ufec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
		if _spec.ID.Value != nil {
				if id, ok := _spec.ID.Value.(string); ok {
					_node.ID = id
				} else {
					return nil, fmt.Errorf("unexpected UserFollowEvent.ID type: %T", _spec.ID.Value)
				}
		}
		ufec.mutation.id = &_node.ID
		ufec.mutation.done = true
	return _node, nil
}

func (ufec *UserFollowEventCreate) createSpec() (*UserFollowEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &UserFollowEvent{config: ufec.config}
		_spec = sqlgraph.NewCreateSpec(userfollowevent.Table, sqlgraph.NewFieldSpec(userfollowevent.FieldID, field.TypeString))
	)
		if id, ok := ufec.mutation.ID(); ok {
			_node.ID = id
			_spec.ID.Value = id
		}
		if value, ok := ufec.mutation.CreatedAt(); ok {
				_spec.SetField(userfollowevent.FieldCreatedAt, field.TypeTime, value)
			_node.CreatedAt = value
		}
		if value, ok := ufec.mutation.UpdatedAt(); ok {
				_spec.SetField(userfollowevent.FieldUpdatedAt, field.TypeTime, value)
			_node.UpdatedAt = value
		}
		if nodes := ufec.mutation.UserIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: userfollowevent.UserTable,
		Columns: []string{ userfollowevent.UserColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
				_node.user_user_follow_events = &nodes[0]
			_spec.Edges = append(_spec.Edges, edge)
		}
		if nodes := ufec.mutation.EventIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: false,
		Table: userfollowevent.EventTable,
		Columns: []string{ userfollowevent.EventColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
				_node.user_follow_event_event = &nodes[0]
			_spec.Edges = append(_spec.Edges, edge)
		}
	return _node, _spec
}
	








// UserFollowEventCreateBulk is the builder for creating many UserFollowEvent entities in bulk.
type UserFollowEventCreateBulk struct {
	config
	err error
	builders []*UserFollowEventCreate
}




	
		



// Save creates the UserFollowEvent entities in the database.
func (ufecb *UserFollowEventCreateBulk) Save(ctx context.Context) ([]*UserFollowEvent, error) {
	if ufecb.err != nil {
		return nil, ufecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ufecb.builders))
	nodes := make([]*UserFollowEvent, len(ufecb.builders))
	mutators := make([]Mutator, len(ufecb.builders))
	for i := range ufecb.builders {
		func(i int, root context.Context) {
			builder := ufecb.builders[i]
				builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserFollowEventMutation)
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
					_, err = mutators[i+1].Mutate(root, ufecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ufecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ufecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ufecb *UserFollowEventCreateBulk) SaveX(ctx context.Context) []*UserFollowEvent {
	v, err := ufecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufecb *UserFollowEventCreateBulk) Exec(ctx context.Context) error {
	_, err := ufecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufecb *UserFollowEventCreateBulk) ExecX(ctx context.Context) {
	if err := ufecb.Exec(ctx); err != nil {
		panic(err)
	}
}
	


	

