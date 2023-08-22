// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"placio-app/ent/business"
	"placio-app/ent/category"
	"placio-app/ent/inventorytype"
	"placio-app/ent/media"
	"placio-app/ent/place"
	"placio-app/ent/placeinventory"
	"placio-app/ent/placeinventoryattribute"
	"placio-app/ent/reservationblock"
	"placio-app/ent/transactionhistory"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaceInventoryCreate is the builder for creating a PlaceInventory entity.
type PlaceInventoryCreate struct {
	config
	mutation *PlaceInventoryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pic *PlaceInventoryCreate) SetName(s string) *PlaceInventoryCreate {
	pic.mutation.SetName(s)
	return pic
}

// SetPrice sets the "price" field.
func (pic *PlaceInventoryCreate) SetPrice(f float64) *PlaceInventoryCreate {
	pic.mutation.SetPrice(f)
	return pic
}

// SetStockQuantity sets the "stock_quantity" field.
func (pic *PlaceInventoryCreate) SetStockQuantity(i int) *PlaceInventoryCreate {
	pic.mutation.SetStockQuantity(i)
	return pic
}

// SetMinStockThreshold sets the "min_stock_threshold" field.
func (pic *PlaceInventoryCreate) SetMinStockThreshold(i int) *PlaceInventoryCreate {
	pic.mutation.SetMinStockThreshold(i)
	return pic
}

// SetNillableMinStockThreshold sets the "min_stock_threshold" field if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillableMinStockThreshold(i *int) *PlaceInventoryCreate {
	if i != nil {
		pic.SetMinStockThreshold(*i)
	}
	return pic
}

// SetSku sets the "sku" field.
func (pic *PlaceInventoryCreate) SetSku(s string) *PlaceInventoryCreate {
	pic.mutation.SetSku(s)
	return pic
}

// SetNillableSku sets the "sku" field if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillableSku(s *string) *PlaceInventoryCreate {
	if s != nil {
		pic.SetSku(*s)
	}
	return pic
}

// SetExpiryDate sets the "expiry_date" field.
func (pic *PlaceInventoryCreate) SetExpiryDate(t time.Time) *PlaceInventoryCreate {
	pic.mutation.SetExpiryDate(t)
	return pic
}

// SetNillableExpiryDate sets the "expiry_date" field if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillableExpiryDate(t *time.Time) *PlaceInventoryCreate {
	if t != nil {
		pic.SetExpiryDate(*t)
	}
	return pic
}

// SetSize sets the "size" field.
func (pic *PlaceInventoryCreate) SetSize(s string) *PlaceInventoryCreate {
	pic.mutation.SetSize(s)
	return pic
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillableSize(s *string) *PlaceInventoryCreate {
	if s != nil {
		pic.SetSize(*s)
	}
	return pic
}

// SetColor sets the "color" field.
func (pic *PlaceInventoryCreate) SetColor(s string) *PlaceInventoryCreate {
	pic.mutation.SetColor(s)
	return pic
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillableColor(s *string) *PlaceInventoryCreate {
	if s != nil {
		pic.SetColor(*s)
	}
	return pic
}

// SetBrand sets the "brand" field.
func (pic *PlaceInventoryCreate) SetBrand(s string) *PlaceInventoryCreate {
	pic.mutation.SetBrand(s)
	return pic
}

// SetNillableBrand sets the "brand" field if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillableBrand(s *string) *PlaceInventoryCreate {
	if s != nil {
		pic.SetBrand(*s)
	}
	return pic
}

// SetPurchaseDate sets the "purchase_date" field.
func (pic *PlaceInventoryCreate) SetPurchaseDate(t time.Time) *PlaceInventoryCreate {
	pic.mutation.SetPurchaseDate(t)
	return pic
}

// SetNillablePurchaseDate sets the "purchase_date" field if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillablePurchaseDate(t *time.Time) *PlaceInventoryCreate {
	if t != nil {
		pic.SetPurchaseDate(*t)
	}
	return pic
}

// SetLastUpdated sets the "last_updated" field.
func (pic *PlaceInventoryCreate) SetLastUpdated(t time.Time) *PlaceInventoryCreate {
	pic.mutation.SetLastUpdated(t)
	return pic
}

// SetNillableLastUpdated sets the "last_updated" field if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillableLastUpdated(t *time.Time) *PlaceInventoryCreate {
	if t != nil {
		pic.SetLastUpdated(*t)
	}
	return pic
}

// SetID sets the "id" field.
func (pic *PlaceInventoryCreate) SetID(s string) *PlaceInventoryCreate {
	pic.mutation.SetID(s)
	return pic
}

// SetPlaceID sets the "place" edge to the Place entity by ID.
func (pic *PlaceInventoryCreate) SetPlaceID(id string) *PlaceInventoryCreate {
	pic.mutation.SetPlaceID(id)
	return pic
}

// SetNillablePlaceID sets the "place" edge to the Place entity by ID if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillablePlaceID(id *string) *PlaceInventoryCreate {
	if id != nil {
		pic = pic.SetPlaceID(*id)
	}
	return pic
}

// SetPlace sets the "place" edge to the Place entity.
func (pic *PlaceInventoryCreate) SetPlace(p *Place) *PlaceInventoryCreate {
	return pic.SetPlaceID(p.ID)
}

// SetInventoryTypeID sets the "inventory_type" edge to the InventoryType entity by ID.
func (pic *PlaceInventoryCreate) SetInventoryTypeID(id string) *PlaceInventoryCreate {
	pic.mutation.SetInventoryTypeID(id)
	return pic
}

// SetNillableInventoryTypeID sets the "inventory_type" edge to the InventoryType entity by ID if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillableInventoryTypeID(id *string) *PlaceInventoryCreate {
	if id != nil {
		pic = pic.SetInventoryTypeID(*id)
	}
	return pic
}

// SetInventoryType sets the "inventory_type" edge to the InventoryType entity.
func (pic *PlaceInventoryCreate) SetInventoryType(i *InventoryType) *PlaceInventoryCreate {
	return pic.SetInventoryTypeID(i.ID)
}

// AddAttributeIDs adds the "attributes" edge to the PlaceInventoryAttribute entity by IDs.
func (pic *PlaceInventoryCreate) AddAttributeIDs(ids ...string) *PlaceInventoryCreate {
	pic.mutation.AddAttributeIDs(ids...)
	return pic
}

// AddAttributes adds the "attributes" edges to the PlaceInventoryAttribute entity.
func (pic *PlaceInventoryCreate) AddAttributes(p ...*PlaceInventoryAttribute) *PlaceInventoryCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pic.AddAttributeIDs(ids...)
}

// AddMediumIDs adds the "media" edge to the Media entity by IDs.
func (pic *PlaceInventoryCreate) AddMediumIDs(ids ...string) *PlaceInventoryCreate {
	pic.mutation.AddMediumIDs(ids...)
	return pic
}

// AddMedia adds the "media" edges to the Media entity.
func (pic *PlaceInventoryCreate) AddMedia(m ...*Media) *PlaceInventoryCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pic.AddMediumIDs(ids...)
}

// AddTransactionHistoryIDs adds the "transaction_histories" edge to the TransactionHistory entity by IDs.
func (pic *PlaceInventoryCreate) AddTransactionHistoryIDs(ids ...string) *PlaceInventoryCreate {
	pic.mutation.AddTransactionHistoryIDs(ids...)
	return pic
}

// AddTransactionHistories adds the "transaction_histories" edges to the TransactionHistory entity.
func (pic *PlaceInventoryCreate) AddTransactionHistories(t ...*TransactionHistory) *PlaceInventoryCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pic.AddTransactionHistoryIDs(ids...)
}

// AddReservationBlockIDs adds the "reservation_blocks" edge to the ReservationBlock entity by IDs.
func (pic *PlaceInventoryCreate) AddReservationBlockIDs(ids ...string) *PlaceInventoryCreate {
	pic.mutation.AddReservationBlockIDs(ids...)
	return pic
}

// AddReservationBlocks adds the "reservation_blocks" edges to the ReservationBlock entity.
func (pic *PlaceInventoryCreate) AddReservationBlocks(r ...*ReservationBlock) *PlaceInventoryCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pic.AddReservationBlockIDs(ids...)
}

// SetBusinessID sets the "business" edge to the Business entity by ID.
func (pic *PlaceInventoryCreate) SetBusinessID(id string) *PlaceInventoryCreate {
	pic.mutation.SetBusinessID(id)
	return pic
}

// SetNillableBusinessID sets the "business" edge to the Business entity by ID if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillableBusinessID(id *string) *PlaceInventoryCreate {
	if id != nil {
		pic = pic.SetBusinessID(*id)
	}
	return pic
}

// SetBusiness sets the "business" edge to the Business entity.
func (pic *PlaceInventoryCreate) SetBusiness(b *Business) *PlaceInventoryCreate {
	return pic.SetBusinessID(b.ID)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (pic *PlaceInventoryCreate) SetCategoryID(id string) *PlaceInventoryCreate {
	pic.mutation.SetCategoryID(id)
	return pic
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (pic *PlaceInventoryCreate) SetNillableCategoryID(id *string) *PlaceInventoryCreate {
	if id != nil {
		pic = pic.SetCategoryID(*id)
	}
	return pic
}

// SetCategory sets the "category" edge to the Category entity.
func (pic *PlaceInventoryCreate) SetCategory(c *Category) *PlaceInventoryCreate {
	return pic.SetCategoryID(c.ID)
}

// Mutation returns the PlaceInventoryMutation object of the builder.
func (pic *PlaceInventoryCreate) Mutation() *PlaceInventoryMutation {
	return pic.mutation
}

// Save creates the PlaceInventory in the database.
func (pic *PlaceInventoryCreate) Save(ctx context.Context) (*PlaceInventory, error) {
	pic.defaults()
	return withHooks(ctx, pic.sqlSave, pic.mutation, pic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pic *PlaceInventoryCreate) SaveX(ctx context.Context) *PlaceInventory {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *PlaceInventoryCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *PlaceInventoryCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *PlaceInventoryCreate) defaults() {
	if _, ok := pic.mutation.LastUpdated(); !ok {
		v := placeinventory.DefaultLastUpdated()
		pic.mutation.SetLastUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *PlaceInventoryCreate) check() error {
	if _, ok := pic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "PlaceInventory.name"`)}
	}
	if _, ok := pic.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "PlaceInventory.price"`)}
	}
	if _, ok := pic.mutation.StockQuantity(); !ok {
		return &ValidationError{Name: "stock_quantity", err: errors.New(`ent: missing required field "PlaceInventory.stock_quantity"`)}
	}
	if _, ok := pic.mutation.LastUpdated(); !ok {
		return &ValidationError{Name: "last_updated", err: errors.New(`ent: missing required field "PlaceInventory.last_updated"`)}
	}
	if v, ok := pic.mutation.ID(); ok {
		if err := placeinventory.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "PlaceInventory.id": %w`, err)}
		}
	}
	return nil
}

func (pic *PlaceInventoryCreate) sqlSave(ctx context.Context) (*PlaceInventory, error) {
	if err := pic.check(); err != nil {
		return nil, err
	}
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected PlaceInventory.ID type: %T", _spec.ID.Value)
		}
	}
	pic.mutation.id = &_node.ID
	pic.mutation.done = true
	return _node, nil
}

func (pic *PlaceInventoryCreate) createSpec() (*PlaceInventory, *sqlgraph.CreateSpec) {
	var (
		_node = &PlaceInventory{config: pic.config}
		_spec = sqlgraph.NewCreateSpec(placeinventory.Table, sqlgraph.NewFieldSpec(placeinventory.FieldID, field.TypeString))
	)
	if id, ok := pic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pic.mutation.Name(); ok {
		_spec.SetField(placeinventory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pic.mutation.Price(); ok {
		_spec.SetField(placeinventory.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := pic.mutation.StockQuantity(); ok {
		_spec.SetField(placeinventory.FieldStockQuantity, field.TypeInt, value)
		_node.StockQuantity = value
	}
	if value, ok := pic.mutation.MinStockThreshold(); ok {
		_spec.SetField(placeinventory.FieldMinStockThreshold, field.TypeInt, value)
		_node.MinStockThreshold = value
	}
	if value, ok := pic.mutation.Sku(); ok {
		_spec.SetField(placeinventory.FieldSku, field.TypeString, value)
		_node.Sku = value
	}
	if value, ok := pic.mutation.ExpiryDate(); ok {
		_spec.SetField(placeinventory.FieldExpiryDate, field.TypeTime, value)
		_node.ExpiryDate = value
	}
	if value, ok := pic.mutation.Size(); ok {
		_spec.SetField(placeinventory.FieldSize, field.TypeString, value)
		_node.Size = value
	}
	if value, ok := pic.mutation.Color(); ok {
		_spec.SetField(placeinventory.FieldColor, field.TypeString, value)
		_node.Color = value
	}
	if value, ok := pic.mutation.Brand(); ok {
		_spec.SetField(placeinventory.FieldBrand, field.TypeString, value)
		_node.Brand = value
	}
	if value, ok := pic.mutation.PurchaseDate(); ok {
		_spec.SetField(placeinventory.FieldPurchaseDate, field.TypeTime, value)
		_node.PurchaseDate = value
	}
	if value, ok := pic.mutation.LastUpdated(); ok {
		_spec.SetField(placeinventory.FieldLastUpdated, field.TypeTime, value)
		_node.LastUpdated = value
	}
	if nodes := pic.mutation.PlaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placeinventory.PlaceTable,
			Columns: []string{placeinventory.PlaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(place.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.place_inventories = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pic.mutation.InventoryTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placeinventory.InventoryTypeTable,
			Columns: []string{placeinventory.InventoryTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(inventorytype.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.inventory_type_place_inventories = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pic.mutation.AttributesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   placeinventory.AttributesTable,
			Columns: []string{placeinventory.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(placeinventoryattribute.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pic.mutation.MediaIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   placeinventory.MediaTable,
			Columns: placeinventory.MediaPrimaryKey,
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
	if nodes := pic.mutation.TransactionHistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   placeinventory.TransactionHistoriesTable,
			Columns: []string{placeinventory.TransactionHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transactionhistory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pic.mutation.ReservationBlocksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   placeinventory.ReservationBlocksTable,
			Columns: []string{placeinventory.ReservationBlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(reservationblock.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pic.mutation.BusinessIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placeinventory.BusinessTable,
			Columns: []string{placeinventory.BusinessColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(business.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.business_place_inventories = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pic.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   placeinventory.CategoryTable,
			Columns: []string{placeinventory.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.category_place_inventories = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PlaceInventoryCreateBulk is the builder for creating many PlaceInventory entities in bulk.
type PlaceInventoryCreateBulk struct {
	config
	builders []*PlaceInventoryCreate
}

// Save creates the PlaceInventory entities in the database.
func (picb *PlaceInventoryCreateBulk) Save(ctx context.Context) ([]*PlaceInventory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*PlaceInventory, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaceInventoryMutation)
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
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *PlaceInventoryCreateBulk) SaveX(ctx context.Context) []*PlaceInventory {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *PlaceInventoryCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *PlaceInventoryCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}
