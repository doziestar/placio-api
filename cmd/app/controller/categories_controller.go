package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (cc *CategoryController) RegisterRoutes(router *gin.RouterGroup) {
	categoryRouter := router.Group("/categories")
	{
		categoryRouter.POST("/", utility.Use(cc.createCategory))
		categoryRouter.PATCH("/:id", utility.Use(cc.updateCategory))
		categoryRouter.DELETE("/:id", utility.Use(cc.deleteCategory))
		categoryRouter.GET("/search", utility.Use(cc.searchByCategory)) // new route
	}
}

// @Summary Create a new category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body ent.Category true "Category"
// @Success 200 {object} ent.Category
// @Router /categories/ [post]
func (cc *CategoryController) createCategory(ctx *gin.Context) error {
	id := ctx.GetString("id")
	name := ctx.GetString("name")
	image := ctx.GetString("image")
	category, err := cc.categoryService.CreateCategory(ctx, id, name, image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, category)
	return nil
}

// @Summary Update an existing category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body ent.Category true "Category"
// @Success 200 {object} ent.Category
// @Router /categories/{id} [patch]
func (cc *CategoryController) updateCategory(ctx *gin.Context) error {
	categoryID := ctx.Param("id")
	name := ctx.GetString("name")
	image := ctx.GetString("image")
	category, err := cc.categoryService.UpdateCategory(ctx, categoryID, name, image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, category)
	return nil
}

// @Summary Delete a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} ent.Category
// @Router /categories/{id} [delete]
func (cc *CategoryController) deleteCategory(ctx *gin.Context) error {
	categoryID := ctx.Param("id")
	err := cc.categoryService.DeleteCategory(ctx, categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Category successfully deleted"})
	return nil
}

// @Summary Search by category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param name query string true "Category name"
// @Success 200 {array} ent.Category
// @Router /categories/search [get]
func (cc *CategoryController) searchByCategory(ctx *gin.Context) error {
	categoryName := ctx.Query("name")
	categories, err := cc.categoryService.GetEntitiesByCategory(ctx, categoryName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, categories)
	return nil
}
