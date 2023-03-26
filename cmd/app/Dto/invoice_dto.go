package Dto

type Invoice struct {
	InvoiceNumber string `json:"invoice_number"`
	InvoiceDate   string `json:"invoice_date"`
	InvoiceDue    string `json:"invoice_due"`
	InvoiceNote   string `json:"invoice_note"`
	InvoiceStatus string `json:"invoice_status"`
	InvoiceTotal  string `json:"invoice_total"`
	InvoicePaid   string `json:"invoice_paid"`
	InvoiceItems  []InvoiceItem
}

type InvoiceItem struct {
	InvoiceItemName        string `json:"invoice_item_name"`
	InvoiceItemDescription string `json:"invoice_item_description"`
	InvoiceItemQuantity    string `json:"invoice_item_quantity"`
	InvoiceItemPrice       string `json:"invoice_item_price"`
	InvoiceItemTotal       string `json:"invoice_item_total"`
}
