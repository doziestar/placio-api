package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"
)

type LikeController struct {
	likeService     service.LikeService
	userPlacesLikes service.UserLikePlaceService
}

func NewLikeController(likeService service.LikeService, userPlacesLikes service.UserLikePlaceService) *LikeController {
	return &LikeController{likeService: likeService, userPlacesLikes: userPlacesLikes}
}

func (likesController *LikeController) RegisterRoutes(router *gin.RouterGroup) {
	likeRouter := router.Group("/likes")
	{
		likeRouter.POST("/user/:userID/post/:postID", likesController.likePost)
		likeRouter.DELETE("/:likeID", likesController.unlikePost)
		likeRouter.GET("/user/:userID", likesController.getUserLikes)
		likeRouter.GET("/post/:postID", likesController.getPostLikes)

		likePlaceRouter := likeRouter.Group("/place")
		{
			likePlaceRouter.POST("/:placeID", utility.Use(likesController.likePlace))
			likePlaceRouter.DELETE("/:userLikePlaceID", utility.Use(likesController.unlikePlace))
			likePlaceRouter.GET("/user/:userID", utility.Use(likesController.getUserLikedPlaces))
			likePlaceRouter.GET("/:placeID", utility.Use(likesController.getPlaceLikes))
			likePlaceRouter.GET("check/:placeID", utility.Use(likesController.checkUserLikesPlace))
		}
	}
}

// @Summary Check if user likes a place
// @ID user-check-like-place
// @Tags Like
// @Produce json
// @Param placeID path string true "Place ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Check if the user likes a Place
// @Success 200 {object} string "Like status returned successfully"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/place/check/{placeID} [get]
func (likesController *LikeController) checkUserLikesPlace(ctx *gin.Context) error {
	placeID := ctx.Param("placeID")
	userID := ctx.GetString("user")

	doesUserLike, err := likesController.userPlacesLikes.CheckIfUserLikesPlace(ctx, userID, placeID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": doesUserLike})
	return nil
}

// @Summary Like a post
// @Description Add a like to a post by a user
// @Tags Like
// @Accept json
// @Produce json
// @Param userID path string true "ID of the user"
// @Param postID path string true "ID of the post"
// @Success 200 {object} Dto.Response "Successfully liked post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/user/{userID}/post/{postID} [post]
func (likesController *LikeController) likePost(ctx *gin.Context) {
	userID := ctx.Param("userID")
	postID := ctx.Param("postID")

	like, err := likesController.likeService.LikePost(ctx, userID, postID)
	if err != nil {

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
// @Success 200 {object} Dto.Response "Successfully unliked post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/{likeID} [delete]
func (likesController *LikeController) unlikePost(ctx *gin.Context) {
	likeID := ctx.Param("likeID")

	err := likesController.likeService.UnlikePost(ctx, likeID)
	if err != nil {

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
// @Success 200 {array} ent.Like "List of likes"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/user/{userID} [get]
func (likesController *LikeController) getUserLikes(ctx *gin.Context) {
	userID := ctx.Param("userID")

	likes, err := likesController.likeService.GetUserLikes(ctx, userID)
	if err != nil {

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
// @Success 200 {array} ent.Like "List of likes"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/post/{postID} [get]
func (likesController *LikeController) getPostLikes(ctx *gin.Context) {
	postID := ctx.Param("postID")

	likes, err := likesController.likeService.GetPostLikes(ctx, postID)
	if err != nil {

		return
	}

	ctx.JSON(http.StatusOK, likes)
}

// @Summary Like a place
// @Description Allows a user to like a specific place
// @Tags Like
// @Accept json
// @Produce json
// @Param userID path string true "ID of the user"
// @Param placeID path string true "ID of the place"
// @Param Authorization header string true "Bearer token"
// @Success 201 {object} ent.UserLikePlace "Successfully liked place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/place/{placeID} [post]
func (likesController *LikeController) likePlace(c *gin.Context) error {
	userID := c.MustGet("user").(string)
	placeID := c.Param("placeID")

	like, err := likesController.userPlacesLikes.LikePlace(c, userID, placeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error(), ""))
		return nil
	}

	c.JSON(http.StatusCreated, like)
	return nil
}

// @Summary Unlike a place
// @Description Allows a user to unlike a specific place
// @Tags Like
// @Accept json
// @Produce json
// @Param userLikePlaceID path string true "ID of the UserLikePlace record"
// @Param Authorization header string true "Provide JWT access token"
// @Success 200 {object} string "Successfully unliked place"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/place/{userLikePlaceID} [delete]
func (likesController *LikeController) unlikePlace(c *gin.Context) error {
	userLikePlaceID := c.Param("userLikePlaceID")
	userId := c.MustGet("user").(string)

	err := likesController.userPlacesLikes.UnlikePlace(c, userId, userLikePlaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error(), ""))
		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "Unliked successfully", ""))
	return nil
}

// @Summary Get user liked places
// @Description Retrieve all places liked by a user
// @Tags Like
// @Accept json
// @Produce json
// @Param userID path string true "ID of the user"
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} ent.UserLikePlace "List of UserLikePlace records"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/user/place/{userID} [get]
func (likesController *LikeController) getUserLikedPlaces(c *gin.Context) error {
	userID := c.Param("userID")

	likes, err := likesController.userPlacesLikes.GetUserLikedPlaces(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error(), ""))
		return nil
	}

	c.JSON(http.StatusOK, likes)
	return nil
}

// @Summary Get place likes
// @Description Retrieve all likes for a place
// @Tags Like
// @Accept json
// @Produce json
// @Param placeID path string true "ID of the place"
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} ent.UserLikePlace "List of UserLikePlace records"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/likes/place/{placeID} [get]
func (likesController *LikeController) getPlaceLikes(c *gin.Context) error {
	placeID := c.Param("placeID")

	likes, err := likesController.userPlacesLikes.GetPlaceLikes(c, placeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error(), ""))
		return nil
	}

	c.JSON(http.StatusOK, likes)
	return nil
}
