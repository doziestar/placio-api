package controller

import (
	"github.com/gofiber/fiber/v2"
	_ "placio-app/Dto"
	"placio-app/service"
	"placio-app/utility"
)

type LikesController struct {
	likesService service.LikeService
}

func NewLikeController(likeService service.LikeService) *LikesController {
	return &LikesController{likesService: likeService}
}

func (lc *LikesController) RegisterRoutes(router fiber.Router) {
	likeRouter := router.Group("/likes")
	{
		//likeRouter.Get("/", lc.getAllLikes)
		//likeRouter.Get("/:id", lc.getLike)
		//likeRouter.Post("/", lc.createLike)
		//likeRouter.Put("/:id", lc.updateLike)
		//likeRouter.Delete("/:id", lc.deleteLike)
		likeRouter.Post("/:postID", utility.Use(lc.likePost))
		likeRouter.Delete("/:postID", utility.Use(lc.unlikePost))
		likeRouter.Get("/:postID/count", utility.Use(lc.getLikeCount))

	}
}

// @Summary Like a post
// @Description Add a like to a post
// @Tags Likes
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Success 201 {object} fiber.Map "Successfully liked post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/{postID} [post]
func (c *LikesController) likePost(ctx *fiber.Ctx) error {
	postID := ctx.Params("postID")

	err := c.likesService.LikePost(postID, "")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Successfully liked post",
	})
}

// @Summary Unlike a post
// @Description Remove a like from a post
// @Tags Likes
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {object} fiber.Map "Successfully unliked post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/{postID} [delete]
func (c *LikesController) unlikePost(ctx *fiber.Ctx) error {
	postID := ctx.Params("postID")

	err := c.likesService.UnlikePost(postID, "")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully unliked post",
	})
}

// @Summary Get like count for a post
// @Description Retrieve the number of likes for a post
// @Tags Likes
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {object} fiber.Map "Successfully retrieved like count"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/{postID}/count [get]
func (c *LikesController) getLikeCount(ctx *fiber.Ctx) error {
	postID := ctx.Params("postID")

	count, err := c.likesService.GetLikeCount(postID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": count,
	})
}
