package tickets

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/models"
	"placio-app/utility"
)

type TicketController struct {
	ticketService TicketService
}

func NewTicketController(ticketService TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

func (tc *TicketController) RegisterRoutes(router *gin.RouterGroup) {
	ticketRouter := router.Group("/tickets")
	{
		ticketRouter.POST("/", utility.Use(tc.createTicket))
		ticketRouter.GET("/:id", utility.Use(tc.getTicket))
		ticketRouter.PUT("/:id", utility.Use(tc.updateTicket))
		ticketRouter.DELETE("/:id", utility.Use(tc.deleteTicket))
		ticketRouter.GET("/event/:eventId", utility.Use(tc.getTicketsByEvent))
	}
}

// createTicket creates a new ticket
// @Summary Create a new ticket
// @Description Create a new ticket for the specified event
// @Tags Ticket
// @Accept json
// @Produce json
// @Param CreateTicketDto body models.Ticket true "Ticket Data"
// @Success 201 {object} models.Ticket "Successfully created ticket"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/tickets/ [post]
func (tc *TicketController) createTicket(ctx *gin.Context) error {
	data := new(models.Ticket)
	if err := ctx.BindJSON(data); err != nil {

		return err
	}

	ticket := &models.Ticket{
		//AttendeeID:     data.AttendeeID,
		EventID: data.EventID,
		//TicketOptionID: data.TicketOptionID,
	}

	err := tc.ticketService.CreateTicket(ticket)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusCreated, ticket)
	return nil
}

// getTicket retrieves a ticket by its ID
// @Summary GET ticket by ID
// @Description GET a ticket by its ID
// @Tags Ticket
// @Accept json
// @Produce json
// @Param id path string true "Ticket ID"
// @Success 200 {object} models.Ticket "Successfully retrieved ticket"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Ticket Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/tickets/{id} [get]
func (tc *TicketController) getTicket(ctx *gin.Context) error {
	ticketID := ctx.Param("id")
	ticket, err := tc.ticketService.GetTicketByEvent(ticketID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Ticket Not Found",
			})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
	}

	ctx.JSON(http.StatusOK, ticket)
	return nil
}

// updateTicket updates a ticket by its ID
// @Summary Update ticket by ID
// @Description Update a ticket by its ID
// @Tags Ticket
// @Accept json
// @Produce json
// @Param id path string true "Ticket ID"
// // @Param UpdateTicketDto body models.Ticket true "Ticket Data"
// // @Success 200 {object} models.Ticket "Successfully updated ticket"
// // @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// // @Failure 404 {object} Dto.ErrorDTO "Ticket Not Found"
// // @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// // @Router /api/v1/tickets/{id} [put]
func (tc *TicketController) updateTicket(ctx *gin.Context) error {
	ticketID := ctx.Param("id")
	data := new(models.Ticket)
	if err := ctx.BindJSON(data); err != nil {

		return err
	}

	ticket := &models.Ticket{
		ID: ticketID,
		//AttendeeID:     data.AttendeeID,
		EventID: data.EventID,
		//TicketOptionID: data.TicketOptionID,
	}
	err := tc.ticketService.UpdateTicket(ticket)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Ticket Not Found",
			})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, ticket)
	return nil
}

// deleteTicket deletes a ticket by its ID
// @Summary Delete ticket by ID
// @Description Delete a ticket by its ID
// @Tags Ticket
// @Accept json
// @Produce json
// @Param id path string true "Ticket ID"
// @Success 204 "Successfully deleted ticket"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Ticket Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/tickets/{id} [delete]
func (tc *TicketController) deleteTicket(ctx *gin.Context) error {
	ticketID := ctx.Param("id")
	err := tc.ticketService.DeleteTicket(ticketID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Ticket Not Found",
			})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

// getTicketsByEvent retrieves all tickets for an event
// @Summary GET tickets by event ID
// @Description GET all tickets for a specific event
// @Tags Ticket
// @Accept json
// @Produce json
// @Param eventId path string true "Event ID"
// @Success 200 {array} models.Ticket "Successfully retrieved tickets"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/tickets/event/{eventId} [get]
func (tc *TicketController) getTicketsByEvent(ctx *gin.Context) error {
	eventID := ctx.Param("eventId")
	tickets, err := tc.ticketService.GetTicketByEvent(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, tickets)
	return nil
}
