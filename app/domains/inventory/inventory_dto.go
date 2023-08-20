package inventory

type InventoryTypeData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type InventoryAttributeData struct {
	InventoryTypeID string `json:"inventory_type_id"`
	IsMandatory     bool   `json:"is_mandatory"`
	DataType        string `json:"data_type"`
}

type PlaceInventoryData struct {
	InventoryTypeID string `json:"inventory_type_id"`
	Quantity        int    `json:"quantity"`
	Location        string `json:"location"`
	Price           int    `json:"price"`
	ExpiryDate      string `json:"expiry_date"`
	Size            string `json:"size"`
	Color           string `json:"color"`
	Brand           string `json:"brand"`
	PurchaseDate    string `json:"purchase_date"`
}

type PlaceInventoryAttributeData struct {
	PlaceInventoryID string `json:"place_inventory_id"`
	Key              string `json:"key"`
	Value            string `json:"value"`
	DataType         string `json:"data_type"`
	// ... add other fields as required
}
