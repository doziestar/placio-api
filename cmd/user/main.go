package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/doziestar/tutis-api/cmd/user/internal/domain"
	"github.com/doziestar/tutis-api/pkg/grpc/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vardius/gocontainer"
	"google.golang.org/grpc"

	"github.com/doziestar/tutis-api/cmd/user/internal/application/config"
	"github.com/doziestar/tutis-api/cmd/user/internal/application/services"
	usergrpc "github.com/doziestar/tutis-api/cmd/user/internal/interfaces/grpc"
	userhttp "github.com/doziestar/tutis-api/cmd/user/internal/interfaces/http"
	userproto "github.com/doziestar/tutis-api/cmd/user/proto"
	"github.com/doziestar/tutis-api/pkg/application"
	"github.com/doziestar/tutis-api/pkg/buildinfo"
	grpcutils "github.com/doziestar/tutis-api/pkg/grpc"
	httputils "github.com/doziestar/tutis-api/pkg/http"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	gocontainer.GlobalContainer = nil // disable global container instance
}

func main() {
	buildinfo.PrintVersionOrContinue()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.FromEnv()
	fmt.Println("CONFIG:", cfg)

	container, err := services.NewServiceContainer(ctx, cfg)
	if err != nil {
		panic(fmt.Errorf("failed to create service container: %w", err))
	}
	defer container.Close()

	if err := domain.RegisterUserDomain(ctx, cfg, container); err != nil {
		panic(err)
	}

	grpcServer := grpcutils.NewServer(
		grpcutils.ServerConfig{
			ServerMinTime: cfg.GRPC.ServerMinTime,
			ServerTime:    cfg.GRPC.ServerTime,
			ServerTimeout: cfg.GRPC.ServerTimeout,
		},
		[]grpc.UnaryServerInterceptor{
			middleware.TransformUnaryOutgoingError(),
			middleware.CountIncomingUnaryRequests(),
			// 	firewall.GrantAccessForUnaryRequest(identity.RoleUser),
		},
		[]grpc.StreamServerInterceptor{
			middleware.TransformStreamOutgoingError(),
			middleware.CountIncomingStreamRequests(),
			// 	firewall.GrantAccessForStreamRequest(identity.RoleUser),
		},
	)

	router := userhttp.NewRouter(
		cfg,
		container.TokenAuthorizer,
		container.UserPersistenceRepository,
		container.CommandBus,
		container.SQL,
		container.Mongo,
		map[string]*grpc.ClientConn{
			"user": container.UserConn,
		},
	)

	grpcUserServer := usergrpc.NewServer(container.CommandBus, container.UserPersistenceRepository)
	userproto.RegisterUserServiceServer(grpcServer, grpcUserServer)

	app := application.New()

	app.AddAdapters(
		httputils.NewAdapter(
			&http.Server{
				Addr:         fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
				ReadTimeout:  cfg.HTTP.ReadTimeout,
				WriteTimeout: cfg.HTTP.WriteTimeout,
				IdleTimeout:  cfg.HTTP.IdleTimeout, // limits server-side the amount of time a Keep-Alive connection will be kept idle before being reused
				Handler:      router,
			},
		),
		grpcutils.NewAdapter(
			"user",
			fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port),
			grpcServer,
		),
	)

	if cfg.App.Environment == "development" {
		app.AddAdapters(
			application.NewDebugAdapter(
				fmt.Sprintf("%s:%d", cfg.Debug.Host, cfg.Debug.Port),
			),
		)
	}

	app.WithShutdownTimeout(cfg.App.ShutdownTimeout)
	app.Run(ctx)
}
