// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/inventoryattribute"
	"placio-app/ent/inventorytype"
	"placio-app/ent/placeinventory"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InventoryTypeUpdate is the builder for updating InventoryType entities.
type InventoryTypeUpdate struct {
	config
	hooks    []Hook
	mutation *InventoryTypeMutation
}

// Where appends a list predicates to the InventoryTypeUpdate builder.
func (itu *InventoryTypeUpdate) Where(ps ...predicate.InventoryType) *InventoryTypeUpdate {
	itu.mutation.Where(ps...)
	return itu
}

// SetName sets the "name" field.
func (itu *InventoryTypeUpdate) SetName(s string) *InventoryTypeUpdate {
	itu.mutation.SetName(s)
	return itu
}

// SetIndustryType sets the "industry_type" field.
func (itu *InventoryTypeUpdate) SetIndustryType(it inventorytype.IndustryType) *InventoryTypeUpdate {
	itu.mutation.SetIndustryType(it)
	return itu
}

// SetNillableIndustryType sets the "industry_type" field if the given value is not nil.
func (itu *InventoryTypeUpdate) SetNillableIndustryType(it *inventorytype.IndustryType) *InventoryTypeUpdate {
	if it != nil {
		itu.SetIndustryType(*it)
	}
	return itu
}

// ClearIndustryType clears the value of the "industry_type" field.
func (itu *InventoryTypeUpdate) ClearIndustryType() *InventoryTypeUpdate {
	itu.mutation.ClearIndustryType()
	return itu
}

// SetMeasurementUnit sets the "measurement_unit" field.
func (itu *InventoryTypeUpdate) SetMeasurementUnit(s string) *InventoryTypeUpdate {
	itu.mutation.SetMeasurementUnit(s)
	return itu
}

// SetNillableMeasurementUnit sets the "measurement_unit" field if the given value is not nil.
func (itu *InventoryTypeUpdate) SetNillableMeasurementUnit(s *string) *InventoryTypeUpdate {
	if s != nil {
		itu.SetMeasurementUnit(*s)
	}
	return itu
}

// ClearMeasurementUnit clears the value of the "measurement_unit" field.
func (itu *InventoryTypeUpdate) ClearMeasurementUnit() *InventoryTypeUpdate {
	itu.mutation.ClearMeasurementUnit()
	return itu
}

// AddAttributeIDs adds the "attributes" edge to the InventoryAttribute entity by IDs.
func (itu *InventoryTypeUpdate) AddAttributeIDs(ids ...string) *InventoryTypeUpdate {
	itu.mutation.AddAttributeIDs(ids...)
	return itu
}

// AddAttributes adds the "attributes" edges to the InventoryAttribute entity.
func (itu *InventoryTypeUpdate) AddAttributes(i ...*InventoryAttribute) *InventoryTypeUpdate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return itu.AddAttributeIDs(ids...)
}

// AddPlaceInventoryIDs adds the "place_inventories" edge to the PlaceInventory entity by IDs.
func (itu *InventoryTypeUpdate) AddPlaceInventoryIDs(ids ...string) *InventoryTypeUpdate {
	itu.mutation.AddPlaceInventoryIDs(ids...)
	return itu
}

// AddPlaceInventories adds the "place_inventories" edges to the PlaceInventory entity.
func (itu *InventoryTypeUpdate) AddPlaceInventories(p ...*PlaceInventory) *InventoryTypeUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return itu.AddPlaceInventoryIDs(ids...)
}

// Mutation returns the InventoryTypeMutation object of the builder.
func (itu *InventoryTypeUpdate) Mutation() *InventoryTypeMutation {
	return itu.mutation
}

// ClearAttributes clears all "attributes" edges to the InventoryAttribute entity.
func (itu *InventoryTypeUpdate) ClearAttributes() *InventoryTypeUpdate {
	itu.mutation.ClearAttributes()
	return itu
}

// RemoveAttributeIDs removes the "attributes" edge to InventoryAttribute entities by IDs.
func (itu *InventoryTypeUpdate) RemoveAttributeIDs(ids ...string) *InventoryTypeUpdate {
	itu.mutation.RemoveAttributeIDs(ids...)
	return itu
}

// RemoveAttributes removes "attributes" edges to InventoryAttribute entities.
func (itu *InventoryTypeUpdate) RemoveAttributes(i ...*InventoryAttribute) *InventoryTypeUpdate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return itu.RemoveAttributeIDs(ids...)
}

// ClearPlaceInventories clears all "place_inventories" edges to the PlaceInventory entity.
func (itu *InventoryTypeUpdate) ClearPlaceInventories() *InventoryTypeUpdate {
	itu.mutation.ClearPlaceInventories()
	return itu
}

// RemovePlaceInventoryIDs removes the "place_inventories" edge to PlaceInventory entities by IDs.
func (itu *InventoryTypeUpdate) RemovePlaceInventoryIDs(ids ...string) *InventoryTypeUpdate {
	itu.mutation.RemovePlaceInventoryIDs(ids...)
	return itu
}

// RemovePlaceInventories removes "place_inventories" edges to PlaceInventory entities.
func (itu *InventoryTypeUpdate) RemovePlaceInventories(p ...*PlaceInventory) *InventoryTypeUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return itu.RemovePlaceInventoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (itu *InventoryTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, itu.sqlSave, itu.mutation, itu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (itu *InventoryTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := itu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (itu *InventoryTypeUpdate) Exec(ctx context.Context) error {
	_, err := itu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itu *InventoryTypeUpdate) ExecX(ctx context.Context) {
	if err := itu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (itu *InventoryTypeUpdate) check() error {
	if v, ok := itu.mutation.IndustryType(); ok {
		if err := inventorytype.IndustryTypeValidator(v); err != nil {
			return &ValidationError{Name: "industry_type", err: fmt.Errorf(`ent: validator failed for field "InventoryType.industry_type": %w`, err)}
		}
	}
	return nil
}

func (itu *InventoryTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := itu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(inventorytype.Table, inventorytype.Columns, sqlgraph.NewFieldSpec(inventorytype.FieldID, field.TypeString))
	if ps := itu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := itu.mutation.Name(); ok {
		_spec.SetField(inventorytype.FieldName, field.TypeString, value)
	}
	if value, ok := itu.mutation.IndustryType(); ok {
		_spec.SetField(inventorytype.FieldIndustryType, field.TypeEnum, value)
	}
	if itu.mutation.IndustryTypeCleared() {
		_spec.ClearField(inventorytype.FieldIndustryType, field.TypeEnum)
	}
	if value, ok := itu.mutation.MeasurementUnit(); ok {
		_spec.SetField(inventorytype.FieldMeasurementUnit, field.TypeString, value)
	}
	if itu.mutation.MeasurementUnitCleared() {
		_spec.ClearField(inventorytype.FieldMeasurementUnit, field.TypeString)
	}
	if itu.mutation.AttributesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := itu.mutation.RemovedAttributesIDs(); len(nodes) > 0 && !itu.mutation.AttributesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := itu.mutation.AttributesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if itu.mutation.PlaceInventoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := itu.mutation.RemovedPlaceInventoriesIDs(); len(nodes) > 0 && !itu.mutation.PlaceInventoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := itu.mutation.PlaceInventoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, itu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inventorytype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	itu.mutation.done = true
	return n, nil
}

// InventoryTypeUpdateOne is the builder for updating a single InventoryType entity.
type InventoryTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InventoryTypeMutation
}

// SetName sets the "name" field.
func (ituo *InventoryTypeUpdateOne) SetName(s string) *InventoryTypeUpdateOne {
	ituo.mutation.SetName(s)
	return ituo
}

// SetIndustryType sets the "industry_type" field.
func (ituo *InventoryTypeUpdateOne) SetIndustryType(it inventorytype.IndustryType) *InventoryTypeUpdateOne {
	ituo.mutation.SetIndustryType(it)
	return ituo
}

// SetNillableIndustryType sets the "industry_type" field if the given value is not nil.
func (ituo *InventoryTypeUpdateOne) SetNillableIndustryType(it *inventorytype.IndustryType) *InventoryTypeUpdateOne {
	if it != nil {
		ituo.SetIndustryType(*it)
	}
	return ituo
}

// ClearIndustryType clears the value of the "industry_type" field.
func (ituo *InventoryTypeUpdateOne) ClearIndustryType() *InventoryTypeUpdateOne {
	ituo.mutation.ClearIndustryType()
	return ituo
}

// SetMeasurementUnit sets the "measurement_unit" field.
func (ituo *InventoryTypeUpdateOne) SetMeasurementUnit(s string) *InventoryTypeUpdateOne {
	ituo.mutation.SetMeasurementUnit(s)
	return ituo
}

// SetNillableMeasurementUnit sets the "measurement_unit" field if the given value is not nil.
func (ituo *InventoryTypeUpdateOne) SetNillableMeasurementUnit(s *string) *InventoryTypeUpdateOne {
	if s != nil {
		ituo.SetMeasurementUnit(*s)
	}
	return ituo
}

// ClearMeasurementUnit clears the value of the "measurement_unit" field.
func (ituo *InventoryTypeUpdateOne) ClearMeasurementUnit() *InventoryTypeUpdateOne {
	ituo.mutation.ClearMeasurementUnit()
	return ituo
}

// AddAttributeIDs adds the "attributes" edge to the InventoryAttribute entity by IDs.
func (ituo *InventoryTypeUpdateOne) AddAttributeIDs(ids ...string) *InventoryTypeUpdateOne {
	ituo.mutation.AddAttributeIDs(ids...)
	return ituo
}

// AddAttributes adds the "attributes" edges to the InventoryAttribute entity.
func (ituo *InventoryTypeUpdateOne) AddAttributes(i ...*InventoryAttribute) *InventoryTypeUpdateOne {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ituo.AddAttributeIDs(ids...)
}

// AddPlaceInventoryIDs adds the "place_inventories" edge to the PlaceInventory entity by IDs.
func (ituo *InventoryTypeUpdateOne) AddPlaceInventoryIDs(ids ...string) *InventoryTypeUpdateOne {
	ituo.mutation.AddPlaceInventoryIDs(ids...)
	return ituo
}

// AddPlaceInventories adds the "place_inventories" edges to the PlaceInventory entity.
func (ituo *InventoryTypeUpdateOne) AddPlaceInventories(p ...*PlaceInventory) *InventoryTypeUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ituo.AddPlaceInventoryIDs(ids...)
}

// Mutation returns the InventoryTypeMutation object of the builder.
func (ituo *InventoryTypeUpdateOne) Mutation() *InventoryTypeMutation {
	return ituo.mutation
}

// ClearAttributes clears all "attributes" edges to the InventoryAttribute entity.
func (ituo *InventoryTypeUpdateOne) ClearAttributes() *InventoryTypeUpdateOne {
	ituo.mutation.ClearAttributes()
	return ituo
}

// RemoveAttributeIDs removes the "attributes" edge to InventoryAttribute entities by IDs.
func (ituo *InventoryTypeUpdateOne) RemoveAttributeIDs(ids ...string) *InventoryTypeUpdateOne {
	ituo.mutation.RemoveAttributeIDs(ids...)
	return ituo
}

// RemoveAttributes removes "attributes" edges to InventoryAttribute entities.
func (ituo *InventoryTypeUpdateOne) RemoveAttributes(i ...*InventoryAttribute) *InventoryTypeUpdateOne {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ituo.RemoveAttributeIDs(ids...)
}

// ClearPlaceInventories clears all "place_inventories" edges to the PlaceInventory entity.
func (ituo *InventoryTypeUpdateOne) ClearPlaceInventories() *InventoryTypeUpdateOne {
	ituo.mutation.ClearPlaceInventories()
	return ituo
}

// RemovePlaceInventoryIDs removes the "place_inventories" edge to PlaceInventory entities by IDs.
func (ituo *InventoryTypeUpdateOne) RemovePlaceInventoryIDs(ids ...string) *InventoryTypeUpdateOne {
	ituo.mutation.RemovePlaceInventoryIDs(ids...)
	return ituo
}

// RemovePlaceInventories removes "place_inventories" edges to PlaceInventory entities.
func (ituo *InventoryTypeUpdateOne) RemovePlaceInventories(p ...*PlaceInventory) *InventoryTypeUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ituo.RemovePlaceInventoryIDs(ids...)
}

// Where appends a list predicates to the InventoryTypeUpdate builder.
func (ituo *InventoryTypeUpdateOne) Where(ps ...predicate.InventoryType) *InventoryTypeUpdateOne {
	ituo.mutation.Where(ps...)
	return ituo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ituo *InventoryTypeUpdateOne) Select(field string, fields ...string) *InventoryTypeUpdateOne {
	ituo.fields = append([]string{field}, fields...)
	return ituo
}

// Save executes the query and returns the updated InventoryType entity.
func (ituo *InventoryTypeUpdateOne) Save(ctx context.Context) (*InventoryType, error) {
	return withHooks(ctx, ituo.sqlSave, ituo.mutation, ituo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ituo *InventoryTypeUpdateOne) SaveX(ctx context.Context) *InventoryType {
	node, err := ituo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ituo *InventoryTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ituo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ituo *InventoryTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ituo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ituo *InventoryTypeUpdateOne) check() error {
	if v, ok := ituo.mutation.IndustryType(); ok {
		if err := inventorytype.IndustryTypeValidator(v); err != nil {
			return &ValidationError{Name: "industry_type", err: fmt.Errorf(`ent: validator failed for field "InventoryType.industry_type": %w`, err)}
		}
	}
	return nil
}

func (ituo *InventoryTypeUpdateOne) sqlSave(ctx context.Context) (_node *InventoryType, err error) {
	if err := ituo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(inventorytype.Table, inventorytype.Columns, sqlgraph.NewFieldSpec(inventorytype.FieldID, field.TypeString))
	id, ok := ituo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "InventoryType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ituo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inventorytype.FieldID)
		for _, f := range fields {
			if !inventorytype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != inventorytype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ituo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ituo.mutation.Name(); ok {
		_spec.SetField(inventorytype.FieldName, field.TypeString, value)
	}
	if value, ok := ituo.mutation.IndustryType(); ok {
		_spec.SetField(inventorytype.FieldIndustryType, field.TypeEnum, value)
	}
	if ituo.mutation.IndustryTypeCleared() {
		_spec.ClearField(inventorytype.FieldIndustryType, field.TypeEnum)
	}
	if value, ok := ituo.mutation.MeasurementUnit(); ok {
		_spec.SetField(inventorytype.FieldMeasurementUnit, field.TypeString, value)
	}
	if ituo.mutation.MeasurementUnitCleared() {
		_spec.ClearField(inventorytype.FieldMeasurementUnit, field.TypeString)
	}
	if ituo.mutation.AttributesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ituo.mutation.RemovedAttributesIDs(); len(nodes) > 0 && !ituo.mutation.AttributesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ituo.mutation.AttributesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ituo.mutation.PlaceInventoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ituo.mutation.RemovedPlaceInventoriesIDs(); len(nodes) > 0 && !ituo.mutation.PlaceInventoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ituo.mutation.PlaceInventoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &InventoryType{config: ituo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ituo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inventorytype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ituo.mutation.done = true
	return _node, nil
}
