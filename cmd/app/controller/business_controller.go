package controller

import (
	"log"
	"net/http"
	"placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"

	"github.com/gin-gonic/gin"
)

type BusinessAccountController struct {
	service service.BusinessAccountService
}

func NewBusinessAccountController(service service.BusinessAccountService) *BusinessAccountController {
	return &BusinessAccountController{service: service}
}

func (bc *BusinessAccountController) RegisterRoutes(router *gin.RouterGroup) {
	businessRouter := router.Group("/business")
	{
		businessRouter.POST("/:businessAccountID/follow/user/:userID", bc.followUser)
		businessRouter.POST("/:businessAccountID/follow/business/:followedID", bc.followBusiness)
		businessRouter.DELETE("/:businessAccountID/unfollow/user/:userID", bc.unfollowUser)
		businessRouter.DELETE("/:businessAccountID/unfollow/business/:followedID", bc.unfollowBusiness)
		businessRouter.GET("/:businessAccountID/followed-contents", bc.getFollowedContents)
		businessRouter.POST("/", bc.createBusinessAccount)
		businessRouter.GET("/user-business-account", bc.getUserBusinessAccounts)
		businessRouter.GET("/:businessAccountID", bc.getBusinessAccount)
		businessRouter.PATCH("/:businessAccountID", bc.updateBusinessAccount)
		businessRouter.DELETE("/:businessAccountID", bc.deleteBusinessAccount)
		businessRouter.POST("/:businessAccountID/user/:userID", bc.associateUserWithBusinessAccount)
		businessRouter.DELETE("/:businessAccountID/user/:userID", bc.removeUserFromBusinessAccount)
		businessRouter.PUT("/:businessAccountID/user/:currentOwnerID/:newOwnerID", bc.transferBusinessAccountOwnership)
		businessRouter.GET("/:businessAccountID/users", bc.getBusinessAccountsForUser)
		businessRouter.GET("/:businessAccountID/associated-users", bc.getUsersForBusinessAccount)
		businessRouter.GET("/", bc.listBusinessAccounts)
	}
}

// @Summary Follow a user by a business
// @ID follow-user
// @Produce json
// @Description Retrieve all comments for a given post
// @Tags Business
// @Accept json
// @Param businessID path string true "Business ID"
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessID}/follow/user/{userID} [post]
func (bc *BusinessAccountController) followUser(c *gin.Context) {
	businessID := c.Param("businessID")
	userID := c.Param("userID")

	if err := bc.service.FollowUser(c, businessID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully followed user"})
}

// @Summary Follow a business by another business
// @ID follow-business
// @Tags Business
// @Produce json
// @Param followerBusinessID path string true "Follower Business ID"
// @Param followedBusinessID path string true "Followed Business ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{followerBusinessID}/follow/business/{followedID} [post]
func (bc *BusinessAccountController) followBusiness(c *gin.Context) {
	followerID := c.Param("followerID")
	followedID := c.Param("followedID")

	if err := bc.service.FollowBusiness(c, followerID, followedID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully followed business"})
}

// @Summary Unfollow a user by a business
// @ID unfollow-user
// @Tags Business
// @Produce json
// @Param businessID path string true "Business ID"
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessID}/unfollow/user/{userID} [delete]
func (bc *BusinessAccountController) unfollowUser(c *gin.Context) {
	businessID := c.Param("businessID")
	userID := c.Param("userID")
	err := bc.service.UnfollowUser(c, businessID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully unfollowed the user"})
}

// @Summary Unfollow a business by another business
// @ID unfollow-business
// @Tags Business
// @Produce json
// @Param followerBusinessID path string true "Follower Business ID"
// @Param followedBusinessID path string true "Followed Business ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{followerBusinessID}/unfollow/business/{followedBusinessID} [delete]
func (bc *BusinessAccountController) unfollowBusiness(c *gin.Context) {
	followerID := c.Param("followerID")
	followedID := c.Param("followedID")
	err := bc.service.UnfollowBusiness(c, followerID, followedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully unfollowed the business"})
}

// @Summary Get contents followed by a business
// @ID get-followed-contents
// @Tags Business
// @Produce json
// @Param businessID path string true "Business ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessID}/followed-contents [get]
func (bc *BusinessAccountController) getFollowedContents(c *gin.Context) {
	businessID := c.Param("businessID")
	posts, err := bc.service.GetFollowedContents(c, businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// @Summary Create a business account
// @ID create-business-account
// @Tags Business
// @Accept  json
// @Produce  json
// @Param Dto.BusinessDto body Dto.BusinessDto true "Business Account Data"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/ [post]
func (bc *BusinessAccountController) createBusinessAccount(c *gin.Context) {
	var businessData Dto.BusinessDto
	if err := c.ShouldBindJSON(&businessData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	business, err := bc.service.CreateBusinessAccount(c, &businessData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, business)
}

// @Summary Get a business account
// @ID get-business-account
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID} [get]
func (bc *BusinessAccountController) getBusinessAccount(c *gin.Context) {
	businessAccountID := c.Param("businessAccountID")
	business, err := bc.service.GetBusinessAccount(c, businessAccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, business)
}

// updateBusinessAccount updates a business's details.
// @Summary Update a business's details
// @Description Get a business's details by ID
// @Tags Business
// @Accept json
// @Produce json
// @Param id path string true "Business ID"
// @Security Bearer
// @Success 200 {object} ent.Business "Successfully retrieved business"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/business/ [patch]
func (bc *BusinessAccountController) updateBusinessAccount(ctx *gin.Context) {
	businessId := ctx.Param("businessAccountID")
	if businessId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Business ID required",
		})
		return
	}

	var business map[string]interface{}
	if err := ctx.ShouldBindJSON(&business); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	businessData, err := bc.service.UpdateBusinessAccount(ctx, businessId, business)
	if err != nil {
		//if errors.Is(err, ent.) {
		//	ctx.JSON(http.StatusNotFound, gin.H{
		//		"error": "Business not found",
		//	})
		//	return
		//}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, businessData)
}

// @Summary Delete a business account
// @ID delete-business-account
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID} [delete]
func (bc *BusinessAccountController) deleteBusinessAccount(c *gin.Context) {
	//businessAccountID := c.Param("businessAccountID")
	//err := bc.service.DeleteBusinessAccount(c, businessAccountID)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted the business account"})
}

// @Summary Get user business accounts
// @ID get-user-business-accounts
// @Tags Business
// @Produce json
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/user-business-account [get]
func (bc *BusinessAccountController) getUserBusinessAccounts(c *gin.Context) {
	log.Println("Get user business accounts")
	businessAccount, err := bc.service.GetUserBusinessAccounts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"businessAccounts": businessAccount})
}

// @Summary Associate user with business account
// @ID associate-user-business-account
// @Produce json
// @Tags Business
// @Param businessAccountID path string true "Business Account ID"
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID}/user/{userID} [post]
func (bc *BusinessAccountController) associateUserWithBusinessAccount(c *gin.Context) {
	// Implementation...
}

// @Summary Remove user from business account
// @ID remove-user-business-account
// @Produce json
// @Tags Business
// @Param businessAccountID path string true "Business Account ID"
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID}/user/{userID} [delete]
func (bc *BusinessAccountController) removeUserFromBusinessAccount(c *gin.Context) {
	// Implementation...
}

// @Summary Transfer business account ownership
// @ID transfer-business-account-ownership
// @Produce json
// @Tags Business
// @Param businessAccountID path string true "Business Account ID"
// @Param currentOwnerID path string true "Current Owner ID"
// @Param newOwnerID path string true "New Owner ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID}/user/{currentOwnerID}/{newOwnerID} [put]
func (bc *BusinessAccountController) transferBusinessAccountOwnership(c *gin.Context) {
	// Implementation...
}

// @Summary Get business accounts for a user
// @ID get-business-accounts-for-user
// @Produce json
// @Tags Business
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID}/users [get]
func (bc *BusinessAccountController) getBusinessAccountsForUser(c *gin.Context) {
	// Implementation...
}

// @Summary Get users for a business account
// @ID get-users-business-account
// @Produce json
// @Tags Business
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID}/users [get]
func (bc *BusinessAccountController) getUsersForBusinessAccount(c *gin.Context) {
	// Implementation...
}

// @Summary List all business accounts
// @ID list-business-accounts
// @Produce json
// @Tags Business
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/ [get]
func (bc *BusinessAccountController) listBusinessAccounts(c *gin.Context) {
	// Implementation...
}
