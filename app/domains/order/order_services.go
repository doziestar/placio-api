package order

import (
	"context"
	"fmt"
	"placio-app/ent"
	"placio-app/ent/order"
	"placio-app/ent/placetable"
	"placio-app/ent/user"

	"github.com/google/uuid"
)

type OrderServicesImpl struct {
	client *ent.Client
}

type OrderServices interface {
	CreateOrder(ctx context.Context, orderDto *ent.Order, tableID, userid string, orderItems OrderWithItemsDTO) (*ent.Order, error)
	UpdateOrder(ctx context.Context, orderID string, orderDto OrderWithItemsDTO) (*ent.Order, error)
	DeleteOrder(ctx context.Context, orderID string) error
	GetOrder(ctx context.Context, orderID string) (*ent.Order, error)
	GetOrders(ctx context.Context, limit, offset int) ([]*ent.Order, error)
	GetOrdersByUserID(ctx context.Context, userID string, limit, offset int) ([]*ent.Order, error)
	GetOrdersByTableID(ctx context.Context, tableID string, limit, offset int) ([]*ent.Order, error)
	//GetOrdersByStatus(ctx context.Context, status string, limit, offset int) ([]*ent.Order, error)
	GetOrdersByTableIDAndStatus(ctx context.Context, tableID, status string, limit, offset int) ([]*ent.Order, error)
	GetOrdersByUserIDAndStatus(ctx context.Context, userID, status string, limit, offset int) ([]*ent.Order, error)
}

func NewOrderServices(client *ent.Client) OrderServices {
	return &OrderServicesImpl{
		client: client,
	}
}

func (o *OrderServicesImpl) CreateOrder(ctx context.Context, orderDto *ent.Order, tableID, userid string, orderItems OrderWithItemsDTO ) (*ent.Order, error) {
	// Initialize totalAmount
	var totalAmount float64

	// Create a new order with the calculated total amount
	orderQuery := o.client.Order.Create().
		SetID(uuid.New().String()).
		SetStatus(orderDto.Status).
		SetTotalAmount(totalAmount). // Set the calculated total amount
		SetAdditionalInfo(orderDto.AdditionalInfo)

	if tableID != "" {
		orderQuery.AddTableIDs(tableID)
	}

	if userid != "" {
		orderQuery.SetUserID(userid)
	}

	newOrder, err := orderQuery.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	// Add order items
	for menuItemID, quantity := range orderItems.Items {
    menuItem, err := o.client.MenuItem.Get(ctx, menuItemID)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch menu item with ID %s: %v", menuItemID, err)
    }
    totalAmount += menuItem.Price * float64(quantity)

    // Create order item
    _, err = o.client.OrderItem.Create().
        SetID(uuid.New().String()).
        SetQuantity(quantity).
        AddOrder(newOrder).
        AddMenuItemIDs(menuItemID).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to create order item for menu item %s: %v", menuItemID, err)
    }
}

	return newOrder, nil
}

func (o *OrderServicesImpl) UpdateOrder(ctx context.Context, orderID string, orderDto OrderWithItemsDTO) (*ent.Order, error) {
	// Fetch the existing order
	existingOrder, err := o.client.Order.Get(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch existing order with ID %s: %v", orderID, err)
	}

	// Start with the existing total amount
	totalAmount := existingOrder.TotalAmount

	// Add the prices of new menu items to the existing total amount
	if len(orderDto.Items) > 0 {
		for menuItemId, qty := range orderDto.Items {
			menuItem, err := o.client.MenuItem.Get(ctx, menuItemId)
			if err != nil {
				return nil, fmt.Errorf("failed to fetch menu item with ID %s: %v", menuItemId, err)
			}
			totalAmount += menuItem.Price * float64(qty)

			// Create and associate a new OrderItem with this order
			_, err = o.client.OrderItem.Create().
				AddOrder(existingOrder).
				AddMenuItemIDs(menuItemId).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to add new order item: %v", err)
			}
		}
	}

	// Prepare the update query
	updateQuery := o.client.Order.UpdateOneID(orderID)

	// Update fields conditionally
	if orderDto.Order.Status != "" {
		updateQuery = updateQuery.SetStatus(orderDto.Order.Status)
	}
	// Update the total amount if new menu items are added
	if len(orderDto.Items) > 0 {
		updateQuery = updateQuery.SetTotalAmount(totalAmount)
	}
	if orderDto.Order.AdditionalInfo != nil {
		updateQuery = updateQuery.SetAdditionalInfo(orderDto.Order.AdditionalInfo)
	}

	// Execute the update query
	updatedOrder, err := updateQuery.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update order: %v", err)
	}

	return updatedOrder, nil
}

func (o *OrderServicesImpl) DeleteOrder(ctx context.Context, orderID string) error {
	err := o.client.Order.DeleteOneID(orderID).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete order with ID %s: %v", orderID, err)
	}
	return nil
}

func (o *OrderServicesImpl) GetOrder(ctx context.Context, orderID string) (*ent.Order, error) {
	order, err := o.client.Order.
		Query().
		Where(order.ID(orderID)).
		WithOrderItems().
		WithTable().
		WithUser().
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order with ID %s: %v", orderID, err)
	}
	return order, nil
}

func (o *OrderServicesImpl) GetOrders(ctx context.Context, limit, offset int) ([]*ent.Order, error) {
	orders, err := o.client.Order.
		Query().
		Limit(limit).
		Offset(offset).
		WithOrderItems().
		WithTable().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %v", err)
	}
	return orders, nil
}

func (o *OrderServicesImpl) GetOrdersByUserID(ctx context.Context, userID string, limit, offset int) ([]*ent.Order, error) {
	orders, err := o.client.Order.
		Query().
		Where(order.HasUserWith(user.ID(userID))).
		Limit(limit).
		Offset(offset).
		WithOrderItems().
		WithTable().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders for user ID %s: %v", userID, err)
	}
	return orders, nil
}

func (o *OrderServicesImpl) GetOrdersByTableID(ctx context.Context, tableID string, limit, offset int) ([]*ent.Order, error) {
	orders, err := o.client.Order.
		Query().
		Where(order.HasTableWith(placetable.ID(tableID))).
		Limit(limit).
		Offset(offset).
		WithOrderItems().
		WithTable().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders for table ID %s: %v", tableID, err)
	}
	return orders, nil
}

func (o *OrderServicesImpl) GetOrdersByTableIDAndStatus(ctx context.Context, tableID, status string, limit, offset int) ([]*ent.Order, error) {
	orders, err := o.client.Order.
		Query().
		Where(order.HasTableWith(placetable.ID(tableID)), order.StatusEQ(order.Status(status))).
		Limit(limit).
		Offset(offset).
		WithOrderItems().
		WithTable().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders for table ID %s with status %s: %v", tableID, status, err)
	}
	return orders, nil
}

func (o *OrderServicesImpl) GetOrdersByUserIDAndStatus(ctx context.Context, userID, status string, limit, offset int) ([]*ent.Order, error) {
	orders, err := o.client.Order.
		Query().
		Where(order.HasUserWith(user.ID(userID)), order.StatusEQ(order.Status(status))).
		Limit(limit).
		Offset(offset).
		WithOrderItems().
		WithTable().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders for user ID %s with status %s: %v", userID, status, err)
	}
	return orders, nil
}

// this takes care "23344": 2, "23345": 3
//func (o *OrderServicesImpl) UpdateOrder(ctx context.Context, orderID string, updateDto *ent.Order, newMenuItems map[string]int) (*ent.Order, error) {
//	// Fetch the existing order
//	existingOrder, err := o.client.Order.Get(ctx, orderID)
//	if err != nil {
//		return nil, fmt.Errorf("failed to fetch existing order with ID %s: %v", orderID, err)
//	}
//
//	// Start with the existing total amount
//	totalAmount := existingOrder.TotalAmount
//
//	// Process new menu items
//	for menuItemID, quantity := range newMenuItems {
//		menuItem, err := o.client.MenuItem.Get(ctx, menuItemID)
//		if err != nil {
//			return nil, fmt.Errorf("failed to fetch menu item with ID %s: %v", menuItemID, err)
//		}
//
//		totalAmount += menuItem.Price * float64(quantity)
//
//		_, err = o.client.OrderItem.Create().
//			SetID(uuid.New().String()).
//			SetQuantity(quantity).
//			SetOrder(existingOrder).
//			SetMenuItemID(menuItemID).
//			Save(ctx)
//		if err != nil {
//			return nil, fmt.Errorf("failed to add new order item: %v", err)
//		}
//	}
//
//	// Update the order with new details
//	updatedOrder, err := o.client.Order.UpdateOneID(orderID).
//		SetStatus(updateDto.Status).
//		SetTotalAmount(totalAmount).
//		SetAdditionalInfo(updateDto.AdditionalInfo).
//		Save(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("failed to update order: %v", err)
//	}
//
//	return updatedOrder, nil
//}

//func (o *OrderServicesImpl) CreateOrder(ctx context.Context, orderDto *ent.Order, tableID, userID string, orderItems map[string]int) (*ent.Order, error) {
//	totalAmount := 0.0
//
//	// Create a new order
//	newOrder, err := o.client.Order.Create().
//		SetID(uuid.New().String()).
//		SetStatus(orderDto.Status).
//		SetTotalAmount(orderDto.TotalAmount).
//		SetAdditionalInfo(orderDto.AdditionalInfo).
//		Save(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create order: %v", err)
//	}
//
//	// Process each menu item and its quantity
//	for menuItemID, quantity := range orderItems {
//		menuItem, err := o.client.MenuItem.Get(ctx, menuItemID)
//		if err != nil {
//			return nil, fmt.Errorf("failed to fetch menu item with ID %s: %v", menuItemID, err)
//		}
//
//		totalAmount += menuItem.Price * float64(quantity)
//
//		_, err = o.client.OrderItem.Create().
//			SetID(uuid.New().String()).
//			SetQuantity(quantity).
//			SetOrder(newOrder).
//			SetMenuItemID(menuItemID).
//			Save(ctx)
//		if err != nil {
//			return nil, fmt.Errorf("failed to add order item: %v", err)
//		}
//	}
//
//	// Update the order with the final total amount
//	_, err = o.client.Order.UpdateOneID(newOrder.ID).
//		SetTotalAmount(totalAmount).
//		Save(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("failed to update order total amount: %v", err)
//	}
//
//	return newOrder, nil
//}
