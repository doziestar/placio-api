package inventory

import (
	"context"
	"placio-app/ent"
)

type RoomService interface {
	GetRoom(ctx context.Context, roomID string) (*ent.Room, error)
	CreateRoom(ctx context.Context, roomData map[string]interface{}) (*ent.Room, error)
	UpdateRoom(ctx context.Context, roomID string, roomData map[string]interface{}) (*ent.Room, error)
	DeleteRoom(ctx context.Context, roomID string) error
}

type RoomServiceImpl struct {
	client *ent.Client
}

func NewRoomService(client *ent.Client) *RoomServiceImpl {
	return &RoomServiceImpl{client: client}
}

func (s *RoomServiceImpl) GetRoom(ctx context.Context, roomID string) (*ent.Room, error) {
	return s.client.Room.Get(ctx, roomID)
}

func (s *RoomServiceImpl) CreateRoom(ctx context.Context, roomData map[string]interface{}) (*ent.Room, error) {
	return s.client.Room.
		Create().
		//SetName(roomData["name"].(string)).
		//SetNillableCapacity(roomData["capacity"].(int)).
		Save(ctx)
}

func (s *RoomServiceImpl) UpdateRoom(ctx context.Context, roomID string, roomData map[string]interface{}) (*ent.Room, error) {
	return s.client.Room.
		UpdateOneID(roomID).
		//SetName(roomData["name"].(string)).
		//SetNillableCapacity(roomData["capacity"].(int)).
		Save(ctx)
}

func (s *RoomServiceImpl) DeleteRoom(ctx context.Context, roomID string) error {
	return s.client.Room.
		DeleteOneID(roomID).
		Exec(ctx)
}
