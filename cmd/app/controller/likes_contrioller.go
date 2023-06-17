package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
)

type LikeController struct {
	likeService service.LikeService
}

func NewLikeController(likeService service.LikeService) *LikeController {
	return &LikeController{likeService: likeService}
}

func (lc *LikeController) RegisterRoutes(router *gin.RouterGroup) {
	likeRouter := router.Group("/likes")
	{
		likeRouter.POST("/user/:userID/post/:postID", lc.likePost)
		likeRouter.DELETE("/:likeID", lc.unlikePost)
		likeRouter.GET("/user/:userID", lc.getUserLikes)
		likeRouter.GET("/post/:postID", lc.getPostLikes)
	}
}

// @Summary Like a post
// @Description Add a like to a post by a user
// @Tags Like
// @Accept json
// @Produce json
// @Param userID path string true "ID of the user"
// @Param postID path string true "ID of the post"
// @Security Bearer
// @Success 200 {object} Dto.Response "Successfully liked post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/user/{userID}/post/{postID} [post]
func (lc *LikeController) likePost(ctx *gin.Context) {
	userID := ctx.Param("userID")
	postID := ctx.Param("postID")

	like, err := lc.likeService.LikePost(ctx, userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, like)
}

// @Summary Unlike a post
// @Description Remove a like from a post
// @Tags Like
// @Accept json
// @Produce json
// @Param likeID path string true "ID of the like"
// @Security Bearer
// @Success 200 {object} Dto.Response "Successfully unliked post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/{likeID} [delete]
func (lc *LikeController) unlikePost(ctx *gin.Context) {
	likeID := ctx.Param("likeID")

	err := lc.likeService.UnlikePost(ctx, likeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully unliked post"})
}

// @Summary Get user likes
// @Description Retrieve all likes by a user
// @Tags Like
// @Accept json
// @Produce json
// @Param userID path string true "ID of the user"
// @Security Bearer
// @Success 200 {array} ent.Like "List of likes"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/user/{userID} [get]
func (lc *LikeController) getUserLikes(ctx *gin.Context) {
	userID := ctx.Param("userID")

	likes, err := lc.likeService.GetUserLikes(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, likes)
}

// @Summary Get post likes
// @Description Retrieve all likes for a post
// @Tags Like
// @Accept json
// @Produce json
// @Param postID path string true "ID of the post"
// @Security Bearer
// @Success 200 {array} ent.Like "List of likes"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/post/{postID} [get]
func (lc *LikeController) getPostLikes(ctx *gin.Context) {
	postID := ctx.Param("postID")

	likes, err := lc.likeService.GetPostLikes(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, likes)
}
