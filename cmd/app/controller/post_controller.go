package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/middleware"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
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
		//postRouter.Get("/", middleware.Verify("user"), pc.getAllPosts)
		postRouter.GET("/:id", middleware.Verify("user"), utility.Use(pc.getPost))
		postRouter.POST("/", middleware.Verify("user"), utility.Use(pc.createPost))
		postRouter.PUT("/:id", middleware.Verify("user"), utility.Use(pc.updatePost))
		postRouter.DELETE("/:id", middleware.Verify("user"), utility.Use(pc.deletePost))

	}
}

// @Summary Create a new post
// @Description Create a new post for the authenticated user
// @Tags Post
// @Accept json
// @Produce json
// @Param CreatePostDto body models.Post true "Post Data"
// @Success 201 {object} models.Post "Successfully created post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/ [post]
func (c *PostController) createPost(ctx *gin.Context) error {
	data := new(models.Post)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
	}

	post := &models.Post{
		Content: data.Content,
		//UserID:  data.UserID,
	}

	newPost, err := c.postService.CreatePost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusCreated, newPost)
	return nil
}

// @Summary Get a post
// @Description Retrieve a post by its ID
// @Tags Post
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {object} models.Post "Successfully retrieved post"
// @Failure 404 {object} Dto.ErrorDTO "Post Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{postID} [get]
func (c *PostController) getPost(ctx *gin.Context) error {
	postID := ctx.Param("postID")

	post, err := c.postService.GetPost(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Post Not Found",
		})
		return nil
	}

	ctx.JSON(http.StatusOK, post)
	return nil
}

// @Summary Update a post
// @Description Update a post by its ID
// @Tags Post
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Param UpdatePostDto body models.Post true "Post Data"
// @Success 200 {object} models.Post "Successfully updated post"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Post Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{postID} [put]
func (c *PostController) updatePost(ctx *gin.Context) error {
	postID := ctx.Param("postID")
	data := new(models.Post)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
	}
	post, err := c.postService.GetPost(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	if post == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Post Not Found",
		})
		return nil
	}

	//post.Content = data.Content
	updatedPost, err := c.postService.UpdatePost(post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, updatedPost)
	return nil
}

// @Summary Delete a post
// @Description Delete a post by its ID
// @Tags Post
// @Produce json
// @Param postID path string true "Post ID"
// @Success 204 "Successfully deleted post"
// @Failure 404 {object} Dto.ErrorDTO "Post Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{postID} [delete]
func (c *PostController) deletePost(ctx *gin.Context) error {
	postID := ctx.Param("postID")
	err := c.postService.DeletePost(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

// @Summary List all posts
// @Description Retrieve all posts
// @Tags Post
// @Produce json
// @Success 200 {array} models.Post "Successfully retrieved posts"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/ [get]
func (c *PostController) listPosts(ctx *gin.Context) error {
	posts, err := c.postService.ListPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}
	ctx.JSON(http.StatusOK, posts)
	return nil

}
