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

func (c *PlaceController) RegisterRoutes(router, routerWithoutAuth *gin.RouterGroup) {
	placeRouter := router.Group("/places")
	placeRouterWithoutAuth := routerWithoutAuth.Group("/places")
	{
		placeRouterWithoutAuth.GET("/:id", utility.Use(c.getPlace))
		placeRouter.POST("/", utility.Use(c.createPlace))
		placeRouterWithoutAuth.GET("/", utility.Use(c.getPlacesByFilters))
		placeRouter.PATCH("/:id", utility.Use(c.updatePlace))
		placeRouter.DELETE("/:id", utility.Use(c.deletePlace))
		placeRouter.POST("/:id/amenities", utility.Use(c.addAmenitiesToPlace))
		placeRouterWithoutAuth.GET("/all", utility.Use(c.getAllPlaces))
	}
}

// @Summary Get a place
// @Description Get a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to get"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully retrieved place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id} [get]
func (c *PlaceController) getPlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	place, err := c.placeService.GetPlace(ctx, id)
	if err != nil {

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
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully created place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/ [post]
func (c *PlaceController) createPlace(ctx *gin.Context) error {
	var placeData Dto.CreatePlaceDTO
	if err := ctx.ShouldBindJSON(&placeData); err != nil {

		return err
	}

	placeData.BusinessID = ctx.Query("business_id")
	if placeData.BusinessID == "" {

		return nil
	}

	place, err := c.placeService.CreatePlace(ctx, placeData)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, place)
	return nil
}

// @Summary Get all places
// @Description Get all places
// @Tags Place
// @Accept json
// @Produce json
// @Param nextPageToken query string false "Token for the next page of results"
// @Param limit query int false "Number of results to return"
// @Success 200 {array} []ent.Place "Successfully retrieved all places"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/all [get]
func (c *PlaceController) getAllPlaces(ctx *gin.Context) error {
	nextPageToken := ctx.Query("nextPageToken")
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	places, token, err := c.placeService.GetAllPlaces(ctx, nextPageToken, limit)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(places, "success", "places retrieved successfully", token))
	return nil
}

// @Summary Add amenities to a place
// @Description Add amenities to a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to add amenities to"
// @Param amenity body Dto.AmenityAdditionDTO true "Amenities to add"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully added amenities to place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id}/amenities [post]
func (c *PlaceController) addAmenitiesToPlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	var amenityDTO Dto.AmenityAdditionDTO
	if err := ctx.ShouldBindJSON(&amenityDTO); err != nil {

		return err
	}

	if err := c.placeService.AddAmenitiesToPlace(ctx, id, amenityDTO.AmenityIDs); err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Amenities added successfully",
	})
	return nil
}

// @Summary Update a place
// @Description Update a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to update"
// @Param place body Dto.UpdatePlaceDTO true "Place data to update"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully updated place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id} [patch]
func (c *PlaceController) updatePlace(ctx *gin.Context) error {
	id := ctx.Param("id")
	var placeData Dto.UpdatePlaceDTO
	if err := ctx.ShouldBindJSON(&placeData); err != nil {

		return err
	}

	place, err := c.placeService.UpdatePlace(ctx, id, placeData)
	if err != nil {

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
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully deleted place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id} [delete]
func (c *PlaceController) deletePlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	err := c.placeService.DeletePlace(ctx, id)
	if err != nil {

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

		return err
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {

		return err
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {

		return err
	}

	places, err := c.placeService.GetPlaces(ctx, &filter, page, pageSize)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, places)
	return nil
}
