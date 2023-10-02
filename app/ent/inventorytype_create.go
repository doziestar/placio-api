// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio_api/inventoryattribute"
	"placio_api/inventorytype"
	"placio_api/placeinventory"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InventoryTypeCreate is the builder for creating a InventoryType entity.
type InventoryTypeCreate struct {
	config
	mutation *InventoryTypeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (itc *InventoryTypeCreate) SetName(s string) *InventoryTypeCreate {
	itc.mutation.SetName(s)
	return itc
}

// SetDescription sets the "description" field.
func (itc *InventoryTypeCreate) SetDescription(s string) *InventoryTypeCreate {
	itc.mutation.SetDescription(s)
	return itc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (itc *InventoryTypeCreate) SetNillableDescription(s *string) *InventoryTypeCreate {
	if s != nil {
		itc.SetDescription(*s)
	}
	return itc
}

// SetMeasurementUnit sets the "measurement_unit" field.
func (itc *InventoryTypeCreate) SetMeasurementUnit(s string) *InventoryTypeCreate {
	itc.mutation.SetMeasurementUnit(s)
	return itc
}

// SetNillableMeasurementUnit sets the "measurement_unit" field if the given value is not nil.
func (itc *InventoryTypeCreate) SetNillableMeasurementUnit(s *string) *InventoryTypeCreate {
	if s != nil {
		itc.SetMeasurementUnit(*s)
	}
	return itc
}

// SetID sets the "id" field.
func (itc *InventoryTypeCreate) SetID(s string) *InventoryTypeCreate {
	itc.mutation.SetID(s)
	return itc
}

// AddAttributeIDs adds the "attributes" edge to the InventoryAttribute entity by IDs.
func (itc *InventoryTypeCreate) AddAttributeIDs(ids ...string) *InventoryTypeCreate {
	itc.mutation.AddAttributeIDs(ids...)
	return itc
}

// AddAttributes adds the "attributes" edges to the InventoryAttribute entity.
func (itc *InventoryTypeCreate) AddAttributes(i ...*InventoryAttribute) *InventoryTypeCreate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return itc.AddAttributeIDs(ids...)
}

// AddPlaceInventoryIDs adds the "place_inventories" edge to the PlaceInventory entity by IDs.
func (itc *InventoryTypeCreate) AddPlaceInventoryIDs(ids ...string) *InventoryTypeCreate {
	itc.mutation.AddPlaceInventoryIDs(ids...)
	return itc
}

// AddPlaceInventories adds the "place_inventories" edges to the PlaceInventory entity.
func (itc *InventoryTypeCreate) AddPlaceInventories(p ...*PlaceInventory) *InventoryTypeCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return itc.AddPlaceInventoryIDs(ids...)
}

// Mutation returns the InventoryTypeMutation object of the builder.
func (itc *InventoryTypeCreate) Mutation() *InventoryTypeMutation {
	return itc.mutation
}

// Save creates the InventoryType in the database.
func (itc *InventoryTypeCreate) Save(ctx context.Context) (*InventoryType, error) {
	return withHooks(ctx, itc.sqlSave, itc.mutation, itc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (itc *InventoryTypeCreate) SaveX(ctx context.Context) *InventoryType {
	v, err := itc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (itc *InventoryTypeCreate) Exec(ctx context.Context) error {
	_, err := itc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itc *InventoryTypeCreate) ExecX(ctx context.Context) {
	if err := itc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (itc *InventoryTypeCreate) check() error {
	if _, ok := itc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`placio_api: missing required field "InventoryType.name"`)}
	}
	if v, ok := itc.mutation.ID(); ok {
		if err := inventorytype.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`placio_api: validator failed for field "InventoryType.id": %w`, err)}
		}
	}
	return nil
}

func (itc *InventoryTypeCreate) sqlSave(ctx context.Context) (*InventoryType, error) {
	if err := itc.check(); err != nil {
		return nil, err
	}
	_node, _spec := itc.createSpec()
	if err := sqlgraph.CreateNode(ctx, itc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected InventoryType.ID type: %T", _spec.ID.Value)
		}
	}
	itc.mutation.id = &_node.ID
	itc.mutation.done = true
	return _node, nil
}

func (itc *InventoryTypeCreate) createSpec() (*InventoryType, *sqlgraph.CreateSpec) {
	var (
		_node = &InventoryType{config: itc.config}
		_spec = sqlgraph.NewCreateSpec(inventorytype.Table, sqlgraph.NewFieldSpec(inventorytype.FieldID, field.TypeString))
	)
	if id, ok := itc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := itc.mutation.Name(); ok {
		_spec.SetField(inventorytype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := itc.mutation.Description(); ok {
		_spec.SetField(inventorytype.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := itc.mutation.MeasurementUnit(); ok {
		_spec.SetField(inventorytype.FieldMeasurementUnit, field.TypeString, value)
		_node.MeasurementUnit = value
	}
	if nodes := itc.mutation.AttributesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inventorytype.AttributesTable,
			Columns: []string{inventorytype.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(inventoryattribute.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := itc.mutation.PlaceInventoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   inventorytype.PlaceInventoriesTable,
			Columns: []string{inventorytype.PlaceInventoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InventoryTypeCreateBulk is the builder for creating many InventoryType entities in bulk.
type InventoryTypeCreateBulk struct {
	config
	err      error
	builders []*InventoryTypeCreate
}

// Save creates the InventoryType entities in the database.
func (itcb *InventoryTypeCreateBulk) Save(ctx context.Context) ([]*InventoryType, error) {
	if itcb.err != nil {
		return nil, itcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(itcb.builders))
	nodes := make([]*InventoryType, len(itcb.builders))
	mutators := make([]Mutator, len(itcb.builders))
	for i := range itcb.builders {
		func(i int, root context.Context) {
			builder := itcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InventoryTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, itcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, itcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, itcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (itcb *InventoryTypeCreateBulk) SaveX(ctx context.Context) []*InventoryType {
	v, err := itcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (itcb *InventoryTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := itcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itcb *InventoryTypeCreateBulk) ExecX(ctx context.Context) {
	if err := itcb.Exec(ctx); err != nil {
		panic(err)
	}
}
