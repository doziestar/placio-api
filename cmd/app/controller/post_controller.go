package controller

import (
	"net/http"
	_ "placio-app/Dto"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{postService: postService}
}

func (pc *PostController) RegisterRoutes(router *gin.RouterGroup) {
	postRouter := router.Group("/posts")
	{
		postRouter.GET("/:id", utility.EnsureValidToken(), utility.Use(pc.getPost))
		postRouter.POST("/", utility.EnsureValidToken(), utility.Use(pc.createPost))
		postRouter.PUT("/:id", utility.EnsureValidToken(), utility.Use(pc.updatePost))
		postRouter.DELETE("/:id", utility.EnsureValidToken(), utility.Use(pc.deletePost))
		postRouter.GET("/:id/comments", utility.EnsureValidToken(), utility.Use(pc.getComments))
		postRouter.POST("/:id/comments", utility.EnsureValidToken(), utility.Use(pc.createComment))
		postRouter.PUT("/comments/:id", utility.EnsureValidToken(), utility.Use(pc.updateComment))
		postRouter.DELETE("/comments/:id", utility.EnsureValidToken(), utility.Use(pc.deleteComment))
		postRouter.POST("/:id/like", utility.EnsureValidToken(), utility.Use(pc.likePost))
		postRouter.POST("/:id/unlike", utility.EnsureValidToken(), utility.Use(pc.unlikePost))
	}
}

// CreatePost creates a new post.
// @Summary Create a new post
// @Description Create a new post for the authenticated user
// @Tags Post
// @Accept json
// @Produce json
// @Param CreatePostDto body models.Post true "Post Data"
// @Security Bearer
// @Success 201 {object} models.Post "Successfully created post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/ [post]
func (pc *PostController) createPost(ctx *gin.Context) error {
	userID := utility.GetUserIDFromContext(ctx)

	data := new(models.Post)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
	}

	post := &models.Post{
		Content: data.Content,
		UserID:  userID,
	}

	newPost, err := pc.postService.CreatePost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusCreated, newPost)
	return nil
}

// GetPost retrieves a post by ID.
// @Summary Get a post
// @Description Get a post by ID
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} models.Post "Successfully retrieved post"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id} [get]
func (pc *PostController) getPost(ctx *gin.Context) error {
	postID := ctx.Param("id")

	post, err := pc.postService.GetPost(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return nil
	}

	ctx.JSON(http.StatusOK, post)
	return nil
}

// UpdatePost updates a post.
// @Summary Update a post
// @Description Update a post by ID
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param UpdatePostDto body models.Post true "Post Data"
// @Security Bearer
// @Success 200 {object} models.Post "Successfully updated post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id} [put]
func (pc *PostController) updatePost(ctx *gin.Context) error {
	userID := utility.GetUserIDFromContext(ctx)
	postID := ctx.Param("id")

	data := new(models.Post)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
	}

	post, err := pc.postService.GetPost(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return nil
	}

	if post.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return nil
	}

	post.Content = data.Content

	updatedPost, err := pc.postService.UpdatePost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, updatedPost)
	return nil
}

// DeletePost deletes a post.
// @Summary Delete a post
// @Description Delete a post by ID
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security Bearer
// @Success 204 "Successfully deleted post"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id} [delete]
func (pc *PostController) deletePost(ctx *gin.Context) error {
	userID := utility.GetUserIDFromContext(ctx)
	postID := ctx.Param("id")

	post, err := pc.postService.GetPost(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return nil
	}

	if post.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return nil
	}

	err = pc.postService.DeletePost(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

// GetComments retrieves comments for a post.
// @Summary Get comments for a post
// @Description Get comments for a post by ID
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {array} models.Comment "Successfully retrieved comments"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id}/comments [get]
func (pc *PostController) getComments(ctx *gin.Context) error {
	postID := ctx.Param("id")

	comments, err := pc.postService.GetComments(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, comments)
	return nil
}

// CreateComment creates a new comment for a post.
// @Summary Create a new comment
// @Description Create a new comment for a post by ID
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param CreateCommentDto body models.Comment true "Comment Data"
// @Security Bearer
// @Success 201 {object} models.Comment "Successfully created comment"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id}/comments [post]
func (pc *PostController) createComment(ctx *gin.Context) error {
	userID := utility.GetUserIDFromContext(ctx)
	postID := ctx.Param("id")

	data := new(models.Comment)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
	}

	comment := &models.Comment{
		PostID:  postID,
		UserID:  userID,
		Content: data.Content,
	}

	newComment, err := pc.postService.CreateComment(postID, comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusCreated, newComment)
	return nil
}

// UpdateComment updates a comment.
// @Summary Update a comment
// @Description Update a comment by ID
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Param UpdateCommentDto body models.Comment true "Comment Data"
// @Security Bearer
// @Success 200 {object} models.Comment "Successfully updated comment"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/comments/{id} [put]
func (pc *PostController) updateComment(ctx *gin.Context) error {
	userID := utility.GetUserIDFromContext(ctx)
	commentID := ctx.Param("id")

	data := new(models.Comment)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
	}

	comment, err := pc.postService.GetComment(commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	if comment == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Comment not found",
		})
		return nil
	}

	if comment.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return nil
	}

	comment.Content = data.Content

	updatedComment, err := pc.postService.UpdateComment(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, updatedComment)
	return nil
}

// DeleteComment deletes a comment.
// @Summary Delete a comment
// @Description Delete a comment by ID
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Security Bearer
// @Success 204 "Successfully deleted comment"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/comments/{id} [delete]
func (pc *PostController) deleteComment(ctx *gin.Context) error {
	userID := utility.GetUserIDFromContext(ctx)
	commentID := ctx.Param("id")

	comment, err := pc.postService.GetComment(commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	if comment == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Comment not found",
		})
		return nil
	}

	if comment.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return nil
	}

	err = pc.postService.DeleteComment(commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

// LikePost likes a post.
// @Summary Like a post
// @Description Like a post by ID
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security Bearer
// @Success 200 "Successfully liked post"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id}/like [post]
func (pc *PostController) likePost(ctx *gin.Context) error {
	userID := utility.GetUserIDFromContext(ctx)
	postID := ctx.Param("id")

	err := pc.postService.LikePost(postID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.Status(http.StatusOK)
	return nil
}

// UnlikePost unlikes a post.
// @Summary Unlike a post
// @Description Unlike a post by ID
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security Bearer
// @Success 200 "Successfully unliked post"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id}/unlike [post]
func (pc *PostController) unlikePost(ctx *gin.Context) error {
	userID := utility.GetUserIDFromContext(ctx)
	postID := ctx.Param("id")

	err := pc.postService.UnlikePost(postID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.Status(http.StatusOK)
	return nil
}
