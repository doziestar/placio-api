package places

import (
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/domains/amenities"
	"placio-app/ent"
	_ "placio-app/ent"
	"placio-app/utility"
	"placio-pkg/errors"
	"strconv"
)

type PlaceController struct {
	placeService PlaceService
	cache        utility.RedisClient
}

func NewPlaceController(placeService PlaceService, cache utility.RedisClient) *PlaceController {
	return &PlaceController{placeService: placeService, cache: cache}
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
		placeRouter.POST("/:id/media", utility.Use(c.addMediaToAPlace))
		placeRouter.DELETE("/:id/media", utility.Use(c.removeMediaToAPlace))
		placeRouterWithoutAuth.GET("/all", utility.Use(c.getAllPlaces))
		placeRouter.POST("/:id/remove_amenities", utility.Use(c.removeAmenitiesFromPlace))
		placeRouter.DELETE("/:id/media/:mediaID", utility.Use(c.removeSingleMediaFromPlace))

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
	// get place from cache
	cacheKey := "place:" + id

	bytes, err := c.cache.GetCache(ctx, cacheKey)
	if err != nil {
		// if the error is redis: nil, just ignore it and fetch from the db
		if err.Error() != "redis: nil" {
			sentry.CaptureException(err)
			return err
		}
	}

	if bytes != nil {
		var data *ent.Place
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			sentry.CaptureException(err)
			return err
		}
		ctx.JSON(http.StatusOK, data)
		return nil
	}

	data, err := c.placeService.GetPlace(ctx, id)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	ctx.JSON(http.StatusOK, data)
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
	var placeData CreatePlaceDTO
	if err := ctx.ShouldBindJSON(&placeData); err != nil {

		return err
	}

	placeData.BusinessID = ctx.Query("business_id")
	if placeData.BusinessID == "" {
		return errors.IDMissing
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

	placesData, token, err := c.placeService.GetAllPlaces(ctx, nextPageToken, limit)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(placesData, "success", "places retrieved successfully", token))
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

	var amenityDTO []amenities.CreateAmenityInput
	if err := ctx.ShouldBindJSON(&amenityDTO); err != nil {
		return err
	}

	if err := c.placeService.AddAmenitiesToPlace(ctx, id, amenityDTO); err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Amenities added successfully",
	})
	return nil
}

// @Summary Remove amenities to a place
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
func (c *PlaceController) removeAmenitiesFromPlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	var amenityDTO amenities.AmenityAdditionDTO
	if err := ctx.ShouldBindJSON(&amenityDTO); err != nil {
		return err
	}

	if err := c.placeService.RemoveAmenitiesFromPlace(ctx, id, amenityDTO.AmenityIDs); err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Amenities removed successfully",
	})
	return nil
}

// @Summary Add Media to a place
// @Description Add media to a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to add amenities to"
// @Param amenity body  true ent.Media "Media to add"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully added amenities to place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id}/media [post]
func (c *PlaceController) addMediaToAPlace(ctx *gin.Context) error {
	id := ctx.Param("id")

	forms, err := ctx.MultipartForm()
	if err != err {
		sentry.CaptureException(err)
		return err
	}
	log.Println("form gotten", forms)

	files, ok := forms.File["files"]
	if !ok || len(files) == 0 {
		sentry.CaptureException(errors.New("files could not be extracted from the places form"))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "files could not be extracted from the places form",
		})
		return err
	}

	log.Println("calling AddMediaToPlace", files)
	if err := c.placeService.AddMediaToPlace(ctx, id, files); err != nil {
		log.Print(err)
		sentry.CaptureException(err)
	}
	log.Println("upload complete")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Media added successfully",
	})
	return nil
}

// @Summary Remove media from a place
// @Description remove media to a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to add amenities to"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully added amenities to place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id}/media [post]
func (c *PlaceController) removeSingleMediaFromPlace(ctx *gin.Context) error {
	placeID := ctx.Param("id")
	mediaID := ctx.Param("mediaID")

	if err := c.placeService.RemoveMediaFromPlace(ctx, placeID, mediaID); err != nil {
		log.Print(err)
		sentry.CaptureException(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to remove media",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Media removed successfully",
	})
	return nil
}

// @Summary Remove Media to a place
// @Description Add media to a place by ID
// @Tags Place
// @Accept json
// @Produce json
// @Param id path string true "ID of the place to add amenities to"
// @Param amenity body  true ent.Media "Media to add"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Place "Successfully added amenities to place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/places/{id}/media [post]
func (c *PlaceController) removeMediaToAPlace(ctx *gin.Context) error {
	type requestBody struct {
		MediaIDs []string `json:"media_ids"`
	}

	var body requestBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return err
	}

	// The place ID can be a path parameter, as before
	placeID := ctx.Param("id")

	go func() {
		if err := c.placeService.RemoveMediaToPlace(ctx, placeID, body.MediaIDs); err != nil {
			sentry.CaptureException(err)
		}
	}()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Media removed successfully",
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
	var placeData UpdatePlaceDTO
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
// @Param filter query places.PlaceFilter false "Filter"
// @Param page query int false "Page Number"
// @Param pageSize query int false "limit"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {array} ent.Place
// @Failure 400 {object} Dto.ErrorDTO "Invalid inputs"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized. Invalid or expired token"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /places [get]
func (c *PlaceController) getPlacesByFilters(ctx *gin.Context) error {
	var filter PlaceFilter
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		return err
	}

	page := ctx.Query("nextPageToken")

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {

		return err
	}

	places, nextPageToken, err := c.placeService.GetPlaces(ctx, &filter, page, limit)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(places, "success", "places retrieved successfully", nextPageToken))
	return nil
}
