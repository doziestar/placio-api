package amenities

type CreateAmenityInput struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type UpdateAmenityInput struct {
	Name *string
	Icon *string
}

type AmenityAdditionDTO struct {
	AmenityIDs []string `json:"amenity_ids"`
}
