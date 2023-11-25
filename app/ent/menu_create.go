// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/category"
	"placio-app/ent/media"
	"placio-app/ent/menu"
	"placio-app/ent/menuitem"
	"placio-app/ent/place"
	"placio-app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuCreate is the builder for creating a Menu entity.
type MenuCreate struct {
	config
	mutation *MenuMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (mc *MenuCreate) SetName(s string) *MenuCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetDeletedAt sets the "deleted_at" field.
func (mc *MenuCreate) SetDeletedAt(s string) *MenuCreate {
	mc.mutation.SetDeletedAt(s)
	return mc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDeletedAt(s *string) *MenuCreate {
	if s != nil {
		mc.SetDeletedAt(*s)
	}
	return mc
}

// SetIsDeleted sets the "is_deleted" field.
func (mc *MenuCreate) SetIsDeleted(b bool) *MenuCreate {
	mc.mutation.SetIsDeleted(b)
	return mc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (mc *MenuCreate) SetNillableIsDeleted(b *bool) *MenuCreate {
	if b != nil {
		mc.SetIsDeleted(*b)
	}
	return mc
}

// SetDescription sets the "description" field.
func (mc *MenuCreate) SetDescription(s string) *MenuCreate {
	mc.mutation.SetDescription(s)
	return mc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDescription(s *string) *MenuCreate {
	if s != nil {
		mc.SetDescription(*s)
	}
	return mc
}

// SetOptions sets the "options" field.
func (mc *MenuCreate) SetOptions(s string) *MenuCreate {
	mc.mutation.SetOptions(s)
	return mc
}

// SetNillableOptions sets the "options" field if the given value is not nil.
func (mc *MenuCreate) SetNillableOptions(s *string) *MenuCreate {
	if s != nil {
		mc.SetOptions(*s)
	}
	return mc
}

// SetFoodType sets the "foodType" field.
func (mc *MenuCreate) SetFoodType(mt menu.FoodType) *MenuCreate {
	mc.mutation.SetFoodType(mt)
	return mc
}

// SetNillableFoodType sets the "foodType" field if the given value is not nil.
func (mc *MenuCreate) SetNillableFoodType(mt *menu.FoodType) *MenuCreate {
	if mt != nil {
		mc.SetFoodType(*mt)
	}
	return mc
}

// SetMenuItemType sets the "menuItemType" field.
func (mc *MenuCreate) SetMenuItemType(mit menu.MenuItemType) *MenuCreate {
	mc.mutation.SetMenuItemType(mit)
	return mc
}

// SetNillableMenuItemType sets the "menuItemType" field if the given value is not nil.
func (mc *MenuCreate) SetNillableMenuItemType(mit *menu.MenuItemType) *MenuCreate {
	if mit != nil {
		mc.SetMenuItemType(*mit)
	}
	return mc
}

// SetDrinkType sets the "drinkType" field.
func (mc *MenuCreate) SetDrinkType(mt menu.DrinkType) *MenuCreate {
	mc.mutation.SetDrinkType(mt)
	return mc
}

// SetNillableDrinkType sets the "drinkType" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDrinkType(mt *menu.DrinkType) *MenuCreate {
	if mt != nil {
		mc.SetDrinkType(*mt)
	}
	return mc
}

// SetDietaryType sets the "dietaryType" field.
func (mc *MenuCreate) SetDietaryType(mt menu.DietaryType) *MenuCreate {
	mc.mutation.SetDietaryType(mt)
	return mc
}

// SetNillableDietaryType sets the "dietaryType" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDietaryType(mt *menu.DietaryType) *MenuCreate {
	if mt != nil {
		mc.SetDietaryType(*mt)
	}
	return mc
}

// SetIsAvailable sets the "is_available" field.
func (mc *MenuCreate) SetIsAvailable(b bool) *MenuCreate {
	mc.mutation.SetIsAvailable(b)
	return mc
}

// SetNillableIsAvailable sets the "is_available" field if the given value is not nil.
func (mc *MenuCreate) SetNillableIsAvailable(b *bool) *MenuCreate {
	if b != nil {
		mc.SetIsAvailable(*b)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MenuCreate) SetUpdatedAt(t time.Time) *MenuCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MenuCreate) SetNillableUpdatedAt(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MenuCreate) SetID(s string) *MenuCreate {
	mc.mutation.SetID(s)
	return mc
}

// AddPlaceIDs adds the "place" edge to the Place entity by IDs.
func (mc *MenuCreate) AddPlaceIDs(ids ...string) *MenuCreate {
	mc.mutation.AddPlaceIDs(ids...)
	return mc
}

// AddPlace adds the "place" edges to the Place entity.
func (mc *MenuCreate) AddPlace(p ...*Place) *MenuCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddPlaceIDs(ids...)
}

// AddCategoryIDs adds the "categories" edge to the Category entity by IDs.
func (mc *MenuCreate) AddCategoryIDs(ids ...string) *MenuCreate {
	mc.mutation.AddCategoryIDs(ids...)
	return mc
}

// AddCategories adds the "categories" edges to the Category entity.
func (mc *MenuCreate) AddCategories(c ...*Category) *MenuCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mc.AddCategoryIDs(ids...)
}

// AddMenuItemIDs adds the "menu_items" edge to the MenuItem entity by IDs.
func (mc *MenuCreate) AddMenuItemIDs(ids ...string) *MenuCreate {
	mc.mutation.AddMenuItemIDs(ids...)
	return mc
}

// AddMenuItems adds the "menu_items" edges to the MenuItem entity.
func (mc *MenuCreate) AddMenuItems(m ...*MenuItem) *MenuCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMenuItemIDs(ids...)
}

// AddMediumIDs adds the "media" edge to the Media entity by IDs.
func (mc *MenuCreate) AddMediumIDs(ids ...string) *MenuCreate {
	mc.mutation.AddMediumIDs(ids...)
	return mc
}

// AddMedia adds the "media" edges to the Media entity.
func (mc *MenuCreate) AddMedia(m ...*Media) *MenuCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMediumIDs(ids...)
}

// AddCreatedByIDs adds the "created_by" edge to the User entity by IDs.
func (mc *MenuCreate) AddCreatedByIDs(ids ...string) *MenuCreate {
	mc.mutation.AddCreatedByIDs(ids...)
	return mc
}

// AddCreatedBy adds the "created_by" edges to the User entity.
func (mc *MenuCreate) AddCreatedBy(u ...*User) *MenuCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mc.AddCreatedByIDs(ids...)
}

// AddUpdatedByIDs adds the "updated_by" edge to the User entity by IDs.
func (mc *MenuCreate) AddUpdatedByIDs(ids ...string) *MenuCreate {
	mc.mutation.AddUpdatedByIDs(ids...)
	return mc
}

// AddUpdatedBy adds the "updated_by" edges to the User entity.
func (mc *MenuCreate) AddUpdatedBy(u ...*User) *MenuCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mc.AddUpdatedByIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mc *MenuCreate) Mutation() *MenuMutation {
	return mc.mutation
}

// Save creates the Menu in the database.
func (mc *MenuCreate) Save(ctx context.Context) (*Menu, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MenuCreate) SaveX(ctx context.Context) *Menu {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MenuCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MenuCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MenuCreate) defaults() {
	if _, ok := mc.mutation.IsDeleted(); !ok {
		v := menu.DefaultIsDeleted
		mc.mutation.SetIsDeleted(v)
	}
	if _, ok := mc.mutation.IsAvailable(); !ok {
		v := menu.DefaultIsAvailable
		mc.mutation.SetIsAvailable(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := menu.DefaultUpdatedAt
		mc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MenuCreate) check() error {
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Menu.name"`)}
	}
	if _, ok := mc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "Menu.is_deleted"`)}
	}
	if v, ok := mc.mutation.FoodType(); ok {
		if err := menu.FoodTypeValidator(v); err != nil {
			return &ValidationError{Name: "foodType", err: fmt.Errorf(`ent: validator failed for field "Menu.foodType": %w`, err)}
		}
	}
	if v, ok := mc.mutation.MenuItemType(); ok {
		if err := menu.MenuItemTypeValidator(v); err != nil {
			return &ValidationError{Name: "menuItemType", err: fmt.Errorf(`ent: validator failed for field "Menu.menuItemType": %w`, err)}
		}
	}
	if v, ok := mc.mutation.DrinkType(); ok {
		if err := menu.DrinkTypeValidator(v); err != nil {
			return &ValidationError{Name: "drinkType", err: fmt.Errorf(`ent: validator failed for field "Menu.drinkType": %w`, err)}
		}
	}
	if v, ok := mc.mutation.DietaryType(); ok {
		if err := menu.DietaryTypeValidator(v); err != nil {
			return &ValidationError{Name: "dietaryType", err: fmt.Errorf(`ent: validator failed for field "Menu.dietaryType": %w`, err)}
		}
	}
	if _, ok := mc.mutation.IsAvailable(); !ok {
		return &ValidationError{Name: "is_available", err: errors.New(`ent: missing required field "Menu.is_available"`)}
	}
	if v, ok := mc.mutation.ID(); ok {
		if err := menu.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Menu.id": %w`, err)}
		}
	}
	return nil
}

func (mc *MenuCreate) sqlSave(ctx context.Context) (*Menu, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Menu.ID type: %T", _spec.ID.Value)
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MenuCreate) createSpec() (*Menu, *sqlgraph.CreateSpec) {
	var (
		_node = &Menu{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(menu.Table, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.DeletedAt(); ok {
		_spec.SetField(menu.FieldDeletedAt, field.TypeString, value)
		_node.DeletedAt = value
	}
	if value, ok := mc.mutation.IsDeleted(); ok {
		_spec.SetField(menu.FieldIsDeleted, field.TypeBool, value)
		_node.IsDeleted = value
	}
	if value, ok := mc.mutation.Description(); ok {
		_spec.SetField(menu.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := mc.mutation.Options(); ok {
		_spec.SetField(menu.FieldOptions, field.TypeString, value)
		_node.Options = value
	}
	if value, ok := mc.mutation.FoodType(); ok {
		_spec.SetField(menu.FieldFoodType, field.TypeEnum, value)
		_node.FoodType = value
	}
	if value, ok := mc.mutation.MenuItemType(); ok {
		_spec.SetField(menu.FieldMenuItemType, field.TypeEnum, value)
		_node.MenuItemType = value
	}
	if value, ok := mc.mutation.DrinkType(); ok {
		_spec.SetField(menu.FieldDrinkType, field.TypeEnum, value)
		_node.DrinkType = value
	}
	if value, ok := mc.mutation.DietaryType(); ok {
		_spec.SetField(menu.FieldDietaryType, field.TypeEnum, value)
		_node.DietaryType = value
	}
	if value, ok := mc.mutation.IsAvailable(); ok {
		_spec.SetField(menu.FieldIsAvailable, field.TypeBool, value)
		_node.IsAvailable = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(menu.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := mc.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.PlaceTable,
			Columns: menu.PlacePrimaryKey,
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
	if nodes := mc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menu.CategoriesTable,
			Columns: menu.CategoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MenuItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menu.MenuItemsTable,
			Columns: menu.MenuItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menu.MediaTable,
			Columns: menu.MediaPrimaryKey,
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
	if nodes := mc.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.CreatedByTable,
			Columns: menu.CreatedByPrimaryKey,
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
	if nodes := mc.mutation.UpdatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.UpdatedByTable,
			Columns: menu.UpdatedByPrimaryKey,
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
	return _node, _spec
}

// MenuCreateBulk is the builder for creating many Menu entities in bulk.
type MenuCreateBulk struct {
	config
	err      error
	builders []*MenuCreate
}

// Save creates the Menu entities in the database.
func (mcb *MenuCreateBulk) Save(ctx context.Context) ([]*Menu, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Menu, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MenuMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MenuCreateBulk) SaveX(ctx context.Context) []*Menu {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MenuCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MenuCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
