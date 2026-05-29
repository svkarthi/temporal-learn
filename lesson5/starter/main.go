package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	"temporal-learn/lesson5/reminder"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client:", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "reminder-workflow-1",
		TaskQueue: reminder.TaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, reminder.ReminderWorkflow, "Buy groceries", 10)
	if err != nil {
		log.Fatalln("Unable to start workflow:", err)
	}

	fmt.Printf("Workflow started: ID=%s RunID=%s\n", we.GetID(), we.GetRunID())
	fmt.Println("Waiting for reminder to fire in 10 seconds... (or run canceller to cancel)")

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Workflow failed:", err)
	}

	fmt.Printf("Result: %s\n", result)
}
