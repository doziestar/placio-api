package events_management

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/ent"
	_ "placio-app/ent"
	"placio-app/utility"
	"placio-pkg/middleware"
	"strconv"
)

type EventController struct {
	service IEventService
	utility utility.IUtility
}

func NewEventController(service IEventService, utility utility.IUtility) *EventController {
	return &EventController{service: service, utility: utility}
}

func (c *EventController) RegisterRoutes(router, routerWithoutAuth *gin.RouterGroup) {
	eventRouter := router.Group("/events")
	eventRouterWithoutAuth := routerWithoutAuth.Group("/events")
	eventRouter.POST("/", middleware.ErrorMiddleware(c.createEvent))
	eventRouter.PATCH("/:eventId", middleware.ErrorMiddleware(c.updateEvent))
	eventRouter.DELETE("/:eventId", middleware.ErrorMiddleware(c.deleteEvent))
	eventRouterWithoutAuth.GET("/:eventId", middleware.ErrorMiddleware(c.getEventByID))
	eventRouterWithoutAuth.GET("/", middleware.ErrorMiddleware(c.getEventsByFilters))
	eventRouter.POST("/:eventId/media", middleware.ErrorMiddleware(c.addMediaToEvent))
	eventRouter.DELETE("/:eventId/media/:mediaID", middleware.ErrorMiddleware(c.removeMediaFromEvent))
	eventRouter.POST("/:eventId/organizers", middleware.ErrorMiddleware(c.addOrganizersToEvent))
	eventRouterWithoutAuth.GET("/:eventId/organizers", middleware.ErrorMiddleware(c.getOrganizersForEvent))
	eventRouter.DELETE("/:eventId/organizers/:organizerId", middleware.ErrorMiddleware(c.removeOrganizerFromEvent))
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
// @Success 200 {object} ent.Event
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /events [post]
func (c *EventController) createEvent(ctx *gin.Context) error {
	var data *ent.Event
	if err := ctx.ShouldBindJSON(&data); err != nil {
		log.Println("error: ", err)
		return err
	}
	businessId := ctx.Query("businessId")

	if businessId == "" {

		return nil
	}

	event, err := c.service.CreateEvent(ctx, businessId, data)
	if err != nil {

		return err
	}
	ctx.JSON(http.StatusCreated, utility.ProcessResponse(event))
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
	var data *ent.Event
	if err := ctx.ShouldBindJSON(&data); err != nil {
		return err
	}
	eventId := ctx.Param("eventId")
	businessId := ctx.Query("businessId")
	event, err := c.service.UpdateEvent(ctx, eventId, businessId, data)
	if err != nil {

		return err
	}
	ctx.JSON(http.StatusOK, utility.ProcessResponse(event))
	return nil
}

func (c *EventController) addMediaToEvent(ctx *gin.Context) error {
	eventID := ctx.Param("id")

	files, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	uploadedFiles, ok := files.File["files"]
	if !ok || len(uploadedFiles) == 0 {
		return errors.New("No files uploaded")
	}

	event, err := c.service.AddMediaToEvent(ctx, eventID, uploadedFiles)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(event))
	return nil
}

// AddOrganizersToEvent handles adding organizers to an event.
func (c *EventController) addOrganizersToEvent(ctx *gin.Context) error {
	eventID := ctx.Param("eventId")
	var organizers []OrganizerInput
	if err := ctx.BindJSON(&organizers); err != nil {
		return err
	}

	if err := c.service.AddOrganizers(ctx, eventID, organizers); err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Organizers added successfully"})
	return nil
}

// GetOrganizersForEvent handles retrieving organizers for an event.
func (c *EventController) getOrganizersForEvent(ctx *gin.Context) error {
	eventID := ctx.Param("eventId")
	organizers, err := c.service.GetOrganizersForEvent(ctx, eventID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(organizers))
	return nil
}

// RemoveOrganizerFromEvent handles removing an organizer from an event.
func (c *EventController) removeOrganizerFromEvent(ctx *gin.Context) error {
	eventID := ctx.Param("eventId")
	organizerID := ctx.Param("organizerId") // Ensure this param name matches your routing definition

	if err := c.service.RemoveOrganizer(ctx, eventID, organizerID); err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Organizer removed successfully"})
	return nil
}

func (c *EventController) removeMediaFromEvent(ctx *gin.Context) error {
	eventID := ctx.Param("id")
	mediaID := ctx.Param("mediaID")

	if err := c.service.RemoveMediaFromEvent(ctx, eventID, mediaID); err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse("", "Media removed successfully"))
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
	var filter EventFilter
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {

		return err
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {

		return err
	}

	events, err := c.service.GetEvents(ctx, &filter, page, pageSize)
	if err != nil {

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
