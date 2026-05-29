package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal-learn/lesson6/greetingsvc"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client:", err)
	}
	defer c.Close()

	w := worker.New(c, greetingsvc.HandlerTaskQueue, worker.Options{})
	w.RegisterNexusService(greetingsvc.NewService())

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Handler worker failed:", err)
	}
}
