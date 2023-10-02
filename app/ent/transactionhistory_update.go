// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/placeinventory"
	"placio-app/ent/predicate"
	"placio-app/ent/transactionhistory"
	"placio-app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransactionHistoryUpdate is the builder for updating TransactionHistory entities.
type TransactionHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionHistoryMutation
}

// Where appends a list predicates to the TransactionHistoryUpdate builder.
func (thu *TransactionHistoryUpdate) Where(ps ...predicate.TransactionHistory) *TransactionHistoryUpdate {
	thu.mutation.Where(ps...)
	return thu
}

// SetTransactionType sets the "transaction_type" field.
func (thu *TransactionHistoryUpdate) SetTransactionType(tt transactionhistory.TransactionType) *TransactionHistoryUpdate {
	thu.mutation.SetTransactionType(tt)
	return thu
}

// SetQuantity sets the "quantity" field.
func (thu *TransactionHistoryUpdate) SetQuantity(i int) *TransactionHistoryUpdate {
	thu.mutation.ResetQuantity()
	thu.mutation.SetQuantity(i)
	return thu
}

// AddQuantity adds i to the "quantity" field.
func (thu *TransactionHistoryUpdate) AddQuantity(i int) *TransactionHistoryUpdate {
	thu.mutation.AddQuantity(i)
	return thu
}

// SetDate sets the "date" field.
func (thu *TransactionHistoryUpdate) SetDate(t time.Time) *TransactionHistoryUpdate {
	thu.mutation.SetDate(t)
	return thu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (thu *TransactionHistoryUpdate) SetNillableDate(t *time.Time) *TransactionHistoryUpdate {
	if t != nil {
		thu.SetDate(*t)
	}
	return thu
}

// SetPlaceInventoryID sets the "place_inventory" edge to the PlaceInventory entity by ID.
func (thu *TransactionHistoryUpdate) SetPlaceInventoryID(id string) *TransactionHistoryUpdate {
	thu.mutation.SetPlaceInventoryID(id)
	return thu
}

// SetNillablePlaceInventoryID sets the "place_inventory" edge to the PlaceInventory entity by ID if the given value is not nil.
func (thu *TransactionHistoryUpdate) SetNillablePlaceInventoryID(id *string) *TransactionHistoryUpdate {
	if id != nil {
		thu = thu.SetPlaceInventoryID(*id)
	}
	return thu
}

// SetPlaceInventory sets the "place_inventory" edge to the PlaceInventory entity.
func (thu *TransactionHistoryUpdate) SetPlaceInventory(p *PlaceInventory) *TransactionHistoryUpdate {
	return thu.SetPlaceInventoryID(p.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (thu *TransactionHistoryUpdate) SetUserID(id string) *TransactionHistoryUpdate {
	thu.mutation.SetUserID(id)
	return thu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (thu *TransactionHistoryUpdate) SetNillableUserID(id *string) *TransactionHistoryUpdate {
	if id != nil {
		thu = thu.SetUserID(*id)
	}
	return thu
}

// SetUser sets the "user" edge to the User entity.
func (thu *TransactionHistoryUpdate) SetUser(u *User) *TransactionHistoryUpdate {
	return thu.SetUserID(u.ID)
}

// Mutation returns the TransactionHistoryMutation object of the builder.
func (thu *TransactionHistoryUpdate) Mutation() *TransactionHistoryMutation {
	return thu.mutation
}

// ClearPlaceInventory clears the "place_inventory" edge to the PlaceInventory entity.
func (thu *TransactionHistoryUpdate) ClearPlaceInventory() *TransactionHistoryUpdate {
	thu.mutation.ClearPlaceInventory()
	return thu
}

// ClearUser clears the "user" edge to the User entity.
func (thu *TransactionHistoryUpdate) ClearUser() *TransactionHistoryUpdate {
	thu.mutation.ClearUser()
	return thu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (thu *TransactionHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, thu.sqlSave, thu.mutation, thu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thu *TransactionHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := thu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (thu *TransactionHistoryUpdate) Exec(ctx context.Context) error {
	_, err := thu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thu *TransactionHistoryUpdate) ExecX(ctx context.Context) {
	if err := thu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thu *TransactionHistoryUpdate) check() error {
	if v, ok := thu.mutation.TransactionType(); ok {
		if err := transactionhistory.TransactionTypeValidator(v); err != nil {
			return &ValidationError{Name: "transaction_type", err: fmt.Errorf(`ent: validator failed for field "TransactionHistory.transaction_type": %w`, err)}
		}
	}
	return nil
}

func (thu *TransactionHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := thu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(transactionhistory.Table, transactionhistory.Columns, sqlgraph.NewFieldSpec(transactionhistory.FieldID, field.TypeString))
	if ps := thu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := thu.mutation.TransactionType(); ok {
		_spec.SetField(transactionhistory.FieldTransactionType, field.TypeEnum, value)
	}
	if value, ok := thu.mutation.Quantity(); ok {
		_spec.SetField(transactionhistory.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := thu.mutation.AddedQuantity(); ok {
		_spec.AddField(transactionhistory.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := thu.mutation.Date(); ok {
		_spec.SetField(transactionhistory.FieldDate, field.TypeTime, value)
	}
	if thu.mutation.PlaceInventoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionhistory.PlaceInventoryTable,
			Columns: []string{transactionhistory.PlaceInventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := thu.mutation.PlaceInventoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionhistory.PlaceInventoryTable,
			Columns: []string{transactionhistory.PlaceInventoryColumn},
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
	if thu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionhistory.UserTable,
			Columns: []string{transactionhistory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := thu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionhistory.UserTable,
			Columns: []string{transactionhistory.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, thu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactionhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	thu.mutation.done = true
	return n, nil
}

// TransactionHistoryUpdateOne is the builder for updating a single TransactionHistory entity.
type TransactionHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionHistoryMutation
}

// SetTransactionType sets the "transaction_type" field.
func (thuo *TransactionHistoryUpdateOne) SetTransactionType(tt transactionhistory.TransactionType) *TransactionHistoryUpdateOne {
	thuo.mutation.SetTransactionType(tt)
	return thuo
}

// SetQuantity sets the "quantity" field.
func (thuo *TransactionHistoryUpdateOne) SetQuantity(i int) *TransactionHistoryUpdateOne {
	thuo.mutation.ResetQuantity()
	thuo.mutation.SetQuantity(i)
	return thuo
}

// AddQuantity adds i to the "quantity" field.
func (thuo *TransactionHistoryUpdateOne) AddQuantity(i int) *TransactionHistoryUpdateOne {
	thuo.mutation.AddQuantity(i)
	return thuo
}

// SetDate sets the "date" field.
func (thuo *TransactionHistoryUpdateOne) SetDate(t time.Time) *TransactionHistoryUpdateOne {
	thuo.mutation.SetDate(t)
	return thuo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (thuo *TransactionHistoryUpdateOne) SetNillableDate(t *time.Time) *TransactionHistoryUpdateOne {
	if t != nil {
		thuo.SetDate(*t)
	}
	return thuo
}

// SetPlaceInventoryID sets the "place_inventory" edge to the PlaceInventory entity by ID.
func (thuo *TransactionHistoryUpdateOne) SetPlaceInventoryID(id string) *TransactionHistoryUpdateOne {
	thuo.mutation.SetPlaceInventoryID(id)
	return thuo
}

// SetNillablePlaceInventoryID sets the "place_inventory" edge to the PlaceInventory entity by ID if the given value is not nil.
func (thuo *TransactionHistoryUpdateOne) SetNillablePlaceInventoryID(id *string) *TransactionHistoryUpdateOne {
	if id != nil {
		thuo = thuo.SetPlaceInventoryID(*id)
	}
	return thuo
}

// SetPlaceInventory sets the "place_inventory" edge to the PlaceInventory entity.
func (thuo *TransactionHistoryUpdateOne) SetPlaceInventory(p *PlaceInventory) *TransactionHistoryUpdateOne {
	return thuo.SetPlaceInventoryID(p.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (thuo *TransactionHistoryUpdateOne) SetUserID(id string) *TransactionHistoryUpdateOne {
	thuo.mutation.SetUserID(id)
	return thuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (thuo *TransactionHistoryUpdateOne) SetNillableUserID(id *string) *TransactionHistoryUpdateOne {
	if id != nil {
		thuo = thuo.SetUserID(*id)
	}
	return thuo
}

// SetUser sets the "user" edge to the User entity.
func (thuo *TransactionHistoryUpdateOne) SetUser(u *User) *TransactionHistoryUpdateOne {
	return thuo.SetUserID(u.ID)
}

// Mutation returns the TransactionHistoryMutation object of the builder.
func (thuo *TransactionHistoryUpdateOne) Mutation() *TransactionHistoryMutation {
	return thuo.mutation
}

// ClearPlaceInventory clears the "place_inventory" edge to the PlaceInventory entity.
func (thuo *TransactionHistoryUpdateOne) ClearPlaceInventory() *TransactionHistoryUpdateOne {
	thuo.mutation.ClearPlaceInventory()
	return thuo
}

// ClearUser clears the "user" edge to the User entity.
func (thuo *TransactionHistoryUpdateOne) ClearUser() *TransactionHistoryUpdateOne {
	thuo.mutation.ClearUser()
	return thuo
}

// Where appends a list predicates to the TransactionHistoryUpdate builder.
func (thuo *TransactionHistoryUpdateOne) Where(ps ...predicate.TransactionHistory) *TransactionHistoryUpdateOne {
	thuo.mutation.Where(ps...)
	return thuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (thuo *TransactionHistoryUpdateOne) Select(field string, fields ...string) *TransactionHistoryUpdateOne {
	thuo.fields = append([]string{field}, fields...)
	return thuo
}

// Save executes the query and returns the updated TransactionHistory entity.
func (thuo *TransactionHistoryUpdateOne) Save(ctx context.Context) (*TransactionHistory, error) {
	return withHooks(ctx, thuo.sqlSave, thuo.mutation, thuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thuo *TransactionHistoryUpdateOne) SaveX(ctx context.Context) *TransactionHistory {
	node, err := thuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (thuo *TransactionHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := thuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thuo *TransactionHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := thuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thuo *TransactionHistoryUpdateOne) check() error {
	if v, ok := thuo.mutation.TransactionType(); ok {
		if err := transactionhistory.TransactionTypeValidator(v); err != nil {
			return &ValidationError{Name: "transaction_type", err: fmt.Errorf(`ent: validator failed for field "TransactionHistory.transaction_type": %w`, err)}
		}
	}
	return nil
}

func (thuo *TransactionHistoryUpdateOne) sqlSave(ctx context.Context) (_node *TransactionHistory, err error) {
	if err := thuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(transactionhistory.Table, transactionhistory.Columns, sqlgraph.NewFieldSpec(transactionhistory.FieldID, field.TypeString))
	id, ok := thuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TransactionHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := thuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transactionhistory.FieldID)
		for _, f := range fields {
			if !transactionhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transactionhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := thuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := thuo.mutation.TransactionType(); ok {
		_spec.SetField(transactionhistory.FieldTransactionType, field.TypeEnum, value)
	}
	if value, ok := thuo.mutation.Quantity(); ok {
		_spec.SetField(transactionhistory.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := thuo.mutation.AddedQuantity(); ok {
		_spec.AddField(transactionhistory.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := thuo.mutation.Date(); ok {
		_spec.SetField(transactionhistory.FieldDate, field.TypeTime, value)
	}
	if thuo.mutation.PlaceInventoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionhistory.PlaceInventoryTable,
			Columns: []string{transactionhistory.PlaceInventoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := thuo.mutation.PlaceInventoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionhistory.PlaceInventoryTable,
			Columns: []string{transactionhistory.PlaceInventoryColumn},
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
	if thuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionhistory.UserTable,
			Columns: []string{transactionhistory.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := thuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionhistory.UserTable,
			Columns: []string{transactionhistory.UserColumn},
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
	_node = &TransactionHistory{config: thuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, thuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactionhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	thuo.mutation.done = true
	return _node, nil
}