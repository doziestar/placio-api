// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/permission"
	"placio-app/ent/place"
	"placio-app/ent/staff"
	"placio-app/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StaffCreate is the builder for creating a Staff entity.
type StaffCreate struct {
	config
	mutation *StaffMutation
	hooks    []Hook
}

// SetPosition sets the "position" field.
func (sc *StaffCreate) SetPosition(s string) *StaffCreate {
	sc.mutation.SetPosition(s)
	return sc
}

// SetNillablePosition sets the "position" field if the given value is not nil.
func (sc *StaffCreate) SetNillablePosition(s *string) *StaffCreate {
	if s != nil {
		sc.SetPosition(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StaffCreate) SetID(s string) *StaffCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sc *StaffCreate) SetUserID(id string) *StaffCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (sc *StaffCreate) SetNillableUserID(id *string) *StaffCreate {
	if id != nil {
		sc = sc.SetUserID(*id)
	}
	return sc
}

// SetUser sets the "user" edge to the User entity.
func (sc *StaffCreate) SetUser(u *User) *StaffCreate {
	return sc.SetUserID(u.ID)
}

// AddPlaceIDs adds the "place" edge to the Place entity by IDs.
func (sc *StaffCreate) AddPlaceIDs(ids ...string) *StaffCreate {
	sc.mutation.AddPlaceIDs(ids...)
	return sc
}

// AddPlace adds the "place" edges to the Place entity.
func (sc *StaffCreate) AddPlace(p ...*Place) *StaffCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPlaceIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (sc *StaffCreate) AddPermissionIDs(ids ...string) *StaffCreate {
	sc.mutation.AddPermissionIDs(ids...)
	return sc
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (sc *StaffCreate) AddPermissions(p ...*Permission) *StaffCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPermissionIDs(ids...)
}

// AddBusinesIDs adds the "business" edge to the Business entity by IDs.
func (sc *StaffCreate) AddBusinesIDs(ids ...string) *StaffCreate {
	sc.mutation.AddBusinesIDs(ids...)
	return sc
}

// AddBusiness adds the "business" edges to the Business entity.
func (sc *StaffCreate) AddBusiness(b ...*Business) *StaffCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return sc.AddBusinesIDs(ids...)
}

// Mutation returns the StaffMutation object of the builder.
func (sc *StaffCreate) Mutation() *StaffMutation {
	return sc.mutation
}

// Save creates the Staff in the database.
func (sc *StaffCreate) Save(ctx context.Context) (*Staff, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StaffCreate) SaveX(ctx context.Context) *Staff {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StaffCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StaffCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StaffCreate) check() error {
	if v, ok := sc.mutation.ID(); ok {
		if err := staff.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Staff.id": %w`, err)}
		}
	}
	return nil
}

func (sc *StaffCreate) sqlSave(ctx context.Context) (*Staff, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Staff.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StaffCreate) createSpec() (*Staff, *sqlgraph.CreateSpec) {
	var (
		_node = &Staff{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(staff.Table, sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Position(); ok {
		_spec.SetField(staff.FieldPosition, field.TypeString, value)
		_node.Position = value
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   staff.UserTable,
			Columns: []string{staff.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_staffs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   staff.PlaceTable,
			Columns: staff.PlacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   staff.PermissionsTable,
			Columns: staff.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   staff.BusinessTable,
			Columns: staff.BusinessPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StaffCreateBulk is the builder for creating many Staff entities in bulk.
type StaffCreateBulk struct {
	config
	err      error
	builders []*StaffCreate
}

// Save creates the Staff entities in the database.
func (scb *StaffCreateBulk) Save(ctx context.Context) ([]*Staff, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Staff, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StaffMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StaffCreateBulk) SaveX(ctx context.Context) []*Staff {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StaffCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StaffCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}