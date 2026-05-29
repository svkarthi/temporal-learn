package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	"temporal-learn/lesson3/approval"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client:", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "approval-workflow-1",
		TaskQueue: approval.TaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, approval.ApprovalWorkflow, "REQ-001")
	if err != nil {
		log.Fatalln("Unable to start workflow:", err)
	}

	fmt.Printf("Workflow started: ID=%s RunID=%s\n", we.GetID(), we.GetRunID())
	fmt.Println("Workflow is now WAITING for a signal. Run the signaler next.")
}
