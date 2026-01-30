// Package utils - models
package utils

import (
	"io"
	"log"
)

func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("Failed to close resource: %v", err)
	}
}
