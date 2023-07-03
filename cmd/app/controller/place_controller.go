package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/Dto"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
)

type PlaceController struct {
	placeService service.PlaceService
}

func NewPlaceController(placeService service.PlaceService) *PlaceController {
	return &PlaceController{placeService: placeService}
}

func (pc *PlaceController) RegisterRoutes(router *gin.RouterGroup) {
	placeRouter := router.Group("/places")
	{
		placeRouter.GET("/:id", pc.getPlace)
		placeRouter.POST("/", pc.createPlace)
		placeRouter.PATCH("/:id", pc.updatePlace)
		placeRouter.DELETE("/:id", pc.deletePlace)
	}
}

// @Summary Get a place
// @Description Get a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to get"
// @Security Bearer
// @Success 200 {object} ent.Place "Successfully retrieved place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id} [get]
func (pc *PlaceController) getPlace(ctx *gin.Context) {
	id := ctx.Param("id")

	place, err := pc.placeService.GetPlace(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, place)
}

// @Summary Create a place
// @Description Create a new place
// @Tags Place
// @Accept json
// @Produce json
// @Param place body Dto.CreatePlaceDTO true "Place to create"
// @Security Bearer
// @Success 200 {object} ent.Place "Successfully created place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/ [post]
func (pc *PlaceController) createPlace(ctx *gin.Context) {
	var placeData Dto.CreatePlaceDTO
	if err := ctx.ShouldBindJSON(&placeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	place, err := pc.placeService.CreatePlace(ctx, placeData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, place)
}

// @Summary Update a place
// @Description Update a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to update"
// @Param place body Dto.UpdatePlaceDTO true "Place data to update"
// @Security Bearer
// @Success 200 {object} ent.Place "Successfully updated place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id} [patch]
func (pc *PlaceController) updatePlace(ctx *gin.Context) {
	id := ctx.Param("id")
	var placeData Dto.UpdatePlaceDTO
	if err := ctx.ShouldBindJSON(&placeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	place, err := pc.placeService.UpdatePlace(ctx, id, placeData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, place)
}

// @Summary Delete a place
// @Description Delete a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to delete"
// @Security Bearer
// @Success 200 {object} ent.Place "Successfully deleted place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id} [delete]
func (pc *PlaceController) deletePlace(ctx *gin.Context) {
	id := ctx.Param("id")

	err := pc.placeService.DeletePlace(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully deleted place",
	})
}
