package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/Dto"
	"placio-app/ent"
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
		ratingRouter.GET("/", utility.Use(rc.listRatings))
	}
}

// createRating creates a new rating
// @Summary Create a new rating
// @Description Create a new rating for the specified event
// @Tags Rating
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param Dto.RatingDTO body Dto.RatingDTO true "Rating Data"
// @Success 201 {object} ent.Rating "Successfully created rating"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings/ [post]
func (rc *RatingController) createRating(ctx *gin.Context) error {
	data := new(Dto.RatingDTO)
	if err := ctx.BindJSON(data); err != nil {

		return err
	}

	newRating, err := rc.ratingService.CreateRating(ctx, data)
	if err != nil {

		return err
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
// @Param Authorization header string true "JWT Token"
// @Param id path string true "Rating ID"
// @Success 200 {object} ent.Rating "Successfully retrieved rating"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Rating Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings/{id} [get]
func (rc *RatingController) getRating(ctx *gin.Context) error {
	ratingID := ctx.Param("id")
	rating, err := rc.ratingService.GetRating(ctx, ratingID)
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Rating Not Found"})
			return err
		}

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
// @Param Authorization header string true "JWT Token"
// @Param id path string true "Rating ID"
// @Param score body int true "New Score"
// @Success 200 {object} ent.Rating "Successfully updated rating"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Rating Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings/{id} [put]
func (rc *RatingController) updateRating(ctx *gin.Context) error {
	ratingID := ctx.Param("id")
	var score int
	if err := ctx.BindJSON(&score); err != nil {

		return err
	}

	updatedRating, err := rc.ratingService.UpdateRating(ctx, ratingID, score)
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Rating Not Found"})
			return err
		}

		return err
	}

	ctx.JSON(http.StatusOK, updatedRating)
	return nil
}

// deleteRating deletes a rating by its ID
// @Summary Delete rating by ID
// @Description Delete a rating by its ID
// @Tags Rating
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param id path string true "Rating ID"
// @Success 204 "Successfully deleted rating"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Rating Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings/{id} [delete]
func (rc *RatingController) deleteRating(ctx *gin.Context) error {
	ratingID := ctx.Param("id")
	err := rc.ratingService.DeleteRating(ctx, ratingID)
	if err != nil {
		if ent.IsNotFound(err) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Rating Not Found"})
			return err
		}

		return err
	}
	ctx.JSON(http.StatusNoContent, nil)
	return nil
}

// ListRatings retrieves all ratings
// @Summary GET all ratings
// @Description Retrieve all ratings
// @Tags Rating
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Success 200 {array} ent.Rating "Successfully retrieved ratings"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings [get]
func (rc *RatingController) listRatings(ctx *gin.Context) error {
	ratings, err := rc.ratingService.ListRatings(ctx)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, ratings)
	return nil
}
