package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func (lc *LikesController) RegisterRoutes(router *gin.RouterGroup) {
	likeRouter := router.Group("/likes")
	{
		//likeRouter.GET("/", lc.getAllLikes)
		//likeRouter.GET("/:id", lc.getLike)
		//likeRouter.Post("/", lc.createLike)
		//likeRouter.Put("/:id", lc.updateLike)
		//likeRouter.Delete("/:id", lc.deleteLike)
		likeRouter.POST("/:postID", utility.Use(lc.likePost))
		likeRouter.DELETE("/:postID", utility.Use(lc.unlikePost))
		likeRouter.GET("/:postID/count", utility.Use(lc.getLikeCount))

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
func (c *LikesController) likePost(ctx *gin.Context) error {
	postID := ctx.Param("postID")

	err := c.likesService.LikePost(postID, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully liked post",
	})
	return nil
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
func (c *LikesController) unlikePost(ctx *gin.Context) error {
	postID := ctx.Param("postID")

	err := c.likesService.UnlikePost(postID, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully unliked post",
	})
	return nil
}

// @Summary GET like count for a post
// @Description Retrieve the number of likes for a post
// @Tags Likes
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {object} fiber.Map "Successfully retrieved like count"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/{postID}/count [get]
func (c *LikesController) getLikeCount(ctx *gin.Context) error {
	postID := ctx.Param("postID")

	count, err := c.likesService.GetLikeCount(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": count,
	})
	return nil
}
