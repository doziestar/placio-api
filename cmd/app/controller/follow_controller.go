package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "placio-app/Dto"
	_ "placio-app/ent"
	"placio-app/service"
	"placio-app/utility"
)

type FollowController struct {
	service service.IFollowService
}

func NewFollowController(service service.IFollowService) *FollowController {
	return &FollowController{service: service}
}

func (fc *FollowController) RegisterRoutes(router *gin.RouterGroup) {
	followRouter := router.Group("/follow")
	{
		followRouter.POST("business/:businessID", utility.Use(fc.followUserToBusiness))
		followRouter.GET("business", utility.Use(fc.getFollowedBusinessesByUser))

		followRouter.POST("user/:userID", utility.Use(fc.followUserToUser))
		followRouter.GET("user", utility.Use(fc.getFollowedUsersByUser))

		followRouter.POST("place/:placeID", utility.Use(fc.followUserToPlace))
		followRouter.GET("place", utility.Use(fc.getFollowedPlacesByUser))

		followRouter.POST("event/:eventID", utility.Use(fc.followUserToEvent))
		followRouter.GET("event", utility.Use(fc.getFollowedEventsByUser))

		followRouter.GET("check/business/:businessID", utility.Use(fc.checkUserFollowsBusiness))
		followRouter.GET("check/user/:userID", utility.Use(fc.checkUserFollowsUser))
		followRouter.GET("check/place/:placeID", utility.Use(fc.checkUserFollowsPlace))
		followRouter.GET("check/event/:eventID", utility.Use(fc.checkUserFollowsEvent))

	}

	unfollowRouter := router.Group("/unfollow")
	{
		unfollowRouter.DELETE("business/:businessID", utility.Use(fc.unfollowUserToBusiness))
		unfollowRouter.DELETE("user/:userID", utility.Use(fc.unfollowUserToUser))
		unfollowRouter.DELETE("place/:placeID", utility.Use(fc.unfollowUserToPlace))
		unfollowRouter.DELETE("event/:eventID", utility.Use(fc.unfollowUserToEvent))
	}

}

// @Summary Check if user follows a business
// @ID user-check-follow-business
// @Tags Follow
// @Produce json
// @Param businessID path string true "Business ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Check if the user follows a specific Business
// @Success 200 {object} string "Follow status returned successfully"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/check/business/{businessID} [get]
func (fc *FollowController) checkUserFollowsBusiness(c *gin.Context) error {
	businessID := c.Param("businessID")
	userID := c.GetString("user")

	follows, err := fc.service.CheckIfUserFollowsBusiness(c, userID, businessID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"follows": follows})
	return nil
}

// @Summary Check if user follows another user
// @ID user-check-follow-user
// @Tags Follow
// @Produce json
// @Param followedID path string true "User ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Check if the user follows another User
// @Success 200 {object} string "Follow status returned successfully"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/check/user/{followedID} [get]
func (fc *FollowController) checkUserFollowsUser(c *gin.Context) error {
	followedID := c.Param("followedID")
	followerID := c.GetString("user")

	follows, err := fc.service.CheckIfUserFollowsUser(c, followerID, followedID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"follows": follows})
	return nil
}

// @Summary Check if user follows a place
// @ID user-check-follow-place
// @Tags Follow
// @Produce json
// @Param placeID path string true "Place ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Check if the user follows a Place
// @Success 200 {object} string "Follow status returned successfully"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/check/place/{placeID} [get]
func (fc *FollowController) checkUserFollowsPlace(c *gin.Context) error {
	placeID := c.Param("placeID")
	userID := c.GetString("user")

	follows, err := fc.service.CheckIfUserFollowsPlace(c, userID, placeID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"follows": follows})
	return nil
}

// @Summary Check if user follows an event
// @ID user-check-follow-event
// @Tags Follow
// @Produce json
// @Param eventID path string true "Event ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Check if the user follows an Event
// @Success 200 {object} string "Follow status returned successfully"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/check/event/{eventID} [get]
func (fc *FollowController) checkUserFollowsEvent(c *gin.Context) error {
	eventID := c.Param("eventID")
	userID := c.GetString("user")

	follows, err := fc.service.CheckIfUserFollowsEvent(c, userID, eventID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"follows": follows})
	return nil
}

// @Summary Follow a business
// @ID user-follow-business
// @Tags Follow
// @Produce json
// @Param businessID path string true "Business ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Follow a specific Business by the user
// @Success 201 {object} string "User successfully followed the business"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/business/{businessID} [post]
func (fc *FollowController) followUserToBusiness(c *gin.Context) error {
	businessID := c.Param("businessID")
	userID := c.GetString("user")

	err := fc.service.FollowUserToBusiness(c, userID, businessID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusCreated, utility.ProcessResponse(nil, "success", "User successfully followed the business", ""))
	return nil
}

// @Summary Unfollow a business
// @ID user-unfollow-business
// @Tags Follow
// @Produce json
// @Param businessID path string true "Business ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Unfollow a specific Business by the user
// @Success 200 {object} string "User successfully unfollowed the business"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /unfollow/business/{businessID} [delete]
func (fc *FollowController) unfollowUserToBusiness(c *gin.Context) error {
	businessID := c.Param("businessID")
	userID := c.GetString("user")

	err := fc.service.UnfollowUserToBusiness(c, userID, businessID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "User successfully unfollowed the business", ""))
	return nil
}

// @Summary Get followed businesses by a user
// @ID get-followed-businesses
// @Tags Follow
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Get all Businesses followed by the user
// @Success 200 {array} ent.Business
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /followed/businesses [get]
func (fc *FollowController) getFollowedBusinessesByUser(c *gin.Context) error {
	userID := c.GetString("user")

	businesses, err := fc.service.GetFollowedBusinessesByUser(c, userID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(businesses, "success", "Successfully retrieved followed businesses", ""))
	return nil
}

// @Summary Follow a user
// @ID user-follow-user
// @Tags Follow
// @Produce json
// @Param userID path string true "User ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Follow a specific User
// @Success 201 {object} string "User successfully followed the user"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/user/{userID} [post]
func (fc *FollowController) followUserToUser(c *gin.Context) error {
	followedID := c.Param("userID")
	followerID := c.GetString("user")

	err := fc.service.FollowUserToUser(c, followerID, followedID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusCreated, utility.ProcessResponse(nil, "success", "User successfully followed the user", ""))
	return nil
}

// @Summary Unfollow a user
// @ID user-unfollow-user
// @Tags Follow
// @Produce json
// @Param userID path string true "User ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Unfollow a specific User
// @Success 200 {object} string "User successfully unfollowed the user"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /unfollow/user/{userID} [delete]
func (fc *FollowController) unfollowUserToUser(c *gin.Context) error {
	followedID := c.Param("userID")
	followerID := c.GetString("user")

	err := fc.service.UnfollowUserToUser(c, followerID, followedID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "User successfully unfollowed the user", ""))
	return nil
}

// @Summary Get followed users by a user
// @ID get-followed-users
// @Tags Follow
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Get all Users followed by the user
// @Success 200 {array} ent.User
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /followed/users [get]
func (fc *FollowController) getFollowedUsersByUser(c *gin.Context) error {
	userID := c.GetString("userID")

	users, err := fc.service.GetFollowedUsersByUser(c, userID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(users, "success", "Successfully retrieved followed users", ""))
	return nil
}

// @Summary Follow a place
// @ID follow-place
// @Tags Follow
// @Produce json
// @Param placeID path string true "Place ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Follow a specific Place by the user
// @Success 201 {object} string "User successfully followed the place"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/place/{placeID} [post]
func (fc *FollowController) followUserToPlace(c *gin.Context) error {
	placeID := c.Param("placeID")
	userID := c.GetString("user")

	err := fc.service.FollowUserToPlace(c, userID, placeID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusCreated, utility.ProcessResponse(nil, "success", "User successfully followed the place", ""))
	return nil
}

// @Summary Unfollow a place
// @ID unfollow-place
// @Tags Follow
// @Produce json
// @Param placeID path string true "Place ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Unfollow a specific Place by the user
// @Success 200 {object} string "User successfully unfollowed the place"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /unfollow/place/{placeID} [delete]
func (fc *FollowController) unfollowUserToPlace(c *gin.Context) error {
	placeID := c.Param("placeID")
	userID := c.GetString("user")

	err := fc.service.UnfollowUserToPlace(c, userID, placeID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "User successfully unfollowed the place", ""))
	return nil
}

// @Summary Get followed places by a user
// @ID get-followed-places
// @Tags Follow
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Get all Places followed by the user
// @Success 200 {array} ent.Place
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /followed/places [get]
func (fc *FollowController) getFollowedPlacesByUser(c *gin.Context) error {
	userID := c.GetString("user")

	places, err := fc.service.GetFollowedPlacesByUser(c, userID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(places, "success", "Successfully retrieved followed places", ""))
	return nil
}

// @Summary Follow an event
// @ID follow-event
// @Tags Follow
// @Produce json
// @Param eventID path string true "Event ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Follow a specific Event by the user
// @Success 201 {object} string "User successfully followed the event"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /follow/event/{eventID} [post]
func (fc *FollowController) followUserToEvent(c *gin.Context) error {
	eventID := c.Param("eventID")
	userID := c.GetString("user")

	err := fc.service.FollowUserToEvent(c, userID, eventID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusCreated, utility.ProcessResponse(nil, "success", "User successfully followed the event", ""))
	return nil
}

// @Summary Unfollow an event
// @ID unfollow-event
// @Tags Follow
// @Produce json
// @Param eventID path string true "Event ID"
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Unfollow a specific Event by the user
// @Success 200 {object} string "User successfully unfollowed the event"
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /unfollow/event/{eventID} [delete]
func (fc *FollowController) unfollowUserToEvent(c *gin.Context) error {
	eventID := c.Param("eventID")
	userID := c.GetString("user")

	err := fc.service.UnfollowUserToEvent(c, userID, eventID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(nil, "success", "User successfully unfollowed the event", ""))
	return nil
}

// @Summary Get followed events by a user
// @ID get-followed-events
// @Tags Follow
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Accept json
// @Description Get all Events followed by the user
// @Success 200 {array} ent.Event
// @Failure 400 {object} Dto.Error
// @Failure 401 {object} Dto.Error
// @Failure 500 {object} Dto.Error
// @Router /followed/events [get]
func (fc *FollowController) getFollowedEventsByUser(c *gin.Context) error {
	userID := c.GetString("user")

	events, err := fc.service.GetFollowedEventsByUser(c, userID)
	if err != nil {

		return nil
	}

	c.JSON(http.StatusOK, utility.ProcessResponse(events, "success", "Successfully retrieved followed events", ""))
	return nil
}
