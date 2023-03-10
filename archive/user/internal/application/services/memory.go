//go:build !persistence_mysql
// +build !persistence_mysql

package services

import (
	"context"

	authproto "placio-api/cmd/auth/proto"
	"placio-api/cmd/user/internal/application/config"
	persistence "placio-api/cmd/user/internal/infrastructure/persistence/memory"
	"placio-api/cmd/user/internal/infrastructure/repository"
	"placio-api/pkg/auth"
	memorycommandbus "placio-api/pkg/commandbus/memory"
	memoryeventbus "placio-api/pkg/eventbus/memory"
	memoryeventstore "placio-api/pkg/eventstore/memory"
	grpcutils "placio-api/pkg/grpc"
)

func init() {
	NewServiceContainer = newMemoryServiceContainer
}

func newMemoryServiceContainer(ctx context.Context, cfg *config.Config) (*ServiceContainer, error) {
	commandBus := memorycommandbus.New(cfg.CommandBus.QueueSize)
	grpcUserConn := grpcutils.NewConnection(
		ctx,
		cfg.GRPC.Host,
		cfg.GRPC.Port,
		grpcutils.ConnectionConfig{
			ConnTime:    cfg.GRPC.ConnTime,
			ConnTimeout: cfg.GRPC.ConnTimeout,
		},
	)
	grpcAuthConn := grpcutils.NewConnection(
		ctx,
		cfg.Auth.Host,
		cfg.GRPC.Port,
		grpcutils.ConnectionConfig{
			ConnTime:    cfg.GRPC.ConnTime,
			ConnTimeout: cfg.GRPC.ConnTimeout,
		},
	)
	eventStore := memoryeventstore.New()
	eventBus := memoryeventbus.New(cfg.EventBus.QueueSize)
	userPersistenceRepository := persistence.NewUserRepository()
	userRepository := repository.NewUserRepository(eventStore, eventBus)
	grpAuthClient := authproto.NewAuthenticationServiceClient(grpcAuthConn)
	authenticator := auth.NewSecretAuthenticator([]byte(cfg.Auth.Secret))
	claimsProvider := auth.NewClaimsProvider(authenticator)
	tokenAuthorizer := auth.NewJWTTokenAuthorizer(grpAuthClient, claimsProvider, authenticator)

	return &ServiceContainer{
		CommandBus:                commandBus,
		UserConn:                  grpcUserConn,
		AuthConn:                  grpcAuthConn,
		EventBus:                  eventBus,
		AuthClient:                grpAuthClient,
		TokenAuthorizer:           tokenAuthorizer,
		UserRepository:            userRepository,
		UserPersistenceRepository: userPersistenceRepository,
		Authenticator:             authenticator,
	}, nil
}
