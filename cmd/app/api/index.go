package api

import (
	"fmt"
	_ "placio-api/docs/app"
	"placio-app/controller"
	"placio-app/middleware"
	"placio-app/models"
	"placio-app/service"
	"placio-app/utility"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func InitializeRoutes(app *gin.Engine, db *gorm.DB) {
	fmt.Println("Initializing routes...")
	app.GET("/docs/*files", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Use(middleware.AuthorizeUser("user"))
	routerGroupV1 := app.Group("/api/v1")

	// instances
	var user models.User
	var account models.Account

	// utility
	newUtils := utility.NewUtility()

	// auth
	authService := service.NewAuthService(db, &models.User{})
	authController := controller.NewAuthController(authService, newUtils)
	authController.RegisterRoutes(routerGroupV1)

	// user
	userService := service.NewUserService(db, &user, &account)
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(routerGroupV1)

	// settings
	store := service.NewSettingsService(db)
	settingsController := controller.NewSettingsController(store)
	settingsController.RegisterRoutes(routerGroupV1, nil)

	// account
	accountService := service.NewAccountService(db, account, user)
	accountController := controller.NewAccountController(accountService, newUtils)
	accountController.RegisterRoutes(routerGroupV1)

	// posts
	postService := service.NewPostService(db)
	postController := controller.NewPostController(postService)
	postController.RegisterRoutes(routerGroupV1)

	// comments
	commentService := service.NewCommentService(db)
	commentController := controller.NewCommentController(commentService)
	commentController.RegisterRoutes(routerGroupV1)

	// media
	mediaService := service.NewMediaService(db)
	mediaController := controller.NewMediaController(mediaService)
	mediaController.RegisterRoutes(routerGroupV1)

	// likes
	likeService := service.NewLikeService(db)
	likeController := controller.NewLikeController(likeService)
	likeController.RegisterRoutes(routerGroupV1)

	// ratings
	ratingService := service.NewRatingService(db)
	ratingController := controller.NewRatingController(ratingService)
	ratingController.RegisterRoutes(routerGroupV1)

	// tickets
	ticketService := service.NewTicketService(db)
	ticketController := controller.NewTicketController(ticketService)
	ticketController.RegisterRoutes(routerGroupV1)

	// attendee
	attendeeService := service.NewAttendeeService(db)
	attendeeController := controller.NewAttendeeController(attendeeService)
	attendeeController.RegisterRoutes(routerGroupV1)

	// ticketOption
	ticketOptionService := service.NewTicketOptionService(db)
	ticketOptionController := controller.NewTicketOptionController(ticketOptionService)
	ticketOptionController.RegisterRoutes(routerGroupV1)

}
