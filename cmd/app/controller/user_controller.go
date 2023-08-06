package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	_ "placio-app/Dto"
	"placio-app/ent"
	_ "placio-app/ent"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	userService service.UserService
	cache       utility.RedisClient
}

func NewUserController(userService service.UserService, cache utility.RedisClient) *UserController {
	return &UserController{userService: userService, cache: cache}
}

func (uc *UserController) RegisterRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("/", utility.Use(uc.GetUser))
		userRouter.GET("/followers", utility.Use(uc.getFollowers))
		userRouter.GET("/:id/followers", utility.Use(uc.getFollowersByUserID))
		userRouter.GET("/likes", utility.Use(uc.getLikes))
		userRouter.GET("/:id/likes", utility.Use(uc.getUserLikesUserID))
		userRouter.PATCH("/", utility.Use(uc.UpdateUser))
		userRouter.GET("/posts", utility.Use(uc.GetPostsByUser))
		userRouter.PATCH("/:id/userinfo", utility.Use(uc.updateAuth0UserInformation))
		userRouter.PATCH("/:id/metadata", utility.Use(uc.updateAuth0UserMetadata))
		userRouter.PATCH("/:id/appdata", utility.Use(uc.updateAuth0AppMetadata))
		userRouter.POST("/business-account", utility.Use(uc.createBusinessAccount))
		userRouter.GET("/:id/business-accounts", utility.Use(uc.getUserBusinessAccounts))
		userRouter.POST("/business-account/:businessAccountID/user/:userID/association", utility.Use(uc.associateUserWithBusinessAccount))
		userRouter.DELETE("/business-account/:businessAccountID/user/:userID/association", utility.Use(uc.removeUserFromBusinessAccount))
		userRouter.GET("/business-account/:businessAccountID/users", utility.Use(uc.getUsersForBusinessAccount))
		// userRouter.GET("/:userID/business-accounts", utility.Use(uc.GetBusinessAccountsForUser))
		userRouter.POST("/business-account/:businessAccountID/user/:userID/role/:action", utility.Use(uc.canPerformAction))
		userRouter.POST("/follow/user/:followerID/:followedID", utility.Use(uc.followUser))
		userRouter.POST("/follow/business/:followerID/:businessID", utility.Use(uc.followBusiness))
		userRouter.DELETE("/unfollow/user/:followerID/:followedID", utility.Use(uc.unfollowUser))
		userRouter.DELETE("/unfollow/business/:followerID/:businessID", utility.Use(uc.unfollowBusiness))
		userRouter.GET("/followed-contents/:userID", utility.Use(uc.getFollowedContents))
	}
}

// @Summary Follow a user
// @Description Follow a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param followerID path string true "ID of the follower"
// @Param followedID path string true "ID of the user to follow"
// @Success 200 {object} Dto.Response "Successfully followed user"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{followerID}/follow/user/{followedID} [post]
func (uc *UserController) followUser(ctx *gin.Context) error {
	followerID := ctx.Param("followerID")
	followedID := ctx.Param("followedID")

	err := uc.userService.FollowUser(ctx, followerID, followedID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully followed user",
	})
	return nil
}

// @Summary Get Followers
// @Description Get followers of a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param nextPageToken query string false "Next page token"
// @Param limit query string false "Limit"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /api/v1/users/followers [get]
func (uc *UserController) getFollowers(ctx *gin.Context) error {
	userId := ctx.GetString("user")
	nextPageToken := ctx.Query("nextPageToken")
	limit := ctx.DefaultQuery("limit", "10")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}
	followers, err := uc.userService.GetFollowers(ctx, userId, nextPageToken, limitInt)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(followers, "success", "Successfully retrieved followers", nextPageToken))
	return nil
}

// @Summary Get User Likes
// @Description Get followers of a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param nextPageToken query string false "Next page token"
// @Param limit query string false "Limit"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /api/v1/users/followers [get]
func (uc *UserController) getLikes(ctx *gin.Context) error {
	userId := ctx.GetString("user")
	nextPageToken := ctx.Query("nextPageToken")
	limit := ctx.DefaultQuery("limit", "10")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}
	followers, err := uc.userService.GetLikes(ctx, userId, nextPageToken, limitInt)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(followers, "success", "Successfully retrieved followers", nextPageToken))
	return nil
}

// @Summary Get User Follwers By User ID
// @Description Get followers of a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param userID path string true "ID of the user"
// @Param nextPageToken query string false "Next page token"
// @Param limit query string false "Limit"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /api/v1/users/{userID}/followers [get]
func (uc *UserController) getFollowersByUserID(ctx *gin.Context) error {
	userID := ctx.Param("userID")
	nextPageToken := ctx.Query("nextPageToken")
	limit := ctx.DefaultQuery("limit", "10")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}
	followers, err := uc.userService.GetFollowers(ctx, userID, nextPageToken, limitInt)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(followers, "success", "Successfully retrieved followers", nextPageToken))
	return nil
}

// @Summary Get User Likes By User ID
// @Description Get followers of a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param userID path string true "ID of the user"
// @Param nextPageToken query string false "Next page token"
// @Param limit query string false "Limit"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /api/v1/users/{userID}/followers [get]
func (uc *UserController) getUserLikesUserID(ctx *gin.Context) error {
	userID := ctx.Param("id")
	nextPageToken := ctx.Query("nextPageToken")
	limit := ctx.DefaultQuery("limit", "10")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}
	followers, err := uc.userService.GetLikes(ctx, userID, nextPageToken, limitInt)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(followers, "success", "Successfully retrieved followers", nextPageToken))
	return nil
}

// @Summary Follow a business
// @Description Follow a business by its ID
// @Tags User
// @Accept json
// @Produce json
// @Param followerID path string true "ID of the follower"
// @Param businessID path string true "ID of the business to follow"
// @Success 200 {object} Dto.Response "Successfully followed business"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{followerID}/follow/business/{businessID} [post]
func (uc *UserController) followBusiness(ctx *gin.Context) error {
	followerID := ctx.Param("followerID")
	businessID := ctx.Param("businessID")

	err := uc.userService.FollowBusiness(ctx, followerID, businessID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully followed business",
	})
	return nil
}

// @Summary Unfollow a user
// @Description Unfollow a user by their ID
// @Tags User
// @Accept json
// @Produce json
// @Param followerID path string true "ID of the follower"
// @Param followedID path string true "ID of the user to unfollow"
// @Success 200 {object} Dto.Response "Successfully unfollowed user"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{followerID}/unfollow/user/{followedID} [delete]
func (uc *UserController) unfollowUser(ctx *gin.Context) error {
	followerID := ctx.Param("followerID")
	followedID := ctx.Param("followedID")

	err := uc.userService.UnfollowUser(ctx, followerID, followedID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully unfollowed user",
	})
	return nil
}

// @Summary Unfollow a business
// @Description Unfollow a business by its ID
// @Tags User
// @Accept json
// @Produce json
// @Param followerID path string true "ID of the follower"
// @Param businessID path string true "ID of the business to unfollow"
// @Success 200 {object} string "Successfully unfollowed business"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{followerID}/unfollow/business/{businessID} [delete]
func (uc *UserController) unfollowBusiness(ctx *gin.Context) error {
	followerID := ctx.Param("followerID")
	businessID := ctx.Param("businessID")

	err := uc.userService.UnfollowBusiness(ctx, followerID, businessID)

	err = uc.userService.UnfollowBusiness(ctx, followerID, businessID)
	if err != nil {
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Successfully unfollowed business",
	})
	return nil
}

// @Summary Get followed contents
// @Description Get posts of users and businesses followed by the user
// @Tags User
// @Accept json
// @Produce json
// @Param userID path string true "ID of the user"
// @Success 200 {array} ent.Post "Successfully retrieved posts"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{userID}/followed-contents [get]
func (uc *UserController) getFollowedContents(ctx *gin.Context) error {
	userID := ctx.Param("userID")

	posts, err := uc.userService.GetFollowedContents(ctx, userID)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, posts)
	return nil
}

// @Summary Update a user's information
// @Description Update a user's information by their Auth0 ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User Auth0 ID"
// @Param userData body models.Auth0UserData true "User data to update"
// @Success 200 {object} ent.User "Successfully updated user information"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{id}/userinfo [patch]
func (uc *UserController) updateAuth0UserInformation(ctx *gin.Context) error {
	auth0ID := ctx.MustGet("auth0_id").(string)
	if auth0ID == "" {
		return nil
	}

	//userToUpdateID := ctx.Param("id")

	// split the auth0ID by |
	provider := utility.SplitString(auth0ID, "|")[0]
	log.Println("provider: ", provider)
	if provider != "auth0" {
		return errors.New("user is not authorized to update user information")
	}

	var userData *models.Auth0UserData
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		userData = nil
	}

	user, err := uc.userService.UpdateAuth0UserInformation(auth0ID, userData)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(user, "success", "Successfully updated user", ""))
	return nil
}

// @Summary Update a user's metadata
// @Description Update a user's metadata by their Auth0 ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User Auth0 ID"
// @Param userMetaData body models.Metadata true "User metadata to update"
// @Success 200 {object} models.User "Successfully updated user metadata"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{id}/metadata [patch]
func (uc *UserController) updateAuth0UserMetadata(ctx *gin.Context) error {
	log.Println("updateAuth0UserMetadata")
	auth0ID := ctx.MustGet("auth0_id").(string)
	if auth0ID == "" {
		return nil
	}

	var userMetaData *models.Metadata
	if err := ctx.ShouldBindJSON(&userMetaData); err != nil {
		userMetaData = nil
	}

	user, err := uc.userService.UpdateAuth0UserMetadata(auth0ID, userMetaData)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(user, "success", "Successfully updated user metadata", ""))
	return nil
}

// @Summary Update a user's app metadata
// @Description Update a user's app metadata by their Auth0 ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User Auth0 ID"
// @Param appData body models.AppMetadata true "App metadata to update"
// @Success 200 {object} ent.User "Successfully updated app metadata"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{id}/appdata [patch]
func (uc *UserController) updateAuth0AppMetadata(ctx *gin.Context) error {
	auth0ID := ctx.MustGet("auth0_id").(string)
	if auth0ID == "" {
		return nil
	}

	var appData *models.AppMetadata
	if err := ctx.ShouldBindJSON(&appData); err != nil {
		appData = nil
	}

	user, err := uc.userService.UpdateAuth0AppMetadata(auth0ID, appData)
	if err != nil {

		return err
	}

	ctx.JSON(http.StatusOK, utility.ProcessResponse(user, "success", "Successfully updated app metadata", ""))
	return nil
}

// GetUser gets a user's details.
// @Summary Get a user's details
// @Description Get a user's details by their Auth0 ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User Auth0 ID"
// @Success 200 {object} ent.User "Successfully retrieved user"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/ [get]
func (uc *UserController) GetUser(ctx *gin.Context) error {
	auth0ID := ctx.MustGet("auth0_id").(string)
	if auth0ID == "" {
		return errors.New("user Auth0 ID required")
	}

	return utility.GetDataFromCache[*ent.User](ctx, &uc.cache, uc.userService.GetUser, auth0ID, fmt.Sprintf("user:%s", auth0ID))

}

// UpdateUser updates a user's details.
// @Summary Update a user's details
// @Description Get a user's details by their Auth0 ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User Auth0 ID"
// @Success 200 {object} ent.User "Successfully retrieved user"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/ [patch]
func (uc *UserController) UpdateUser(ctx *gin.Context) error {
	userId := ctx.MustGet("user").(string)
	log.Println("UpdateUser", ctx.Request.URL.Path, ctx.Request.Method, userId)
	if userId == "" {
		return nil
	}

	var user map[string]interface{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil
	}

	userData, err := uc.userService.UpdateUser(ctx, userId, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return err
	}

	ctx.JSON(http.StatusOK, userData)
	return nil
}

// createBusinessAccount creates a new business account and associates it with the user.
// @Summary Create a new business account
// @Description Create a new business account for the authenticated user
// @Tags User
// @Accept json
// @Produce json
// @Param name body string true "Business Account Name"
// @Success 201 {object} ent.Business "Successfully created business account"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/business-account [post]
func (uc *UserController) createBusinessAccount(ctx *gin.Context) error {
	auth0ID := ctx.MustGet("user").(string)

	var businessAccount models.BusinessAccount
	if err := ctx.ShouldBindJSON(&businessAccount); err != nil {
		return err
	}

	role := "admin" // Define role or get it from somewhere
	user, err := uc.userService.GetUser(ctx, auth0ID)
	if err != nil {
		return err
	}

	log.Println("CreateBusinessAccount", ctx.Request.URL.Path, ctx.Request.Method, auth0ID, user.ID, businessAccount.Name, role)

	//newBusinessAccount, err := uc.userService.CreateBusinessAccount(user.UserID, businessAccount.Name, role)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"error": "Could not create business account",
	//	})
	//	return err
	//}
	//
	//ctx.JSON(http.StatusCreated, newBusinessAccount)
	return nil
}

// GetPostsByUser retrieves all posts by the user.
// @Summary Retrieve posts by user
// @Description Get posts by the authenticated user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} Dto.PostResponseDto "Successfully retrieved posts"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{userID}/posts [get]
func (uc *UserController) GetPostsByUser(ctx *gin.Context) error {
	log.Println("GetPostsByUser", ctx.Request.URL.Path, ctx.Request.Method)
	// Get the user ID from the URL path parameters
	userID := ctx.MustGet("user").(string)
	log.Println("GetPostsByUser and userId", ctx.Request.URL.Path, ctx.Request.Method, userID)

	// Call the service method to get the posts
	posts, err := uc.userService.GetPostsByUser(ctx, userID)
	log.Println("GetPostsByUser and userId", ctx.Request.URL.Path, ctx.Request.Method, userID, posts)
	if err != nil {
		return err
	}
	//
	//// If posts were found, convert them to the DTO format
	//postsDto := make([]Dto.PostResponseDto, len(posts))
	//for i, post := range posts {
	//	postsDto[i] = Dto.PostResponseDto{
	//		ID:      post.ID,
	//		Content: post.Content,
	//		User:    post.Edges.User,
	//		//Business:  post.Edges.Business,
	//		CreatedAt: post.CreatedAt,
	//		Medias:    make([]Dto.MediaDto, len(post.Edges.Medias)),
	//	}
	//
	//	for j, media := range post.Edges.Medias {
	//		postsDto[i].Medias[j] = Dto.MediaDto{
	//			Type: media.MediaType,
	//			URL:  media.URL,
	//		}
	//	}
	//}

	// Return the posts in the response
	ctx.JSON(http.StatusOK, posts)
	return nil
}

// getUserBusinessAccounts retrieves all the business accounts associated with a specific user.
// @Summary Get all business accounts for a user
// @Description Get all business accounts associated with a specific user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} ent.Business "Successfully retrieved business accounts"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{id}/business-accounts [get]
func (uc *UserController) getUserBusinessAccounts(ctx *gin.Context) error {
	userID := ctx.Param("id")
	if userID == "" {
		return nil
	}

	//businessAccounts, err := uc.userService.GetUserBusinessAccounts(userID)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"error": "Could not retrieve business accounts",
	//	})
	//	return err
	//}
	//
	//ctx.JSON(http.StatusOK, businessAccounts)
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
// @Success 200 {array} ent.User "Successfully retrieved users"
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
// @Success 200 {array} ent.Business "Successfully retrieved business accounts"
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
// @Success 200 {object} Dto.PermissionDTO "Successfully checked permission"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/users/{userID}/business-account/{businessAccountID}/can-perform-action [post]
func (uc *UserController) canPerformAction(ctx *gin.Context) error {
	// TODO: Implement this method
	return nil
}
