package tickets

type TicketPurchaseDTO struct {
	UserID         string                 // The ID of the user making the purchase
	EventID        string                 // The ID of the event for which tickets are being purchased
	TicketOptions  []TicketOptionPurchase // Details of ticket options being purchased
	PromotionCode  string                 // Optional promotion or discount code
	PaymentDetails PaymentDetails         // Payment information
}

// TicketOptionPurchase captures the details of a specific ticket option being purchased, including quantity.
type TicketOptionPurchase struct {
	OptionID string // The ID of the ticket option
	Quantity int    // The number of tickets of this option being purchased
}

// PaymentDetails encapsulates payment information necessary for processing the ticket purchase.
type PaymentDetails struct {
	Method      string       // Payment method, e.g., "credit_card", "paypal"
	Amount      float64      // Total payment amount
	Currency    string       // Currency code, e.g., "USD", "EUR"
	CardDetails *CardDetails // Details for credit card payments, nil if not applicable
	PayPalEmail string       // PayPal email, if using PayPal
	// Additional fields can be added here for other payment methods
}

// CardDetails provides the necessary details for processing credit card payments.
type CardDetails struct {
	CardNumber     string // Credit card number
	ExpirationDate string // Expiration date in "MM/YY" format
	CVC            string // Card verification code
	CardholderName string // Name of the cardholder
}
