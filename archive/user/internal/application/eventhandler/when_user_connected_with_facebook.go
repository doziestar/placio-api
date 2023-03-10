package eventhandler

import (
	"context"
	"time"

	"placio-api/cmd/user/internal/domain/user"
	"placio-api/cmd/user/internal/infrastructure/persistence"
	"placio-api/pkg/commandbus"
	"placio-api/pkg/domain"
	apperrors "placio-api/pkg/errors"
	"placio-api/pkg/eventbus"
	"placio-api/pkg/executioncontext"
)

// WhenUserConnectedWithFacebook handles event
func WhenUserConnectedWithFacebook(repository persistence.UserRepository, cb commandbus.CommandBus) eventbus.EventHandler {
	fn := func(parentCtx context.Context, event *domain.Event) error {
		ctx, cancel := context.WithTimeout(parentCtx, time.Second*120)
		defer cancel()

		e := event.Payload.(*user.ConnectedWithFacebook)

		if err := repository.UpdateFacebookID(ctx, e.ID.String(), e.FacebookID); err != nil {
			return apperrors.Wrap(err)
		}

		if executioncontext.Has(ctx, executioncontext.LIVE) {
			if err := cb.Publish(ctx, user.RequestAccessToken{
				ID:           e.ID,
				RedirectPath: e.RedirectPath,
			}); err != nil {
				return apperrors.Wrap(err)
			}
		}

		return nil
	}

	return fn
}