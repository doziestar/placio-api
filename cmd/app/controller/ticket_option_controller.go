package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	_ "placio-app/Dto"
	"placio-app/models"
	"placio-app/service"
)

type TicketOptionController struct {
	ticketOptionService service.TicketOptionService
}

func NewTicketOptionController(ticketOptionService service.TicketOptionService) *TicketOptionController {
	return &TicketOptionController{ticketOptionService: ticketOptionService}
}

func (toc *TicketOptionController) RegisterRoutes(router fiber.Router) {
	ticketOptionRouter := router.Group("/ticketOptions")
	{
		ticketOptionRouter.Post("/", toc.createTicketOption)
		ticketOptionRouter.Get("/:id", toc.getTicketOption)
		ticketOptionRouter.Put("/:id", toc.updateTicketOption)
		ticketOptionRouter.Delete("/:id", toc.deleteTicketOption)
		ticketOptionRouter.Get("/event/:eventID", toc.getTicketOptionsByEvent)
	}
}

// Implement the handlers (createTicketOption, getTicketOption, updateTicketOption, deleteTicketOption, getTicketOptionsByEvent) similar to the CommentController handlers.

// createTicketOption creates a new ticket option
// @Summary Create a new ticket option
// @Description Create a new ticket option for the specified event
// @Tags TicketOption
// @Accept json
// @Produce json
// @Param CreateTicketOptionDto body models.TicketOption true "Ticket Option Data"
// @Success 201 {object} models.TicketOption "Successfully created ticket option"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ticket-options/ [post]
func (toc *TicketOptionController) createTicketOption(ctx *fiber.Ctx) error {
	data := new(models.TicketOption)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	ticketOption := &models.TicketOption{
		EventID:  data.EventID,
		Name:     data.Name,
		Price:    data.Price,
		Quantity: data.Quantity,
	}

	newTicketOption := toc.ticketOptionService.CreateTicketOption(ticketOption)
	if newTicketOption == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(newTicketOption)
}

// getTicketOption retrieves a ticket option by its ID
// @Summary Get ticket option by ID
// @Description Get a ticket option by its ID
// @Tags TicketOption
// @Accept json
// @Produce json
// @Param id path string true "Ticket Option ID"
// @Success 200 {object} models.TicketOption "Successfully retrieved ticket option"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Ticket Option Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ticket-options/{id} [get]
func (toc *TicketOptionController) getTicketOption(ctx *fiber.Ctx) error {
	ticketOptionID := ctx.Params("id")
	ticketOption, err := toc.ticketOptionService.GetTicketOptionsByEvent(ticketOptionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Ticket Option Not Found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(ticketOption)
}

// updateTicketOption updates a ticket option by its ID
// @Summary Update ticket option by ID
// @Description Update a ticket option by its ID
// @Tags TicketOption
// @Accept json
// @Produce json
// @Param id path string true "Ticket Option ID"
// @Param UpdateTicketOptionDto body models.TicketOption true "Ticket Option Data"
// @Success 200 {object} models.TicketOption "Successfully updated ticket option"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Ticket Option Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ticket-options/{id} [put]
func (toc *TicketOptionController) updateTicketOption(ctx *fiber.Ctx) error {
	ticketOptionID := ctx.Params("id")
	data := new(models.TicketOption)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}
	data.ID = ticketOptionID
	//updatedTicketOption, err := toc.ticketOptionService.UpdateTicketOption(data)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
	//			"error": "Ticket Option Not Found",
	//		})
	//	}
	//	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Internal Server Error",
	//	})
	//}

	return ctx.Status(fiber.StatusOK).JSON(data)
}

// deleteTicketOption deletes a ticket option by its ID
// @Summary Delete ticket option by ID
// @Description Delete a ticket option by its ID
// @Tags TicketOption
// @Accept json
// @Produce json
// @Param id path string true "Ticket Option ID"
// @Success 204 "Successfully deleted ticket option"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Ticket Option Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ticket-options/{id} [delete]
func (toc *TicketOptionController) deleteTicketOption(ctx *fiber.Ctx) error {
	ticketOptionID := ctx.Params("id")
	if err := toc.ticketOptionService.DeleteTicketOption(ticketOptionID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Ticket Option Not Found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}

// getTicketOptionsByEvent retrieves all ticket options for an event by its ID
// @Summary Get ticket options by event ID
// @Description Get all ticket options for an event by its ID
// @Tags TicketOption
// @Accept json
// @Produce json
// @Param eventId path string true "Event ID"
// @Success 200 {array} models.TicketOption "Successfully retrieved ticket options for event"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/events/{eventId}/ticket-options [get]
func (toc *TicketOptionController) getTicketOptionsByEvent(ctx *fiber.Ctx) error {
	eventID := ctx.Params("eventId")
	ticketOptions, err := toc.ticketOptionService.GetTicketOptionsByEvent(eventID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(ticketOptions)
}
