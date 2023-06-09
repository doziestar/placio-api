package controller

import (
	"errors"
	"log"
	"net/http"
	_ "placio-app/Dto"
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
		userRouter.GET("/", utility.Use(uc.GetUser))
		userRouter.PATCH("/:id", utility.Use(uc.updateAuth0UserData))
		userRouter.POST("/business-account", utility.Use(uc.createBusinessAccount))
		userRouter.GET("/:id/business-accounts", utility.Use(uc.getUserBusinessAccounts))
		userRouter.POST("/:userID/business-account/:businessAccountID/association", utility.Use(uc.associateUserWithBusinessAccount))
		userRouter.DELETE("/:userID/business-account/:businessAccountID/association", utility.Use(uc.removeUserFromBusinessAccount))
		userRouter.GET("/business-account/:businessAccountID/users", utility.Use(uc.getUsersForBusinessAccount))
		// userRouter.GET("/:userID/business-accounts", utility.Use(uc.GetBusinessAccountsForUser))
		userRouter.POST("/:userID/business-account/:businessAccountID/role/:action", utility.Use(uc.canPerformAction))
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
	auth0ID := ctx.MustGet("user").(string)
	log.Println("GetUser", ctx.Request.URL.Path, ctx.Request.Method, auth0ID)
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

// updateAuth0UserData updates a user's details in Auth0.
// @Summary Update a user's details
// @Description Update a user's details by their Auth0 ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User Auth0 ID"
// @Param userData body models.Auth0UserData false "User data to update"
// @Param appData body models.AppMetadata false "App metadata to update"
// @Param userMetaData body models.Metadata false "User metadata to update"
// @Security Bearer
// @Success 200 {object} models.User "Successfully updated user"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{id} [patch]
func (uc *UserController) updateAuth0UserData(ctx *gin.Context) error {
	auth0ID := ctx.MustGet("user").(string)
	log.Println("UpdateAuth0UserData", ctx.Request.URL.Path, ctx.Request.Method, auth0ID)
	if auth0ID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User Auth0 ID required",
		})
		return nil
	}

	var userData *models.Auth0UserData
	var appData *models.AppMetadata
	var userMetaData *models.Metadata
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		userData = nil
	}
	if err := ctx.ShouldBindJSON(&appData); err != nil {
		appData = nil
	}
	if err := ctx.ShouldBindJSON(&userMetaData); err != nil {
		userMetaData = nil
	}

	//user, err := uc.userService.UpdateAuth0UserData(auth0ID, ctx.Request.Header.Get("Authorization"), userData, appData, userMetaData)
	user, err := uc.userService.UpdateAuth0UserData(auth0ID, "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InliX2tBaHRaOWQ3VEt6VHc2cEJVayJ9.eyJpc3MiOiJodHRwczovL2Rldi1xbGIwbHYzZC51cy5hdXRoMC5jb20vIiwic3ViIjoiZnVqdWJFSlBCdGxnMzhGdE51dUM0MTNQaktsN3VwbnhAY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vYXBpLnBhbG5pZ2h0LmNvbSIsImlhdCI6MTY4NjI3NjQ0OSwiZXhwIjoxNjg4ODY4NDQ5LCJhenAiOiJmdWp1YkVKUEJ0bGczOEZ0TnV1QzQxM1BqS2w3dXBueCIsInNjb3BlIjoicmVhZDpldmVudHMgcmVhZDpwbGFjZXMgd3JpdGU6ZXZlbnRzIHdyaXRlIHBsYWNlcyBkZWxldGU6ZXZlbnRzIGRlbGV0ZTpwbGFjZXMgYXBpIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIiwicGVybWlzc2lvbnMiOlsicmVhZDpldmVudHMiLCJyZWFkOnBsYWNlcyIsIndyaXRlOmV2ZW50cyIsIndyaXRlIHBsYWNlcyIsImRlbGV0ZTpldmVudHMiLCJkZWxldGU6cGxhY2VzIiwiYXBpIl19.mvNEB7rxWC09gBBNdMd5sZ-XuLCbOrTM3AQfQCBMnAlkPC_KqYbQ3psMEqInMNksd2DAvPz_RLb_UmZoX4F-tQ-5sHxnswwTIW6UYyP64PDjo-X4VWRqhRa9n0zKPVCHHVlO3iOnEZ7prG0ctVfze5ZziyH3SQ5t-AKa87b4R2JFWJwqUwpDsEqsnmXrtM-2hv33As3_aN_hVmWS8wKJuEw1gOrr_l6tdGD_BqwtxFwgql2vYKivIXrOZfO5WpYlt8ROHi8lcnoG2FAN7zMzDfGgXfh7l4dv4XsUjQmjGR0bkvXn0YGSFN-e31aPc-53RbvkK5rYfFd0k5UmcZA0KA", userData, appData, userMetaData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(user, "success", "Successfully updated user"))
	return nil
}

// createBusinessAccount creates a new business account and associates it with the user.
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
func (uc *UserController) createBusinessAccount(ctx *gin.Context) error {
	auth0ID := ctx.MustGet("user").(string)

	var businessAccount models.BusinessAccount
	if err := ctx.ShouldBindJSON(&businessAccount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return err
	}

	role := "admin" // Define role or get it from somewhere
	user, err := uc.userService.GetUser(auth0ID)

	log.Println("CreateBusinessAccount", ctx.Request.URL.Path, ctx.Request.Method, auth0ID, user.UserID, businessAccount.Name, role)

	newBusinessAccount, err := uc.userService.CreateBusinessAccount(user.UserID, businessAccount.Name, role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create business account",
		})
		return err
	}

	ctx.JSON(http.StatusCreated, newBusinessAccount)
	return nil
}

// getUserBusinessAccounts retrieves all the business accounts associated with a specific user.
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
func (uc *UserController) getUserBusinessAccounts(ctx *gin.Context) error {
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

// associateUserWithBusinessAccount associates a user with a specific business account.
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
func (uc *UserController) associateUserWithBusinessAccount(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}

// removeUserFromBusinessAccount removes a user's association with a specific business account.
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
func (uc *UserController) removeUserFromBusinessAccount(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}

// getUsersForBusinessAccount retrieves all the users associated with a specific business account.
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
func (uc *UserController) getUsersForBusinessAccount(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}

// getBusinessAccountsForUser retrieves all the business accounts a user is associated with.
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
func (uc *UserController) getBusinessAccountsForUser(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}

// canPerformAction checks if a user can perform a certain action based on their role.
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
func (uc *UserController) canPerformAction(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}
