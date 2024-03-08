package tickets

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/utility"
	"placio-pkg/middleware"
)

type ticketController struct {
	service ITicketService
}

func NewTicketController(service ITicketService) *ticketController {
	return &ticketController{service: service}
}

func (c *ticketController) RegisterRoutes(router, routerWithoutAuth *gin.RouterGroup) {
	ticketRouter := router.Group("/tickets")
	ticketRouterWithoutAuth := routerWithoutAuth.Group("/tickets")

	// Ticket Option Management
	ticketRouter.POST("/options", middleware.ErrorMiddleware(c.createTicketOption))
	ticketRouter.PATCH("/options/:optionId", middleware.ErrorMiddleware(c.updateTicketOption))
	ticketRouter.DELETE("/options/:optionId", middleware.ErrorMiddleware(c.deleteTicketOption))
	ticketRouterWithoutAuth.GET("/event/:eventId/options", middleware.ErrorMiddleware(c.getTicketOptionsForEvent))
	ticketRouter.POST("/options/:optionId/media", middleware.ErrorMiddleware(c.addMediaToTicketOption))
	ticketRouter.DELETE("/options/:optionId/media/:mediaID", middleware.ErrorMiddleware(c.removeMediaFromTicketOption))

	// Ticket Management
	ticketRouter.POST("/", middleware.ErrorMiddleware(c.purchaseTicket))
	ticketRouter.POST("/:ticketId/validate", middleware.ErrorMiddleware(c.validateTicket))
	ticketRouter.POST("/:ticketId/cancel", middleware.ErrorMiddleware(c.cancelTicket))
	ticketRouter.POST("/:ticketId/transfer", middleware.ErrorMiddleware(c.transferTicket))
	ticketRouterWithoutAuth.GET("/:ticketId", middleware.ErrorMiddleware(c.getTicketDetails))
	ticketRouterWithoutAuth.GET("/user/:userId", middleware.ErrorMiddleware(c.getTicketsByUser))

	// Attendee Management
	routerWithoutAuth.GET("/events/:eventId/attendees", middleware.ErrorMiddleware(c.listAttendeesForEvent))
	ticketRouter.POST("/tickets/:ticketId/assign", middleware.ErrorMiddleware(c.assignTicketToAttendee))
}

// CreateTicketOption godoc
// @Summary Create a new ticket option
// @Description Create a new ticket option for an event
// @Accept json
// @Produce json
// @Tags Ticket
// @Param eventId path string true "Event ID"
// @Param data body TicketOptionDTO true "Ticket Option Data"
// @Success 200 {object} TicketOptionDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /tickets/options [post]
func (c *ticketController) createTicketOption(ctx *gin.Context) error {
	return nil
}

// UpdateTicketOption godoc
// @Summary Update a ticket option
// @Description Update a ticket option for an event
// @Accept json
// @Produce json
// @Tags Ticket
// @Param optionId path string true "Option ID"
// @Param data body TicketOptionDTO true "Ticket Option Data"
// @Success 200 {object} TicketOptionDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /tickets/options/{optionId} [patch]
func (c *ticketController) updateTicketOption(ctx *gin.Context) error {
	return nil
}

// DeleteTicketOption godoc
// @Summary Delete a ticket option
// @Description Delete a ticket option for an event
// @Accept json
// @Produce json
// @Tags Ticket
// @Param optionId path string true "Option ID"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /tickets/options/{optionId} [delete]
func (c *ticketController) deleteTicketOption(ctx *gin.Context) error {
	return nil
}

// GetTicketOptionsForEvent godoc
// @Summary Get ticket options for an event
// @Description Get ticket options for an event
// @Accept json
// @Produce json
// @Tags Ticket
// @Param eventId path string true "Event ID"
// @Success 200 {object} []TicketOptionDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tickets/event/{eventId}/options [get]
func (c *ticketController) getTicketOptionsForEvent(ctx *gin.Context) error {
	return nil
}

// PurchaseTicket godoc
// @Summary Purchase a ticket
// @Description Purchase a ticket for an event
// @Accept json
// @Produce json
// @Tags Ticket
// @Param data body TicketPurchaseDTO true "Ticket Purchase Data"
// @Success 200 {object} TicketDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /tickets [post]
func (c *ticketController) purchaseTicket(ctx *gin.Context) error {
	return nil
}

// ValidateTicket godoc
// @Summary Validate a ticket
// @Description Validate a ticket for an event
// @Accept json
// @Produce json
// @Tags Ticket
// @Param ticketId path string true "Ticket ID"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /tickets/{ticketId}/validate [post]
func (c *ticketController) validateTicket(ctx *gin.Context) error {
	return nil
}

// CancelTicket godoc
// @Summary Cancel a ticket
// @Description Cancel a ticket for an event
// @Accept json
// @Produce json
// @Tags Ticket
// @Param ticketId path string true "Ticket ID"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /tickets/{ticketId}/cancel [post]
func (c *ticketController) cancelTicket(ctx *gin.Context) error {
	return nil
}

// TransferTicket godoc
// @Summary Transfer a ticket
// @Description Transfer a ticket for an event to another user
// @Accept json
// @Produce json
// @Tags Ticket
// @Param ticketId path string true "Ticket ID"
// @Param toUserId path string true "To User ID"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /tickets/{ticketId}/transfer/{toUserId} [post]
func (c *ticketController) transferTicket(ctx *gin.Context) error {
	return nil
}

// GetTicketsByUser godoc
// @Summary Get tickets by user
// @Description Get tickets for a user
// @Accept json
// @Produce json
// @Tags Ticket
// @Param userId path string true "User ID"
// @Success 200 {object} []TicketDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /tickets/user/{userId} [get]
func (c *ticketController) getTicketsByUser(ctx *gin.Context) error {

	return nil
}

// GetTicketDetails godoc
// @Summary Get ticket details
// @Description Get details of a ticket
// @Accept json
// @Produce json
// @Tags Ticket
// @Param ticketId path string true "Ticket ID"
// @Success 200 {object} TicketDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tickets/{ticketId} [get]
func (c *ticketController) getTicketDetails(ctx *gin.Context) error {
	return nil
}

// AddMediaToTicketOption godoc
// @Summary Add media to a ticket option
// @Description Add media files to a ticket option for an event
// @Accept multipart/form-data
// @Produce json
// @Tags TicketOption
// @Param ticketOptionID path string true "Ticket Option ID"
// @Param files formData file true "Media Files"
// @Success 200 {object} TicketOptionDTO "Successfully added media to ticket option"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Security ApiKeyAuth
// @Router /tickets/options/{ticketOptionID}/media [post]
func (c *ticketController) addMediaToTicketOption(ctx *gin.Context) error {
	// Extract ticketOptionID from URL
	ticketOptionID := ctx.Param("ticketOptionID")
	// Assume files are uploaded with form key 'files'
	form, _ := ctx.MultipartForm()
	files := form.File["files"]

	// Call the service layer method
	updatedTicketOption, err := c.service.AddMediaToTicketOption(ctx, ticketOptionID, files)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(updatedTicketOption))
	return nil
}

// RemoveMediaFromTicketOption godoc
// @Summary Remove media from a ticket option
// @Description Remove a media file from a ticket option for an event
// @Produce json
// @Tags TicketOption
// @Param ticketOptionID path string true "Ticket Option ID"
// @Param mediaID path string true "Media ID"
// @Success 200 "Successfully removed media from ticket option"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Security ApiKeyAuth
// @Router /tickets/options/{ticketOptionID}/media/{mediaID} [delete]
func (c *ticketController) removeMediaFromTicketOption(ctx *gin.Context) error {
	ticketOptionID := ctx.Param("ticketOptionID")
	mediaID := ctx.Param("mediaID")

	err := c.service.RemoveMediaFromTicketOption(ctx, ticketOptionID, mediaID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Media removed successfully"})
	return nil
}

// ListAttendeesForEvent godoc
// @Summary List attendees for an event
// @Description Get a list of all users who have purchased tickets for the specified event
// @Produce json
// @Tags Event
// @Param eventId path string true "Event ID"
// @Success 200 {array} UserDTO "List of attendees"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /events/{eventId}/attendees [get]
func (c *ticketController) listAttendeesForEvent(ctx *gin.Context) error {
	eventId := ctx.Param("eventId")

	attendees, err := c.service.ListAttendeesForEvent(ctx, eventId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(attendees))
	return nil
}

// AssignTicketToAttendee godoc
// @Summary Assign a ticket to an attendee
// @Description Assign a ticket to a different user by specifying the ticket and user IDs
// @Accept json
// @Produce json
// @Tags Ticket
// @Param ticketId path string true "Ticket ID"
// @Param data body AssignTicketDTO true "Assign Ticket Data"
// @Success 200 "Ticket assigned successfully"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /tickets/{ticketId}/assign [post]
func (c *ticketController) assignTicketToAttendee(ctx *gin.Context) error {
	ticketId := ctx.Param("ticketId")
	attendeeId := ctx.Query("attendeeId")

	if attendeeId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "attendeeId is required"})
		return nil
	}

	err := c.service.AssignTicketToAttendee(ctx, ticketId, attendeeId)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ticket assigned successfully"})
	return nil
}
