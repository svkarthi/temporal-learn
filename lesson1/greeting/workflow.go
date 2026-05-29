package greeting

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) ([]string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	f1 := workflow.ExecuteActivity(ctx, ComposeGreeting1, name)
	f2 := workflow.ExecuteActivity(ctx, ComposeGreeting2, name)
	f3 := workflow.ExecuteActivity(ctx, ComposeGreeting3, name)

	var g1, g2, g3 string
	if err := f1.Get(ctx, &g1); err != nil {
		return nil, err
	}
	if err := f2.Get(ctx, &g2); err != nil {
		return nil, err
	}
	if err := f3.Get(ctx, &g3); err != nil {
		return nil, err
	}

	return []string{g1, g2, g3}, nil
}
