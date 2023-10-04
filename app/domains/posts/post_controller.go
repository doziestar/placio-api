package posts

import (
	"net/http"
	_ "placio-app/Dto"
	"placio-app/domains/business"
	"placio-app/domains/media"
	"placio-app/domains/users"
	"placio-app/ent"
	_ "placio-app/ent"
	"placio-app/ent/post"
	"placio-app/models"
	"placio-app/utility"
	appErr "placio-pkg/errors"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService            PostService
	userService            users.UserService
	businessAccountService business.BusinessAccountService
	mediaService           media.MediaService
}

func NewPostController(postService PostService, userService users.UserService, businessAccountService business.BusinessAccountService, mediaService media.MediaService) *PostController {
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
		postRouter.GET("/", utility.Use(pc.getPostFeeds))
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
// @Success 201 {object} Dto.PostResponseDto "Successfully created post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/ [post]
// createPost creates a new post.
func (pc *PostController) createPost(ctx *gin.Context) error {
	// Extract the user from the context
	userId := ctx.MustGet("user").(string)
	businessAccountId := ctx.Query("businessAccountId")

	// Bind the incoming JSON to a new PostDto instance
	data := new(PostDto)
	if err := ctx.BindJSON(data); err != nil {

		return err
	}

	privacy := post.Privacy(data.Privacy)

	// Create a new Post instance
	post := &ent.Post{
		Content: data.Content,
		ID:      models.GenerateID(),
		Privacy: privacy,
	}

	// Create the post
	newPost, err := pc.postService.CreatePost(ctx, post, userId, businessAccountId, data.Medias)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusCreated, utility.ProcessResponse(newPost, "success", "Post created successfully"))
	return nil
}

// GetPostFeeds retrieves all posts for the authenticated user.
// @Summary Get post realtime_server
// @Description Retrieve all posts for the authenticated user
// @Tags Post
// @Accept json
// @Produce json
// @Param businessId query string false "Business ID"
// @Param limit query string false "Limit"
// @Param offset query string false "Offset"
// @Success 200 {object} []Dto.PostResponseDto
// @Failure 400 {object} Dto.ErrorDTO
// @Failure 401 {object} Dto.ErrorDTO
// @Failure 500 {object} Dto.ErrorDTO
// @Router /api/v1/posts/ [get]
func (pc *PostController) getPostFeeds(ctx *gin.Context) error {
	// Extract the user from the context

	// Get the posts
	posts, err := pc.postService.GetPostFeeds(ctx)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(posts, "success"))
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

		return err
	}

	if post == nil {

		return nil
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(post, "success", "Post retrieved successfully"))
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

		return err
	}

	postId := ctx.Param("id")

	// Check if the post exists and belongs to the user
	postData, err := pc.postService.GetPost(ctx, postId)
	if err != nil {
		if ent.IsNotFound(err) {

			return err
		}

		return err
	}

	if postData.Edges.User.ID != user.ID {
		return appErr.ErrUnauthorized

	}

	// Bind the incoming JSON to a new PostDto instance
	data := new(PostDto)
	if err := ctx.BindJSON(data); err != nil || data.Content == "" {
		return err
	}

	privacy := post.Privacy(data.Privacy)

	// Create a new Post instance
	postData = &ent.Post{
		Content: data.Content,
		ID:      postId,
		Privacy: privacy,
	}

	// Update the post
	updatedPost, err := pc.postService.UpdatePost(ctx, postData)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(updatedPost, "success", "Post updated successfully"))
	return nil
}

// DeletePost deletes an existing post.
// @Summary Delete a post
// @Description Delete an existing post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
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
		if ent.IsNotFound(err) {

			return nil
		}

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "Post deleted successfully"))
	return nil
}

// GetCommentsByPost retrieves all comments for a given post.
// @Summary Get comments by post
// @Description Retrieve all comments for a given post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
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
