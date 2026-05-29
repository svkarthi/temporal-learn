package order

import (
	"go.temporal.io/sdk/workflow"
)

const (
	TaskQueue      = "order-task-queue"
	SignalName     = "order-status-signal"
	QueryName      = "order-status-query"
)

func OrderWorkflow(ctx workflow.Context, orderID string) (string, error) {
	status := "pending"

	err := workflow.SetQueryHandler(ctx, QueryName, func() (string, error) {
		return status, nil
	})
	if err != nil {
		return "", err
	}

	statusCh := workflow.GetSignalChannel(ctx, SignalName)

	for status != "delivered" {
		statusCh.Receive(ctx, &status)
		workflow.GetLogger(ctx).Info("Order status updated", "orderID", orderID, "status", status)
	}

	return "Order " + orderID + " completed with status: " + status, nil
}
