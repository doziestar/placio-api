package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	_ "placio-api/docs/app"
	"placio-app/domains/amenities"
	"placio-app/domains/booking"
	"placio-app/domains/business"
	"placio-app/domains/cache"
	"placio-app/domains/categories"
	"placio-app/domains/comments"
	"placio-app/domains/events_management"
	"placio-app/domains/faq"
	"placio-app/domains/feature_releases"
	"placio-app/domains/feedback"
	"placio-app/domains/follow"
	"placio-app/domains/inventory"
	"placio-app/domains/like"
	"placio-app/domains/media"
	"placio-app/domains/notifications"
	"placio-app/domains/order"
	"placio-app/domains/places"
	"placio-app/domains/posts"
	"placio-app/domains/recommendations"
	"placio-app/domains/reviews"
	"placio-app/domains/search"
	"placio-app/domains/smartGym"
	"placio-app/domains/smartMenu"
	"placio-app/domains/smartRoom"
	"placio-app/domains/users"
	"placio-app/domains/websites"
	"placio-app/ent"
	"placio-app/utility"
	"placio-pkg/kafka"
	"placio-pkg/middleware"

	"github.com/cloudinary/cloudinary-go/v2"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func requestBodyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		fmt.Println("Request Body:", string(body))

		c.Next()
	}
}

func InitializeRoutes(app *gin.Engine, client *ent.Client) {
	//app.Use(requestBodyLogger())
	routerGroupV1 := app.Group("/api/v1")
	routerGroupV1WithoutAuth := app.Group("/api/v1")
	{
		routerGroupV1.GET("/docs/*files", ginSwagger.WrapHandler(swaggerfiles.Handler))

		routerGroupV1.GET("/ready", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ready",
			})
		})

		//redisClient := utility.NewRedisClient(os.Getenv("REDIS_URL"), 0, utility.CacheDuration)
		redisClient := utility.NewRedisClient("redis://default:a3677c1a7b84402eb34efd55ad3cf059@golden-colt-33790.upstash.io:33790", 0)
		_ = redisClient.ConnectRedis()

		cld, _ := cloudinary.NewFromParams("placio", "312498583624125", "k4XSQwWuhi3Vy7QAw7Qn0mUaW0s")
		brokers := []string{"glad-ocelot-13748-eu2-kafka.upstash.io:9092"}
		topic := "post_created"
		username := "Z2xhZC1vY2Vsb3QtMTM3NDgkiJbJsYDFiX7WFPdq0E1rXMVgyy2z-P46ix43a8g"
		password := "MmI0ZmY0MTAtZTU1OS00MjQ0LTkyMmItYjM1MjdhNWY4OThl"

		producer := kafka.NewProducer(brokers, topic, username, password)

		searchService, _ := search.NewSearchService(client)
		searchController := search.NewSearchController(searchService)
		searchController.RegisterRoutes(routerGroupV1)

		routerGroupV1.Use(middleware.EnsureValidToken())
		routerGroupV1WithoutAuth.Use(middleware.EnsureValidTokenButAllowAccess())

		cacheService := cache.NewCacheService(client, *redisClient, searchService)

		// notifications
		notificationService := notifications.NewNotificationService(client)
		notificationController := notifications.NewNotificationController(notificationService)
		notificationController.RegisterRoutes(routerGroupV1)

		// user
		userService := users.NewUserService(client, redisClient, searchService)
		userController := users.NewUserController(userService, *redisClient)
		userController.RegisterRoutes(routerGroupV1)

		// media
		mediaService := media.NewMediaService(client, cld)
		mediaController := media.NewMediaController(mediaService)
		mediaController.RegisterRoutes(routerGroupV1WithoutAuth)

		// comments
		commentService := comments.NewCommentService(client, notificationService)
		commentController := comments.NewCommentController(commentService, userService)
		commentController.RegisterRoutes(routerGroupV1)

		// likes
		likeService := like.NewLikeService(client, redisClient)
		userPlacesLikesService := like.NewUserLikePlaceService(client, redisClient, searchService, cacheService)
		likeController := like.NewLikeController(likeService, userPlacesLikesService)
		likeController.RegisterRoutes(routerGroupV1)

		// follow
		followService := follow.NewFollowService(client, cacheService)
		followController := follow.NewFollowController(followService)
		followController.RegisterRoutes(routerGroupV1)

		// places
		placeService := places.NewPlaceService(client, searchService, userPlacesLikesService, *followService, *redisClient, mediaService)
		placeController := places.NewPlaceController(placeService, *redisClient)
		placeController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)

		// reservations
		reservationService := booking.NewReservationService(client)
		reservationController := booking.NewReservationController(reservationService)
		reservationController.RegisterRoutes(routerGroupV1)

		// room
		//roomService := inventory.NewRoomService(client)
		//roomController := inventory.NewRoomController(roomService)
		//roomController.RegisterRoutes(routerGroupV1)

		// menu
		menuService := inventory.NewMenuService(client)
		menuController := inventory.NewMenuController(menuService)
		menuController.RegisterRoutes(routerGroupV1)

		//booking
		bookingService := booking.NewBookingService(client)
		bookingController := booking.NewBookingController(bookingService)
		bookingController.RegisterRoutes(routerGroupV1)

		// feedback
		helpService := feedback.NewHelpService(client)
		helpController := feedback.NewHelpController(helpService)
		helpController.RegisterRoutes(routerGroupV1)

		// category
		categoryService := categories.NewCategoryService(client, mediaService)
		categoryController := categories.NewCategoryController(categoryService)
		categoryController.RegisterRoutes(routerGroupV1)

		// business
		businessService := business.NewBusinessAccountService(client, searchService, redisClient, placeService)
		businessController := business.NewBusinessAccountController(businessService, *redisClient)
		businessController.RegisterRoutes(routerGroupV1)

		// posts
		postService := posts.NewPostService(client, redisClient, mediaService, producer)
		postController := posts.NewPostController(postService, userService, businessService, mediaService)
		postController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)

		// events
		eventService := events_management.NewEventService(client, searchService)
		eventController := events_management.NewEventController(eventService, utility.NewUtility())
		eventController.RegisterRoutes(routerGroupV1)

		// amenities
		amenityService := amenities.NewAmenityService(client)
		amenityController := amenities.NewAmenityController(amenityService, redisClient)
		amenityController.RegisterRoutes(routerGroupV1)

		// faq
		faqService := faq.NewFAQService(client, redisClient)
		faqController := faq.NewFAQController(faqService)
		faqController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)

		// Review
		reviewService := reviews.NewReviewService(client, mediaService)
		reviewController := reviews.NewReviewController(reviewService, mediaService)
		reviewController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)

		// inventory
		inventoryService := inventory.NewInventoryService(client, cacheService, &mediaService, placeService, &businessService)
		inventoryController := inventory.NewInventoryController(inventoryService, *redisClient)
		inventoryController.RegisterRoutes(routerGroupV1)

		// Feature Releases
		featureReleaseService := feature_releases.NewFeatureReleaseService(client, *redisClient)
		featureReleaseController := feature_releases.NewFeatureReleaseController(featureReleaseService, *redisClient)
		featureReleaseController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)

		// smart menu
		smartMenuService := smartMenu.NewSmartMenuService(client, mediaService, cld)
		smartMenuController := smartMenu.NewSmartMenuController(smartMenuService)
		smartMenuController.RegisterRoutes(routerGroupV1WithoutAuth, routerGroupV1)

		// smart room
		smartRoomService := smartRoom.NewSmartRoomService(client, mediaService)
		smartRoomController := smartRoom.NewSmartRoomController(smartRoomService)
		smartRoomController.RegisterRoutes(routerGroupV1WithoutAuth, routerGroupV1)

		// recommendations
		recommendationService := recommendations.NewRecommendations(client, userService, placeService)
		recommendationController := recommendations.NewRecommendationController(*recommendationService)
		recommendationController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)

		// website
		websiteService := websites.NewWebsiteService(client, businessService, userService, mediaService)
		websiteController := websites.NewWebsiteController(websiteService)
		websiteController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)

		// order
		orderService := order.NewOrderServices(client)
		orderController := order.NewOrderController(orderService)
		orderController.RegisterRoutes(routerGroupV1)

		// fitness
		smartFitnessService := smartGym.NewSmartFitnessService(client)
		smartFitnessController := smartGym.NewSmartFitnessController(smartFitnessService)
		smartFitnessController.RegisterRoutes(routerGroupV1, routerGroupV1WithoutAuth)
	}

}
