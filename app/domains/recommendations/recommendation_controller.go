package recommendations

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/utility"
	"placio-pkg/middleware"
)

type RecommendationController struct {
	recommendationService RecommendationService
}

func NewRecommendationController(recommendationService RecommendationService) *RecommendationController {
	return &RecommendationController{recommendationService: recommendationService}
}

func (rc *RecommendationController) RegisterRoutes(router *gin.RouterGroup) {
	recommendationRouter := router.Group("/recommendations")
	{
		recommendationRouter.GET("/places", middleware.ErrorMiddleware(rc.getPlacesRecommendations))
		recommendationRouter.GET("/restaurants", middleware.ErrorMiddleware(rc.getRestaurantsRecommendations))
		recommendationRouter.GET("/hotels", middleware.ErrorMiddleware(rc.getHotelsRecommendations))
		recommendationRouter.GET("/inventory", middleware.ErrorMiddleware(rc.getInventoryRecommendations))
	}
}

// GetPlacesRecommendations returns a list of places recommendations.
// @Summary Get places recommendations
// @Description Get places recommendations for the authenticated user
// @Tags Recommendation
// @Accept json
// @Produce json
// @Success 200 {object} []PlaceResponseDto "Successfully retrieved places recommendations"
// @Failure 400 {object} ErrorDTO "Bad Request"
func (rc *RecommendationController) getPlacesRecommendations(c *gin.Context) error {
	places, err := rc.recommendationService.GetPlacesRecommendations(c)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, utility.ProcessResponse(places))
	return nil
}

// GetRestaurantsRecommendations returns a list of restaurants recommendations.
// @Summary Get restaurants recommendations
// @Description Get restaurants recommendations for the authenticated user
// @Tags Recommendation
// @Accept json
// @Produce json
// @Success 200 {object} []PlaceResponseDto "Successfully retrieved restaurants recommendations"
// @Failure 400 {object} ErrorDTO "Bad Request"
func (rc *RecommendationController) getRestaurantsRecommendations(c *gin.Context) error {
	restaurants, err := rc.recommendationService.GetRestaurantsRecommendations(c)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, utility.ProcessResponse(restaurants))
	return nil
}

// GetHotelsRecommendations returns a list of hotels recommendations.
// @Summary Get hotels recommendations
// @Description Get hotels recommendations for the authenticated user
// @Tags Recommendation
// @Accept json
// @Produce json
// @Success 200 {object} []PlaceResponseDto "Successfully retrieved hotels recommendations"
// @Failure 400 {object} ErrorDTO "Bad Request"
func (rc *RecommendationController) getHotelsRecommendations(c *gin.Context) error {
	hotels, err := rc.recommendationService.GetHotelsRecommendations(c)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, utility.ProcessResponse(hotels))
	return nil
}

// GetInventoryRecommendations returns a list of inventory recommendations.
// @Summary Get inventory recommendations
// @Description Get inventory recommendations for the authenticated user
// @Tags Recommendation
// @Accept json
// @Produce json
// @Success 200 {object} []InventoryResponseDto "Successfully retrieved inventory recommendations"
// @Failure 400 {object} ErrorDTO "Bad Request"
func (rc *RecommendationController) getInventoryRecommendations(c *gin.Context) error {
	inventory, err := rc.recommendationService.GetInventoryRecommendations(c)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, utility.ProcessResponse(inventory))
	return nil
}
