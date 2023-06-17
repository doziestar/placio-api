package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"
)

type RoomController struct {
	roomService service.RoomService
}

func NewRoomController(roomService service.RoomService) *RoomController {
	return &RoomController{roomService: roomService}
}

func (rc *RoomController) RegisterRoutes(router *gin.RouterGroup) {
	roomRouter := router.Group("/rooms")
	{
		roomRouter.GET("/:id", utility.Use(rc.getRoom))
		roomRouter.POST("/", utility.Use(rc.createRoom))
		roomRouter.PATCH("/:id", utility.Use(rc.updateRoom))
		roomRouter.DELETE("/:id", utility.Use(rc.deleteRoom))
	}
}

// @Summary Get a room
// @Description Get a room by its ID
// @Tags Room
// @Accept json
// @Produce json
// @Param id path string true "ID of the room"
// @Success 200 {object} ent.Room "Successfully retrieved room"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/rooms/{id} [get]
func (rc *RoomController) getRoom(ctx *gin.Context) error {
	id := ctx.Param("id")

	room, err := rc.roomService.GetRoom(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, room)
	return nil
}

// @Summary Create a room
// @Description Create a new room
// @Tags Room
// @Accept json
// @Produce json
// @Param room body map[string]interface{} true "Room data"
// @Success 200 {object} ent.Room "Successfully created room"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/rooms/ [post]
func (rc *RoomController) createRoom(ctx *gin.Context) error {
	var roomData map[string]interface{}
	if err := ctx.ShouldBindJSON(&roomData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	room, err := rc.roomService.CreateRoom(ctx, roomData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, room)
	return nil
}

// @Summary Update a room
// @Description Update a room by its ID
// @Tags Room
// @Accept json
// @Produce json
// @Param id path string true "ID of the room"
// @Param room body map[string]interface{} true "Room data"
// @Success 200 {object} ent.Room "Successfully updated room"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/rooms/{id} [patch]
func (rc *RoomController) updateRoom(ctx *gin.Context) error {
	id := ctx.Param("id")
	var roomData map[string]interface{}
	if err := ctx.ShouldBindJSON(&roomData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	room, err := rc.roomService.UpdateRoom(ctx, id, roomData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, room)
	return nil
}

// @Summary Delete a room
// @Description Delete a room by its ID
// @Tags Room
// @Accept json
// @Produce json
// @Param id path string true "ID of the room"
// @Success 200 {object} gin.H "Successfully deleted room"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/rooms/{id} [delete]
func (rc *RoomController) deleteRoom(ctx *gin.Context) error {
	id := ctx.Param("id")

	err := rc.roomService.DeleteRoom(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully deleted room",
	})

	return nil
}
