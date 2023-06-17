package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
)

type ReservationController struct {
	reservationService service.ReservationService
}

func NewReservationController(reservationService service.ReservationService) *ReservationController {
	return &ReservationController{reservationService: reservationService}
}

func (rc *ReservationController) RegisterRoutes(router *gin.RouterGroup) {
	reservationRouter := router.Group("/reservations")
	{
		reservationRouter.GET("/:id", rc.GetReservation)
		reservationRouter.POST("/", rc.CreateReservation)
		reservationRouter.PATCH("/:id", rc.UpdateReservation)
		reservationRouter.DELETE("/:id", rc.DeleteReservation)
	}
}

// @Summary Get a reservation
// @Description Get a reservation by ID
// @Tags Reservation
// @Accept json
// @Produce json
// @Param id path string true "ID of the reservation to get"
// @Security Bearer
// @Success 200 {object} ent.Reservation "Successfully retrieved reservation"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/reservations/{id} [get]
func (rc *ReservationController) GetReservation(ctx *gin.Context) {
	id := ctx.Param("id")

	reservation, err := rc.reservationService.GetReservation(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, reservation)
}

// @Summary Create a reservation
// @Description Create a new reservation
// @Tags Reservation
// @Accept json
// @Produce json
// @Param reservation body ent.Reservation true "Reservation to create"
// @Security Bearer
// @Success 200 {object} ent.Reservation "Successfully created reservation"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/reservations/ [post]
func (rc *ReservationController) CreateReservation(ctx *gin.Context) {
	var reservationData map[string]interface{}
	if err := ctx.ShouldBindJSON(&reservationData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservation, err := rc.reservationService.CreateReservation(ctx, reservationData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, reservation)
}

// @Summary Update a reservation
// @Description Update a reservation by ID
// @Tags Reservation
// @Accept json
// @Produce json
// @Param id path string true "ID of the reservation to update"
// @Param reservation body ent.Reservation true "Reservation data to update"
// @Security Bearer
// @Success 200 {object} ent.Reservation "Successfully updated reservation"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/reservations/{id} [patch]
func (rc *ReservationController) UpdateReservation(ctx *gin.Context) {
	id := ctx.Param("id")
	var reservationData map[string]interface{}
	if err := ctx.ShouldBindJSON(&reservationData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservation, err := rc.reservationService.UpdateReservation(ctx, id, reservationData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, reservation)
}

// @Summary Delete a reservation
// @Description Delete a reservation by ID
// @Tags Reservation
// @Accept json
// @Produce json
// @Param id path string true "ID of the reservation to delete"
// @Security Bearer
// @Success 200 {object} ent.Reservation "Successfully deleted reservation"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/reservations/{id} [delete]
func (rc *ReservationController) DeleteReservation(ctx *gin.Context) {
	id := ctx.Param("id")

	err := rc.reservationService.DeleteReservation(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully deleted reservation",
	})
}
