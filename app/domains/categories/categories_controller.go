package categories

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/utility"
	"placio-pkg/errors"
	"placio-pkg/middleware"
	"strconv"
)

type CategoryController struct {
	categoryService CategoryService
}

func NewCategoryController(categoryService CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (cc *CategoryController) RegisterRoutes(router *gin.RouterGroup) {
	categoryRouter := router.Group("/categories")
	{
		categoryRouter.POST("/", middleware.ErrorMiddleware(cc.createCategory))
		categoryRouter.PATCH("/:id", middleware.ErrorMiddleware(cc.updateCategory))
		categoryRouter.DELETE("/:id", middleware.ErrorMiddleware(cc.deleteCategory))
		categoryRouter.GET("/search", middleware.ErrorMiddleware(cc.searchByCategory))
		categoryRouter.GET("/", middleware.ErrorMiddleware(cc.getAllCategories))

		// New routes for User, Business, and Place
		categoryRouter.POST("/:categoryID/users/:userID", middleware.ErrorMiddleware(cc.assignUserToCategory))
		categoryRouter.GET("/:categoryID/users", middleware.ErrorMiddleware(cc.getUsersByCategory))
		categoryRouter.POST("/:categoryID/businesses/:businessID", middleware.ErrorMiddleware(cc.assignBusinessToCategory))
		categoryRouter.GET("/:categoryID/businesses", middleware.ErrorMiddleware(cc.getBusinessesByCategory))
		categoryRouter.POST("/:categoryID/places/:placeID", middleware.ErrorMiddleware(cc.assignPlaceToCategory))
		categoryRouter.GET("/:categoryID/places", middleware.ErrorMiddleware(cc.getPlacesByCategory))
	}
}

// @Summary Create a new category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body Dto.CreateCategoryRequest true "Category"
// @Success 200 {object} ent.Category
// @Router /categories/ [post]
func (cc *CategoryController) createCategory(ctx *gin.Context) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	name := form.Value["name"][0]
	image := form.File["files"]
	icon := form.Value["icon"][0]

	category, err := cc.categoryService.CreateCategory(ctx, icon, name, image)
	if err != nil {
		return errors.LogAndReturnError(err)
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
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	name := form.Value["name"][0]
	image := form.File["files"]
	icon := form.Value["icon"][0]
	category, err := cc.categoryService.UpdateCategory(ctx, categoryID, name, image, icon)
	if err != nil {
		return errors.LogAndReturnError(err)
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
		return errors.LogAndReturnError(err)
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
	nextPageToken := ctx.Query("nextPageToken")
	limiter := ctx.Query("limit")

	limit, err := strconv.Atoi(limiter)
	if err != nil {
		return err
	}
	categories, err := cc.categoryService.GetEntitiesByCategory(ctx, categoryName, nextPageToken, limit)
	if err != nil {
		return errors.LogAndReturnError(err)
	}
	ctx.JSON(http.StatusOK, categories)
	return nil
}

// @Summary Get all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Param name query string true "Category name"
// @Success 200 {array} ent.Category
// @Router /categories/ [get]
func (cc *CategoryController) getAllCategories(ctx *gin.Context) error {
	categories, err := cc.categoryService.GetAllCategories(ctx)
	if err != nil {
		return errors.LogAndReturnError(err)
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
		return errors.LogAndReturnError(err)
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
		return errors.LogAndReturnError(err)
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
		return errors.LogAndReturnError(err)
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
	categoryID := ctx.Param("categoryName")
	nextPageToken := ctx.Query("nextPageToken")
	limiter := ctx.Query("limit")

	limit, err := strconv.Atoi(limiter)
	businesses, nextPageToken, err := cc.categoryService.GetBusinessesByCategory(ctx, categoryID, nextPageToken, limit)
	if err != nil {
		return errors.LogAndReturnError(err)
	}
	ctx.JSON(http.StatusOK, utility.ProcessResponse(businesses, "success", "retrieve businesses successfully", nextPageToken))
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
		return errors.LogAndReturnError(err)
	}
	ctx.JSON(http.StatusOK, assignment)
	return nil
}

// @Summary Get places by category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param categoryName path string true "Category name"
// @Success 200 {array} ent.Place
// @Router /categories/{categoryID}/places [get]
func (cc *CategoryController) getPlacesByCategory(ctx *gin.Context) error {
	categoryID := ctx.Param("categoryName")
	nextPageToken := ctx.Query("nextPageToken")
	limiter := ctx.Query("limit")

	limit, err := strconv.Atoi(limiter)
	places, nextPageToken, err := cc.categoryService.GetPlacesByCategory(ctx, categoryID, nextPageToken, limit)
	if err != nil {
		return errors.LogAndReturnError(err)
	}
	ctx.JSON(http.StatusOK, utility.ProcessResponse(places, "success", "retrieve places successfully", nextPageToken))
	return nil
}
