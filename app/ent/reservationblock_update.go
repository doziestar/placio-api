// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio_api/placeinventory"
	"placio_api/predicate"
	"placio_api/reservationblock"
	"placio_api/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReservationBlockUpdate is the builder for updating ReservationBlock entities.
type ReservationBlockUpdate struct {
	config
	hooks    []Hook
	mutation *ReservationBlockMutation
}

// Where appends a list predicates to the ReservationBlockUpdate builder.
func (rbu *ReservationBlockUpdate) Where(ps ...predicate.ReservationBlock) *ReservationBlockUpdate {
	rbu.mutation.Where(ps...)
	return rbu
}

// SetStartTime sets the "start_time" field.
func (rbu *ReservationBlockUpdate) SetStartTime(t time.Time) *ReservationBlockUpdate {
	rbu.mutation.SetStartTime(t)
	return rbu
}

// SetEndTime sets the "end_time" field.
func (rbu *ReservationBlockUpdate) SetEndTime(t time.Time) *ReservationBlockUpdate {
	rbu.mutation.SetEndTime(t)
	return rbu
}

// SetStatus sets the "status" field.
func (rbu *ReservationBlockUpdate) SetStatus(r reservationblock.Status) *ReservationBlockUpdate {
	rbu.mutation.SetStatus(r)
	return rbu
}

// SetPlaceInventoryID sets the "place_inventory" edge to the PlaceInventory entity by ID.
func (rbu *ReservationBlockUpdate) SetPlaceInventoryID(id string) *ReservationBlockUpdate {
	rbu.mutation.SetPlaceInventoryID(id)
	return rbu
}

// SetNillablePlaceInventoryID sets the "place_inventory" edge to the PlaceInventory entity by ID if the given value is not nil.
func (rbu *ReservationBlockUpdate) SetNillablePlaceInventoryID(id *string) *ReservationBlockUpdate {
	if id != nil {
		rbu = rbu.SetPlaceInventoryID(*id)
	}
	return rbu
}

// SetPlaceInventory sets the "place_inventory" edge to the PlaceInventory entity.
func (rbu *ReservationBlockUpdate) SetPlaceInventory(p *PlaceInventory) *ReservationBlockUpdate {
	return rbu.SetPlaceInventoryID(p.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rbu *ReservationBlockUpdate) SetUserID(id string) *ReservationBlockUpdate {
	rbu.mutation.SetUserID(id)
	return rbu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (rbu *ReservationBlockUpdate) SetNillableUserID(id *string) *ReservationBlockUpdate {
	if id != nil {
		rbu = rbu.SetUserID(*id)
	}
	return rbu
}

// SetUser sets the "user" edge to the User entity.
func (rbu *ReservationBlockUpdate) SetUser(u *User) *ReservationBlockUpdate {
	return rbu.SetUserID(u.ID)
}

// Mutation returns the ReservationBlockMutation object of the builder.
func (rbu *ReservationBlockUpdate) Mutation() *ReservationBlockMutation {
	return rbu.mutation
}

// ClearPlaceInventory clears the "place_inventory" edge to the PlaceInventory entity.
func (rbu *ReservationBlockUpdate) ClearPlaceInventory() *ReservationBlockUpdate {
	rbu.mutation.ClearPlaceInventory()
	return rbu
}

// ClearUser clears the "user" edge to the User entity.
func (rbu *ReservationBlockUpdate) ClearUser() *ReservationBlockUpdate {
	rbu.mutation.ClearUser()
	return rbu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rbu *ReservationBlockUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rbu.sqlSave, rbu.mutation, rbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rbu *ReservationBlockUpdate) SaveX(ctx context.Context) int {
	affected, err := rbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rbu *ReservationBlockUpdate) Exec(ctx context.Context) error {
	_, err := rbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rbu *ReservationBlockUpdate) ExecX(ctx context.Context) {
	if err := rbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rbu *ReservationBlockUpdate) check() error {
	if v, ok := rbu.mutation.Status(); ok {
		if err := reservationblock.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`placio_api: validator failed for field "ReservationBlock.status": %w`, err)}
		}
	}
	return nil
}

func (rbu *ReservationBlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rbu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(reservationblock.Table, reservationblock.Columns, sqlgraph.NewFieldSpec(reservationblock.FieldID, field.TypeString))
	if ps := rbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rbu.mutation.StartTime(); ok {
		_spec.SetField(reservationblock.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := rbu.mutation.EndTime(); ok {
		_spec.SetField(reservationblock.FieldEndTime, field.TypeTime, value)
	}
	if value, ok := rbu.mutation.Status(); ok {
		_spec.SetField(reservationblock.FieldStatus, field.TypeEnum, value)
	}
	if rbu.mutation.PlaceInventoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservationblock.PlaceInventoryTable,
			Columns: []string{reservationblock.PlaceInventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rbu.mutation.PlaceInventoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservationblock.PlaceInventoryTable,
			Columns: []string{reservationblock.PlaceInventoryColumn},
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
	if rbu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservationblock.UserTable,
			Columns: []string{reservationblock.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rbu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservationblock.UserTable,
			Columns: []string{reservationblock.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reservationblock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rbu.mutation.done = true
	return n, nil
}

// ReservationBlockUpdateOne is the builder for updating a single ReservationBlock entity.
type ReservationBlockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReservationBlockMutation
}

// SetStartTime sets the "start_time" field.
func (rbuo *ReservationBlockUpdateOne) SetStartTime(t time.Time) *ReservationBlockUpdateOne {
	rbuo.mutation.SetStartTime(t)
	return rbuo
}

// SetEndTime sets the "end_time" field.
func (rbuo *ReservationBlockUpdateOne) SetEndTime(t time.Time) *ReservationBlockUpdateOne {
	rbuo.mutation.SetEndTime(t)
	return rbuo
}

// SetStatus sets the "status" field.
func (rbuo *ReservationBlockUpdateOne) SetStatus(r reservationblock.Status) *ReservationBlockUpdateOne {
	rbuo.mutation.SetStatus(r)
	return rbuo
}

// SetPlaceInventoryID sets the "place_inventory" edge to the PlaceInventory entity by ID.
func (rbuo *ReservationBlockUpdateOne) SetPlaceInventoryID(id string) *ReservationBlockUpdateOne {
	rbuo.mutation.SetPlaceInventoryID(id)
	return rbuo
}

// SetNillablePlaceInventoryID sets the "place_inventory" edge to the PlaceInventory entity by ID if the given value is not nil.
func (rbuo *ReservationBlockUpdateOne) SetNillablePlaceInventoryID(id *string) *ReservationBlockUpdateOne {
	if id != nil {
		rbuo = rbuo.SetPlaceInventoryID(*id)
	}
	return rbuo
}

// SetPlaceInventory sets the "place_inventory" edge to the PlaceInventory entity.
func (rbuo *ReservationBlockUpdateOne) SetPlaceInventory(p *PlaceInventory) *ReservationBlockUpdateOne {
	return rbuo.SetPlaceInventoryID(p.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rbuo *ReservationBlockUpdateOne) SetUserID(id string) *ReservationBlockUpdateOne {
	rbuo.mutation.SetUserID(id)
	return rbuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (rbuo *ReservationBlockUpdateOne) SetNillableUserID(id *string) *ReservationBlockUpdateOne {
	if id != nil {
		rbuo = rbuo.SetUserID(*id)
	}
	return rbuo
}

// SetUser sets the "user" edge to the User entity.
func (rbuo *ReservationBlockUpdateOne) SetUser(u *User) *ReservationBlockUpdateOne {
	return rbuo.SetUserID(u.ID)
}

// Mutation returns the ReservationBlockMutation object of the builder.
func (rbuo *ReservationBlockUpdateOne) Mutation() *ReservationBlockMutation {
	return rbuo.mutation
}

// ClearPlaceInventory clears the "place_inventory" edge to the PlaceInventory entity.
func (rbuo *ReservationBlockUpdateOne) ClearPlaceInventory() *ReservationBlockUpdateOne {
	rbuo.mutation.ClearPlaceInventory()
	return rbuo
}

// ClearUser clears the "user" edge to the User entity.
func (rbuo *ReservationBlockUpdateOne) ClearUser() *ReservationBlockUpdateOne {
	rbuo.mutation.ClearUser()
	return rbuo
}

// Where appends a list predicates to the ReservationBlockUpdate builder.
func (rbuo *ReservationBlockUpdateOne) Where(ps ...predicate.ReservationBlock) *ReservationBlockUpdateOne {
	rbuo.mutation.Where(ps...)
	return rbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rbuo *ReservationBlockUpdateOne) Select(field string, fields ...string) *ReservationBlockUpdateOne {
	rbuo.fields = append([]string{field}, fields...)
	return rbuo
}

// Save executes the query and returns the updated ReservationBlock entity.
func (rbuo *ReservationBlockUpdateOne) Save(ctx context.Context) (*ReservationBlock, error) {
	return withHooks(ctx, rbuo.sqlSave, rbuo.mutation, rbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rbuo *ReservationBlockUpdateOne) SaveX(ctx context.Context) *ReservationBlock {
	node, err := rbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rbuo *ReservationBlockUpdateOne) Exec(ctx context.Context) error {
	_, err := rbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rbuo *ReservationBlockUpdateOne) ExecX(ctx context.Context) {
	if err := rbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rbuo *ReservationBlockUpdateOne) check() error {
	if v, ok := rbuo.mutation.Status(); ok {
		if err := reservationblock.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`placio_api: validator failed for field "ReservationBlock.status": %w`, err)}
		}
	}
	return nil
}

func (rbuo *ReservationBlockUpdateOne) sqlSave(ctx context.Context) (_node *ReservationBlock, err error) {
	if err := rbuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(reservationblock.Table, reservationblock.Columns, sqlgraph.NewFieldSpec(reservationblock.FieldID, field.TypeString))
	id, ok := rbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`placio_api: missing "ReservationBlock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reservationblock.FieldID)
		for _, f := range fields {
			if !reservationblock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("placio_api: invalid field %q for query", f)}
			}
			if f != reservationblock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rbuo.mutation.StartTime(); ok {
		_spec.SetField(reservationblock.FieldStartTime, field.TypeTime, value)
	}
	if value, ok := rbuo.mutation.EndTime(); ok {
		_spec.SetField(reservationblock.FieldEndTime, field.TypeTime, value)
	}
	if value, ok := rbuo.mutation.Status(); ok {
		_spec.SetField(reservationblock.FieldStatus, field.TypeEnum, value)
	}
	if rbuo.mutation.PlaceInventoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservationblock.PlaceInventoryTable,
			Columns: []string{reservationblock.PlaceInventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rbuo.mutation.PlaceInventoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservationblock.PlaceInventoryTable,
			Columns: []string{reservationblock.PlaceInventoryColumn},
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
	if rbuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservationblock.UserTable,
			Columns: []string{reservationblock.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rbuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservationblock.UserTable,
			Columns: []string{reservationblock.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ReservationBlock{config: rbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reservationblock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rbuo.mutation.done = true
	return _node, nil
}
