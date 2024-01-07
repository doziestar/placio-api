// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/media"
	"placio-app/ent/place"
	"placio-app/ent/plan"
	"placio-app/ent/price"
	"placio-app/ent/subscription"
	"placio-app/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlanCreate is the builder for creating a Plan entity.
type PlanCreate struct {
	config
	mutation *PlanMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PlanCreate) SetName(s string) *PlanCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *PlanCreate) SetNillableName(s *string) *PlanCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *PlanCreate) SetDescription(s string) *PlanCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PlanCreate) SetNillableDescription(s *string) *PlanCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetOverview sets the "overview" field.
func (pc *PlanCreate) SetOverview(s string) *PlanCreate {
	pc.mutation.SetOverview(s)
	return pc
}

// SetNillableOverview sets the "overview" field if the given value is not nil.
func (pc *PlanCreate) SetNillableOverview(s *string) *PlanCreate {
	if s != nil {
		pc.SetOverview(*s)
	}
	return pc
}

// SetFeatures sets the "features" field.
func (pc *PlanCreate) SetFeatures(s []string) *PlanCreate {
	pc.mutation.SetFeatures(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PlanCreate) SetID(s string) *PlanCreate {
	pc.mutation.SetID(s)
	return pc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (pc *PlanCreate) AddUserIDs(ids ...string) *PlanCreate {
	pc.mutation.AddUserIDs(ids...)
	return pc
}

// AddUsers adds the "users" edges to the User entity.
func (pc *PlanCreate) AddUsers(u ...*User) *PlanCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddUserIDs(ids...)
}

// AddBusinessIDs adds the "businesses" edge to the Business entity by IDs.
func (pc *PlanCreate) AddBusinessIDs(ids ...string) *PlanCreate {
	pc.mutation.AddBusinessIDs(ids...)
	return pc
}

// AddBusinesses adds the "businesses" edges to the Business entity.
func (pc *PlanCreate) AddBusinesses(b ...*Business) *PlanCreate {
	ids := make([]string, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddBusinessIDs(ids...)
}

// AddPlaceIDs adds the "places" edge to the Place entity by IDs.
func (pc *PlanCreate) AddPlaceIDs(ids ...string) *PlanCreate {
	pc.mutation.AddPlaceIDs(ids...)
	return pc
}

// AddPlaces adds the "places" edges to the Place entity.
func (pc *PlanCreate) AddPlaces(p ...*Place) *PlanCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPlaceIDs(ids...)
}

// AddMediumIDs adds the "media" edge to the Media entity by IDs.
func (pc *PlanCreate) AddMediumIDs(ids ...string) *PlanCreate {
	pc.mutation.AddMediumIDs(ids...)
	return pc
}

// AddMedia adds the "media" edges to the Media entity.
func (pc *PlanCreate) AddMedia(m ...*Media) *PlanCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pc.AddMediumIDs(ids...)
}

// AddPriceIDs adds the "prices" edge to the Price entity by IDs.
func (pc *PlanCreate) AddPriceIDs(ids ...string) *PlanCreate {
	pc.mutation.AddPriceIDs(ids...)
	return pc
}

// AddPrices adds the "prices" edges to the Price entity.
func (pc *PlanCreate) AddPrices(p ...*Price) *PlanCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPriceIDs(ids...)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the Subscription entity by IDs.
func (pc *PlanCreate) AddSubscriptionIDs(ids ...string) *PlanCreate {
	pc.mutation.AddSubscriptionIDs(ids...)
	return pc
}

// AddSubscriptions adds the "subscriptions" edges to the Subscription entity.
func (pc *PlanCreate) AddSubscriptions(s ...*Subscription) *PlanCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddSubscriptionIDs(ids...)
}

// Mutation returns the PlanMutation object of the builder.
func (pc *PlanCreate) Mutation() *PlanMutation {
	return pc.mutation
}

// Save creates the Plan in the database.
func (pc *PlanCreate) Save(ctx context.Context) (*Plan, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlanCreate) SaveX(ctx context.Context) *Plan {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlanCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlanCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlanCreate) check() error {
	if v, ok := pc.mutation.ID(); ok {
		if err := plan.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Plan.id": %w`, err)}
		}
	}
	return nil
}

func (pc *PlanCreate) sqlSave(ctx context.Context) (*Plan, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Plan.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PlanCreate) createSpec() (*Plan, *sqlgraph.CreateSpec) {
	var (
		_node = &Plan{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(plan.Table, sqlgraph.NewFieldSpec(plan.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(plan.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(plan.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Overview(); ok {
		_spec.SetField(plan.FieldOverview, field.TypeString, value)
		_node.Overview = value
	}
	if value, ok := pc.mutation.Features(); ok {
		_spec.SetField(plan.FieldFeatures, field.TypeJSON, value)
		_node.Features = value
	}
	if nodes := pc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.UsersTable,
			Columns: plan.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BusinessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.BusinessesTable,
			Columns: plan.BusinessesPrimaryKey,
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
	if nodes := pc.mutation.PlacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   plan.PlacesTable,
			Columns: plan.PlacesPrimaryKey,
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
	if nodes := pc.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   plan.MediaTable,
			Columns: []string{plan.MediaColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(media.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PricesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   plan.PricesTable,
			Columns: []string{plan.PricesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(price.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   plan.SubscriptionsTable,
			Columns: []string{plan.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscription.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlanCreateBulk is the builder for creating many Plan entities in bulk.
type PlanCreateBulk struct {
	config
	err      error
	builders []*PlanCreate
}

// Save creates the Plan entities in the database.
func (pcb *PlanCreateBulk) Save(ctx context.Context) ([]*Plan, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Plan, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlanMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlanCreateBulk) SaveX(ctx context.Context) []*Plan {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlanCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlanCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
