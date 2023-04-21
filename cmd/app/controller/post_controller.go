package controller

import (
	"github.com/gofiber/fiber/v2"
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

func (pc *PostController) RegisterRoutes(router fiber.Router) {
	postRouter := router.Group("/posts")
	{
		//postRouter.Get("/", middleware.Verify("user"), pc.getAllPosts)
		postRouter.Get("/:id", middleware.Verify("user"), utility.Use(pc.getPost))
		postRouter.Post("/", middleware.Verify("user"), utility.Use(pc.createPost))
		postRouter.Put("/:id", middleware.Verify("user"), utility.Use(pc.updatePost))
		postRouter.Delete("/:id", middleware.Verify("user"), utility.Use(pc.deletePost))

	}
}

// @Summary Create a new post
// @Description Create a new post for the authenticated user
// @Tags Post
// @Accept json
// @Produce json
// @Param CreatePostDto body models.Post true "Post Data"
// @Success 201 {object} models.Post "Successfully created post"
// @Failure 400 {object} dto.ErrorDTO "Bad Request"
// @Failure 500 {object} dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/ [post]
func (c *PostController) createPost(ctx *fiber.Ctx) error {
	data := new(models.Post)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	post := &models.Post{
		Content: data.Content,
		//UserID:  data.UserID,
	}

	newPost, err := c.postService.CreatePost(post)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(newPost)
}

// @Summary Get a post
// @Description Retrieve a post by its ID
// @Tags Post
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {object} models.Post "Successfully retrieved post"
// @Failure 404 {object} dto.ErrorDTO "Post Not Found"
// @Failure 500 {object} dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{postID} [get]
func (c *PostController) getPost(ctx *fiber.Ctx) error {
	postID := ctx.Params("postID")

	post, err := c.postService.GetPost(postID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	if post == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post Not Found",
		})
	}

	return ctx.JSON(post)
}

// @Summary Update a post
// @Description Update a post by its ID
// @Tags Post
// @Accept json
// @Produce json
// @Param postID path string true "Post ID"
// @Param UpdatePostDto body dto.UpdatePostDto true "Post Data"
// @Success 200 {object} models.Post "Successfully updated post"
// @Failure 400 {object} dto.ErrorDTO "Bad Request"
// @Failure 404 {object} dto.ErrorDTO "Post Not Found"
// @Failure 500 {object} dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{postID} [put]
func (c *PostController) updatePost(ctx *fiber.Ctx) error {
	postID := ctx.Params("postID")
	data := new(models.Post)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	post, err := c.postService.GetPost(postID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	if post == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post Not Found",
		})
	}

	//post.Content = data.Content
	updatedPost, err := c.postService.UpdatePost(post)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.JSON(updatedPost)
}

// @Summary Delete a post
// @Description Delete a post by its ID
// @Tags Post
// @Produce json
// @Param postID path string true "Post ID"
// @Success 204 "Successfully deleted post"
// @Failure 404 {object} dto.ErrorDTO "Post Not Found"
// @Failure 500 {object} dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/{postID} [delete]
func (c *PostController) deletePost(ctx *fiber.Ctx) error {
	postID := ctx.Params("postID")
	err := c.postService.DeletePost(postID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// @Summary List all posts
// @Description Retrieve all posts
// @Tags Post
// @Produce json
// @Success 200 {array} models.Post "Successfully retrieved posts"
// @Failure 500 {object} dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/posts/ [get]
func (c *PostController) listPosts(ctx *fiber.Ctx) error {
	posts, err := c.postService.ListPosts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return ctx.JSON(posts)

}
