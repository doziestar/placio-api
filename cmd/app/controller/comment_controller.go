package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	_ "placio-app/Dto"
	"placio-app/models"
	"placio-app/service"
)

type CommentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

func (cc *CommentController) RegisterRoutes(router *gin.RouterGroup) {
	commentRouter := router.Group("/comments")
	{
		//commentRouter.Get("/", cc.getAllComments)
		commentRouter.Get("/:id", cc.getComment)
		commentRouter.Post("/", cc.createComment)
		commentRouter.Put("/:id", cc.updateComment)
		commentRouter.Delete("/:id", cc.deleteComment)
	}
}

// @Summary Create a new comment
// @Description Create a new comment for the specified post
// @Tags Comment
// @Accept json
// @Produce json
// @Param CreateCommentDto body models.Comment true "Comment Data"
// @Success 201 {object} models.Comment "Successfully created comment"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/comments/ [post]
func (c *CommentController) createComment(ctx *fiber.Ctx) error {
	data := new(models.Comment)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	comment := &models.Comment{
		Content: data.Content,
		//UserID:  data.UserID,
		PostID: data.PostID,
	}

	newComment, err := c.commentService.CreateComment(comment)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(newComment)
}

// @Summary Get a comment
// @Description Retrieve a comment by its ID
// @Tags Comment
// @Produce json
// @Param commentID path string true "Comment ID"
// @Success 200 {object} models.Comment "Successfully retrieved comment"
// @Failure 404 {object} Dto.ErrorDTO "Comment Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/comments/{commentID} [get]
func (c *CommentController) getComment(ctx *fiber.Ctx) error {
	commentID := ctx.Params("commentID")

	comment, err := c.commentService.GetComment(commentID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	if comment == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Comment Not Found",
		})
	}

	return ctx.JSON(comment)
}

// @Summary Update a comment
// @Description Update a comment by its ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param commentID path string true "Comment ID"
// @Param UpdateCommentDto body models.Comment true "Comment Data"
// @Success 200 {object} models.Comment "Successfully updated comment"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 404 {object} Dto.ErrorDTO "Comment Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/comments/{commentID} [put]
func (c *CommentController) updateComment(ctx *fiber.Ctx) error {
	commentID := ctx.Params("commentID")
	data := new(models.Comment)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	comment, err := c.commentService.GetComment(commentID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	if comment == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Comment Not Found",
		})
	}

	//comment.Content = data.Content
	updatedComment, err := c.commentService.UpdateComment(comment)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.JSON(updatedComment)
}

// @Summary Delete a comment
// @Description Delete a comment by its ID
// @Tags Comment
// @Produce json
// @Param commentID path string true "Comment ID"
// @Success 204 "Successfully deleted comment"
// @Failure 404 {object} Dto.ErrorDTO "Comment Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/comments/{commentID} [delete]
func (c *CommentController) deleteComment(ctx *fiber.Ctx) error {
	commentID := ctx.Params("commentID")
	err := c.commentService.DeleteComment(commentID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// @Summary List all comments for a post
// @Description Retrieve all comments for the specified post
// @Tags Comment
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {array} models.Comment "Successfully retrieved comments"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/comments/post/{postID} [get]
func (c *CommentController) listComments(ctx *fiber.Ctx) error {
	postID := ctx.Params("postID")
	comments, err := c.commentService.ListComments(postID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return ctx.JSON(comments)
}
