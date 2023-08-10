package models

type Subscription struct {
	ID   string `gorm:"primaryKey"`
	Plan string
}

func (s *Subscription) TableName() string {
	return "subscriptions"
}

func (s *Subscription) GetSubscriptionID() string {
	return s.ID
}

//
//func init() {
//	stripe.Key = "sk_test_..."
//}
//
//// Retrieve a Stripe subscription
//func RetrieveSubscription(id string) (*stripe.Subscription, error) {
//	subscription, err := stripe.Sub.Get(id, &stripe.SubscriptionParams{
//		Expand: []string{"latest_invoice.payment_intent"},
//	})
//	if err != nil {
//		return nil, err
//	}
//	return subscription, nil
//}
//
//// Upgrade or downgrade a Stripe subscription to a different plan
//func UpdateSubscription(subscription *stripe.Subscription, plan string) (*stripe.Subscription, error) {
//	if len(subscription.Items.Data) == 0 {
//		return nil, errors.New("no subscription items found")
//	}
//	subscription, err := sub.Update(subscription.ID, &stripe.SubscriptionParams{
//		Items: []*stripe.SubscriptionItemsParams{
//			{
//				ID:   subscription.Items.Data[0].ID,
//				Plan: &plan,
//			},
//		},
//	})
//	if err != nil {
//		return nil, err
//	}
//	return subscription, nil
//}
//
//// Cancel a Stripe subscription
//func CancelSubscription(id string) (*stripe.Subscription, error) {
//	subscription, err := sub.Cancel(id, &stripe.SubscriptionCancelParams{})
//	if err != nil {
//		return nil, err
//	}
//	return subscription, nil
//}
//
//// Retrieve a Stripe customer
//func RetrieveCustomer(id string) (*stripe.Customer, error) {
//	customer, err := cust.Get(id, nil)
//	if err != nil {
//		return nil, err
//	}
//	return customer, nil
//}
//
//// Create a new Stripe customer
//func CreateCustomer(email string, token string) (*stripe.Customer, error) {
//	customer, err := cust.New(&stripe.CustomerParams{
//		Email: &email,
//		Source: &stripe.SourceParams{
//			Token: &token,
//		},
//	})
//	if err != nil {
//		return nil, err
//	}
//	return customer, nil
//}
//
//// Update a Stripe customer's card details
//func UpdateCustomer(id string, token string) (*stripe.Customer, error) {
//	customer, err := stripe.Customer{}.Update(id, &stripe.CustomerParams{
//		Source: &stripe.SourceParams{
//			Token: &token,
//		},
//	})
//	if err != nil {
//		return nil, err
//	}
//	return customer, nil
//}
//
//// List the invoices paid by a Stripe customer
//func ListCustomerInvoices(id string, limit int64) ([]*stripe.Invoice, error) {
//	params := &stripe.InvoiceListParams{
//		Customer: &id,
//		Limit:    &limit,
//	}
//	invoices := make([]*stripe.Invoice, 0)
//	i := inv.List(params)
//	for i.Next() {
//		invoices = append(invoices, i.Invoice())
//	}
//	if err := i.Err(); err != nil {
//		return nil, err
//	}
//	return invoices, nil
//}
//
//// Subscribe a Stripe customer to a plan
//func SubscribeCustomer(id string, plan string) (*stripe.Subscription, error) {
//	subscription, err := sub.New(&stripe.SubscriptionParams{
//		Customer: &id,
//		Items: []*stripe.SubscriptionItemsParams{
//			{
//				Plan: &plan,
//			},
//		},
//		EnableIncompletePayments: stripe.Bool(true),
//		Expand: []string{"latest_invoice.payment_intent"},
//	})
//	if err != nil {
//		return nil, err
//	}
//	return subscription, nil
//}
//
