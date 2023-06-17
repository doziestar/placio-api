// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/event"
	"placio-app/ent/ticket"
	"placio-app/ent/ticketoption"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TicketCreate is the builder for creating a Ticket entity.
type TicketCreate struct {
	config
	mutation *TicketMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (tc *TicketCreate) SetCreatedAt(t time.Time) *TicketCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (tc *TicketCreate) SetNillableCreatedAt(t *time.Time) *TicketCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updatedAt" field.
func (tc *TicketCreate) SetUpdatedAt(t time.Time) *TicketCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (tc *TicketCreate) SetNillableUpdatedAt(t *time.Time) *TicketCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TicketCreate) SetID(s string) *TicketCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (tc *TicketCreate) SetEventID(id string) *TicketCreate {
	tc.mutation.SetEventID(id)
	return tc
}

// SetNillableEventID sets the "event" edge to the Event entity by ID if the given value is not nil.
func (tc *TicketCreate) SetNillableEventID(id *string) *TicketCreate {
	if id != nil {
		tc = tc.SetEventID(*id)
	}
	return tc
}

// SetEvent sets the "event" edge to the Event entity.
func (tc *TicketCreate) SetEvent(e *Event) *TicketCreate {
	return tc.SetEventID(e.ID)
}

// AddTicketOptionIDs adds the "ticket_options" edge to the TicketOption entity by IDs.
func (tc *TicketCreate) AddTicketOptionIDs(ids ...string) *TicketCreate {
	tc.mutation.AddTicketOptionIDs(ids...)
	return tc
}

// AddTicketOptions adds the "ticket_options" edges to the TicketOption entity.
func (tc *TicketCreate) AddTicketOptions(t ...*TicketOption) *TicketCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTicketOptionIDs(ids...)
}

// Mutation returns the TicketMutation object of the builder.
func (tc *TicketCreate) Mutation() *TicketMutation {
	return tc.mutation
}

// Save creates the Ticket in the database.
func (tc *TicketCreate) Save(ctx context.Context) (*Ticket, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TicketCreate) SaveX(ctx context.Context) *Ticket {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TicketCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TicketCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TicketCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := ticket.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := ticket.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TicketCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Ticket.createdAt"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Ticket.updatedAt"`)}
	}
	return nil
}

func (tc *TicketCreate) sqlSave(ctx context.Context) (*Ticket, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Ticket.ID type: %T", _spec.ID.Value)
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TicketCreate) createSpec() (*Ticket, *sqlgraph.CreateSpec) {
	var (
		_node = &Ticket{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(ticket.Table, sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeString))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(ticket.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(ticket.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.EventTable,
			Columns: []string{ticket.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.event_tickets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TicketOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   ticket.TicketOptionsTable,
			Columns: []string{ticket.TicketOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticketoption.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TicketCreateBulk is the builder for creating many Ticket entities in bulk.
type TicketCreateBulk struct {
	config
	builders []*TicketCreate
}

// Save creates the Ticket entities in the database.
func (tcb *TicketCreateBulk) Save(ctx context.Context) ([]*Ticket, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Ticket, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TicketMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TicketCreateBulk) SaveX(ctx context.Context) []*Ticket {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TicketCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TicketCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
