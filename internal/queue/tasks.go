// Packgae queue - Task payload
// task will receive email id from api handler
// this is a data structure that will be stored in Redis, we just need to send the EmailID
// depends on it, worker will query the database for the email details
package queue

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const (
	TypeEmailSend = "email:send"
)

// EmailSendPayload is the data needed to send an email
type EmailSendPayload struct {
	EmailID string `json:"email_id"`
}

// NewEmailSendTask creates a new task for sending an email
func NewEmailSendTask(emailID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailSendPayload{EmailID: emailID})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal email task payload: %w", err)
	}

	// task options: we can specify things like max retries here
	return asynq.NewTask(TypeEmailSend, payload, asynq.MaxRetry(3)), nil
}
