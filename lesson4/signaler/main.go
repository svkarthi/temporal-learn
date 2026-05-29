package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.temporal.io/sdk/client"

	"temporal-learn/lesson4/order"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run lesson4/signaler/main.go <status>")
		fmt.Println("  status: processing | shipped | delivered")
		os.Exit(1)
	}

	status := os.Args[1]

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client:", err)
	}
	defer c.Close()

	err = c.SignalWorkflow(context.Background(), "order-workflow-1", "", order.SignalName, status)
	if err != nil {
		log.Fatalln("Unable to send signal:", err)
	}

	fmt.Printf("Signal sent: status → %s\n", status)
}
