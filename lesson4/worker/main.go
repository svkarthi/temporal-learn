package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal-learn/lesson4/order"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client:", err)
	}
	defer c.Close()

	w := worker.New(c, order.TaskQueue, worker.Options{})
	w.RegisterWorkflow(order.OrderWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Worker failed:", err)
	}
}
