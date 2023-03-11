package eventhandler

import (
	"context"
	"time"

	"placio-api/cmd/user/internal/domain/user"
	"placio-api/cmd/user/internal/infrastructure/persistence"
	"placio-pkg/domain"
	apperrors "placio-pkg/errors"
	"placio-pkg/eventbus"
)

// WhenUserEmailAddressWasChanged handles event
func WhenUserEmailAddressWasChanged(repository persistence.UserRepository) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*user.EmailAddressWasChanged)

		if err := repository.UpdateEmail(ctx, e.ID.String(), string(e.Email)); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return fn
}
