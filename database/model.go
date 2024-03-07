package database

import (
	"time"

	"github.com/google/uuid"
)

type Command struct {
	ID        uuid.UUID `json:"id" bson:"id"`
	Command   string    `json:"command" bson:"command"`
	Args      []string  `json:"args" bson:"args"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
