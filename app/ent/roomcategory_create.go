// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/amenity"
	"placio-app/ent/media"
	"placio-app/ent/place"
	"placio-app/ent/room"
	"placio-app/ent/roomcategory"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoomCategoryCreate is the builder for creating a RoomCategory entity.
type RoomCategoryCreate struct {
	config
	mutation *RoomCategoryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rcc *RoomCategoryCreate) SetName(s string) *RoomCategoryCreate {
	rcc.mutation.SetName(s)
	return rcc
}

// SetDescription sets the "description" field.
func (rcc *RoomCategoryCreate) SetDescription(s string) *RoomCategoryCreate {
	rcc.mutation.SetDescription(s)
	return rcc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rcc *RoomCategoryCreate) SetNillableDescription(s *string) *RoomCategoryCreate {
	if s != nil {
		rcc.SetDescription(*s)
	}
	return rcc
}

// SetPrice sets the "price" field.
func (rcc *RoomCategoryCreate) SetPrice(s string) *RoomCategoryCreate {
	rcc.mutation.SetPrice(s)
	return rcc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (rcc *RoomCategoryCreate) SetNillablePrice(s *string) *RoomCategoryCreate {
	if s != nil {
		rcc.SetPrice(*s)
	}
	return rcc
}

// SetID sets the "id" field.
func (rcc *RoomCategoryCreate) SetID(s string) *RoomCategoryCreate {
	rcc.mutation.SetID(s)
	return rcc
}

// AddPlaceIDs adds the "place" edge to the Place entity by IDs.
func (rcc *RoomCategoryCreate) AddPlaceIDs(ids ...string) *RoomCategoryCreate {
	rcc.mutation.AddPlaceIDs(ids...)
	return rcc
}

// AddPlace adds the "place" edges to the Place entity.
func (rcc *RoomCategoryCreate) AddPlace(p ...*Place) *RoomCategoryCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rcc.AddPlaceIDs(ids...)
}

// AddRoomIDs adds the "rooms" edge to the Room entity by IDs.
func (rcc *RoomCategoryCreate) AddRoomIDs(ids ...string) *RoomCategoryCreate {
	rcc.mutation.AddRoomIDs(ids...)
	return rcc
}

// AddRooms adds the "rooms" edges to the Room entity.
func (rcc *RoomCategoryCreate) AddRooms(r ...*Room) *RoomCategoryCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcc.AddRoomIDs(ids...)
}

// AddMediumIDs adds the "media" edge to the Media entity by IDs.
func (rcc *RoomCategoryCreate) AddMediumIDs(ids ...string) *RoomCategoryCreate {
	rcc.mutation.AddMediumIDs(ids...)
	return rcc
}

// AddMedia adds the "media" edges to the Media entity.
func (rcc *RoomCategoryCreate) AddMedia(m ...*Media) *RoomCategoryCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return rcc.AddMediumIDs(ids...)
}

// AddAmenityIDs adds the "amenities" edge to the Amenity entity by IDs.
func (rcc *RoomCategoryCreate) AddAmenityIDs(ids ...string) *RoomCategoryCreate {
	rcc.mutation.AddAmenityIDs(ids...)
	return rcc
}

// AddAmenities adds the "amenities" edges to the Amenity entity.
func (rcc *RoomCategoryCreate) AddAmenities(a ...*Amenity) *RoomCategoryCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rcc.AddAmenityIDs(ids...)
}

// Mutation returns the RoomCategoryMutation object of the builder.
func (rcc *RoomCategoryCreate) Mutation() *RoomCategoryMutation {
	return rcc.mutation
}

// Save creates the RoomCategory in the database.
func (rcc *RoomCategoryCreate) Save(ctx context.Context) (*RoomCategory, error) {
	return withHooks(ctx, rcc.sqlSave, rcc.mutation, rcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rcc *RoomCategoryCreate) SaveX(ctx context.Context) *RoomCategory {
	v, err := rcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcc *RoomCategoryCreate) Exec(ctx context.Context) error {
	_, err := rcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcc *RoomCategoryCreate) ExecX(ctx context.Context) {
	if err := rcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rcc *RoomCategoryCreate) check() error {
	if _, ok := rcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "RoomCategory.name"`)}
	}
	if v, ok := rcc.mutation.ID(); ok {
		if err := roomcategory.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "RoomCategory.id": %w`, err)}
		}
	}
	return nil
}

func (rcc *RoomCategoryCreate) sqlSave(ctx context.Context) (*RoomCategory, error) {
	if err := rcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected RoomCategory.ID type: %T", _spec.ID.Value)
		}
	}
	rcc.mutation.id = &_node.ID
	rcc.mutation.done = true
	return _node, nil
}

func (rcc *RoomCategoryCreate) createSpec() (*RoomCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &RoomCategory{config: rcc.config}
		_spec = sqlgraph.NewCreateSpec(roomcategory.Table, sqlgraph.NewFieldSpec(roomcategory.FieldID, field.TypeString))
	)
	if id, ok := rcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rcc.mutation.Name(); ok {
		_spec.SetField(roomcategory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rcc.mutation.Description(); ok {
		_spec.SetField(roomcategory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rcc.mutation.Price(); ok {
		_spec.SetField(roomcategory.FieldPrice, field.TypeString, value)
		_node.Price = value
	}
	if nodes := rcc.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   roomcategory.PlaceTable,
			Columns: roomcategory.PlacePrimaryKey,
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
	if nodes := rcc.mutation.RoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   roomcategory.RoomsTable,
			Columns: roomcategory.RoomsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rcc.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   roomcategory.MediaTable,
			Columns: roomcategory.MediaPrimaryKey,
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
	if nodes := rcc.mutation.AmenitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   roomcategory.AmenitiesTable,
			Columns: roomcategory.AmenitiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(amenity.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoomCategoryCreateBulk is the builder for creating many RoomCategory entities in bulk.
type RoomCategoryCreateBulk struct {
	config
	err      error
	builders []*RoomCategoryCreate
}

// Save creates the RoomCategory entities in the database.
func (rccb *RoomCategoryCreateBulk) Save(ctx context.Context) ([]*RoomCategory, error) {
	if rccb.err != nil {
		return nil, rccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rccb.builders))
	nodes := make([]*RoomCategory, len(rccb.builders))
	mutators := make([]Mutator, len(rccb.builders))
	for i := range rccb.builders {
		func(i int, root context.Context) {
			builder := rccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoomCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, rccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rccb *RoomCategoryCreateBulk) SaveX(ctx context.Context) []*RoomCategory {
	v, err := rccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rccb *RoomCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := rccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rccb *RoomCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := rccb.Exec(ctx); err != nil {
		panic(err)
	}
}
