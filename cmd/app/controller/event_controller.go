package controller

import (
	"github.com/gofiber/fiber/v2"
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
	eventRouter.Get("/get/:eventId", c.getEventID)
	eventRouter.Get("/get/location/:locationId", c.getEventByLocation)
	eventRouter.Get("/get/category/:categoryId", c.getEventByCategory)
	eventRouter.Get("/get/date/:date", c.getEventByDate)
	eventRouter.Delete("/delete/:eventId", c.deleteEvent)
	eventRouter.Put("/update/:eventId", c.updateEvent)
	eventRouter.Get("/participants/:eventId", c.getEventParticipants)

}

func (c *EventController) createEvent(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) getEventID(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) getEventByLocation(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) getEventByCategory(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) getEventByDate(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) deleteEvent(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) updateEvent(ctx *fiber.Ctx) error {
	return nil
}

func (c *EventController) getEventParticipants(ctx *fiber.Ctx) error {
	return nil
}
