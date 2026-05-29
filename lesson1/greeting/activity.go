package greeting

import (
	"context"
	"fmt"
)

func ComposeGreeting1(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}

func ComposeGreeting2(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Good to see you, %s!", name), nil
}

func ComposeGreeting3(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Welcome to Temporal, %s!", name), nil
}
