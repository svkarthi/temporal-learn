package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	"temporal-learn/lesson1/greeting"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow-1",
		TaskQueue: greeting.TaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, greeting.GreetingWorkflow, "Karthik")
	if err != nil {
		log.Fatalln("Unable to start workflow:", err)
	}

	fmt.Printf("Started workflow: ID=%s, RunID=%s\n", we.GetID(), we.GetRunID())

	var results []string
	err = we.Get(context.Background(), &results)
	if err != nil {
		log.Fatalln("Workflow failed:", err)
	}

	for i, r := range results {
		fmt.Printf("Greeting %d: %s\n", i+1, r)
	}
}
