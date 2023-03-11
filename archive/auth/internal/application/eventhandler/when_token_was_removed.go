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

// WhenTokenWasRemoved handles event
func WhenTokenWasRemoved(repository persistence.TokenRepository) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*token.WasRemoved)

		if err := repository.Delete(ctx, e.ID.String()); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
