package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
)

type AttendeeController struct {
	attendeeService service.AttendeeService
}

func NewAttendeeController(attendeeService service.AttendeeService) *AttendeeController {
	return &AttendeeController{attendeeService: attendeeService}
}

func (ac *AttendeeController) RegisterRoutes(router *gin.RouterGroup) {
	attendeeRouter := router.Group("/attendees")
	{
		attendeeRouter.POST("/", utility.Use(ac.addAttendee))
		attendeeRouter.GET("/:id", utility.Use(ac.getAttendee))
		attendeeRouter.PUT("/:id", utility.Use(ac.updateAttendee))
		attendeeRouter.DELETE("/:id", utility.Use(ac.removeAttendee))
		attendeeRouter.GET("/event/:eventID", utility.Use(ac.getAttendeesByEvent))
	}
}

// addAttendee adds an attendee to an event
// @Summary Add attendee to an event
// @Description Add an attendee to the specified event
// @Tags Attendee
// @Accept json
// @Produce json
// @Param AddAttendeeDto body models.Attendee true "Attendee Data"
// @Success 201 {object} models.Attendee "Successfully added attendee"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/attendees/ [post]
func (ac *AttendeeController) addAttendee(ctx *gin.Context) error {
	data := new(models.Attendee)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
	}

	attendee := &models.Attendee{
		EventID: data.EventID,
		UserID:  data.UserID,
	}

	newAttendee := ac.attendeeService.AddAttendee(attendee)
	if newAttendee == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return errors.New("internal server error")
	}

	ctx.JSON(http.StatusCreated, newAttendee)
	return nil
}

// getAttendee retrieves an attendee by their ID
// @Summary Get attendee by ID
// @Description Get attendee details by their ID
// @Tags Attendee
// @Accept json
// @Produce json
// @Param id path string true "Attendee ID"
// @Success 200 {object} models.Attendee "Successfully retrieved attendee"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Attendee Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/attendees/{id} [get]
func (ac *AttendeeController) getAttendee(ctx *gin.Context) error {
	attendeeID := ctx.Param("id")
	attendee, err := ac.attendeeService.GetAttendeesByEvent(attendeeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Attendee Not Found",
			})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
	}

	ctx.JSON(http.StatusOK, attendee)
	return nil
}

// updateAttendee updates an attendee's details
// @Summary Update attendee details
// @Description Update an attendee's details by their ID
// @Tags Attendee
// @Accept json
// @Produce json
// @Param id path string true "Attendee ID"
// @Param UpdateAttendeeDto body models.Attendee true "Attendee Data"
// @Success 200 {object} models.Attendee "Successfully updated attendee"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Attendee Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/attendees/{id} [put]
func (ac *AttendeeController) updateAttendee(ctx *gin.Context) error {
	attendeeID := ctx.Param("id")
	data := new(models.Attendee)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
	}
	attendee, err := ac.attendeeService.GetAttendeesByEvent(attendeeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Attendee Not Found",
			})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	//attendee.EventID = data.EventID
	//attendee.UserID = data.UserID
	//
	//updatedAttendee, err := ac.attendeeService.UpdateAttendee(attendee)
	//if err != nil {
	//	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Internal Server Error",
	//	})
	//}

	ctx.JSON(http.StatusOK, attendee)
	return nil
}

// removeAttendee removes an attendee from an event
// @Summary Remove attendee from an event
// @Description Remove an attendee from the specified event by their ID
// @Tags Attendee
// @Accept json
// @Produce json
// @Param id path string true "Attendee ID"
// @Success 204 "Successfully removed attendee"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Attendee Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/attendees/{id} [delete]
func (ac *AttendeeController) removeAttendee(ctx *gin.Context) error {
	attendeeID := ctx.Param("id")
	err := ac.attendeeService.RemoveAttendee(attendeeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "Attendee Not Found",
			})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
	}
	ctx.JSON(http.StatusNoContent, nil)
	return nil
}

// getAttendeesByEvent retrieves all attendees for a specific event
// @Summary Get attendees by event
// @Description Get all attendees for a specific event by event ID
// @Tags Attendee
// @Accept json
// @Produce json
// @Param eventId path string true "Event ID"
// @Success 200 {array} models.Attendee "Successfully retrieved attendees"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/events/{eventId}/attendees [get]
func (ac *AttendeeController) getAttendeesByEvent(ctx *gin.Context) error {
	eventID := ctx.Param("eventId")
	attendees, err := ac.attendeeService.GetAttendeesByEvent(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, attendees)
	return nil
}
