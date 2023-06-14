package controller

import (
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
		businessRouter.POST("/:businessID/follow/user/:userID", bc.followUser)
		businessRouter.POST("/:followerID/follow/business/:followedID", bc.followBusiness)
		businessRouter.DELETE("/:businessID/unfollow/user/:userID", bc.unfollowUser)
		businessRouter.DELETE("/:followerID/unfollow/business/:followedID", bc.unfollowBusiness)
		businessRouter.GET("/:businessID/followed-contents", bc.getFollowedContents)
		businessRouter.POST("/", bc.createBusinessAccount)
		businessRouter.GET("/:businessAccountID", bc.getBusinessAccount)
		businessRouter.PUT("/:businessAccountID", bc.updateBusinessAccount)
		businessRouter.DELETE("/:businessAccountID", bc.deleteBusinessAccount)
		businessRouter.GET("/user/:userID", bc.getUserBusinessAccounts)
		businessRouter.POST("/:businessAccountID/user/:userID", bc.associateUserWithBusinessAccount)
		businessRouter.DELETE("/:businessAccountID/user/:userID", bc.removeUserFromBusinessAccount)
		businessRouter.PUT("/:businessAccountID/user/:currentOwnerID/:newOwnerID", bc.transferBusinessAccountOwnership)
		businessRouter.GET("/:businessAccountID/users", bc.getBusinessAccountsForUser)
		businessRouter.GET("/:businessAccountID/users", bc.getUsersForBusinessAccount)
		businessRouter.GET("/", bc.listBusinessAccounts)
	}
}

// @Summary Follow a user by a business
// @ID follow-user
// @Produce json
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
// @Produce json
// @Param followerID path string true "Follower Business ID"
// @Param followedID path string true "Followed Business ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{followerID}/follow/business/{followedID} [post]
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
// @Produce json
// @Param businessID path string true "Business ID"
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{businessID}/unfollow/user/{userID} [delete]
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
// @Produce json
// @Param followerID path string true "Follower Business ID"
// @Param followedID path string true "Followed Business ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{followerID}/unfollow/business/{followedID} [delete]
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
// @Produce json
// @Param businessID path string true "Business ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{businessID}/followed-contents [get]
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
// @Accept  json
// @Produce  json
// @Param ent.Business body ent.Business true "Business Account Data"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router / [post]
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
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{businessAccountID} [get]
func (bc *BusinessAccountController) getBusinessAccount(c *gin.Context) {
	businessAccountID := c.Param("businessAccountID")
	business, err := bc.service.GetBusinessAccount(c, businessAccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, business)
}

// @Summary Update a business account
// @ID update-business-account
// @Accept  json
// @Produce  json
// @Param businessAccountID path string true "Business Account ID"
// @Param ent.Business body ent.Business true "Business Account Data"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{businessAccountID} [put]
func (bc *BusinessAccountController) updateBusinessAccount(c *gin.Context) {
	//businessAccountID := c.Param("businessAccountID")
	//var businessData Dto.BusinessDto
	//if err := c.ShouldBindJSON(&businessData); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//businessData.ID = businessAccountID
	//business, err := bc.service.UpdateBusinessAccount(c, &businessData)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//c.JSON(http.StatusOK, business)
}

// @Summary Delete a business account
// @ID delete-business-account
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{businessAccountID} [delete]
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
// @Produce json
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /user/{userID} [get]
func (bc *BusinessAccountController) getUserBusinessAccounts(c *gin.Context) {
	// Implementation...
}

// @Summary Associate user with business account
// @ID associate-user-business-account
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{businessAccountID}/user/{userID} [post]
func (bc *BusinessAccountController) associateUserWithBusinessAccount(c *gin.Context) {
	// Implementation...
}

// @Summary Remove user from business account
// @ID remove-user-business-account
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{businessAccountID}/user/{userID} [delete]
func (bc *BusinessAccountController) removeUserFromBusinessAccount(c *gin.Context) {
	// Implementation...
}

// @Summary Transfer business account ownership
// @ID transfer-business-account-ownership
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param currentOwnerID path string true "Current Owner ID"
// @Param newOwnerID path string true "New Owner ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{businessAccountID}/user/{currentOwnerID}/{newOwnerID} [put]
func (bc *BusinessAccountController) transferBusinessAccountOwnership(c *gin.Context) {
	// Implementation...
}

// @Summary Get business accounts for a user
// @ID get-business-accounts-for-user
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{businessAccountID}/users [get]
func (bc *BusinessAccountController) getBusinessAccountsForUser(c *gin.Context) {
	// Implementation...
}

// @Summary Get users for a business account
// @ID get-users-business-account
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /{businessAccountID}/users [get]
func (bc *BusinessAccountController) getUsersForBusinessAccount(c *gin.Context) {
	// Implementation...
}

// @Summary List all business accounts
// @ID list-business-accounts
// @Produce json
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router / [get]
func (bc *BusinessAccountController) listBusinessAccounts(c *gin.Context) {
	// Implementation...
}
