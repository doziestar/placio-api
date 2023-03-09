//go:build persistence_mysql
// +build persistence_mysql

package services

import (
	"context"

	authproto "github.com/doziestar/tutis-api/cmd/auth/proto"
	"github.com/doziestar/tutis-api/cmd/user/internal/application/config"
	persistence "github.com/doziestar/tutis-api/cmd/user/internal/infrastructure/persistence/mysql"
	"github.com/doziestar/tutis-api/cmd/user/internal/infrastructure/repository"
	"github.com/doziestar/tutis-api/pkg/auth"
	memorycommandbus "github.com/doziestar/tutis-api/pkg/commandbus/memory"
	apperrors "github.com/doziestar/tutis-api/pkg/errors"
	memoryeventbus "github.com/doziestar/tutis-api/pkg/eventbus/memory"
	mysqleventstore "github.com/doziestar/tutis-api/pkg/eventstore/mysql"
	grpcutils "github.com/doziestar/tutis-api/pkg/grpc"
	"github.com/doziestar/tutis-api/pkg/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	NewServiceContainer = newMYSQLServiceContainer
}

func newMYSQLServiceContainer(ctx context.Context, cfg *config.Config) (*ServiceContainer, error) {
	commandBus := memorycommandbus.New(cfg.CommandBus.QueueSize)
	sqlConn := mysql.NewConnection(
		ctx,
		mysql.ConnectionConfig{
			Host:            cfg.MYSQL.Host,
			Port:            cfg.MYSQL.Port,
			User:            cfg.MYSQL.User,
			Pass:            cfg.MYSQL.Pass,
			Database:        cfg.MYSQL.Database,
			ConnMaxLifetime: cfg.MYSQL.ConnMaxLifetime,
			MaxIdleConns:    cfg.MYSQL.MaxIdleConns,
			MaxOpenConns:    cfg.MYSQL.MaxOpenConns,
		},
	)
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
	eventStore, err := mysqleventstore.New(ctx, "user_events", sqlConn)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	eventBus := memoryeventbus.New(cfg.EventBus.QueueSize)
	userPersistenceRepository, err := persistence.NewUserRepository(ctx, sqlConn)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	userRepository := repository.NewUserRepository(eventStore, eventBus)
	grpAuthClient := authproto.NewAuthenticationServiceClient(grpcAuthConn)
	authenticator := auth.NewSecretAuthenticator([]byte(cfg.Auth.Secret))
	claimsProvider := auth.NewClaimsProvider(authenticator)
	tokenAuthorizer := auth.NewJWTTokenAuthorizer(grpAuthClient, claimsProvider, authenticator)

	return &ServiceContainer{
		SQL:                       sqlConn,
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
