package services

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"placio-api/cmd/auth/internal/application/config"
	"placio-api/cmd/auth/internal/domain/client"
	"placio-api/cmd/auth/internal/domain/token"
	authpersistence "placio-api/cmd/auth/internal/infrastructure/persistence"
	"placio-api/pkg/auth"
	"placio-api/pkg/commandbus"
	"placio-api/pkg/eventbus"
)

type containerFactory func(ctx context.Context, cfg *config.Config) (*ServiceContainer, error)

// NewServiceContainer creates new container
var NewServiceContainer containerFactory

type ServiceContainer struct {
	SQL   *sql.DB
	Mongo *mongo.Client

	CommandBus                  commandbus.CommandBus
	EventBus                    eventbus.EventBus
	AuthConn                    *grpc.ClientConn
	TokenRepository             token.Repository
	ClientRepository            client.Repository
	TokenPersistenceRepository  authpersistence.TokenRepository
	ClientPersistenceRepository authpersistence.ClientRepository
	Authenticator               auth.Authenticator
	OAuth2Manager               oauth2.Manager
	TokenAuthorizer             auth.TokenAuthorizer
}

func (c *ServiceContainer) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(3)

	var errs []error
	go func() {
		defer wg.Done()
		if c.SQL != nil {
			if err := c.SQL.Close(); err != nil {
				errs = append(errs, err)
			}
		}
	}()
	go func() {
		defer wg.Done()
		if c.Mongo != nil {
			if err := c.Mongo.Disconnect(ctx); err != nil {
				errs = append(errs, err)
			}
		}
	}()
	go func() {
		defer wg.Done()
		if c.AuthConn != nil {
			if err := c.AuthConn.Close(); err != nil {
				errs = append(errs, err)
			}
		}
	}()

	wg.Wait()

	var closeErr error
	for _, err := range errs {
		if closeErr == nil {
			closeErr = err
		} else {
			closeErr = fmt.Errorf("%v | %v", closeErr, err)
		}
	}

	if closeErr != nil {
		return closeErr
	}

	return ctx.Err()
}
