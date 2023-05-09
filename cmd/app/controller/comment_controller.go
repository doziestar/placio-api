package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
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
		commentRouter.GET("/:id", utility.Use(cc.getComment))
		commentRouter.POST("/", utility.Use(cc.createComment))
		commentRouter.PUT("/:id", utility.Use(cc.updateComment))
		commentRouter.DELETE("/:id", utility.Use(cc.deleteComment))
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
func (c *CommentController) createComment(ctx *gin.Context) error {
	data := new(models.Comment)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})

		return err
	}

	comment := &models.Comment{
		Content: data.Content,
		//UserID:  data.UserID,
		PostID: data.PostID,
	}

	newComment, err := c.commentService.CreateComment(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusCreated, newComment)
	return nil
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
func (c *CommentController) getComment(ctx *gin.Context) error {
	commentID := ctx.Param("commentID")

	comment, err := c.commentService.GetComment(commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	if comment == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Comment Not Found",
		})
		return nil
	}

	ctx.JSON(http.StatusOK, comment)
	return nil
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
func (c *CommentController) updateComment(ctx *gin.Context) error {
	commentID := ctx.Param("commentID")
	data := new(models.Comment)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return err
	}

	comment, err := c.commentService.GetComment(commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	if comment == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Comment Not Found",
		})
		return nil
	}

	//comment.Content = data.Content
	updatedComment, err := c.commentService.UpdateComment(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, updatedComment)
	return nil
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
func (c *CommentController) deleteComment(ctx *gin.Context) error {
	commentID := ctx.Param("commentID")
	err := c.commentService.DeleteComment(commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.Status(http.StatusNoContent)
	return nil
}

// @Summary List all comments for a post
// @Description Retrieve all comments for the specified post
// @Tags Comment
// @Produce json
// @Param postID path string true "Post ID"
// @Success 200 {array} models.Comment "Successfully retrieved comments"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/comments/post/{postID} [get]
func (c *CommentController) listComments(ctx *gin.Context) error {
	postID := ctx.Param("postID")
	comments, err := c.commentService.ListComments(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, comments)
	return nil
}
