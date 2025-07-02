package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"queue-broker/internal/handler/http"
	"queue-broker/internal/infrastructure"
	"queue-broker/internal/usecase"
)

func main() {
	var (
		port           = flag.Int("port", 8080, "Port to listen on")
		defaultTimeout = flag.Int("timeout", 30, "Default timeout in seconds (0 to disable)")
		maxQueues      = flag.Int("max-queues", 100, "Maximum number of queues")
		maxMessages    = flag.Int("max-messages", 1000, "Maximum messages per queue")
	)
	flag.Parse()

	timeout := time.Duration(*defaultTimeout) * time.Second
	if *defaultTimeout == 0 {
		timeout = time.Hour * 24
	}

	repo := infrastructure.NewMemoryQueueRepository(*maxQueues, *maxMessages)
	service := usecase.NewQueueService(repo)
	queueHandler := handler.NewQueueHandler(service, timeout)

	http.Handle("/queue/", queueHandler)

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Queue broker listening on port %d\n", *port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
