package feature_releases

import (
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/ent"
	"placio-app/utility"
	"strconv"
)

type FeatureReleaseController struct {
	featureService FeatureReleaseService
	cache          utility.RedisClient
}

func NewFeatureReleaseController(featureService FeatureReleaseService, cache utility.RedisClient) *FeatureReleaseController {
	return &FeatureReleaseController{featureService: featureService, cache: cache}
}

func (c *FeatureReleaseController) RegisterRoutes(router, routerWithoutAuth *gin.RouterGroup) {
	featureRouter := router.Group("/features")
	featureRouterWithoutAuth := routerWithoutAuth.Group("/features")
	{
		featureRouter.GET("/:id", utility.Use(c.getFeature))
		featureRouter.POST("/", utility.Use(c.createFeature))
		featureRouterWithoutAuth.GET("/", utility.Use(c.listFeatures))
		featureRouter.PATCH("/:id", utility.Use(c.updateFeature))
		featureRouter.DELETE("/:id", utility.Use(c.deleteFeature))
		featureRouter.PATCH("/:id/state", utility.Use(c.setFeatureState))
	}
}

// @Summary Get a feature release
// @Description Get a feature by ID
// @Tags Feature
// @Accept json
// @Produce json
// @Param id path string true "ID of the feature to get"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.FeatureRelease "Successfully retrieved feature"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/features/{id} [get]
func (c *FeatureReleaseController) getFeature(ctx *gin.Context) error {
	id := ctx.Param("id")
	// get feature from cache
	cacheKey := "feature:" + id

	bytes, err := c.cache.GetCache(ctx, cacheKey)
	if err != nil {
		// if the error is redis: nil, just ignore it and fetch from the db
		if err.Error() != "redis: nil" {
			sentry.CaptureException(err)
			return err
		}
	}

	if bytes != nil {
		var data *ent.FeatureRelease
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			sentry.CaptureException(err)
			return err
		}
		ctx.JSON(http.StatusOK, data)
		return nil
	}

	data, err := c.featureService.GetFeature(ctx, id)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}

	ctx.JSON(http.StatusOK, data)
	return nil

}

// createFeature creates a new feature release.
func (c *FeatureReleaseController) createFeature(ctx *gin.Context) error {
	var featureData FeatureReleaseDTO
	if err := ctx.ShouldBindJSON(&featureData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload."})
		return nil
	}

	newFeature, err := c.featureService.CreateFeature(ctx, featureData)
	if err != nil {
		sentry.CaptureException(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create feature."})
		return nil
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(newFeature))
	return nil
}

// listFeatures lists all feature releases with optional pagination.
func (c *FeatureReleaseController) listFeatures(ctx *gin.Context) error {
	lastID := ctx.DefaultQuery("lastId", "")
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	features, nextID, err := c.featureService.ListFeatures(ctx, lastID, limit)
	if err != nil {
		sentry.CaptureException(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list features."})
		return nil
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(features, "success", "features list", nextID))
	return nil
}

// updateFeature updates an existing feature release.
func (c *FeatureReleaseController) updateFeature(ctx *gin.Context) error {
	id := ctx.Param("id")

	var featureData FeatureReleaseDTO
	if err := ctx.ShouldBindJSON(&featureData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload."})
		return nil
	}

	updatedFeature, err := c.featureService.UpdateFeature(ctx, id, featureData)
	if err != nil {
		sentry.CaptureException(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update feature."})
		return nil
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(updatedFeature))
	return nil
}

// deleteFeature deletes a feature release.
func (c *FeatureReleaseController) deleteFeature(ctx *gin.Context) error {
	id := ctx.Param("id")

	if err := c.featureService.DeleteFeature(ctx, id); err != nil {
		sentry.CaptureException(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete feature."})
		return nil
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Feature deleted successfully."})
	return nil
}

// setFeatureState updates the state of a feature.
func (c *FeatureReleaseController) setFeatureState(ctx *gin.Context) error {
	id := ctx.Param("id")

	type stateDTO struct {
		State string `json:"state" binding:"required"`
	}

	var state stateDTO
	if err := ctx.ShouldBindJSON(&state); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload."})
		return nil
	}
	if err := c.featureService.SetFeatureState(ctx, id, state.State); err != nil {
		sentry.CaptureException(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set feature state."})
		return nil
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Feature state updated successfully."})
	return nil
}
