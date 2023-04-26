package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	_ "placio-app/Dto"
	"placio-app/models"

	"placio-app/service"
)

type TicketController struct {
	ticketService service.TicketService
}

func NewTicketController(ticketService service.TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

func (tc *TicketController) RegisterRoutes(router *gin.RouterGroup) {
	ticketRouter := router.Group("/tickets")
	{
		ticketRouter.Post("/", tc.createTicket)
		ticketRouter.Get("/:id", tc.getTicket)
		ticketRouter.Put("/:id", tc.updateTicket)
		ticketRouter.Delete("/:id", tc.deleteTicket)
		ticketRouter.Get("/event/:eventId", tc.getTicketsByEvent)
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
func (tc *TicketController) createTicket(ctx *fiber.Ctx) error {
	data := new(models.Ticket)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	ticket := &models.Ticket{
		//AttendeeID:     data.AttendeeID,
		EventID: data.EventID,
		//TicketOptionID: data.TicketOptionID,
	}

	err := tc.ticketService.CreateTicket(ticket)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(ticket)
}

// getTicket retrieves a ticket by its ID
// @Summary Get ticket by ID
// @Description Get a ticket by its ID
// @Tags Ticket
// @Accept json
// @Produce json
// @Param id path string true "Ticket ID"
// @Success 200 {object} models.Ticket "Successfully retrieved ticket"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Ticket Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/tickets/{id} [get]
func (tc *TicketController) getTicket(ctx *fiber.Ctx) error {
	ticketID := ctx.Params("id")
	ticket, err := tc.ticketService.GetTicketByEvent(ticketID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Ticket Not Found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(ticket)
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
func (tc *TicketController) updateTicket(ctx *fiber.Ctx) error {
	ticketID := ctx.Params("id")
	data := new(models.Ticket)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
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
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Ticket Not Found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON("Ticket updated successfully")
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
func (tc *TicketController) deleteTicket(ctx *fiber.Ctx) error {
	ticketID := ctx.Params("id")
	err := tc.ticketService.DeleteTicket(ticketID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Ticket Not Found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON("Ticket deleted successfully")
}

// getTicketsByEvent retrieves all tickets for an event
// @Summary Get tickets by event ID
// @Description Get all tickets for a specific event
// @Tags Ticket
// @Accept json
// @Produce json
// @Param eventId path string true "Event ID"
// @Success 200 {array} models.Ticket "Successfully retrieved tickets"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/tickets/event/{eventId} [get]
func (tc *TicketController) getTicketsByEvent(ctx *fiber.Ctx) error {
	eventID := ctx.Params("eventId")
	tickets, err := tc.ticketService.GetTicketByEvent(eventID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(tickets)
}
