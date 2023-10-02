// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/event"
	"placio-app/ent/ticketoption"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TicketOptionCreate is the builder for creating a TicketOption entity.
type TicketOptionCreate struct {
	config
	mutation *TicketOptionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (toc *TicketOptionCreate) SetCreatedAt(t time.Time) *TicketOptionCreate {
	toc.mutation.SetCreatedAt(t)
	return toc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (toc *TicketOptionCreate) SetNillableCreatedAt(t *time.Time) *TicketOptionCreate {
	if t != nil {
		toc.SetCreatedAt(*t)
	}
	return toc
}

// SetUpdatedAt sets the "updatedAt" field.
func (toc *TicketOptionCreate) SetUpdatedAt(t time.Time) *TicketOptionCreate {
	toc.mutation.SetUpdatedAt(t)
	return toc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (toc *TicketOptionCreate) SetNillableUpdatedAt(t *time.Time) *TicketOptionCreate {
	if t != nil {
		toc.SetUpdatedAt(*t)
	}
	return toc
}

// SetID sets the "id" field.
func (toc *TicketOptionCreate) SetID(s string) *TicketOptionCreate {
	toc.mutation.SetID(s)
	return toc
}

// SetEventID sets the "event" edge to the Event entity by ID.
func (toc *TicketOptionCreate) SetEventID(id string) *TicketOptionCreate {
	toc.mutation.SetEventID(id)
	return toc
}

// SetNillableEventID sets the "event" edge to the Event entity by ID if the given value is not nil.
func (toc *TicketOptionCreate) SetNillableEventID(id *string) *TicketOptionCreate {
	if id != nil {
		toc = toc.SetEventID(*id)
	}
	return toc
}

// SetEvent sets the "event" edge to the Event entity.
func (toc *TicketOptionCreate) SetEvent(e *Event) *TicketOptionCreate {
	return toc.SetEventID(e.ID)
}

// Mutation returns the TicketOptionMutation object of the builder.
func (toc *TicketOptionCreate) Mutation() *TicketOptionMutation {
	return toc.mutation
}

// Save creates the TicketOption in the database.
func (toc *TicketOptionCreate) Save(ctx context.Context) (*TicketOption, error) {
	toc.defaults()
	return withHooks(ctx, toc.sqlSave, toc.mutation, toc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (toc *TicketOptionCreate) SaveX(ctx context.Context) *TicketOption {
	v, err := toc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (toc *TicketOptionCreate) Exec(ctx context.Context) error {
	_, err := toc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (toc *TicketOptionCreate) ExecX(ctx context.Context) {
	if err := toc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (toc *TicketOptionCreate) defaults() {
	if _, ok := toc.mutation.CreatedAt(); !ok {
		v := ticketoption.DefaultCreatedAt()
		toc.mutation.SetCreatedAt(v)
	}
	if _, ok := toc.mutation.UpdatedAt(); !ok {
		v := ticketoption.DefaultUpdatedAt()
		toc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (toc *TicketOptionCreate) check() error {
	if _, ok := toc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "TicketOption.createdAt"`)}
	}
	if _, ok := toc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "TicketOption.updatedAt"`)}
	}
	return nil
}

func (toc *TicketOptionCreate) sqlSave(ctx context.Context) (*TicketOption, error) {
	if err := toc.check(); err != nil {
		return nil, err
	}
	_node, _spec := toc.createSpec()
	if err := sqlgraph.CreateNode(ctx, toc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected TicketOption.ID type: %T", _spec.ID.Value)
		}
	}
	toc.mutation.id = &_node.ID
	toc.mutation.done = true
	return _node, nil
}

func (toc *TicketOptionCreate) createSpec() (*TicketOption, *sqlgraph.CreateSpec) {
	var (
		_node = &TicketOption{config: toc.config}
		_spec = sqlgraph.NewCreateSpec(ticketoption.Table, sqlgraph.NewFieldSpec(ticketoption.FieldID, field.TypeString))
	)
	if id, ok := toc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := toc.mutation.CreatedAt(); ok {
		_spec.SetField(ticketoption.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := toc.mutation.UpdatedAt(); ok {
		_spec.SetField(ticketoption.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := toc.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticketoption.EventTable,
			Columns: []string{ticketoption.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.event_ticket_options = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TicketOptionCreateBulk is the builder for creating many TicketOption entities in bulk.
type TicketOptionCreateBulk struct {
	config
	err      error
	builders []*TicketOptionCreate
}

// Save creates the TicketOption entities in the database.
func (tocb *TicketOptionCreateBulk) Save(ctx context.Context) ([]*TicketOption, error) {
	if tocb.err != nil {
		return nil, tocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tocb.builders))
	nodes := make([]*TicketOption, len(tocb.builders))
	mutators := make([]Mutator, len(tocb.builders))
	for i := range tocb.builders {
		func(i int, root context.Context) {
			builder := tocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TicketOptionMutation)
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
					_, err = mutators[i+1].Mutate(root, tocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tocb *TicketOptionCreateBulk) SaveX(ctx context.Context) []*TicketOption {
	v, err := tocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tocb *TicketOptionCreateBulk) Exec(ctx context.Context) error {
	_, err := tocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tocb *TicketOptionCreateBulk) ExecX(ctx context.Context) {
	if err := tocb.Exec(ctx); err != nil {
		panic(err)
	}
}