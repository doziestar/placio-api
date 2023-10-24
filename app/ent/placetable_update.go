// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/order"
	"placio-app/ent/place"
	"placio-app/ent/placetable"
	"placio-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaceTableUpdate is the builder for updating PlaceTable entities.
type PlaceTableUpdate struct {
	config
	hooks    []Hook
	mutation *PlaceTableMutation
}

// Where appends a list predicates to the PlaceTableUpdate builder.
func (ptu *PlaceTableUpdate) Where(ps ...predicate.PlaceTable) *PlaceTableUpdate {
	ptu.mutation.Where(ps...)
	return ptu
}

// SetNumber sets the "number" field.
func (ptu *PlaceTableUpdate) SetNumber(i int) *PlaceTableUpdate {
	ptu.mutation.ResetNumber()
	ptu.mutation.SetNumber(i)
	return ptu
}

// AddNumber adds i to the "number" field.
func (ptu *PlaceTableUpdate) AddNumber(i int) *PlaceTableUpdate {
	ptu.mutation.AddNumber(i)
	return ptu
}

// SetQrCode sets the "qr_code" field.
func (ptu *PlaceTableUpdate) SetQrCode(s string) *PlaceTableUpdate {
	ptu.mutation.SetQrCode(s)
	return ptu
}

// SetNillableQrCode sets the "qr_code" field if the given value is not nil.
func (ptu *PlaceTableUpdate) SetNillableQrCode(s *string) *PlaceTableUpdate {
	if s != nil {
		ptu.SetQrCode(*s)
	}
	return ptu
}

// ClearQrCode clears the value of the "qr_code" field.
func (ptu *PlaceTableUpdate) ClearQrCode() *PlaceTableUpdate {
	ptu.mutation.ClearQrCode()
	return ptu
}

// SetPlaceID sets the "place" edge to the Place entity by ID.
func (ptu *PlaceTableUpdate) SetPlaceID(id string) *PlaceTableUpdate {
	ptu.mutation.SetPlaceID(id)
	return ptu
}

// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
func (ptu *PlaceTableUpdate) SetNillablePlaceID(id *string) *PlaceTableUpdate {
	if id != nil {
		ptu = ptu.SetPlaceID(*id)
	}
	return ptu
}

// SetPlace sets the "place" edge to the Place entity.
func (ptu *PlaceTableUpdate) SetPlace(p *Place) *PlaceTableUpdate {
	return ptu.SetPlaceID(p.ID)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (ptu *PlaceTableUpdate) AddOrderIDs(ids ...string) *PlaceTableUpdate {
	ptu.mutation.AddOrderIDs(ids...)
	return ptu
}

// AddOrders adds the "orders" edges to the Order entity.
func (ptu *PlaceTableUpdate) AddOrders(o ...*Order) *PlaceTableUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ptu.AddOrderIDs(ids...)
}

// Mutation returns the PlaceTableMutation object of the builder.
func (ptu *PlaceTableUpdate) Mutation() *PlaceTableMutation {
	return ptu.mutation
}

// ClearPlace clears the "place" edge to the Place entity.
func (ptu *PlaceTableUpdate) ClearPlace() *PlaceTableUpdate {
	ptu.mutation.ClearPlace()
	return ptu
}

// ClearOrders clears all "orders" edges to the Order entity.
func (ptu *PlaceTableUpdate) ClearOrders() *PlaceTableUpdate {
	ptu.mutation.ClearOrders()
	return ptu
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (ptu *PlaceTableUpdate) RemoveOrderIDs(ids ...string) *PlaceTableUpdate {
	ptu.mutation.RemoveOrderIDs(ids...)
	return ptu
}

// RemoveOrders removes "orders" edges to Order entities.
func (ptu *PlaceTableUpdate) RemoveOrders(o ...*Order) *PlaceTableUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ptu.RemoveOrderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptu *PlaceTableUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ptu.sqlSave, ptu.mutation, ptu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptu *PlaceTableUpdate) SaveX(ctx context.Context) int {
	affected, err := ptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptu *PlaceTableUpdate) Exec(ctx context.Context) error {
	_, err := ptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptu *PlaceTableUpdate) ExecX(ctx context.Context) {
	if err := ptu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ptu *PlaceTableUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(placetable.Table, placetable.Columns, sqlgraph.NewFieldSpec(placetable.FieldID, field.TypeString))
	if ps := ptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptu.mutation.Number(); ok {
		_spec.SetField(placetable.FieldNumber, field.TypeInt, value)
	}
	if value, ok := ptu.mutation.AddedNumber(); ok {
		_spec.AddField(placetable.FieldNumber, field.TypeInt, value)
	}
	if value, ok := ptu.mutation.QrCode(); ok {
		_spec.SetField(placetable.FieldQrCode, field.TypeString, value)
	}
	if ptu.mutation.QrCodeCleared() {
		_spec.ClearField(placetable.FieldQrCode, field.TypeString)
	}
	if ptu.mutation.PlaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placetable.PlaceTable,
			Columns: []string{placetable.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placetable.PlaceTable,
			Columns: []string{placetable.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ptu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   placetable.OrdersTable,
			Columns: placetable.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !ptu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   placetable.OrdersTable,
			Columns: placetable.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptu.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   placetable.OrdersTable,
			Columns: placetable.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{placetable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ptu.mutation.done = true
	return n, nil
}

// PlaceTableUpdateOne is the builder for updating a single PlaceTable entity.
type PlaceTableUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlaceTableMutation
}

// SetNumber sets the "number" field.
func (ptuo *PlaceTableUpdateOne) SetNumber(i int) *PlaceTableUpdateOne {
	ptuo.mutation.ResetNumber()
	ptuo.mutation.SetNumber(i)
	return ptuo
}

// AddNumber adds i to the "number" field.
func (ptuo *PlaceTableUpdateOne) AddNumber(i int) *PlaceTableUpdateOne {
	ptuo.mutation.AddNumber(i)
	return ptuo
}

// SetQrCode sets the "qr_code" field.
func (ptuo *PlaceTableUpdateOne) SetQrCode(s string) *PlaceTableUpdateOne {
	ptuo.mutation.SetQrCode(s)
	return ptuo
}

// SetNillableQrCode sets the "qr_code" field if the given value is not nil.
func (ptuo *PlaceTableUpdateOne) SetNillableQrCode(s *string) *PlaceTableUpdateOne {
	if s != nil {
		ptuo.SetQrCode(*s)
	}
	return ptuo
}

// ClearQrCode clears the value of the "qr_code" field.
func (ptuo *PlaceTableUpdateOne) ClearQrCode() *PlaceTableUpdateOne {
	ptuo.mutation.ClearQrCode()
	return ptuo
}

// SetPlaceID sets the "place" edge to the Place entity by ID.
func (ptuo *PlaceTableUpdateOne) SetPlaceID(id string) *PlaceTableUpdateOne {
	ptuo.mutation.SetPlaceID(id)
	return ptuo
}

// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
func (ptuo *PlaceTableUpdateOne) SetNillablePlaceID(id *string) *PlaceTableUpdateOne {
	if id != nil {
		ptuo = ptuo.SetPlaceID(*id)
	}
	return ptuo
}

// SetPlace sets the "place" edge to the Place entity.
func (ptuo *PlaceTableUpdateOne) SetPlace(p *Place) *PlaceTableUpdateOne {
	return ptuo.SetPlaceID(p.ID)
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (ptuo *PlaceTableUpdateOne) AddOrderIDs(ids ...string) *PlaceTableUpdateOne {
	ptuo.mutation.AddOrderIDs(ids...)
	return ptuo
}

// AddOrders adds the "orders" edges to the Order entity.
func (ptuo *PlaceTableUpdateOne) AddOrders(o ...*Order) *PlaceTableUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ptuo.AddOrderIDs(ids...)
}

// Mutation returns the PlaceTableMutation object of the builder.
func (ptuo *PlaceTableUpdateOne) Mutation() *PlaceTableMutation {
	return ptuo.mutation
}

// ClearPlace clears the "place" edge to the Place entity.
func (ptuo *PlaceTableUpdateOne) ClearPlace() *PlaceTableUpdateOne {
	ptuo.mutation.ClearPlace()
	return ptuo
}

// ClearOrders clears all "orders" edges to the Order entity.
func (ptuo *PlaceTableUpdateOne) ClearOrders() *PlaceTableUpdateOne {
	ptuo.mutation.ClearOrders()
	return ptuo
}

// RemoveOrderIDs removes the "orders" edge to Order entities by IDs.
func (ptuo *PlaceTableUpdateOne) RemoveOrderIDs(ids ...string) *PlaceTableUpdateOne {
	ptuo.mutation.RemoveOrderIDs(ids...)
	return ptuo
}

// RemoveOrders removes "orders" edges to Order entities.
func (ptuo *PlaceTableUpdateOne) RemoveOrders(o ...*Order) *PlaceTableUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ptuo.RemoveOrderIDs(ids...)
}

// Where appends a list predicates to the PlaceTableUpdate builder.
func (ptuo *PlaceTableUpdateOne) Where(ps ...predicate.PlaceTable) *PlaceTableUpdateOne {
	ptuo.mutation.Where(ps...)
	return ptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptuo *PlaceTableUpdateOne) Select(field string, fields ...string) *PlaceTableUpdateOne {
	ptuo.fields = append([]string{field}, fields...)
	return ptuo
}

// Save executes the query and returns the updated PlaceTable entity.
func (ptuo *PlaceTableUpdateOne) Save(ctx context.Context) (*PlaceTable, error) {
	return withHooks(ctx, ptuo.sqlSave, ptuo.mutation, ptuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptuo *PlaceTableUpdateOne) SaveX(ctx context.Context) *PlaceTable {
	node, err := ptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptuo *PlaceTableUpdateOne) Exec(ctx context.Context) error {
	_, err := ptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptuo *PlaceTableUpdateOne) ExecX(ctx context.Context) {
	if err := ptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ptuo *PlaceTableUpdateOne) sqlSave(ctx context.Context) (_node *PlaceTable, err error) {
	_spec := sqlgraph.NewUpdateSpec(placetable.Table, placetable.Columns, sqlgraph.NewFieldSpec(placetable.FieldID, field.TypeString))
	id, ok := ptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PlaceTable.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, placetable.FieldID)
		for _, f := range fields {
			if !placetable.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != placetable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptuo.mutation.Number(); ok {
		_spec.SetField(placetable.FieldNumber, field.TypeInt, value)
	}
	if value, ok := ptuo.mutation.AddedNumber(); ok {
		_spec.AddField(placetable.FieldNumber, field.TypeInt, value)
	}
	if value, ok := ptuo.mutation.QrCode(); ok {
		_spec.SetField(placetable.FieldQrCode, field.TypeString, value)
	}
	if ptuo.mutation.QrCodeCleared() {
		_spec.ClearField(placetable.FieldQrCode, field.TypeString)
	}
	if ptuo.mutation.PlaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placetable.PlaceTable,
			Columns: []string{placetable.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placetable.PlaceTable,
			Columns: []string{placetable.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ptuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   placetable.OrdersTable,
			Columns: placetable.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !ptuo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   placetable.OrdersTable,
			Columns: placetable.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptuo.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   placetable.OrdersTable,
			Columns: placetable.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PlaceTable{config: ptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{placetable.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ptuo.mutation.done = true
	return _node, nil
}
