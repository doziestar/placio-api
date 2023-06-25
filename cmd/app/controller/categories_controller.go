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

		// New routes for User, Business, and Place
		categoryRouter.POST("/:categoryID/users/:userID", utility.Use(cc.assignUserToCategory))
		categoryRouter.GET("/:categoryID/users", utility.Use(cc.getUsersByCategory))
		categoryRouter.POST("/:categoryID/businesses/:businessID", utility.Use(cc.assignBusinessToCategory))
		categoryRouter.GET("/:categoryID/businesses", utility.Use(cc.getBusinessesByCategory))
		categoryRouter.POST("/:categoryID/places/:placeID", utility.Use(cc.assignPlaceToCategory))
		categoryRouter.GET("/:categoryID/places", utility.Use(cc.getPlacesByCategory))
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

// @Summary Assign a user to a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param categoryID path string true "Category ID"
// @Param userID path string true "User ID"
// @Success 200 {object} ent.CategoryAssignment
// @Router /categories/{categoryID}/users/{userID} [post]
func (cc *CategoryController) assignUserToCategory(ctx *gin.Context) error {
	userID := ctx.Param("userID")
	categoryID := ctx.Param("categoryID")
	assignment, err := cc.categoryService.AssignUserToCategory(ctx, userID, categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, assignment)
	return nil
}

// @Summary Get users by category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param categoryID path string true "Category ID"
// @Success 200 {array} ent.User
// @Router /categories/{categoryID}/users [get]
func (cc *CategoryController) getUsersByCategory(ctx *gin.Context) error {
	categoryID := ctx.Param("categoryID")
	users, err := cc.categoryService.GetUsersByCategory(ctx, categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, users)
	return nil
}

// @Summary Assign a business to a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param categoryID path string true "Category ID"
// @Param businessID path string true "Business ID"
// @Success 200 {object} ent.CategoryAssignment
// @Router /categories/{categoryID}/businesses/{businessID} [post]
func (cc *CategoryController) assignBusinessToCategory(ctx *gin.Context) error {
	businessID := ctx.Param("businessID")
	categoryID := ctx.Param("categoryID")
	assignment, err := cc.categoryService.AssignBusinessToCategory(ctx, businessID, categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, assignment)
	return nil
}

// @Summary Get businesses by category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param categoryID path string true "Category ID"
// @Success 200 {array} ent.Business
// @Router /categories/{categoryID}/businesses [get]
func (cc *CategoryController) getBusinessesByCategory(ctx *gin.Context) error {
	categoryID := ctx.Param("categoryID")
	businesses, err := cc.categoryService.GetBusinessesByCategory(ctx, categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, businesses)
	return nil
}

// @Summary Assign a place to a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param categoryID path string true "Category ID"
// @Param placeID path string true "Place ID"
// @Success 200 {object} ent.CategoryAssignment
// @Router /categories/{categoryID}/places/{placeID} [post]
func (cc *CategoryController) assignPlaceToCategory(ctx *gin.Context) error {
	placeID := ctx.Param("placeID")
	categoryID := ctx.Param("categoryID")
	assignment, err := cc.categoryService.AssignPlaceToCategory(ctx, placeID, categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, assignment)
	return nil
}

// @Summary Get places by category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param categoryID path string true "Category ID"
// @Success 200 {array} ent.Place
// @Router /categories/{categoryID}/places [get]
func (cc *CategoryController) getPlacesByCategory(ctx *gin.Context) error {
	categoryID := ctx.Param("categoryID")
	places, err := cc.categoryService.GetPlacesByCategory(ctx, categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return err
	}
	ctx.JSON(http.StatusOK, places)
	return nil
}
