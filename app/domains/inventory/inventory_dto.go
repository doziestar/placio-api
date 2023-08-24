package inventory

type inventoryTypeData struct {
	Name            string       `json:"name"`
	Description     string       `json:"description"`
	MeasurementUnit string       `json:"measurement_unit"`
	IndustryType    IndustryType `json:"industry_type"`
}

type inventoryAttributeData struct {
	InventoryTypeID string `json:"inventory_type_id"`
	IsMandatory     bool   `json:"is_mandatory"`
	DataType        string `json:"data_type"`
	Name            string `json:"name"`
}

type placeInventoryData struct {
	InventoryTypeID string `json:"inventory_type_id"`
	Quantity        int    `json:"quantity"`
	Location        string `json:"location"`
	Price           int    `json:"price"`
	ExpiryDate      string `json:"expiry_date"`
	Size            string `json:"size"`
	Color           string `json:"color"`
	Brand           string `json:"brand"`
	PurchaseDate    string `json:"purchase_date"`
	Name            string `json:"name"`
}

type placeInventoryAttributeData struct {
	PlaceInventoryID     string `json:"place_inventory_id"`
	Key                  string `json:"key"`
	Value                string `json:"value"`
	DataType             string `json:"data_type"`
	InventoryAttributeID string `json:"inventory_attribute_id"`
}
