package main

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 10,
		},
	)

	mux := asynq.NewServeMux()
	// TODO: register handlers
	// mux.HandleFunc("email:send", handleEmailSend)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}

func handleEmailSend(ctx context.Context, t *asynq.Task) error {
	// TODO: implement email sending logic
	return nil
}
