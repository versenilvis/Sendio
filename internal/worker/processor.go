// Package worker handles async email processing tasks
package worker

import (
	"github.com/versenilvis/sendio/internal/renderer"
	"github.com/versenilvis/sendio/internal/repository"
)

// Processor handles email sending tasks from the queue
type Processor struct {
	repo          *repository.Registry
	renderer      *renderer.Renderer
	encryptionKey string
}

func NewProcessor(repo *repository.Registry, renderer *renderer.Renderer, encryptionKey string) *Processor {
	return &Processor{
		repo:          repo,
		renderer:      renderer,
		encryptionKey: encryptionKey,
	}
}
