package greeting

import (
	"context"
	"errors"
	"fmt"

	"go.temporal.io/sdk/activity"
)

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	attempt := activity.GetInfo(ctx).Attempt
	logger := activity.GetLogger(ctx)
	logger.Info("ComposeGreeting attempt", "attempt", attempt)

	if attempt < 3 {
		return "", errors.New("simulated flaky failure")
	}

	return fmt.Sprintf("Hello, %s! Welcome to Temporal.", name), nil
}
