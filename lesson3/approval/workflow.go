package approval

import (
	"go.temporal.io/sdk/workflow"
)

const SignalName = "approval-signal"
const TaskQueue = "approval-task-queue"

type SignalPayload struct {
	Approved bool
	Reason   string
}

func ApprovalWorkflow(ctx workflow.Context, requestID string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Approval workflow started", "requestID", requestID)

	var signal SignalPayload
	receiveChan := workflow.GetSignalChannel(ctx, SignalName)
	receiveChan.Receive(ctx, &signal)

	if signal.Approved {
		return "Request " + requestID + " approved", nil
	}

	return "Request " + requestID + " rejected: " + signal.Reason, nil
}
