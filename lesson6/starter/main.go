package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	"temporal-learn/lesson6/caller"
	"temporal-learn/lesson6/greetingsvc"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client:", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "nexus-caller-workflow-1",
		TaskQueue: greetingsvc.CallerTaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, caller.CallerWorkflow, "Karthik")
	if err != nil {
		log.Fatalln("Unable to start workflow:", err)
	}

	fmt.Printf("Workflow started: ID=%s RunID=%s\n", we.GetID(), we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Workflow failed:", err)
	}

	fmt.Printf("Result: %s\n", result)
}
