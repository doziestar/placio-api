package Dto

type CreateAmenityInput struct {
	Name string
	Icon string
}

type UpdateAmenityInput struct {
	Name *string
	Icon *string
}
