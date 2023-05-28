package controller

import (
	"errors"
	"net/http"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) RegisterRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("/:id", utility.Use(uc.GetUser))
		userRouter.POST("/business-account", utility.Use(uc.CreateBusinessAccount))
		userRouter.GET("/:id/business-accounts", utility.Use(uc.GetUserBusinessAccounts))
		userRouter.POST("/:userID/business-account/:businessAccountID/association", utility.Use(uc.AssociateUserWithBusinessAccount))
		userRouter.DELETE("/:userID/business-account/:businessAccountID/association", utility.Use(uc.RemoveUserFromBusinessAccount))
		userRouter.GET("/business-account/:businessAccountID/users", utility.Use(uc.GetUsersForBusinessAccount))
		userRouter.GET("/:userID/business-accounts", utility.Use(uc.GetBusinessAccountsForUser))
		userRouter.POST("/:userID/business-account/:businessAccountID/role/:action", utility.Use(uc.CanPerformAction))
	}
}

// GetUser gets a user's details.
// @Summary Get a user's details
// @Description Get a user's details by their Auth0 ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User Auth0 ID"
// @Security Bearer
// @Success 200 {object} models.User "Successfully retrieved user"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{id} [get]
func (uc *UserController) GetUser(ctx *gin.Context) error {
	auth0ID := ctx.Param("id")
	if auth0ID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User Auth0 ID required",
		})
		return nil
	}

	user, err := uc.userService.GetUser(auth0ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return nil
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, user)
	return nil
}


// CreateBusinessAccount creates a new business account and associates it with the user.
// @Summary Create a new business account
// @Description Create a new business account for the authenticated user
// @Tags User
// @Accept json
// @Produce json
// @Param name body string true "Business Account Name"
// @Security Bearer
// @Success 201 {object} models.BusinessAccount "Successfully created business account"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/business-account [post]
func (uc *UserController) CreateBusinessAccount(ctx *gin.Context) error {
	userID := ctx.MustGet("user").(string)
	role := ctx.MustGet("role").(string)

	var businessAccount models.BusinessAccount
	if err := ctx.ShouldBindJSON(&businessAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return err
	}

	newBusinessAccount, err := uc.userService.CreateBusinessAccount(userID, businessAccount.Name, role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create business account",
		})
		return err
	}

	ctx.JSON(http.StatusCreated, newBusinessAccount)
	return nil
}

// GetUserBusinessAccounts retrieves all the business accounts associated with a specific user.
// @Summary Get all business accounts for a user
// @Description Get all business accounts associated with a specific user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Security Bearer
// @Success 200 {array} models.BusinessAccount "Successfully retrieved business accounts"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{id}/business-accounts [get]
func (uc *UserController) GetUserBusinessAccounts(ctx *gin.Context) error {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID required",
		})
		return nil
	}

	businessAccounts, err := uc.userService.GetUserBusinessAccounts(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not retrieve business accounts",
		})
		return err
	}

	ctx.JSON(http.StatusOK, businessAccounts)
	return nil
}

// AssociateUserWithBusinessAccount associates a user with a specific business account.
// @Summary Associate a user with a business account
// @Description Associate a user with a specific business account
// @Tags User
// @Accept json
// @Produce json
// @Param userID path uint true "User ID"
// @Param businessAccountID path uint true "Business Account ID"
// @Param role body string true "Role"
// @Security Bearer
// @Success 204 "Successfully associated user with business account"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{userID}/business-account/{businessAccountID}/association [post]
func (uc *UserController) AssociateUserWithBusinessAccount(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}

// RemoveUserFromBusinessAccount removes a user's association with a specific business account.
// @Summary Remove a user's association with a business account
// @Description Remove a user's association with a specific business account
// @Tags User
// @Accept json
// @Produce json
// @Param userID path uint true "User ID"
// @Param businessAccountID path uint true "Business Account ID"
// @Security Bearer
// @Success 204 "Successfully removed user's association with business account"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{userID}/business-account/{businessAccountID}/association [delete]
func (uc *UserController) RemoveUserFromBusinessAccount(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}

// GetUsersForBusinessAccount retrieves all the users associated with a specific business account.
// @Summary Get all users for a business account
// @Description Get all users associated with a specific business account
// @Tags User
// @Accept json
// @Produce json
// @Param businessAccountID path uint true "Business Account ID"
// @Security Bearer
// @Success 200 {array} models.User "Successfully retrieved users"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/business-account/{businessAccountID}/users [get]
func (uc *UserController) GetUsersForBusinessAccount(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}

// GetBusinessAccountsForUser retrieves all the business accounts a user is associated with.
// @Summary Get all business accounts a user is associated with
// @Description Get all business accounts a user is associated with
// @Tags User
// @Accept json
// @Produce json
// @Param userID path uint true "User ID"
// @Security Bearer
// @Success 200 {array} models.BusinessAccount "Successfully retrieved business accounts"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{userID}/business-accounts [get]
func (uc *UserController) GetBusinessAccountsForUser(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}

// CanPerformAction checks if a user can perform a certain action based on their role.
// @Summary Check if a user can perform an action
// @Description Check if a user can perform a certain action based on their role in the business account
// @Tags User
// @Accept json
// @Produce json
// @Param userID path uint true "User ID"
// @Param businessAccountID path uint true "Business Account ID"
// @Param action body string true "Action"
// @Security Bearer
// @Success 200 {object} Dto.PermissionDTO "Successfully checked permission"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{userID}/business-account/{businessAccountID}/can-perform-action [post]
func (uc *UserController) CanPerformAction(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}
