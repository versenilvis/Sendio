// Package worker - Processor
// it will be called by the worker when a task is received (TODO: from redis)
package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/versenilvis/sendio/internal/database"
	"github.com/versenilvis/sendio/internal/logger"
	"github.com/versenilvis/sendio/internal/queue"
	"go.uber.org/zap"
)

type Processor struct {
	db *database.DB
}

func NewProcessor(db *database.DB) *Processor {
	return &Processor{
		db: db,
	}
}

// ProcessTask handles the email:send tasks
func (p *Processor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload queue.EmailSendPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("could not unmarshal payload: %w", err)
	}

	logger.Log.Info("Processing email task", zap.String("email_id", payload.EmailID))

	// TODO Step 1: Fetch email details from DB
	// TODO Step 2: Fetch SMTP config
	// TODO Step 3: Render template if needed
	// TODO Step 4: Send email using go-mail
	// TODO Step 5: Update email status in DB

	return nil
}
