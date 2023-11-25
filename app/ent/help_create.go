// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/help"
	"placio-app/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HelpCreate is the builder for creating a Help entity.
type HelpCreate struct {
	config
	mutation *HelpMutation
	hooks    []Hook
}

// SetCategory sets the "category" field.
func (hc *HelpCreate) SetCategory(s string) *HelpCreate {
	hc.mutation.SetCategory(s)
	return hc
}

// SetSubject sets the "subject" field.
func (hc *HelpCreate) SetSubject(s string) *HelpCreate {
	hc.mutation.SetSubject(s)
	return hc
}

// SetBody sets the "body" field.
func (hc *HelpCreate) SetBody(s string) *HelpCreate {
	hc.mutation.SetBody(s)
	return hc
}

// SetMedia sets the "media" field.
func (hc *HelpCreate) SetMedia(s string) *HelpCreate {
	hc.mutation.SetMedia(s)
	return hc
}

// SetNillableMedia sets the "media" field if the given value is not nil.
func (hc *HelpCreate) SetNillableMedia(s *string) *HelpCreate {
	if s != nil {
		hc.SetMedia(*s)
	}
	return hc
}

// SetStatus sets the "status" field.
func (hc *HelpCreate) SetStatus(s string) *HelpCreate {
	hc.mutation.SetStatus(s)
	return hc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hc *HelpCreate) SetNillableStatus(s *string) *HelpCreate {
	if s != nil {
		hc.SetStatus(*s)
	}
	return hc
}

// SetUserID sets the "user_id" field.
func (hc *HelpCreate) SetUserID(s string) *HelpCreate {
	hc.mutation.SetUserID(s)
	return hc
}

// SetID sets the "id" field.
func (hc *HelpCreate) SetID(s string) *HelpCreate {
	hc.mutation.SetID(s)
	return hc
}

// SetUser sets the "user" edge to the User entity.
func (hc *HelpCreate) SetUser(u *User) *HelpCreate {
	return hc.SetUserID(u.ID)
}

// Mutation returns the HelpMutation object of the builder.
func (hc *HelpCreate) Mutation() *HelpMutation {
	return hc.mutation
}

// Save creates the Help in the database.
func (hc *HelpCreate) Save(ctx context.Context) (*Help, error) {
	hc.defaults()
	return withHooks(ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HelpCreate) SaveX(ctx context.Context) *Help {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HelpCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HelpCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HelpCreate) defaults() {
	if _, ok := hc.mutation.Status(); !ok {
		v := help.DefaultStatus
		hc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HelpCreate) check() error {
	if _, ok := hc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Help.category"`)}
	}
	if _, ok := hc.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "Help.subject"`)}
	}
	if _, ok := hc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Help.body"`)}
	}
	if _, ok := hc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Help.status"`)}
	}
	if _, ok := hc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Help.user_id"`)}
	}
	if v, ok := hc.mutation.UserID(); ok {
		if err := help.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Help.user_id": %w`, err)}
		}
	}
	if _, ok := hc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Help.user"`)}
	}
	return nil
}

func (hc *HelpCreate) sqlSave(ctx context.Context) (*Help, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Help.ID type: %T", _spec.ID.Value)
		}
	}
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HelpCreate) createSpec() (*Help, *sqlgraph.CreateSpec) {
	var (
		_node = &Help{config: hc.config}
		_spec = sqlgraph.NewCreateSpec(help.Table, sqlgraph.NewFieldSpec(help.FieldID, field.TypeString))
	)
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hc.mutation.Category(); ok {
		_spec.SetField(help.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := hc.mutation.Subject(); ok {
		_spec.SetField(help.FieldSubject, field.TypeString, value)
		_node.Subject = value
	}
	if value, ok := hc.mutation.Body(); ok {
		_spec.SetField(help.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := hc.mutation.Media(); ok {
		_spec.SetField(help.FieldMedia, field.TypeString, value)
		_node.Media = value
	}
	if value, ok := hc.mutation.Status(); ok {
		_spec.SetField(help.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if nodes := hc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   help.UserTable,
			Columns: []string{help.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HelpCreateBulk is the builder for creating many Help entities in bulk.
type HelpCreateBulk struct {
	config
	err      error
	builders []*HelpCreate
}

// Save creates the Help entities in the database.
func (hcb *HelpCreateBulk) Save(ctx context.Context) ([]*Help, error) {
	if hcb.err != nil {
		return nil, hcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Help, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HelpMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HelpCreateBulk) SaveX(ctx context.Context) []*Help {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HelpCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HelpCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
