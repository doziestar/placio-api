package amenities

type CreateAmenityInput struct {
	Name string
	Icon string
}

type UpdateAmenityInput struct {
	Name *string
	Icon *string
}

type AmenityAdditionDTO struct {
	AmenityIDs []string `json:"amenity_ids"`
}
