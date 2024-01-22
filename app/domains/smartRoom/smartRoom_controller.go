package smartRoom

import (
	"fmt"
	"log"
	"net/http"
	"placio-app/ent"
	"placio-app/utility"
	"placio-pkg/errors"
	"placio-pkg/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SmartRoomController struct {
	smartRoomService ISmartRoom
}

func NewSmartRoomController(smartRoomService ISmartRoom) *SmartRoomController {
	return &SmartRoomController{smartRoomService: smartRoomService}
}

func (c *SmartRoomController) RegisterRoutes(router, routerWithAuth *gin.RouterGroup) {
	const (
		placeIDParam    = "placeId"
		categoryIDParam = "categoryId"
		roomIDParam     = "roomId"
	)

	roomCategoryRouterWithAuth := routerWithAuth.Group("/roomCategories")
	routerCategoryRouterWithoutAuth := router.Group("/roomCategories")
	{
		roomCategoryRouterWithAuth.POST(fmt.Sprintf("/place/:%s", placeIDParam), middleware.ErrorMiddleware(c.createRoomCategory))
		routerCategoryRouterWithoutAuth.GET(fmt.Sprintf("/place/:%s", placeIDParam), middleware.ErrorMiddleware(c.getRoomCategories))
		routerCategoryRouterWithoutAuth.GET(fmt.Sprintf("/:%s", categoryIDParam), middleware.ErrorMiddleware(c.getRoomCategoryByID))
		roomCategoryRouterWithAuth.PUT(fmt.Sprintf("/:%s", categoryIDParam), middleware.ErrorMiddleware(c.updateRoomCategory))
		roomCategoryRouterWithAuth.DELETE(fmt.Sprintf("/:%s", categoryIDParam), middleware.ErrorMiddleware(c.deleteRoomCategory))
		roomCategoryRouterWithAuth.PATCH(fmt.Sprintf("/:%s/restore", categoryIDParam), middleware.ErrorMiddleware(c.restoreRoomCategory))
	}

	roomRouterWithAuth := routerWithAuth.Group("/rooms")
	routerRouterWithoutAuth := router.Group("/rooms")
	{
		roomRouterWithAuth.POST(fmt.Sprintf("/category/:%s", categoryIDParam), middleware.ErrorMiddleware(c.createRoom))
		routerRouterWithoutAuth.GET(fmt.Sprintf("/category/:%s", categoryIDParam), middleware.ErrorMiddleware(c.getRooms))
		routerRouterWithoutAuth.GET(fmt.Sprintf("/:%s", roomIDParam), middleware.ErrorMiddleware(c.getRoomByID))
		routerRouterWithoutAuth.GET(fmt.Sprintf("/place/:%s", placeIDParam), middleware.ErrorMiddleware(c.getRoomsByPlaceID))
		roomRouterWithAuth.PUT(fmt.Sprintf("/:%s", roomIDParam), middleware.ErrorMiddleware(c.updateRoom))
		roomRouterWithAuth.DELETE(fmt.Sprintf("/:%s", roomIDParam), middleware.ErrorMiddleware(c.deleteRoom))
		roomRouterWithAuth.PATCH(fmt.Sprintf("/:%s/restore", roomIDParam), middleware.ErrorMiddleware(c.restoreRoom))
		roomRouterWithAuth.POST(fmt.Sprintf("/:%s/generateQRCode", roomIDParam), middleware.ErrorMiddleware(c.generateRoomQRCode))
	}
}

func (c *SmartRoomController) createRoomCategory(ctx *gin.Context) error {
	log.Println("createRoomCategory")
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	log.Println("createRoomCategory", form)
	placeID := ctx.Param("placeId")
	log.Println("createRoomCategory", placeID)

	if placeID == "" {
		return errors.IDMissing
	}

	log.Println("createRoomCategory", placeID)

	var roomCategory ent.RoomCategory

	log.Println("createRoomCategory", roomCategory)

	if medias, exists := form.File["medias"]; !exists || len(medias) == 0 {
		return errors.ErrMediaMissing
	}

	log.Println("createRoomCategory", roomCategory)
	log.Println("createRoomCategory", form.Value)
	if name, exists := form.Value["name"]; exists {
		roomCategory.Name = name[0]
	}

	log.Println("createRoomCategory", roomCategory)

	if description, exists := form.Value["description"]; exists {
		roomCategory.Description = description[0]
	}

	log.Println("createRoomCategory", roomCategory)

	if price, exists := form.Value["price"]; exists {
		roomCategory.Price = price[0]
	}

	log.Println("roomCategory", roomCategory)

	roomCategoryData, err := c.smartRoomService.CreateRoomCategory(ctx, placeID, &roomCategory, form.File["medias"])
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusCreated, utility.ProcessResponse(roomCategoryData))
	return nil
}

func (c *SmartRoomController) getRoomCategories(ctx *gin.Context) error {
	placeID := ctx.Param("placeId")

	if placeID == "" {
		return errors.IDMissing
	}

	roomCategories, err := c.smartRoomService.GetRoomCategories(ctx, placeID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(roomCategories))
	return nil
}

func (c *SmartRoomController) getRoomCategoryByID(ctx *gin.Context) error {
	categoryID := ctx.Param("categoryId")

	if categoryID == "" {
		return errors.IDMissing
	}

	roomCategory, err := c.smartRoomService.GetRoomCategoryByID(ctx, categoryID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(roomCategory))
	return nil
}

func (c *SmartRoomController) updateRoomCategory(ctx *gin.Context) error {
	var roomCategory *ent.RoomCategory

	if err := ctx.ShouldBindJSON(&roomCategory); err != nil {
		return err
	}

	categoryID := ctx.Param("categoryId")

	if categoryID == "" {
		return errors.IDMissing
	}

	roomCategory, err := c.smartRoomService.UpdateRoomCategory(ctx, categoryID, roomCategory)

	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(roomCategory))

	return nil
}

func (c *SmartRoomController) deleteRoomCategory(ctx *gin.Context) error {
	categoryID := ctx.Param("categoryId")

	if categoryID == "" {
		return errors.IDMissing
	}

	err := c.smartRoomService.DeleteRoomCategory(ctx, categoryID)

	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(nil))
	return nil
}

func (c *SmartRoomController) restoreRoomCategory(ctx *gin.Context) error {
	categoryID := ctx.Param("categoryId")

	if categoryID == "" {
		return errors.IDMissing
	}

	//err := c.smartRoomService.RestoreRoomCategory(ctx, categoryID)
	//
	//if err != nil {
	//	return err
	//}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(nil))
	return nil
}

func (c *SmartRoomController) createRoom(ctx *gin.Context) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	categoryID := ctx.Param("categoryId")

	if categoryID == "" {
		return errors.IDMissing
	}

	var room ent.Room

	if len(form.File["medias"]) == 0 {
		return errors.ErrMediaMissing
	}

	if name, exists := form.Value["name"]; exists {
		room.Name = name[0]
	}

	if description, exists := form.Value["description"]; exists {
		room.Description = description[0]
	}

	if price, exists := form.Value["price"]; exists {
		priceStr := price[0]
		priceFloat, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return err
		}
		room.RoomPrice = priceFloat
	}

	if roomNumber, exists := form.Value["roomNumber"]; exists {
		roomNum, err := strconv.ParseInt(roomNumber[0], 10, 64)
		if err != nil {
			return errors.New("Error parsing roomNumber")
		}
		if roomNum == 0 {
			return errors.New("Invalid roomNumber, it can't be 0")
		}
		room.RoomNumber = int(roomNum)
	} else {
		return errors.New("Missing roomNumber")
	}

	//if extras, exists := form.Value["extras"]; exists {
	//	room.Extras = extras[0]
	//}

	if availability, exists := form.Value["availability"]; exists {
		availabilityStr := availability[0]
		availabilityBool, err := strconv.ParseBool(availabilityStr)
		if err != nil {
			room.Availability = true
		}
		room.Availability = availabilityBool
	}

	roomData, err := c.smartRoomService.CreateRoom(ctx, categoryID, &room, form.File["medias"])
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusCreated, utility.ProcessResponse(roomData))
	return nil
}

func (c *SmartRoomController) getRooms(ctx *gin.Context) error {
	categoryID := ctx.Param("categoryId")
	if categoryID == "" {
		return errors.IDMissing
	}

	rooms, err := c.smartRoomService.GetRooms(ctx, categoryID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(rooms))
	return nil
}

func (c *SmartRoomController) getRoomByID(ctx *gin.Context) error {
	roomID := ctx.Param("roomId")
	if roomID == "" {
		return errors.IDMissing
	}

	room, err := c.smartRoomService.GetRoomByID(ctx, roomID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(room))
	return nil
}

func (c *SmartRoomController) getRoomsByPlaceID(ctx *gin.Context) error {
	placeId := ctx.Param("placeId")
	if placeId == "" {
		return errors.IDMissing
	}

	rooms, err := c.smartRoomService.GetRoomByPlaceID(ctx, placeId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(rooms))
	return nil
}

func (c *SmartRoomController) updateRoom(ctx *gin.Context) error {
	var roomDto *ent.Room
	if err := ctx.ShouldBindJSON(&roomDto); err != nil {
		return err
	}

	roomID := ctx.Param("roomId")
	if roomID == "" {
		return errors.IDMissing
	}

	updatedRoom, err := c.smartRoomService.UpdateRoom(ctx, roomID, roomDto)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(updatedRoom))
	return nil
}

func (c *SmartRoomController) deleteRoom(ctx *gin.Context) error {
	roomID := ctx.Param("roomId")
	if roomID == "" {
		return errors.IDMissing
	}

	err := c.smartRoomService.DeleteRoom(ctx, roomID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse("Room successfully deleted"))
	return nil
}

func (c *SmartRoomController) restoreRoom(ctx *gin.Context) error {
	roomID := ctx.Param("roomId")
	if roomID == "" {
		return errors.IDMissing
	}

	restoredRoom, err := c.smartRoomService.RestoreRoom(ctx, roomID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(restoredRoom))
	return nil
}

func (c *SmartRoomController) generateRoomQRCode(ctx *gin.Context) error {
	roomID := ctx.Param("roomId")
	if roomID == "" {
		return errors.IDMissing
	}

	qrCodeURL, err := c.smartRoomService.GenerateRoomQRCode(ctx, roomID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(qrCodeURL))
	return nil
}
