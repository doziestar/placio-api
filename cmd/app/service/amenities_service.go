package service

import (
	"context"
	"github.com/google/uuid"
	"placio-app/Dto"
	"placio-app/ent"
)

type AmenityService interface {
	CreateAmenity(input Dto.CreateAmenityInput) (*ent.Amenity, error)
	GetAmenity(id string) (*ent.Amenity, error)
	UpdateAmenity(id string, input Dto.UpdateAmenityInput) (*ent.Amenity, error)
	DeleteAmenity(id string) error
	GetAllAmenities() ([]*ent.Amenity, error)
}

type amenityServiceImpl struct {
	client *ent.Client
}

func NewAmenityService(client *ent.Client) AmenityService {
	return &amenityServiceImpl{client: client}
}

func (s *amenityServiceImpl) CreateAmenity(input Dto.CreateAmenityInput) (*ent.Amenity, error) {
	a, err := s.client.Amenity.
		Create().
		SetID(uuid.New().String()).
		SetName(input.Name).
		SetIcon(input.Icon).
		Save(context.Background())

	return a, err
}

func (s *amenityServiceImpl) GetAmenity(id string) (*ent.Amenity, error) {
	a, err := s.client.Amenity.
		Get(context.Background(), id)

	return a, err
}

func (s *amenityServiceImpl) UpdateAmenity(id string, input Dto.UpdateAmenityInput) (*ent.Amenity, error) {
	upd := s.client.Amenity.UpdateOneID(id)

	if input.Name != nil {
		upd = upd.SetName(*input.Name)
	}

	if input.Icon != nil {
		upd = upd.SetIcon(*input.Icon)
	}

	a, err := upd.Save(context.Background())

	return a, err
}

func (s *amenityServiceImpl) DeleteAmenity(id string) error {
	err := s.client.Amenity.
		DeleteOneID(id).
		Exec(context.Background())

	return err
}

func (s *amenityServiceImpl) GetAllAmenities() ([]*ent.Amenity, error) {
	a, err := s.client.Amenity.
		Query().
		All(context.Background())

	return a, err
}
