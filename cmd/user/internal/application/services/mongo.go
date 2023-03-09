//go:build persistence_mongodb
// +build persistence_mongodb

package services

import (
	"context"
	"fmt"

	authproto "github.com/doziestar/tutis-api/cmd/auth/proto"
	"github.com/doziestar/tutis-api/cmd/user/internal/application/config"
	persistence "github.com/doziestar/tutis-api/cmd/user/internal/infrastructure/persistence/mongo"
	"github.com/doziestar/tutis-api/cmd/user/internal/infrastructure/repository"
	"github.com/doziestar/tutis-api/pkg/auth"
	memorycommandbus "github.com/doziestar/tutis-api/pkg/commandbus/memory"
	apperrors "github.com/doziestar/tutis-api/pkg/errors"
	memoryeventbus "github.com/doziestar/tutis-api/pkg/eventbus/memory"
	mongoeventstore "github.com/doziestar/tutis-api/pkg/eventstore/mongo"
	grpcutils "github.com/doziestar/tutis-api/pkg/grpc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	NewServiceContainer = newMongoServiceContainer
}

func newMongoServiceContainer(ctx context.Context, cfg *config.Config) (*ServiceContainer, error) {
	commandBus := memorycommandbus.New(cfg.CommandBus.QueueSize)
	mongoConnection, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.MongoDB.User, cfg.MongoDB.Pass, cfg.MongoDB.Host, cfg.MongoDB.Port),
	))
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	mongoDB := mongoConnection.Database(cfg.MongoDB.Database)
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
	eventStore, err := mongoeventstore.New(ctx, "events", mongoDB)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	eventBus := memoryeventbus.New(cfg.EventBus.QueueSize)
	userPersistenceRepository, err := persistence.NewUserRepository(ctx, mongoDB)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	userRepository := repository.NewUserRepository(eventStore, eventBus)
	grpAuthClient := authproto.NewAuthenticationServiceClient(grpcAuthConn)
	authenticator := auth.NewSecretAuthenticator([]byte(cfg.Auth.Secret))
	claimsProvider := auth.NewClaimsProvider(authenticator)
	tokenAuthorizer := auth.NewJWTTokenAuthorizer(grpAuthClient, claimsProvider, authenticator)

	return &ServiceContainer{
		Mongo:                     mongoConnection,
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
