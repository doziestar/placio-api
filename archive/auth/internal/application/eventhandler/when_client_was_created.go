package eventhandler

import (
	"context"
	"time"

	"placio-api/cmd/auth/internal/domain/client"
	"placio-api/cmd/auth/internal/infrastructure/persistence"
	"placio-pkg/domain"
	apperrors "placio-pkg/errors"
	"placio-pkg/eventbus"
)

// WhenClientWasCreated handles event
func WhenClientWasCreated(repository persistence.ClientRepository) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*client.WasCreated)

		if err := repository.Add(ctx, e); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
