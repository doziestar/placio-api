package eventhandler

import (
	"context"
	"time"

	"placio-api/cmd/auth/internal/domain/token"
	"placio-api/cmd/auth/internal/infrastructure/persistence"
	"placio-pkg/domain"
	apperrors "placio-pkg/errors"
	"placio-pkg/eventbus"
)

// WhenTokenWasCreated handles event
func WhenTokenWasCreated(repository persistence.TokenRepository) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*token.WasCreated)

		if err := repository.Add(ctx, e); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
