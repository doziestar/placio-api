package http

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/doziestar/tutis-api/cmd/auth/internal/application/config"
	"github.com/doziestar/tutis-api/cmd/auth/internal/infrastructure/persistence"
	"github.com/doziestar/tutis-api/cmd/auth/internal/interfaces/http/handlers"
	"github.com/doziestar/tutis-api/pkg/auth"
	"github.com/doziestar/tutis-api/pkg/commandbus"
	httpmiddleware "github.com/doziestar/tutis-api/pkg/http/middleware"
	httpauthenticator "github.com/doziestar/tutis-api/pkg/http/middleware/authenticator"
	"github.com/doziestar/tutis-api/pkg/http/response/json"
	"github.com/doziestar/tutis-api/pkg/identity"
	"github.com/vardius/gorouter/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"gopkg.in/oauth2.v4/server"
)

// NewRouter provides new router
func NewRouter(
	cfg *config.Config,
	tokenAuthorizer auth.TokenAuthorizer,
	server *server.Server,
	commandBus commandbus.CommandBus,
	sqlConn *sql.DB, mongoConn *mongo.Client,
	grpcConnectionMap map[string]*grpc.ClientConn,
	tokenRepository persistence.TokenRepository,
	clientRepository persistence.ClientRepository,
) http.Handler {
	authenticator := httpauthenticator.NewToken(tokenAuthorizer.Auth)

	// Global middleware
	router := gorouter.New(
		httpmiddleware.Recover(),
		httpmiddleware.WithMetadata(),
		httpmiddleware.Logger(),
		httpmiddleware.XSS(),
		httpmiddleware.HSTS(),
		authenticator.FromHeader("Restricted"),
		authenticator.FromQuery("authToken"),
		authenticator.FromCookie("at"),
		httpmiddleware.CORS(
			cfg.HTTP.Origins,
			cfg.App.Environment == "development",
		),
		httpmiddleware.LimitRequestBody(int64(10<<20)), // 10 MB is a lot of text.
		httpmiddleware.Metrics(),
		httpmiddleware.RateLimit(10, 10, 3*time.Minute), // 5 of requests per second with bursts of at most 10 requests
	)
	router.NotFound(json.NotFound())
	router.NotAllowed(json.NotAllowed())

	authorizeHandler := handlers.BuildAuthorizeHandler(server)
	router.GET("/authorize", authorizeHandler)
	router.POST("/authorize", authorizeHandler)
	router.POST("/token", handlers.BuildTokenHandler(server))

	router.POST("/dispatch/client/{command}", handlers.BuildClientCommandDispatchHandler(commandBus))
	router.POST("/dispatch/token/{command}", handlers.BuildTokenCommandDispatchHandler(commandBus))

	router.GET("/clients", handlers.BuildListClientsHandler(clientRepository))
	router.GET("/clients/{clientID}", handlers.BuildGetClientHandler(clientRepository))
	router.GET("/clients/{clientID}/tokens", handlers.BuildListTokensHandler(tokenRepository, clientRepository))
	router.GET("/users/{userID}/tokens", handlers.BuildListUserAuthTokensHandler(tokenRepository))

	// middleware applies to whole subtrees
	router.USE(http.MethodGet, "/users", httpmiddleware.GrantAccessFor(identity.PermissionTokenRead))
	router.USE(http.MethodGet, "/clients", httpmiddleware.GrantAccessFor(identity.PermissionClientRead))
	router.USE(http.MethodPost, "/dispatch", httpmiddleware.GrantAccessFor(identity.PermissionClientWrite))

	mainRouter := gorouter.New()
	mainRouter.NotFound(json.NotFound())
	mainRouter.NotAllowed(json.NotAllowed())

	// We do not want to apply middleware for this handlers
	// Liveness probes are to indicate that your application is running
	mainRouter.GET("/health", handlers.BuildLivenessHandler())
	// Readiness is meant to check if your application is ready to serve traffic
	mainRouter.GET("/readiness", handlers.BuildReadinessHandler(sqlConn, mongoConn, grpcConnectionMap))

	mainRouter.Mount("/v1", router)

	return mainRouter
}
