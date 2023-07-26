package api

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"net/http"
	_ "placio-api/docs/app"
	"placio-app/controller"
	"placio-app/ent"
	"placio-app/middleware"
	"placio-app/service"
	"placio-app/utility"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func JWTMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		validatedClaims, ok := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		if !ok || validatedClaims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		//
		//if ok && validatedClaims != nil {
		//	fmt.Println("Issuer:", validatedClaims.RegisteredClaims.Issuer)
		//	fmt.Println("Subject:", validatedClaims.RegisteredClaims.Subject)
		//	fmt.Println("Audience:", validatedClaims.RegisteredClaims.Audience)
		//	fmt.Println("Expiration Time:", validatedClaims.RegisteredClaims.Expiry)
		//	fmt.Println("Not Before Time:", validatedClaims.RegisteredClaims.NotBefore)
		//	fmt.Println("Issued At Time:", validatedClaims.RegisteredClaims.IssuedAt)
		//}

		c.Set("user", validatedClaims.RegisteredClaims.Subject)
		c.Next()
	}
}

func InitializeRoutes(app *gin.Engine, client *ent.Client) {
	routerGroupV1 := app.Group("/api/v1")
	routerGroupV1WithoutAuth := app.Group("/api/v1")
	{
		routerGroupV1.GET("/docs/*files", ginSwagger.WrapHandler(swaggerfiles.Handler))

		routerGroupV1.GET("/ready", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ready",
			})
		})

		routerGroupV1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "health",
			})
		})

		//redisClient := utility.NewRedisClient(os.Getenv("REDIS_URL"), 0, utility.CacheDuration)
		redisClient := utility.NewRedisClient("redis://default:a3677c1a7b84402eb34efd55ad3cf059@golden-colt-33790.upstash.io:33790", 0, utility.CacheDuration)
		_ = redisClient.ConnectRedis()

		cld, _ := cloudinary.NewFromParams("placio", "312498583624125", "k4XSQwWuhi3Vy7QAw7Qn0mUaW0s")

		searchService, _ := service.NewSearchService(client)
		searchController := controller.NewSearchController(searchService)
		searchController.RegisterRoutes(routerGroupV1)

		routerGroupV1.Use(middleware.EnsureValidToken(client))
		routerGroupV1WithoutAuth.Use(middleware.EnsureValidTokenButAllowAccess(client))

		// utility
		//newUtils := utility.NewUtility()

		// user
		userService := service.NewUserService(client, redisClient, searchService)
		userController := controller.NewUserController(userService)
		userController.RegisterRoutes(routerGroupV1)

		// media
		mediaService := service.NewMediaService(client, cld)
		mediaController := controller.NewMediaController(mediaService)
		mediaController.RegisterRoutes(routerGroupV1WithoutAuth)

		// comments
		commentService := service.NewCommentService(client)
		commentController := controller.NewCommentController(commentService, userService)
		commentController.RegisterRoutes(routerGroupV1)

		// likes
		likeService := service.NewLikeService(client, redisClient)
		userPlacesLikesService := service.NewUserLikePlaceService(client, redisClient)
		likeController := controller.NewLikeController(likeService, userPlacesLikesService)
		likeController.RegisterRoutes(routerGroupV1)

		// follow
		followService := service.NewFollowService(client)
		followController := controller.NewFollowController(followService)
		followController.RegisterRoutes(routerGroupV1)

		// places
		placeService := service.NewPlaceService(client, searchService, userPlacesLikesService, *followService)
		placeController := controller.NewPlaceController(placeService)
		placeController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)

		// reservations
		reservationService := service.NewReservationService(client)
		reservationController := controller.NewReservationController(reservationService)
		reservationController.RegisterRoutes(routerGroupV1)

		// room
		roomService := service.NewRoomService(client)
		roomController := controller.NewRoomController(roomService)
		roomController.RegisterRoutes(routerGroupV1)

		// menu
		menuService := service.NewMenuService(client)
		menuController := controller.NewMenuController(menuService)
		menuController.RegisterRoutes(routerGroupV1)

		//booking
		bookingService := service.NewBookingService(client)
		bookingController := controller.NewBookingController(bookingService)
		bookingController.RegisterRoutes(routerGroupV1)

		// feedback
		helpService := service.NewHelpService(client)
		helpController := controller.NewHelpController(helpService)
		helpController.RegisterRoutes(routerGroupV1)

		// category
		categoryService := service.NewCategoryService(client, mediaService)
		categoryController := controller.NewCategoryController(categoryService)
		categoryController.RegisterRoutes(routerGroupV1)

		// business
		businessService := service.NewBusinessAccountService(client, searchService, redisClient, placeService)
		businessController := controller.NewBusinessAccountController(businessService)
		businessController.RegisterRoutes(routerGroupV1)

		// posts
		postService := service.NewPostService(client, redisClient)
		postController := controller.NewPostController(postService, userService, businessService, mediaService)
		postController.RegisterRoutes(routerGroupV1)

		// events
		eventService := service.NewEventService(client, searchService)
		eventController := controller.NewEventController(eventService, utility.NewUtility())
		eventController.RegisterRoutes(routerGroupV1)

		// amenities
		amenityService := service.NewAmenityService(client)
		amenityController := controller.NewAmenityController(amenityService, redisClient)
		amenityController.RegisterRoutes(routerGroupV1)

		// faq
		faqService := service.NewFAQService(client, redisClient)
		faqController := controller.NewFAQController(faqService)
		faqController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)

		// Review
		reviewService := service.NewReviewService(client, mediaService)
		reviewController := controller.NewReviewController(reviewService, mediaService)
		reviewController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)

		// ratings
		//ratingService := service.NewRatingService(client)
		//ratingController := controller.NewRatingController(ratingService)
		//ratingController.RegisterRoutes(routerGroupV1)

		//// tickets
		//ticketService := service.NewTicketService(db)
		//ticketController := controller.NewTicketController(ticketService)
		//ticketController.RegisterRoutes(routerGroupV1)
		//
		//// attendee
		//attendeeService := service.NewAttendeeService(db)
		//attendeeController := controller.NewAttendeeController(attendeeService)
		//attendeeController.RegisterRoutes(routerGroupV1)
		//
		//// ticketOption
		//ticketOptionService := service.NewTicketOptionService(db)
		//ticketOptionController := controller.NewTicketOptionController(ticketOptionService)
		//ticketOptionController.RegisterRoutes(routerGroupV1)

	}

}
