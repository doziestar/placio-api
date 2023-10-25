// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/order"
	"placio-app/ent/orderitem"
	"placio-app/ent/placetable"
	"placio-app/ent/predicate"
	"placio-app/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetCreatedAt sets the "created_at" field.
func (ou *OrderUpdate) SetCreatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetCreatedAt(t)
	return ou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableCreatedAt(t *time.Time) *OrderUpdate {
	if t != nil {
		ou.SetCreatedAt(*t)
	}
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrderUpdate) SetUpdatedAt(t time.Time) *OrderUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetStatus sets the "status" field.
func (ou *OrderUpdate) SetStatus(o order.Status) *OrderUpdate {
	ou.mutation.SetStatus(o)
	return ou
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableStatus(o *order.Status) *OrderUpdate {
	if o != nil {
		ou.SetStatus(*o)
	}
	return ou
}

// SetTotalAmount sets the "total_amount" field.
func (ou *OrderUpdate) SetTotalAmount(f float64) *OrderUpdate {
	ou.mutation.ResetTotalAmount()
	ou.mutation.SetTotalAmount(f)
	return ou
}

// AddTotalAmount adds f to the "total_amount" field.
func (ou *OrderUpdate) AddTotalAmount(f float64) *OrderUpdate {
	ou.mutation.AddTotalAmount(f)
	return ou
}

// SetAdditionalInfo sets the "additional_info" field.
func (ou *OrderUpdate) SetAdditionalInfo(m map[string]interface{}) *OrderUpdate {
	ou.mutation.SetAdditionalInfo(m)
	return ou
}

// ClearAdditionalInfo clears the value of the "additional_info" field.
func (ou *OrderUpdate) ClearAdditionalInfo() *OrderUpdate {
	ou.mutation.ClearAdditionalInfo()
	return ou
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ou *OrderUpdate) SetUserID(id string) *OrderUpdate {
	ou.mutation.SetUserID(id)
	return ou
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ou *OrderUpdate) SetNillableUserID(id *string) *OrderUpdate {
	if id != nil {
		ou = ou.SetUserID(*id)
	}
	return ou
}

// SetUser sets the "user" edge to the User entity.
func (ou *OrderUpdate) SetUser(u *User) *OrderUpdate {
	return ou.SetUserID(u.ID)
}

// AddOrderItemIDs adds the "order_items" edge to the OrderItem entity by IDs.
func (ou *OrderUpdate) AddOrderItemIDs(ids ...string) *OrderUpdate {
	ou.mutation.AddOrderItemIDs(ids...)
	return ou
}

// AddOrderItems adds the "order_items" edges to the OrderItem entity.
func (ou *OrderUpdate) AddOrderItems(o ...*OrderItem) *OrderUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.AddOrderItemIDs(ids...)
}

// AddTableIDs adds the "table" edge to the PlaceTable entity by IDs.
func (ou *OrderUpdate) AddTableIDs(ids ...string) *OrderUpdate {
	ou.mutation.AddTableIDs(ids...)
	return ou
}

// AddTable adds the "table" edges to the PlaceTable entity.
func (ou *OrderUpdate) AddTable(p ...*PlaceTable) *OrderUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.AddTableIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ou *OrderUpdate) ClearUser() *OrderUpdate {
	ou.mutation.ClearUser()
	return ou
}

// ClearOrderItems clears all "order_items" edges to the OrderItem entity.
func (ou *OrderUpdate) ClearOrderItems() *OrderUpdate {
	ou.mutation.ClearOrderItems()
	return ou
}

// RemoveOrderItemIDs removes the "order_items" edge to OrderItem entities by IDs.
func (ou *OrderUpdate) RemoveOrderItemIDs(ids ...string) *OrderUpdate {
	ou.mutation.RemoveOrderItemIDs(ids...)
	return ou
}

// RemoveOrderItems removes "order_items" edges to OrderItem entities.
func (ou *OrderUpdate) RemoveOrderItems(o ...*OrderItem) *OrderUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.RemoveOrderItemIDs(ids...)
}

// ClearTable clears all "table" edges to the PlaceTable entity.
func (ou *OrderUpdate) ClearTable() *OrderUpdate {
	ou.mutation.ClearTable()
	return ou
}

// RemoveTableIDs removes the "table" edge to PlaceTable entities by IDs.
func (ou *OrderUpdate) RemoveTableIDs(ids ...string) *OrderUpdate {
	ou.mutation.RemoveTableIDs(ids...)
	return ou
}

// RemoveTable removes "table" edges to PlaceTable entities.
func (ou *OrderUpdate) RemoveTable(p ...*PlaceTable) *OrderUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ou.RemoveTableIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	ou.defaults()
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeString))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ou.mutation.TotalAmount(); ok {
		_spec.SetField(order.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := ou.mutation.AddedTotalAmount(); ok {
		_spec.AddField(order.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := ou.mutation.AdditionalInfo(); ok {
		_spec.SetField(order.FieldAdditionalInfo, field.TypeJSON, value)
	}
	if ou.mutation.AdditionalInfoCleared() {
		_spec.ClearField(order.FieldAdditionalInfo, field.TypeJSON)
	}
	if ou.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
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
	if ou.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: order.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedOrderItemsIDs(); len(nodes) > 0 && !ou.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: order.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.OrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: order.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.TableCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   order.TableTable,
			Columns: order.TablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placetable.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedTableIDs(); len(nodes) > 0 && !ou.mutation.TableCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   order.TableTable,
			Columns: order.TablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placetable.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.TableIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   order.TableTable,
			Columns: order.TablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placetable.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetCreatedAt sets the "created_at" field.
func (ouo *OrderUpdateOne) SetCreatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetCreatedAt(t)
	return ouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCreatedAt(t *time.Time) *OrderUpdateOne {
	if t != nil {
		ouo.SetCreatedAt(*t)
	}
	return ouo
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrderUpdateOne) SetUpdatedAt(t time.Time) *OrderUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetStatus sets the "status" field.
func (ouo *OrderUpdateOne) SetStatus(o order.Status) *OrderUpdateOne {
	ouo.mutation.SetStatus(o)
	return ouo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableStatus(o *order.Status) *OrderUpdateOne {
	if o != nil {
		ouo.SetStatus(*o)
	}
	return ouo
}

// SetTotalAmount sets the "total_amount" field.
func (ouo *OrderUpdateOne) SetTotalAmount(f float64) *OrderUpdateOne {
	ouo.mutation.ResetTotalAmount()
	ouo.mutation.SetTotalAmount(f)
	return ouo
}

// AddTotalAmount adds f to the "total_amount" field.
func (ouo *OrderUpdateOne) AddTotalAmount(f float64) *OrderUpdateOne {
	ouo.mutation.AddTotalAmount(f)
	return ouo
}

// SetAdditionalInfo sets the "additional_info" field.
func (ouo *OrderUpdateOne) SetAdditionalInfo(m map[string]interface{}) *OrderUpdateOne {
	ouo.mutation.SetAdditionalInfo(m)
	return ouo
}

// ClearAdditionalInfo clears the value of the "additional_info" field.
func (ouo *OrderUpdateOne) ClearAdditionalInfo() *OrderUpdateOne {
	ouo.mutation.ClearAdditionalInfo()
	return ouo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ouo *OrderUpdateOne) SetUserID(id string) *OrderUpdateOne {
	ouo.mutation.SetUserID(id)
	return ouo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableUserID(id *string) *OrderUpdateOne {
	if id != nil {
		ouo = ouo.SetUserID(*id)
	}
	return ouo
}

// SetUser sets the "user" edge to the User entity.
func (ouo *OrderUpdateOne) SetUser(u *User) *OrderUpdateOne {
	return ouo.SetUserID(u.ID)
}

// AddOrderItemIDs adds the "order_items" edge to the OrderItem entity by IDs.
func (ouo *OrderUpdateOne) AddOrderItemIDs(ids ...string) *OrderUpdateOne {
	ouo.mutation.AddOrderItemIDs(ids...)
	return ouo
}

// AddOrderItems adds the "order_items" edges to the OrderItem entity.
func (ouo *OrderUpdateOne) AddOrderItems(o ...*OrderItem) *OrderUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.AddOrderItemIDs(ids...)
}

// AddTableIDs adds the "table" edge to the PlaceTable entity by IDs.
func (ouo *OrderUpdateOne) AddTableIDs(ids ...string) *OrderUpdateOne {
	ouo.mutation.AddTableIDs(ids...)
	return ouo
}

// AddTable adds the "table" edges to the PlaceTable entity.
func (ouo *OrderUpdateOne) AddTable(p ...*PlaceTable) *OrderUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.AddTableIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ouo *OrderUpdateOne) ClearUser() *OrderUpdateOne {
	ouo.mutation.ClearUser()
	return ouo
}

// ClearOrderItems clears all "order_items" edges to the OrderItem entity.
func (ouo *OrderUpdateOne) ClearOrderItems() *OrderUpdateOne {
	ouo.mutation.ClearOrderItems()
	return ouo
}

// RemoveOrderItemIDs removes the "order_items" edge to OrderItem entities by IDs.
func (ouo *OrderUpdateOne) RemoveOrderItemIDs(ids ...string) *OrderUpdateOne {
	ouo.mutation.RemoveOrderItemIDs(ids...)
	return ouo
}

// RemoveOrderItems removes "order_items" edges to OrderItem entities.
func (ouo *OrderUpdateOne) RemoveOrderItems(o ...*OrderItem) *OrderUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.RemoveOrderItemIDs(ids...)
}

// ClearTable clears all "table" edges to the PlaceTable entity.
func (ouo *OrderUpdateOne) ClearTable() *OrderUpdateOne {
	ouo.mutation.ClearTable()
	return ouo
}

// RemoveTableIDs removes the "table" edge to PlaceTable entities by IDs.
func (ouo *OrderUpdateOne) RemoveTableIDs(ids ...string) *OrderUpdateOne {
	ouo.mutation.RemoveTableIDs(ids...)
	return ouo
}

// RemoveTable removes "table" edges to PlaceTable entities.
func (ouo *OrderUpdateOne) RemoveTable(p ...*PlaceTable) *OrderUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ouo.RemoveTableIDs(ids...)
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	ouo.defaults()
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		v := order.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeString))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ouo.mutation.TotalAmount(); ok {
		_spec.SetField(order.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := ouo.mutation.AddedTotalAmount(); ok {
		_spec.AddField(order.FieldTotalAmount, field.TypeFloat64, value)
	}
	if value, ok := ouo.mutation.AdditionalInfo(); ok {
		_spec.SetField(order.FieldAdditionalInfo, field.TypeJSON, value)
	}
	if ouo.mutation.AdditionalInfoCleared() {
		_spec.ClearField(order.FieldAdditionalInfo, field.TypeJSON)
	}
	if ouo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
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
	if ouo.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: order.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedOrderItemsIDs(); len(nodes) > 0 && !ouo.mutation.OrderItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: order.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.OrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: order.OrderItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.TableCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   order.TableTable,
			Columns: order.TablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placetable.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedTableIDs(); len(nodes) > 0 && !ouo.mutation.TableCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   order.TableTable,
			Columns: order.TablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placetable.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.TableIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   order.TableTable,
			Columns: order.TablePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placetable.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
