package entities

import (
	"time"
)

type Problem struct {
  Id string
	Summary     string
	Description string
	CreatedBy   string
	CreatedAt   time.Time
}
