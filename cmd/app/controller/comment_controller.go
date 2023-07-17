package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"placio-app/Dto"
	_ "placio-app/Dto"
	"placio-app/service"
	"placio-app/utility"
)

type CommentController struct {
	commentService service.CommentService
	userService    service.UserService
}

func NewCommentController(commentService service.CommentService, userService service.UserService) *CommentController {
	return &CommentController{commentService: commentService, userService: userService}
}

func (cc *CommentController) RegisterRoutes(router *gin.RouterGroup) {
	commentRouter := router.Group("/comments")
	{
		//commentRouter.Get("/", cc.getAllComments)
		//commentRouter.GET("/:id", utility.Use(cc.getComment))
		commentRouter.POST("/:postId", utility.Use(cc.createComment))
		commentRouter.PUT("/:id", utility.Use(cc.updateComment))
		commentRouter.DELETE("/:id", utility.Use(cc.deleteComment))
	}
}

// CreateComment creates a new comment for a post.
// @Summary Create a new comment for a post
// @Description Create a new comment for a post by the authenticated user
// @Tags Comment
// @Accept json
// @Produce json
// @Param CreateCommentDto body Dto.CommentDto true "Comment Data"

// @Success 201 {object} Dto.CommentResponseDto "Successfully created comment"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/comments/:postId [post]
func (cc *CommentController) createComment(ctx *gin.Context) error {
	// Extract the user from the context
	authOID := ctx.MustGet("auth0_id").(string)
	postId := ctx.Param("postId")
	user, err := cc.userService.GetUser(ctx, authOID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	// Bind the incoming JSON to a new CommentDto instance
	data := new(Dto.CommentDto)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return err
	}

	// Create a new Comment instance
	newComment, err := cc.commentService.CreateComment(ctx, user.ID, postId, data.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	ctx.JSON(http.StatusCreated, utility.ProcessResponse(newComment, "Success", "Successfully created comment"))
	return nil
}

// UpdateComment updates an existing comment.
// @Summary Update a comment
// @Description Update an existing comment by the authenticated user
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Param UpdateCommentDto body Dto.CommentDto true "Comment Data"

// @Success 200 {object} Dto.CommentResponseDto "Successfully updated comment"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Comment Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/comments/{id} [put]
func (cc *CommentController) updateComment(ctx *gin.Context) error {
	// Extract the user from the context
	authOID := ctx.MustGet("auth0_id").(string)
	user, err := cc.userService.GetUser(ctx, authOID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utility.ProcessError(err))
		return err
	}

	// Extract the commentID from the path
	commentID := ctx.Param("id")

	// Bind the incoming JSON to a new CommentDto instance
	data := new(Dto.CommentDto)
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, utility.ProcessError(err))
		return err
	}

	// Update the Comment instance
	updatedComment, err := cc.commentService.UpdateComment(ctx, user.ID, commentID, data.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utility.ProcessError(err))
		return err
	}

	// Create a response struct
	//response := Dto.CommentResponseDto{
	//	ID:        updatedComment.ID,
	//	Content:   updatedComment.Content,
	//	User:      updatedComment.User,
	//	CreatedAt: updatedComment.CreatedAt,
	//}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(updatedComment, "Success", "Successfully updated comment"))
	return nil
}

// DeleteComment deletes an existing comment.
// @Summary Delete a comment
// @Description Delete an existing comment by the authenticated user
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"

// @Success 204 "Successfully deleted comment"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 404 {object} Dto.ErrorDTO "Comment Not Found"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/comments/{id} [delete]
func (cc *CommentController) deleteComment(ctx *gin.Context) error {
	// Extract the user from the context
	authOID := ctx.MustGet("auth0_id").(string)
	user, err := cc.userService.GetUser(ctx, authOID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	// Extract the commentID from the path
	commentID := ctx.Param("id")

	// Delete the Comment
	err = cc.commentService.DeleteComment(ctx, user.ID, commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}

	ctx.JSON(http.StatusNoContent, utility.ProcessResponse(nil, "Success", "Successfully deleted comment"))
	return nil
}
