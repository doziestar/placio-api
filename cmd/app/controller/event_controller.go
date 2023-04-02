package controller

import (
	"github.com/gofiber/fiber/v2"
	_ "placio-app/Dto"
	_ "placio-app/models"
	"placio-app/service"
	"placio-app/utility"
)

type EventController struct {
	service service.IEventService
	utility utility.IUtility
}

func NewEventController(service service.IEventService, utility utility.IUtility) *EventController {
	return &EventController{service: service, utility: utility}
}

func (c *EventController) RegisterRoutes(router fiber.Router) {
	eventRouter := router.Group("/event")
	eventRouter.Post("/create", c.createEvent)
	eventRouter.Get("/:eventId", c.getEventID)
	eventRouter.Get("/get/location/:locationId", c.getEventByLocation)
	eventRouter.Get("/get/category/:categoryId", c.getEventByCategory)
	eventRouter.Get("/get/date/:date", c.getEventByDate)
	eventRouter.Delete("/delete/:eventId", c.deleteEvent)
	eventRouter.Put("/update/:eventId", c.updateEvent)
	eventRouter.Get("/participants/:eventId", c.getEventParticipants)

}

// CreateEvent godoc
// @Summary Create Event
// @Description Create Event
// @Tags Event
// @Accept  json
// @Produce  json
// @Param data body Dto.EventDto true "Event Data"
// @Success 200 {object} models.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /event/create [post]
func (c *EventController) createEvent(ctx *fiber.Ctx) error {
	return nil
}

// GetEventID godoc
// @Summary Get Event By ID
// @Description Get Event By ID
// @Tags Event
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Success 200 {object} models.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /event/{eventId} [get]
func (c *EventController) getEventID(ctx *fiber.Ctx) error {
	return nil
}

// GetEventByLocation godoc
// @Summary Get Event By Location
// @Description Get Event By Location
// @Tags Event
// @Accept  json
// @Produce  json
// @Param address path string true "Location Address"
// @Success 200 {object} []models.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /event/get/location/{locationId} [get]
func (c *EventController) getEventByLocation(ctx *fiber.Ctx) error {
	return nil
}

// GetEventByCategory godoc
// @Summary Get Event By Category
// @Description Get Event By Category
// @Tags Event
// @Accept  json
// @Produce  json
// @Param categoryId path string true "Category ID"
// @Success 200 {object} []models.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /event/get/category/{categoryId} [get]
func (c *EventController) getEventByCategory(ctx *fiber.Ctx) error {
	return nil
}

// GetEventByDate godoc
// @Summary Get Event By Date
// @Description Get Event By Date
// @Tags Event
// @Accept  json
// @Produce  json
// @Param date path string true "Event Date"
// @Success 200 {object} []models.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /event/get/date/{date} [get]
func (c *EventController) getEventByDate(ctx *fiber.Ctx) error {
	return nil
}

// DeleteEvent godoc
// @Summary Delete Event
// @Description Delete Event
// @Tags Event
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Success 200 {object} Dto.SuccessDTO
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /event/delete/{eventId} [delete]
func (c *EventController) deleteEvent(ctx *fiber.Ctx) error {
	return nil
}

// UpdateEvent godoc
// @Summary Update Event
// @Description Update Event
// @Tags Event
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Param data body Dto.EventDto true "Event Data"
// @Success 200 {object} models.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /event/update/{eventId} [put]
func (c *EventController) updateEvent(ctx *fiber.Ctx) error {
	return nil
}

// GetEventParticipants godoc
// @Summary Get Event Participants
// @Description Get Event Participants
// @Tags Event
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Success 200 {object} []models.User
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /event/participants/{eventId} [get]
func (c *EventController) getEventParticipants(ctx *fiber.Ctx) error {
	return nil
}
