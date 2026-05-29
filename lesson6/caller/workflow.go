package caller

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"temporal-learn/lesson6/greetingsvc"
)

func CallerWorkflow(ctx workflow.Context, name string) (string, error) {
	c := workflow.NewNexusClient(greetingsvc.EndpointName, greetingsvc.ServiceName)

	fut := c.ExecuteOperation(ctx, greetingsvc.ComposeOperation, name, workflow.NexusOperationOptions{
		ScheduleToCloseTimeout: 10 * time.Second,
	})

	var result string
	if err := fut.Get(ctx, &result); err != nil {
		return "", err
	}

	return result, nil
}
