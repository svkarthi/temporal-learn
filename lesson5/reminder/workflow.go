package reminder

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

const (
	TaskQueue    = "reminder-task-queue"
	CancelSignal = "cancel-signal"
)

func ReminderWorkflow(ctx workflow.Context, task string, delaySeconds int) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Reminder scheduled", "task", task, "delaySeconds", delaySeconds)

	timerCtx, cancelTimer := workflow.WithCancel(ctx)
	timer := workflow.NewTimer(timerCtx, time.Duration(delaySeconds)*time.Second)
	cancelCh := workflow.GetSignalChannel(ctx, CancelSignal)

	var result string

	selector := workflow.NewSelector(ctx)

	selector.AddFuture(timer, func(f workflow.Future) {
		result = "Reminder fired: " + task
		logger.Info("Timer fired", "task", task)
	})

	selector.AddReceive(cancelCh, func(c workflow.ReceiveChannel, more bool) {
		cancelTimer()
		result = "Reminder cancelled: " + task
		logger.Info("Reminder cancelled", "task", task)
	})

	selector.Select(ctx)

	return result, nil
}
