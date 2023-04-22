package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	_ "placio-app/Dto"
	"placio-app/models"
	"placio-app/service"
)

type RatingController struct {
	ratingService service.RatingService
}

func NewRatingController(ratingService service.RatingService) *RatingController {
	return &RatingController{ratingService: ratingService}
}

func (rc *RatingController) RegisterRoutes(router fiber.Router) {
	ratingRouter := router.Group("/ratings")
	{
		ratingRouter.Post("/", rc.createRating)
		ratingRouter.Get("/:id", rc.getRating)
		ratingRouter.Put("/:id", rc.updateRating)
		ratingRouter.Delete("/:id", rc.deleteRating)
		ratingRouter.Get("/event/:eventID", rc.getRatingsByEvent)
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
func (rc *RatingController) createRating(ctx *fiber.Ctx) error {
	data := new(models.Rating)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	rating := &models.Rating{
		EventID: data.EventID,
		UserID:  data.UserID,
		//RateValue: data.RateValue,
	}

	newRating := rc.ratingService.CreateRating(rating)
	if newRating == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(newRating)
}

// getRating retrieves a rating by its ID
// @Summary Get rating by ID
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
func (rc *RatingController) getRating(ctx *fiber.Ctx) error {
	ratingID := ctx.Params("id")
	rating, err := rc.ratingService.GetRatingsByEvent(ratingID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Rating Not Found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(rating)
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
func (rc *RatingController) updateRating(ctx *fiber.Ctx) error {
	ratingID := ctx.Params("id")
	data := new(models.Rating)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	_, err := rc.ratingService.GetRatingsByEvent(ratingID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Rating Not Found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	//rating.RateValue = data.RateValue
	//updatedRating, err := rc.ratingService.UpdateRating(rating)
	//if err != nil {
	//	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Internal Server Error",
	//	})
	//}

	return ctx.Status(fiber.StatusOK).JSON("updatedRating")
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
func (rc *RatingController) deleteRating(ctx *fiber.Ctx) error {
	ratingID := ctx.Params("id")
	err := rc.ratingService.DeleteRating(ratingID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Rating Not Found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return ctx.Status(fiber.StatusNoContent).SendString("")
}

// getRatingsByEvent retrieves all ratings for a given event
// @Summary Get all ratings for an event
// @Description Retrieve all ratings for a given event
// @Tags Rating
// @Accept json
// @Produce json
// @Param eventID path string true "Event ID"
// @Success 200 {array} models.Rating "Successfully retrieved ratings"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/ratings/event/{eventID} [get]
func (rc *RatingController) getRatingsByEvent(ctx *fiber.Ctx) error {
	eventID := ctx.Params("eventID")
	ratings, err := rc.ratingService.GetRatingsByEvent(eventID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(ratings)
}
