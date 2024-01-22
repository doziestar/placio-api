package smartRoom

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"placio-app/domains/media"
	"placio-app/ent"
	"placio-app/ent/place"
	"placio-app/ent/room"
	"placio-app/ent/roomcategory"
	"placio-app/utility"
	"placio-pkg/errors"

	"github.com/google/uuid"
)

type ISmartRoom interface {
	CreateRoomCategory(ctx context.Context, placeId string, roomCategoryDto *ent.RoomCategory, mediaFiles []*multipart.FileHeader) (*ent.RoomCategory, error)
	GetRoomCategories(ctx context.Context, placeId string) ([]*ent.RoomCategory, error)
	GetRoomCategoryByID(ctx context.Context, categoryId string) (*ent.RoomCategory, error)
	UpdateRoomCategory(ctx context.Context, categoryId string, roomCategoryDto *ent.RoomCategory) (*ent.RoomCategory, error)
	DeleteRoomCategory(ctx context.Context, categoryId string) error
	RestoreRoomCategory(ctx context.Context, categoryId string) (*ent.RoomCategory, error)

	CreateRoom(ctx context.Context, categoryId string, roomDto *ent.Room, mediaFiles []*multipart.FileHeader) (*ent.Room, error)
	GetRooms(ctx context.Context, categoryId string) ([]*ent.Room, error)
	GetRoomByID(ctx context.Context, roomId string) (*ent.Room, error)
	UpdateRoom(ctx context.Context, roomId string, roomDto *ent.Room) (*ent.Room, error)
	DeleteRoom(ctx context.Context, roomId string) error
	RestoreRoom(ctx context.Context, roomId string) (*ent.Room, error)
	GetRoomByPlaceID(ctx context.Context, placeId string) ([]*ent.Room, error)

	GenerateRoomQRCode(ctx context.Context, roomId string) (string, error)
}

type SmartRoomService struct {
	client       *ent.Client
	mediaService media.MediaService
}

func NewSmartRoomService(client *ent.Client, mediaService media.MediaService) ISmartRoom {
	return &SmartRoomService{client: client, mediaService: mediaService}
}

func (s *SmartRoomService) CreateRoomCategory(ctx context.Context, placeId string, roomCategoryDto *ent.RoomCategory, mediaFiles []*multipart.FileHeader) (*ent.RoomCategory, error) {
	log.Printf("Creating room category for place ID '%s'", placeId)
	// Check for empty placeId or roomCategoryDto
	if placeId == "" || roomCategoryDto == nil {
		return nil, errors.New("placeId and roomCategory data must be provided")
	}

	// First, check if the room category already exists for the given place
	existingCategory, err := s.client.RoomCategory.
		Query().
		Where(roomcategory.NameEQ(roomCategoryDto.Name), roomcategory.HasPlaceWith(place.IDEQ(placeId))).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("error querying existing room category: %w", err)
	}
	if existingCategory != nil {
		return nil, fmt.Errorf("room category '%s' already exists for place ID '%s'", roomCategoryDto.Name, placeId)
	}

	// Create the room category
	roomCategoryCreate := s.client.RoomCategory.
		Create().
		SetID(uuid.New().String()).
		SetName(roomCategoryDto.Name).
		SetDescription(roomCategoryDto.Description).
		AddPlaceIDs(placeId)

	// Save the new room category
	newRoomCategory, err := roomCategoryCreate.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating room category: %w", err)
	}

	// Handle media files asynchronously if provided
	if len(mediaFiles) > 0 {
		go func(roomCategoryID string, mediaFiles []*multipart.FileHeader) {
			// Use a separate context for asynchronous operation
			asyncCtx := context.Background()

			media, err := s.mediaService.UploadAndCreateMedia(asyncCtx, mediaFiles)
			if err != nil {
				log.Printf("error uploading media for room category ID '%s': %v", roomCategoryID, err)
				return
			}

			_, err = s.client.RoomCategory.
				UpdateOneID(roomCategoryID).
				AddMedia(media...).
				Save(asyncCtx)
			if err != nil {
				log.Printf("error adding media to room category ID '%s': %v", roomCategoryID, err)
			}
		}(newRoomCategory.ID, mediaFiles)
	}

	return newRoomCategory, nil
}

func (s *SmartRoomService) GetRoomCategories(ctx context.Context, placeId string) ([]*ent.RoomCategory, error) {
	if placeId == "" {
		return nil, errors.New("placeId must be provided")
	}

	roomCategories, err := s.client.RoomCategory.
		Query().
		Where(roomcategory.HasPlaceWith(place.IDEQ(placeId))).
		WithMedia().
		WithRooms(func(roomQuery *ent.RoomQuery) {
			roomQuery.WithMedia()
			roomQuery.WithReservations()
			roomQuery.WithBookings()

		}).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying room categories for place ID '%s': %w", placeId, err)
	}

	if len(roomCategories) == 0 {
		return nil, fmt.Errorf("no room categories found for place ID '%s'", placeId)
	}

	return roomCategories, nil
}

func (s *SmartRoomService) GetRoomCategoryByID(ctx context.Context, categoryId string) (*ent.RoomCategory, error) {
	if categoryId == "" {
		return nil, errors.New("categoryId must be provided")
	}

	roomCategory, err := s.client.RoomCategory.
		Query().
		Where(roomcategory.IDEQ(categoryId)).
		WithMedia().
		WithRooms(func(roomQuery *ent.RoomQuery) {
			roomQuery.WithMedia()
			roomQuery.WithReservations()
			roomQuery.WithBookings()

		}).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("room category with ID '%s' not found", categoryId)
		}
		return nil, fmt.Errorf("error querying room category with ID '%s': %w", categoryId, err)
	}

	return roomCategory, nil
}

func (s *SmartRoomService) UpdateRoomCategory(ctx context.Context, categoryId string, roomCategoryDto *ent.RoomCategory) (*ent.RoomCategory, error) {
	if categoryId == "" {
		return nil, errors.New("categoryId must be provided")
	}
	if roomCategoryDto == nil {
		return nil, errors.New("roomCategory data must be provided")
	}

	// First, check if the room category exists
	exists, err := s.client.RoomCategory.
		Query().
		Where(roomcategory.IDEQ(categoryId)).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("error checking existence of room category with ID '%s': %w", categoryId, err)
	}
	if exists == nil {
		return nil, fmt.Errorf("room category with ID '%s' does not exist", categoryId)
	}

	if roomCategoryDto.Name == "" {
		roomCategoryDto.Name = exists.Name
	}

	if roomCategoryDto.Price == "" {
		roomCategoryDto.Price = exists.Price
	}

	if roomCategoryDto.Description == "" {
		roomCategoryDto.Description = exists.Description
	}

	// Update operation
	updateOp := s.client.RoomCategory.
		UpdateOneID(categoryId).
		SetName(roomCategoryDto.Name).
		SetPrice(roomCategoryDto.Price).
		SetDescription(roomCategoryDto.Description)

	updatedCategory, err := updateOp.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("error updating room category with ID '%s': %w", categoryId, err)
	}

	return updatedCategory, nil
}

func (s *SmartRoomService) DeleteRoomCategory(ctx context.Context, categoryId string) error {
	if categoryId == "" {
		return errors.New("categoryId must be provided")
	}

	// First, check if the room category exists
	exists, err := s.client.RoomCategory.
		Query().
		Where(roomcategory.IDEQ(categoryId)).
		First(ctx)
	if err != nil {
		return fmt.Errorf("error checking existence of room category with ID '%s': %w", categoryId, err)
	}
	if exists == nil {
		return fmt.Errorf("room category with ID '%s' does not exist", categoryId)
	}

	// Delete operation
	err = s.client.RoomCategory.
		DeleteOneID(categoryId).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("error deleting room category with ID '%s': %w", categoryId, err)
	}

	return nil
}

func (s SmartRoomService) RestoreRoomCategory(ctx context.Context, categoryId string) (*ent.RoomCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SmartRoomService) CreateRoom(ctx context.Context, categoryId string, roomDto *ent.Room, mediaFiles []*multipart.FileHeader) (*ent.Room, error) {
	if categoryId == "" {
		return nil, errors.New("categoryId must be provided")
	}
	if roomDto == nil {
		return nil, errors.New("room data must be provided")
	}

	// Check if the category exists
	exists, err := s.client.RoomCategory.
		Query().
		Where(roomcategory.IDEQ(categoryId)).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("error checking room category existence: %w", err)
	}
	if exists == nil {
		return nil, fmt.Errorf("room category with ID '%s' does not exist", categoryId)
	}

	log.Printf("Creating room for category ID '%s'", roomDto)
	// check if a room with the same room name or room number already exist, if yes return an error
	roomAlreadyExists, err := s.client.Room.
		Query().
		Where(room.RoomNumberEQ(roomDto.RoomNumber)).
		Where(room.HasRoomCategoryWith(roomcategory.IDEQ(categoryId))).
		First(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, fmt.Errorf("error checking room existence: %w", err)
		}
	}

	if roomAlreadyExists != nil {
		return nil, fmt.Errorf("room with room number '%d' already exists", roomDto.RoomNumber)
	}

	// Create room
	room, err := s.client.Room.
		Create().
		SetID(uuid.New().String()).
		SetRoomNumber(roomDto.RoomNumber).
		SetRoomType(roomDto.RoomType).
		SetRoomStatus(roomDto.RoomStatus).
		SetRoomRating(roomDto.RoomRating).
		SetRoomPrice(roomDto.RoomPrice).
		SetAvailability(roomDto.Availability).
		SetImage(roomDto.Image).
		AddRoomCategoryIDs(categoryId).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating room: %w", err)
	}

	// Handle media files asynchronously if provided
	if len(mediaFiles) > 0 {
		go func(roomID string, mediaFiles []*multipart.FileHeader) {
			asyncCtx := context.Background()
			media, mediaErr := s.mediaService.UploadAndCreateMedia(asyncCtx, mediaFiles)
			if mediaErr != nil {
				log.Printf("error uploading media for room ID '%s': %v", roomID, mediaErr)
				return
			}

			_, mediaErr = s.client.Room.
				UpdateOneID(roomID).
				AddMedia(media...).
				Save(asyncCtx)
			if mediaErr != nil {
				log.Printf("error adding media to room ID '%s': %v", roomID, mediaErr)
			}
		}(room.ID, mediaFiles)
	}

	return room, nil
}

func (s *SmartRoomService) GetRooms(ctx context.Context, categoryId string) ([]*ent.Room, error) {
	if categoryId == "" {
		return nil, errors.New("categoryId must be provided")
	}

	rooms, err := s.client.Room.
		Query().
		Where(room.HasRoomCategoryWith(roomcategory.IDEQ(categoryId))).
		WithMedia().
		WithReservations().
		WithBookings().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error retrieving rooms: %w", err)
	}

	return rooms, nil
}

func (s *SmartRoomService) GetRoomByID(ctx context.Context, roomId string) (*ent.Room, error) {
	if roomId == "" {
		return nil, errors.New("roomId must be provided")
	}

	room, err := s.client.Room.
		Query().
		Where(room.IDEQ(roomId)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("error retrieving room: %w", err)
	}

	return room, nil
}

func (s *SmartRoomService) GetRoomByPlaceID(ctx context.Context, placeId string) ([]*ent.Room, error) {
	if placeId == "" {
		return nil, errors.New("placeId must be provided")
	}

	rooms, err := s.client.Room.
		Query().
		Where(room.HasRoomCategoryWith(roomcategory.HasPlaceWith(place.IDEQ(placeId)))).
		WithRoomCategory(func(query *ent.RoomCategoryQuery) {
			query.WithPlace(func(query *ent.PlaceQuery) {
				query.WithMedias()
			})
			query.WithMedia()
		}).
		WithMedia().
		WithReservations().
		WithBookings().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error retrieving rooms: %w", err)
	}

	return rooms, nil
}

func (s *SmartRoomService) UpdateRoom(ctx context.Context, roomId string, roomDto *ent.Room) (*ent.Room, error) {
	if roomId == "" || roomDto == nil {
		return nil, errors.New("roomId and room data must be provided")
	}

	updateOp := s.client.Room.
		UpdateOneID(roomId).
		//SetRoomNumber(*roomDto.RoomNumber).
		SetRoomType(roomDto.RoomType).
		SetRoomStatus(roomDto.RoomStatus).
		SetRoomRating(roomDto.RoomRating).
		SetRoomPrice(roomDto.RoomPrice).
		SetAvailability(roomDto.Availability).
		SetImage(roomDto.Image)

	updatedRoom, err := updateOp.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("error updating room: %w", err)
	}

	return updatedRoom, nil
}

func (s *SmartRoomService) DeleteRoom(ctx context.Context, roomId string) error {
	if roomId == "" {
		return errors.New("roomId must be provided")
	}

	err := s.client.Room.
		DeleteOneID(roomId).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("error deleting room: %w", err)
	}

	return nil
}

func (s SmartRoomService) RestoreRoom(ctx context.Context, roomId string) (*ent.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SmartRoomService) GenerateRoomQRCode(ctx context.Context, roomId string) (string, error) {
	if roomId == "" {
		return "", errors.New("roomId must be provided")
	}

	roomURL := fmt.Sprintf("https://placio.io/places/rooms/%s", roomId)
	qrCodeURL, err := utility.GenerateAndUploadQRCode(ctx, roomURL, roomId)
	if err != nil {
		return "", fmt.Errorf("error generating and uploading QR code for room ID '%s': %w", roomId, err)
	}

	// Update room with QR code URL
	_, err = s.client.Room.
		UpdateOneID(roomId).
		SetQrCode(qrCodeURL).
		Save(ctx)
	if err != nil {
		return "", fmt.Errorf("error updating room with QR code URL: %w", err)
	}

	return qrCodeURL, nil
}
