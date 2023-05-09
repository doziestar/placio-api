package controller

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	_ "gorm.io/gorm"
	"net/http"
	"placio-app/middleware"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"
	"placio-pkg/logger"
)

//ref
//https://blog.cloudnativefolks.org/oauth-20-implementation-in-golang
//
// https://github.com/gofiber/jwt
// https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j
// https://github.com/markbates/goth/tree/master/providers/twitterv2
// https://github.com/gofiber/fiber/issues/292

//users := []models.User{}

type UserController struct {
	service service.IUser
}

func NewUserController(service service.IUser) *UserController {
	return &UserController{service: service}
} // GET /api/users/:id/messages_sent - Retrieve a list of messages sent by a specific user by ID

func (c *UserController) RegisterRoutes(app *gin.RouterGroup) {
	userGroup := app.Group("/users")

	userGroup.POST("/", utility.Use(c.CreateUser))
	userGroup.GET("/", middleware.Verify("user"), utility.Use(c.getAllUsers))
	userGroup.GET("/me", middleware.Verify("user"), utility.Use(c.GetMe))
	userGroup.GET("/exists", utility.Use(c.checkIfUserExist))
	userGroup.GET("/:id", middleware.Verify("user"), utility.Use(c.getUserByID))
	userGroup.PUT("/:id", middleware.Verify("user"), utility.Use(c.UpdateUser))
	userGroup.DELETE("/:id", middleware.Verify("user"), utility.Use(c.deleteUser))
	userGroup.GET("/:id/sent_messages", middleware.Verify("user"), utility.Use(c.GetMessagesSent))
	userGroup.GET("/:id/received_messages", middleware.Verify("user"), utility.Use(c.GetMessagesReceived))
	userGroup.GET("/:id/conversations", middleware.Verify("user"), utility.Use(c.GetConversations))
	userGroup.GET("/:id/groups", middleware.Verify("user"), utility.Use(c.GetGroups))
	userGroup.GET("/:id/sent_voice_notes", middleware.Verify("user"), utility.Use(c.GetVoiceNotesSent))
	userGroup.GET("/:id/received_voice_notes", middleware.Verify("user"), utility.Use(c.GetVoiceNotesReceived))
	userGroup.GET("/:id/notifications", middleware.Verify("user"), utility.Use(c.GetUserNotifications))
	userGroup.GET("/:id/bookings", middleware.Verify("user"), utility.Use(c.GetUserBookings))
	userGroup.GET("/:id/payments", middleware.Verify("user"), utility.Use(c.GetUserPayments))
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Accept */*
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /api/v1/users [post]
func (c *UserController) CreateUser(ctx *gin.Context) error {
	user := new(models.User)
	if err := ctx.ShouldBindJSON(user); err != nil {
		logger.Error(ctx, fmt.Sprintf("error while binding user data: %v", err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	//user.ID = len(users) + 1
	//users = append(users, *user)
	ctx.JSON(http.StatusOK, user)
	return nil
}

// getAllUsers godoc
// @Summary Retrieve a list of users
// @Description Retrieve a list of users
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} []models.User
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users [get]
func (c *UserController) getAllUsers(ctx *gin.Context) error {
	//return c.JSON(users)
	ctx.JSON(http.StatusOK, "users")
	return nil
}

// getUserByID godoc
// @Summary Retrieve a user by ID
// @Description Retrieve a user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id} [get]
func (c *UserController) getUserByID(ctx *gin.Context) error {
	// get user id from the path
	id := ctx.Param("id")
	// get user from the user service
	user, err := c.service.GetUserByID(id)
	if err != nil {
		sentry.CaptureException(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	items, err := utility.RemoveSensitiveFields(user)
	if err != nil {
		sentry.CaptureException(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": items})
	return nil
}

// checkIfUserExist godoc
// @Summary Check if user exists
// @Description Check if user exists
// @Tags User
// @Accept */*
// @Produce json
// @Param username query string false "Username"
// @Param email query string false "Email"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/check_user [get]c
func (c *UserController) checkIfUserExist(ctx *gin.Context) error {
	username := ctx.Query("username")
	email := ctx.Query("email")

	user, err := c.service.CheckIfUserNameOrEmailExists(username, email)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusOK, gin.H{"status": "ok", "exist": false})
			return nil
		}
		sentry.CaptureException(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "exist": true, "user": user})
	return nil
}

// UpdateUser godoc
// @Summary Update an existing user by ID
// @Description Update an existing user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id} [put]
func (c *UserController) UpdateUser(ctx *gin.Context) error {
	//id := c.Params("id")
	//for i, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		if err := c.BodyParser(&users[i]); err != nil {
	//			return err
	//		}
	//		return c.JSON(user)
	//	}
	//}
	ctx.JSON(http.StatusOK, "user")
	return nil
}

// deleteUser godoc
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id} [delete]
func (c *UserController) deleteUser(ctx *gin.Context) error {
	//id := c.Params("id")
	//for i, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		users = append(users[:i], users[i+1:]...)
	//		return c.SendStatus(fiber.StatusNoContent)
	//	}
	//}
	ctx.JSON(http.StatusOK, "user")
	return nil
}

// GetMessagesSent godoc
// @Summary Retrieve a list of messages sent by a specific user by ID
// @Description Retrieve a list of messages sent by a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Message
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/messages_sent [get]
func (c *UserController) GetMessagesSent(ctx *gin.Context) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.MessagesSent)
	//	}
	//}
	ctx.JSON(http.StatusOK, "user")
	return nil
}

// GetMessagesReceived godoc
// @Summary Retrieve a list of messages received by a specific user by ID
// @Description Retrieve a list of messages received by a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Message
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/messages_received [get]
func (controller *UserController) GetMessagesReceived(c *gin.Context) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.MessagesReceived)
	//	}
	//}
	c.JSON(http.StatusOK, "user")
	return nil
}

// GetConversations godoc
// @Summary Retrieve a list of conversations a specific user is a participant in by ID
// @Description Retrieve a list of conversations a specific user is a participant in by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Conversation
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/conversations [get]
func (controller *UserController) GetConversations(c *gin.Context) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Conversations)
	//	}
	//}

	c.JSON(http.StatusOK, "user")
	return nil
}

// GetGroups godoc
// @Summary Retrieve a list of groups a specific user is a member of by ID
// @Description Retrieve a list of groups a specific user is a member of by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Group
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/groups [get]
func (controller *UserController) GetGroups(c *gin.Context) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Groups)
	//	}
	//}
	c.JSON(http.StatusOK, "user")
	return nil
}

// GetVoiceNotesSent godoc
// @Summary Retrieve a list of voice notes sent by a specific user by ID
// @Description Retrieve a list of voice notes sent by a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.VoiceNote
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/voice_notes_sent [get]
func (controller *UserController) GetVoiceNotesSent(c *gin.Context) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.VoiceNotesSent)
	//	}
	//}
	c.JSON(http.StatusOK, "user")
	return nil
}

// GetVoiceNotesReceived godoc
// @Summary Retrieve a list of voice notes received by a specific user by ID
// @Description Retrieve a list of voice notes received by a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.VoiceNote
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/voice_notes_received [get]
func (controller *UserController) GetVoiceNotesReceived(c *gin.Context) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.VoiceNotesReceived)
	//	}
	//}
	c.JSON(http.StatusOK, "user")
	return nil
}

// GetVoiceNotes godoc
// @Summary Retrieve a list of voice notes a specific user is a participant in by ID
// @Description Retrieve a list of voice notes a specific user is a participant in by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.VoiceNote
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/voice_notes [get]
func (controller *UserController) GetVoiceNotes(c *gin.Context) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.VoiceNotes)
	//	}
	//}
	c.JSON(http.StatusOK, "user")
	return nil
}

// GetVoiceNote godoc
// @Summary Retrieve a specific voice note by ID
// @Description Retrieve a specific voice note by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "Voice Note ID"
// @Success 200 {object} models.VoiceNote
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/voice_notes/{id} [get]
func (controller *UserController) GetVoiceNote(c *gin.Context) error {
	//id := c.Params("id")
	//for _, voiceNote := range voiceNotes {
	//	if strconv.Itoa(voiceNote.ID) == id {
	//		return c.JSON(voiceNote)
	//	}
	//}
	c.JSON(http.StatusOK, "user")
	return nil
}

// GetUserNotifications godoc
// @Summary Retrieve a list of notifications for a specific user by ID
// @Description Retrieve a list of notifications for a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Notification
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/notifications [get]
func (controller *UserController) GetUserNotifications(c *gin.Context) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Notifications)
	//	}
	//}
	c.JSON(http.StatusOK, "user")
	return nil
}

// GetUserBookings godoc
// @Summary Retrieve a list of bookings for a specific user by ID
// @Description Retrieve a list of bookings for a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Booking
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/bookings [get]
func (controller *UserController) GetUserBookings(c *gin.Context) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Bookings)
	//	}
	//}
	c.JSON(http.StatusOK, "user")
	return nil
}

// GetUserPayments godoc
// @Summary Retrieve a list of payments for a specific user by ID
// @Description Retrieve a list of payments for a specific user by ID
// @Tags User
// @Accept */*
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []models.Payment
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/{id}/payments [get]
func (controller *UserController) GetUserPayments(c *gin.Context) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Payments)
	//	}
	//}
	c.JSON(http.StatusOK, "user")
	return nil
}

// GetMe godoc
// @Summary Retrieve the current user
// @Description Retrieve the current user
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/users/me [get]
func (c *UserController) GetMe(ctx *gin.Context) error {
	userIdValue, exist := ctx.Get("user")
	if !exist {
		sentry.CaptureMessage("User not found")
		ctx.JSON(http.StatusInternalServerError, fiber.Map{"message": "User not found"})
	}
	logger.Info(context.Background(), fmt.Sprintf("GetMe: %s", userIdValue))
	userId, ok := userIdValue.(string)
	if !ok {
		sentry.CaptureMessage("User ID is not a string")
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "User ID has an invalid type"})
		return nil
	}
	userData, err := c.service.GetLoggedInUser(userId)
	if err != nil {
		sentry.CaptureException(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "User not found"})
		return nil
	}

	userInfo, err := utility.RemoveSensitiveFields(userData)
	if err != nil {
		sentry.CaptureException(err)
	}
	ctx.JSON(http.StatusOK, userInfo)
	return nil
}
