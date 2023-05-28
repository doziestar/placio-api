package controller

import (
	"errors"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TicketOptionController struct {
	ticketOptionService service.TicketOptionService
}

func NewTicketOptionController(ticketOptionService service.TicketOptionService) *TicketOptionController {
	return &TicketOptionController{ticketOptionService: ticketOptionService}
}

func (toc *TicketOptionController) RegisterRoutes(router *gin.RouterGroup) {
	ticketOptionRouter := router.Group("/ticketOptions")
	// ticketOptionRouter.Use(middleware.Verify("user"))
	ticketOptionRouter.POST("/", utility.Use(toc.createTicketOption))
	ticketOptionRouter.GET("/:id", utility.Use(toc.getTicketOption))
	ticketOptionRouter.PUT("/:id", utility.Use(toc.updateTicketOption))
	ticketOptionRouter.DELETE("/:id", utility.Use(toc.deleteTicketOption))
	ticketOptionRouter.GET("/event/:eventID", utility.Use(toc.getTicketOptionsByEvent))
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
func (toc *TicketOptionController) createTicketOption(ctx *gin.Context) error {
	data := new(models.TicketOption)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
	}

	ticketOption := &models.TicketOption{
		EventID:  data.EventID,
		Name:     data.Name,
		Price:    data.Price,
		Quantity: data.Quantity,
	}

	newTicketOption := toc.ticketOptionService.CreateTicketOption(ticketOption)
	if newTicketOption == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return errors.New("ticket option not created")
	}

	ctx.JSON(http.StatusCreated, newTicketOption)
	return nil
}

// getTicketOption retrieves a ticket option by its ID
// @Summary GET ticket option by ID
// @Description GET a ticket option by its ID
// @Tags TicketOption
// @Accept json
// @Produce json
// @Param id path string true "Ticket Option ID"
// @Success 200 {object} models.TicketOption "Successfully retrieved ticket option"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Ticket Option Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ticket-options/{id} [get]
func (toc *TicketOptionController) getTicketOption(ctx *gin.Context) error {
	ticketOptionID := ctx.Param("id")
	ticketOption, err := toc.ticketOptionService.GetTicketOptionsByEvent(ticketOptionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Ticket Option Not Found",
			})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, ticketOption)
	return nil
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
func (toc *TicketOptionController) updateTicketOption(ctx *gin.Context) error {
	ticketOptionID := ctx.Param("id")
	data := new(models.TicketOption)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
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

	ctx.JSON(http.StatusOK, data)
	return nil
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
func (toc *TicketOptionController) deleteTicketOption(ctx *gin.Context) error {
	ticketOptionID := ctx.Param("id")
	if err := toc.ticketOptionService.DeleteTicketOption(ticketOptionID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Ticket Option Not Found",
			})
			return err

		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusNoContent, nil)
	return nil
}

// getTicketOptionsByEvent retrieves all ticket options for an event by its ID
// @Summary GET ticket options by event ID
// @Description GET all ticket options for an event by its ID
// @Tags TicketOption
// @Accept json
// @Produce json
// @Param eventId path string true "Event ID"
// @Success 200 {array} models.TicketOption "Successfully retrieved ticket options for event"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/events/{eventId}/ticket-options [get]
func (toc *TicketOptionController) getTicketOptionsByEvent(ctx *gin.Context) error {
	eventID := ctx.Param("eventId")
	ticketOptions, err := toc.ticketOptionService.GetTicketOptionsByEvent(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Ticket Options Not Found",
			})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, ticketOptions)
	return nil
}
