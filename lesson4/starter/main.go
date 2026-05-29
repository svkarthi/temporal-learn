package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	"temporal-learn/lesson4/order"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client:", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "order-workflow-1",
		TaskQueue: order.TaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, order.OrderWorkflow, "ORD-001")
	if err != nil {
		log.Fatalln("Unable to start workflow:", err)
	}

	fmt.Printf("Workflow started: ID=%s RunID=%s\n", we.GetID(), we.GetRunID())
	fmt.Println("Workflow is RUNNING. Use the querier to check status, signaler to advance it.")
}
