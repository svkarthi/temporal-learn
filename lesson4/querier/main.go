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

	resp, err := c.QueryWorkflow(context.Background(), "order-workflow-1", "", order.QueryName)
	if err != nil {
		log.Fatalln("Unable to query workflow:", err)
	}

	var status string
	if err := resp.Get(&status); err != nil {
		log.Fatalln("Unable to decode query result:", err)
	}

	fmt.Printf("Current order status: %s\n", status)
}
