package memory

import (
	"context"
	"errors"
	"runtime"
	"testing"
	"time"

	apperrors "github.com/doziestar/tutis-api/pkg/errors"

	"github.com/doziestar/tutis-api/pkg/domain"
)

type commandMock struct{}

func (c *commandMock) GetName() string {
	return "command"
}

func TestNew(t *testing.T) {
	bus := New(runtime.NumCPU())

	if bus == nil {
		t.Fail()
	}
}

func TestSubscribePublish(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	bus := New(runtime.NumCPU())

	bus.Subscribe(ctx, "command", func(ctx context.Context, _ domain.Command) error {
		return nil
	})

	bus.Publish(ctx, &commandMock{})

	if err := bus.Publish(ctx, &commandMock{}); err != nil {
		t.Error(err)
	}
}

func TestUnsubscribe(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	bus := New(runtime.NumCPU())

	handler := func(ctx context.Context, _ domain.Command) error {
		t.Fail()

		return nil
	}

	bus.Subscribe(ctx, "command", handler)
	bus.Unsubscribe(ctx, "command")

	if err := bus.Publish(ctx, &commandMock{}); err != nil && !errors.Is(err, apperrors.ErrTimeout) {
		t.Error(err)
	}
}
