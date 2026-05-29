package greetingsvc

import (
	"context"
	"fmt"

	"github.com/nexus-rpc/sdk-go/nexus"
)

const (
	ServiceName      = "greeting-service"
	HandlerTaskQueue = "nexus-handler-task-queue"
	CallerTaskQueue  = "nexus-caller-task-queue"
	EndpointName     = "greeting-endpoint"
)

var ComposeOperation = nexus.NewSyncOperation("compose", func(ctx context.Context, name string, opts nexus.StartOperationOptions) (string, error) {
	return fmt.Sprintf("Hello from Nexus, %s!", name), nil
})

func NewService() *nexus.Service {
	s := nexus.NewService(ServiceName)
	s.Register(ComposeOperation)
	return s
}
