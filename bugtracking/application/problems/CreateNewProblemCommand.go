package problems

import (
  "time"
)

type CreateNewProblemCommand struct {
  Summary string
  Description string
  CreatedBy string
  CreatedAt time.Time
}
