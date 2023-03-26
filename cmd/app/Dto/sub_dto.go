package Dto

type Subscription struct {
	SubscriptionId            string `json:"subscription_id"`
	SubscriptionName          string `json:"subscription_name"`
	SubscriptionDescription   string `json:"subscription_description"`
	SubscriptionPrice         string `json:"subscription_price"`
	SubscriptionCurrency      string `json:"subscription_currency"`
	SubscriptionInterval      string `json:"subscription_interval"`
	SubscriptionIntervalCount string `json:"subscription_interval_count"`
}
