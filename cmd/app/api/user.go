package api

import (
	"placio-app/controller"
	"placio-app/utility"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	apiRouter := app.Group("/api")
	v1 := apiRouter.Group("/v1/users")
	{
		// POST /api/users - Create a new user
		v1.Post("/", utility.Use(controller.CreateUser))

		// GET /api/users/:id - Retrieve a specific user by ID
		v1.Get("/:id", utility.Use(controller.GetUserByID))

		// PUT /api/users/:id - Update an existing user by ID
		v1.Put("/:id", utility.Use(controller.UpdateUser))

		// DELETE /api/users/:id - Delete a user by ID
		v1.Delete("/:id", utility.Use(controller.DeleteUser))

		// GET /api/users/:id/messages_sent - Retrieve a list of messages sent by a specific user by ID
		v1.Get("/:id/messages_sent", utility.Use(controller.GetMessagesSent))

		// GET /api/users/:id/messages_received - Retrieve a list of messages received by a specific user by ID
		v1.Get("/:id/messages_received", utility.Use(controller.GetMessagesReceived))

		// GET /api/users/:id/conversations - Retrieve a list of conversations a specific user is a participant in by ID
		v1.Get("/:id/conversations", utility.Use(controller.GetConversations))

		// GET /api/users/:id/groups - Retrieve a list of groups a specific user is a member of by ID
		v1.Get("/:id/groups", utility.Use(controller.GetGroups))

		// GET /api/users/:id/voice_notes_sent - Retrieve a list of voice notes sent by a specific user by ID
		v1.Get("/:id/voice_notes_sent", utility.Use(controller.GetVoiceNotesSent))

		// GET /api/users/:id/voice_notes_received - Retrieve a list of voice notes received by a specific user by ID
		v1.Get("/:id/voice_notes_received", utility.Use(controller.GetVoiceNotesReceived))

		// GET /api/users/:id/notifications - Retrieve a list of notifications for a specific user by ID
		v1.Get("/:id/notifications", utility.Use(controller.GetUserNotifications))

		// GET /api/users/:id/bookings - Retrieve a list of bookings made by a specific user by ID
		v1.Get("/:id/bookings", utility.Use(controller.GetUserBookings))

		// GET /api/users/:id/payments - Retrieve a list of payments made by a specific user by ID
		v1.Get("/:id/payments", utility.Use(controller.GetUserPayments))
	}
}

// api.post('/api/user', limiter(throttle.signup), use(userController.create));

// api.get('/api/user', auth.verify('user', 'user.read'), use(userController.get));

// api.get('/api/user/account', auth.verify('user'), use(userController.account));

// api.get(
// 	'/api/user/permissions',
// 	auth.verify('user', 'user.read'),
// 	use(userController.permissions)
// );

// api.patch(
// 	'/api/user',
// 	auth.verify('user', 'user.update'),
// 	use(userController.update)
// );

// api.put(
// 	'/api/user/password',
// 	auth.verify('user'),
// 	use(userController.password)
// );

// api.put('/api/user/2fa', auth.verify('user'), use(userController['2fa']));

// api.post(
// 	'/api/user/2fa/verify',
// 	auth.verify('user'),
// 	use(userController['2fa'].verify)
// );

// api.delete(
// 	'/api/user',
// 	auth.verify('user', 'user.delete'),
// 	use(userController.delete)
// );