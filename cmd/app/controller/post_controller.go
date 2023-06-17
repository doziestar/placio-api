package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/Dto"
	_ "placio-app/Dto"
	"placio-app/ent"
	_ "placio-app/ent"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
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
		postRouter.GET("/:id/comments", utility.Use(pc.getCommentsByPost))
		//postRouter.GET("/:id/user", utility.Use(pc.getPostsByUser))

		//postRouter.PUT("/:id", utility.Use(pc.updatePost))
		//postRouter.DELETE("/:id", utility.Use(pc.deletePost))
		//postRouter.GET("/:id/comments", utility.Use(pc.getComments))
		//postRouter.POST("/:id/comments", utility.Use(pc.createComment))
		//postRouter.PUT("/comments/:id", utility.Use(pc.updateComment))
		//postRouter.DELETE("/comments/:id", utility.Use(pc.deleteComment))
		//postRouter.POST("/:id/like", utility.Use(pc.likePost))
		//postRouter.POST("/:id/unlike", utility.Use(pc.unlikePost))
	}
}

// CreatePost creates a new post.
// @Summary Create a new post
// @Description Create a new post for the authenticated user
// @Tags Post
// @Accept json
// @Produce json
// @Param CreatePostDto body Dto.PostDto true "Post Data"
// @Security Bearer
// @Success 201 {object} Dto.PostResponseDto "Successfully created post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/ [post]
// createPost creates a new post.
func (pc *PostController) createPost(ctx *gin.Context) error {
	// Extract the user from the context
	authOID := ctx.MustGet("auth0_id").(string)
	user, err := pc.userService.GetUser(ctx, authOID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	// Extract the BusinessAccountID from the query parameters, if it exists
	var businessID *string
	var businessAccount *ent.Business
	businessAccountId := ctx.Query("businessAccountId")
	if businessAccountId != "" {
		// Verify that the business account exists
		businessAccount, err = pc.businessAccountService.GetBusinessAccount(ctx, businessAccountId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Business Account ID"})
			return err
		}
		businessID = &businessAccount.ID
	}

	// Bind the incoming JSON to a new PostDto instance
	data := new(Dto.PostDto)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return err
	}

	// Create a new Post instance
	post := &ent.Post{
		Content: data.Content,
		ID:      models.GenerateID(),
	}

	// Create the post
	newPost, err := pc.postService.CreatePost(ctx, post, user.ID, businessID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	// Handle media files, if any were provided
	medias := make([]*ent.Media, len(data.Medias))
	for i, mediaDto := range data.Medias {
		media := &ent.Media{
			ID:        utility.GenerateID(),
			URL:       mediaDto.URL,
			MediaType: mediaDto.Type,
		}
		// Save the media in the database
		createdMedia, err := pc.mediaService.CreateMedia(ctx, media)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return err
		}
		medias[i] = createdMedia
	}

	// Create a response struct
	response := Dto.PostResponseDto{
		ID:        newPost.ID,
		Content:   newPost.Content,
		User:      user,
		Business:  businessAccount,
		CreatedAt: newPost.CreatedAt,
		Medias:    make([]Dto.MediaDto, len(medias)),
	}

	for i, media := range medias {
		response.Medias[i] = Dto.MediaDto{
			Type: media.MediaType,
			URL:  media.URL,
		}
	}

	ctx.JSON(http.StatusCreated, response)
	return nil
}

// GetPost retrieves a post by ID.
// @Summary Get a post
// @Description Get a post by ID
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} ent.Post "Successfully retrieved post"
// @Failure 404 {object} Dto.ErrorDTO "Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id} [get]
func (pc *PostController) getPost(ctx *gin.Context) error {
	postID := ctx.Param("id")

	post, err := pc.postService.GetPost(ctx, postID)
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

// UpdatePost updates an existing post.
// @Summary Update a post
// @Description Update an existing post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param UpdatePostDto body Dto.PostDto true "Post Data"
// @Security Bearer
// @Success 200 {object} Dto.PostResponseDto "Successfully updated post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Post Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id} [put]
func (pc *PostController) updatePost(ctx *gin.Context) error {
	// Extract the user from the context
	authOID := ctx.MustGet("auth0_id").(string)
	user, err := pc.userService.GetUser(ctx, authOID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	postId := ctx.Param("id")

	// Bind the incoming JSON to a new PostDto instance
	data := new(Dto.PostDto)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return err
	}

	// Create a new Post instance
	post := &ent.Post{
		Content: data.Content,
		ID:      postId,
	}

	// Update the post
	updatedPost, err := pc.postService.UpdatePost(ctx, post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	// Create a response struct
	response := Dto.PostResponseDto{
		ID:        updatedPost.ID,
		Content:   updatedPost.Content,
		User:      user,
		CreatedAt: updatedPost.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
	return nil
}

// DeletePost deletes an existing post.
// @Summary Delete a post
// @Description Delete an existing post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security Bearer
// @Success 200 {object} string "Successfully deleted post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Post Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id} [delete]
func (pc *PostController) deletePost(ctx *gin.Context) error {
	postID := ctx.Param("id")

	err := pc.postService.DeletePost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "post deleted successfully"})
	return nil
}

// GetCommentsByPost retrieves all comments for a given post.
// @Summary Get comments by post
// @Description Retrieve all comments for a given post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security Bearer
// @Success 200 {object} []ent.Comment "Successfully retrieved comments"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Post Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{id}/comments [get]
func (pc *PostController) getCommentsByPost(ctx *gin.Context) error {
	postID := ctx.Param("id")

	comments, err := pc.postService.GetCommentsByPost(ctx, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	ctx.JSON(http.StatusOK, comments)
	return nil
}

// GetPostsByUser retrieves all posts for a given user.
// @Summary Get posts by user
// @Description Retrieve all posts for a given user
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "User ID"
