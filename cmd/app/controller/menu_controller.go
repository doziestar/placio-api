package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"
)

type MenuController struct {
	menuService service.MenuService
}

func NewMenuController(menuService service.MenuService) *MenuController {
	return &MenuController{menuService: menuService}
}

func (mc *MenuController) RegisterRoutes(router *gin.RouterGroup) {
	menuRouter := router.Group("/menus")
	{
		menuRouter.GET("/:id", utility.Use(mc.getMenu))
		menuRouter.POST("/", utility.Use(mc.createMenu))
		menuRouter.PATCH("/:id", utility.Use(mc.updateMenu))
		menuRouter.DELETE("/:id", utility.Use(mc.deleteMenu))
	}
}

// @Summary Get a menu
// @Description Get a menu by its ID
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path string true "ID of the menu"
// @Success 200 {object} ent.Menu "Successfully retrieved menu"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/menus/{id} [get]
func (mc *MenuController) getMenu(ctx *gin.Context) error {
	id := ctx.Param("id")

	menu, err := mc.menuService.GetMenu(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, menu)
	return nil
}

// @Summary Create a menu
// @Description Create a new menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param menu body map[string]interface{} true "Menu data"
// @Success 200 {object} ent.Menu "Successfully created menu"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/menus/ [post]
func (mc *MenuController) createMenu(ctx *gin.Context) error {
	var menuData map[string]interface{}
	if err := ctx.ShouldBindJSON(&menuData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	menu, err := mc.menuService.CreateMenu(ctx, menuData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, menu)
	return nil
}

// @Summary Update a menu
// @Description Update a menu by its ID
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path string true "ID of the menu"
// @Param menu body map[string]interface{} true "Menu data"
// @Success 200 {object} ent.Menu "Successfully updated menu"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/menus/{id} [patch]
func (mc *MenuController) updateMenu(ctx *gin.Context) error {
	id := ctx.Param("id")
	var menuData map[string]interface{}
	if err := ctx.ShouldBindJSON(&menuData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	menu, err := mc.menuService.UpdateMenu(ctx, id, menuData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, menu)
	return nil
}

// @Summary Delete a menu
// @Description Delete a menu by its ID
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path string true "ID of the menu"
// @Success 200 {object} gin.H "Successfully deleted menu"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/menus/{id} [delete]
func (mc *MenuController) deleteMenu(ctx *gin.Context) error {
	// ... implementation omitted for brevity
	return nil
}
