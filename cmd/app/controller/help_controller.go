package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/ent"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"
)

type HelpController struct {
	helpService service.HelpService
}

func (uc *HelpController) RegisterRoutes(router *gin.RouterGroup) {
	// ...
	helpRouter := router.Group("/helps")
	{
		helpRouter.POST("/", utility.Use(uc.createHelp))
		helpRouter.PATCH("/:id/resolve", utility.Use(uc.resolveHelp))
		helpRouter.DELETE("/:id", utility.Use(uc.deleteHelp))
	}
	// ...
}

// @Summary Create a help request
// @Description Create a new help request
// @Tags Help
// @Accept json
// @Produce json
// @Param userID body string true "ID of the user creating the help request"
// @Param category body string true "Category of the help request"
// @Param subject body string true "Subject of the help request"
// @Param body body string true "Body of the help request"
// @Param media body string false "Optional media associated with the help request"
// @Security Bearer
// @Success 200 {object} ent.Help "Successfully created help request"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/helps [post]
func (uc *HelpController) createHelp(ctx *gin.Context) error {
	userID := ctx.MustGet("user_d").(string)
	var helpDto ent.Help
	if err := ctx.ShouldBindJSON(&helpDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	help, err := uc.helpService.CreateHelp(ctx, userID, helpDto.Category, helpDto.Subject, helpDto.Body, helpDto.Media)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, help)
	return nil
}

// @Summary Resolve a help request
// @Description Resolve an existing help request
// @Tags Help
// @Accept json
// @Produce json
// @Param id path string true "ID of the help request to resolve"
// @Security Bearer
// @Success 200 {object} ent.Help "Successfully resolved help request"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/helps/{id}/resolve [patch]
func (uc *HelpController) resolveHelp(ctx *gin.Context) error {
	helpID := ctx.Param("id")
	help, err := uc.helpService.ResolveHelp(ctx, helpID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, help)
	return nil
}

// @Summary Delete a help request
// @Description Delete an existing help request
// @Tags Help
// @Accept json
// @Produce json
// @Param id path string true "ID of the help request to delete"
// @Security Bearer
// @Success 200 {object} Dto.Response "Successfully deleted help request"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/helps/{id} [delete]
func (uc *HelpController) deleteHelp(ctx *gin.Context) error {
	helpID := ctx.Param("id")
	err := uc.helpService.DeleteHelp(ctx, helpID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Help request successfully deleted"})
	return nil
}
