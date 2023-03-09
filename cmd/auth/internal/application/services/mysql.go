//go:build persistence_mysql
// +build persistence_mysql

package services

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"placio-api/cmd/auth/internal/application/config"
	appoauth2 "placio-api/cmd/auth/internal/application/services/oauth2"
	persistence "placio-api/cmd/auth/internal/infrastructure/persistence/mysql"
	"placio-api/cmd/auth/internal/infrastructure/repository"
	authproto "placio-api/cmd/auth/proto"
	"placio-api/pkg/auth"
	memorycommandbus "placio-api/pkg/commandbus/memory"
	apperrors "placio-api/pkg/errors"
	memoryeventbus "placio-api/pkg/eventbus/memory"
	mysqleventstore "placio-api/pkg/eventstore/mysql"
	grpcutils "placio-api/pkg/grpc"
	"placio-api/pkg/mysql"
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
	grpcAuthConn := grpcutils.NewConnection(
		ctx,
		cfg.GRPC.Host,
		cfg.GRPC.Port,
		grpcutils.ConnectionConfig{
			ConnTime:    cfg.GRPC.ConnTime,
			ConnTimeout: cfg.GRPC.ConnTimeout,
		},
	)
	eventStore, err := mysqleventstore.New(ctx, "auth_events", sqlConn)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	eventBus := memoryeventbus.New(cfg.EventBus.QueueSize)
	tokenRepository := repository.NewTokenRepository(eventStore, eventBus)
	clientRepository := repository.NewClientRepository(eventStore, eventBus)
	tokenPersistenceRepository, err := persistence.NewTokenRepository(ctx, sqlConn)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	clientPersistenceRepository, err := persistence.NewClientRepository(ctx, cfg, sqlConn)
	if err != nil {
		return nil, apperrors.Wrap(err)
	}
	tokenStore := appoauth2.NewTokenStore(tokenPersistenceRepository, tokenRepository)
	authenticator := auth.NewSecretAuthenticator([]byte(cfg.App.Secret))
	grpAuthClient := authproto.NewAuthenticationServiceClient(grpcAuthConn)
	claimsProvider := auth.NewClaimsProvider(authenticator)
	manager := appoauth2.NewManager(tokenStore, clientPersistenceRepository, authenticator, clientPersistenceRepository)
	tokenAuthorizer := auth.NewJWTTokenAuthorizer(grpAuthClient, claimsProvider, authenticator)

	return &ServiceContainer{
		SQL:                         sqlConn,
		CommandBus:                  commandBus,
		EventBus:                    eventBus,
		Authenticator:               authenticator,
		OAuth2Manager:               manager,
		AuthConn:                    grpcAuthConn,
		TokenAuthorizer:             tokenAuthorizer,
		TokenRepository:             tokenRepository,
		ClientRepository:            clientRepository,
		TokenPersistenceRepository:  tokenPersistenceRepository,
		ClientPersistenceRepository: clientPersistenceRepository,
	}, nil
}
