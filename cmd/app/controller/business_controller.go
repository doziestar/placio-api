package controller

import (
	"errors"
	"log"
	"net/http"
	"placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"

	"github.com/gin-gonic/gin"
)

type BusinessAccountController struct {
	service      service.BusinessAccountService
	eventService service.EventService
}

func NewBusinessAccountController(service service.BusinessAccountService) *BusinessAccountController {
	return &BusinessAccountController{service: service}
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
		businessRouter.POST("/:businessAccountID/user/:userID", utility.Use(bc.associateUserWithBusinessAccount))
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
	}
}

// @Summary Get Places and Events associated with a Business Account
// @ID get-places-and-events-associated-with-business-account
// @Tags Business
// @Produce json
// @Param businessAccountID path string true "Business Account ID"
// @Param All query bool false "All"
// @Security Bearer
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

	placesAndEvents, err := bc.service.GetPlacesAndEventsAssociatedWithBusinessAccount(c, businessAccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
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
// @Security Bearer
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Add a new Event to a Business Account
// @Success 200 {object} ent.Event
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Router /business/{businessAccountID}/event [post]
func (bc *BusinessAccountController) addANewEventToBusinessAccount(c *gin.Context) error {
	businessAccountID := c.Param("businessAccountID")

	var eventDto Dto.EventDTO
	if err := c.ShouldBindJSON(&eventDto); err != nil {
		c.JSON(http.StatusBadRequest, utility.ProcessResponse(nil, "failed", err.Error()))
		return err
	}

	event, err := bc.eventService.CreateEvent(c, businessAccountID, eventDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.ProcessResponse(nil, "failed", err.Error()))
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
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessID}/follow/user/{userID} [post]
func (bc *BusinessAccountController) followUser(c *gin.Context) error {
	businessID := c.Param("businessID")
	userID := c.Param("userID")

	if err := bc.service.FollowUser(c, businessID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{followerBusinessID}/follow/business/{followedID} [post]
func (bc *BusinessAccountController) followBusiness(c *gin.Context) error {
	followerID := c.Param("followerID")
	followedID := c.Param("followedID")

	if err := bc.service.FollowBusiness(c, followerID, followedID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessID}/followed-contents [get]
func (bc *BusinessAccountController) getFollowedContents(c *gin.Context) error {
	businessID := c.Param("businessID")
	posts, err := bc.service.GetFollowedContents(c, businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/ [post]
func (bc *BusinessAccountController) createBusinessAccount(c *gin.Context) error {
	var businessData Dto.BusinessDto
	if err := c.ShouldBindJSON(&businessData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	business, err := bc.service.CreateBusinessAccount(c, &businessData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Success 200 {object} Dto.Response
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.ErrorDto
// @Router /business/{businessAccountID} [get]
func (bc *BusinessAccountController) getBusinessAccount(c *gin.Context) error {
	businessAccountID := c.Param("businessAccountID")
	business, err := bc.service.GetBusinessAccount(c, businessAccountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Security Bearer
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
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted the business account"})
	return nil
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
func (bc *BusinessAccountController) getUserBusinessAccounts(c *gin.Context) error {
	log.Println("Get user business accounts")
	businessAccount, err := bc.service.GetUserBusinessAccounts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	c.JSON(http.StatusOK, gin.H{"businessAccounts": businessAccount})
	return nil
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
func (bc *BusinessAccountController) associateUserWithBusinessAccount(c *gin.Context) error {
	// Implementation...
	return nil
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
func (bc *BusinessAccountController) removeUserFromBusinessAccount(c *gin.Context) error {
	// Implementation...
	return nil
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
func (bc *BusinessAccountController) transferBusinessAccountOwnership(c *gin.Context) error {
	// Implementation...
	return nil
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
func (bc *BusinessAccountController) getBusinessAccountsForUser(c *gin.Context) error {
	// Implementation...
	return nil
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
func (bc *BusinessAccountController) getUsersForBusinessAccount(c *gin.Context) error {
	// Implementation...
	return nil
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
func (bc *BusinessAccountController) listBusinessAccounts(c *gin.Context) error {
	// Implementation...
	return nil
}
