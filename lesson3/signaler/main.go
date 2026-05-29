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

	signal := approval.SignalPayload{
		Approved: true,
		Reason:   "",
	}

	err = c.SignalWorkflow(context.Background(), "approval-workflow-1", "", approval.SignalName, signal)
	if err != nil {
		log.Fatalln("Unable to send signal:", err)
	}

	fmt.Println("Signal sent!")
}
