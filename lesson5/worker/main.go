package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal-learn/lesson5/reminder"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client:", err)
	}
	defer c.Close()

	w := worker.New(c, reminder.TaskQueue, worker.Options{})
	w.RegisterWorkflow(reminder.ReminderWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Worker failed:", err)
	}
}
