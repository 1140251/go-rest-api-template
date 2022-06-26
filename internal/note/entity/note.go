package entity

import (
	"time"
)

// Note -.
type Note struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewNote(message string) *Note {
	return &Note{Message: message}
}
