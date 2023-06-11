package controller

import (
	"errors"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
	"time"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService            service.PostService
	userService            service.UserService
	businessAccountService service.BusinessAccountService
	mediaService           service.MediaService
}

func NewPostController(postService service.PostService, userService service.UserService, businessAccountService service.BusinessAccountService, mediaService service.MediaService) *PostController {
	return &PostController{postService: postService, userService: userService, businessAccountService: businessAccountService, mediaService: mediaService}
}

func (pc *PostController) RegisterRoutes(router *gin.RouterGroup) {
	postRouter := router.Group("/posts")
	{
		postRouter.GET("/:id", utility.Use(pc.getPost))
		postRouter.POST("/", utility.Use(pc.createPost))
		postRouter.PUT("/:id", utility.Use(pc.updatePost))
		postRouter.DELETE("/:id", utility.Use(pc.deletePost))
		postRouter.GET("/:id/comments", utility.Use(pc.getComments))
		postRouter.POST("/:id/comments", utility.Use(pc.createComment))
		postRouter.PUT("/comments/:id", utility.Use(pc.updateComment))
		postRouter.DELETE("/comments/:id", utility.Use(pc.deleteComment))
		postRouter.POST("/:id/like", utility.Use(pc.likePost))
		postRouter.POST("/:id/unlike", utility.Use(pc.unlikePost))
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
// createPost creates a new post.
func (pc *PostController) createPost(ctx *gin.Context) error {
	// Extract the user from the context
	authOID := ctx.MustGet("user").(string)
	user, err := pc.userService.GetUser(authOID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	// Extract the BusinessAccountID from the query parameters, if it exists
	businessAccountId := ctx.Query("businessAccountId")

	// Bind the incoming JSON to a new Post instance
	data := new(models.Post)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return err
	}

	// Validate the content of the post
	if data.Content == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Content cannot be empty"})
		return errors.New("content cannot be empty")
	}

	// If a BusinessAccountID was provided, verify that it exists
	if businessAccountId != "" {
		if _, err := pc.businessAccountService.GetBusinessAccount(businessAccountId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Business Account ID"})
			return err
		}
	}

	// Create a new Post instance
	post := &models.Post{
		Content:   data.Content,
		UserID:    user.UserID,
		ID:        models.GenerateID(),
		CreatedAt: time.Now(),
	}

	// If a BusinessAccountID was provided, associate the post with the business account
	if businessAccountId != "" {
		post.BusinessAccountID = businessAccountId
	}

	// Create the post
	newPost, err := pc.postService.CreatePost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	// Handle media files, if any were provided
	// (This is a placeholder - you'll need to replace it with your actual media handling code)
	for _, media := range data.Medias {
		// Upload the media file to Cloudinary (or another service)
		uploadedMediaURL := "..." // replace with actual URL from Cloudinary
		media.URL = uploadedMediaURL
		// Associate the media with the post
		media.PostID = newPost.ID
		// Save the media in the database
		if _, err := pc.mediaService.CreateMedia(&media); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return err
		}
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
	//userID, ok := utility.GetUserIDFromContext(ctx)
	//if !ok {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{
	//		"error": "Unauthorized",
	//	})
	//	return nil
	//}
	auth0ID := ctx.MustGet("user").(string)
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

	user, err := pc.userService.GetUser(auth0ID)

	if post.UserID != user.UserID {
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
	//userID, ok := utility.GetUserIDFromContext(ctx)
	//if !ok {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{
	//		"error": "Unauthorized",
	//	})
	//	return nil
	//}
	userID := ctx.MustGet("userId").(string)
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

	comments, err := pc.postService.GetComments(postID, 0, 0, "created_at", map[string]interface{}{})
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
	//userID, ok := utility.GetUserIDFromContext(ctx)
	//if !ok {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{
	//		"error": "Unauthorized",
	//	})
	//	return nil
	//}
	userID := ctx.MustGet("userId").(string)
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
	//userID, ok := utility.GetUserIDFromContext(ctx)
	//if !ok {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{
	//		"error": "Unauthorized",
	//	})
	//	return nil
	//}
	userID := ctx.MustGet("userId").(string)
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
	//userID, ok := utility.GetUserIDFromContext(ctx)
	//if !ok {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{
	//		"error": "Unauthorized",
	//	})
	//	return nil
	//}
	userID := ctx.MustGet("userId").(string)
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
	//userID, ok := utility.GetUserIDFromContext(ctx)
	//if !ok {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{
	//		"error": "Unauthorized",
	//	})
	//	return nil
	//}
	userID := ctx.MustGet("userId").(string)
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
	//userID, ok := utility.GetUserIDFromContext(ctx)
	//if !ok {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{
	//		"error": "Unauthorized",
	//	})
	//	return nil
	//}
	userID := ctx.MustGet("userId").(string)
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
