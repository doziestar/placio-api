package api

import (
	"fmt"
	"net/http"
	"net/url"
	_ "placio-api/docs/app"
	"placio-app/controller"
	"placio-app/service"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func JWTMiddleware() gin.HandlerFunc {
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

func InitializeRoutes(app *gin.Engine, db *gorm.DB) {
	fmt.Println("Initializing routes...")
	app.GET("/docs/*files", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//issuerURL, _ := url.Parse(os.Getenv("AUTH0_ISSUER_URL"))
	issuerURL, _ := url.Parse("https://dev-qlb0lv3d.us.auth0.com/")
	//audience := os.Getenv("AUTH0_AUDIENCE")

	provider := jwks.NewCachingProvider(issuerURL, time.Duration(5*time.Minute))

	jwtValidator, _ := validator.New(provider.KeyFunc,
		validator.RS256,
		//issuerURL.String(),
		"https://dev-qlb0lv3d.us.auth0.com/",
		[]string{"https://api.palnight.com"},
	)

	jwtMiddleware := jwtmiddleware.New(jwtValidator.ValidateToken)
	app.Use(adapter.Wrap(jwtMiddleware.CheckJWT))
	app.Use(JWTMiddleware())

	routerGroupV1 := app.Group("/api/v1")

	//// instances
	//var user models.User
	//var account models.Account

	// utility
	//newUtils := utility.NewUtility()
	//// auth
	//authService := service.NewAuthService(db, &models.User{})
	//authController := controller.NewAuthController(authService, newUtils)
	//authController.RegisterRoutes(routerGroupV1)
	//
	//// user
	//userService := service.NewUserService(db, &user, &account)
	//userController := controller.NewUserController(userService)
	//userController.RegisterRoutes(routerGroupV1)

	// settings
	//store := service.NewSettingsService(db)
	//settingsController := controller.NewSettingsController(store)
	//settingsController.RegisterRoutes(routerGroupV1, nil)
	//
	//// account
	//accountService := service.NewAccountService(db, account, user)
	//accountController := controller.NewAccountController(accountService, newUtils)
	//accountController.RegisterRoutes(routerGroupV1)

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

	// user
	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(routerGroupV1)

}
