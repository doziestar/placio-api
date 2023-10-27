// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"placio-app/ent/orderitem"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// OrderItem is the model entity for the OrderItem schema.
type OrderItem struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// PricePerItem holds the value of the "price_per_item" field.
	PricePerItem float64 `json:"price_per_item,omitempty"`
	// TotalPrice holds the value of the "total_price" field.
	TotalPrice float64 `json:"total_price,omitempty"`
	// Options holds the value of the "options" field.
	Options []string `json:"options,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderItemQuery when eager-loading is set.
	Edges        OrderItemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderItemEdges holds the relations/edges for other nodes in the graph.
type OrderItemEdges struct {
	// Order holds the value of the order edge.
	Order []*Order `json:"order,omitempty"`
	// MenuItem holds the value of the menu_item edge.
	MenuItem []*MenuItem `json:"menu_item,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading.
func (e OrderItemEdges) OrderOrErr() ([]*Order, error) {
	if e.loadedTypes[0] {
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// MenuItemOrErr returns the MenuItem value or an error if the edge
// was not loaded in eager-loading.
func (e OrderItemEdges) MenuItemOrErr() ([]*MenuItem, error) {
	if e.loadedTypes[1] {
		return e.MenuItem, nil
	}
	return nil, &NotLoadedError{edge: "menu_item"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldOptions:
			values[i] = new([]byte)
		case orderitem.FieldPricePerItem, orderitem.FieldTotalPrice:
			values[i] = new(sql.NullFloat64)
		case orderitem.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case orderitem.FieldID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderItem fields.
func (oi *OrderItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				oi.ID = value.String
			}
		case orderitem.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				oi.Quantity = int(value.Int64)
			}
		case orderitem.FieldPricePerItem:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price_per_item", values[i])
			} else if value.Valid {
				oi.PricePerItem = value.Float64
			}
		case orderitem.FieldTotalPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_price", values[i])
			} else if value.Valid {
				oi.TotalPrice = value.Float64
			}
		case orderitem.FieldOptions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field options", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &oi.Options); err != nil {
					return fmt.Errorf("unmarshal field options: %w", err)
				}
			}
		default:
			oi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderItem.
// This includes values selected through modifiers, order, etc.
func (oi *OrderItem) Value(name string) (ent.Value, error) {
	return oi.selectValues.Get(name)
}

// QueryOrder queries the "order" edge of the OrderItem entity.
func (oi *OrderItem) QueryOrder() *OrderQuery {
	return NewOrderItemClient(oi.config).QueryOrder(oi)
}

// QueryMenuItem queries the "menu_item" edge of the OrderItem entity.
func (oi *OrderItem) QueryMenuItem() *MenuItemQuery {
	return NewOrderItemClient(oi.config).QueryMenuItem(oi)
}

// Update returns a builder for updating this OrderItem.
// Note that you need to call OrderItem.Unwrap() before calling this method if this OrderItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrderItem) Update() *OrderItemUpdateOne {
	return NewOrderItemClient(oi.config).UpdateOne(oi)
}

// Unwrap unwraps the OrderItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrderItem) Unwrap() *OrderItem {
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderItem is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderItem) String() string {
	var builder strings.Builder
	builder.WriteString("OrderItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oi.ID))
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", oi.Quantity))
	builder.WriteString(", ")
	builder.WriteString("price_per_item=")
	builder.WriteString(fmt.Sprintf("%v", oi.PricePerItem))
	builder.WriteString(", ")
	builder.WriteString("total_price=")
	builder.WriteString(fmt.Sprintf("%v", oi.TotalPrice))
	builder.WriteString(", ")
	builder.WriteString("options=")
	builder.WriteString(fmt.Sprintf("%v", oi.Options))
	builder.WriteByte(')')
	return builder.String()
}

// OrderItems is a parsable slice of OrderItem.
type OrderItems []*OrderItem