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

	err = c.SignalWorkflow(context.Background(), "reminder-workflow-1", "", reminder.CancelSignal, nil)
	if err != nil {
		log.Fatalln("Unable to send cancel signal:", err)
	}

	fmt.Println("Cancel signal sent!")
}
