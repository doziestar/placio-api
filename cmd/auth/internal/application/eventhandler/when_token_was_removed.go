package eventhandler

import (
	"context"
	"time"

	"github.com/doziestar/tutis-api/cmd/auth/internal/domain/token"
	"github.com/doziestar/tutis-api/cmd/auth/internal/infrastructure/persistence"
	"github.com/doziestar/tutis-api/pkg/domain"
	apperrors "github.com/doziestar/tutis-api/pkg/errors"
	"github.com/doziestar/tutis-api/pkg/eventbus"
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
