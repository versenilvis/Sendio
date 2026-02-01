// Package queue - Client
package queue

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/versenilvis/sendio/internal/logger"
	"go.uber.org/zap"
)

type QueueClient struct {
	client *asynq.Client
}

// NewQueueClient creates a new queue client
// it will connect to redis with the given redisAddr
func NewQueueClient(redisAddr string) *QueueClient {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	return &QueueClient{
		client: client,
	}
}
// Close will close the queue client
func (q *QueueClient) Close() error {
	return q.client.Close()
}

// EnqueueEmailSend puts an email sending job into the queue
// this function will call NewEmailSendTask to create a task
// then store it into redis with q.client.Enqueue(task)
func (q *QueueClient) EnqueueEmailSend(emailID string) error {
	task, err := NewEmailSendTask(emailID)
	if err != nil {
		return err
	}

	info, err := q.client.Enqueue(task)
	if err != nil {
		return fmt.Errorf("could not enqueue task: %w", err)
	}

	logger.Log.Debug("Enqueued email task",
		zap.String("task_id", info.ID),
		zap.String("queue", info.Queue),
		zap.String("email_id", emailID),
	)
	return nil
}
