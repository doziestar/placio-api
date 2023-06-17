package controller

import "placio-app/service"

type RoomController struct {
	roomService service.RoomService
}

func NewRoomController(roomService service.RoomService) *RoomController {
	return &RoomController{roomService: roomService}
}
