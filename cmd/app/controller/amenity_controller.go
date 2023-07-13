package controller

import (
	"encoding/json"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/Dto"
	"placio-app/service"
	"placio-app/utility"
	"time"
)

type AmenityController struct {
	service      service.AmenityService
	mediaService service.MediaService
	cache        *utility.RedisClient
}

func NewAmenityController(s service.AmenityService, cache *utility.RedisClient) *AmenityController {
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

func (c *AmenityController) createAmenity(ctx *gin.Context) error {
	var input Dto.CreateAmenityInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	form, _ := ctx.MultipartForm()
	files := form.File["icons"]

	filePaths := make([]string, len(files))
	for i, file := range files {
		// Save the file to a temporary path before uploading
		tempFilePath := "/tmp/" + file.Filename
		ctx.SaveUploadedFile(file, tempFilePath)
		filePaths[i] = tempFilePath
	}

	uploadParams := uploader.UploadParams{Folder: "your/folder"}
	urls, err := c.mediaService.UploadFiles(ctx, filePaths, uploadParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	// If you're storing multiple icon URLs, you should modify the input and Amenity structures accordingly
	input.Icon = urls[0].URL

	amenity, err := c.service.CreateAmenity(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	amenityJSON, _ := json.Marshal(amenity)
	c.cache.SetCache(ctx, amenity.ID, string(amenityJSON), 1*time.Hour)
	c.cache.DeleteCache(ctx, "all_amenities")

	ctx.JSON(http.StatusOK, amenity)
	return nil
}

func (c *AmenityController) getAmenity(ctx *gin.Context) error {
	id := ctx.Param("id")

	amenity, err := c.service.GetAmenity(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Amenity not found"})
		return err
	}

	ctx.JSON(http.StatusOK, amenity)
	return nil
}

func (c *AmenityController) updateAmenity(ctx *gin.Context) error {
	var input Dto.UpdateAmenityInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	id := ctx.Param("id")

	file, err := ctx.FormFile("icon")
	if err == nil {
		// Save the file to a temporary path before uploading
		tempFilePath := "/tmp/" + file.Filename
		ctx.SaveUploadedFile(file, tempFilePath)

		uploadParams := uploader.UploadParams{Folder: "your/folder"}
		mediaInfo, err := c.mediaService.UploadFiles(ctx, []string{tempFilePath}, uploadParams)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Amenity not found"})
		return err
	}

	amenityJSON, _ := json.Marshal(amenity)
	c.cache.SetCache(ctx, id, string(amenityJSON), 1*time.Hour)
	c.cache.DeleteCache(ctx, "all_amenities")

	ctx.JSON(http.StatusOK, amenity)
	return nil
}

func (c *AmenityController) deleteAmenity(ctx *gin.Context) error {
	id := ctx.Param("id")

	err := c.service.DeleteAmenity(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Amenity not found"})
		return err
	}

	c.cache.DeleteCache(ctx, id)
	c.cache.DeleteCache(ctx, "all_amenities")

	ctx.JSON(http.StatusOK, gin.H{"result": "Amenity deleted successfully"})
	return nil
}

func (c *AmenityController) getAllAmenities(ctx *gin.Context) error {
	amenities, err := c.service.GetAllAmenities()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	amenitiesJSON, _ := json.Marshal(amenities)
	c.cache.SetCache(ctx, "all_amenities", string(amenitiesJSON), 1*time.Hour)

	ctx.JSON(http.StatusOK, amenities)
	return nil
}
