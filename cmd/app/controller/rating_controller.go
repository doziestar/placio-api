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

type RatingController struct {
	ratingService service.RatingService
}

func NewRatingController(ratingService service.RatingService) *RatingController {
	return &RatingController{ratingService: ratingService}
}

func (rc *RatingController) RegisterRoutes(router *gin.RouterGroup) {
	ratingRouter := router.Group("/ratings")
	{
		ratingRouter.POST("/", utility.Use(rc.createRating))
		ratingRouter.GET("/:id", utility.Use(rc.getRating))
		ratingRouter.PUT("/:id", utility.Use(rc.updateRating))
		ratingRouter.DELETE("/:id", utility.Use(rc.deleteRating))
		ratingRouter.GET("/event/:eventID", utility.Use(rc.getRatingsByEvent))
	}
}

// createRating creates a new rating
// @Summary Create a new rating
// @Description Create a new rating for the specified event
// @Tags Rating
// @Accept json
// @Produce json
// @Param CreateRatingDto body models.Rating true "Rating Data"
// @Success 201 {object} models.Rating "Successfully created rating"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings/ [post]
func (rc *RatingController) createRating(ctx *gin.Context) error {
	data := new(models.Rating)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	rating := &models.Rating{
		EventID: data.EventID,
		UserID:  data.UserID,
		//RateValue: data.RateValue,
	}

	newRating := rc.ratingService.CreateRating(rating)
	if newRating == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return errors.New("internal server error")
	}

	ctx.JSON(http.StatusCreated, newRating)
	return nil

}

// getRating retrieves a rating by its ID
// @Summary GET rating by ID
// @Description Retrieve a rating by its ID
// @Tags Rating
// @Accept json
// @Produce json
// @Param id path string true "Rating ID"
// @Success 200 {object} models.Rating "Successfully retrieved rating"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Rating Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings/{id} [get]
func (rc *RatingController) getRating(ctx *gin.Context) error {
	ratingID := ctx.Param("id")
	rating, err := rc.ratingService.GetRatingsByEvent(ratingID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Rating Not Found"})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	ctx.JSON(http.StatusOK, rating)
	return nil
}

// updateRating updates a rating by its ID
// @Summary Update rating by ID
// @Description Update a rating by its ID
// @Tags Rating
// @Accept json
// @Produce json
// @Param id path string true "Rating ID"
// @Param UpdateRatingDto body models.Rating true "Rating Data"
// @Success 200 {object} models.Rating "Successfully updated rating"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Rating Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings/{id} [put]
func (rc *RatingController) updateRating(ctx *gin.Context) error {
	ratingID := ctx.Param("id")
	data := new(models.Rating)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	_, err := rc.ratingService.GetRatingsByEvent(ratingID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Rating Not Found"})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	//rating.RateValue = data.RateValue
	//updatedRating, err := rc.ratingService.UpdateRating(rating)
	//if err != nil {
	//	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Internal Server Error",
	//	})
	//}

	ctx.JSON(http.StatusOK, data)
	return nil
}

// deleteRating deletes a rating by its ID
// @Summary Delete rating by ID
// @Description Delete a rating by its ID
// @Tags Rating
// @Accept json
// @Produce json
// @Param id path string true "Rating ID"
// @Success 204 "Successfully deleted rating"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Rating Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings/{id} [delete]
func (rc *RatingController) deleteRating(ctx *gin.Context) error {
	ratingID := ctx.Param("id")
	err := rc.ratingService.DeleteRating(ratingID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Rating Not Found"})
			return err
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusNoContent, nil)
	return nil
}

// getRatingsByEvent retrieves all ratings for a given event
// @Summary GET all ratings for an event
// @Description Retrieve all ratings for a given event
// @Tags Rating
// @Accept json
// @Produce json
// @Param eventID path string true "Event ID"
// @Success 200 {array} models.Rating "Successfully retrieved ratings"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings/event/{eventID} [get]
func (rc *RatingController) getRatingsByEvent(ctx *gin.Context) error {
	eventID := ctx.Param("eventID")
	ratings, err := rc.ratingService.GetRatingsByEvent(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	ctx.JSON(http.StatusOK, ratings)
	return nil
}
