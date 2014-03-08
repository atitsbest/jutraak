package problems

import (
    "time"
)

type CreateNewProblemCommand struct {
    Summary     string
    Description string
    Tags        []string
    CreatedBy   string
    CreatedAt   time.Time
}
