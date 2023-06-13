package models

import (
	"time"
)

type Ticket struct {
	ID          int
	Title       string
	Description string
	Status      int
	Created     time.Time
	Resolved    time.Time
}
