package events_management

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	_ "placio-app/Dto"
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
	eventRouter.PATCH("/:eventId/cancel", middleware.ErrorMiddleware(c.cancelEvent))
	eventRouter.PATCH("/:eventId/makeActive", middleware.ErrorMiddleware(c.makeEventActive))
	eventRouter.DELETE("/:eventId", middleware.ErrorMiddleware(c.deleteEvent))
	eventRouterWithoutAuth.GET("/:eventId", middleware.ErrorMiddleware(c.getEventByID))
	eventRouterWithoutAuth.GET("/", middleware.ErrorMiddleware(c.getEventsByFilters))
	eventRouter.POST("/:eventId/media", middleware.ErrorMiddleware(c.addMediaToEvent))
	eventRouter.DELETE("/:eventId/media/:mediaID", middleware.ErrorMiddleware(c.removeMediaFromEvent))
	eventRouter.POST("/:eventId/organizers", middleware.ErrorMiddleware(c.addOrganizersToEvent))
	eventRouterWithoutAuth.GET("/:eventId/organizers", middleware.ErrorMiddleware(c.getOrganizersForEvent))
	eventRouterWithoutAuth.GET("/organizers/:organizerId", middleware.ErrorMiddleware(c.getEventsByOrganizerId))
	eventRouter.DELETE("/:eventId/organizers/:organizerId", middleware.ErrorMiddleware(c.removeOrganizerFromEvent))
}

// TODO: Add the rest of the methods
// 8. test add ticket to event
// 9. test get ticket by event id
// 10. test get ticket by id
// 11. test update ticket
// 12. test delete ticket
// 13. test get event participants
// 14. add event participants
// 15. remove event participants
// 17. add user ticket purchase
// 18. get user ticket purchase
// 19. update user ticket purchase
// 20. delete user ticket purchase
// 21. add ability for user to cancel event purchase
// 22. add notifications to event
// 23. automatically send notifications to event participants
// 24. add a button to create a post with event
// 25. add ability to add event to calendar

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
	var data *EventDTO
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
	var data *EventDTO
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

// CancelEvent godoc
// @Summary Cancel Event
// @Description Cancel Event
// @Tags Event
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Param Authorization header string true "
// @Success 200 {string} string "Event Cancelled"
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /events/{eventId}/cancel [patch]
func (c *EventController) cancelEvent(ctx *gin.Context) error {
	eventId := ctx.Param("eventId")
	err := c.service.CancelEvent(ctx, eventId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse("", "Event Cancelled"))
	return nil

}

// MakeEventActive godoc
// @Summary Make Event Active
// @Description Make Event Active
// @Tags Event
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Param Authorization header string true "
// @Success 200 {string} string "Event Activated"
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /events/{eventId}/makeActive [patch]
func (c *EventController) makeEventActive(ctx *gin.Context) error {
	eventId := ctx.Param("eventId")
	err := c.service.MakeEventActive(ctx, eventId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse("", "Event Activated"))
	return nil
}

func (c *EventController) addMediaToEvent(ctx *gin.Context) error {
	eventID := ctx.Param("eventId")

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
	var organizers []OrganizerInfo
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

// GetEventsByOrganizerId handles retrieving events for an organizer.
func (c *EventController) getEventsByOrganizerId(ctx *gin.Context) error {
	organizerID := ctx.Param("organizerId")
	events, err := c.service.GetEventsByOrganizerID(ctx, organizerID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(events))
	return nil
}

// RemoveOrganizerFromEvent handles removing an organizer from an event.
func (c *EventController) removeOrganizerFromEvent(ctx *gin.Context) error {
	eventID := ctx.Param("eventId")
	organizerID := ctx.Param("organizerId")

	if err := c.service.RemoveOrganizer(ctx, eventID, organizerID); err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Organizer removed successfully"})
	return nil
}

func (c *EventController) removeMediaFromEvent(ctx *gin.Context) error {
	eventID := ctx.Param("eventId")
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

	businessId := ctx.Query("businessId")

	if businessId != "" {
		events, err := c.service.GetEventByBusinessID(ctx, businessId)
		if err != nil {
			return err
		}
		ctx.JSON(http.StatusOK, utility.ProcessResponse(events))
		return nil
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

	ctx.JSON(http.StatusOK, utility.ProcessResponse(events))
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
	ctx.JSON(http.StatusOK, utility.ProcessResponse(event))
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
	ctx.JSON(http.StatusNoContent, utility.ProcessResponse("", "Deleted"))
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
