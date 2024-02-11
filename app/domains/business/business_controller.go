package business

import (
	"errors"
	"log"
	"net/http"
	"placio-app/domains/events_management"
	"placio-app/ent"
	_ "placio-app/ent"
	"placio-app/utility"

	"github.com/gin-gonic/gin"
)

type BusinessAccountController struct {
	service      BusinessAccountService
	eventService events_management.EventService
	cache        utility.RedisClient
}

func NewBusinessAccountController(service BusinessAccountService, cache utility.RedisClient) *BusinessAccountController {
	return &BusinessAccountController{service: service, cache: cache}
}

func (bc *BusinessAccountController) RegisterRoutes(router *gin.RouterGroup) {
	businessRouter := router.Group("/business")
	{
		businessRouter.POST("/:businessAccountID/follow/user/:userID", utility.Use(bc.followUser))
		businessRouter.POST("/:businessAccountID/follow/business/:followedID", utility.Use(bc.followBusiness))
		businessRouter.DELETE("/:businessAccountID/unfollow/user/:userID", utility.Use(bc.unfollowUser))
		businessRouter.DELETE("/:businessAccountID/unfollow/business/:followedID", utility.Use(bc.unfollowBusiness))
		businessRouter.GET("/:businessAccountID/followed-contents", utility.Use(bc.getFollowedContents))
		businessRouter.POST("/", utility.Use(bc.createBusinessAccount))
		businessRouter.GET("/user-business-account", utility.Use(bc.getUserBusinessAccounts))
		businessRouter.GET("/:businessAccountID", utility.Use(bc.getBusinessAccount))
		businessRouter.PATCH("/:businessAccountID", utility.Use(bc.updateBusinessAccount))
		businessRouter.DELETE("/:businessAccountID", utility.Use(bc.deleteBusinessAccount))
		businessRouter.DELETE("/:businessAccountID/user/:userID", utility.Use(bc.removeUserFromBusinessAccount))
		businessRouter.PUT("/:businessAccountID/user/:currentOwnerID/:newOwnerID", utility.Use(bc.transferBusinessAccountOwnership))
		businessRouter.GET("/:businessAccountID/users", utility.Use(bc.getBusinessAccountsForUser))
		businessRouter.GET("/:businessAccountID/associated-users", utility.Use(bc.getUsersForBusinessAccount))
		businessRouter.GET("/", utility.Use(bc.listBusinessAccounts))
		businessRouter.GET("/:businessAccountID/associated", utility.Use(bc.getPlacesAndEventsAssociatedWithBusinessAccount))
		//businessRouter.POST("/:businessAccountID/place/:placeID", utility.Use(bc.associatePlaceWithBusinessAccount))
		//businessRouter.DELETE("/:businessAccountID/place/:placeID", utility.Use(bc.removePlaceFromBusinessAccount))
		//businessRouter.POST("/:businessAccountID/event/:eventID", utility.Use(bc.associateEventWithBusinessAccount))
		//businessRouter.DELETE("/:businessAccountID/event/:eventID", utility.Use(bc.removeEventFromBusinessAccount))
		businessRouter.POST("/:businessAccountID/event/", utility.Use(bc.addANewEventToBusinessAccount))
		businessRouter.POST("/:businessAccountID/team-members/:userID", utility.Use(bc.addTeamMember))
		businessRouter.GET("/:businessAccountID/team-members", utility.Use(bc.listTeamMembers))
		businessRouter.DELETE("/:businessAccountID/team-members/:userID", utility.Use(bc.removeTeamMember))
		businessRouter.PATCH("/:businessAccountID/team-members/:userID", utility.Use(bc.editTeamMember))
		businessRouter.GET("/:businessAccountID/team-members/search", utility.Use(bc.searchTeamMembers))

	}
}

// @Summary Add a team member to a Business Account
// @ID add-team-member
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param userID path string true "User ID"
// @Param data body Dto.TeamMember true "Team Member"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /business/{businessAccountID}/team-members/{userID} [post]
func (bc *BusinessAccountController) addTeamMember(c *gin.Context) error {
	//log.Println("addTeamMember")
	businessAccountID := c.Param("businessAccountID")
	userID := c.Param("userID")
	adminUser := c.MustGet("user").(string)

	// role and permissions are sent in the request body
	var teamMember TeamMember
	if err := c.ShouldBindJSON(&teamMember); err != nil {
		return err
	}

	err := bc.service.AddTeamMember(c, adminUser, userID, businessAccountID, teamMember.Role, teamMember.Permission)
	if err != nil {

		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "Team member added successfully", ""))
	return nil
}

// @Summary List all team members of a Business Account
// @ID list-team-members
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Success 200 {array} Dto.TeamMember
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /business/{businessAccountID}/team-members [get]
func (bc *BusinessAccountController) listTeamMembers(c *gin.Context) error {
	businessAccountID := c.Param("businessAccountID")

	teamMembers, err := bc.service.ListTeamMembers(c, businessAccountID)
	if err != nil {

		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(teamMembers, "success", "", ""))
	return nil
}

// @Summary Edit a team member in a Business Account
// @ID edit-team-member
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param userID path string true "User ID"
// @Param role body string true "Role"
// @Param permissions body string true "Permissions"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /business/{businessAccountID}/team-members/{userID} [patch]
func (bc *BusinessAccountController) editTeamMember(c *gin.Context) error {
	businessAccountID := c.Param("businessAccountID")
	userID := c.Param("userID")

	var teamMember TeamMember
	if err := c.ShouldBindJSON(&teamMember); err != nil {

		return err
	}

	err := bc.service.EditTeamMember(c, userID, businessAccountID, teamMember.Role, teamMember.Permission)
	if err != nil {

		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse("Team member edited successfully", "success", "", ""))
	return nil
}

// @Summary Remove a team member from a Business Account
// @ID remove-team-member
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param userID path string true "User ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /business/{businessAccountID}/team-members/{userID} [delete]
func (bc *BusinessAccountController) removeTeamMember(c *gin.Context) error {
	businessAccountID := c.Param("businessAccountID")
	userID := c.Param("userID")

	err := bc.service.RemoveTeamMember(c, userID, businessAccountID)
	if err != nil {

		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse("Team member removed successfully", "success", "", ""))
	return nil
}

// @Summary Search team members in a Business Account
// @ID search-team-members
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param searchText query string true "Search Text"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Success 200 {array} Dto.TeamMember
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /business/{businessAccountID}/team-members/search [get]
func (bc *BusinessAccountController) searchTeamMembers(c *gin.Context) error {
	businessAccountID := c.Param("businessAccountID")
	searchText := c.Query("searchText")

	teamMembers, err := bc.service.SearchTeamMembers(c, businessAccountID, searchText)
	if err != nil {

		return err
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(teamMembers, "success", "", ""))
	return nil
}

// @Summary Get Places and Events associated with a Business Account
// @ID get-places-and-events-associated-with-business-account
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param All query bool false "All"
// @Param Authorization header string true "Bearer token"
// @QueryParam page query int false "Page Number"
// @QueryParam limit query int false "Page Size"
// @QueryParam sort query string false "Sort By"
// @Accept json
// @Description Retrieve All Places and Events associated with a Business Account
// @Success 200 {object} Dto.BusinessAccountPlacesAndEvents
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /business/{businessAccountID}/associated [get]
func (bc *BusinessAccountController) getPlacesAndEventsAssociatedWithBusinessAccount(c *gin.Context) error {
	businessAccountID := c.Param("businessAccountID")
	relatedType := c.Query("relatedType")

	placesAndEvents, err := bc.service.GetPlacesAndEventsAssociatedWithBusinessAccount(c, relatedType, businessAccountID)
	if err != nil {

		return err
	}

	c.JSON(http.StatusOK, placesAndEvents)
	return nil
}

// @Summary Add a new Event to a Business Account
// @ID add-a-new-event-to-business-account
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param Dto.EventDTO body Dto.EventDTO true "Event DTO"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Add a new Event to a Business Account
// @Success 200 {object} ent.Event
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Router /business/{businessAccountID}/event [post]
func (bc *BusinessAccountController) addANewEventToBusinessAccount(c *gin.Context) error {
	businessAccountID := c.Param("businessAccountID")

	var eventDto *ent.Event
	if err := c.ShouldBindJSON(&eventDto); err != nil {

		return err
	}

	event, err := bc.eventService.CreateEvent(c, businessAccountID, eventDto)
	if err != nil {

		return err
	}

	c.JSON(http.StatusOK, event)
	return nil
}

// @Summary Follow a user by a business
// @ID follow-user
// @Produce json
// @Description Retrieve all comments for a given post
// @Tags Business
// @Accept json
// @Param businessID path string true "Business ID"
// @Param userID path string true "User ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessID}/follow/user/{userID} [post]
func (bc *BusinessAccountController) followUser(c *gin.Context) error {
	businessID := c.Param("businessID")
	userID := c.Param("userID")

	if err := bc.service.FollowUser(c, businessID, userID); err != nil {

		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully followed user"})
	return nil
}

// @Summary Follow a business by another business
// @ID follow-business
// @Tags Business
// @Produce json
// @Param followerBusinessID path string true "Follower Business ID"
// @Param followedBusinessID path string true "Followed Business ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{followerBusinessID}/follow/business/{followedID} [post]
func (bc *BusinessAccountController) followBusiness(c *gin.Context) error {
	followerID := c.Param("followerID")
	followedID := c.Param("followedID")

	if err := bc.service.FollowBusiness(c, followerID, followedID); err != nil {

		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully followed business"})
	return nil
}

// @Summary Unfollow a user by a business
// @ID unfollow-user
// @Tags Business
// @Produce json
// @Param businessID path string true "Business ID"
// @Param userID path string true "User ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessID}/unfollow/user/{userID} [delete]
func (bc *BusinessAccountController) unfollowUser(c *gin.Context) error {
	businessID := c.Param("businessID")
	userID := c.Param("userID")
	err := bc.service.UnfollowUser(c, businessID, userID)
	if err != nil {

		return err
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully unfollowed the user"})
	return nil
}

// @Summary Unfollow a business by another business
// @ID unfollow-business
// @Tags Business
// @Produce json
// @Param followerBusinessID path string true "Follower Business ID"
// @Param followedBusinessID path string true "Followed Business ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{followerBusinessID}/unfollow/business/{followedBusinessID} [delete]
func (bc *BusinessAccountController) unfollowBusiness(c *gin.Context) error {
	followerID := c.Param("followerID")
	followedID := c.Param("followedID")
	err := bc.service.UnfollowBusiness(c, followerID, followedID)
	if err != nil {

		return err
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully unfollowed the business"})
	return nil
}

// @Summary Get contents followed by a business
// @ID get-followed-contents
// @Tags Business
// @Produce json
// @Param businessID path string true "Business ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessID}/followed-contents [get]
func (bc *BusinessAccountController) getFollowedContents(c *gin.Context) error {
	businessID := c.Param("businessID")
	posts, err := bc.service.GetFollowedContents(c, businessID)
	if err != nil {

		return err
	}
	c.JSON(http.StatusOK, posts)
	return nil
}

// @Summary Create a business account
// @ID create-business-account
// @Tags Business
// @Accept  json
// @Produce  json
// @Param Dto.BusinessDto body Dto.BusinessDto true "Business Account Data"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/ [post]
func (bc *BusinessAccountController) createBusinessAccount(c *gin.Context) error {
	var businessData BusinessDto
	if err := c.ShouldBindJSON(&businessData); err != nil {
		log.Println("error", err.Error())
		return err
	}
	business, err := bc.service.CreateBusinessAccount(c, &businessData)
	if err != nil {

		return err
	}
	c.JSON(http.StatusCreated, business)
	return nil
}

// @Summary Get a business account
// @ID get-business-account
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID} [get]
func (bc *BusinessAccountController) getBusinessAccount(c *gin.Context) error {
	businessAccountID := c.Param("businessAccountID")
	business, err := bc.service.GetBusinessAccount(c, businessAccountID)
	if err != nil {

		return err
	}
	c.JSON(http.StatusOK, business)
	return nil
}

// updateBusinessAccount updates a business's details.
// @Summary Update a business's details
// @Description Get a business's details by ID
// @Tags Business
// @Accept json
// @Produce json
// @Param id path string true "Business ID"
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} ent.Business "Successfully retrieved business"
// @Failure 400 {object} Dto.ErrorDTO "Bad Request"
// @Failure 401 {object} Dto.ErrorDTO "Unauthorized"
// @Failure 500 {object} Dto.ErrorDTO "Internal Server Error"
// @Router /api/v1/business/ [patch]
func (bc *BusinessAccountController) updateBusinessAccount(ctx *gin.Context) error {
	businessId := ctx.Param("businessAccountID")
	if businessId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Business ID required",
		})
		return errors.New("business ID required")
	}

	var business map[string]interface{}
	if err := ctx.ShouldBindJSON(&business); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return err
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
		return err
	}

	ctx.JSON(http.StatusOK, businessData)
	return nil
}

// @Summary Delete a business account
// @ID delete-business-account
// @Tags Business
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID} [delete]
func (bc *BusinessAccountController) deleteBusinessAccount(c *gin.Context) error {
	//businessAccountID := c.Param("businessAccountID")
	//err := bc.service.DeleteBusinessAccount(c, businessAccountID)
	//if err != nil {
	//
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted the business account"})
	return nil
}

// @Summary Get user business accounts
// @ID get-user-business-accounts
// @Tags Business
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/user-business-account [get]
func (bc *BusinessAccountController) getUserBusinessAccounts(c *gin.Context) error {
	log.Println("Get user business accounts")
	// get user id from context
	userID := c.GetString("user")

	userBusinessAccounts, err := bc.service.GetUserBusinessAccounts(c, userID)
	if err != nil {
		log.Println("Error getting user business accounts", err)
		return err
	}

	c.JSON(http.StatusOK, userBusinessAccounts)
	return nil
}

// @Summary Remove user from business account
// @ID remove-user-business-account
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Tags Business
// @Param businessAccountID path string true "Business Account ID"
// @Param userID path string true "User ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID}/user/{userID} [delete]
func (bc *BusinessAccountController) removeUserFromBusinessAccount(c *gin.Context) error {
	// Implementation...
	return nil
}

// @Summary Transfer business account ownership
// @ID transfer-business-account-ownership
// @Produce json
// @Tags Business
// @Param Authorization header string true "Bearer token"
// @Param businessAccountID path string true "Business Account ID"
// @Param currentOwnerID path string true "Current Owner ID"
// @Param newOwnerID path string true "New Owner ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID}/user/{currentOwnerID}/{newOwnerID} [put]
func (bc *BusinessAccountController) transferBusinessAccountOwnership(c *gin.Context) error {
	// Implementation...
	return nil
}

// @Summary Get business accounts for a user
// @ID get-business-accounts-for-user
// @Produce json
// @Tags Business
// @Param Authorization header string true "Bearer token"
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID}/users [get]
func (bc *BusinessAccountController) getBusinessAccountsForUser(c *gin.Context) error {
	// Implementation...
	return nil
}

// @Summary Get users for a business account
// @ID get-users-business-account
// @Produce json
// @Tags Business
// @Param Authorization header string true "Bearer token"
// @Param businessAccountID path string true "Business Account ID"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID}/users [get]
func (bc *BusinessAccountController) getUsersForBusinessAccount(c *gin.Context) error {
	// Implementation...
	return nil
}

// @Summary List all business accounts
// @ID list-business-accounts
// @Produce json
// @Tags Business
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/ [get]
func (bc *BusinessAccountController) listBusinessAccounts(c *gin.Context) error {
	// Implementation...
	return nil
}
