// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/category"
	"placio-app/ent/categoryassignment"
	"placio-app/ent/place"
	"placio-app/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryAssignmentCreate is the builder for creating a CategoryAssignment entity.
type CategoryAssignmentCreate struct {
	config
	mutation *CategoryAssignmentMutation
	hooks    []Hook
}

// SetEntityID sets the "entity_id" field.
func (cac *CategoryAssignmentCreate) SetEntityID(s string) *CategoryAssignmentCreate {
	cac.mutation.SetEntityID(s)
	return cac
}

// SetNillableEntityID sets the "entity_id" field if the given value is not nil.
func (cac *CategoryAssignmentCreate) SetNillableEntityID(s *string) *CategoryAssignmentCreate {
	if s != nil {
		cac.SetEntityID(*s)
	}
	return cac
}

// SetEntityType sets the "entity_type" field.
func (cac *CategoryAssignmentCreate) SetEntityType(s string) *CategoryAssignmentCreate {
	cac.mutation.SetEntityType(s)
	return cac
}

// SetNillableEntityType sets the "entity_type" field if the given value is not nil.
func (cac *CategoryAssignmentCreate) SetNillableEntityType(s *string) *CategoryAssignmentCreate {
	if s != nil {
		cac.SetEntityType(*s)
	}
	return cac
}

// SetCategoryID sets the "category_id" field.
func (cac *CategoryAssignmentCreate) SetCategoryID(s string) *CategoryAssignmentCreate {
	cac.mutation.SetCategoryID(s)
	return cac
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (cac *CategoryAssignmentCreate) SetNillableCategoryID(s *string) *CategoryAssignmentCreate {
	if s != nil {
		cac.SetCategoryID(*s)
	}
	return cac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cac *CategoryAssignmentCreate) SetUserID(id string) *CategoryAssignmentCreate {
	cac.mutation.SetUserID(id)
	return cac
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cac *CategoryAssignmentCreate) SetNillableUserID(id *string) *CategoryAssignmentCreate {
	if id != nil {
		cac = cac.SetUserID(*id)
	}
	return cac
}

// SetUser sets the "user" edge to the User entity.
func (cac *CategoryAssignmentCreate) SetUser(u *User) *CategoryAssignmentCreate {
	return cac.SetUserID(u.ID)
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (cac *CategoryAssignmentCreate) SetBusinessID(id string) *CategoryAssignmentCreate {
	cac.mutation.SetBusinessID(id)
	return cac
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (cac *CategoryAssignmentCreate) SetNillableBusinessID(id *string) *CategoryAssignmentCreate {
	if id != nil {
		cac = cac.SetBusinessID(*id)
	}
	return cac
}

// SetBusiness sets the "business" edge to the Business entity.
func (cac *CategoryAssignmentCreate) SetBusiness(b *Business) *CategoryAssignmentCreate {
	return cac.SetBusinessID(b.ID)
}

// SetPlaceID sets the "place" edge to the Place entity by ID.
func (cac *CategoryAssignmentCreate) SetPlaceID(id string) *CategoryAssignmentCreate {
	cac.mutation.SetPlaceID(id)
	return cac
}

// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
func (cac *CategoryAssignmentCreate) SetNillablePlaceID(id *string) *CategoryAssignmentCreate {
	if id != nil {
		cac = cac.SetPlaceID(*id)
	}
	return cac
}

// SetPlace sets the "place" edge to the Place entity.
func (cac *CategoryAssignmentCreate) SetPlace(p *Place) *CategoryAssignmentCreate {
	return cac.SetPlaceID(p.ID)
}

// SetCategory sets the "category" edge to the Category entity.
func (cac *CategoryAssignmentCreate) SetCategory(c *Category) *CategoryAssignmentCreate {
	return cac.SetCategoryID(c.ID)
}

// Mutation returns the CategoryAssignmentMutation object of the builder.
func (cac *CategoryAssignmentCreate) Mutation() *CategoryAssignmentMutation {
	return cac.mutation
}

// Save creates the CategoryAssignment in the database.
func (cac *CategoryAssignmentCreate) Save(ctx context.Context) (*CategoryAssignment, error) {
	return withHooks(ctx, cac.sqlSave, cac.mutation, cac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cac *CategoryAssignmentCreate) SaveX(ctx context.Context) *CategoryAssignment {
	v, err := cac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cac *CategoryAssignmentCreate) Exec(ctx context.Context) error {
	_, err := cac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cac *CategoryAssignmentCreate) ExecX(ctx context.Context) {
	if err := cac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cac *CategoryAssignmentCreate) check() error {
	return nil
}

func (cac *CategoryAssignmentCreate) sqlSave(ctx context.Context) (*CategoryAssignment, error) {
	if err := cac.check(); err != nil {
		return nil, err
	}
	_node, _spec := cac.createSpec()
	if err := sqlgraph.CreateNode(ctx, cac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected CategoryAssignment.ID type: %T", _spec.ID.Value)
		}
	}
	cac.mutation.id = &_node.ID
	cac.mutation.done = true
	return _node, nil
}

func (cac *CategoryAssignmentCreate) createSpec() (*CategoryAssignment, *sqlgraph.CreateSpec) {
	var (
		_node = &CategoryAssignment{config: cac.config}
		_spec = sqlgraph.NewCreateSpec(categoryassignment.Table, sqlgraph.NewFieldSpec(categoryassignment.FieldID, field.TypeString))
	)
	if value, ok := cac.mutation.EntityType(); ok {
		_spec.SetField(categoryassignment.FieldEntityType, field.TypeString, value)
		_node.EntityType = value
	}
	if nodes := cac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   categoryassignment.UserTable,
			Columns: []string{categoryassignment.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EntityID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cac.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   categoryassignment.BusinessTable,
			Columns: []string{categoryassignment.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EntityID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cac.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   categoryassignment.PlaceTable,
			Columns: []string{categoryassignment.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EntityID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cac.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   categoryassignment.CategoryTable,
			Columns: []string{categoryassignment.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CategoryAssignmentCreateBulk is the builder for creating many CategoryAssignment entities in bulk.
type CategoryAssignmentCreateBulk struct {
	config
	err      error
	builders []*CategoryAssignmentCreate
}

// Save creates the CategoryAssignment entities in the database.
func (cacb *CategoryAssignmentCreateBulk) Save(ctx context.Context) ([]*CategoryAssignment, error) {
	if cacb.err != nil {
		return nil, cacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cacb.builders))
	nodes := make([]*CategoryAssignment, len(cacb.builders))
	mutators := make([]Mutator, len(cacb.builders))
	for i := range cacb.builders {
		func(i int, root context.Context) {
			builder := cacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryAssignmentMutation)
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
					_, err = mutators[i+1].Mutate(root, cacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cacb *CategoryAssignmentCreateBulk) SaveX(ctx context.Context) []*CategoryAssignment {
	v, err := cacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cacb *CategoryAssignmentCreateBulk) Exec(ctx context.Context) error {
	_, err := cacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cacb *CategoryAssignmentCreateBulk) ExecX(ctx context.Context) {
	if err := cacb.Exec(ctx); err != nil {
		panic(err)
	}
}
