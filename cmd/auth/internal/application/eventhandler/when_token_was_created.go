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
