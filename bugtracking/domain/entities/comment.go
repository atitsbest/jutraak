package entities

import (
    . "github.com/atitsbest/jutraak/bugtracking/domain/valueobjects"
    "time"
)

type Comment struct {
    Text        string
    CreatedAt   time.Time
    CreatedBy   string
    Attachments []*Attachment
}

// CTR
func NewComment(text string, who string, attachments []*Attachment) *Comment {
    return &Comment{
        Text:        text,
        CreatedBy:   who,
        CreatedAt:   time.Now(),
        Attachments: attachments,
    }
}
