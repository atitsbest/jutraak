package entities

import (
    "time"
)

type Comment struct {
    Text      string
    CreatedAt time.Time
    CreatedBy string
}

// CTR
func NewComment(text string, who string) *Comment {
    return &Comment{
        Text:      text,
        CreatedBy: who,
        CreatedAt: time.Now(),
    }
}
