package amenities

type CreateAmenityInput struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Amenity struct {
	Amenities []CreateAmenityInput `json:"amenities"`
}

type UpdateAmenityInput struct {
	Name *string
	Icon *string
}

type AmenityAdditionDTO struct {
	AmenityIDs []string `json:"amenity_ids"`
}
