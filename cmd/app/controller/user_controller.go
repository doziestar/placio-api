package controller

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	_ "gorm.io/gorm"
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

func (c *UserController) RegisterRoutes(app fiber.Router) {
	userGroup := app.Group("/users")

	userGroup.Post("/", utility.Use(c.CreateUser))
	userGroup.Get("/", middleware.Verify("user"), utility.Use(c.getAllUsers))
	userGroup.Get("/me", middleware.Verify("user"), utility.Use(c.GetMe))
	userGroup.Get("/check_user", utility.Use(c.checkIfUserExist))
	userGroup.Get("/:userId", middleware.Verify("user"), utility.Use(c.getUserByID))
	userGroup.Put("/:id", middleware.Verify("user"), utility.Use(c.UpdateUser))
	userGroup.Delete("/:id", middleware.Verify("user"), utility.Use(c.deleteUser))
	userGroup.Get("/:id/messages_sent", middleware.Verify("user"), utility.Use(c.GetMessagesSent))
	userGroup.Get("/:id/messages_received", middleware.Verify("user"), utility.Use(c.GetMessagesReceived))
	userGroup.Get("/:id/conversations", middleware.Verify("user"), utility.Use(c.GetConversations))
	userGroup.Get("/:id/groups", middleware.Verify("user"), utility.Use(c.GetGroups))
	userGroup.Get("/:id/voice_notes_sent", middleware.Verify("user"), utility.Use(c.GetVoiceNotesSent))
	userGroup.Get("/:id/voice_notes_received", middleware.Verify("user"), utility.Use(c.GetVoiceNotesReceived))
	userGroup.Get("/:id/notifications", middleware.Verify("user"), utility.Use(c.GetUserNotifications))
	userGroup.Get("/:id/bookings", middleware.Verify("user"), utility.Use(c.GetUserBookings))
	userGroup.Get("/:id/payments", middleware.Verify("user"), utility.Use(c.GetUserPayments))

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
func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}
	//user.ID = len(users) + 1
	//users = append(users, *user)
	return ctx.JSON(user)
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
func (c *UserController) getAllUsers(ctx *fiber.Ctx) error {
	//return c.JSON(users)
	return ctx.JSON(fiber.Map{"status": "ok"})
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
func (c *UserController) getUserByID(ctx *fiber.Ctx) error {
	// get user id from the path
	id := ctx.Params("userId")
	// get user from the user service
	user, err := c.service.GetUserByID(id)
	if err != nil {
		sentry.CaptureException(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	items, err := utility.RemoveSensitiveFields(user)
	if err != nil {
		sentry.CaptureException(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"data":   items,
	})
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
func (c *UserController) checkIfUserExist(ctx *fiber.Ctx) error {
	username := ctx.Query("username")
	email := ctx.Query("email")

	user, err := c.service.CheckIfUserNameOrEmailExists(username, email)
	if err != nil {
		if err.Error() == "user not found" {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "exist": false})
		}
		sentry.CaptureException(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"exist":  user,
	})
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
func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	//id := c.Params("id")
	//for i, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		if err := c.BodyParser(&users[i]); err != nil {
	//			return err
	//		}
	//		return c.JSON(user)
	//	}
	//}
	return ctx.Status(fiber.StatusNotFound).SendString("User not found")
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
func (c *UserController) deleteUser(ctx *fiber.Ctx) error {
	//id := c.Params("id")
	//for i, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		users = append(users[:i], users[i+1:]...)
	//		return c.SendStatus(fiber.StatusNoContent)
	//	}
	//}
	return ctx.Status(fiber.StatusNotFound).SendString("User not found")
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
func (c *UserController) GetMessagesSent(ctx *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.MessagesSent)
	//	}
	//}
	return ctx.Status(fiber.StatusNotFound).SendString("User not found")
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
func (controller *UserController) GetMessagesReceived(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.MessagesReceived)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
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
func (controller *UserController) GetConversations(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Conversations)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
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
func (controller *UserController) GetGroups(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Groups)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
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
func (controller *UserController) GetVoiceNotesSent(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.VoiceNotesSent)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
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
func (controller *UserController) GetVoiceNotesReceived(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.VoiceNotesReceived)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
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
func (controller *UserController) GetVoiceNotes(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.VoiceNotes)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
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
func (controller *UserController) GetVoiceNote(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, voiceNote := range voiceNotes {
	//	if strconv.Itoa(voiceNote.ID) == id {
	//		return c.JSON(voiceNote)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("Voice note not found")
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
func (controller *UserController) GetUserNotifications(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Notifications)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
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
func (controller *UserController) GetUserBookings(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Bookings)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
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
func (controller *UserController) GetUserPayments(c *fiber.Ctx) error {
	//id := c.Params("id")
	//for _, user := range users {
	//	if strconv.Itoa(user.ID) == id {
	//		return c.JSON(user.Payments)
	//	}
	//}
	return c.Status(fiber.StatusNotFound).SendString("User not found")
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
func (c *UserController) GetMe(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user").(string)
	logger.Info(context.Background(), fmt.Sprintf("GetMe: %s", userId))
	userData, err := c.service.GetLoggedInUser(userId)
	if err != nil {
		sentry.CaptureException(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	userInfo, err := utility.RemoveSensitiveFields(userData)
	if err != nil {
		sentry.CaptureException(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(userInfo)
}
