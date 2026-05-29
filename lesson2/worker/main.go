package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal-learn/lesson2/greeting"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	w := worker.New(c, greeting.TaskQueue, worker.Options{})
	w.RegisterWorkflow(greeting.GreetingWorkflow)
	w.RegisterActivity(greeting.ComposeGreeting)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Worker failed:", err)
	}
}
