package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"placio-app/Dto"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"
	"strconv"
)

type EventController struct {
	service service.IEventService
	utility utility.IUtility
}

func NewEventController(service service.IEventService, utility utility.IUtility) *EventController {
	return &EventController{service: service, utility: utility}
}

func (c *EventController) RegisterRoutes(router *gin.RouterGroup) {
	eventRouter := router.Group("/events")
	eventRouter.POST("/", utility.Use(c.createEvent))
	eventRouter.PATCH("/:eventId", utility.Use(c.updateEvent))
	eventRouter.DELETE("/:eventId", utility.Use(c.deleteEvent))
	eventRouter.GET("/:eventId", utility.Use(c.getEventByID))
	eventRouter.GET("/", utility.Use(c.getEventsByFilters))
	//eventRouter.Get("/participants/:eventId", c.getEventParticipants)

}

// CreateEvent godoc
// @Summary Create Event
// @Description Create Event
// @Tags Event
// @Accept  json
// @Produce  json
// @Param data body Dto.EventDTO true "Event Data"
// @Param businessId query string false "Business ID"
// @Param Authorization header string true "Bearer Token"
// @Security Bearer
// @Success 200 {object} ent.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /events [post]
func (c *EventController) createEvent(ctx *gin.Context) error {
	var data Dto.EventDTO
	if err := ctx.ShouldBindJSON(&data); err != nil {
		return err
	}
	businessId := ctx.Query("businessId")
	event, err := c.service.CreateEvent(ctx, businessId, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}
	ctx.JSON(http.StatusCreated, event)
	return nil
}

// UpdateEvent godoc
// @Summary Update Event
// @Description Update Event
// @Tags Event
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Param businessId query string false "Business ID"
// @Param Authorization header string true "Bearer Token"
// @Param data body Dto.EventDTO true "Event Data"
// @Success 200 {object} ent.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /events/{eventId} [put]
func (c *EventController) updateEvent(ctx *gin.Context) error {
	var data Dto.EventDTO
	if err := ctx.ShouldBindJSON(&data); err != nil {
		return err
	}
	eventId := ctx.Param("eventId")
	businessId := ctx.Query("businessId")
	event, err := c.service.UpdateEvent(ctx, eventId, businessId, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return err
	}
	ctx.JSON(http.StatusOK, event)
	return nil
}

// GetEvents godoc
// @Summary Get all Events
// @Description Get Events
// @Tags Event
// @Accept  json
// @Produce  json
// @Param filter query service.EventFilter false "Filter"
// @Param page query int false "Page Number"
// @Param pageSize query int false "Page Size"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {array} ent.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /events [get]
func (c *EventController) getEventsByFilters(ctx *gin.Context) error {
	var filter service.EventFilter
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return err
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return err
	}

	events, err := c.service.GetEvents(ctx, &filter, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, events)
	return nil
}

// @Summary Get Event By ID
// @Description Get a single event by its ID
// @Tags Event
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} ent.Event
// @Failure 400 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /events/{eventId} [get]
func (c *EventController) getEventByID(ctx *gin.Context) error {
	eventId := ctx.Param("eventId")
	event, err := c.service.GetEventByID(ctx, eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, event)
	return nil
}

// @Summary Delete Event
// @Description Delete an existing event
// @Tags Event
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {string} string "Deleted"
// @Failure 400 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /events/{eventId} [delete]
func (c *EventController) deleteEvent(ctx *gin.Context) error {
	eventId := ctx.Param("eventId")
	err := c.service.DeleteEvent(ctx, eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	ctx.JSON(http.StatusOK, "Deleted")
	return nil
}

// GetEventParticipants godoc
// @Summary Get Event Participants
// @Description Get Event Participants
// @Tags Event
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Success 200 {object} []ent.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /events/participants/{eventId} [get]
func (c *EventController) getEventParticipants(ctx *fiber.Ctx) error {
	return nil
}
