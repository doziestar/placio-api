




// Code generated by ent, DO NOT EDIT.



package ent



	
import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
		"placio-app/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
			"database/sql/driver"
			"entgo.io/ent/dialect/sql"
			"entgo.io/ent/dialect/sql/sqlgraph"
			"entgo.io/ent/dialect/sql/sqljson"
			"entgo.io/ent/schema/field"
			 "placio-app/ent/order"
			 "placio-app/ent/user"
			 "placio-app/ent/orderitem"
			 "placio-app/ent/placetable"

)







// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks []Hook
}


	





	


	
	
	// SetCreatedAt sets the "created_at" field.
	func (oc *OrderCreate) SetCreatedAt(t time.Time) *OrderCreate {
		oc.mutation.SetCreatedAt(t)
		return oc
	}

	
	
	
	
	
	
		// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
		func (oc *OrderCreate) SetNillableCreatedAt(t *time.Time) *OrderCreate {
			if t != nil {
				oc.SetCreatedAt(*t)
			}
			return oc
		}
	

	

	

	

	
	
	// SetUpdatedAt sets the "updated_at" field.
	func (oc *OrderCreate) SetUpdatedAt(t time.Time) *OrderCreate {
		oc.mutation.SetUpdatedAt(t)
		return oc
	}

	
	
	
	
	
	
		// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
		func (oc *OrderCreate) SetNillableUpdatedAt(t *time.Time) *OrderCreate {
			if t != nil {
				oc.SetUpdatedAt(*t)
			}
			return oc
		}
	

	

	

	

	
	
	// SetStatus sets the "status" field.
	func (oc *OrderCreate) SetStatus(o order.Status) *OrderCreate {
		oc.mutation.SetStatus(o)
		return oc
	}

	
	
	
	
	
	
		// SetNillableStatus sets the "status" field if the given value is not nil.
		func (oc *OrderCreate) SetNillableStatus(o *order.Status) *OrderCreate {
			if o != nil {
				oc.SetStatus(*o)
			}
			return oc
		}
	

	

	

	

	
	
	// SetTotalAmount sets the "total_amount" field.
	func (oc *OrderCreate) SetTotalAmount(f float64) *OrderCreate {
		oc.mutation.SetTotalAmount(f)
		return oc
	}

	
	
	
	
	
	

	

	

	

	
	
	// SetAdditionalInfo sets the "additional_info" field.
	func (oc *OrderCreate) SetAdditionalInfo(m map[string]interface {}) *OrderCreate {
		oc.mutation.SetAdditionalInfo(m)
		return oc
	}

	
	
	
	
	
	

	

	

	

	
	
	// SetDeletedAt sets the "deleted_at" field.
	func (oc *OrderCreate) SetDeletedAt(s string) *OrderCreate {
		oc.mutation.SetDeletedAt(s)
		return oc
	}

	
	
	
	
	
	
		// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
		func (oc *OrderCreate) SetNillableDeletedAt(s *string) *OrderCreate {
			if s != nil {
				oc.SetDeletedAt(*s)
			}
			return oc
		}
	

	

	

	

	
	
	// SetIsDeleted sets the "is_deleted" field.
	func (oc *OrderCreate) SetIsDeleted(b bool) *OrderCreate {
		oc.mutation.SetIsDeleted(b)
		return oc
	}

	
	
	
	
	
	
		// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
		func (oc *OrderCreate) SetNillableIsDeleted(b *bool) *OrderCreate {
			if b != nil {
				oc.SetIsDeleted(*b)
			}
			return oc
		}
	

	

	

	

	
	
	// SetID sets the "id" field.
	func (oc *OrderCreate) SetID(s string) *OrderCreate {
		oc.mutation.SetID(s)
		return oc
	}

	
	
	
	
	
	

	

	

	



	
	
	
	
	
		// SetUserID sets the "user" edge to the User entity by ID.
		func (oc *OrderCreate) SetUserID(id string) *OrderCreate {
			oc.mutation.SetUserID(id)
			return oc
		}
	
	
		
		// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
		func (oc *OrderCreate) SetNillableUserID(id *string) *OrderCreate {
			if id != nil {
				oc = oc.SetUserID(*id)
			}
			return oc
		}
	
	
	
	
	// SetUser sets the "user" edge to the User entity.
	func (oc *OrderCreate) SetUser(u *User) *OrderCreate {
		return oc.SetUserID(u.ID)
	}

	
	
	
	
	
		// AddOrderItemIDs adds the "order_items" edge to the OrderItem entity by IDs.
		func (oc *OrderCreate) AddOrderItemIDs(ids ... string) *OrderCreate {
			oc.mutation.AddOrderItemIDs(ids ...)
			return oc
		}
	
	
	
	
	
	// AddOrderItems adds the "order_items" edges to the OrderItem entity.
	func (oc *OrderCreate) AddOrderItems(o ...*OrderItem) *OrderCreate {
		ids := make([]string, len(o))
			for i := range o {
				ids[i] = o[i].ID
			}
			return oc.AddOrderItemIDs(ids...)
	}

	
	
	
	
	
		// AddTableIDs adds the "table" edge to the PlaceTable entity by IDs.
		func (oc *OrderCreate) AddTableIDs(ids ... string) *OrderCreate {
			oc.mutation.AddTableIDs(ids ...)
			return oc
		}
	
	
	
	
	
	// AddTable adds the "table" edges to the PlaceTable entity.
	func (oc *OrderCreate) AddTable(p ...*PlaceTable) *OrderCreate {
		ids := make([]string, len(p))
			for i := range p {
				ids[i] = p[i].ID
			}
			return oc.AddTableIDs(ids...)
	}


// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}




// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
			oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

	// defaults sets the default values of the builder before save.
	func (oc *OrderCreate) defaults() {
				if _, ok := oc.mutation.CreatedAt(); !ok {
					v := order.DefaultCreatedAt()
					oc.mutation.SetCreatedAt(v)
				}
				if _, ok := oc.mutation.UpdatedAt(); !ok {
					v := order.DefaultUpdatedAt()
					oc.mutation.SetUpdatedAt(v)
				}
				if _, ok := oc.mutation.Status(); !ok {
					v := order.DefaultStatus
					oc.mutation.SetStatus(v)
				}
				if _, ok := oc.mutation.IsDeleted(); !ok {
					v := order.DefaultIsDeleted
					oc.mutation.SetIsDeleted(v)
				}
	}


// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
					if _, ok := oc.mutation.CreatedAt(); !ok {
						return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Order.created_at"`)}
					}
					if _, ok := oc.mutation.UpdatedAt(); !ok {
						return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Order.updated_at"`)}
					}
					if _, ok := oc.mutation.Status(); !ok {
						return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Order.status"`)}
					}
			if v, ok := oc.mutation.Status(); ok {
				if err := order.StatusValidator(v); err != nil {
					return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
				}
			}
					if _, ok := oc.mutation.TotalAmount(); !ok {
						return &ValidationError{Name: "total_amount", err: errors.New(`ent: missing required field "Order.total_amount"`)}
					}
					if _, ok := oc.mutation.IsDeleted(); !ok {
						return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "Order.is_deleted"`)}
					}
			if v, ok := oc.mutation.ID(); ok {
				if err := order.IDValidator(v); err != nil {
					return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Order.id": %w`, err)}
				}
			}
	return nil
}


	
	




func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec  := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
		if _spec.ID.Value != nil {
				if id, ok := _spec.ID.Value.(string); ok {
					_node.ID = id
				} else {
					return nil, fmt.Errorf("unexpected Order.ID type: %T", _spec.ID.Value)
				}
		}
		oc.mutation.id = &_node.ID
		oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeString))
	)
		if id, ok := oc.mutation.ID(); ok {
			_node.ID = id
			_spec.ID.Value = id
		}
		if value, ok := oc.mutation.CreatedAt(); ok {
				_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
			_node.CreatedAt = value
		}
		if value, ok := oc.mutation.UpdatedAt(); ok {
				_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
			_node.UpdatedAt = value
		}
		if value, ok := oc.mutation.Status(); ok {
				_spec.SetField(order.FieldStatus, field.TypeEnum, value)
			_node.Status = value
		}
		if value, ok := oc.mutation.TotalAmount(); ok {
				_spec.SetField(order.FieldTotalAmount, field.TypeFloat64, value)
			_node.TotalAmount = value
		}
		if value, ok := oc.mutation.AdditionalInfo(); ok {
				_spec.SetField(order.FieldAdditionalInfo, field.TypeJSON, value)
			_node.AdditionalInfo = value
		}
		if value, ok := oc.mutation.DeletedAt(); ok {
				_spec.SetField(order.FieldDeletedAt, field.TypeString, value)
			_node.DeletedAt = value
		}
		if value, ok := oc.mutation.IsDeleted(); ok {
				_spec.SetField(order.FieldIsDeleted, field.TypeBool, value)
			_node.IsDeleted = value
		}
		if nodes := oc.mutation.UserIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: order.UserTable,
		Columns: []string{ order.UserColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
				_node.user_orders = &nodes[0]
			_spec.Edges = append(_spec.Edges, edge)
		}
		if nodes := oc.mutation.OrderItemsIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2M,
		Inverse: false,
		Table: order.OrderItemsTable,
		Columns: order.OrderItemsPrimaryKey,
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges = append(_spec.Edges, edge)
		}
		if nodes := oc.mutation.TableIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2M,
		Inverse: true,
		Table: order.TableTable,
		Columns: order.TablePrimaryKey,
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: sqlgraph.NewFieldSpec(placetable.FieldID, field.TypeString),
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges = append(_spec.Edges, edge)
		}
	return _node, _spec
}
	








// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	err error
	builders []*OrderCreate
}




	
		



// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
				builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
	


	

