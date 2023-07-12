package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/Dto"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"
	"strconv"
)

type PlaceController struct {
	placeService service.PlaceService
}

func NewPlaceController(placeService service.PlaceService) *PlaceController {
	return &PlaceController{placeService: placeService}
}

func (c *PlaceController) RegisterRoutes(router *gin.RouterGroup) {
	placeRouter := router.Group("/places")
	{
		placeRouter.GET("/:id", utility.Use(c.getPlace))
		placeRouter.POST("/", utility.Use(c.createPlace))
		placeRouter.GET("/", utility.Use(c.getPlacesByFilters))
		placeRouter.PATCH("/:id", utility.Use(c.updatePlace))
		placeRouter.DELETE("/:id", utility.Use(c.deletePlace))
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
func (c *PlaceController) getPlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	place, err := c.placeService.GetPlace(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return err
	}

	ctx.JSON(http.StatusOK, place)
	return nil
}

// @Summary Create a place
// @Description Create a new place
// @Tags Place
// @Accept json
// @Produce json
// @Param place body Dto.CreatePlaceDTO true "Place to create"
// @Param business_id query string true "ID of the business to create the place for"
// @Security Bearer
// @Success 200 {object} ent.Place "Successfully created place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/ [post]
func (c *PlaceController) createPlace(ctx *gin.Context) error {
	var placeData Dto.CreatePlaceDTO
	if err := ctx.ShouldBindJSON(&placeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	placeData.BusinessID = ctx.Query("business_id")
	if placeData.BusinessID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Business ID is required",
		})
		return nil
	}

	place, err := c.placeService.CreatePlace(ctx, placeData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return err
	}

	ctx.JSON(http.StatusOK, place)
	return nil
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
func (c *PlaceController) updatePlace(ctx *gin.Context) error {
	id := ctx.Param("id")
	var placeData Dto.UpdatePlaceDTO
	if err := ctx.ShouldBindJSON(&placeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	place, err := c.placeService.UpdatePlace(ctx, id, placeData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, place)
	return nil
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
func (c *PlaceController) deletePlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	err := c.placeService.DeletePlace(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully deleted place",
	})
	return nil
}

// GetPlaces godoc
// @Summary Get all Places
// @Description Get Places by applying various filters (ID, Name, Type, Country, City, State, Tags, Features)
// @Tags Place
// @Accept  json
// @Produce  json
// @Param filter query service.PlaceFilter false "Filter"
// @Param page query int false "Page Number"
// @Param pageSize query int false "Page Size"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {array} ent.Place
// @Failure 400 {object} Dto.ErrorDTO "Invalid inputs"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized. Invalid or expired token"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /places [get]
func (c *PlaceController) getPlacesByFilters(ctx *gin.Context) error {
	var filter service.PlaceFilter
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return err
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return err
	}

	places, err := c.placeService.GetPlaces(ctx, &filter, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, places)
	return nil
}
