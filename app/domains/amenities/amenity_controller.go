package amenities

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"placio-app/domains/media"
	_ "placio-app/ent"
	"placio-app/utility"
)

type AmenityController struct {
	service      AmenityService
	mediaService media.MediaService
	cache        *utility.RedisClient
}

func NewAmenityController(s AmenityService, cache *utility.RedisClient) *AmenityController {
	return &AmenityController{service: s, cache: cache}
}

func (c *AmenityController) RegisterRoutes(r *gin.RouterGroup) {
	amenityRouterGroup := r.Group("/amenities")
	amenityRouterGroup.POST("/", utility.Use(c.createAmenity))
	amenityRouterGroup.GET("/:id", utility.Use(c.getAmenity))
	amenityRouterGroup.PUT("/:id", utility.Use(c.updateAmenity))
	amenityRouterGroup.DELETE("/:id", utility.Use(c.deleteAmenity))
	amenityRouterGroup.GET("/", utility.Use(c.getAllAmenities))
}

// @Summary Create a new amenity
// @Description Create a new amenity with provided information
// @Tags Amenity
// @Accept json
// @Produce json
// @Param icons formData file true "Icon files for amenity"
// @Param amenity body Dto.CreateAmenityInput true "Amenity information"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Amenity "Successfully created amenity"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/amenities [post]
func (c *AmenityController) createAmenity(ctx *gin.Context) error {
	var input CreateAmenityInput
	if err := ctx.ShouldBindJSON(&input); err != nil {

		return err
	}

	amenity, err := c.service.CreateAmenity(input)
	if err != nil {

		return err
	}

	amenityJSON, _ := json.Marshal(amenity)
	c.cache.SetCache(ctx, amenity.ID, string(amenityJSON))
	c.cache.DeleteCache(ctx, "all_amenities")

	ctx.JSON(http.StatusOK, amenity)
	return nil
}

// @Summary Get an amenity
// @Description Get an amenity by ID
// @Tags Amenity
// @Accept json
// @Produce json
// @Param id path string true "ID of the amenity"
// @Success 200 {object} ent.Amenity "Successfully retrieved amenity"
// @Failure 404 {object} Dto.ErrorDTO "Amenity not found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/amenities/{id} [get]
func (c *AmenityController) getAmenity(ctx *gin.Context) error {
	id := ctx.Param("id")

	amenity, err := c.service.GetAmenity(id)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, amenity)
	return nil
}

// @Summary Update an amenity
// @Description Update an amenity by ID
// @Tags Amenity
// @Accept json
// @Produce json
// @Param id path string true "ID of the amenity to update"
// @Param icon formData file true "New icon file for amenity"
// @Param amenity body Dto.UpdateAmenityInput true "New amenity information"
// @Success 200 {object} ent.Amenity "Successfully updated amenity"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Amenity not found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/amenities/{id} [put]
func (c *AmenityController) updateAmenity(ctx *gin.Context) error {
	var input UpdateAmenityInput
	if err := ctx.ShouldBindJSON(&input); err != nil {

		return err
	}

	id := ctx.Param("id")

	file, err := ctx.FormFile("icon")
	if err == nil {
		// Save the file to a temporary path before uploading
		tempFilePath := "/tmp/" + file.Filename
		ctx.SaveUploadedFile(file, tempFilePath)

		//uploadParams := uploader.UploadParams{Folder: "your/folder"}
		mediaInfo, err := c.mediaService.UploadFiles(ctx, []*multipart.FileHeader{file})
		if err != nil {

			return err
		}

		// Assign the uploaded file's URL to the Icon field
		input.Icon = &mediaInfo[0].URL
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return err
	}

	amenity, err := c.service.UpdateAmenity(id, input)
	if err != nil {

		return err
	}

	amenityJSON, _ := json.Marshal(amenity)
	c.cache.SetCache(ctx, id, string(amenityJSON))
	c.cache.DeleteCache(ctx, "all_amenities")

	ctx.JSON(http.StatusOK, amenity)
	return nil
}

// @Summary Delete an amenity
// @Description Delete an amenity by ID
// @Tags Amenity
// @Accept json
// @Produce json
// @Param id path string true "ID of the amenity to delete"
// @Param Authorization header string true "JWT token"
// @Success 200 {object} string "Successfully deleted amenity"
// @Failure 404 {object} Dto.ErrorDTO "Amenity not found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/amenities/{id} [delete]
func (c *AmenityController) deleteAmenity(ctx *gin.Context) error {
	id := ctx.Param("id")

	err := c.service.DeleteAmenity(id)
	if err != nil {

		return err
	}

	c.cache.DeleteCache(ctx, id)
	c.cache.DeleteCache(ctx, "all_amenities")

	ctx.JSON(http.StatusOK, gin.H{"success": "Amenity deleted successfully"})
	return nil
}

// @Summary Get all amenities
// @Description Get all amenities
// @Tags Amenity
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT token"
// @Success 200 {object} []ent.Amenity "Successfully retrieved amenities"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Failure 401 {object} Dto.ErrorDTO
// @Router /api/v1/amenities [get]
func (c *AmenityController) getAllAmenities(ctx *gin.Context) error {
	amenities, err := c.service.GetAllAmenities()
	if err != nil {

		return err
	}

	amenitiesJSON, _ := json.Marshal(amenities)
	c.cache.SetCache(ctx, "all_amenities", string(amenitiesJSON))

	ctx.JSON(http.StatusOK, amenities)
	return nil
}
