package api

import (
	"github.com/gofiber/fiber/v2"
	"placio-app/controller"
)

func UserRoutes(app *fiber.App) {
	apiRouter := app.Group("/api")
	v1 := apiRouter.Group("/v1/users")
	{
		// POST /api/users - Create a new user
		v1.Post("/", controller.CreateUser)

		// GET /api/users/:id - Retrieve a specific user by ID
		v1.Get("/:id", controller.GetUserByID)

		// PUT /api/users/:id - Update an existing user by ID
		v1.Put("/:id", controller.UpdateUser)

		// DELETE /api/users/:id - Delete a user by ID
		v1.Delete("/:id", controller.DeleteUser)

		// GET /api/users/:id/messages_sent - Retrieve a list of messages sent by a specific user by ID
		v1.Get("/:id/messages_sent", controller.GetMessagesSent)

		// GET /api/users/:id/messages_received - Retrieve a list of messages received by a specific user by ID
		v1.Get("/:id/messages_received", controller.GetMessagesReceived)

		// GET /api/users/:id/conversations - Retrieve a list of conversations a specific user is a participant in by ID
		v1.Get("/:id/conversations", controller.GetConversations)

		// GET /api/users/:id/groups - Retrieve a list of groups a specific user is a member of by ID
		v1.Get("/:id/groups", controller.GetGroups)

		// GET /api/users/:id/voice_notes_sent - Retrieve a list of voice notes sent by a specific user by ID
		v1.Get("/:id/voice_notes_sent", controller.GetVoiceNotesSent)

		// GET /api/users/:id/voice_notes_received - Retrieve a list of voice notes received by a specific user by ID
		v1.Get("/:id/voice_notes_received", controller.GetVoiceNotesReceived)

		// GET /api/users/:id/notifications - Retrieve a list of notifications for a specific user by ID
		v1.Get("/:id/notifications", controller.GetUserNotifications)

		// GET /api/users/:id/bookings - Retrieve a list of bookings made by a specific user by ID
		v1.Get("/:id/bookings", controller.GetUserBookings)

		// GET /api/users/:id/payments - Retrieve a list of payments made by a specific user by ID
		v1.Get("/:id/payments", controller.GetUserPayments)
	}
}
